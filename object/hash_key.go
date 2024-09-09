package object

import "hash/fnv"

type HashKey uint64

func (b *Boolean) HashKey() HashKey {
	if b.Value {
		return 1
	}
	return 0
}

func (n Number) HashKey() HashKey {
	return HashKey(n.Value)
}

func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return HashKey(h.Sum64())
}
