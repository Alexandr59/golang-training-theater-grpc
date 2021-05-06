package api

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type UserServer struct {
	data *data.UserData
}

func NewUserServer(a data.UserData) *UserServer {
	return &UserServer{data: &a}
}

func (u UserServer) CreateUser(ctx context.Context, in *pb.UserRequest) (*pb.IdUserResponse, error) {
	entity := data.User{
		AccountId:   int(in.GetAccountId()),
		FirstName:   in.GetFirstName(),
		LastName:    in.GetLastName(),
		RoleId:      int(in.GetRoleId()),
		LocationId:  int(in.GetLocationId()),
		PhoneNumber: in.GetPhoneNumber(),
	}
	id, err := u.data.AddUser(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"user": entity,
		}).Warningf("got an error when tried to create user: %s", err)
		return &pb.IdUserResponse{Id: -1}, fmt.Errorf("got an error when tried to create user: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"user": entity,
	}).Info("create user")
	return &pb.IdUserResponse{Id: int64(id)}, nil
}

func (u UserServer) DeleteUser(ctx context.Context, in *pb.IdUserRequest) (*pb.StatusUserResponse, error) {
	entity := new(data.User)
	entity.Id = int(in.Id)
	err := u.data.DeleteUser(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete user: %s", err)
		return &pb.StatusUserResponse{Message: "got an error when tried to delete user"},
			fmt.Errorf("got an error when tried to delete user: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("user deletion was successful")
	return &pb.StatusUserResponse{Message: "user deletion was successful"}, nil
}

func (u UserServer) UpdateUser(ctx context.Context, in *pb.UserRequest) (*pb.StatusUserResponse, error) {
	entity := data.User{
		Id:          int(in.GetId()),
		AccountId:   int(in.GetAccountId()),
		FirstName:   in.GetFirstName(),
		LastName:    in.GetLastName(),
		RoleId:      int(in.GetRoleId()),
		LocationId:  int(in.GetLocationId()),
		PhoneNumber: in.GetPhoneNumber(),
	}
	err := u.data.UpdateUser(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"user": entity,
		}).Warningf("got an error when tried to update user: %s", err)
		return &pb.StatusUserResponse{Message: "got an error when tried to update user"},
			fmt.Errorf("got an error when tried to update user: %w", err)
	}
	log.WithFields(log.Fields{
		"user": entity,
	}).Info("user update was successful")
	return &pb.StatusUserResponse{Message: "user update was successful"}, nil
}

func (u UserServer) GetUser(ctx context.Context, in *pb.IdUserRequest) (*pb.UserResponse, error) {
	entity := new(data.User)
	entity.Id = int(in.Id)
	entry, err := u.data.FindByIdUser(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get user: %s", err)
		return &pb.UserResponse{},
			fmt.Errorf("got an error when tried to get user: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("user was successfully received")
	return &pb.UserResponse{
		Id:          int64(entry.Id),
		AccountId:   int64(entry.AccountId),
		FirstName:   entry.FirstName,
		LastName:    entry.LastName,
		RoleId:      int64(entry.RoleId),
		LocationId:  int64(entry.LocationId),
		PhoneNumber: entry.PhoneNumber,
	}, nil
}

func (u UserServer) GetAllUsers(ctx context.Context, in *pb.UsersRequest) (*pb.JsonUsersResponse, error) {
	account := new(data.Account)
	account.Id = int(in.GetId())
	users, err := u.data.ReadAllUsers(*account)
	if err != nil {
		log.WithFields(log.Fields{
			"users": users,
		}).Warningf("got an error when tried to get users: %s", err)
		return &pb.JsonUsersResponse{Json: ""},
			fmt.Errorf("got an error when tried to get users: %w", err)
	}
	json, err := json.Marshal(users)
	if err != nil {
		log.WithFields(log.Fields{
			"users": users,
		}).Warningf("got an error when tried to get json users: %s", err)
		return &pb.JsonUsersResponse{Json: ""},
			fmt.Errorf("got an error when tried to get json users: %w", err)
	}
	return &pb.JsonUsersResponse{Json: string(json)}, nil
}
