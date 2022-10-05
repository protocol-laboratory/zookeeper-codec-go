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
	"os"
	"testing"
)

// TestDecodeSnapshot zookeeper without traffic
func TestDecodeSnapshot0(t *testing.T) {
	bytes, err := os.ReadFile("snapshot.0")
	assert.Nil(t, err)
	s, err := DecodeSnapshot(bytes)
	assert.Nil(t, err)
	assert.Len(t, s.AclMap, 1)
}
