package counter

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"
)

func TestCounterBuffer(t *testing.T) {
	expected := bytes.Count([]byte(testData), []byte("Go"))
	r := bytes.NewReader([]byte(testData))

	count, err := counterBuffer(r, []byte("Go"))
	if err != nil {
		if err != io.EOF {
			t.Fatal(err)
		}
	}

	t.Logf("%d\t%d", count, expected)
	if count != expected {
		t.Fatalf("count expected and given missmatch. given: `%d`, expected: `%d`", count, expected)
	}
}

func BenchmarkCunterBuffer(b *testing.B) {
	var crit = []byte("GO")

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		r := bytes.NewReader([]byte(testData))
		b.StartTimer()

		counterBuffer(r, crit)
	}
}

func TestExecutor(t *testing.T) {
	requests := make(RequestChan)

	responses := Executor(requests, 5, func(r *Request) (int, error) {
		response, err := http.DefaultClient.Get(r.Url)
		if err != nil {
			return 0, err
		}
		defer response.Body.Close()

		return 42, nil
	})

	go func() {
		for _, u := range testUrls {
			requests <- NewRequest(u, "GO")
		}
		close(requests)
	}()

	for r := range responses {
		log.Println(r)
	}
}
