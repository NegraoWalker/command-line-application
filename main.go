package main

import (
	"fmt"
	"log"
	"mod-command-line/app"
	"os"
)

func main() {
	fmt.Println("IPs ou Servidores:")
	application := app.Gerar()
	erro := application.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
