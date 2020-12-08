package main

import "github.com/GeraldLoeffler/go-cloud-native/tracing"

func main() {
	requestQueue := make(chan *tracing.Request, 5) // size determines number of concurrent requests in flight

	go tracing.SimulateClientsSendingRequestsAndReceivingResponses(requestQueue)

	tracing.SimulateServerReceivingRequestsAndSendingResponses(requestQueue)
}
