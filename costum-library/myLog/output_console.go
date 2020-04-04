package mylog

import (
	"fmt"
	"os"
)

type consoleOutput struct {
	files *os.File
}

func newconsoleOutput() (c *consoleOutput) {
	//实例化对象
	c = new(consoleOutput)
	c.files = os.Stdout
	return
}

func (p *consoleOutput) run() {
	data := logQueues.queue
	for v := range data {
		writeInfo := fmt.Sprintf("(%s) %s:%s (%s in %s at line %d)\n", v.time,
			getLevelText(v.level), v.msg, v.funcName, v.filenName, v.lineNum)
		fmt.Fprintf(p.files, writeInfo)
	}
}

func (p *consoleOutput) closeAll() {
}
