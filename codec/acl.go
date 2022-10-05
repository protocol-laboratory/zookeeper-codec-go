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

type Acl struct {
	Perms int
	Id    *Id
}

func DecodeAcl(bytes []byte) (acl *Acl, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = PanicToError(r, debug.Stack())
			acl = nil
		}
	}()
	acl = &Acl{}
	idx := 0
	acl.Perms, idx = readInt(bytes, idx)
	acl.Id, idx = readId(bytes, idx)
	return acl, nil
}

func (a *Acl) ByteLength() int {
	return 4 + a.Id.ByteLength()
}

func readAcl(bytes []byte, idx int) (*Acl, int) {
	acl, err := DecodeAcl(bytes[idx:])
	if err != nil {
		panic(err)
	}
	return acl, idx + acl.ByteLength()
}
