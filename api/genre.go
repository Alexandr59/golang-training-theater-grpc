package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type GenreServer struct {
	data *data.GenreData
}

func NewGenreServer(a data.GenreData) *GenreServer {
	return &GenreServer{data: &a}
}

func (g GenreServer) CreateGenre(ctx context.Context, in *pb.GenreRequest) (*pb.IdGenreResponse, error) {
	if err := checkGenreRequest(in); err != nil {
		log.WithFields(log.Fields{
			"genre": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdGenreResponse{Id: -1}, err
	}
	entity := data.Genre{
		Name: in.GetName(),
	}
	id, err := g.data.AddGenre(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"genre": entity,
		}).Warningf("got an error when tried to create genre: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create genre: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdGenreResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdGenreResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"genre": entity,
	}).Info("account has been successfully created")
	return &pb.IdGenreResponse{Id: int64(id)}, nil
}

func (g GenreServer) DeleteGenre(ctx context.Context, in *pb.IdGenreRequest) (*pb.StatusGenreResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"genre": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusGenreResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Genre)
	entity.Id = int(in.GetId())
	err := g.data.DeleteGenre(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete genre: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete genre: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusGenreResponse{Message: "got an error when tried to delete genre"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusGenreResponse{Message: "got an error when tried to delete genre"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("genre deletion was successful")
	return &pb.StatusGenreResponse{Message: "genre deletion was successful"}, nil
}

func (g GenreServer) UpdateGenre(ctx context.Context, in *pb.GenreRequest) (*pb.StatusGenreResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"genre": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusGenreResponse{Message: "empty fields error"}, err
	}
	if err := checkGenreRequest(in); err != nil {
		log.WithFields(log.Fields{
			"genre": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusGenreResponse{Message: "empty fields error"}, err
	}
	entity := data.Genre{
		Id:   int(in.GetId()),
		Name: in.GetName(),
	}
	err := g.data.UpdateGenre(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"genre": entity,
		}).Warningf("got an error when tried to update genre: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update genre: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusGenreResponse{Message: "got an error when tried to update genre"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusGenreResponse{Message: "got an error when tried to update genre"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"genre": entity,
	}).Info("genre update was successful")
	return &pb.StatusGenreResponse{Message: "genre update was successful"}, nil
}

func (g GenreServer) GetGenre(ctx context.Context, in *pb.IdGenreRequest) (*pb.GenreResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"genre": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.GenreResponse{}, err
	}
	entity := new(data.Genre)
	entity.Id = int(in.Id)
	entry, err := g.data.FindByIdGenre(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get genre: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get genre: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.GenreResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.GenreResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("account was successfully received")
	return &pb.GenreResponse{
		Id:   int64(entry.Id),
		Name: entry.Name,
	}, nil
}

func checkGenreRequest(in *pb.GenreRequest) error {
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
