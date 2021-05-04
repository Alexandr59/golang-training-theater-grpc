package api

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"
	pb "golang-training-theater-grpc/proto/go_proto"
	//"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	"golang-training-theater-grpc/pkg/data"
)

type GenreServer struct {
	data *data.GenreData
}

func NewGenreServer(a data.GenreData) *GenreServer {
	return &GenreServer{data: &a}
}

func (g GenreServer) CreateGenre(ctx context.Context, in *pb.GenreRequest) (*pb.IdGenreResponse, error) {
	entity := data.Genre{
		Name: in.GetName(),
	}
	id, err := g.data.AddGenre(entity)
	if err != nil {
		//_, err := writer.Write([]byte("got an error when tried to create genre"))
		log.WithFields(log.Fields{
			"genre": entity,
		}).Warningf("got an error when tried to create genre: %s", err)
		return &pb.IdGenreResponse{Id: -1}, fmt.Errorf("got an error when tried to create genre: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"genre": entity,
	}).Info("create genre")
	return &pb.IdGenreResponse{Id: int64(id)}, nil
}

func (g GenreServer) DeleteGenre(ctx context.Context, in *pb.IdGenreRequest) (*pb.StatusGenreResponse, error) {
	entity := new(data.Genre)
	entity.Id = int(in.GetId())
	err := g.data.DeleteGenre(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete genre: %s", err)
		return &pb.StatusGenreResponse{Message: "got an error when tried to delete genre"},
			fmt.Errorf("got an error when tried to delete genre: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("genre deletion was successful")
	return &pb.StatusGenreResponse{Message: "genre deletion was successful"}, nil
}

func (g GenreServer) UpdateGenre(ctx context.Context, in *pb.GenreRequest) (*pb.StatusGenreResponse, error) {
	entity := data.Genre{
		Id:   int(in.GetId()),
		Name: in.GetName(),
	}
	err := g.data.UpdateGenre(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"genre": entity,
		}).Warningf("got an error when tried to update genre: %s", err)
		return &pb.StatusGenreResponse{Message: "got an error when tried to update genre"},
			fmt.Errorf("got an error when tried to update genre: %w", err)
	}
	log.WithFields(log.Fields{
		"genre": entity,
	}).Info("genre update was successful")
	return &pb.StatusGenreResponse{Message: "genre update was successful"}, nil
}

func (g GenreServer) GetGenre(ctx context.Context, in *pb.IdGenreRequest) (*pb.GenreResponse, error) {
	entity := new(data.Genre)
	entity.Id = int(in.Id)
	entry, err := g.data.FindByIdGenre(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get genre: %s", err)
		return &pb.GenreResponse{},
			fmt.Errorf("got an error when tried to get genre: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("account was successfully received")
	return &pb.GenreResponse{
		Id:   int64(entry.Id),
		Name: entry.Name,
	}, nil
}
