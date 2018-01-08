package service

import (
	"bitbucket.org/stack-rox/apollo/apollo/db"
	"bitbucket.org/stack-rox/apollo/apollo/detection"
	"bitbucket.org/stack-rox/apollo/generated/api/v1"
	"bitbucket.org/stack-rox/apollo/pkg/registries"
	"bitbucket.org/stack-rox/apollo/pkg/secrets"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// NewRegistryService returns the RegistryService API.
func NewRegistryService(storage db.RegistryStorage, detection *detection.Detector) *RegistryService {
	return &RegistryService{
		storage:  storage,
		detector: detection,
	}
}

// RegistryService is the struct that manages the Registry API
type RegistryService struct {
	storage  db.RegistryStorage
	detector *detection.Detector
}

// RegisterServiceServer registers this service with the given gRPC Server.
func (s *RegistryService) RegisterServiceServer(grpcServer *grpc.Server) {
	v1.RegisterRegistryServiceServer(grpcServer, s)
}

// RegisterServiceHandlerFromEndpoint registers this service with the given gRPC Gateway endpoint.
func (s *RegistryService) RegisterServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return v1.RegisterRegistryServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// GetRegistries retrieves all registries that matches the request filters
func (s *RegistryService) GetRegistries(ctx context.Context, request *v1.GetRegistriesRequest) (*v1.GetRegistriesResponse, error) {
	registries, err := s.storage.GetRegistries(request)
	if err != nil {
		return nil, err
	}

	if request.RequestorIsAgent {
		return &v1.GetRegistriesResponse{Registries: registries}, nil
	}

	// Remove secrets for other API accessors.
	for _, r := range registries {
		r.Config = secrets.ScrubSecrets(r.Config)
	}
	return &v1.GetRegistriesResponse{Registries: registries}, nil
}

// PutRegistry updates a registry in the system
func (s *RegistryService) PutRegistry(ctx context.Context, request *v1.Registry) (*empty.Empty, error) {
	// creates and validates the configuration
	registry, err := registries.CreateRegistry(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.storage.UpdateRegistry(request); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	s.detector.UpdateRegistry(registry)
	return &empty.Empty{}, nil
}

// PostRegistry inserts a new registry into the system if it doesn't already exist
func (s *RegistryService) PostRegistry(ctx context.Context, request *v1.Registry) (*empty.Empty, error) {
	// creates and validates the configuration
	registry, err := registries.CreateRegistry(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	if err := s.storage.AddRegistry(request); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	s.detector.UpdateRegistry(registry)
	return &empty.Empty{}, nil
}
