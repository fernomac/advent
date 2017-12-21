package main

type state struct {
	buffer []byte
	start  int
}

func (s *state) Spin(n int) {
	temp := s.start - n
	for temp < 0 {
		temp += len(s.buffer)
	}
	s.start = temp
}

func (s *state) Exchange(a, b int) {
	ai := (a + s.start) % len(s.buffer)
	bi := (b + s.start) % len(s.buffer)
	s.buffer[ai], s.buffer[bi] = s.buffer[bi], s.buffer[ai]
}

func (s *state) Partner(a, b byte) {
	for i := 0; i < len(s.buffer); i++ {
		if s.buffer[i] == a {
			s.buffer[i] = b
		} else if s.buffer[i] == b {
			s.buffer[i] = a
		}
	}
}

func (s *state) String() string {
	return string(s.buffer[s.start:]) + string(s.buffer[:s.start])
}
