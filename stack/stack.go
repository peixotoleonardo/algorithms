package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// É uma coleção de elementos que são inseridos
// e removidos seguindo a premissa de que o
// último a entrar é o primeiro a sair (LIFO).
type Stack[T any] struct {
	items []T
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{items: make([]T, 0, capacity)}
}

// Adiciona i no topo da pilha
func (s *Stack[T]) Push(i T) *Stack[T] {
	s.items = append(s.items, i)

	return s
}

// Remove e retorna o elemento do topo da
// pilha
func (s *Stack[T]) Pop() (T, bool) {
	var zero T

	if len(s.items) == 0 {
		return zero, false
	}

	last := len(s.items) - 1

	v := s.items[last]

	s.items[last] = zero

	s.items = s.items[:last]

	return v, true
}

// Retorna o tamanho da pilha
func (s Stack[T]) Size() int {
	return len(s.items)
}

func (s Stack[T]) String() string {
	var str strings.Builder

	if len(s.items) == 0 {
		return "| pilha vazia |\n"
	}

	for i := len(s.items) - 1; i >= 0; i-- {
		fmt.Fprintf(&str, "| %2v |\n", s.items[i])
	}

	return str.String()
}

func main() {
	s := NewStack[int](100)

	for i := range 10 {
		s.Push(i)
	}

	b := bufio.NewReader(os.Stdin)

	for {

		if _, ok := s.Pop(); ok {
			fmt.Println(s)
		} else {
			break
		}

		b.ReadString('\n')
	}
}
