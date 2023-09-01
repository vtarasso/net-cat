package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func HandleConn(conn net.Conn) {
	Greetings(conn)

	eyn := "[ENTER YOUR NAME]: "
	fmt.Fprint(conn, eyn)

	name := ""
	nickName := bufio.NewScanner(conn)
	for nickName.Scan() {
		name = nickName.Text()
		name = strings.TrimSpace(name)

		if !CheckValid(name) {
			conn.Write([]byte(redColor + "Incorrect input.\nPlease use latin letters or numbers or special characters\n" + resetColor))
			conn.Write([]byte(eyn))
			continue
		} else if len(name) <= 2 || len(name) >= 16 {
			conn.Write([]byte(redColor + "Incorrect input.\nPlease enter a nickname between 3 and 15 characters long\n" + resetColor))
			conn.Write([]byte(eyn))

		} else if !IsUniqueName(name) {
			conn.Write([]byte(redColor + "Incorrect input\nNickname is used. Please create new nickname\n" + resetColor))
			conn.Write([]byte(eyn))

		} else {
			break
		}

	}

	mutex.Lock()
	clients[conn] = name
	mutex.Unlock()

	allstory := []Message{}

	for _, allhistory := range history.arrhistory {
		allstory = append(allstory, allhistory)
	}

	for _, w := range allstory {
		if w.Text != "{" || w.Text != "}" {
			fmt.Fprintln(conn, w.Text)
		} else {
			break
		}
	}

	if len(clients) <= 10 {
		mutex.Lock()
		messages <- Message{
			Name: "",
			Text: cyanColor + name + resetColor + greenColor + " has joined the chat!" + resetColor,
		}
		history.arrhistory = append(history.arrhistory, Message{Text: cyanColor + name + resetColor + greenColor + " has joined the chat!" + resetColor})
		mutex.Unlock()
	}

	input := bufio.NewScanner(conn)
	var newMess Message

	for input.Scan() {
		text := input.Text()

		if !CheckValid(text) {
			conn.Write([]byte(redColor + "Incorrect input.\nPlease use latin letters or numbers or special characters\n" + resetColor))
			conn.Write([]byte((fmt.Sprintf(MakeFormat(name, "")))))
			continue
		} else if strings.TrimSpace(text) == "" || len(text) >= 301 {
			conn.Write([]byte(redColor + "Incorrect input.\nPlease enter a message with content from 1(min) to 300(max) characters\n" + resetColor))
			conn.Write([]byte((fmt.Sprintf(MakeFormat(name, "")))))
			continue
		} else {
			for {
				newMess = Message{
					Name: name,
					Text: MakeFormat(name, text),
				}
				messages <- newMess
				history.arrhistory = append(history.arrhistory, Message{Text: MakeFormat(name, text)})
				break
			}
		}

	}
	mutex.Lock()
	if len(clients) <= 10 {
		newMess = Message{
			Name: name,
			Text: cyanColor + name + resetColor + greenColor + " has left the chat!" + resetColor,
		}

		messages <- newMess
	}
	delete(clients, conn)
	mutex.Unlock()
	conn.Close()
}
