package tracing

type Request struct {
	input  int
	output chan int
}

func NewRequest(input int) *Request {
	return &Request{input, make(chan int)}
}
