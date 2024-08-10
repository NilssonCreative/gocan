package can

import (
	"log"

	"github.com/angelodlfrtr/go-can"
	"github.com/angelodlfrtr/go-can/transports"
)

type CanBus struct {
	*can.Bus
}

func Connect(port string, baudrate int) CanBus {
	// Configure transport
	tr := &transports.USBCanAnalyzer{
		Port:     port,
		BaudRate: baudrate,
	}

	// Open bus
	bus := can.NewBus(tr)

	if err := bus.Open(); err != nil {
		log.Fatal(err)
	}

	return CanBus{
		Bus: bus,
	}
}
