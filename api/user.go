package api

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	if err := checkUserRequest(in); err != nil {
		log.WithFields(log.Fields{
			"user": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdUserResponse{Id: -1}, err
	}
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
		s := status.Newf(codes.Internal, "got an error when tried to create user: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdUserResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdUserResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"user": entity,
	}).Info("create user")
	return &pb.IdUserResponse{Id: int64(id)}, nil
}

func (u UserServer) DeleteUser(ctx context.Context, in *pb.IdUserRequest) (*pb.StatusUserResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"User": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusUserResponse{Message: "empty fields error"}, err
	}
	entity := new(data.User)
	entity.Id = int(in.Id)
	err := u.data.DeleteUser(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete user: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete delete: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusUserResponse{Message: "got an error when tried to delete delete"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusUserResponse{Message: "got an error when tried to delete delete"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("user deletion was successful")
	return &pb.StatusUserResponse{Message: "user deletion was successful"}, nil
}

func (u UserServer) UpdateUser(ctx context.Context, in *pb.UserRequest) (*pb.StatusUserResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"User": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusUserResponse{Message: "empty fields error"}, err
	}
	if err := checkUserRequest(in); err != nil {
		log.WithFields(log.Fields{
			"user": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusUserResponse{Message: "empty fields error"}, err
	}
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
		s := status.Newf(codes.Internal, "got an error when tried to update user: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusUserResponse{Message: "got an error when tried to update user"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusUserResponse{Message: "got an error when tried to update user"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"user": entity,
	}).Info("user update was successful")
	return &pb.StatusUserResponse{Message: "user update was successful"}, nil
}

func (u UserServer) GetUser(ctx context.Context, in *pb.IdUserRequest) (*pb.UserResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"User": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.UserResponse{}, err
	}
	entity := new(data.User)
	entity.Id = int(in.Id)
	entry, err := u.data.FindByIdUser(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get user: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get user: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.UserResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.UserResponse{}, errWithDetails.Err()
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
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"User": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.JsonUsersResponse{Json: ""}, err
	}
	account := new(data.Account)
	account.Id = int(in.GetId())
	users, err := u.data.ReadAllUsers(*account)
	if err != nil {
		log.WithFields(log.Fields{
			"users": users,
		}).Warningf("got an error when tried to get users: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get users: %w", err)
		return &pb.JsonUsersResponse{Json: ""}, s.Err()
	}
	json, err := json.Marshal(users)
	if err != nil {
		log.WithFields(log.Fields{
			"users": users,
		}).Warningf("got an error when tried to get json users: %s", err)
		s := status.Newf(codes.Unknown, "got an error when tried to get json users: %w", err)
		return &pb.JsonUsersResponse{Json: ""}, s.Err()
	}
	return &pb.JsonUsersResponse{Json: string(json)}, nil
}

func checkUserRequest(in *pb.UserRequest) error {
	if in.GetAccountId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {AccountId}: %s", in.GetAccountId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetFirstName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {FirstName}: %s", in.GetFirstName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetLastName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {LastName}: %s", in.GetLastName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetRoleId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {RoleId}: %s", in.GetRoleId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetLocationId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {LocationId}: %s", in.GetLocationId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetPhoneNumber() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {PhoneNumber}: %s", in.GetPhoneNumber())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
