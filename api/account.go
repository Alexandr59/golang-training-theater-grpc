package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type AccountServer struct {
	data *data.AccountData
}

func NewAccountServer(a data.AccountData) *AccountServer {
	return &AccountServer{data: &a}
}

func (a AccountServer) CreateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.IdAccountResponse, error) {
	if err := checkAccountRequest(in); err != nil {
		log.WithFields(log.Fields{
			"account": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdAccountResponse{Id: -1}, err
	}
	entity := data.Account{
		FirstName:   in.GetFirstName(),
		LastName:    in.GetLastName(),
		PhoneNumber: in.GetPhoneNumber(),
		Email:       in.GetEmail(),
	}
	id, err := a.data.AddAccount(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"account": entity,
		}).Warningf("got an error when tried to create account: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create account: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdAccountResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdAccountResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"account": entity,
	}).Info("account has been successfully created")
	return &pb.IdAccountResponse{Id: int64(id)}, nil
}

func (a AccountServer) DeleteAccount(ctx context.Context, in *pb.IdAccountRequest) (*pb.StatusAccountResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"account": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusAccountResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Account)
	entity.Id = int(in.Id)
	err := a.data.DeleteAccount(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete account: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete account: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusAccountResponse{Message: "got an error when tried to delete account"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusAccountResponse{Message: "got an error when tried to delete account"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("account deletion was successful")
	return &pb.StatusAccountResponse{Message: "account deletion was successful"}, nil
}

func (a AccountServer) UpdateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.StatusAccountResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"account": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusAccountResponse{Message: "empty fields error"}, err
	}
	if err := checkAccountRequest(in); err != nil {
		log.WithFields(log.Fields{
			"account": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusAccountResponse{Message: "empty fields error"}, err
	}
	entity := data.Account{
		Id:          int(in.GetId()),
		FirstName:   in.GetFirstName(),
		LastName:    in.GetLastName(),
		PhoneNumber: in.GetPhoneNumber(),
		Email:       in.GetEmail(),
	}
	err := a.data.UpdateAccount(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"account": entity,
		}).Warningf("got an error when tried to update account: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update account: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusAccountResponse{Message: "got an error when tried to update account"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusAccountResponse{Message: "got an error when tried to update account"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"account": entity,
	}).Info("account update was successful")
	return &pb.StatusAccountResponse{Message: "account update was successful"}, nil
}

func (a AccountServer) GetAccount(ctx context.Context, in *pb.IdAccountRequest) (*pb.AccountResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"account": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.AccountResponse{}, err
	}
	entity := new(data.Account)
	entity.Id = int(in.Id)
	entry, err := a.data.FindByIdAccount(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get account: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get account: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.AccountResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.AccountResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("account was successfully received")
	return &pb.AccountResponse{
		Id:          int64(entry.Id),
		FirstName:   entry.FirstName,
		LastName:    entry.LastName,
		PhoneNumber: entry.PhoneNumber,
		Email:       entry.Email,
	}, nil
}

func checkAccountRequest(in *pb.AccountRequest) error {
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
	if in.GetPhoneNumber() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {PhoneNumber}: %s", in.GetPhoneNumber())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetEmail() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Email}: %s", in.GetEmail())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}

func checkId(id int64) error {
	if id <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Id}: %s", id)
		return s.Err()
	}
	return nil
}
