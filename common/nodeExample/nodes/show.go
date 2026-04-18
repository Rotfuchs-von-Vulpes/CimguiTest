package nodes

import (
	"fmt"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"
)

type Show struct {
	Id     int32
	InId   int32
	amount float32
}

func (s *Show) Init(data any) {
	if d, ok := data.(Show); ok {
		s.Id = d.Id
		s.InId = d.InId
	} else {
		s.Id = IdGen()
		s.InId = IdGen()
	}
}

func (s *Show) Show() {
	imnodes.BeginNode(s.Id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Show")
	imnodes.EndNodeTitleBar()

	imnodes.BeginInputAttribute(s.InId)
	imgui.Text(fmt.Sprintf("%.1f", s.amount))
	imnodes.EndInputAttribute()

	imnodes.EndNode()
}

func (s *Show) GetOutput(id int32) (bool, Value) {
	return false, Value{t_null, nil}
}

func (s *Show) SetInput(id int32, input Value) bool {
	if id == s.InId && input.Typ == t_float32 {
		s.amount = input.Data.(float32)
		return true
	}
	return false
}

func (s *Show) OutputList() []int32 {
	return []int32{}
}

func (s *Show) InputList() []int32 {
	return []int32{s.InId}
}

func (s *Show) Type() NodeKind {
	return NodeShow
}

type ShowColor struct {
	Id    int32
	InId  int32
	color [3]float32
}

func (s *ShowColor) Init(data any) {
	if d, ok := data.(ShowColor); ok {
		s.Id = d.Id
		s.InId = d.InId
	} else {
		s.Id = IdGen()
		s.InId = IdGen()
	}
}

func (s *ShowColor) Show() {
	imnodes.BeginNode(s.Id)

	imnodes.BeginNodeTitleBar()
	imgui.Text("Color")
	imnodes.EndNodeTitleBar()

	imnodes.BeginInputAttribute(s.InId)
	imgui.Text("Color")
	imgui.ColorButton("Color", imgui.NewVec4(s.color[0], s.color[1], s.color[2], 1))
	imnodes.EndInputAttribute()

	imnodes.EndNode()
}

func (s *ShowColor) GetOutput(id int32) (bool, Value) {
	return true, Value{t_3float32, s.color}
}

func (s *ShowColor) SetInput(id int32, input Value) bool {
	if id == s.InId && input.Typ == t_3float32 {
		s.color = input.Data.([3]float32)
		return true
	}
	return false
}

func (s *ShowColor) OutputList() []int32 {
	return []int32{}
}

func (s *ShowColor) InputList() []int32 {
	return []int32{s.InId}
}

func (s *ShowColor) Type() NodeKind {
	return NodeShowColor
}
