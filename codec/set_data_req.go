// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package codec

import "runtime/debug"

type SetDataReq struct {
	TransactionId int
	OpCode        OpCode
	Path          string
	Data          []byte
	Version       int
}

func DecodeSetDataReq(bytes []byte) (req *SetDataReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &SetDataReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	req.Path, idx = readPath(bytes, idx)
	req.Data, idx = readData(bytes, idx)
	req.Version, idx = readVersion(bytes, idx)
	return req, nil
}

func (s *SetDataReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode + StrLen(s.Path) + BytesLen(s.Data) + LenVersion
	return length
}

func (s *SetDataReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, s.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, s.TransactionId)
	idx = putOpCode(bytes, idx, OP_SET_DATA)
	idx = putPath(bytes, idx, s.Path)
	idx = putData(bytes, idx, s.Data)
	idx = putVersion(bytes, idx, s.Version)
	return bytes
}
