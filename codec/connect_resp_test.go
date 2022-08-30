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

func TestDecodeConnectResp(t *testing.T) {
	bytes := testHex2Bytes(t, "00000000000075300100020abac20001000000109f9f123fc926a4013f35eca7ee76386100")
	resp, err := DecodeConnectResp(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 0, resp.ProtocolVersion)
	assert.Equal(t, 30_000, resp.Timeout)
	assert.Equal(t, int64(72059839144132609), resp.SessionId)
	assert.Len(t, resp.Password, 16)
}

func TestEncodeConnectResp(t *testing.T) {
	resp := &ConnectResp{
		ProtocolVersion: 0,
		Timeout:         30_000,
		SessionId:       72059839144132609,
		Password:        testHex2Bytes(t, "9f9f123fc926a4013f35eca7ee763861"),
		ReadOnly:        false,
	}
	bytes := resp.Bytes(false)
	assert.Equal(t, testHex2Bytes(t, "00000000000075300100020abac20001000000109f9f123fc926a4013f35eca7ee76386100"), bytes)
}
