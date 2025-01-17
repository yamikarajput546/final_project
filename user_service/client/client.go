//Use only for testing purposes to test the server without a gatway client
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jcromanu/final_project/user_service/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	createUserReq := &pb.CreateUserRequest{User: &pb.User{Name: "juan", Age: 31, PwdHash: "ooooo", Parent: []string{}, AdditionalInformation: "informacion que cura"}}
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatal("Error on simple rpc ", err)
	}
	defer conn.Close()
	client := pb.NewUserServiceClient(conn)
	response, err := client.CreateUser(context.Background(), createUserReq)
	if err != nil {
		log.Fatal("Error on simple rpc ", err)
	}
	fmt.Print("User id:", response.User.Id)
}
