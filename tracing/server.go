package tracing

import (
	"context"
	"fmt"
	"log"

	"github.com/opentracing/opentracing-go"
)

const (
	tagKeyInput  = "input"
	tagKeyOutput = "output"
)

func ServerMain(requestQueue chan *Request) {
	log.Print("Initializing Open Tracing implementation")
	tracer, closer := CreateOpenTracingTracer("tracing-server")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer) // required for integration with context API

	log.Print("Server now accepting requests")
	simulateServerReceivingRequestsAndSendingResponses(requestQueue)
}

func simulateServerReceivingRequestsAndSendingResponses(requestQueue chan *Request) {
	for r := range requestQueue { // accept all requests until no more coming
		go handleRequest(r)
	}
}

func handleRequest(r *Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	span, ctx := opentracing.StartSpanFromContext(ctx, "handleRequest")
	defer span.Finish()
	span.SetTag(tagKeyInput, r.input)

	span.LogKV(logKeyEvent, logValEventBeforeCall, logKeyFunction, "processRequest")
	output := processRequest(ctx, r.input)
	span.LogKV(logKeyEvent, logValEventAfterReturn, logKeyFunction, "processRequest")
	span.SetTag(tagKeyOutput, output)

	r.output <- output
}

func processRequest(ctx context.Context, input int) (output int) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "processRequest")
	defer span.Finish()
	span.SetTag(tagKeyInput, input)

	span.LogKV(logKeyEvent, logValEventBeforeCall, logKeyFunction, "calculateOutput")
	output = calculateOutput(ctx, input)
	span.LogKV(logKeyEvent, logValEventAfterReturn, logKeyFunction, "calculateOutput")
	span.SetTag(tagKeyOutput, output)

	return
}

func calculateOutput(ctx context.Context, input int) (output int) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "calculateOutput")
	defer span.Finish()
	span.SetTag(tagKeyInput, input)

	output = input * 2
	span.LogKV(logKeyEvent, "calculated", logKeyMessage, fmt.Sprintf("Calculated from input %v output %v", input, output))
	span.SetTag(tagKeyOutput, output)

	return
}
