package main

import (
	"fmt"
	"log"

	bc "github.com/arferreira/banco-central-brasil/bancocentral"
)

func main() {
	cambio := bc.NewCambio()
	dataDolarPtax, err := cambio.GetDolarCompraPtax()
	if err != nil {
		log.Println("Error getting Dólar PTAX date:", err)
		return
	}
	fmt.Printf("Dólar PTAX em %.2f\n", dataDolarPtax)

}
