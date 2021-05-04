package api

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"

	//"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	"golang-training-theater-grpc/pkg/data"
	pb "golang-training-theater-grpc/proto/go_proto"
)

type SectorServer struct {
	data *data.SectorData
}

func NewSectorServer(a data.SectorData) *SectorServer {
	return &SectorServer{data: &a}
}

func (s SectorServer) CreateSector(ctx context.Context, in *pb.SectorRequest) (*pb.IdSectorResponse, error) {
	entity := data.Sector{
		Name: in.GetName(),
	}
	id, err := s.data.AddSector(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"sector": entity,
		}).Warningf("got an error when tried to create sector: %s", err)
		return &pb.IdSectorResponse{Id: -1}, fmt.Errorf("got an error when tried to create sector: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"sector": entity,
	}).Info("create sector")
	return &pb.IdSectorResponse{Id: int64(id)}, nil
}

func (s SectorServer) DeleteSector(ctx context.Context, in *pb.IdSectorRequest) (*pb.StatusSectorResponse, error) {
	entity := new(data.Sector)
	entity.Id = int(in.Id)
	err := s.data.DeleteSector(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete sector: %s", err)
		return &pb.StatusSectorResponse{Message: "got an error when tried to delete sector"},
			fmt.Errorf("got an error when tried to delete sector: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("sector deletion was successful")
	return &pb.StatusSectorResponse{Message: "sector deletion was successful"}, nil
}

func (s SectorServer) UpdateSector(ctx context.Context, in *pb.SectorRequest) (*pb.StatusSectorResponse, error) {
	entity := data.Sector{
		Id:   int(in.GetId()),
		Name: in.GetName(),
	}
	err := s.data.UpdateSector(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"sector": entity,
		}).Warningf("got an error when tried to update sector: %s", err)
		return &pb.StatusSectorResponse{Message: "got an error when tried to update sector"},
			fmt.Errorf("got an error when tried to update sector: %w", err)
	}
	log.WithFields(log.Fields{
		"sector": entity,
	}).Info("sector update was successful")
	return &pb.StatusSectorResponse{Message: "sector update was successful"}, nil
}

func (s SectorServer) GetSector(ctx context.Context, in *pb.IdSectorRequest) (*pb.SectorResponse, error) {
	entity := new(data.Sector)
	entity.Id = int(in.Id)
	entry, err := s.data.FindByIdSector(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get sector: %s", err)
		return &pb.SectorResponse{},
			fmt.Errorf("got an error when tried to get sector: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("sector was successfully received")
	return &pb.SectorResponse{
		Id:   int64(entry.Id),
		Name: entry.Name,
	}, nil
}
