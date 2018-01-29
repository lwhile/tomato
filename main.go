package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/lwhile/ttomato/ttomato"
)

const (
	defaultMinutes = 25
)

type arg struct {
	ctrl    string
	name    string
	minutes int
}

func parseArgs(args []string) (*arg, error) {
	argRet := &arg{}

	if len(args) > 0 {
		argRet.ctrl = args[0]
	}

	args = args[1:]
	if len(args) > 0 {
		argRet.name = args[0]
	} else {
		argRet.name = fmt.Sprintf("tomato@%d", time.Now().Unix())
	}

	args = args[1:]
	if len(args) > 0 {
		minutes, err := strconv.Atoi(args[0])
		if err != nil {
			return nil, err
		}
		argRet.minutes = minutes
	} else {
		argRet.minutes = defaultMinutes
	}

	return argRet, nil
}

func behaveCtrl(arg *arg) error {
	switch arg.ctrl {
	case "new":
		tomato := ttomato.New(arg.name, arg.minutes)
		return tomato.Start()
	}
	return nil
}

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("Use arguments `[new(n)]` to controll ttomato")
		return
	}
	arg, err := parseArgs(args[1:])
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := behaveCtrl(arg); err != nil {
		fmt.Println(err)
		return
	}

	select {}
}
