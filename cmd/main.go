package main

import (
	"fmt"
	"log"
	"time"

	"github.com/lwhile/tomato/store"

	"github.com/lwhile/tomato"
	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	defaultMinutes = 25
)

var (
	name  = kingpin.Flag("name", "").Short('n').String()
	dur   = kingpin.Flag("duration", "").Short('d').Int()
	actor = kingpin.Arg("actor", "").Required().String()
)

// support `new` argument now
func actorCtrl() error {
	switch *actor {
	case "new":
		return new()
	}
	return fmt.Errorf("Actore `%s` not supported", *actor)
}

func new() error {
	t := tomato.New(*name, *dur)
	t.Start()

	// tomato finish
	<-t.Done
	return store.DefaultStore.Save(t)
}

func main() {
	kingpin.Parse()

	if *name == "" {
		*name = fmt.Sprintf("tomato@%d", time.Now().Unix())
	}
	if *dur == 0 {
		*dur = defaultMinutes
	}

	err := actorCtrl()
	if err != nil {
		log.Fatal(err)
	}
}
