package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type RoleServer struct {
	data *data.RoleData
}

func NewRoleServer(a data.RoleData) *RoleServer {
	return &RoleServer{data: &a}
}

func (r RoleServer) CreateRole(ctx context.Context, in *pb.RoleRequest) (*pb.IdRoleResponse, error) {
	if err := checkRoleRequest(in); err != nil {
		log.WithFields(log.Fields{
			"role": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdRoleResponse{Id: -1}, err
	}
	entity := data.Role{
		Name: in.GetName(),
	}
	id, err := r.data.AddRole(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"role": entity,
		}).Warningf("got an error when tried to create role: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create role: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdRoleResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdRoleResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"role": entity,
	}).Info("role has been successfully created")
	return &pb.IdRoleResponse{Id: int64(id)}, nil
}

func (r RoleServer) DeleteRole(ctx context.Context, in *pb.IdRoleRequest) (*pb.StatusRoleResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"role": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusRoleResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Role)
	entity.Id = int(in.Id)
	err := r.data.DeleteRole(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete role: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete role: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusRoleResponse{Message: "got an error when tried to delete role"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusRoleResponse{Message: "got an error when tried to delete role"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("role deletion was successful")
	return &pb.StatusRoleResponse{Message: "role deletion was successful"}, nil
}

func (r RoleServer) UpdateRole(ctx context.Context, in *pb.RoleRequest) (*pb.StatusRoleResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"role": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusRoleResponse{Message: "empty fields error"}, err
	}
	if err := checkRoleRequest(in); err != nil {
		log.WithFields(log.Fields{
			"role": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusRoleResponse{Message: "empty fields error"}, err
	}
	entity := data.Role{
		Id:   int(in.GetId()),
		Name: in.GetName(),
	}
	err := r.data.UpdateRole(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"role": entity,
		}).Warningf("got an error when tried to update role: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update role: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusRoleResponse{Message: "got an error when tried to update role"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusRoleResponse{Message: "got an error when tried to update role"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"role": entity,
	}).Info("role update was successful")
	return &pb.StatusRoleResponse{Message: "role update was successful"}, nil
}

func (r RoleServer) GetRole(ctx context.Context, in *pb.IdRoleRequest) (*pb.RoleResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"role": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.RoleResponse{}, err
	}
	entity := new(data.Role)
	entity.Id = int(in.Id)
	entry, err := r.data.FindByIdRole(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get role: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get role: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.RoleResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.RoleResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("role was successfully received")
	return &pb.RoleResponse{
		Id:   int64(entry.Id),
		Name: entry.Name,
	}, nil
}

func checkRoleRequest(in *pb.RoleRequest) error {
	if in.GetName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Name}: %s", in.GetName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
