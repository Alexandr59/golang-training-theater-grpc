package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type PlaceServer struct {
	data *data.PlaceData
}

func NewPlaceServer(a data.PlaceData) *PlaceServer {
	return &PlaceServer{data: &a}
}

func (p PlaceServer) CreatePlace(ctx context.Context, in *pb.PlaceRequest) (*pb.IdPlaceResponse, error) {
	if err := checkPlaceRequest(in); err != nil {
		log.WithFields(log.Fields{
			"place": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdPlaceResponse{Id: -1}, err
	}
	entity := data.Place{
		SectorId: int(in.GetSectorId()),
		Name:     in.GetName(),
	}
	id, err := p.data.AddPlace(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"place": entity,
		}).Warningf("got an error when tried to create place: %s", err)
		s := status.Newf(codes.Canceled, "got an error when tried to create place: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdPlaceResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdPlaceResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"place": entity,
	}).Info("place has been successfully created")
	return &pb.IdPlaceResponse{Id: int64(id)}, nil
}

func (p PlaceServer) DeletePlace(ctx context.Context, in *pb.IdPlaceRequest) (*pb.StatusPlaceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"place": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPlaceResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Place)
	entity.Id = int(in.Id)
	err := p.data.DeletePlace(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete place: %s", err)
		s := status.Newf(codes.Canceled, "got an error when tried to delete place: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusPlaceResponse{Message: "got an error when tried to delete place"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusPlaceResponse{Message: "got an error when tried to delete place"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("place deletion was successful")
	return &pb.StatusPlaceResponse{Message: "place deletion was successful"}, nil
}

func (p PlaceServer) UpdatePlace(ctx context.Context, in *pb.PlaceRequest) (*pb.StatusPlaceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"place": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPlaceResponse{Message: "empty fields error"}, err
	}
	if err := checkPlaceRequest(in); err != nil {
		log.WithFields(log.Fields{
			"place": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPlaceResponse{Message: "empty fields error"}, err
	}
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
		s := status.Newf(codes.Canceled, "got an error when tried to update place: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusPlaceResponse{Message: "got an error when tried to update place"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusPlaceResponse{Message: "got an error when tried to update place"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"place": entity,
	}).Info("place update was successful")
	return &pb.StatusPlaceResponse{Message: "place update was successful"}, nil
}

func (p PlaceServer) GetPlace(ctx context.Context, in *pb.IdPlaceRequest) (*pb.PlaceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"place": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.PlaceResponse{}, err
	}
	entity := new(data.Place)
	entity.Id = int(in.Id)
	entry, err := p.data.FindByIdPlace(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get place: %s", err)
		s := status.Newf(codes.Canceled, "got an error when tried to get place: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.PlaceResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.PlaceResponse{}, errWithDetails.Err()
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

func checkPlaceRequest(in *pb.PlaceRequest) error {
	if in.GetSectorId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {SectorId}: %s", in.GetSectorId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
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
