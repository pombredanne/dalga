package main

import (
	"code.google.com/p/gcfg"
	"flag"
	"fmt"
	"github.com/cenkalti/dalga"
	"log"
)

var (
	configPath  = flag.String("c", "", "config file path")
	createTable = flag.Bool("t", false, "create table for storing jobs")
)

func main() {
	var err error
	flag.Parse()

	// Read config
	c := dalga.NewConfig()
	if *configPath != "" {
		err = gcfg.ReadFileInto(c, *configPath)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Read config: ", c)
	}

	// Initialize Dalga object
	d := dalga.NewDalga(c)

	// Create jobs table
	if *createTable {
		err = d.CreateTable()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Table created successfully")
		return
	}

	// Run Dalga
	err = d.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
