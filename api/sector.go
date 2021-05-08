package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type SectorServer struct {
	data *data.SectorData
}

func NewSectorServer(a data.SectorData) *SectorServer {
	return &SectorServer{data: &a}
}

func (s SectorServer) CreateSector(ctx context.Context, in *pb.SectorRequest) (*pb.IdSectorResponse, error) {
	if err := checkSectorRequest(in); err != nil {
		log.WithFields(log.Fields{
			"sector": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdSectorResponse{Id: -1}, err
	}
	entity := data.Sector{
		Name: in.GetName(),
	}
	id, err := s.data.AddSector(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"sector": entity,
		}).Warningf("got an error when tried to create sector: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create sector: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdSectorResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdSectorResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"sector": entity,
	}).Info("sector has been successfully created")
	return &pb.IdSectorResponse{Id: int64(id)}, nil
}

func (s SectorServer) DeleteSector(ctx context.Context, in *pb.IdSectorRequest) (*pb.StatusSectorResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"sector": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusSectorResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Sector)
	entity.Id = int(in.Id)
	err := s.data.DeleteSector(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete sector: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete sector: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusSectorResponse{Message: "got an error when tried to delete sector"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusSectorResponse{Message: "got an error when tried to delete sector"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("sector deletion was successful")
	return &pb.StatusSectorResponse{Message: "sector deletion was successful"}, nil
}

func (s SectorServer) UpdateSector(ctx context.Context, in *pb.SectorRequest) (*pb.StatusSectorResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"sector": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusSectorResponse{Message: "empty fields error"}, err
	}
	if err := checkSectorRequest(in); err != nil {
		log.WithFields(log.Fields{
			"sector": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusSectorResponse{Message: "empty fields error"}, err
	}
	entity := data.Sector{
		Id:   int(in.GetId()),
		Name: in.GetName(),
	}
	err := s.data.UpdateSector(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"sector": entity,
		}).Warningf("got an error when tried to update sector: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update sector: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusSectorResponse{Message: "got an error when tried to update sector"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusSectorResponse{Message: "got an error when tried to update sector"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"sector": entity,
	}).Info("sector update was successful")
	return &pb.StatusSectorResponse{Message: "sector update was successful"}, nil
}

func (s SectorServer) GetSector(ctx context.Context, in *pb.IdSectorRequest) (*pb.SectorResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"sector": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.SectorResponse{}, err
	}
	entity := new(data.Sector)
	entity.Id = int(in.Id)
	entry, err := s.data.FindByIdSector(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get sector: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get sector: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.SectorResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.SectorResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("sector was successfully received")
	return &pb.SectorResponse{
		Id:   int64(entry.Id),
		Name: entry.Name,
	}, nil
}

func checkSectorRequest(in *pb.SectorRequest) error {
	if in.GetName() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Name}: %s", in.GetName())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
