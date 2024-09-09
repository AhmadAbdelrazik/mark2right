package app

import "errors"

type Stack []string

var (
	ErrEmptyStack = errors.New("stack is full")
)

func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s Stack) Push(n string) {
	s = append(s, n)
}

func (s Stack) Pop() {
	if len(s) == 0 {
		return
	}

	s = s[:len(s)-1]
}

func (s Stack) Top() (string, error) {
	if s.Empty() {
		return "", ErrEmptyStack
	}

	return s[len(s)-1], nil
}
