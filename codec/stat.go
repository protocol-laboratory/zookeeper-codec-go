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

import "runtime/debug"

type Stat struct {
	CreatedZxId              int64
	LastModifiedZxId         int64
	Created                  int64
	LastModified             int64
	Version                  int
	ChildVersion             int
	AclVersion               int
	EphemeralOwner           int64
	DataLength               int
	NumChildren              int
	LastModifiedChildrenZxId int64
}

func DecodeStat(bytes []byte) (stat *Stat, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			stat = nil
		}
	}()
	stat = &Stat{}
	idx := 0
	stat.CreatedZxId, idx = readCreatedZxId(bytes, idx)
	stat.LastModifiedZxId, idx = readLastModifiedZxId(bytes, idx)
	stat.Created, idx = readCreated(bytes, idx)
	stat.LastModified, idx = readLastModified(bytes, idx)
	stat.Version, idx = readVersion(bytes, idx)
	stat.ChildVersion, idx = readChildVersion(bytes, idx)
	stat.AclVersion, idx = readAclVersion(bytes, idx)
	stat.EphemeralOwner, idx = readEphemeralOwner(bytes, idx)
	stat.DataLength, idx = readDataLength(bytes, idx)
	stat.NumChildren, idx = readNumChildren(bytes, idx)
	stat.LastModifiedChildrenZxId, idx = readLastModifiedChildrenZxId(bytes, idx)
	return stat, nil
}

func (s *Stat) BytesLength() int {
	length := 0
	length += LenCreatedZxId + LenLastModifiedZxId
	length += LenCreated + LenLastModified + LenVersion + LenChildVersion + LenAclVersion + LenEphemeralOwner
	length += LenDataLength + LenNumberOfChildren + LenLastModifiedZxId
	return length
}

func (s *Stat) Bytes() []byte {
	bytes := make([]byte, s.BytesLength())
	idx := 0
	idx = putCreatedZxId(bytes, idx, s.CreatedZxId)
	idx = putLastModifiedZxId(bytes, idx, s.LastModifiedZxId)
	idx = putCreated(bytes, idx, s.Created)
	idx = putLastModified(bytes, idx, s.LastModified)
	idx = putVersion(bytes, idx, s.Version)
	idx = putChildVersion(bytes, idx, s.ChildVersion)
	idx = putAclVersion(bytes, idx, s.AclVersion)
	idx = putEphemeralOwner(bytes, idx, s.EphemeralOwner)
	idx = putDataLength(bytes, idx, s.DataLength)
	idx = putNumChildren(bytes, idx, s.NumChildren)
	idx = putLastModifiedChildrenZxId(bytes, idx, s.LastModifiedChildrenZxId)
	return bytes
}

func readStat(bytes []byte, idx int) (*Stat, int) {
	stat, err := DecodeStat(bytes[idx:])
	if err != nil {
		panic(err)
	}
	return stat, idx + stat.BytesLength()
}

func putStat(bytes []byte, idx int, stat *Stat) int {
	idx = putBytesWithoutLen(bytes, idx, stat.Bytes())
	return idx
}
