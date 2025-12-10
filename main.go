package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"strings"

	"github.com/hajimehoshi/oto"
	"thinkmoney.co.uk/music/songs"
)

const notes = "c,c#,d,d#,e,f,f#,g,g#,a,a#,b,C,C#,D,D#,E,F,F#,G,G#,A,A#,B"

func main() {
	fmt.Println("initialising player")
	p := newPlayer(meantToBeAPianoButSoundsMoreLikeSomeSortOfGuitar)
	defer p.Close()

	fmt.Println("playing. press Ctrl-C to shut me up")
	p.playSong(songs.TwelveDays())
}

func newPlayer(instrument instrument) *player {
	noteMap := map[string]float64{}
	notes := strings.Split(notes, ",")
	for i,n := range notes {
		noteMap[n] = 440 * math.Pow(2, float64(i)/12.0)
	}

	c, err := oto.NewContext(48000, 2, 2, 8192)
	if err != nil {
		panic(err)
	}
	p := c.NewPlayer()

	return &player{
		playerStream: p,
		noteMap: noteMap,
		instrument: instrument,
	}
}

type player struct {
	playerStream io.WriteCloser
	noteMap map[string]float64
	instrument instrument
}

func (p *player) Close() {
	p.playerStream.Close()
}

func (p *player) playSong(song string) {
	song += ".."
	i := 0
	for i < len(song) - 3 {
		note := song[i:i+1]
		next1 := song[i+1]
		lenChar := song[i+1]
		len := 1
		if next1 == '#' {
			note = song[i:i+2]
			lenChar = song[i+2]
			i++
		}
		if lenChar >= '1' && lenChar <= '9' {
			len = int(lenChar - '0')
			i++
		}
		i++
		p.playNote(note, len)
	}
}

func (p *player) playNote(note string, semiquavers int) {
	freq := p.getFreq(note)
	p.playSound(freq, float64(semiquavers * 240.0)-100)
	p.playSilence(100)
}

func (p *player) getFreq(note string) float64{
	freq, ok := p.noteMap[note]
	if !ok {
		panic("Invalid note")
	}
	return freq
}

func (p *player) playSound(freq float64, durationMS float64) {
	adjustedFreq := float64(freq) / 48000.0
	for i:=float64(0);i< durationMS * 48;i++ {
		binary.Write(p.playerStream, binary.LittleEndian, int16(p.instrument(adjustedFreq, 32767, i)))
	}
}

func (p *player) playSilence(durationMS float64) {
	for i:=float64(0);i< durationMS * 48;i++ {
		binary.Write(p.playerStream, binary.LittleEndian, int16(0))
	}
}


