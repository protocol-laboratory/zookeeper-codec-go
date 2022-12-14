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

type OpCode int32

const (
	OP_ERROR          OpCode = -1
	OP_CREATE_SESSION OpCode = -10
	OP_CLOSE_SESSION  OpCode = -11
)

const (
	OP_NOTIFICATION OpCode = iota
	OP_CREATE
	OP_DELETE
	OP_EXISTS
	OP_GET_DATA
	OP_SET_DATA
	OP_GET_ACL
	OP_SET_ACL
	OP_GET_CHILDREN
	OP_SYNC
	OP_PING
	OP_GET_CHILDREN2
	OP_CHECK
	OP_MULTI
	OP_CREATE2
	OP_RECONFIG
	OP_CHECK_WATCHES
	OP_REMOVE_WATCHES
	OP_CREATE_CONTAINER
	OP_DELETE_CONTAINER
	OP_CREATE_TTL
	OP_MULTI_READ
)

const (
	OP_AUTH OpCode = iota + 100
	OP_SET_WATCHES
	OP_SASL
	OP_GET_EPHEMERALS
	OP_GET_ALL_CHILDREN_NUMBER
	OP_SET_WATCHES2
	OP_ADD_WATCH
	OP_WHO_AM_I
)
