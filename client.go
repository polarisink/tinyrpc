// Copyright 2022 <mzh.scnu@qq.com>. All rights reserved.
// Use of this source code is governed by a git
// license that can be found in the LICENSE file.

package tinyrpc

import (
	"io"
	"net/rpc"

	"github.com/zehuamama/tinyrpc/codec"
	"github.com/zehuamama/tinyrpc/compressor"
)

// Client rpc client based on net/rpc implementation
type Client struct {
	*rpc.Client
}

//Option provides options for rpc client
type Option func(o *options)

type options struct {
	compressType compressor.CompressType
}

// WithCompress set client compression format
func WithCompress(c compressor.CompressType) Option {
	return func(o *options) {
		o.compressType = c
	}
}

// NewClient Create a new rpc client
func NewClient(conn io.ReadWriteCloser, opts ...Option) *Client {
	options := options{
		compressType: compressor.Raw,
	}
	for _, option := range opts {
		option(&options)
	}
	return &Client{rpc.NewClientWithCodec(
		codec.NewClientCodec(conn, options.compressType))}
}

// Call synchronously calls the rpc function
func (c *Client) Call(serviceMethod string, args interface{}, reply interface{}) error {
	return c.Client.Call(serviceMethod, args, reply)
}

// AsyncCall asynchronously calls the rpc function and returns a channel of *rpc.Call
func (c *Client) AsyncCall(serviceMethod string, args interface{}, reply interface{}) chan *rpc.Call {
	return c.Go(serviceMethod, args, reply, nil).Done
}
