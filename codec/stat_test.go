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

func TestDecodeStat(t *testing.T) {
	bytes := testHex2Bytes(t, "0000000000000005000000000000000600000182ec1b539f00000182ec1b53ac000000010000000000000000000000000000000000000005000000000000000000000005")
	stat, err := DecodeStat(bytes)
	assert.Nil(t, err)
	assert.Equal(t, int64(5), stat.CreatedZxId)
	assert.Equal(t, int64(6), stat.LastModifiedZxId)
	assert.Equal(t, int64(1661818590111), stat.Created)
	assert.Equal(t, int64(1661818590124), stat.LastModified)
	assert.Equal(t, 1, stat.Version)
	assert.Equal(t, 0, stat.ChildVersion)
	assert.Equal(t, 0, stat.AclVersion)
	assert.Equal(t, int64(0), stat.EphemeralOwner)
	assert.Equal(t, 5, stat.DataLength)
	assert.Equal(t, 0, stat.NumChildren)
	assert.Equal(t, int64(5), stat.LastModifiedChildrenZxId)
}

func TestEncodeStat(t *testing.T) {
	resp := &Stat{
		CreatedZxId:              5,
		LastModifiedZxId:         6,
		Created:                  1661818590111,
		LastModified:             1661818590124,
		Version:                  1,
		ChildVersion:             0,
		AclVersion:               0,
		EphemeralOwner:           0,
		DataLength:               5,
		NumChildren:              0,
		LastModifiedChildrenZxId: 5,
	}
	bytes := resp.Bytes()
	assert.Equal(t, testHex2Bytes(t, "0000000000000005000000000000000600000182ec1b539f00000182ec1b53ac000000010000000000000000000000000000000000000005000000000000000000000005"), bytes)
}
