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

// This file is for zookeeper code bytes type. Format method as alpha order.

func readData(bytes []byte, idx int) ([]byte, int) {
	return readBytes(bytes, idx)
}

func putData(bytes []byte, idx int, dstBytes []byte) int {
	return putBytes(bytes, idx, dstBytes)
}

func readPassword(bytes []byte, idx int) ([]byte, int) {
	return readBytes(bytes, idx)
}

func putPassword(bytes []byte, idx int, dstBytes []byte) int {
	return putBytes(bytes, idx, dstBytes)
}
