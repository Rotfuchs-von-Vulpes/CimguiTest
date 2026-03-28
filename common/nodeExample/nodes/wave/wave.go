package wave

import "math"

type Oscillator struct {
	Freq       float32
	Amp        float32
	Phase      float32
	ammount    float32
	sampleRate float32
}

func NewOscillator(freq, amp, phase, rate float32) Oscillator {
	return Oscillator{freq, amp, phase, 0, rate}
}

func (s *Oscillator) Sample() (sample float32) {
	sample = s.Amp * float32(math.Sin(float64(s.ammount+s.Phase)))

	s.ammount += 2 * math.Pi * s.Freq / s.sampleRate
	if s.ammount > 2*math.Pi {
		s.ammount -= 2 * math.Pi
	}

	return sample
}
