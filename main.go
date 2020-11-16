package main

import (
	"github.com/Cheep2Workshop/MyMathFunction/mmf"
)

const (
	queryServiceAddress = "open-match-query.open-match.svc.cluster.local:50503"
	serverPort          = 50502
)

func main() {
	mmf.Start(queryServiceAddress, serverPort)
}
