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
	Id          int32
	OutId       int32
	InId        int32
	FreqA       float32
	Freq        float32
	Amp         float32
	Phase       float32
	osc         wave
	plot        [SAMPLE_COUNT]float32
	refreshTime float64
	last        float32
}

func (s *Oscillator) Init(data any) {
	if d, ok := data.(Oscillator); ok {
		s.Id = d.Id
		s.InId = d.InId
		s.OutId = d.OutId
		s.Amp = d.Amp
		s.Phase = d.Phase
		s.Freq = d.Freq
		s.FreqA = d.FreqA
	} else {
		s.Id = IdGen()
		s.InId = IdGen()
		s.OutId = IdGen()
		s.Amp = 1
		s.Phase = 0
		s.Freq = 4
		s.FreqA = 0
	}
	s.last = 0
	s.osc = newOscillator(s.Freq, s.Amp, s.Phase, 100)
	for i := range s.plot {
		s.plot[i] = 0
	}
	s.refreshTime = 0
}

func (s *Oscillator) Show() {
	imnodes.BeginNode(s.Id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Oscillator")
	imnodes.EndNodeTitleBar()

	imnodes.BeginInputAttribute(s.InId)
	imgui.Text("Frequency Modulation")
	imgui.DragFloatV("Ammount", &s.FreqA, 1, -100, 100, "%.1f", 0)
	imnodes.EndInputAttribute()
	if imgui.DragFloatV("Frequency", &s.Freq, 0.25, 0, 100, "%.1f", 0) {
		s.osc.Freq = s.Freq
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

	imnodes.BeginOutputAttribute(s.OutId)
	imgui.Text("Out")
	imgui.PlotLinesFloatPtrV("Data", &s.plot[0], SAMPLE_COUNT, 0, "", -1, 1, imgui.NewVec2(150, 25), 4)
	imnodes.EndOutputAttribute()

	imnodes.EndNode()
}

func (s *Oscillator) GetOutput(id int32) (bool, Value) {
	if id == s.OutId {
		return true, Value{t_float32, s.last}
	}
	return false, Value{t_null, nil}
}

func (s *Oscillator) SetInput(id int32, input Value) bool {
	if s.InId == id && input.Typ == t_float32 {
		s.osc.Freq = s.Freq + s.FreqA*input.Data.(float32)
		return true
	}
	return false
}

func (s *Oscillator) OutputList() []int32 {
	return []int32{s.OutId}
}

func (s *Oscillator) InputList() []int32 {
	return []int32{s.InId}
}

func (s *Oscillator) Type() NodeKind {
	return NodeOscillator
}
