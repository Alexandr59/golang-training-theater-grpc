package api

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type PerformanceServer struct {
	data *data.PerformanceData
}

func NewPerformanceServer(a data.PerformanceData) *PerformanceServer {
	return &PerformanceServer{data: &a}
}

func (p PerformanceServer) CreatePerformance(ctx context.Context, in *pb.PerformanceRequest) (*pb.IdPerformanceResponse, error) {
	entity := data.Performance{
		AccountId: int(in.GetAccountId()),
		Name:      in.GetName(),
		GenreId:   int(in.GetGenreId()),
		Duration:  in.GetDuration(),
	}
	id, err := p.data.AddPerformance(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"performance": entity,
		}).Warningf("got an error when tried to create performance: %s", err)
		return &pb.IdPerformanceResponse{Id: -1}, fmt.Errorf("got an error when tried to create performance: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"performance": entity,
	}).Info("create performance")
	return &pb.IdPerformanceResponse{Id: int64(id)}, nil
}

func (p PerformanceServer) DeletePerformance(ctx context.Context, in *pb.IdPerformanceRequest) (*pb.StatusPerformanceResponse, error) {
	entity := new(data.Performance)
	entity.Id = int(in.Id)
	err := p.data.DeletePerformance(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete performance: %s", err)
		return &pb.StatusPerformanceResponse{Message: "got an error when tried to delete performance"},
			fmt.Errorf("got an error when tried to delete performance: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("performance deletion was successful")
	return &pb.StatusPerformanceResponse{Message: "performance deletion was successful"}, nil
}

func (p PerformanceServer) UpdatePerformance(ctx context.Context, in *pb.PerformanceRequest) (*pb.StatusPerformanceResponse, error) {
	entity := data.Performance{
		Id:        int(in.GetId()),
		AccountId: int(in.GetAccountId()),
		Name:      in.GetName(),
		GenreId:   int(in.GetGenreId()),
		Duration:  in.GetDuration(),
	}
	err := p.data.UpdatePerformance(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"performance": entity,
		}).Warningf("got an error when tried to update performance: %s", err)
		return &pb.StatusPerformanceResponse{Message: "got an error when tried to update performance"},
			fmt.Errorf("got an error when tried to update performance: %w", err)
	}
	log.WithFields(log.Fields{
		"performance": entity,
	}).Info("performance update was successful")
	return &pb.StatusPerformanceResponse{Message: "performance update was successful"}, nil
}

func (p PerformanceServer) GetPerformance(ctx context.Context, in *pb.IdPerformanceRequest) (*pb.PerformanceResponse, error) {
	entity := new(data.Performance)
	entity.Id = int(in.Id)
	entry, err := p.data.FindByIdPerformance(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get performance: %s", err)
		return &pb.PerformanceResponse{},
			fmt.Errorf("got an error when tried to get performance: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("performance was successfully received")
	return &pb.PerformanceResponse{
		Id:        int64(entry.Id),
		AccountId: int64(entry.AccountId),
		Name:      entry.Name,
		GenreId:   int64(entry.GenreId),
		Duration:  entry.Duration,
	}, nil
}
