package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type LocationServer struct {
	data *data.LocationData
}

func NewLocationServer(a data.LocationData) *LocationServer {
	return &LocationServer{data: &a}
}

func (l LocationServer) CreateLocation(ctx context.Context, in *pb.LocationRequest) (*pb.IdLocationResponse, error) {
	if err := checkLocationRequest(in); err != nil {
		log.WithFields(log.Fields{
			"location": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdLocationResponse{Id: -1}, err
	}
	entity := data.Location{
		AccountId:   int(in.GetAccountId()),
		Address:     in.GetAddress(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	id, err := l.data.AddLocation(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"location": entity,
		}).Warningf("got an error when tried to create location: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create location: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdLocationResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdLocationResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"location": entity,
	}).Info("location has been successfully created")
	return &pb.IdLocationResponse{Id: int64(id)}, nil
}

func (l LocationServer) DeleteLocation(ctx context.Context, in *pb.IdLocationRequest) (*pb.StatusLocationResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"location": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusLocationResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Location)
	entity.Id = int(in.Id)
	err := l.data.DeleteLocation(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete location: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete location: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusLocationResponse{Message: "got an error when tried to delete location"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusLocationResponse{Message: "got an error when tried to delete location"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("location deletion was successful")
	return &pb.StatusLocationResponse{Message: "location deletion was successful"}, nil
}

func (l LocationServer) UpdateLocation(ctx context.Context, in *pb.LocationRequest) (*pb.StatusLocationResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"location": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusLocationResponse{Message: "empty fields error"}, err
	}
	if err := checkLocationRequest(in); err != nil {
		log.WithFields(log.Fields{
			"location": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusLocationResponse{Message: "empty fields error"}, err
	}
	entity := data.Location{
		Id:          int(in.GetId()),
		AccountId:   int(in.GetAccountId()),
		Address:     in.GetAddress(),
		PhoneNumber: in.GetPhoneNumber(),
	}
	err := l.data.UpdateLocation(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"location": entity,
		}).Warningf("got an error when tried to update location: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update location: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusLocationResponse{Message: "got an error when tried to update location"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusLocationResponse{Message: "got an error when tried to update location"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"location": entity,
	}).Info("location update was successful")
	return &pb.StatusLocationResponse{Message: "location update was successful"}, nil
}

func (l LocationServer) GetLocation(ctx context.Context, in *pb.IdLocationRequest) (*pb.LocationResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"location": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.LocationResponse{}, err
	}
	entity := new(data.Location)
	entity.Id = int(in.Id)
	entry, err := l.data.FindByIdLocation(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get location: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get location: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.LocationResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.LocationResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("account was successfully received")
	return &pb.LocationResponse{
		Id:          int64(entry.Id),
		AccountId:   int64(entry.AccountId),
		Address:     entry.Address,
		PhoneNumber: entry.PhoneNumber,
	}, nil
}

func checkLocationRequest(in *pb.LocationRequest) error {
	if in.GetAccountId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {AccountId}: %s", in.GetAccountId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetAddress() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Address}: %s", in.GetAddress())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetPhoneNumber() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {PhoneNumber}: %s", in.GetPhoneNumber())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
