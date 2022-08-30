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

func TestDecodeCloseResp(t *testing.T) {
	bytes := testHex2Bytes(t, "00000003000000000000000700000000")
	resp, err := DecodeCloseResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 3, resp.TransactionId)
	assert.Equal(t, int64(7), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
}

func TestEncodeCloseResp(t *testing.T) {
	resp := &CloseResp{
		TransactionId: 3,
		ZxId:          int64(7),
		Error:         0,
	}
	bytes := resp.Bytes(false)
	assert.Equal(t, testHex2Bytes(t, "00000003000000000000000700000000"), bytes)
}
