package main

import (
	"debug/macho"
	"log"
)

func mainPlatform(pathExe string) (gopclntab []byte, gosymtab []byte, text uint64) {
	file, err := macho.Open(pathExe)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if gopclntab, err = file.Section("__gopclntab").Data(); err != nil {
		log.Fatalln(err)
	}
	if gosymtab, err = file.Section("__gosymtab").Data(); err != nil {
		log.Fatalln(err)
	}
	text = file.Section("__text").Addr

	return
}
