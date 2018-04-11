package pb

func NewResponse(url, criteria string, cnt int, err error) *Response {
	r := &CountResponse{
		Request: &Request{url, criteria},
		Count:   uint32(cnt),
	}

	if err != nil {
		r.Error = err.Error()
	}

	return &Response{&Response_Count{r}}
}

func NewSummaryResponse(sec float64, count, failed, success int, err error) *Response {
	r := &Summary{
		SecondsElapsed: float32(sec),
		Count:          uint32(count),
		Failed:         uint32(failed),
		Success:        uint32(success),
	}

	if err != nil {
		r.Error = err.Error()
	}

	return &Response{&Response_Summary{r}}
}
