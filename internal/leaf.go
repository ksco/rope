package internal

import "github.com/ksco/rope/i"

type Leaf struct {
	v []byte
}

func NewLeaf(v []byte) *Leaf {
	return &Leaf{v}
}

func (l *Leaf) Weight() int {
	return len(l.v)
}

func (l *Leaf) Val() []byte {
	return l.v
}

func (l *Leaf) Split(i int) i.Rope {
	return &Concat{&Leaf{l.v[:i]}, &Leaf{l.v[i:]}, l.Weight()}
}

func (l *Leaf) Sub(i, j int) i.Rope {
	return &Leaf{l.v[i:j]}
}
