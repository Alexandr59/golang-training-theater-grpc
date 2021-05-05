package api

import (
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	//pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
	pb "golang-training-theater-grpc/proto/go_proto"
)

type PosterServer struct {
	data *data.PosterData
}

func NewPosterServer(a data.PosterData) *PosterServer {
	return &PosterServer{data: &a}
}

func (p PosterServer) CreatePoster(ctx context.Context, in *pb.PosterRequest) (*pb.IdPosterResponse, error) {
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
		return &pb.IdPosterResponse{Id: -1}, fmt.Errorf("got an error when tried to create poster: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"poster": entity,
	}).Info("create poster")
	return &pb.IdPosterResponse{Id: int64(id)}, nil
}

func (p PosterServer) DeletePoster(ctx context.Context, in *pb.IdPosterRequest) (*pb.StatusPosterResponse, error) {
	entity := new(data.Poster)
	entity.Id = int(in.Id)
	err := p.data.DeletePoster(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete poster: %s", err)
		return &pb.StatusPosterResponse{Message: "got an error when tried to delete poster"},
			fmt.Errorf("got an error when tried to delete poster: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("poster deletion was successful")
	return &pb.StatusPosterResponse{Message: "poster deletion was successful"}, nil
}

func (p PosterServer) UpdatePoster(ctx context.Context, in *pb.PosterRequest) (*pb.StatusPosterResponse, error) {
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
		return &pb.StatusPosterResponse{Message: "got an error when tried to update poster"},
			fmt.Errorf("got an error when tried to update poster: %w", err)
	}
	log.WithFields(log.Fields{
		"poster": entity,
	}).Info("poster update was successful")
	return &pb.StatusPosterResponse{Message: "poster update was successful"}, nil
}

func (p PosterServer) GetPoster(ctx context.Context, in *pb.IdPosterRequest) (*pb.PosterResponse, error) {
	entity := new(data.Poster)
	entity.Id = int(in.Id)
	entry, err := p.data.FindByIdPoster(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get poster: %s", err)
		return &pb.PosterResponse{},
			fmt.Errorf("got an error when tried to get poster: %w", err)
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
		}).Warningf("got an error when tried to get poster: %s", err)
		return &pb.JsonResponse{Json: ""},
			fmt.Errorf("got an error when tried to get poster: %w", err)
	}
	json, err := json.Marshal(posters)
	if err != nil {
		log.WithFields(log.Fields{
			"posters": posters,
		}).Warningf("got an error when tried to get json posters: %s", err)
		return &pb.JsonResponse{Json: ""},
			fmt.Errorf("got an error when tried to get json poster: %w", err)
	}
	return &pb.JsonResponse{Json: string(json)}, nil
}
