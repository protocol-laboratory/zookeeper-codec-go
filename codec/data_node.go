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

type DataNode struct {
	Path     string
	Data     []byte
	Acl      int64
	Stat     *StatPersisted
	Children []*DataNode
}

func readDataNode(bytes []byte, idx int) (dataNode *DataNode, nextIdx int) {
	dataNode = &DataNode{}
	dataNode.Data, idx = readBytes(bytes, idx)
	dataNode.Acl, idx = readInt64(bytes, idx)
	dataNode.Stat, idx = readStatPersisted(bytes, idx)
	return dataNode, idx
}
