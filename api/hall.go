package api

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	//pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
	pb "golang-training-theater-grpc/proto/go_proto"
)

type HallServer struct {
	data *data.HallData
}

func NewHallServer(a data.HallData) *HallServer {
	return &HallServer{data: &a}
}

func (h HallServer) CreateHall(ctx context.Context, in *pb.HallRequest) (*pb.IdHallResponse, error) {
	entity := data.Hall{
		AccountId:  int(in.GetAccountId()),
		Name:       in.GetName(),
		Capacity:   int(in.GetCapacity()),
		LocationId: int(in.GetLocationId()),
	}
	id, err := h.data.AddHall(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"hall": entity,
		}).Warningf("got an error when tried to create hall: %s", err)
		return &pb.IdHallResponse{Id: -1}, fmt.Errorf("got an error when tried to create hall: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"hall": entity,
	}).Info("create hall")
	return &pb.IdHallResponse{Id: int64(id)}, nil
}

func (h HallServer) DeleteHall(ctx context.Context, in *pb.IdHallRequest) (*pb.StatusHallResponse, error) {
	entity := new(data.Hall)
	entity.Id = int(in.Id)
	err := h.data.DeleteHall(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete hall: %s", err)
		return &pb.StatusHallResponse{Message: "got an error when tried to delete hall"},
			fmt.Errorf("got an error when tried to delete hall: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("hall deletion was successful")
	return &pb.StatusHallResponse{Message: "hall deletion was successful"}, nil
}

func (h HallServer) UpdateHall(ctx context.Context, in *pb.HallRequest) (*pb.StatusHallResponse, error) {
	entity := data.Hall{
		Id:         int(in.GetId()),
		AccountId:  int(in.GetAccountId()),
		Name:       in.GetName(),
		Capacity:   int(in.GetCapacity()),
		LocationId: int(in.GetLocationId()),
	}
	err := h.data.UpdateHall(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"hall": entity,
		}).Warningf("got an error when tried to update hall: %s", err)
		return &pb.StatusHallResponse{Message: "got an error when tried to update hall"},
			fmt.Errorf("got an error when tried to update hall: %w", err)
	}
	log.WithFields(log.Fields{
		"hall": entity,
	}).Info("hall update was successful")
	return &pb.StatusHallResponse{Message: "hall update was successful"}, nil
}

func (h HallServer) GetHall(ctx context.Context, in *pb.IdHallRequest) (*pb.HallResponse, error) {
	entity := new(data.Hall)
	entity.Id = int(in.Id)
	entry, err := h.data.FindByIdHall(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get hall: %s", err)
		return &pb.HallResponse{},
			fmt.Errorf("got an error when tried to get hall: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("hall was successfully received")
	return &pb.HallResponse{
		Id:         int64(entry.Id),
		AccountId:  int64(entry.AccountId),
		Name:       entry.Name,
		Capacity:   int64(entry.Capacity),
		LocationId: int64(entry.LocationId),
	}, nil
}
