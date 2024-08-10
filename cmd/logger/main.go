package main

const (
	PORT     = "com5"
	BAUDRATE = 2000000
)

func main() {

	bus := can.Connect(PORT, BAUDRATE)
}
