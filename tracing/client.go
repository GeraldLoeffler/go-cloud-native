package tracing

import (
	"log"
	"time"
)

func SimulateClientsSendingRequestsAndReceivingResponses(requestQueue chan *Request) {
	time.Sleep(100 * time.Millisecond)
	for n := 1; n <= 10; n++ {
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
