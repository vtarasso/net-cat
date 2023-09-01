package functions

import (
	"net"
	"sync"
)

type Message struct {
	Name string
	Text string
}

type History struct {
	arrhistory []Message
}

var (
	clients  = make(map[net.Conn]string) // Все подключенные клиенты
	messages = make(chan Message)        // Все входящие сообщения клиента
	mutex    sync.Mutex
	history  = History{}
)

var (
	redColor   = "\033[91m" // error
	greenColor = "\033[92m" // connected and disconnected user
	cyanColor  = "\033[96m" // nickname user connecting and disconnected
	resetColor = "\033[0m"  // reset
)
