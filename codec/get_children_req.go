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

type GetChildrenReq struct {
	TransactionId int
	OpCode        OpCode
	Path          string
	Watch         bool
}

func DecodeGetChildrenReq(bytes []byte) (req *GetChildrenReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &GetChildrenReq{}
	idx := 0
	req.TransactionId, idx = readTransactionId(bytes, idx)
	req.OpCode, idx = readOpCode(bytes, idx)
	req.Path, idx = readPath(bytes, idx)
	req.Watch, idx = readWatch(bytes, idx)
	return req, nil
}

func (g *GetChildrenReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenTransactionId + LenOpCode + StrLen(g.Path) + LenWatch
	return length
}

func (g *GetChildrenReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, g.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putTransactionId(bytes, idx, g.TransactionId)
	idx = putOpCode(bytes, idx, OP_GET_CHILDREN)
	idx = putPath(bytes, idx, g.Path)
	idx = putWatch(bytes, idx, g.Watch)
	return bytes
}
