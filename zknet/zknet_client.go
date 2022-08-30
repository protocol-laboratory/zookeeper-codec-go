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

package zknet

import (
	"fmt"
	"github.com/protocol-laboratory/zookeeper-codec-go/codec"
	"net"
	"sync"
)

type ZookeeperNetClientConfig struct {
	Host             string
	Port             int
	BufferMax        int
	SendQueueSize    int
	PendingQueueSize int
}

func (z ZookeeperNetClientConfig) addr() string {
	return fmt.Sprintf("%s:%d", z.Host, z.Port)
}

type sendRequest struct {
	bytes    []byte
	callback func([]byte, error)
}

type ZookeeperNetClient struct {
	conn         net.Conn
	eventsChan   chan *sendRequest
	pendingQueue chan *sendRequest
	buffer       *buffer
	closeCh      chan struct{}
}

type buffer struct {
	max    int
	bytes  []byte
	cursor int
}

func (z *ZookeeperNetClient) Connect(req *codec.ConnectReq) (*codec.ConnectResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := codec.DecodeConnectResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperNetClient) Create(req *codec.CreateReq) (*codec.CreateResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := codec.DecodeCreateResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperNetClient) Delete(req *codec.DeleteReq) (*codec.DeleteResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := codec.DecodeDeleteResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperNetClient) Exists(req *codec.ExistsReq) (*codec.ExistsResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := codec.DecodeExistsResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperNetClient) GetData(req *codec.GetDataReq) (*codec.GetDataResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := codec.DecodeGetDataResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperNetClient) SetData(req *codec.SetDataReq) (*codec.SetDataResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := codec.DecodeSetDataResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperNetClient) GetChildren(req *codec.GetChildrenReq) (*codec.GetChildrenResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := codec.DecodeGetChildrenResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperNetClient) CloseSession(req *codec.CloseReq) (*codec.CloseResp, error) {
	bytes, err := z.Send(req.Bytes(true))
	if err != nil {
		return nil, err
	}
	resp, err := codec.DecodeCloseResp(bytes)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (z *ZookeeperNetClient) Send(bytes []byte) ([]byte, error) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	var result []byte
	var err error
	z.sendAsync(bytes, func(resp []byte, e error) {
		result = resp
		err = e
		wg.Done()
	})
	wg.Wait()
	return result[4:], err
}

func (z *ZookeeperNetClient) sendAsync(bytes []byte, callback func([]byte, error)) {
	sr := &sendRequest{
		bytes:    bytes,
		callback: callback,
	}
	z.eventsChan <- sr
}

func (z *ZookeeperNetClient) read() {
	for {
		select {
		case req := <-z.pendingQueue:
			n, err := z.conn.Read(z.buffer.bytes[z.buffer.cursor:])
			if err != nil {
				req.callback(nil, err)
				z.closeCh <- struct{}{}
				break
			}
			z.buffer.cursor += n
			if z.buffer.cursor < 4 {
				continue
			}
			length := int(z.buffer.bytes[3]) | int(z.buffer.bytes[2])<<8 | int(z.buffer.bytes[1])<<16 | int(z.buffer.bytes[0])<<24 + 4
			if z.buffer.cursor < length {
				continue
			}
			if length > z.buffer.max {
				req.callback(nil, fmt.Errorf("response length %d is too large", length))
				z.closeCh <- struct{}{}
				break
			}
			req.callback(z.buffer.bytes[:length], nil)
			z.buffer.cursor -= length
			copy(z.buffer.bytes[:z.buffer.cursor], z.buffer.bytes[length:])
		case <-z.closeCh:
			return
		}
	}
}

func (z *ZookeeperNetClient) write() {
	for {
		select {
		case req := <-z.eventsChan:
			n, err := z.conn.Write(req.bytes)
			if err != nil {
				req.callback(nil, err)
				z.closeCh <- struct{}{}
				break
			}
			if n != len(req.bytes) {
				req.callback(nil, fmt.Errorf("write %d bytes, but expect %d bytes", n, len(req.bytes)))
				z.closeCh <- struct{}{}
				break
			}
			z.pendingQueue <- req
		case <-z.closeCh:
			return
		}
	}
}

func (z *ZookeeperNetClient) Close() {
	_ = z.conn.Close()
	z.closeCh <- struct{}{}
}

func NewZkNetClient(config ZookeeperNetClientConfig) (*ZookeeperNetClient, error) {
	conn, err := net.Dial("tcp", config.addr())
	if err != nil {
		return nil, err
	}
	if config.SendQueueSize == 0 {
		config.SendQueueSize = 1000
	}
	if config.PendingQueueSize == 0 {
		config.PendingQueueSize = 1000
	}
	if config.BufferMax == 0 {
		config.BufferMax = 512 * 1024
	}
	z := &ZookeeperNetClient{}
	z.conn = conn
	z.eventsChan = make(chan *sendRequest, config.SendQueueSize)
	z.pendingQueue = make(chan *sendRequest, config.PendingQueueSize)
	z.buffer = &buffer{
		max:    config.BufferMax,
		bytes:  make([]byte, config.BufferMax),
		cursor: 0,
	}
	z.closeCh = make(chan struct{})
	go func() {
		z.read()
	}()
	go func() {
		z.write()
	}()
	return z, nil
}
