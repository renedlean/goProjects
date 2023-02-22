package main

import (
	"appComando/app"
	"fmt"
	"log"
	"os"
)

func main()  {
	fmt.Println("PONTO DE INICIO")

	aplicacao := app.Gerar()
	erro :=	aplicacao.Run(os.Args)

		if erro != nil {
			log.Fatal(erro)
		}

}