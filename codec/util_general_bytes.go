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

func readBytes(bytes []byte, idx int) ([]byte, int) {
	length, idx := readInt(bytes, idx)
	return bytes[idx : idx+length], idx + length
}

func putBytes(bytes []byte, idx int, srcBytes []byte) int {
	idx = putInt(bytes, idx, len(srcBytes))
	copy(bytes[idx:], srcBytes)
	return idx + len(srcBytes)
}

func putBytesWithoutLen(bytes []byte, idx int, srcBytes []byte) int {
	copy(bytes[idx:], srcBytes)
	return idx + len(srcBytes)
}

func BytesLen(bytes []byte) int {
	return 4 + len(bytes)
}
