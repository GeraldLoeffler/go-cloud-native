package main

import (
	"log"
	"time"
)

func main() {
	requestQueue := make(chan *Request, 5) // size determines number of concurrent requests in flight

	go simulateClientsSendingRequestsAndReceivingResponses(requestQueue)

	simulateServerReceivingRequestsAndSendingResponses(requestQueue)
}

type Request struct {
	input  int
	output chan int
}

func NewRequest(input int) *Request {
	return &Request{input, make(chan int)}
}

func simulateClientsSendingRequestsAndReceivingResponses(requestQueue chan *Request) {
	time.Sleep(100 * time.Millisecond)
	for n := 0; n < 10; n++ {
		go simulateClientSendingRequestAndReceivingResponse(n, requestQueue)
	}
	time.Sleep(1 * time.Second)
	close(requestQueue) // no more requests
}

func simulateClientSendingRequestAndReceivingResponse(input int, requestQueue chan *Request) {
	requ := NewRequest(input)
	log.Print("Client sending request with ", input)
	requestQueue <- requ
	go func() {
		output := <-requ.output
		log.Print("Client received response with output for ", requ.input, " as ", output)
	}()
}

func simulateServerReceivingRequestsAndSendingResponses(requestQueue chan *Request) {
	for r := range requestQueue { // accept all requests until no more coming
		go handleRequest(r)
	}
}

func handleRequest(r *Request) {
	log.Print("Handling request")
	defer log.Print("Handled request")
	input := r.input
	output := processRequest(input)
	r.output <- output
}

func processRequest(input int) (output int) {
	log.Print("Processing request with input ", input)
	defer func() { log.Print("Processed ", input, " to ", output) }()
	output = calculateOutput(input)
	return
}

func calculateOutput(input int) (output int) {
	log.Print("Calculating output for ", input)
	defer func() { log.Print("Calculated output for ", input, " as ", output) }()
	output = input * 2
	return
}
