package api

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type PriceServer struct {
	data *data.PriceData
}

func NewPriceServer(a data.PriceData) *PriceServer {
	return &PriceServer{data: &a}
}

func (p PriceServer) CreatePrice(ctx context.Context, in *pb.PriceRequest) (*pb.IdPriceResponse, error) {
	entity := data.Price{
		AccountId:     int(in.GetAccountId()),
		SectorId:      int(in.GetSectorId()),
		PerformanceId: int(in.GetPerformanceId()),
		Price:         int(in.GetPrice()),
	}
	id, err := p.data.AddPrice(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"price": entity,
		}).Warningf("got an error when tried to create price: %s", err)
		return &pb.IdPriceResponse{Id: -1}, fmt.Errorf("got an error when tried to create price: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"price": entity,
	}).Info("create price")
	return &pb.IdPriceResponse{Id: int64(id)}, nil
}

func (p PriceServer) DeletePrice(ctx context.Context, in *pb.IdPriceRequest) (*pb.StatusPriceResponse, error) {
	entity := new(data.Price)
	entity.Id = int(in.Id)
	err := p.data.DeletePrice(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete price: %s", err)
		return &pb.StatusPriceResponse{Message: "got an error when tried to delete price"},
			fmt.Errorf("got an error when tried to delete price: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("price deletion was successful")
	return &pb.StatusPriceResponse{Message: "price deletion was successful"}, nil
}

func (p PriceServer) UpdatePrice(ctx context.Context, in *pb.PriceRequest) (*pb.StatusPriceResponse, error) {
	entity := data.Price{
		Id:            int(in.GetId()),
		AccountId:     int(in.GetAccountId()),
		SectorId:      int(in.GetSectorId()),
		PerformanceId: int(in.GetPerformanceId()),
		Price:         int(in.GetPrice()),
	}
	err := p.data.UpdatePrice(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"price": entity,
		}).Warningf("got an error when tried to update price: %s", err)
		return &pb.StatusPriceResponse{Message: "got an error when tried to update price"},
			fmt.Errorf("got an error when tried to update price: %w", err)
	}
	log.WithFields(log.Fields{
		"price": entity,
	}).Info("price update was successful")
	return &pb.StatusPriceResponse{Message: "price update was successful"}, nil
}

func (p PriceServer) GetPrice(ctx context.Context, in *pb.IdPriceRequest) (*pb.PriceResponse, error) {
	entity := new(data.Price)
	entity.Id = int(in.Id)
	entry, err := p.data.FindByIdPrice(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get price: %s", err)
		return &pb.PriceResponse{},
			fmt.Errorf("got an error when tried to get price: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("price was successfully received")
	return &pb.PriceResponse{
		Id:            int64(entry.Id),
		AccountId:     int64(entry.AccountId),
		SectorId:      int64(entry.SectorId),
		PerformanceId: int64(entry.PerformanceId),
		Price:         int64(entry.Price),
	}, nil
}
