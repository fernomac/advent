package main

type buffer []byte

func (b buffer) Swap(i, j int) {
	i = i % len(b)
	j = j % len(b)
	b[i], b[j] = b[j], b[i]
}

func (b buffer) Reverse(from, to int) {
	for from < to {
		b.Swap(from, to)
		from++
		to--
	}
}

func hash(size int, lengths []byte, rounds int) []byte {
	buf := buffer(make([]byte, size))
	for i := 0; i < size; i++ {
		buf[i] = byte(i)
	}

	cur := 0
	skip := 0

	for i := 0; i < rounds; i++ {
		for _, length := range lengths {
			buf.Reverse(cur, (cur+int(length))-1)
			cur += (skip + int(length))
			skip++
		}
	}

	return buf
}

func collapse(input []byte) []byte {
	result := make([]byte, 16)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			result[i] ^= input[(i*16)+j]
		}
	}
	return result
}

func knothash(input string) []byte {
	lengths := append([]byte(input), []byte{17, 31, 73, 47, 23}...)
	result := hash(256, lengths, 64)
	return collapse(result)
}
