package main

import (
	"log"
	"os"

	"webup/internal/gdoc"
)

func main() {
	//raw, _ := os.ReadFile("dump2.json")
	//https://elitesports.tcus.edu.tw/project/0813AIEvent.html
	id := "1zDMdNztd20tlLttEWEBuYOTH6-qUSmgMiZmykmMfjhQ"
	raw, err := gdoc.Request(id)
	if err != nil {
		log.Fatalln(err)
	}
	os.WriteFile("dump2.json", raw, 0666)
	result, err := gdoc.Parse(raw)
	if err != nil {
		log.Fatalln(err)
	}
	if err := os.WriteFile("out.html", []byte(result), 0666); err != nil {
		log.Fatalln(err)
	}
}
