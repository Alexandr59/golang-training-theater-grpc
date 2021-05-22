package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type ScheduleServer struct {
	data *data.ScheduleData
}

func NewScheduleServer(a data.ScheduleData) *ScheduleServer {
	return &ScheduleServer{data: &a}
}

func (s ScheduleServer) CreateSchedule(ctx context.Context, in *pb.ScheduleRequest) (*pb.IdScheduleResponse, error) {
	if err := checkScheduleRequest(in); err != nil {
		log.WithFields(log.Fields{
			"schedule": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdScheduleResponse{Id: -1}, err
	}
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
		s := status.Newf(codes.Internal, "got an error when tried to create schedule: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdScheduleResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdScheduleResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"schedule": entity,
	}).Info("schedule has been successfully created")
	return &pb.IdScheduleResponse{Id: int64(id)}, nil
}

func (s ScheduleServer) DeleteSchedule(ctx context.Context, in *pb.IdScheduleRequest) (*pb.StatusScheduleResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"schedule": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusScheduleResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Schedule)
	entity.Id = int(in.Id)
	err := s.data.DeleteSchedule(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete schedule: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete schedule: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusScheduleResponse{Message: "got an error when tried to delete schedule"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusScheduleResponse{Message: "got an error when tried to delete schedule"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("schedule deletion was successful")
	return &pb.StatusScheduleResponse{Message: "schedule deletion was successful"}, nil
}

func (s ScheduleServer) UpdateSchedule(ctx context.Context, in *pb.ScheduleRequest) (*pb.StatusScheduleResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"schedule": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusScheduleResponse{Message: "empty fields error"}, err
	}
	if err := checkScheduleRequest(in); err != nil {
		log.WithFields(log.Fields{
			"schedule": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusScheduleResponse{Message: "empty fields error"}, err
	}
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
		s := status.Newf(codes.Internal, "got an error when tried to update schedule: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusScheduleResponse{Message: "got an error when tried to update schedule"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusScheduleResponse{Message: "got an error when tried to update schedule"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"schedule": entity,
	}).Info("schedule update was successful")
	return &pb.StatusScheduleResponse{Message: "schedule update was successful"}, nil
}

func (s ScheduleServer) GetSchedule(ctx context.Context, in *pb.IdScheduleRequest) (*pb.ScheduleResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"schedule": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.ScheduleResponse{}, err
	}
	entity := new(data.Schedule)
	entity.Id = int(in.Id)
	entry, err := s.data.FindByIdSchedule(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get schedule: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get schedule: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.ScheduleResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.ScheduleResponse{}, errWithDetails.Err()
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

func checkScheduleRequest(in *pb.ScheduleRequest) error {
	if in.GetAccountId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {AccountId}: %s", in.GetAccountId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetPerformanceId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {PerformanceId}: %s", in.GetPerformanceId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetDate() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Date}: %s", in.GetDate())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetHallId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {HallId}: %s", in.GetHallId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
