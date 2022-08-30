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

type CreateResp struct {
	TransactionId int
	ZxId          int64
	Error         ErrorCode
	Path          string
}

func DecodeCreateResp(bytes []byte) (resp *CreateResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			resp = nil
		}
	}()
	resp = &CreateResp{}
	idx := 0
	resp.TransactionId, idx = readTransactionId(bytes, idx)
	resp.ZxId, idx = readZxId(bytes, idx)
	resp.Error, idx = readError(bytes, idx)
	if resp.Error == EC_OK {
		resp.Path, idx = readPath(bytes, idx)
	}
	return resp, nil
}

func (c *CreateResp) BytesLength() int {
	length := 0
	length += LenTransactionId + LenZxId + LenError
	if c.Error == EC_OK {
		length += StrLen(c.Path)
	}
	return length
}

func (c *CreateResp) Bytes() []byte {
	bytes := make([]byte, c.BytesLength())
	idx := 0
	idx = putTransactionId(bytes, idx, c.TransactionId)
	idx = putZxId(bytes, idx, c.ZxId)
	idx = putError(bytes, idx, c.Error)
	if c.Error == EC_OK {
		idx = putPath(bytes, idx, c.Path)
	}
	return bytes
}
