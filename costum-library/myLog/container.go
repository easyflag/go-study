package mylog

type logData struct {
	time      string
	date      string
	level     uint8
	msg       string
	filenName string
	funcName  string
	lineNum   int
}

type logQueue struct {
	queue chan *logData
}

var logQueues *logQueue

func newLogQueue(amount int) *logQueue {
	q := new(logQueue)
	q.queue = make(chan *logData, amount)
	return q
}
