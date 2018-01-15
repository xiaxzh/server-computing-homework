package main

import (
	"os"

	"github.com/freakkid/service-agenda/service/server"
	flag "github.com/spf13/pflag"
)

// PORT is defalut to set 8080
const PORT string = "8080"

func main() {
	pPort := flag.StringP("port", "p", PORT, "PORT for listening")
	flag.Parse()
	var port string
	if len(*pPort) != 0 {
		port = *pPort
	} else {
		tPort := os.Getenv("PORT")
		if len(tPort) != 0 {
			port = tPort
		}
	}

	serverInstance := server.NewServer()
	serverInstance.Run(":" + port)
}
