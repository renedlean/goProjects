package app

import (
	
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerar vai retornar a aplicação de linha de comando

func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "aplicação de linha de comando"
	app.Usage = "Busca IPS e nome de servidor"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name: "IP",
			Usage: "busca IPS",
			Flags: flags,
			Action: buscarIps,
				
		},
		{
			Name: "servidores",
			Usage: "Buscar nome dos servidores",
			Flags: flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps (c *cli.Context){
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _ , ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context){
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores{
		fmt.Println(servidor)
	}
}