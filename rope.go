package rope

import (
	"github.com/ksco/rope/i"
	"github.com/ksco/rope/internal"
)

func NewRope(v []byte) i.Rope {
	return internal.NewLeaf(v)
}

func NewRopeFromStr(v string) i.Rope {
	return internal.NewLeaf([]byte(v))
}

func doAppend(cb func(v any) i.Rope, l i.Rope, rest []any) i.Rope {
	res := l
	for _, r := range rest {
		res = internal.ConcatTwo(res, cb(r))
	}
	return res
}

func toAnySlice[T any](vs []T) []any {
	res := make([]any, 0, len(vs))
	for _, v := range vs {
		res = append(res, v)
	}
	return res
}

func Append(l i.Rope, rest ...i.Rope) i.Rope {
	return doAppend(func(v any) i.Rope {
		return v.(i.Rope)
	}, l, toAnySlice(rest))
}

func AppendBytes(l i.Rope, rest ...[]byte) i.Rope {
	return doAppend(func(v any) i.Rope {
		return NewRope(v.([]byte))
	}, l, toAnySlice(rest))
}

func AppendStr(l i.Rope, rest ...string) i.Rope {
	return doAppend(func(v any) i.Rope {
		return NewRope([]byte(v.(string)))
	}, l, toAnySlice(rest))
}

func Insert(r i.Rope, i int, v i.Rope) i.Rope {
	if i >= Len(r) {
		return Append(r, v)
	}

	s := r.Split(i).(*internal.Concat)
	return Append(s.L(), v, s.R())
}

func Bytes(r i.Rope) []byte {
	return r.Val()
}

func Str(r i.Rope) string {
	return string(Bytes(r))
}

func Len(r i.Rope) int {
	return r.Weight()
}

func Sub(r i.Rope, i, j int) i.Rope {
	return r.Sub(i, j)
}
