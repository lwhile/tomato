package ttomato

import (
	"fmt"
	"time"

	"log"
)

// Tomato model definition
type Tomato struct {
	Name      string
	Minutes   int
	StartTime time.Time

	seconds int
	currLoc int
	done    chan struct{}
}

// New return a *Tomato
func New(name string, minutes int) *Tomato {
	return &Tomato{
		Name:    name,
		Minutes: minutes,

		seconds: minutes * 60,
		done:    make(chan struct{}),
	}
}

// Start a tomato
func (t *Tomato) Start() error {
	t.StartTime = time.Now()
	fmt.Printf("Start tomato(%d minutes) %s at %v\n", t.Minutes, t.Name, t.StartTime)
	go t.running()
	return nil
}

func (t *Tomato) running() {
	if err := setBoundary(); err != nil {
		log.Fatal(err)
	}

	t.currLoc = 0
	durSec := time.Duration(float32(t.seconds)/float32(printWidth)*1000) * time.Millisecond
	ticker := time.NewTicker(durSec)
	for {
		select {
		case <-ticker.C:
			t.currLoc++
			t.triggerPrint()
		case <-t.done:
			fmt.Printf("\nFinish the tomato %s\n", t.Name)
			return
		}
	}
}

func (t *Tomato) triggerPrint() {
	lastOne := false
	if t.currLoc >= int(printWidth) {
		lastOne = true
	}
	if t.currLoc%2 != 0 {
		printOneTomato(lastOne)
	}
	if lastOne {
		close(t.done)
	}
}

// Stop a tomato
func (t *Tomato) Stop() error {
	return nil
}
