package functions

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func MakeFormat(name, txt string) string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "]" + "[" + name + "]" + ": " + strings.TrimSpace(txt)
}

func Greetings(conn net.Conn) {
	logo, err := os.ReadFile("pinguin.txt")
	if err != nil {
		fmt.Fprint(conn, redColor+"ERROR: Chat-logo is not loaded!\n"+resetColor)
		return
	}
	fmt.Fprintln(conn, string(logo))
}

func CheckValid(str string) bool {
	for _, c := range str {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}

func IsUniqueName(resname string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	for _, userName := range clients {
		if userName == resname {
			return false
		}
	}
	return true
}
