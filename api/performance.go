package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	if err := checkPerformanceRequest(in); err != nil {
		log.WithFields(log.Fields{
			"performance": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdPerformanceResponse{Id: -1}, err
	}
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
		s := status.Newf(codes.Internal, "got an error when tried to create performance: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdPerformanceResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdPerformanceResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"performance": entity,
	}).Info("performance has been successfully created")
	return &pb.IdPerformanceResponse{Id: int64(id)}, nil
}

func (p PerformanceServer) DeletePerformance(ctx context.Context, in *pb.IdPerformanceRequest) (*pb.StatusPerformanceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"performance": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPerformanceResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Performance)
	entity.Id = int(in.Id)
	err := p.data.DeletePerformance(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete performance: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete performance: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusPerformanceResponse{Message: "got an error when tried to delete performance"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusPerformanceResponse{Message: "got an error when tried to delete performance"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("performance deletion was successful")
	return &pb.StatusPerformanceResponse{Message: "performance deletion was successful"}, nil
}

func (p PerformanceServer) UpdatePerformance(ctx context.Context, in *pb.PerformanceRequest) (*pb.StatusPerformanceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"performance": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPerformanceResponse{Message: "empty fields error"}, err
	}
	if err := checkPerformanceRequest(in); err != nil {
		log.WithFields(log.Fields{
			"performance": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPerformanceResponse{Message: "empty fields error"}, err
	}
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
		s := status.Newf(codes.Internal, "got an error when tried to update performance: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusPerformanceResponse{Message: "got an error when tried to update performance"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusPerformanceResponse{Message: "got an error when tried to update performance"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"performance": entity,
	}).Info("performance update was successful")
	return &pb.StatusPerformanceResponse{Message: "performance update was successful"}, nil
}

func (p PerformanceServer) GetPerformance(ctx context.Context, in *pb.IdPerformanceRequest) (*pb.PerformanceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"performance": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.PerformanceResponse{}, err
	}
	entity := new(data.Performance)
	entity.Id = int(in.Id)
	entry, err := p.data.FindByIdPerformance(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get performance: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get performance: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.PerformanceResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.PerformanceResponse{}, errWithDetails.Err()
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

func checkPerformanceRequest(in *pb.PerformanceRequest) error {
	if in.GetAccountId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {AccountId}: %s", in.GetAccountId())
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
	if in.GetGenreId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {GenreId}: %s", in.GetGenreId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetDuration() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Duration}: %s", in.GetDuration())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
