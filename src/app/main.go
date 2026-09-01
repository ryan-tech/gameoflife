package main

import (
	"strings"
	"fmt"
	"gameOfLife/src/game"
	"net"
	"grpc"
)


func main() {

	// Starts gRPC server
	// lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	// if err != nil {
	// log.Fatalf("failed to listen: %v", err)
	// }
	// var opts []grpc.ServerOption
	
	// grpcServer := grpc.NewServer(opts...)
	// pb.RegisterRouteGuideServer(grpcServer, newServer())
	// grpcServer.Serve(lis)

	// myGrid := game.Grid{}
	// // Create grid for the game
	// myGrid.initGrid(10,10)

	// // 2x2 square -> should be stable
	// // myGrid.toggleCell(2,2)
	// // myGrid.toggleCell(2,3)
	// // myGrid.toggleCell(3,2)
	// // myGrid.toggleCell(3,3)

	// // 1x3 rectangle -> alternates
	// myGrid.toggleCell(3,2)
	// myGrid.toggleCell(3,3)
	// myGrid.toggleCell(3,4)
	// fmt.Println("Starting simulation")
	// myGrid.printGrid()

	// for i := 0; i < 5; i++ {
	// 	myGrid.incrementStep()
	// }
}