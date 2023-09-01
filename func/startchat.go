package functions

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
)

func StartChat() {
	localhost := os.Args[1:]
	listenarr := []string{}
	listenErr := redColor + "[USAGE]: ./TCPChat $port" + resetColor

	listenstr := ""
	if len(localhost) < 1 {
		listenarr = append(listenarr, "8989")
	} else if len(localhost) > 1 {
		fmt.Println(listenErr)
		return
	} else {
		for _, check := range localhost {
			if check >= "0" && check <= "9" {
				num, err := strconv.Atoi(check)
				if err != nil {
					log.Fatal(err)
				}
				if num >= 1024 && num <= 65535 {
					listenarr = append(listenarr, strconv.Itoa(num))
				} else if num > 65535 || num < 1024 {
					fmt.Println(redColor + "ERROR: not correct port, input from port 1024 to 65535" + resetColor)
					return
				}
			} else {
				for i := 0; i <= len(localhost); i++ {
					for i := 0; i <= len(localhost); i++ {
						if localhost[0][i] < '0' || localhost[0][i] > '9' {
							listenstr = listenstr + listenErr
							fmt.Println(redColor + listenstr + " ERROR: please use only numbers!" + resetColor)
							return
						}
					}
				}
			}
		}
	}

	for _, localstr := range listenarr {
		listenstr = listenstr + "localhost:" + localstr
	}
	listener, err := net.Listen("tcp", listenstr)

	log.Println("Listening on the port:", listenstr)

	if err != nil {
		log.Fatal(err)
	}

	go Broadcaster()
	for {
		conn, err := listener.Accept() // консоль открывает
		if err != nil {
			log.Print(err)
			continue
		}

		mutex.Lock()
		go HandleConn(conn)
		if len(clients) == 10 {
			conn.Write([]byte(redColor + "The chat is not available, please try again later!" + resetColor))
			conn.Close()
		}
		mutex.Unlock()
	}
}
