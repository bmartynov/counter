package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/bmartynov/counter/counter"
	"github.com/bmartynov/counter/pb"
)

const (
	envKeyPort        = "PORT"
	envKeyConcurrency = "CONCURRENCY"
)

type response struct {
	summary *counter.Summary
	err     error
}

func newResponse(summary *counter.Summary, err error) response {
	return response{summary, err}
}

func execute(
	concurrency int,
	srv pb.Counter_CountServer,
	requestChan counter.RequestChan,
	summaryChan chan response,
) {

	summaryChan <- newResponse(counter.Collector(
		counter.Executor(
			requestChan,
			concurrency,
			counter.HttpSource),

		func(r *counter.Response) error {
			response := pb.NewResponse(r.Request.Url, r.Request.Criteria, r.Count, r.Error)

			return srv.Send(response)
		},
	))
}

func feeder(srv pb.Counter_CountServer, requestChan counter.RequestChan) {
	for {
		req, err := srv.Recv()
		if err != nil {
			close(requestChan)
			return
		}

		requestChan <- counter.NewRequest(req.Url, req.Criteria)
	}
}

type counterService int

func (s *counterService) Count(cs pb.Counter_CountServer) (err error) {
	var startTime = time.Now()

	var requestChan = make(counter.RequestChan)
	var summaryChan = make(chan response)

	go execute((int)(*s), cs, requestChan, summaryChan)

	go feeder(cs, requestChan)

	summary := <-summaryChan

	elapsedSec := time.Now().Sub(startTime).Seconds()

	response := pb.NewSummaryResponse(elapsedSec,
		summary.summary.Count,
		summary.summary.Failed,
		summary.summary.Success,
		summary.err)

	return cs.Send(response)
}

func NewService(concurrency int) pb.CounterServer {
	s := counterService(concurrency)

	return &s
}

func main() {
	port := os.Getenv(envKeyPort)
	if port == "" {
		log.Fatal("PORT environment variable was not set")
	}

	concurrency, err := strconv.Atoi(
		os.Getenv(envKeyConcurrency))

	if err != nil {
		log.Fatal("CONCURRENCY environment variable was not set or incorrect")
	}



	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	grpcServer := grpc.NewServer()

	pb.RegisterCounterServer(
		grpcServer,
		NewService(concurrency),
	)

	err = grpcServer.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
