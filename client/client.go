package main

import (
	"context"
	"flag"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/bmartynov/counter/counter"
	"github.com/bmartynov/counter/pb"
)

var address = flag.String("address", "localhost:8888", "address")
var criteria = flag.String("criteria", "Go", "criteria to find")

func init() {
	flag.Parse()
}

func main() {
	counterConn, err := grpc.Dial(*address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot establish clientConn: `%s`", err)
	}
	defer counterConn.Close()

	cs, err := pb.NewCounterClient(counterConn).
		Count(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	go feeder(counter.Producer(os.Stdin, *criteria), cs)

	for {
		response, err := cs.Recv()
		if err != nil {
			return
		}
		log.Println(response.String())
	}
}

//feeder - send requests to counter service
func feeder(requests counter.RequestChan, cClient pb.Counter_CountClient) {
	for {
		select {
		case req := <-requests:
			if req == nil {
				cClient.CloseSend()
				return
			}

			if err := cClient.Send(&pb.Request{
				Url:      req.Url,
				Criteria: req.Criteria,
			}); err != nil {
				log.Fatal(err)
			}
		}
	}
}
