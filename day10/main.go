package main

import "fmt"
import "encoding/hex"

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

func hash1(input []byte) int {
	result := hash(256, input, 1)
	return int(result[0]) * int(result[1])
}

func hash64(input string) string {
	lengths := append([]byte(input), []byte{17, 31, 73, 47, 23}...)
	result := hash(256, lengths, 64)
	return hex.EncodeToString(collapse(result))
}

func main() {
	{
		// Part 1: Ingest the input bytes directly and hash once.
		fmt.Println(hash1([]byte{88, 88, 211, 106, 141, 1, 78, 254, 2, 111, 77, 255, 90, 0, 54, 205}))
	}

	{
		// Part 2: Ingest as ASCII with a trailer and hash 64 times.
		fmt.Println(hash64(""))
		fmt.Println(hash64("AoC 2017"))
		fmt.Println(hash64("88,88,211,106,141,1,78,254,2,111,77,255,90,0,54,205"))
	}
}
