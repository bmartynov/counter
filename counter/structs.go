package counter

import "fmt"

type Request struct {
	Url      string
	Criteria string
}

func (r *Request) String() string {
	return fmt.Sprintf("%s(%s)", r.Url, r.Criteria)
}

type Response struct {
	Request *Request
	Count   int
	Error   error
}

func (r *Response) String() string {
	if r.Error == nil {
		return fmt.Sprintf("[%s] -> Ok(%d)", r.Request, r.Count)
	}
	return fmt.Sprintf("[%s] -> Err(%s)", r.Request, r.Error)
}

type Summary struct {
	Count   int
	Failed  int
	Success int
}

func (s *Summary) String() string {
	return fmt.Sprintf("Count: `%d`, Failed: `%d`, Success: `%d`", s.Count, s.Failed, s.Success)
}

type RequestChan chan *Request
type ResponseChan chan *Response

func NewRequest(u string, criteria string) *Request {
	return &Request{Url: u, Criteria: criteria}
}

func NewResponse(r *Request, err error, count int) *Response {
	return &Response{r, count, err}
}

