package main

var opPriority map[byte]int

func opPriorityInit() {
	opPriority = make(map[byte]int)
	opPriority['('] = 0
	opPriority[')'] = 0
	opPriority['+'] = 1
	opPriority['-'] = 1
	opPriority['*'] = 2
	opPriority['/'] = 2
	opPriority['^'] = 3
}
