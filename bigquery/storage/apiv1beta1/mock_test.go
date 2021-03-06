// Copyright 2018 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// AUTO-GENERATED CODE. DO NOT EDIT.

package storage

import (
	emptypb "github.com/golang/protobuf/ptypes/empty"
	storagepb "google.golang.org/genproto/googleapis/cloud/bigquery/storage/v1beta1"
)

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/api/option"
	status "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	gstatus "google.golang.org/grpc/status"
)

var _ = io.EOF
var _ = ptypes.MarshalAny
var _ status.Status

type mockBigQueryStorageServer struct {
	// Embed for forward compatibility.
	// Tests will keep working if more methods are added
	// in the future.
	storagepb.BigQueryStorageServer

	reqs []proto.Message

	// If set, all calls return this error.
	err error

	// responses to return if err == nil
	resps []proto.Message
}

func (s *mockBigQueryStorageServer) CreateReadSession(ctx context.Context, req *storagepb.CreateReadSessionRequest) (*storagepb.ReadSession, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*storagepb.ReadSession), nil
}

func (s *mockBigQueryStorageServer) ReadRows(req *storagepb.ReadRowsRequest, stream storagepb.BigQueryStorage_ReadRowsServer) error {
	md, _ := metadata.FromIncomingContext(stream.Context())
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return s.err
	}
	for _, v := range s.resps {
		if err := stream.Send(v.(*storagepb.ReadRowsResponse)); err != nil {
			return err
		}
	}
	return nil
}

func (s *mockBigQueryStorageServer) BatchCreateReadSessionStreams(ctx context.Context, req *storagepb.BatchCreateReadSessionStreamsRequest) (*storagepb.BatchCreateReadSessionStreamsResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*storagepb.BatchCreateReadSessionStreamsResponse), nil
}

func (s *mockBigQueryStorageServer) FinalizeStream(ctx context.Context, req *storagepb.FinalizeStreamRequest) (*emptypb.Empty, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*emptypb.Empty), nil
}

func (s *mockBigQueryStorageServer) SplitReadStream(ctx context.Context, req *storagepb.SplitReadStreamRequest) (*storagepb.SplitReadStreamResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	if xg := md["x-goog-api-client"]; len(xg) == 0 || !strings.Contains(xg[0], "gl-go/") {
		return nil, fmt.Errorf("x-goog-api-client = %v, expected gl-go key", xg)
	}
	s.reqs = append(s.reqs, req)
	if s.err != nil {
		return nil, s.err
	}
	return s.resps[0].(*storagepb.SplitReadStreamResponse), nil
}

// clientOpt is the option tests should use to connect to the test server.
// It is initialized by TestMain.
var clientOpt option.ClientOption

var (
	mockBigQueryStorage mockBigQueryStorageServer
)

