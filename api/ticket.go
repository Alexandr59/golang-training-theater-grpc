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

type TicketServer struct {
	data *data.TicketData
}

func NewTicketServer(a data.TicketData) *TicketServer {
	return &TicketServer{data: &a}
}

func (t TicketServer) CreateTicket(ctx context.Context, in *pb.TicketRequest) (*pb.IdTicketResponse, error) {
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
		return &pb.IdTicketResponse{Id: -1}, fmt.Errorf("got an error when tried to create ticket: %w", err)
	}
	entity.Id = id
	log.WithFields(log.Fields{
		"ticket": entity,
	}).Info("create ticket")
	return &pb.IdTicketResponse{Id: int64(id)}, nil
}

func (t TicketServer) DeleteTicket(ctx context.Context, in *pb.IdTicketRequest) (*pb.StatusTicketResponse, error) {
	entity := new(data.Ticket)
	entity.Id = int(in.Id)
	err := t.data.DeleteTicket(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to delete ticket: %s", err)
		return &pb.StatusTicketResponse{Message: "got an error when tried to delete ticket"},
			fmt.Errorf("got an error when tried to delete ticket: %w", err)
	}
	log.WithFields(log.Fields{
		"id": entity.Id,
	}).Info("ticket deletion was successful")
	return &pb.StatusTicketResponse{Message: "ticket deletion was successful"}, nil
}

func (t TicketServer) UpdateTicket(ctx context.Context, in *pb.TicketRequest) (*pb.StatusTicketResponse, error) {
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
		return &pb.StatusTicketResponse{Message: "got an error when tried to update ticket"},
			fmt.Errorf("got an error when tried to update ticket: %w", err)
	}
	log.WithFields(log.Fields{
		"ticket": entity,
	}).Info("ticket update was successful")
	return &pb.StatusTicketResponse{Message: "ticket update was successful"}, nil
}

func (t TicketServer) GetTicket(ctx context.Context, in *pb.IdTicketRequest) (*pb.TicketResponse, error) {
	entity := new(data.Ticket)
	entity.Id = int(in.Id)
	entry, err := t.data.FindByIdTicket(*entity)
	if err != nil {
		log.WithFields(log.Fields{
			"id": entity.Id,
		}).Warningf("got an error when tried to get ticket: %s", err)
		return &pb.TicketResponse{},
			fmt.Errorf("got an error when tried to get ticket: %w", err)
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
		return &pb.JsonTicketsResponse{Json: ""},
			fmt.Errorf("got an error when tried to get tickets: %w", err)
	}
	json, err := json.Marshal(tickets)
	if err != nil {
		log.WithFields(log.Fields{
			"tickets": tickets,
		}).Warningf("got an error when tried to get json tickets: %s", err)
		return &pb.JsonTicketsResponse{Json: ""},
			fmt.Errorf("got an error when tried to get json tickets: %w", err)
	}
	return &pb.JsonTicketsResponse{Json: string(json)}, nil
}
