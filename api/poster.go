package api

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type PosterServer struct {
	data *data.PosterData
}

func NewPosterServer(a data.PosterData) *PosterServer {
	return &PosterServer{data: &a}
}

func (p PosterServer) CreatePoster(ctx context.Context, in *pb.PosterRequest) (*pb.IdPosterResponse, error) {
	if err := checkPosterRequest(in); err != nil {
		log.WithFields(log.Fields{
			"poster": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdPosterResponse{Id: -1}, err
	}
	entity := data.Poster{
		AccountId:  int(in.GetAccountId()),
		ScheduleId: int(in.GetScheduleId()),
		Comment:    in.GetComment(),
	}
	id, err := p.data.AddPoster(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"poster": entity,
		}).Warningf("got an error when tried to create poster: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create poster: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdPosterResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdPosterResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"poster": entity,
	}).Info("poster has been successfully created")
	return &pb.IdPosterResponse{Id: int64(id)}, nil
}

func (p PosterServer) DeletePoster(ctx context.Context, in *pb.IdPosterRequest) (*pb.StatusPosterResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"poster": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPosterResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Poster)
	entity.Id = int(in.Id)
	err := p.data.DeletePoster(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete poster: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete poster: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusPosterResponse{Message: "got an error when tried to delete poster"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusPosterResponse{Message: "got an error when tried to delete poster"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("poster deletion was successful")
	return &pb.StatusPosterResponse{Message: "poster deletion was successful"}, nil
}

func (p PosterServer) UpdatePoster(ctx context.Context, in *pb.PosterRequest) (*pb.StatusPosterResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"poster": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPosterResponse{Message: "empty fields error"}, err
	}
	if err := checkPosterRequest(in); err != nil {
		log.WithFields(log.Fields{
			"poster": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPosterResponse{Message: "empty fields error"}, err
	}
	entity := data.Poster{
		Id:         int(in.GetId()),
		AccountId:  int(in.GetAccountId()),
		ScheduleId: int(in.GetScheduleId()),
		Comment:    in.GetComment(),
	}
	err := p.data.UpdatePoster(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"poster": entity,
		}).Warningf("got an error when tried to update poster: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update poster: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusPosterResponse{Message: "got an error when tried to update poster"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusPosterResponse{Message: "got an error when tried to update poster"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"poster": entity,
	}).Info("poster update was successful")
	return &pb.StatusPosterResponse{Message: "poster update was successful"}, nil
}

func (p PosterServer) GetPoster(ctx context.Context, in *pb.IdPosterRequest) (*pb.PosterResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"poster": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.PosterResponse{}, err
	}
	entity := new(data.Poster)
	entity.Id = int(in.Id)
	entry, err := p.data.FindByIdPoster(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get poster: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get poster: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.PosterResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.PosterResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("poster was successfully received")
	return &pb.PosterResponse{
		Id:         int64(entry.Id),
		AccountId:  int64(entry.AccountId),
		ScheduleId: int64(entry.ScheduleId),
		Comment:    entry.Comment,
	}, nil
}

func (p PosterServer) GetAllPosters(ctx context.Context, in *pb.Request) (*pb.JsonResponse, error) {
	posters, err := p.data.ReadAllPosters()
	if err != nil {
		log.WithFields(log.Fields{
			"posters": posters,
		}).Warningf("got an error when tried to get posters: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get posters: %w", err)
		return &pb.JsonResponse{Json: ""}, s.Err()
	}
	jsonPosters, err := json.Marshal(posters)
	if err != nil {
		log.WithFields(log.Fields{
			"posters": posters,
		}).Warningf("got an error when tried to get json posters: %s", err)
		s := status.Newf(codes.Unknown, "got an error when tried to get json posters: %w", err)
		return &pb.JsonResponse{Json: ""}, s.Err()
	}
	log.WithFields(log.Fields{
		"poster": jsonPosters,
	}).Info("posters was successfully received")
	return &pb.JsonResponse{Json: string(jsonPosters)}, nil
}

func checkPosterRequest(in *pb.PosterRequest) error {
	if in.GetAccountId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {AccountId}: %s", in.GetAccountId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetScheduleId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {ScheduleId}: %s", in.GetScheduleId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetComment() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Comment}: %s", in.GetComment())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
