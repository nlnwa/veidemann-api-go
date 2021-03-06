// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package frontier

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// FrontierClient is the client API for Frontier service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FrontierClient interface {
	// Start crawling seed
	CrawlSeed(ctx context.Context, in *CrawlSeedRequest, opts ...grpc.CallOption) (*CrawlExecutionId, error)
	// Request a URI from the Frontiers queue
	// Used by a Harvester to fetch a new page
	GetNextPage(ctx context.Context, opts ...grpc.CallOption) (Frontier_GetNextPageClient, error)
	// The number of busy CrawlHostGroups which essentially is the number of web pages currently downloading
	BusyCrawlHostGroupCount(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountResponse, error)
	// Total number of queued URI's
	QueueCountTotal(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountResponse, error)
	// Number of queued URI's for a CrawlExecution
	QueueCountForCrawlExecution(ctx context.Context, in *CrawlExecutionId, opts ...grpc.CallOption) (*CountResponse, error)
	// Number of queued URI's for a CrawlHostGroup
	QueueCountForCrawlHostGroup(ctx context.Context, in *CrawlHostGroup, opts ...grpc.CallOption) (*CountResponse, error)
}

type frontierClient struct {
	cc grpc.ClientConnInterface
}

func NewFrontierClient(cc grpc.ClientConnInterface) FrontierClient {
	return &frontierClient{cc}
}

func (c *frontierClient) CrawlSeed(ctx context.Context, in *CrawlSeedRequest, opts ...grpc.CallOption) (*CrawlExecutionId, error) {
	out := new(CrawlExecutionId)
	err := c.cc.Invoke(ctx, "/veidemann.api.frontier.v1.Frontier/CrawlSeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontierClient) GetNextPage(ctx context.Context, opts ...grpc.CallOption) (Frontier_GetNextPageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Frontier_serviceDesc.Streams[0], "/veidemann.api.frontier.v1.Frontier/GetNextPage", opts...)
	if err != nil {
		return nil, err
	}
	x := &frontierGetNextPageClient{stream}
	return x, nil
}

type Frontier_GetNextPageClient interface {
	Send(*PageHarvest) error
	Recv() (*PageHarvestSpec, error)
	grpc.ClientStream
}

type frontierGetNextPageClient struct {
	grpc.ClientStream
}

func (x *frontierGetNextPageClient) Send(m *PageHarvest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *frontierGetNextPageClient) Recv() (*PageHarvestSpec, error) {
	m := new(PageHarvestSpec)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *frontierClient) BusyCrawlHostGroupCount(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.frontier.v1.Frontier/BusyCrawlHostGroupCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontierClient) QueueCountTotal(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.frontier.v1.Frontier/QueueCountTotal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontierClient) QueueCountForCrawlExecution(ctx context.Context, in *CrawlExecutionId, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.frontier.v1.Frontier/QueueCountForCrawlExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *frontierClient) QueueCountForCrawlHostGroup(ctx context.Context, in *CrawlHostGroup, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/veidemann.api.frontier.v1.Frontier/QueueCountForCrawlHostGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FrontierServer is the server API for Frontier service.
// All implementations must embed UnimplementedFrontierServer
// for forward compatibility
type FrontierServer interface {
	// Start crawling seed
	CrawlSeed(context.Context, *CrawlSeedRequest) (*CrawlExecutionId, error)
	// Request a URI from the Frontiers queue
	// Used by a Harvester to fetch a new page
	GetNextPage(Frontier_GetNextPageServer) error
	// The number of busy CrawlHostGroups which essentially is the number of web pages currently downloading
	BusyCrawlHostGroupCount(context.Context, *empty.Empty) (*CountResponse, error)
	// Total number of queued URI's
	QueueCountTotal(context.Context, *empty.Empty) (*CountResponse, error)
	// Number of queued URI's for a CrawlExecution
	QueueCountForCrawlExecution(context.Context, *CrawlExecutionId) (*CountResponse, error)
	// Number of queued URI's for a CrawlHostGroup
	QueueCountForCrawlHostGroup(context.Context, *CrawlHostGroup) (*CountResponse, error)
	mustEmbedUnimplementedFrontierServer()
}

// UnimplementedFrontierServer must be embedded to have forward compatible implementations.
type UnimplementedFrontierServer struct {
}

func (UnimplementedFrontierServer) CrawlSeed(context.Context, *CrawlSeedRequest) (*CrawlExecutionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CrawlSeed not implemented")
}
func (UnimplementedFrontierServer) GetNextPage(Frontier_GetNextPageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetNextPage not implemented")
}
func (UnimplementedFrontierServer) BusyCrawlHostGroupCount(context.Context, *empty.Empty) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BusyCrawlHostGroupCount not implemented")
}
func (UnimplementedFrontierServer) QueueCountTotal(context.Context, *empty.Empty) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueCountTotal not implemented")
}
func (UnimplementedFrontierServer) QueueCountForCrawlExecution(context.Context, *CrawlExecutionId) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueCountForCrawlExecution not implemented")
}
func (UnimplementedFrontierServer) QueueCountForCrawlHostGroup(context.Context, *CrawlHostGroup) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueueCountForCrawlHostGroup not implemented")
}
func (UnimplementedFrontierServer) mustEmbedUnimplementedFrontierServer() {}

// UnsafeFrontierServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FrontierServer will
// result in compilation errors.
type UnsafeFrontierServer interface {
	mustEmbedUnimplementedFrontierServer()
}

func RegisterFrontierServer(s grpc.ServiceRegistrar, srv FrontierServer) {
	s.RegisterService(&_Frontier_serviceDesc, srv)
}

func _Frontier_CrawlSeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrawlSeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontierServer).CrawlSeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.frontier.v1.Frontier/CrawlSeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontierServer).CrawlSeed(ctx, req.(*CrawlSeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontier_GetNextPage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FrontierServer).GetNextPage(&frontierGetNextPageServer{stream})
}

type Frontier_GetNextPageServer interface {
	Send(*PageHarvestSpec) error
	Recv() (*PageHarvest, error)
	grpc.ServerStream
}

type frontierGetNextPageServer struct {
	grpc.ServerStream
}

func (x *frontierGetNextPageServer) Send(m *PageHarvestSpec) error {
	return x.ServerStream.SendMsg(m)
}

func (x *frontierGetNextPageServer) Recv() (*PageHarvest, error) {
	m := new(PageHarvest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Frontier_BusyCrawlHostGroupCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontierServer).BusyCrawlHostGroupCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.frontier.v1.Frontier/BusyCrawlHostGroupCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontierServer).BusyCrawlHostGroupCount(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontier_QueueCountTotal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontierServer).QueueCountTotal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.frontier.v1.Frontier/QueueCountTotal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontierServer).QueueCountTotal(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontier_QueueCountForCrawlExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrawlExecutionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontierServer).QueueCountForCrawlExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.frontier.v1.Frontier/QueueCountForCrawlExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontierServer).QueueCountForCrawlExecution(ctx, req.(*CrawlExecutionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Frontier_QueueCountForCrawlHostGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrawlHostGroup)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FrontierServer).QueueCountForCrawlHostGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/veidemann.api.frontier.v1.Frontier/QueueCountForCrawlHostGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FrontierServer).QueueCountForCrawlHostGroup(ctx, req.(*CrawlHostGroup))
	}
	return interceptor(ctx, in, info, handler)
}

var _Frontier_serviceDesc = grpc.ServiceDesc{
	ServiceName: "veidemann.api.frontier.v1.Frontier",
	HandlerType: (*FrontierServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CrawlSeed",
			Handler:    _Frontier_CrawlSeed_Handler,
		},
		{
			MethodName: "BusyCrawlHostGroupCount",
			Handler:    _Frontier_BusyCrawlHostGroupCount_Handler,
		},
		{
			MethodName: "QueueCountTotal",
			Handler:    _Frontier_QueueCountTotal_Handler,
		},
		{
			MethodName: "QueueCountForCrawlExecution",
			Handler:    _Frontier_QueueCountForCrawlExecution_Handler,
		},
		{
			MethodName: "QueueCountForCrawlHostGroup",
			Handler:    _Frontier_QueueCountForCrawlHostGroup_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetNextPage",
			Handler:       _Frontier_GetNextPage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "frontier/v1/frontier.proto",
}
