package nodes

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"
)

type Constant struct {
	Id     int32
	Amount int32
	OutId  int32
}

func (s *Constant) Init(data any) {
	if d, ok := data.(Constant); ok {
		s.Id = d.Id
		s.OutId = d.OutId
		s.Amount = d.Amount
	} else {
		s.Id = IdGen()
		s.OutId = IdGen()
		s.Amount = 10
	}
}

func (s *Constant) Show() {
	imnodes.BeginNode(s.Id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Constant")
	imnodes.EndNodeTitleBar()

	imnodes.BeginOutputAttribute(s.OutId)
	imgui.InputInt("Output", &s.Amount)
	imnodes.EndOutputAttribute()

	imnodes.EndNode()
}

func (s *Constant) GetOutput(id int32) (bool, Value) {
	if id == s.OutId {
		return true, Value{t_float32, float32(s.Amount)}
	}
	return false, Value{t_null, nil}
}

func (s *Constant) SetInput(id int32, data Value) bool {
	return false
}

func (s *Constant) OutputList() []int32 {
	return []int32{s.OutId}
}

func (s *Constant) InputList() []int32 {
	return []int32{}
}

func (s *Constant) Type() NodeKind {
	return NodeConstant
}

type ColorConstant struct {
	Id    int32
	OutId int32
	Color [3]float32
}

func (s *ColorConstant) Init(data any) {
	if d, ok := data.(ColorConstant); ok {
		s.Id = d.Id
		s.OutId = d.OutId
		s.Color = d.Color
	} else {
		s.Id = IdGen()
		s.OutId = IdGen()
		s.Color = [3]float32{1, 1, 1}
	}
}

func (s *ColorConstant) Show() {
	imnodes.BeginNode(s.Id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Color Contant")
	imnodes.EndNodeTitleBar()

	imnodes.BeginOutputAttribute(s.OutId)
	imgui.ColorEdit3("Output color", &s.Color)
	imnodes.EndInputAttribute()

	imnodes.EndNode()
}

func (s *ColorConstant) GetOutput(id int32) (bool, Value) {
	return true, Value{t_3float32, s.Color}
}

func (s *ColorConstant) SetInput(id int32, data Value) bool {
	return true
}

func (s *ColorConstant) OutputList() []int32 {
	return []int32{s.OutId}
}

func (s *ColorConstant) InputList() []int32 {
	return []int32{}
}

func (s *ColorConstant) Type() NodeKind {
	return NodeColor
}
