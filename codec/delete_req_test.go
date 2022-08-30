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

func TestDecodeDeleteReq(t *testing.T) {
	bytes := testHex2Bytes(t, "00000001000000020000000c2f7a6b2d6e6f74666f756e6400000000")
	req, err := DecodeDeleteReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, req.TransactionId)
	assert.Equal(t, OP_DELETE, req.OpCode)
	assert.Equal(t, "/zk-notfound", req.Path)
	assert.Equal(t, 0, req.Version)
}

func TestEncodeDeleteReq(t *testing.T) {
	req := &DeleteReq{
		TransactionId: 1,
		OpCode:        OP_DELETE,
		Path:          "/zk-notfound",
		Version:       0,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testHex2Bytes(t, "00000001000000020000000c2f7a6b2d6e6f74666f756e6400000000"), bytes)
}
