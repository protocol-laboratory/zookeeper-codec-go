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

type ConnectReq struct {
	ProtocolVersion int
	LastZxidSeen    int64
	Timeout         int
	SessionId       int64
	Password        []byte
	ReadOnly        bool
}

func DecodeConnectReq(bytes []byte) (req *ConnectReq, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			req = nil
		}
	}()
	req = &ConnectReq{}
	idx := 0
	req.ProtocolVersion, idx = readProtocolVersion(bytes, idx)
	req.LastZxidSeen, idx = readLastZxidSeen(bytes, idx)
	req.Timeout, idx = readTimeout(bytes, idx)
	req.SessionId, idx = readSessionId(bytes, idx)
	req.Password, idx = readPassword(bytes, idx)
	req.ReadOnly, idx = readReadOnly(bytes, idx)
	return req, nil
}

func (c *ConnectReq) BytesLength(containLen bool) int {
	length := 0
	if containLen {
		length += LenLength
	}
	length += LenProtocolVersion + LenLastZxidSeen + LenTimeout + LenSessionId + BytesLen(c.Password) + LenReadonly
	return length
}

func (c *ConnectReq) Bytes(containLen bool) []byte {
	bytes := make([]byte, c.BytesLength(containLen))
	idx := 0
	if containLen {
		idx = putInt(bytes, idx, len(bytes)-4)
	}
	idx = putProtocolVersion(bytes, idx, c.ProtocolVersion)
	idx = putLastZxidSeen(bytes, idx, c.LastZxidSeen)
	idx = putTimeout(bytes, idx, c.Timeout)
	idx = putSessionId(bytes, idx, c.SessionId)
	idx = putPassword(bytes, idx, c.Password)
	idx = putBool(bytes, idx, c.ReadOnly)
	return bytes
}
