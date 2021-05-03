package api

import (
	"context"
	"fmt"
	"log"

	//"github.com/Alexandr59/golang-training-theater-grpc/pkg/data"
	"golang-training-theater-grpc/pkg/data"
	pb "golang-training-theater-grpc/proto/go_proto"
)

//type accountAPI struct {
//	data *data.AccountData
//}

type AccountServer struct {
	data *data.AccountData
}

func NewAccountServer(a data.AccountData) *AccountServer {
	return &AccountServer{data: &a}
}

func (a AccountServer) CreateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.AccountResponse, error) {
	fmt.Println("Create")
	log.Printf(in.Account)
	//entity := new(data.Account)
	////err := json.NewDecoder(request.Body).Decode(&entity)
	//err := json.NewDecoder(request.GetAccount()).Decode(&entity)
	//if err != nil {
	//	log.Printf("failed reading JSON: %s\n", err)
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//if entity == nil {
	//	log.Printf("failed empty JSON\n")
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//id, err := a.data.AddAccount(*entity)
	//if err != nil {
	//	_, err := writer.Write([]byte("got an error when tried to create account"))
	//	if err != nil {
	//		log.Println(err)
	//		writer.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//}
	//err = json.NewEncoder(writer).Encode(id)
	//if err != nil {
	//	log.Println(err)
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//writer.WriteHeader(http.StatusCreated)
	return &pb.AccountResponse{Account: "Hello Create YYYYYYYYYYYYYYYYYYYYYY"}, nil
}

func (a AccountServer) DeleteAccount(ctx context.Context, in *pb.IdRequest) (*pb.AccountResponse, error) {
	fmt.Println("Delete")
	log.Printf(string(in.Id))
	//entity := new(data.Account)
	//if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
	//	entity.Id = n
	//} else {
	//	log.Printf("failed reading id: %s", err)
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//err := a.data.DeleteAccount(*entity)
	//if err != nil {
	//	_, err := writer.Write([]byte("got an error when tried to delete account"))
	//	if err != nil {
	//		log.Println(err)
	//		writer.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//}
	return &pb.AccountResponse{Account: "Hello Delete YYYYYYYYYYYYYYYYYYYYYY"}, nil
}

func (a AccountServer) UpdateAccount(ctx context.Context, in *pb.AccountRequest) (*pb.AccountResponse, error) {
	fmt.Println("Update")
	log.Printf(in.Account)
	//entity := new(data.Account)
	//err := json.NewDecoder(request.Body).Decode(&entity)
	//if err != nil {
	//	log.Printf("failed reading JSON: %s\n", err)
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//if entity == nil {
	//	log.Printf("failed empty JSON\n")
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//err = a.data.UpdateAccount(*entity)
	//if err != nil {
	//	_, err := writer.Write([]byte("got an error when tried to update account"))
	//	if err != nil {
	//		log.Println(err)
	//		writer.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//}
	return &pb.AccountResponse{Account: "Hello Update YYYYYYYYYYYYYYYYYYYYYY"}, nil
}

func (a AccountServer) GetAccount(ctx context.Context, in *pb.IdRequest) (*pb.AccountResponse, error) {
	fmt.Println("Get")
	log.Printf(string(in.Id))
	//entity := new(data.Account)
	//if n, err := strconv.Atoi(request.URL.Query().Get("id")); err == nil {
	//	entity.Id = n
	//} else {
	//	log.Printf("failed reading id: %s", err)
	//	writer.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//entry, err := a.data.FindByIdAccount(*entity)
	//if err != nil {
	//	_, err := writer.Write([]byte("got an error when tried to get account"))
	//	if err != nil {
	//		log.Println(err)
	//		writer.WriteHeader(http.StatusInternalServerError)
	//		return
	//	}
	//}
	//err = json.NewEncoder(writer).Encode(entry)
	//if err != nil {
	//	log.Println(err)
	//	writer.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	return &pb.AccountResponse{Account: "Hello Get YYYYYYYYYYYYYYYYYYYYYY"}, nil
}
