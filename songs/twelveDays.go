package songs

import "bytes"

func TwelveDays() string {
	b := bytes.Buffer{}
	for i:=1;i<=12;i++{
		addTwelveDaysVerse(&b, i)
	}
	return b.String()
}

func addTwelveDaysVerse(b *bytes.Buffer, v int) {
	addTwelveDaysFirstLine(b, v==1)
	if v >= 5 {
		for range v-5 {
			addVerseLine(b, true)
		}
		addRingsToTurtleDoves(b)
	} else {
		for range v-1 {
			addVerseLine(b, false)
		}
	}
	addPartridge(b, v==1)
}

func addTwelveDaysFirstLine(b *bytes.Buffer, isFirstVerse bool) {
	b.WriteString("ggg2CCC2bCDEFDE")
	if isFirstVerse{
		b.WriteRune('3')
	} else {
		b.WriteRune('4')
	}
}

func addPartridge(b *bytes.Buffer, isFirstVerse bool) {
	if !isFirstVerse {
		b.WriteRune('E')
	}
	b.WriteString("FG2AFECD2C6")
}

func addRingsToTurtleDoves(b *bytes.Buffer) {
	b.WriteString("G4A2F#2G6GFEDC2F2a2C2DCbag2")
}

func addVerseLine(b *bytes.Buffer, ing bool) {
	b.WriteString("G2DEF")
	if ing{
		b.WriteRune('D')
	} else {
		b.WriteRune('2')
	}

}