package main

import (
	"Hydra/hydracommlayer"
	"Hydra/hydracommlayer/hydraproto"
	"flag"
	"log"
	"strings"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8000", "address? host:port")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runServer(addr string)  {
	c := hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	recvChan, err := c.ListenAndDecode(addr)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range recvChan {
		log.Println("Received: ", msg)
	}
}

func runClient(addr string) {
	c := hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	ship := &hydraproto.Ship{
		Shipname:  "Hydra",
		CaptainName: "Jala",
		Crew: []*hydraproto.Ship_CrewMember {
			{1, "Kelvin", 5, "Pilot",  },
			{2, "Jade", 4, "Tech",  },
			{3, "Wally", 3, "Engineer",  },
		},
	}
}
