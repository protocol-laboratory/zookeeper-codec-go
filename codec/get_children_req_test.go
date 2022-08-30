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

func TestDecodeGetChildrenReq(t *testing.T) {
	bytes := testHex2Bytes(t, "0000000100000008000000012f00")
	req, err := DecodeGetChildrenReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, req.TransactionId)
	assert.Equal(t, OP_GET_CHILDREN, req.OpCode)
	assert.Equal(t, "/", req.Path)
	assert.False(t, req.Watch)
}

func TestEncodeGetChildrenReq(t *testing.T) {
	req := &GetChildrenReq{
		TransactionId: 1,
		OpCode:        OP_GET_CHILDREN,
		Path:          "/",
		Watch:         false,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testHex2Bytes(t, "0000000100000008000000012f00"), bytes)
}
