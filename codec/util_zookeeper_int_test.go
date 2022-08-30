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

func TestDecodeProtocolVersion(t *testing.T) {
	bytes := testHex2Bytes(t, "00000000")
	protocolVersion, idx := readProtocolVersion(bytes, 0)
	assert.Equal(t, 4, idx)
	assert.Equal(t, 0, protocolVersion)
}

func TestEncodeProtocolVersion(t *testing.T) {
	bytes := make([]byte, 4)
	putProtocolVersion(bytes, 0, 0)
	assert.Equal(t, testHex2Bytes(t, "00000000"), bytes)
}
