package main

import (
	"fmt"
	"log"
	"mod-command-line/app"
	"os"
)

func main() {
	fmt.Println("Ponto de partida da aplicação")
	application := app.Gerar()
	erro := application.Run(os.Args)

	if erro != nil {
		log.Fatal(erro)
	}
}
