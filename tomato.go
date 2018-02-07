package tomato

import (
	"fmt"
	"os"
	"time"

	"log"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

const (
	soundFile = "sound/sound0.wav"
)

// Tomato model definition
type Tomato struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Minutes   int       `json:"minutes"`
	StartTime time.Time `json:"start_time"`

	seconds int
	currLoc int
	Done    chan struct{} `json:"-"`
}

// New return a *Tomato
func New(name string, minutes int) *Tomato {
	return &Tomato{
		Name:    name,
		Minutes: minutes,

		seconds: minutes * 60,
		Done:    make(chan struct{}),
	}
}

// Start a tomato
func (t *Tomato) Start() {
	t.StartTime = time.Now()
	fmt.Printf("Start tomato(%d minutes) %s at %v\n", t.Minutes, t.Name, t.StartTime)
	t.running()
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
		case <-t.Done:
			ticker.Stop()
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
		close(t.Done)
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

	s, format, err := wav.Decode(fp)
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
