package nodes

import (
	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"
)

type Constant struct {
	id     int32
	amount int32
	outId  int32
}

func (s *Constant) Init() {
	s.id = IdGen()
	s.outId = IdGen()
	s.amount = 10
}

func (s *Constant) Show() {
	imnodes.BeginNode(s.id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Constant")
	imnodes.EndNodeTitleBar()

	imnodes.BeginOutputAttribute(s.outId)
	imgui.InputInt("Output", &s.amount)
	imnodes.EndOutputAttribute()

	imnodes.EndNode()
}

func (s *Constant) GetOutput(id int32) (bool, Value) {
	if id == s.outId {
		return true, Value{t_float32, float32(s.amount)}
	}
	return false, Value{t_null, nil}
}

func (s *Constant) SetInput(id int32, data Value) bool {
	return false
}

func (s *Constant) OutputList() []int32 {
	return []int32{s.outId}
}

func (s *Constant) InputList() []int32 {
	return []int32{}
}

type ColorConstant struct {
	id    int32
	outId int32
	color [3]float32
}

func (s *ColorConstant) Init() {
	s.id = IdGen()
	s.outId = IdGen()
	s.color = [3]float32{1, 1, 1}
}

func (s *ColorConstant) Show() {
	imnodes.BeginNode(s.id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Color Contant")
	imnodes.EndNodeTitleBar()

	imnodes.BeginOutputAttribute(s.outId)
	imgui.ColorEdit3("Output color", &s.color)
	imnodes.EndInputAttribute()

	imnodes.EndNode()
}

func (s *ColorConstant) GetOutput(id int32) (bool, Value) {
	return true, Value{t_3float32, s.color}
}

func (s *ColorConstant) SetInput(id int32, data Value) bool {
	return true
}

func (s *ColorConstant) OutputList() []int32 {
	return []int32{s.outId}
}

func (s *ColorConstant) InputList() []int32 {
	return []int32{}
}
