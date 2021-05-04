package api

import (
	"context"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"

	"golang-training-theater-grpc/pkg/data"
	pb "golang-training-theater-grpc/proto/go_proto"
)

func RegisterAllServiceServer(server *grpc.Server, conn *gorm.DB) {
	pb.RegisterAccountServiceServer(server, NewAccountServer(*data.NewAccountData(conn)))
	pb.RegisterGenreServiceServer(server, NewGenreServer(*data.NewGenreData(conn)))
	pb.RegisterHallServiceServer(server, NewHallServer(*data.NewHallData(conn)))
	pb.RegisterLocationServiceServer(server, NewLocationServer(*data.NewLocationData(conn)))
	pb.RegisterPerformanceServiceServer(server, NewPerformanceServer(*data.NewPerformanceData(conn)))
	pb.RegisterPlaceServiceServer(server, NewPlaceServer(*data.NewPlaceData(conn)))
	pb.RegisterPosterServiceServer(server, NewPosterServer(*data.NewPosterData(conn)))
	pb.RegisterPriceServiceServer(server, NewPriceServer(*data.NewPriceData(conn)))
	pb.RegisterRoleServiceServer(server, NewRoleServer(*data.NewRoleData(conn)))
	pb.RegisterScheduleServiceServer(server, NewScheduleServer(*data.NewScheduleData(conn)))
	pb.RegisterSectorServiceServer(server, NewSectorServer(*data.NewSectorData(conn)))
	pb.RegisterTicketServiceServer(server, NewTicketServer(*data.NewTicketData(conn)))
	pb.RegisterUserServiceServer(server, NewUserServer(*data.NewUserData(conn)))
}

func RegisterAllServiceHandler(background context.Context, grpcMux *runtime.ServeMux, conn *grpc.ClientConn) {
	err := pb.RegisterAccountServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterGenreServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterHallServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterLocationServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterPerformanceServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterPlaceServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterPosterServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterPriceServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterRoleServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterScheduleServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterSectorServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterTicketServiceHandler(background, grpcMux, conn)
	fatal(err)
	err = pb.RegisterUserServiceHandler(background, grpcMux, conn)
	fatal(err)
}

func fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
