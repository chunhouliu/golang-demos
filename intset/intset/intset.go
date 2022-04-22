package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Has(x int) bool {
	idx, offset := x/64, uint(x%64)
	return idx < len(s.words) && (s.words[idx]&(1<<offset)) != 0
}

func (s *IntSet) Len() int {
	cnt := 0
	for _, word := range s.words {
		for word != 0 {
			cnt++
			word &= word - 1
		}
	}
	return cnt
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var res IntSet
	res.UnionWith(s)
	return &res
}

func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

func (s *IntSet) Add(x int) {
	idx, offset := x/64, uint(x%64)
	for len(s.words) <= idx {
		s.words = append(s.words, 0)
	}
	s.words[idx] |= 1 << offset
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, word := range t.words {
		if i < len(s.words) {
			s.words[i] |= word
		} else {
			s.words = append(s.words, word)
		}
	}
}

func (s *IntSet) Intersect(t *IntSet) {
	for idx := range s.words {
		if idx >= len(t.words) {
			s.words[idx] = 0
		} else {
			s.words[idx] &= t.words[idx]
		}
	}
}

func (s *IntSet) Elems() []int {
	var res []int
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				res = append(res, i*64+j)
			}
		}
	}
	return res
}

func (s *IntSet) Remove(x int) {
	idx, offset := x/64, uint(x%64)
	if idx >= len(s.words) {
		return
	}
	s.words[idx] &= ^(1 << offset)
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", i*64+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
