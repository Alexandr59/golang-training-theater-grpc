package api

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"

	//"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	"golang-training-theater-grpc/pkg/data"
	pb "golang-training-theater-grpc/proto/go_proto"
)

type PlaceServer struct {
	data *data.PlaceData
}

func NewPlaceServer(a data.PlaceData) *PlaceServer {
	return &PlaceServer{data: &a}
}

func (p PlaceServer) CreatePlace(ctx context.Context, in *pb.PlaceRequest) (*pb.IdPlaceResponse, error) {
	entity := data.Place{
		SectorId: int(in.GetSectorId()),
		Name:     in.GetName(),
	}
	id, err := p.data.AddPlace(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"place": entity,
		}).Warningf("got an error when tried to create place: %s", err)
		return &pb.IdPlaceResponse{Id: -1}, fmt.Errorf("got an error when tried to create place: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"place": entity,
	}).Info("create place")
	return &pb.IdPlaceResponse{Id: int64(id)}, nil
}

func (p PlaceServer) DeletePlace(ctx context.Context, in *pb.IdPlaceRequest) (*pb.StatusPlaceResponse, error) {
	entity := new(data.Place)
	entity.Id = int(in.Id)
	err := p.data.DeletePlace(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete place: %s", err)
		return &pb.StatusPlaceResponse{Message: "got an error when tried to delete place"},
			fmt.Errorf("got an error when tried to delete place: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("place deletion was successful")
	return &pb.StatusPlaceResponse{Message: "place deletion was successful"}, nil
}

func (p PlaceServer) UpdatePlace(ctx context.Context, in *pb.PlaceRequest) (*pb.StatusPlaceResponse, error) {
	entity := data.Place{
		Id:       int(in.GetId()),
		SectorId: int(in.GetSectorId()),
		Name:     in.GetName(),
	}
	err := p.data.UpdatePlace(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"place": entity,
		}).Warningf("got an error when tried to update place: %s", err)
		return &pb.StatusPlaceResponse{Message: "got an error when tried to update place"},
			fmt.Errorf("got an error when tried to update place: %w", err)
	}
	log.WithFields(log.Fields{
		"place": entity,
	}).Info("place update was successful")
	return &pb.StatusPlaceResponse{Message: "place update was successful"}, nil
}

func (p PlaceServer) GetPlace(ctx context.Context, in *pb.IdPlaceRequest) (*pb.PlaceResponse, error) {
	entity := new(data.Place)
	entity.Id = int(in.Id)
	entry, err := p.data.FindByIdPlace(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get place: %s", err)
		return &pb.PlaceResponse{},
			fmt.Errorf("got an error when tried to get place: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("place was successfully received")
	return &pb.PlaceResponse{
		Id:       int64(entry.Id),
		SectorId: int64(entry.SectorId),
		Name:     entry.Name,
	}, nil
}
