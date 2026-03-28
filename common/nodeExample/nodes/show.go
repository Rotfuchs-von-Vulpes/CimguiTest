package nodes

import (
	"fmt"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"
)

type Show struct {
	id     int32
	inId   int32
	amount float32
}

func (s *Show) Init() {
	s.id = IdGen()
	s.inId = IdGen()
}

func (s *Show) Show() {
	imnodes.BeginNode(s.id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Show")
	imnodes.EndNodeTitleBar()

	imnodes.BeginInputAttribute(s.inId)
	imgui.Text(fmt.Sprintf("%.1f", s.amount))
	imnodes.EndInputAttribute()

	imnodes.EndNode()
}

func (s *Show) GetOutput(id int32) (bool, Value) {
	return false, Value{t_null, nil}
}

func (s *Show) SetInput(id int32, input Value) bool {
	if id == s.inId && input.typ == t_float32 {
		s.amount = input.data.(float32)
		return true
	}
	return false
}

func (s *Show) OutputList() []int32 {
	return []int32{}
}

func (s *Show) InputList() []int32 {
	return []int32{s.inId}
}

type ShowColor struct {
	id    int32
	inId  int32
	color [3]float32
}

func (s *ShowColor) Init() {
	s.id = IdGen()
	s.inId = IdGen()
}

func (s *ShowColor) Show() {
	imnodes.BeginNode(s.id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Color")
	imnodes.EndNodeTitleBar()

	imnodes.BeginInputAttribute(s.inId)
	imgui.Text("Color")
	imgui.ColorButton("Color", imgui.NewVec4(s.color[0], s.color[1], s.color[2], 1))
	imnodes.EndInputAttribute()

	imnodes.EndNode()
}

func (s *ShowColor) GetOutput(id int32) (bool, Value) {
	return false, Value{t_null, nil}
}

func (s *ShowColor) SetInput(id int32, input Value) bool {
	if id == s.inId && input.typ == t_3float32 {
		s.color = input.data.([3]float32)
		return true
	}
	return false
}

func (s *ShowColor) OutputList() []int32 {
	return []int32{}
}

func (s *ShowColor) InputList() []int32 {
	return []int32{s.inId}
}
