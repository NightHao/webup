package main

import (
	"log"
	"webup/internal/gdrive"
)

func main() {
	folder := "1GXeYQOvNvDvqhtSZW4mOhJCBgcG-d_6r"
	log.Fatalln(gdrive.List(folder))
}
