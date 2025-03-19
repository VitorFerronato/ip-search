package main

import (
	"fmt"
	"ip-search/app"
	"log"
	"os"
)

func main() {
	fmt.Println("========")
	aplication := app.Generate()

	if erro := aplication.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
