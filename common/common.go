// Package common contains code commonly used in the examples section of cimgui-go.
package common

import (
	"CimguiTest/common/renderer"
	_ "embed"
	"fmt"

	cte "github.com/AllenDang/cimgui-go/ImGuiColorTextEdit"
	"github.com/AllenDang/cimgui-go/imgui"

	_ "github.com/AllenDang/cimgui-go/immarkdown"
	_ "github.com/AllenDang/cimgui-go/imnodes"
	_ "github.com/AllenDang/cimgui-go/impl/opengl3"
	"github.com/AllenDang/cimgui-go/implot"
	"github.com/AllenDang/cimgui-go/utils"
)

var (
	showDemoWindow bool
	value1         int32
	value2         int32
	value3         int32
	values         [2]int32 = [2]int32{value1, value2}
	content        string   = "Let me try"
	r              float32
	g              float32
	b              float32
	a              float32
	color4         [4]float32 = [4]float32{r, g, b, a}
	selected       bool
	barValues      []int64
	textEditor     *cte.TextEditor
)

// Initialize prepares global variables. Call before anything related to backend.
func Initialize() {
	for i := 0; i < 10; i++ {
		barValues = append(barValues, int64(i+1))
	}
}

func InputTextCallback(data imgui.InputTextCallbackData) int {
	fmt.Println("got call back")
	return 0
}

var clearColor [3]float32
var objectColor [3]float32

func AfterCreateContext() {
	implot.CreateContext()
	textEditor = cte.NewTextEditor()
	textEditor.SetLanguageDefinition(cte.Cpp)
	textEditor.SetText(`// Colorize a C++ file
#include <iostream>
ImGui::Text("Hello World")`)

	clearColor = [3]float32{0.0, 0.0, 0.0}
	objectColor = [3]float32{1.0, 0.5, 0.2}

	renderer.Init()
}

func BeforeDestroyContext() {
	renderer.Nuke()
	implot.DestroyContext()
}

func Loop() {
	renderer.Render(clearColor, objectColor)
	imgui.ClearSizeCallbackPool()

	// viewport := imgui.MainViewport()
	// imgui.SetNextWindowPos(viewport.Pos())
	// imgui.SetNextWindowSize(viewport.Size())
	// flags := imgui.WindowFlagsNoDecoration | imgui.WindowFlagsNoMove
	// if imgui.BeginV("Fullscreen window", nil, flags) {

	// }
	// imgui.End()
	ShowWidgetsDemo()
	ShowPictureLoadingDemo()
	ShowImPlotDemo()
	ShowCTEDemo()
}

func ShowWidgetsDemo() {
	if showDemoWindow {
		imgui.ShowDemoWindowV(&showDemoWindow)
	}

	imgui.SetNextWindowSizeV(imgui.NewVec2(300, 300), imgui.CondOnce)

	imgui.SetNextWindowSizeConstraintsV(imgui.NewVec2(300, 300), imgui.NewVec2(500, 500), func(data *imgui.SizeCallbackData) {
	}, 0)

	imgui.Begin("Window 1")
	if imgui.ButtonV("Click Me", imgui.NewVec2(80, 20)) {
		fmt.Println("Button clicked")
	}
	imgui.TextUnformatted("Unformatted text")
	imgui.Checkbox("Show demo window", &showDemoWindow)
	if imgui.BeginCombo("Combo", "Combo preview") {
		imgui.SelectableBoolPtr("Item 1", &selected)
		imgui.SelectableBool("Item 2")
		imgui.SelectableBool("Item 3")
		imgui.EndCombo()
	}

	if imgui.RadioButtonBool("Radio button1", selected) {
		selected = true
	}

	imgui.SameLine()

	if imgui.RadioButtonBool("Radio button2", !selected) {
		selected = false
	}

	imgui.InputTextWithHint("Name", "write your name here", &content, 0, InputTextCallback)
	imgui.Text(content)
	imgui.SliderInt("Slider int", &value3, 0, 100)
	imgui.DragInt("Drag int", &value1)
	imgui.DragInt2("Drag int2", &values)
	value1 = values[0]
	imgui.ColorEdit4("Color Edit3", &color4)
	imgui.End()
}

func ShowPictureLoadingDemo() {
	// demo of showing a picture
	tex := *imgui.NewTextureRefTextureID(imgui.TextureID(renderer.Image()))
	basePos := imgui.MainViewport().Pos()
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+60, 600), imgui.CondOnce, imgui.NewVec2(0, 0))
	imgui.SetNextWindowSizeV(imgui.NewVec2(300, 185), imgui.CondOnce)
	imgui.Begin("Image")
	imgui.ImageWithBgV(tex, renderer.Size(), imgui.NewVec2(0, 1), imgui.NewVec2(1, 0), imgui.NewVec4(0, 0, 0, 0), imgui.NewVec4(1, 1, 1, 1))
	imgui.ColorEdit3("clear color", &clearColor)   // Edit 3 floats representing a color
	imgui.ColorEdit3("object color", &objectColor) // Edit 3 floats representing a color
	imgui.End()
}

func ShowImPlotDemo() {
	basePos := imgui.MainViewport().Pos()
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+400, basePos.Y+60), imgui.CondOnce, imgui.NewVec2(0, 0))
	imgui.SetNextWindowSizeV(imgui.NewVec2(500, 300), imgui.CondOnce)
	imgui.Begin("Plot window")
	if implot.BeginPlotV("Plot", imgui.NewVec2(-1, -1), 0) {
		implot.PlotBarsS64PtrInt("Bar", utils.SliceToPtr(barValues), int32(len(barValues)))
		implot.PlotLineS64PtrInt("Line", utils.SliceToPtr(barValues), int32(len(barValues)))
		implot.EndPlot()
	}
	imgui.End()
}

func ShowCTEDemo() {
	basePos := imgui.MainViewport().Pos()
	imgui.SetNextWindowPosV(imgui.NewVec2(basePos.X+800, basePos.Y+260), imgui.CondOnce, imgui.NewVec2(0, 0))
	imgui.SetNextWindowSizeV(imgui.NewVec2(250, 400), imgui.CondOnce)
	imgui.Begin("Color Text Edit")

	if textEditor.Render("Color Text Edit") {
	}

	imgui.End()
}
