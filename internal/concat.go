package internal

import (
	"github.com/ksco/rope/i"
)

type Concat struct {
	l i.Rope
	r i.Rope
	w int
}

func ConcatTwo(l, r i.Rope) *Concat {
	return &Concat{l, r, l.Weight() + r.Weight()}
}

func (c *Concat) Weight() int {
	return c.w
}

func (c *Concat) Val() []byte {
	return append(c.l.Val(), c.r.Val()...)
}

func (c *Concat) Split(i int) i.Rope {
	lw := c.l.Weight()
	if i == lw {
		return c
	} else if i < lw {
		c1 := c.l.Split(i).(*Concat)
		return &Concat{c1.l, ConcatTwo(c1.r, c.r), c.w}
	} else {
		c1 := c.r.Split(i - lw).(*Concat)
		return &Concat{ConcatTwo(c.l, c1.l), c1.r, c.w}
	}
}

func (c *Concat) Sub(i, j int) i.Rope {
	return c.Split(i).(*Concat).r.Split(j - i).(*Concat).l
}

func (c *Concat) L() i.Rope {
	return c.l
}

func (c *Concat) R() i.Rope {
	return c.r
}
