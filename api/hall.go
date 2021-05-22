package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type HallServer struct {
	data *data.HallData
}

func NewHallServer(a data.HallData) *HallServer {
	return &HallServer{data: &a}
}

func (h HallServer) CreateHall(ctx context.Context, in *pb.HallRequest) (*pb.IdHallResponse, error) {
	if err := checkHallRequest(in); err != nil {
		log.WithFields(log.Fields{
			"hall": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdHallResponse{Id: -1}, err
	}
	entity := data.Hall{
		AccountId:  int(in.GetAccountId()),
		Name:       in.GetName(),
		Capacity:   int(in.GetCapacity()),
		LocationId: int(in.GetLocationId()),
	}
	id, err := h.data.AddHall(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"hall": entity,
		}).Warningf("got an error when tried to create hall: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create hall: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdHallResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdHallResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"hall": entity,
	}).Info("hall has been successfully created")
	return &pb.IdHallResponse{Id: int64(id)}, nil
}

func (h HallServer) DeleteHall(ctx context.Context, in *pb.IdHallRequest) (*pb.StatusHallResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"hall": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusHallResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Hall)
	entity.Id = int(in.Id)
	err := h.data.DeleteHall(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete hall: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete account: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusHallResponse{Message: "got an error when tried to delete hall"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusHallResponse{Message: "got an error when tried to delete hall"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("hall deletion was successful")
	return &pb.StatusHallResponse{Message: "hall deletion was successful"}, nil
}

func (h HallServer) UpdateHall(ctx context.Context, in *pb.HallRequest) (*pb.StatusHallResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"hall": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusHallResponse{Message: "empty fields error"}, err
	}
	if err := checkHallRequest(in); err != nil {
		log.WithFields(log.Fields{
			"hall": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusHallResponse{Message: "empty fields error"}, err
	}
	entity := data.Hall{
		Id:         int(in.GetId()),
		AccountId:  int(in.GetAccountId()),
		Name:       in.GetName(),
		Capacity:   int(in.GetCapacity()),
		LocationId: int(in.GetLocationId()),
	}
	err := h.data.UpdateHall(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"hall": entity,
		}).Warningf("got an error when tried to update hall: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update hall: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusHallResponse{Message: "got an error when tried to update hall"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusHallResponse{Message: "got an error when tried to update hall"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"hall": entity,
	}).Info("hall update was successful")
	return &pb.StatusHallResponse{Message: "hall update was successful"}, nil
}

func (h HallServer) GetHall(ctx context.Context, in *pb.IdHallRequest) (*pb.HallResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"hall": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.HallResponse{}, err
	}
	entity := new(data.Hall)
	entity.Id = int(in.Id)
	entry, err := h.data.FindByIdHall(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get hall: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get hall: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.HallResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.HallResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("hall was successfully received")
	return &pb.HallResponse{
		Id:         int64(entry.Id),
		AccountId:  int64(entry.AccountId),
		Name:       entry.Name,
		Capacity:   int64(entry.Capacity),
		LocationId: int64(entry.LocationId),
	}, nil
}

func checkHallRequest(in *pb.HallRequest) error {
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
	if in.GetCapacity() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Capacity}: %s", in.GetCapacity())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetLocationId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {LocationId}: %s", in.GetLocationId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
