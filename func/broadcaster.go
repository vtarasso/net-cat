package functions

import (
	"fmt"
	"log"
)

func Broadcaster() {
	for {
		select {
		case msg := <-messages:
			mutex.Lock()
			for conn, user := range clients {
				if user != msg.Name {
					_, err := fmt.Fprint(conn, "\033[2K\r"+msg.Text+"\n")
					if err != nil {
						log.Println(err)
						return
					}
				}
				_, err := fmt.Fprint(conn, "\033[2K\r"+MakeFormat(user, ""))
				if err != nil {
					log.Println(err)
					return
				}
			}
			mutex.Unlock()
		}
	}
}
