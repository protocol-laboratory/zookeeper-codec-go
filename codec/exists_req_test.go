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

func TestDecodeExistsReq(t *testing.T) {
	bytes := testHex2Bytes(t, "00000001000000030000000b2f65786973742d7465737400")
	req, err := DecodeExistsReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, req.TransactionId)
	assert.Equal(t, OP_EXISTS, req.OpCode)
	assert.Equal(t, "/exist-test", req.Path)
	assert.False(t, req.Watch)
}

func TestEncodeExistsReq(t *testing.T) {
	req := &ExistsReq{
		TransactionId: 1,
		OpCode:        OP_EXISTS,
		Path:          "/exist-test",
		Watch:         false,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testHex2Bytes(t, "00000001000000030000000b2f65786973742d7465737400"), bytes)
}
