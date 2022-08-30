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

func TestDecodeConnectReq(t *testing.T) {
	bytes := testHex2Bytes(t, "000000000000000000000000000075300000000000000000000000100000000000000000000000000000000000")
	req, err := DecodeConnectReq(bytes)
	assert.Nil(t, err)
	assert.Equal(t, 0, req.ProtocolVersion)
	assert.Equal(t, int64(0), req.LastZxidSeen)
	assert.Equal(t, 30_000, req.Timeout)
	assert.Equal(t, int64(0), req.SessionId)
	assert.Len(t, req.Password, 16)
}

func TestEncodeConnectReq(t *testing.T) {
	req := &ConnectReq{
		ProtocolVersion: 0,
		LastZxidSeen:    0,
		Timeout:         30_000,
		SessionId:       0,
		Password:        PasswordEmpty,
		ReadOnly:        false,
	}
	bytes := req.Bytes(false)
	assert.Equal(t, testHex2Bytes(t, "000000000000000000000000000075300000000000000000000000100000000000000000000000000000000000"), bytes)
}
