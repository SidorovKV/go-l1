package main

func main() {
	socket := &USBtypeASocket{}
	connector := &USBtypeCConnector{"somedata"}
	adapter := &USBAdaperTypeCToTypeA{&USBtypeCSocket{}, &USBtypeAConnector{}}
	adapter.Combine(connector)
	adapter.Connect(socket)
	println(socket.ReadData())

}

type USBtypeCConnector struct {
	data string
}

func (c *USBtypeCConnector) Connect(socket *USBtypeCSocket) {
	socket.connector = c
	println("Connected to", socket)
}

type USBtypeAConnector struct {
	data string
}

func (c *USBtypeAConnector) Connect(socket *USBtypeASocket) {
	socket.connector = c
	println("Connected to", socket)
}

type USBtypeASocket struct {
	connector *USBtypeAConnector
}

func (s *USBtypeASocket) ReadData() string {
	return s.connector.data
}

type USBtypeCSocket struct {
	connector *USBtypeCConnector
}

func (s *USBtypeCSocket) ReadData() string {
	return s.connector.data
}

type USBAdaperTypeCToTypeA struct {
	socketC *USBtypeCSocket
	typeA   *USBtypeAConnector
}

func (a *USBAdaperTypeCToTypeA) Combine(connector *USBtypeCConnector) {
	a.socketC.connector = connector
	a.typeA.data = a.socketC.ReadData()
}

func (a *USBAdaperTypeCToTypeA) Connect(socket *USBtypeASocket) {
	a.typeA.Connect(socket)
	println("Connected to", a.typeA)
}
