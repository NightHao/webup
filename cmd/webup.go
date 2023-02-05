package main

import (
	"context"
	"golang.org/x/oauth2/google"
	"log"
	"webup/internal/gdoc"

	"io"
	"net/http"
	"os"
)

func clientMustFromFile(fn string) *http.Client {
	cred, err := os.ReadFile(fn)
	if err != nil {
		log.Fatalln(err)
	}
	conf, err := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/documents.readonly")
	if err != nil {
		log.Fatalln(err)
	}
	return conf.Client(context.Background())
}

func request(id string) []byte {
	client := clientMustFromFile("cred.json")
	req, _ := http.NewRequest("GET", "https://docs.googleapis.com/v1/documents/"+id, nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body
}

func main() {
	raw, _ := os.ReadFile("dump2.json")
	//https://elitesports.tcus.edu.tw/project/0813AIEvent.html
	//id := "1zDMdNztd20tlLttEWEBuYOTH6-qUSmgMiZmykmMfjhQ"
	//raw := request(id)
	//os.WriteFile("dump2.json", raw, 0666)
	if err := os.WriteFile("out.html", []byte(gdoc.Parse(raw)), 0666); err != nil {
		log.Fatalln(err)
	}
}
