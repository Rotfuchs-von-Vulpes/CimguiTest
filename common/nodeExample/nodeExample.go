package nodeExample

import (
	"slices"
	"unsafe"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"

	n "CimguiTest/common/nodeExample/nodes"
)

type link struct {
	id        int32
	start     int32
	end       int32
	startNode *NodeBlock
	endNode   *NodeBlock
}

type connection struct {
	id   int32
	data float32
}

func (s *connection) peek() float32 {
	return s.data
}

func (s *connection) poke(in float32) {
	s.data = in
}

type NodeBlock interface {
	Init()
	Show()
	InputList() []int32
	OutputList() []int32
	Type() n.NodeKind
	GetOutput(id int32) (bool, n.Value)
	SetInput(id int32, data n.Value) bool
}

var (
	nodes []NodeBlock
	links []link
)

func addNode(kind n.NodeKind) {
	var node NodeBlock
	switch kind {
	case n.NodeConstant:
		node = &n.Constant{}
	case n.NodeOscillator:
		node = &n.Oscillator{}
	case n.NodeShow:
		node = &n.Show{}
	case n.NodeColor:
		node = &n.ColorConstant{}
	case n.NodeShowColor:
		node = &n.ShowColor{}
	case n.NodeColorMixer:
		node = &n.MixerColor{}
	default:
		return
	}
	node.Init()
	nodes = append(nodes, node)
}

func removeLink(id int32) {
	if len(links) > 0 {
		for i, link := range links {
			if link.id == id {
				links[i] = links[len(links)-1]
				break
			}
		}
		links = links[:len(links)-1]
	}
}

func findLinkIdByStart(id int32) (bool, int32) {
	for _, link := range links {
		if link.start == id {
			return true, link.id
		}
	}
	return false, 0
}

func findLinkIdByEnd(id int32) (bool, int32) {
	for _, link := range links {
		if link.end == id {
			return true, link.id
		}
	}
	return false, 0
}

func findInputById(id int32) *NodeBlock {
	for _, node := range nodes {
		inputs := node.InputList()
		if slices.Contains(inputs, id) {
			return &node
		}
	}

	return nil
}

func findOutputById(id int32) *NodeBlock {
	for _, node := range nodes {
		outputs := node.OutputList()
		if slices.Contains(outputs, id) {
			return &node
		}
	}

	return nil
}

type nodeData struct {
	typ     n.NodeKind
	id      int32
	inputs  []int32
	outputs []int32
}

type linkData struct {
	id    int32
	start int32
	end   int32
}

type nodesData struct {
	nodes []nodeData
	links []linkData
}

func saveData() {
	var linksList []linkData

	for _, link := range links {
		linksList = append(linksList, linkData{link.id, link.start, link.end})
	}
}

var ctx *imnodes.EditorContext
var first = true

func Init() {
	addNode(n.NodeOscillator)
	addNode(n.NodeOscillator)
	addNode(n.NodeColor)
	addNode(n.NodeColor)
	addNode(n.NodeColorMixer)
	addNode(n.NodeShowColor)
	ctx = imnodes.EditorContextCreate()
}

func Show() {
	for _, link := range links {
		start := *link.startNode
		end := *link.endNode
		if ok, data := start.GetOutput(link.start); ok {
			if !end.SetInput(link.end, data) {
				removeLink(link.id)
			}
		} else {
			removeLink(link.id)
		}
	}

	basePos := imgui.MainViewport().Pos()
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+440, basePos.Y+440), imgui.CondOnce, imgui.NewVec2(0, 0))
	imgui.SetNextWindowSizeV(imgui.NewVec2(650, 400), imgui.CondOnce)

	imgui.Begin("ImNodes Demo")
	imnodes.EditorContextSet(ctx)
	if first {
		first = false
		imnodes.LoadCurrentEditorStateFromIniFile("imnodes.ini")
	}

	imnodes.BeginNodeEditor()
	if imgui.IsItemHovered() && imgui.IsMouseDoubleClicked(imgui.MouseButtonLeft) {
		imnodes.StyleColorsClassic()
	}

	imgui.PushItemWidth(100)
	for _, node := range nodes {
		node.Show()
	}
	imgui.PopItemWidth()

	for _, link := range links {
		imnodes.Link(link.id, link.start, link.end)
	}

	imnodes.ClearMiniMapNodeHoveringCallbackPool()
	imnodes.MiniMapV(0.25, imnodes.MiniMapLocationBottomRight, func(arg0 int32, arg1 unsafe.Pointer) {}, imnodes.MiniMapNodeHoveringCallbackUserData{})

	imnodes.EndNodeEditor()

	imnodes.SaveEditorStateToIniFile(ctx, "imnodes.ini")

	{
		var link link
		if imnodes.IsLinkCreatedBoolPtr(&link.start, &link.end) {
			if ok, linkId := findLinkIdByEnd(link.end); ok {
				removeLink(linkId)
			}
			link.id = n.IdGen()
			if startNode := findOutputById(link.start); startNode != nil {
				if endNode := findInputById(link.end); endNode != nil {
					link.startNode = startNode
					link.endNode = endNode
					links = append(links, link)
				}
			}
		}
	}

	{
		var linkId int32
		if imnodes.IsLinkDestroyed(&linkId) {
			removeLink(linkId)
		}
	}

	imgui.End()
}

var Color [3]float32

func Get() [3]float32 {
	for _, node := range nodes {
		if node.Type() == n.NodeShowColor {
			if ok, color := node.GetOutput(0); ok {
				if final, ok := color.Data.([3]float32); ok {
					return final
				}
			}
		}
	}
	return Color
}
