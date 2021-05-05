package api

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	//pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
	pb "golang-training-theater-grpc/proto/go_proto"
)

type ScheduleServer struct {
	data *data.ScheduleData
}

func NewScheduleServer(a data.ScheduleData) *ScheduleServer {
	return &ScheduleServer{data: &a}
}

func (s ScheduleServer) CreateSchedule(ctx context.Context, in *pb.ScheduleRequest) (*pb.IdScheduleResponse, error) {
	entity := data.Schedule{
		AccountId:     int(in.GetAccountId()),
		PerformanceId: int(in.GetPerformanceId()),
		Date:          in.GetDate(),
		HallId:        int(in.GetHallId()),
	}
	id, err := s.data.AddSchedule(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"schedule": entity,
		}).Warningf("got an error when tried to create schedule: %s", err)
		return &pb.IdScheduleResponse{Id: -1}, fmt.Errorf("got an error when tried to create schedule: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"schedule": entity,
	}).Info("create schedule")
	return &pb.IdScheduleResponse{Id: int64(id)}, nil
}

func (s ScheduleServer) DeleteSchedule(ctx context.Context, in *pb.IdScheduleRequest) (*pb.StatusScheduleResponse, error) {
	entity := new(data.Schedule)
	entity.Id = int(in.Id)
	err := s.data.DeleteSchedule(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete schedule: %s", err)
		return &pb.StatusScheduleResponse{Message: "got an error when tried to delete schedule"},
			fmt.Errorf("got an error when tried to delete schedule: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("schedule deletion was successful")
	return &pb.StatusScheduleResponse{Message: "schedule deletion was successful"}, nil
}

func (s ScheduleServer) UpdateSchedule(ctx context.Context, in *pb.ScheduleRequest) (*pb.StatusScheduleResponse, error) {
	entity := data.Schedule{
		Id:            int(in.GetId()),
		AccountId:     int(in.GetAccountId()),
		PerformanceId: int(in.GetPerformanceId()),
		Date:          in.GetDate(),
		HallId:        int(in.GetHallId()),
	}
	err := s.data.UpdateSchedule(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"schedule": entity,
		}).Warningf("got an error when tried to update schedule: %s", err)
		return &pb.StatusScheduleResponse{Message: "got an error when tried to update schedule"},
			fmt.Errorf("got an error when tried to update schedule: %w", err)
	}
	log.WithFields(log.Fields{
		"schedule": entity,
	}).Info("schedule update was successful")
	return &pb.StatusScheduleResponse{Message: "schedule update was successful"}, nil
}
func (s ScheduleServer) GetSchedule(ctx context.Context, in *pb.IdScheduleRequest) (*pb.ScheduleResponse, error) {
	entity := new(data.Schedule)
	entity.Id = int(in.Id)
	entry, err := s.data.FindByIdSchedule(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get schedule: %s", err)
		return &pb.ScheduleResponse{},
			fmt.Errorf("got an error when tried to get schedule: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("schedule was successfully received")
	return &pb.ScheduleResponse{
		Id:            int64(entry.Id),
		AccountId:     int64(entry.AccountId),
		PerformanceId: int64(entry.PerformanceId),
		Date:          entry.Date,
		HallId:        int64(entry.HallId),
	}, nil
}
