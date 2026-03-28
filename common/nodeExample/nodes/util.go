package nodes

var lastId int32 = 0

func IdGen() int32 {
	lastId += 1
	return lastId
}

type DataType int

const (
	t_null DataType = iota
	t_float32
	t_3float32
)

type Value struct {
	typ  DataType
	data any
}
