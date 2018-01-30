package ttomato

import (
	"fmt"
	"os"
	"time"

	"log"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const (
	soundFile = "sound/sound1.mp3"
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
	t.running()
	return nil
}

func (t *Tomato) running() {
	if err := setBoundary(); err != nil {
		log.Fatal(err)
	}

	t.currLoc = 0
	durSec := t.calTickerDur()
	ticker := time.NewTicker(durSec)
	for {
		select {
		case <-ticker.C:
			t.currLoc++
			t.triggerPrint()
		case <-t.done:
			t.finish()
			return
		}
	}
}

// calculate the ticker duration
func (t *Tomato) calTickerDur() time.Duration {
	return time.Duration(float32(t.seconds)/float32(printWidth)*1000) * time.Millisecond
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

func (t *Tomato) finish() {
	fmt.Printf("\nFinish the tomato %s\n", t.Name)
	t.playSound()
}

func (t *Tomato) playSound() {
	// play the sound
	fp, err := os.Open(soundFile)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	s, format, err := mp3.Decode(fp)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	go speaker.Play(beep.Seq(s, beep.Callback(func() { close(done) })))

	<-done
}

// Stop a tomato
func (t *Tomato) Stop() error {
	return nil
}
