package tracing

import (
	"context"
	"log"
)

func SimulateServerReceivingRequestsAndSendingResponses(requestQueue chan *Request) {
	for r := range requestQueue { // accept all requests until no more coming
		go handleRequest(r)
	}
}

func handleRequest(r *Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Print("Handling request")
	defer log.Print("Handled request")

	input := r.input
	output := processRequest(ctx, input)
	r.output <- output
}

func processRequest(ctx context.Context, input int) (output int) {
	ctx = context.WithValue(ctx, "input", input)

	log.Print("Processing request with input ", input)
	defer func() { log.Print("Processed ", input, " to ", output) }()

	output = calculateOutput(ctx, input)
	return
}

func calculateOutput(ctx context.Context, input int) (output int) {
	log.Print("Calculating output for ", input)
	defer func() { log.Print("Calculated output for ", input, " as ", output) }()

	output = input * 2
	return
}
