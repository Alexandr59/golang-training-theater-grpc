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

type TicketServer struct {
	data *data.TicketData
}

func NewTicketServer(a data.TicketData) *TicketServer {
	return &TicketServer{data: &a}
}

func (t TicketServer) CreateTicket(ctx context.Context, in *pb.TicketRequest) (*pb.IdTicketResponse, error) {
	if err := checkTicketRequest(in); err != nil {
		log.WithFields(log.Fields{
			"ticket": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.IdTicketResponse{Id: -1}, err
	}
	entity := data.Ticket{
		AccountId:   int(in.GetAccountId()),
		ScheduleId:  int(in.GetScheduleId()),
		PlaceId:     int(in.GetPlaceId()),
		DateOfIssue: in.GetDateOfIssue(),
		Paid:        in.GetPaid(),
		Reservation: in.GetReservation(),
		Destroyed:   in.GetDestroyed(),
	}
	id, err := t.data.AddTicket(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"ticket": entity,
		}).Warningf("got an error when tried to create ticket: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to create ticket: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.IdTicketResponse{Id: -1}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.IdTicketResponse{Id: -1}, errWithDetails.Err()
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"ticket": entity,
	}).Info("ticket has been successfully created")
	return &pb.IdTicketResponse{Id: int64(id)}, nil
}

func (t TicketServer) DeleteTicket(ctx context.Context, in *pb.IdTicketRequest) (*pb.StatusTicketResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"ticket": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusTicketResponse{Message: "empty fields error"}, err
	}
	entity := new(data.Ticket)
	entity.Id = int(in.Id)
	err := t.data.DeleteTicket(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete ticket: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to delete ticket: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusTicketResponse{Message: "got an error when tried to delete ticket"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusTicketResponse{Message: "got an error when tried to delete ticket"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("ticket deletion was successful")
	return &pb.StatusTicketResponse{Message: "ticket deletion was successful"}, nil
}

func (t TicketServer) UpdateTicket(ctx context.Context, in *pb.TicketRequest) (*pb.StatusTicketResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"ticket": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusTicketResponse{Message: "empty fields error"}, err
	}
	if err := checkTicketRequest(in); err != nil {
		log.WithFields(log.Fields{
			"ticket": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.StatusTicketResponse{Message: "empty fields error"}, err
	}
	entity := data.Ticket{
		Id:          int(in.GetId()),
		AccountId:   int(in.GetAccountId()),
		ScheduleId:  int(in.GetScheduleId()),
		PlaceId:     int(in.GetPlaceId()),
		DateOfIssue: in.GetDateOfIssue(),
		Paid:        in.GetPaid(),
		Reservation: in.GetReservation(),
		Destroyed:   in.GetDestroyed(),
	}
	err := t.data.UpdateTicket(entity)
	if err != nil {
		log.WithFields(log.Fields{
			"ticket": entity,
		}).Warningf("got an error when tried to update ticket: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to update ticket: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.StatusTicketResponse{Message: "got an error when tried to update ticket"},
				status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.StatusTicketResponse{Message: "got an error when tried to update ticket"}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"ticket": entity,
	}).Info("ticket update was successful")
	return &pb.StatusTicketResponse{Message: "ticket update was successful"}, nil
}

func (t TicketServer) GetTicket(ctx context.Context, in *pb.IdTicketRequest) (*pb.TicketResponse, error) {
	if err := checkId(in.GetId()); err != nil {
		log.WithFields(log.Fields{
			"ticket": in,
		}).Warningf("empty fields error: %s", err)
		return &pb.TicketResponse{}, err
	}
	entity := new(data.Ticket)
	entity.Id = int(in.Id)
	entry, err := t.data.FindByIdTicket(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get ticket: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get ticket: %s, with error: %w", in, err)
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return &pb.TicketResponse{}, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return &pb.TicketResponse{}, errWithDetails.Err()
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("ticket was successfully received")
	return &pb.TicketResponse{
		Id:          int64(entry.Id),
		AccountId:   int64(entry.AccountId),
		ScheduleId:  int64(entry.ScheduleId),
		PlaceId:     int64(entry.PlaceId),
		DateOfIssue: entry.DateOfIssue,
		Paid:        entry.Paid,
		Reservation: entry.Reservation,
		Destroyed:   entry.Destroyed,
	}, nil
}

func (t TicketServer) GetAllTickets(ctx context.Context, in *pb.TicketsRequest) (*pb.JsonTicketsResponse, error) {
	tickets, err := t.data.ReadAllTickets()
	if err != nil {
		log.WithFields(log.Fields{
			"tickets": tickets,
		}).Warningf("got an error when tried to get tickets: %s", err)
		s := status.Newf(codes.Internal, "got an error when tried to get tickets: %w", err)
		return &pb.JsonTicketsResponse{Json: ""}, s.Err()
	}
	json, err := json.Marshal(tickets)
	if err != nil {
		log.WithFields(log.Fields{
			"tickets": tickets,
		}).Warningf("got an error when tried to get json tickets: %s", err)
		s := status.Newf(codes.Unknown, "got an error when tried to get json tickets: %w", err)
		return &pb.JsonTicketsResponse{Json: ""}, s.Err()
	}
	return &pb.JsonTicketsResponse{Json: string(json)}, nil
}

func checkTicketRequest(in *pb.TicketRequest) error {
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
	if in.GetPlaceId() <= 0 {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {PlaceId}: %s", in.GetPlaceId())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	if in.GetDateOfIssue() == "" {
		s := status.Newf(codes.InvalidArgument, "didn't specify the field {DateOfIssue}: %s", in.GetDateOfIssue())
		errWithDetails, err := s.WithDetails(in)
		if err != nil {
			return status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return errWithDetails.Err()
	}
	return nil
}
