package api

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type LocationServer struct {
	data *data.LocationData
}

func NewLocationServer(a data.LocationData) *LocationServer {
	return &LocationServer{data: &a}
}

func (l LocationServer) CreateLocation(ctx context.Context, in *pb.LocationRequest) (*pb.IdLocationResponse, error) {
	entity := data.Location{
		AccountId:   int(in.GetAccountId()),
		Address:     in.GetAddress(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	id, err := l.data.AddLocation(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"location": entity,
		}).Warningf("got an error when tried to create location: %s", err)
		return &pb.IdLocationResponse{Id: -1}, fmt.Errorf("got an error when tried to create location: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"location": entity,
	}).Info("create location")
	return &pb.IdLocationResponse{Id: int64(id)}, nil
}

func (l LocationServer) DeleteLocation(ctx context.Context, in *pb.IdLocationRequest) (*pb.StatusLocationResponse, error) {
	entity := new(data.Location)
	entity.Id = int(in.Id)
	err := l.data.DeleteLocation(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete location: %s", err)
		return &pb.StatusLocationResponse{Message: "got an error when tried to delete location"},
			fmt.Errorf("got an error when tried to delete location: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("location deletion was successful")
	return &pb.StatusLocationResponse{Message: "location deletion was successful"}, nil
}

func (l LocationServer) UpdateLocation(ctx context.Context, in *pb.LocationRequest) (*pb.StatusLocationResponse, error) {
	entity := data.Location{
		Id:          int(in.GetId()),
		AccountId:   int(in.GetAccountId()),
		Address:     in.GetAddress(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	err := l.data.UpdateLocation(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"location": entity,
		}).Warningf("got an error when tried to update location: %s", err)
		return &pb.StatusLocationResponse{Message: "got an error when tried to update location"},
			fmt.Errorf("got an error when tried to update location: %w", err)
	}
	log.WithFields(log.Fields{
		"location": entity,
	}).Info("location update was successful")
	return &pb.StatusLocationResponse{Message: "location update was successful"}, nil
}

func (l LocationServer) GetLocation(ctx context.Context, in *pb.IdLocationRequest) (*pb.LocationResponse, error) {
	entity := new(data.Location)
	entity.Id = int(in.Id)
	entry, err := l.data.FindByIdLocation(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get location: %s", err)
		return &pb.LocationResponse{},
			fmt.Errorf("got an error when tried to get location: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("account was successfully received")
	return &pb.LocationResponse{
		Id:          int64(entry.Id),
		AccountId:   int64(entry.AccountId),
		Address:     entry.Address,
		PhoneNumber: entry.PhoneNumber,
	}, nil
}
