package songs

import (
	"fmt"
)

var (
	dayNames = []string{
		"zeroeth", "first", "second", "third", "fourth", "fifth",
		"sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth",
	}
	gifts = []string{
		"nothing",
		"an incident raised by JB",
		"two nasty bugs",
		"three pen tests",
		"four APIs",
		"five Go things",
		"six Pete's a-playing",
		"seven Sri's a-singing",
		"eight Kim's a-clapping",
		"nine Hannah's humming",
		"ten Steve's a-stomping",
		"eleven pipelines piping",
		"twelve app releases",
	}
)

type SongLine struct {
	Notes  string
	Lyrics string
}

type Song struct {
	Title string
	Lines []SongLine
}

func TwelveDays() Song {
	song := Song{
		Title: "Twelve Days of Think Money",
	}
	for i := 1; i <= 12; i++ {
		addTwelveDaysVerse(&song, i)
	}
	return song
}

func addTwelveDaysVerse(s *Song, v int) {
	if v > 1 {
		s.Lines = append(s.Lines, SongLine{})
	}

	addTwelveDaysFirstLine(s, v)
	if v >= 5 {
		for l := v; l > 5; l-- {
			addVerseLine(s, l)
		}
		addRingsToTurtleDoves(s)
	} else {
		for l := v; l > 1; l-- {
			addVerseLine(s, l)
		}
	}
	addPartridge(s, v > 1)
}

func addTwelveDaysFirstLine(s *Song, v int) {
	lastNoteLength := 4
	if v == 1 {
		lastNoteLength = 3
	}
	s.Lines = append(s.Lines, SongLine{
		Lyrics: fmt.Sprintf("On the %s day of Christmas, thinkmoney gave to me", dayNames[v]),
		Notes:  fmt.Sprintf("ggg2CCC2bCDEFDE%d", lastNoteLength),
	})
}

func addPartridge(s *Song, includeAnd bool) {
	line := SongLine{
		Lyrics: gifts[1] + ".",
		Notes:  "FG2AFECD2C6",
	}
	if includeAnd {
		line.Lyrics = "and " + line.Lyrics
		line.Notes = "E" + line.Notes
	}
	s.Lines = append(s.Lines, line)
}

func addRingsToTurtleDoves(s *Song) {
	s.Lines = append(s.Lines,
		SongLine{Lyrics: gifts[5], Notes: "G4A2F#2G6"},
		SongLine{Lyrics: gifts[4], Notes: "GFEDC2"},
		SongLine{Lyrics: gifts[3], Notes: "F2a2C2"},
		SongLine{Lyrics: gifts[2], Notes: "DCbag2"},
	)
}

func addVerseLine(s *Song, v int) {
	line := SongLine{
		Lyrics: gifts[v],
		Notes:  "G2DEF",
	}
	if v > 5 {
		line.Notes += "D" // extra note for "ing"
	} else {
		line.Notes += "2" // double length of last note
	}
	s.Lines = append(s.Lines, line)
}
