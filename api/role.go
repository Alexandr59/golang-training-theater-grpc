package api

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	//pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
	pb "golang-training-theater-grpc/proto/go_proto"
)

type RoleServer struct {
	data *data.RoleData
}

func NewRoleServer(a data.RoleData) *RoleServer {
	return &RoleServer{data: &a}
}

func (r RoleServer) CreateRole(ctx context.Context, in *pb.RoleRequest) (*pb.IdRoleResponse, error) {
	entity := data.Role{
		Name: in.GetName(),
	}
	id, err := r.data.AddRole(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"role": entity,
		}).Warningf("got an error when tried to create role: %s", err)
		return &pb.IdRoleResponse{Id: -1}, fmt.Errorf("got an error when tried to create role: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"role": entity,
	}).Info("create role")
	return &pb.IdRoleResponse{Id: int64(id)}, nil
}

func (r RoleServer) DeleteRole(ctx context.Context, in *pb.IdRoleRequest) (*pb.StatusRoleResponse, error) {
	entity := new(data.Role)
	entity.Id = int(in.Id)
	err := r.data.DeleteRole(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete role: %s", err)
		return &pb.StatusRoleResponse{Message: "got an error when tried to delete role"},
			fmt.Errorf("got an error when tried to delete role: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("role deletion was successful")
	return &pb.StatusRoleResponse{Message: "role deletion was successful"}, nil
}

func (r RoleServer) UpdateRole(ctx context.Context, in *pb.RoleRequest) (*pb.StatusRoleResponse, error) {
	entity := data.Role{
		Id:   int(in.GetId()),
		Name: in.GetName(),
	}
	err := r.data.UpdateRole(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"role": entity,
		}).Warningf("got an error when tried to update role: %s", err)
		return &pb.StatusRoleResponse{Message: "got an error when tried to update role"},
			fmt.Errorf("got an error when tried to update role: %w", err)
	}
	log.WithFields(log.Fields{
		"account": entity,
	}).Info("role update was successful")
	return &pb.StatusRoleResponse{Message: "role update was successful"}, nil
}

func (r RoleServer) GetRole(ctx context.Context, in *pb.IdRoleRequest) (*pb.RoleResponse, error) {
	entity := new(data.Role)
	entity.Id = int(in.Id)
	entry, err := r.data.FindByIdRole(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get role: %s", err)
		return &pb.RoleResponse{},
			fmt.Errorf("got an error when tried to get role: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("role was successfully received")
	return &pb.RoleResponse{
		Id:   int64(entry.Id),
		Name: entry.Name,
	}, nil
}
