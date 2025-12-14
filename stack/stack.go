package main

import (
	"fmt"
	"strings"
)

// É uma coleção de elementos que são inseridos
// e removidos seguindo a premissa de que o
// último a entrar é o primeiro a sair (LIFO).
type Stack struct {
	size  int
	top   int
	items []int
}

// Adiciona i no topo da pilha
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
	s.size = len(s.items)
	s.top = s.size - 1
}

// Remove e retorna o elemento do topo da
// pilha
func (s *Stack) Pop() int {
	top := s.items[s.top]

	s.size -= 1
	s.top -= 1

	s.items = s.items[:s.size]

	return top
}

// Retorna o tamanho da pilha
func (s Stack) Size() int {
	return s.size
}

func (s Stack) String() string {
	var str strings.Builder

	if s.size == 0 {
		return "| pilha vazia |\n"
	}

	for i := s.top; i >= 0; i-- {
		if i == s.top {
			fmt.Fprintf(&str, "| %2d | <- topo\n", s.items[i])
		} else {
			fmt.Fprintf(&str, "| %2d |\n", s.items[i])
		}
	}

	return str.String()
}

func main() {
	var s Stack

	for i := range 100 {
		s.Push(i)
		fmt.Println(s)
	}

	for s.Size() > 0 {
		s.Pop()
		fmt.Println(s)
	}
}
