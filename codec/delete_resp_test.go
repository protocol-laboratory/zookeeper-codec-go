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

func TestDecodeDeleteResp(t *testing.T) {
	bytes := testHex2Bytes(t, "00000002000000000000002600000000")
	resp, err := DecodeDeleteResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 2, resp.TransactionId)
	assert.Equal(t, int64(38), resp.ZxId)
	assert.Equal(t, EC_OK, resp.Error)
}

func TestEncodeDeleteResp(t *testing.T) {
	resp := &DeleteResp{
		TransactionId: 2,
		ZxId:          38,
		Error:         EC_OK,
	}
	bytes := resp.Bytes()
	assert.Equal(t, testHex2Bytes(t, "00000002000000000000002600000000"), bytes)
}

func TestDecodeDeleteRespNoNodeExist(t *testing.T) {
	bytes := testHex2Bytes(t, "000000010000000000000025ffffff9b")
	resp, err := DecodeDeleteResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 1, resp.TransactionId)
	assert.Equal(t, int64(37), resp.ZxId)
	assert.Equal(t, EC_NoNodeError, resp.Error)
}

func TestEncodeDeleteRespNoNodeExist(t *testing.T) {
	resp := &DeleteResp{
		TransactionId: 1,
		ZxId:          37,
		Error:         EC_NoNodeError,
	}
	bytes := resp.Bytes()
	assert.Equal(t, testHex2Bytes(t, "000000010000000000000025ffffff9b"), bytes)
}
