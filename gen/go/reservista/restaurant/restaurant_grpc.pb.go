// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: reservista/restaurant/restaurant.proto

package proto_restaurant

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RestaurantClient is the client API for Restaurant service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestaurantClient interface {
	AddRestaurant(ctx context.Context, in *RestaurantObject, opts ...grpc.CallOption) (*StatusResponse, error)
	GetAllRestaurants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RestaurantListResponse, error)
	GetRestaurant(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*RestaurantObject, error)
	UpdateRestById(ctx context.Context, in *RestaurantObject, opts ...grpc.CallOption) (*StatusResponse, error)
	DeleteRestaurantById(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	UploadPhotos(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	DeletePhoto(ctx context.Context, in *DeletePhotoRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	SearchRestaurants(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	GetRestaurantSuggestions(ctx context.Context, in *SuggestionRequest, opts ...grpc.CallOption) (*RestaurantListResponse, error)
}

type restaurantClient struct {
	cc grpc.ClientConnInterface
}

func NewRestaurantClient(cc grpc.ClientConnInterface) RestaurantClient {
	return &restaurantClient{cc}
}

func (c *restaurantClient) AddRestaurant(ctx context.Context, in *RestaurantObject, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/AddRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantClient) GetAllRestaurants(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*RestaurantListResponse, error) {
	out := new(RestaurantListResponse)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/GetAllRestaurants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantClient) GetRestaurant(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*RestaurantObject, error) {
	out := new(RestaurantObject)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/GetRestaurant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantClient) UpdateRestById(ctx context.Context, in *RestaurantObject, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/UpdateRestById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantClient) DeleteRestaurantById(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/DeleteRestaurantById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantClient) UploadPhotos(ctx context.Context, in *UploadPhotoRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/UploadPhotos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantClient) DeletePhoto(ctx context.Context, in *DeletePhotoRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/DeletePhoto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantClient) SearchRestaurants(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/SearchRestaurants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *restaurantClient) GetRestaurantSuggestions(ctx context.Context, in *SuggestionRequest, opts ...grpc.CallOption) (*RestaurantListResponse, error) {
	out := new(RestaurantListResponse)
	err := c.cc.Invoke(ctx, "/restaurant.Restaurant/GetRestaurantSuggestions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RestaurantServer is the server API for Restaurant service.
// All implementations must embed UnimplementedRestaurantServer
// for forward compatibility
type RestaurantServer interface {
	AddRestaurant(context.Context, *RestaurantObject) (*StatusResponse, error)
	GetAllRestaurants(context.Context, *Empty) (*RestaurantListResponse, error)
	GetRestaurant(context.Context, *IDRequest) (*RestaurantObject, error)
	UpdateRestById(context.Context, *RestaurantObject) (*StatusResponse, error)
	DeleteRestaurantById(context.Context, *IDRequest) (*StatusResponse, error)
	UploadPhotos(context.Context, *UploadPhotoRequest) (*StatusResponse, error)
	DeletePhoto(context.Context, *DeletePhotoRequest) (*StatusResponse, error)
	SearchRestaurants(context.Context, *SearchRequest) (*SearchResponse, error)
	GetRestaurantSuggestions(context.Context, *SuggestionRequest) (*RestaurantListResponse, error)
	mustEmbedUnimplementedRestaurantServer()
}

// UnimplementedRestaurantServer must be embedded to have forward compatible implementations.
type UnimplementedRestaurantServer struct {
}

func (UnimplementedRestaurantServer) AddRestaurant(context.Context, *RestaurantObject) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRestaurant not implemented")
}
func (UnimplementedRestaurantServer) GetAllRestaurants(context.Context, *Empty) (*RestaurantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRestaurants not implemented")
}
func (UnimplementedRestaurantServer) GetRestaurant(context.Context, *IDRequest) (*RestaurantObject, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurant not implemented")
}
func (UnimplementedRestaurantServer) UpdateRestById(context.Context, *RestaurantObject) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRestById not implemented")
}
func (UnimplementedRestaurantServer) DeleteRestaurantById(context.Context, *IDRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRestaurantById not implemented")
}
func (UnimplementedRestaurantServer) UploadPhotos(context.Context, *UploadPhotoRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadPhotos not implemented")
}
func (UnimplementedRestaurantServer) DeletePhoto(context.Context, *DeletePhotoRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhoto not implemented")
}
func (UnimplementedRestaurantServer) SearchRestaurants(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchRestaurants not implemented")
}
func (UnimplementedRestaurantServer) GetRestaurantSuggestions(context.Context, *SuggestionRequest) (*RestaurantListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRestaurantSuggestions not implemented")
}
func (UnimplementedRestaurantServer) mustEmbedUnimplementedRestaurantServer() {}

// UnsafeRestaurantServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestaurantServer will
// result in compilation errors.
type UnsafeRestaurantServer interface {
	mustEmbedUnimplementedRestaurantServer()
}

func RegisterRestaurantServer(s grpc.ServiceRegistrar, srv RestaurantServer) {
	s.RegisterService(&Restaurant_ServiceDesc, srv)
}

func _Restaurant_AddRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).AddRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/AddRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).AddRestaurant(ctx, req.(*RestaurantObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restaurant_GetAllRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).GetAllRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/GetAllRestaurants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).GetAllRestaurants(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restaurant_GetRestaurant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).GetRestaurant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/GetRestaurant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).GetRestaurant(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restaurant_UpdateRestById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestaurantObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).UpdateRestById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/UpdateRestById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).UpdateRestById(ctx, req.(*RestaurantObject))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restaurant_DeleteRestaurantById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).DeleteRestaurantById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/DeleteRestaurantById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).DeleteRestaurantById(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restaurant_UploadPhotos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadPhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).UploadPhotos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/UploadPhotos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).UploadPhotos(ctx, req.(*UploadPhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restaurant_DeletePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePhotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).DeletePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/DeletePhoto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).DeletePhoto(ctx, req.(*DeletePhotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restaurant_SearchRestaurants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).SearchRestaurants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/SearchRestaurants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).SearchRestaurants(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Restaurant_GetRestaurantSuggestions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuggestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RestaurantServer).GetRestaurantSuggestions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/restaurant.Restaurant/GetRestaurantSuggestions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RestaurantServer).GetRestaurantSuggestions(ctx, req.(*SuggestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Restaurant_ServiceDesc is the grpc.ServiceDesc for Restaurant service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Restaurant_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "restaurant.Restaurant",
	HandlerType: (*RestaurantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRestaurant",
			Handler:    _Restaurant_AddRestaurant_Handler,
		},
		{
			MethodName: "GetAllRestaurants",
			Handler:    _Restaurant_GetAllRestaurants_Handler,
		},
		{
			MethodName: "GetRestaurant",
			Handler:    _Restaurant_GetRestaurant_Handler,
		},
		{
			MethodName: "UpdateRestById",
			Handler:    _Restaurant_UpdateRestById_Handler,
		},
		{
			MethodName: "DeleteRestaurantById",
			Handler:    _Restaurant_DeleteRestaurantById_Handler,
		},
		{
			MethodName: "UploadPhotos",
			Handler:    _Restaurant_UploadPhotos_Handler,
		},
		{
			MethodName: "DeletePhoto",
			Handler:    _Restaurant_DeletePhoto_Handler,
		},
		{
			MethodName: "SearchRestaurants",
			Handler:    _Restaurant_SearchRestaurants_Handler,
		},
		{
			MethodName: "GetRestaurantSuggestions",
			Handler:    _Restaurant_GetRestaurantSuggestions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservista/restaurant/restaurant.proto",
}
