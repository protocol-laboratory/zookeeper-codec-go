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

type StatPersisted struct {
	CreatedZxId      int64
	LastModifiedZxId int64
	Created          int64
	LastModified     int64
	Version          int
	ChildVersion     int
	AclVersion       int
	EphemeralOwner   int64
	PeerZxId         int64
}

func DecodeStatPersisted(bytes []byte) (stat *StatPersisted, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			stat = nil
		}
	}()
	stat = &StatPersisted{}
	idx := 0
	stat.CreatedZxId, idx = readCreatedZxId(bytes, idx)
	stat.LastModifiedZxId, idx = readLastModifiedZxId(bytes, idx)
	stat.Created, idx = readCreated(bytes, idx)
	stat.LastModified, idx = readLastModified(bytes, idx)
	stat.Version, idx = readVersion(bytes, idx)
	stat.ChildVersion, idx = readChildVersion(bytes, idx)
	stat.AclVersion, idx = readAclVersion(bytes, idx)
	stat.EphemeralOwner, idx = readEphemeralOwner(bytes, idx)
	stat.PeerZxId, idx = readPeerZxId(bytes, idx)
	return stat, nil
}

func (s *StatPersisted) BytesLength() int {
	length := 0
	length += LenCreatedZxId + LenLastModifiedZxId
	length += LenCreated + LenLastModified + LenVersion + LenChildVersion + LenAclVersion + LenEphemeralOwner
	length += LenPeerZxId
	return length
}

func readStatPersisted(bytes []byte, idx int) (*StatPersisted, int) {
	stat, err := DecodeStatPersisted(bytes[idx:])
	if err != nil {
		panic(err)
	}
	return stat, idx + stat.BytesLength()
}
