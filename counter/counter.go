package counter

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"net/http"
	"sync"
)

const (
	defaultBufSize = 64
)

type HandlerFn func(request *Request) (int, error)

func counterBuffer(r io.Reader, criteria []byte) (count int, err error) {
	var buf [defaultBufSize]byte
	var sz, cnt, tailIdx int
	var cLen = len(criteria)

	for {
		sz, err = r.Read(buf[:])
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			return
		}

		cnt = bytes.Count(buf[:sz], criteria)
		if cnt > 0 {
			count += cnt
			tailIdx = bytes.LastIndex(buf[:sz], criteria) + cLen
		} else {
			tailIdx = sz - cLen
		}

		copy(buf[:sz-tailIdx], buf[tailIdx:sz])
	}

	return
}

func Counter(body io.Reader, criteria []byte) (int, error) {
	return counterBuffer(body, criteria)
}

func Producer(src io.Reader, criteria string) (requests RequestChan) {
	requests = make(RequestChan)

	go func(r *bufio.Reader) {
		for {
			u, err := r.ReadString('\n')
			if err != nil {
				close(requests)
				break
			}
			requests <- NewRequest(strings.TrimSpace(u), criteria)
		}
	}(bufio.NewReader(src))

	return
}

func Executor(
	requests RequestChan,
	concurrency int,
	handler HandlerFn,
) (rChan ResponseChan) {
	rChan = make(ResponseChan)

	go func() {
		wg := sync.WaitGroup{}
		guard := make(chan struct{}, concurrency)

		for {
			select {
			case r, ok := <-requests:
				if !ok {
					wg.Wait()
					close(rChan)
					return
				}

				guard <- struct{}{}
				wg.Add(1)

				go func() {
					defer func() {
						wg.Done()
						<-guard
					}()

					count, err := handler(r)

					rChan <- NewResponse(r, err, count)
				}()
			}
		}
	}()

	return rChan
}

func Collector(responses ResponseChan, cb func(r *Response) error) (rs *Summary, err error) {
	rs = &Summary{}

	for {
		select {
		case r, ok := <-responses:
			if !ok {
				return
			}

			if err = cb(r); err != nil {
				return
			}

			if r.Error == nil {
				rs.Count += r.Count
				rs.Success++
			} else {
				rs.Failed++
			}
		}
	}
}

func HttpSource(r *Request) (int, error) {
	response, err := http.DefaultClient.Get(r.Url)
	if err != nil {
		return 0, err
	}

	defer response.Body.Close()

	return Counter(response.Body, []byte(r.Criteria))
}
