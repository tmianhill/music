package main

import "math"

type instrument func(freq float64, amp float64, i float64) float64

func pureTone(freq float64, amp float64, i float64) float64 {
	return amp*math.Sin(i*freq)
}

func asdrEnvelope(i, a, d, s, decayRate float64) float64 {
	if i < a {
		return i/a
	} else if i < a+d {
		return s + (1.0-s)*math.Exp(-(i-a)/d*decayRate)
	} else {
		return s * math.Exp(-(i-a-d)/24000)
	}
}

func meantToBeAPianoButSoundsMoreLikeSomeSortOfGuitar(freq float64, amp float64, i float64) float64 {
	// Frequency-dependent harmonic scaling
	harmonicScale := math.Max(0.3, 1.0-freq/1000.0)
	
	attackNoise := 0.0
	if i < 240 {
		attackNoise = math.Sin(i*0.1) * amp * 0.05 * (240-i)/240
	}
	
	// Inharmonic partials with individual decay rates
	fundamental := pureTone(freq/2, amp*0.5, i) * asdrEnvelope(i, 80, 3600, 0.35, 3.0)
	harmonic2 := pureTone(freq*1.005, amp*0.25*harmonicScale, i) * asdrEnvelope(i, 80, 2400, 0.2, 4.0)
	harmonic3 := pureTone(freq*1.51, amp*0.167*harmonicScale, i) * asdrEnvelope(i, 80, 1800, 0.15, 5.0)
	harmonic4 := pureTone(freq*2.025, amp*0.1*harmonicScale, i) * asdrEnvelope(i, 80, 1200, 0.1, 6.0)
	harmonic5 := pureTone(freq*2.54, amp*0.05*harmonicScale, i) * asdrEnvelope(i, 80, 800, 0.05, 7.0)
	
	// Sympathetic resonance
	sympathetic := pureTone(freq/4, amp*0.02, i) * asdrEnvelope(i, 80, 4800, 0.4, 2.0) +
				   pureTone(freq*0.75, amp*0.025, i) * asdrEnvelope(i, 80, 2000, 0.2, 4.5)
	
	return (fundamental + harmonic2 + harmonic3 + harmonic4 + harmonic5 + sympathetic + attackNoise) / 1.4
}
