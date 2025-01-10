package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca IPs e Nomes de Servidores na Web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "walkernegrao.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços na Web",
			Flags:  flags,
			Action: buscaIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nome de servidores na Web",
			Flags:  flags,
			Action: buscaServidores,
		},
	}

	return app
}

func buscaIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

func buscaServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}

}
