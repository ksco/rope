package i

type Rope interface {
	Weight() int
	Val() []byte
	Split(i int) Rope
	Sub(i, j int) Rope
}