func TestMain(m *testing.M) {
	flag.Parse()

	serv := grpc.NewServer()
	storagepb.RegisterBigQueryStorageServer(serv, &mockBigQueryStorage)

	lis, err := net.Listen("tcp", "localhost:0")
	if err != nil {
		log.Fatal(err)
	}
	go serv.Serve(lis)

	conn, err := grpc.Dial(lis.Addr().String(), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	clientOpt = option.WithGRPCConn(conn)

	os.Exit(m.Run())
}

func TestBigQueryStorageCreateReadSession(t *testing.T) {
	var name string = "name3373707"
	var expectedResponse = &storagepb.ReadSession{
		Name: name,
	}

	mockBigQueryStorage.err = nil
	mockBigQueryStorage.reqs = nil

	mockBigQueryStorage.resps = append(mockBigQueryStorage.resps[:0], expectedResponse)

	var tableReference *storagepb.TableReference = &storagepb.TableReference{}
	var parent string = "parent-995424086"
	var request = &storagepb.CreateReadSessionRequest{
		TableReference: tableReference,
		Parent:         parent,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateReadSession(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockBigQueryStorage.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestBigQueryStorageCreateReadSessionError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockBigQueryStorage.err = gstatus.Error(errCode, "test error")

	var tableReference *storagepb.TableReference = &storagepb.TableReference{}
	var parent string = "parent-995424086"
	var request = &storagepb.CreateReadSessionRequest{
		TableReference: tableReference,
		Parent:         parent,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.CreateReadSession(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestBigQueryStorageReadRows(t *testing.T) {
	var expectedResponse *storagepb.ReadRowsResponse = &storagepb.ReadRowsResponse{}

	mockBigQueryStorage.err = nil
	mockBigQueryStorage.reqs = nil

	mockBigQueryStorage.resps = append(mockBigQueryStorage.resps[:0], expectedResponse)

	var readPosition *storagepb.StreamPosition = &storagepb.StreamPosition{}
	var request = &storagepb.ReadRowsRequest{
		ReadPosition: readPosition,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	stream, err := c.ReadRows(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := stream.Recv()

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockBigQueryStorage.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestBigQueryStorageReadRowsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockBigQueryStorage.err = gstatus.Error(errCode, "test error")

	var readPosition *storagepb.StreamPosition = &storagepb.StreamPosition{}
	var request = &storagepb.ReadRowsRequest{
		ReadPosition: readPosition,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	stream, err := c.ReadRows(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := stream.Recv()

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestBigQueryStorageBatchCreateReadSessionStreams(t *testing.T) {
	var expectedResponse *storagepb.BatchCreateReadSessionStreamsResponse = &storagepb.BatchCreateReadSessionStreamsResponse{}

	mockBigQueryStorage.err = nil
	mockBigQueryStorage.reqs = nil

	mockBigQueryStorage.resps = append(mockBigQueryStorage.resps[:0], expectedResponse)

	var session *storagepb.ReadSession = &storagepb.ReadSession{}
	var requestedStreams int32 = 1017221410
	var request = &storagepb.BatchCreateReadSessionStreamsRequest{
		Session:          session,
		RequestedStreams: requestedStreams,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchCreateReadSessionStreams(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockBigQueryStorage.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestBigQueryStorageBatchCreateReadSessionStreamsError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockBigQueryStorage.err = gstatus.Error(errCode, "test error")

	var session *storagepb.ReadSession = &storagepb.ReadSession{}
	var requestedStreams int32 = 1017221410
	var request = &storagepb.BatchCreateReadSessionStreamsRequest{
		Session:          session,
		RequestedStreams: requestedStreams,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.BatchCreateReadSessionStreams(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
func TestBigQueryStorageFinalizeStream(t *testing.T) {
	var expectedResponse *emptypb.Empty = &emptypb.Empty{}

	mockBigQueryStorage.err = nil
	mockBigQueryStorage.reqs = nil

	mockBigQueryStorage.resps = append(mockBigQueryStorage.resps[:0], expectedResponse)

	var stream *storagepb.Stream = &storagepb.Stream{}
	var request = &storagepb.FinalizeStreamRequest{
		Stream: stream,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.FinalizeStream(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockBigQueryStorage.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

}

func TestBigQueryStorageFinalizeStreamError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockBigQueryStorage.err = gstatus.Error(errCode, "test error")

	var stream *storagepb.Stream = &storagepb.Stream{}
	var request = &storagepb.FinalizeStreamRequest{
		Stream: stream,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	err = c.FinalizeStream(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
}
func TestBigQueryStorageSplitReadStream(t *testing.T) {
	var expectedResponse *storagepb.SplitReadStreamResponse = &storagepb.SplitReadStreamResponse{}

	mockBigQueryStorage.err = nil
	mockBigQueryStorage.reqs = nil

	mockBigQueryStorage.resps = append(mockBigQueryStorage.resps[:0], expectedResponse)

	var originalStream *storagepb.Stream = &storagepb.Stream{}
	var request = &storagepb.SplitReadStreamRequest{
		OriginalStream: originalStream,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SplitReadStream(context.Background(), request)

	if err != nil {
		t.Fatal(err)
	}

	if want, got := request, mockBigQueryStorage.reqs[0]; !proto.Equal(want, got) {
		t.Errorf("wrong request %q, want %q", got, want)
	}

	if want, got := expectedResponse, resp; !proto.Equal(want, got) {
		t.Errorf("wrong response %q, want %q)", got, want)
	}
}

func TestBigQueryStorageSplitReadStreamError(t *testing.T) {
	errCode := codes.PermissionDenied
	mockBigQueryStorage.err = gstatus.Error(errCode, "test error")

	var originalStream *storagepb.Stream = &storagepb.Stream{}
	var request = &storagepb.SplitReadStreamRequest{
		OriginalStream: originalStream,
	}

	c, err := NewBigQueryStorageClient(context.Background(), clientOpt)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.SplitReadStream(context.Background(), request)

	if st, ok := gstatus.FromError(err); !ok {
		t.Errorf("got error %v, expected grpc error", err)
	} else if c := st.Code(); c != errCode {
		t.Errorf("got error code %q, want %q", c, errCode)
	}
	_ = resp
}
