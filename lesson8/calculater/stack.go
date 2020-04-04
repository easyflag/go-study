package main

import "fmt"

//StackF64 xxx
type StackF64 struct {
	data [32]float64
	top  int
}

//push push in one data
func (s *StackF64) push(d float64) {
	s.data[s.top] = d
	s.top++
}

//pop pop out one data
func (s *StackF64) pop() (float64, error) {
	if s.top == 0 {
		err := fmt.Errorf("stack is empty")
		return 0, err
	}
	s.top--
	return s.data[s.top], nil
}

//getTop get top data
func (s *StackF64) getTop() (float64, error) {
	if s.top == 0 {
		err := fmt.Errorf("stack is empty")
		return 0, err
	}
	return s.data[s.top-1], nil
}

//clear clear all data
func (s *StackF64) clear() {
	for i := 0; i < len(s.data); i++ {
		s.data[i] = 0
	}
}

//StackStr xxx
type StackStr struct {
	data [32]string
	top  int
}

//push push in one data
func (s *StackStr) push(str string) {
	s.data[s.top] = str
	s.top++
}

//pop pop out one data
func (s *StackStr) pop() (string, error) {
	if s.top == 0 {
		err := fmt.Errorf("stack is empty")
		return "", err
	}
	s.top--
	return s.data[s.top], nil
}

//getTop get top data
func (s *StackStr) getTop() (string, error) {
	if s.top == 0 {
		err := fmt.Errorf("stack is empty")
		return "", err
	}
	return s.data[s.top-1], nil
}

//clear clear all data
func (s *StackStr) clear() {
	for i := 0; i < len(s.data); i++ {
		s.data[i] = ""
	}
}
