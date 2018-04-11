package main

import (
	"os"
	"log"
	"flag"
	"github.com/bmartynov/counter/counter"
)

var concurrency = flag.Int("concurrency", 5, "concurrency")
var criteria = flag.String("criteria", "Go", "criteria to find")

func init() {
	flag.Parse()
}

func main() {
	s, err := counter.Collector(
		counter.Executor(
			counter.Producer(os.Stdin, *criteria),
			*concurrency,
			counter.HttpSource,
		),
		func(r *counter.Response) (err error) {
			log.Printf("Count for %s: %d", r.Request.Url, r.Count)

			return
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Total: %d", s.Count)
}
