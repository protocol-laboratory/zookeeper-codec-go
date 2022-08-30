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

// This file is for zookeeper code int type. Format method as alpha order.

func readAclVersion(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putAclVersion(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readChildVersion(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putChildVersion(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readDataLength(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putDataLength(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readError(bytes []byte, idx int) (ErrorCode, int) {
	ec, i := readInt(bytes, idx)
	return ErrorCode(ec), i
}

func putError(bytes []byte, idx int, x ErrorCode) int {
	return putInt(bytes, idx, int(x))
}

func readFlags(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putFlags(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readNumChildren(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putNumChildren(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readPermission(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putPermission(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readProtocolVersion(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putProtocolVersion(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readVersion(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putVersion(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readTimeout(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putTimeout(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}

func readTransactionId(bytes []byte, idx int) (int, int) {
	return readInt(bytes, idx)
}

func putTransactionId(bytes []byte, idx int, x int) int {
	return putInt(bytes, idx, x)
}
