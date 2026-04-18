package nodes

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"
)

type MixerColor struct {
	Id         int32
	InId1      int32
	InId2      int32
	InAmmoutId int32
	OutId      int32
	ammount    float32
	inColor1   [3]float32
	inColor2   [3]float32
	outColor   [3]float32
}

func (s *MixerColor) Init(data any) {
	if d, ok := data.(MixerColor); ok {
		s.Id = d.Id
		s.InId1 = d.InId1
		s.InId2 = d.InId2
		s.InAmmoutId = d.InAmmoutId
		s.OutId = d.OutId
	} else {
		s.Id = IdGen()
		s.InId1 = IdGen()
		s.InId2 = IdGen()
		s.InAmmoutId = IdGen()
		s.OutId = IdGen()
	}
}

func (s *MixerColor) Show() {
	for i := range s.outColor {
		s.outColor[i] = s.ammount*s.inColor1[i] + (1-s.ammount)*s.inColor2[i]
	}

	imnodes.BeginNode(s.Id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Color Mixer")
	imnodes.EndNodeTitleBar()

	imnodes.BeginInputAttribute(s.InId1)
	imgui.Text("Color input 1")
	imgui.ColorButton("Color", imgui.NewVec4(s.inColor1[0], s.inColor1[1], s.inColor1[2], 1))
	imnodes.EndInputAttribute()

	imnodes.BeginInputAttribute(s.InId2)
	imgui.Text("Color input 2")
	imgui.ColorButton("Color", imgui.NewVec4(s.inColor2[0], s.inColor2[1], s.inColor2[2], 1))
	imnodes.EndInputAttribute()

	imnodes.BeginInputAttribute(s.InAmmoutId)
	imgui.Text("Mixer coefficient")
	imnodes.EndInputAttribute()

	imnodes.BeginOutputAttribute(s.OutId)
	imgui.Text("Color output")
	imgui.ColorButton("Color", imgui.NewVec4(s.outColor[0], s.outColor[1], s.outColor[2], 1))
	imnodes.EndInputAttribute()

	imnodes.EndNode()
}

func (s *MixerColor) GetOutput(id int32) (bool, Value) {
	return true, Value{t_3float32, s.outColor}
}

func (s *MixerColor) SetInput(id int32, input Value) bool {
	if id == s.InId1 && input.Typ == t_3float32 {
		s.inColor1 = input.Data.([3]float32)
		return true
	} else if id == s.InId2 && input.Typ == t_3float32 {
		s.inColor2 = input.Data.([3]float32)
		return true
	} else if id == s.InAmmoutId && input.Typ == t_float32 {
		s.ammount = input.Data.(float32)
		s.ammount = s.ammount/2 + 0.5
		if s.ammount > 1 {
			s.ammount = 1
		} else if s.ammount < 0 {
			s.ammount = 0
		}
		return true
	}
	return false
}

func (s *MixerColor) OutputList() []int32 {
	return []int32{s.OutId}
}

func (s *MixerColor) InputList() []int32 {
	return []int32{s.InId1, s.InId2, s.InAmmoutId}
}

func (s *MixerColor) Type() NodeKind {
	return NodeColorMixer
}
