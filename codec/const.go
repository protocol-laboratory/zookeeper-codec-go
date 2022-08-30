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

const (
	LenAclVersion       = 4
	LenArray            = 4
	LenChildVersion     = 4
	LenCreated          = LenTime
	LenCreatedZxId      = LenZxId
	LenDataLength       = 4
	LenEphemeralOwner   = 8
	LenError            = 4
	LenFlags            = 4
	LenLastModified     = LenTime
	LenLastModifiedZxId = LenZxId
	LenLastZxidSeen     = LenZxId
	LenLength           = 4
	LenNumberOfChildren = 4
	LenOpCode           = 4
	LenPermission       = 4
	LenProtocolVersion  = 4
	LenReadonly         = 1
	LenSessionId        = 8
	LenTime             = 8
	LenTimeout          = 4
	LenTransactionId    = 4
	LenVersion          = 4
	LenWatch            = 1
	LenZxId             = 8
)

var (
	PasswordEmpty = make([]byte, 16)
)
