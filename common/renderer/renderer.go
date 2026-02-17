package renderer

import (
	"CimguiTest/util"
	_ "embed"
	"fmt"
	"strings"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/go-gl/gl/v3.3-core/gl"
)

const FLOAT_SIZE = 4

//go:embed shaders/triangle.vert
var vertexShader string

//go:embed shaders/triangle.frag
var fragmentShader string

type uniforms struct {
	color int32
}

type renderer struct {
	shaderHandle uint32
	VAO, VBO     uint32
	uniforms     uniforms
}

type windowScreen struct {
	width, height int32
}

type frameBuffer struct {
	FBO           uint32
	colorBuffer   uint32
	depth         uint32
	width, height int32
}

var w windowScreen
var r renderer
var f frameBuffer

func glError(handle uint32, statusType uint32, getIV func(uint32, uint32, *int32), getInfoLog func(uint32, int32, *int32, *uint8), failureMsg string) {
	var status int32
	getIV(handle, statusType, &status)
	if status == gl.FALSE {
		var logLength int32
		getIV(handle, gl.INFO_LOG_LENGTH, &logLength)

		infoLog := strings.Repeat("\x00", int(logLength))
		getInfoLog(handle, logLength, nil, gl.Str(infoLog))
		fmt.Println(failureMsg+"\n", infoLog)
	}
}

func Init() {
	glShaderSource := func(handle uint32, source string) {
		csource, free := gl.Strs(source + "\x00")
		defer free()

		gl.ShaderSource(handle, 1, csource, nil)
	}

	if err := gl.Init(); err != nil {
		panic(err)
	}

	w.width = 1200
	w.height = 900
	f.width = 100
	f.height = 100

	gl.ClearColor(0.2, 0.3, 0.3, 1.0)
	gl.Enable(gl.CULL_FACE)
	gl.CullFace(gl.BACK)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.GenTextures(1, &f.colorBuffer)
	gl.BindTexture(gl.TEXTURE_2D, f.colorBuffer)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGB16F, f.width, f.height, 0, gl.RGB, gl.UNSIGNED_BYTE, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.ActiveTexture(gl.TEXTURE1)

	gl.GenTextures(1, &f.depth)
	gl.BindTexture(gl.TEXTURE_2D, f.depth)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.DEPTH_COMPONENT32, f.width, f.height, 0, gl.DEPTH_COMPONENT, gl.UNSIGNED_INT, nil)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_BORDER)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_BORDER)
	borderColor := []float32{1, 1, 1, 1}
	gl.TexParameterfv(gl.TEXTURE_2D, gl.TEXTURE_BORDER_COLOR, &borderColor[0])
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.BindTexture(gl.TEXTURE_2D, 0)

	r.shaderHandle = gl.CreateProgram()
	vertHandle := gl.CreateShader(gl.VERTEX_SHADER)
	fragHandle := gl.CreateShader(gl.FRAGMENT_SHADER)
	glShaderSource(vertHandle, vertexShader)
	glShaderSource(fragHandle, fragmentShader)
	gl.CompileShader(vertHandle)
	glError(vertHandle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog, "Vertex shader error")
	gl.CompileShader(fragHandle)
	glError(vertHandle, gl.COMPILE_STATUS, gl.GetShaderiv, gl.GetShaderInfoLog, "Fragment shader error")
	gl.AttachShader(r.shaderHandle, vertHandle)
	gl.AttachShader(r.shaderHandle, fragHandle)
	gl.LinkProgram(r.shaderHandle)
	glError(r.shaderHandle, gl.LINK_STATUS, gl.GetProgramiv, gl.GetProgramInfoLog, "Linking program error")
	gl.DeleteShader(vertHandle)
	gl.DeleteShader(fragHandle)

	r.uniforms.color = gl.GetUniformLocation(r.shaderHandle, util.Str("color"))

	vertices := []float32{
		-0.5, -0.5, 0.0,
		0.5, -0.5, 0.0,
		0.0, 0.5, 0.0,
	}

	gl.GenVertexArrays(1, &r.VAO)
	gl.GenBuffers(1, &r.VBO)
	gl.BindVertexArray(r.VAO)

	gl.BindBuffer(gl.ARRAY_BUFFER, r.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, int(FLOAT_SIZE)*len(vertices), gl.Ptr(&vertices[0]), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*FLOAT_SIZE, nil)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	gl.GenFramebuffers(1, &f.FBO)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FBO)

	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, f.colorBuffer, 0)
	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, f.depth, 0)

	if gl.CheckFramebufferStatus(gl.FRAMEBUFFER) != gl.FRAMEBUFFER_COMPLETE {
		fmt.Println(gl.CheckFramebufferStatus(gl.FRAMEBUFFER))
		panic("Framebuffer error")
	}

	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.BindTexture(gl.TEXTURE_2D, 0)
	gl.BindRenderbuffer(gl.RENDERBUFFER, 0)
}

func Render(clearColor [3]float32, objectColor [3]float32) {
	gl.Viewport(0, 0, f.width, f.height)
	gl.BindFramebuffer(gl.FRAMEBUFFER, f.FBO)
	gl.FramebufferTexture2D(gl.FRAMEBUFFER, gl.COLOR_ATTACHMENT0, gl.TEXTURE_2D, f.colorBuffer, 0)
	gl.FramebufferTexture(gl.FRAMEBUFFER, gl.DEPTH_ATTACHMENT, f.depth, 0)

	gl.ClearColor(clearColor[0], clearColor[1], clearColor[2], 1.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.UseProgram(r.shaderHandle)
	gl.Uniform3f(r.uniforms.color, objectColor[0], objectColor[1], objectColor[2])
	gl.BindVertexArray(r.VAO)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.BindFramebuffer(gl.FRAMEBUFFER, 0)
	gl.Viewport(0, 0, w.width, w.height)
	gl.Clear(gl.COLOR_BUFFER_BIT)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func Nuke() {
	gl.UseProgram(r.shaderHandle)
	gl.DeleteVertexArrays(1, &r.VAO)
	gl.DeleteBuffers(1, &r.VBO)
	gl.DeleteProgram(r.shaderHandle)
}

func Image() uint32 {
	return f.colorBuffer
}

func Size() imgui.Vec2 {
	return imgui.NewVec2(float32(f.width), float32(f.height))
}
