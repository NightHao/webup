package main

import (
	"fmt"
	"log"
	"webup/internal/cms"
)

func main() {
	id := "12CkaxfCn4RMs1gmt3tJxtYq-As0F22ssP6XndGhWDsY"

	items, err := cms.GetMenu(id)
	if err != nil {
		log.Fatalln(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
