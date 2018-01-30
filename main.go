package main

import (
	"fmt"
	"time"

	"github.com/lwhile/ttomato/ttomato"

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
		tomato := ttomato.New(*name, *dur)
		return tomato.Start()
	}
	return fmt.Errorf("Actore %s not supported", *actor)
}

func main() {
	kingpin.Parse()

	if *name == "" {
		*name = fmt.Sprintf("tomato@%d", time.Now().Unix())
	}
	if *dur == 0 {
		*dur = defaultMinutes
	}
	actorCtrl()
}
