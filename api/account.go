package api

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	//"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	"golang-training-theater-grpc/pkg/data"
	//pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
	pb "golang-training-theater-grpc/proto/go_proto"
)

type AccountServer struct {
	data *data.AccountData
}

func NewAccountServer(a data.AccountData) *AccountServer {
	return &AccountServer{data: &a}
}

func (a AccountServer) CreateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.IdAccountResponse, error) {
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
		return &pb.IdAccountResponse{Id: -1}, fmt.Errorf("got an error when tried to create account: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"account": entity,
	}).Info("create account")
	return &pb.IdAccountResponse{Id: int64(id)}, nil
}

func (a AccountServer) DeleteAccount(ctx context.Context, in *pb.IdAccountRequest) (*pb.StatusAccountResponse, error) {
	entity := new(data.Account)
	entity.Id = int(in.Id)
	err := a.data.DeleteAccount(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete account: %s", err)
		return &pb.StatusAccountResponse{Message: "got an error when tried to delete account"},
			fmt.Errorf("got an error when tried to delete account: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("account deletion was successful")
	return &pb.StatusAccountResponse{Message: "account deletion was successful"}, nil
}

func (a AccountServer) UpdateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.StatusAccountResponse, error) {
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
		return &pb.StatusAccountResponse{Message: "got an error when tried to update account"},
			fmt.Errorf("got an error when tried to update account: %w", err)
	}
	log.WithFields(log.Fields{
		"account": entity,
	}).Info("account update was successful")
	return &pb.StatusAccountResponse{Message: "account update was successful"}, nil
}

func (a AccountServer) GetAccount(ctx context.Context, in *pb.IdAccountRequest) (*pb.AccountResponse, error) {
	entity := new(data.Account)
	entity.Id = int(in.Id)
	entry, err := a.data.FindByIdAccount(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get account: %s", err)
		return &pb.AccountResponse{},
			fmt.Errorf("got an error when tried to get account: %w", err)
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
