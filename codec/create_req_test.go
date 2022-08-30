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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeCreateReq(t *testing.T) {
	bytes := testHex2Bytes(t, "0000000100000001000000082f7a6b2d746573740000000568656c6c6f000000010000001f00000005776f726c6400000006616e796f6e6500000000")
	req, err := DecodeCreateReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, req.TransactionId)
	assert.Equal(t, OP_CREATE, req.OpCode)
	assert.Equal(t, "/zk-test", req.Path)
	assert.Equal(t, []byte("hello"), req.Data)
	assert.Len(t, req.Permissions, 1)
	assert.Equal(t, 31, req.Permissions[0])
	assert.Equal(t, "world", req.Scheme)
	assert.Equal(t, "anyone", req.Credentials)
	assert.Equal(t, 0, req.Flags)
}

func TestEncodeCreateReq(t *testing.T) {
	req := &CreateReq{
		TransactionId: 1,
		OpCode:        OP_CREATE,
		Path:          "/zk-test",
		Data:          []byte("hello"),
		Permissions:   []int{31},
		Scheme:        "world",
		Credentials:   "anyone",
		Flags:         0,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testHex2Bytes(t, "0000000100000001000000082f7a6b2d746573740000000568656c6c6f000000010000001f00000005776f726c6400000006616e796f6e6500000000"), bytes)
}
