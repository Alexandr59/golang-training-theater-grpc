package api

import (
	"context"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	pb "github.com/Alexandr59/golang-training-theater-grpc/proto/go_proto"
)

type PriceServer struct {
	data *data.PriceData
}

func NewPriceServer(a data.PriceData) *PriceServer {
	return &PriceServer{data: &a}
}

func (p PriceServer) CreatePrice(ctx context.Context, in *pb.PriceRequest) (*pb.IdPriceResponse, error) {
	if err := checkPriceRequest(in); err != nil {
		log.WithFields(log.Fields{
			"price": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdPriceResponse{Id: -1}, err
	}
	entity := data.Price{
		AccountId:     int(in.GetAccountId()),
		SectorId:      int(in.GetSectorId()),
		PerformanceId: int(in.GetPerformanceId()),
		Price:         int(in.GetPrice()),
	}
	id, err := p.data.AddPrice(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"price": entity,
		}).Warningf("got an error when tried to create price: %s", err)
		s := status.Newf(codes.Canceled, "got an error when tried to create price: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdPriceResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdPriceResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"price": entity,
	}).Info("price has been successfully created")
	return &pb.IdPriceResponse{Id: int64(id)}, nil
}

func (p PriceServer) DeletePrice(ctx context.Context, in *pb.IdPriceRequest) (*pb.StatusPriceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"price": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPriceResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Price)
	entity.Id = int(in.Id)
	err := p.data.DeletePrice(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete price: %s", err)
		s := status.Newf(codes.Canceled, "got an error when tried to delete price: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusPriceResponse{Message: "got an error when tried to delete price"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusPriceResponse{Message: "got an error when tried to delete price"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("price deletion was successful")
	return &pb.StatusPriceResponse{Message: "price deletion was successful"}, nil
}

func (p PriceServer) UpdatePrice(ctx context.Context, in *pb.PriceRequest) (*pb.StatusPriceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"price": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPriceResponse{Message: "empty fields error"}, err
	}
	if err := checkPriceRequest(in); err != nil {
		log.WithFields(log.Fields{
			"price": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusPriceResponse{Message: "empty fields error"}, err
	}
	entity := data.Price{
		Id:            int(in.GetId()),
		AccountId:     int(in.GetAccountId()),
		SectorId:      int(in.GetSectorId()),
		PerformanceId: int(in.GetPerformanceId()),
		Price:         int(in.GetPrice()),
	}
	err := p.data.UpdatePrice(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"price": entity,
		}).Warningf("got an error when tried to update price: %s", err)
		s := status.Newf(codes.Canceled, "got an error when tried to update price: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusPriceResponse{Message: "got an error when tried to update price"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusPriceResponse{Message: "got an error when tried to update price"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"price": entity,
	}).Info("price update was successful")
	return &pb.StatusPriceResponse{Message: "price update was successful"}, nil
}

func (p PriceServer) GetPrice(ctx context.Context, in *pb.IdPriceRequest) (*pb.PriceResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"price": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.PriceResponse{}, err
	}
	entity := new(data.Price)
	entity.Id = int(in.Id)
	entry, err := p.data.FindByIdPrice(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get price: %s", err)
		s := status.Newf(codes.Canceled, "got an error when tried to get price: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.PriceResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.PriceResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("price was successfully received")
	return &pb.PriceResponse{
		Id:            int64(entry.Id),
		AccountId:     int64(entry.AccountId),
		SectorId:      int64(entry.SectorId),
		PerformanceId: int64(entry.PerformanceId),
		Price:         int64(entry.Price),
	}, nil
}

func checkPriceRequest(in *pb.PriceRequest) error {
	if in.GetAccountId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {AccountId}: %s", in.GetAccountId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetSectorId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {SectorId}: %s", in.GetSectorId())
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
	if in.GetPrice() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {Price}: %s", in.GetPrice())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
