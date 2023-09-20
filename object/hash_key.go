package object

import "hash/fnv"

type HashKey uint64

func (b *Boolean) HashKey() HashKey {
	if b.Value {
		return 1
	} else {
		return 0
	}
}

func (i Number[T]) HashKey() HashKey {
	return HashKey(i.Value)
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey(h.Sum64())
}
