package nodes

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"

	"math"
)

type wave struct {
	Freq       float32
	Amp        float32
	Phase      float32
	ammount    float32
	sampleRate float32
}

func newOscillator(freq, amp, phase, rate float32) wave {
	return wave{freq, amp, phase, 0, rate}
}

func (s *wave) sample() (sample float32) {
	sample = s.Amp * float32(math.Sin(float64(s.ammount+s.Phase)))

	s.ammount += 2 * math.Pi * s.Freq / s.sampleRate
	if s.ammount > 2*math.Pi {
		s.ammount -= 2 * math.Pi
	}

	return sample
}

const SAMPLE_COUNT int32 = 100

type Oscillator struct {
	id          int32
	outId       int32
	inId        int32
	freqA       float32
	freq        float32
	osc         wave
	plot        [SAMPLE_COUNT]float32
	refreshTime float64
	last        float32
}

func (s *Oscillator) Init() {
	s.id = IdGen()
	s.inId = IdGen()
	s.outId = IdGen()
	s.last = 0
	s.freq = 4
	s.osc = newOscillator(s.freq, 1, 0, 100)
	for i := range s.plot {
		s.plot[i] = 0
	}
	s.refreshTime = 0
}

func (s *Oscillator) Show() {
	imnodes.BeginNode(s.id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Oscillator")
	imnodes.EndNodeTitleBar()

	imnodes.BeginInputAttribute(s.inId)
	imgui.Text("Frequency Modulation")
	imgui.DragFloatV("Ammount", &s.freqA, 1, -100, 100, "%.1f", 0)
	imnodes.EndInputAttribute()
	if imgui.DragFloatV("Frequency", &s.freq, 0.25, 0, 100, "%.1f", 0) {
		s.osc.Freq = s.freq
	}
	imgui.DragFloatV("Phase", &s.osc.Phase, 0.25, 0, 100, "%.1f", 0)
	imgui.DragFloatV("Amp", &s.osc.Amp, 0.25, 0, 1, "%.1f", 0)

	{
		if s.refreshTime == 0 {
			s.refreshTime = imgui.Time()
		}
		newSamples := []float32{}
		for s.refreshTime < imgui.Time() {
			newSamples = append(newSamples, s.osc.sample())
			s.refreshTime += 1.0 / 60
		}
		if len(newSamples) > 0 {
			s.last = newSamples[len(newSamples)-1]
		}

		shift := len(newSamples)
		if shift > 0 {
			if shift >= int(SAMPLE_COUNT) {
				copy(s.plot[:], newSamples[shift-len(s.plot):])
			} else {
				copy(s.plot[:], s.plot[shift:])
				copy(s.plot[int(SAMPLE_COUNT)-shift:], newSamples)
			}
		}
	}

	imnodes.BeginOutputAttribute(s.outId)
	imgui.Text("Out")
	imgui.PlotLinesFloatPtrV("Data", &s.plot[0], SAMPLE_COUNT, 0, "", -1, 1, imgui.NewVec2(150, 25), 4)
	imnodes.EndOutputAttribute()

	imnodes.EndNode()
}

func (s *Oscillator) GetOutput(id int32) (bool, Value) {
	if id == s.outId {
		return true, Value{t_float32, s.last}
	}
	return false, Value{t_null, nil}
}

func (s *Oscillator) SetInput(id int32, input Value) bool {
	if s.inId == id && input.typ == t_float32 {
		s.osc.Freq = s.freq + s.freqA*input.data.(float32)
		return true
	}
	return false
}

func (s *Oscillator) OutputList() []int32 {
	return []int32{s.outId}
}

func (s *Oscillator) InputList() []int32 {
	return []int32{s.inId}
}
