package nodes

var LastId int32 = 0

func IdGen() int32 {
	LastId += 1
	return LastId
}

type DataType int

const (
	t_null DataType = iota
	t_float32
	t_3float32
)

type Value struct {
	Typ  DataType
	Data any
}

type NodeKind int

const (
	NodeConstant NodeKind = iota
	NodeShow
	NodeColor
	NodeShowColor
	NodeOscillator
	// n_WaveShaper
	NodeColorMixer
)
