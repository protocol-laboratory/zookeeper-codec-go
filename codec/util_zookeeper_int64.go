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

// This file is for zookeeper code int64 type. Format method as alpha order.

func readCreated(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putCreated(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readCreatedZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putCreatedZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readEphemeralOwner(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putEphemeralOwner(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readLastModified(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putLastModified(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readLastModifiedChildrenZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putLastModifiedChildrenZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readLastModifiedZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putLastModifiedZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readLastZxidSeen(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putLastZxidSeen(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readSessionId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putSessionId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}

func readZxId(bytes []byte, idx int) (int64, int) {
	return readInt64(bytes, idx)
}

func putZxId(bytes []byte, idx int, x int64) int {
	return putInt64(bytes, idx, x)
}
