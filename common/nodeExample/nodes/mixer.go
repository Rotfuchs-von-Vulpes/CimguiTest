package nodes

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"
)

type MixerColor struct {
	id         int32
	inId1      int32
	inId2      int32
	inAmmoutId int32
	outId      int32
	ammount    float32
	inColor1   [3]float32
	inColor2   [3]float32
	outColor   [3]float32
}

func (s *MixerColor) Init() {
	s.id = IdGen()
	s.inId1 = IdGen()
	s.inId2 = IdGen()
	s.inAmmoutId = IdGen()
	s.outId = IdGen()
}

func (s *MixerColor) Show() {

	for i := range s.outColor {
		s.outColor[i] = s.ammount*s.inColor1[i] + (1-s.ammount)*s.inColor2[i]
	}

	imnodes.BeginNode(s.id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Color Mixer")
	imnodes.EndNodeTitleBar()

	imnodes.BeginInputAttribute(s.inId1)
	imgui.Text("Color input 1")
	imgui.ColorButton("Color", imgui.NewVec4(s.inColor1[0], s.inColor1[1], s.inColor1[2], 1))
	imnodes.EndInputAttribute()

	imnodes.BeginInputAttribute(s.inId2)
	imgui.Text("Color input 2")
	imgui.ColorButton("Color", imgui.NewVec4(s.inColor2[0], s.inColor2[1], s.inColor2[2], 1))
	imnodes.EndInputAttribute()

	imnodes.BeginInputAttribute(s.inAmmoutId)
	imgui.Text("Mixer coefficient")
	imnodes.EndInputAttribute()

	imnodes.BeginOutputAttribute(s.outId)
	imgui.Text("Color output")
	imgui.ColorButton("Color", imgui.NewVec4(s.outColor[0], s.outColor[1], s.outColor[2], 1))
	imnodes.EndInputAttribute()

	imnodes.EndNode()
}

func (s *MixerColor) GetOutput(id int32) (bool, Value) {
	return true, Value{t_3float32, s.outColor}
}

func (s *MixerColor) SetInput(id int32, input Value) bool {
	if id == s.inId1 && input.Typ == t_3float32 {
		s.inColor1 = input.Data.([3]float32)
		return true
	} else if id == s.inId2 && input.Typ == t_3float32 {
		s.inColor2 = input.Data.([3]float32)
		return true
	} else if id == s.inAmmoutId && input.Typ == t_float32 {
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
	return []int32{s.outId}
}

func (s *MixerColor) InputList() []int32 {
	return []int32{s.inId1, s.inId2, s.inAmmoutId}
}

func (s *MixerColor) Type() NodeKind {
	return NodeColorMixer
}
