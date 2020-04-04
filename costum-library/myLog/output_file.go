package mylog

import (
	"fmt"
	"os"
	"time"
)

const (
	defaultFileSize = 10240
)

type fileStruct struct {
	number    uint
	creatDate string
	level     uint8
	size      int
	file      *os.File
}

type fileOutput struct {
	path  string
	name  string
	files [6]fileStruct
}

func newFileOutput(path string, name string) (f *fileOutput) {
	//测试路径是否合法
	fileName := fmt.Sprintf("%s\\%s.log", path, name)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(fmt.Sprintf("open %s failed,error:%s", fileName, err))
	}
	fmt.Fprintf(file, "path ok\n")
	file.Close()

	//实例化对象
	f = new(fileOutput)
	f.name = name
	f.path = path
	for i := range f.files {
		f.files[i].level = uint8(i)
		err = f.newFile(uint8(i))
		if err != nil {
			panic(fmt.Sprintf("error:%s", err))
		}
	}

	return
}

func (p *fileOutput) newFile(level uint8) error {
	if level < 0 || level > 5 {
		err := fmt.Errorf("level out bounds")
		return err
	}

	nowDate := time.Now().Format("2006-01-02")
	p.files[level].number++
	p.files[level].creatDate = nowDate
	p.files[level].size = 0
	fileName := fmt.Sprintf("%s\\%s.%s #%d.%s", p.path, p.name, nowDate,
		p.files[level].number, getLevelText(level))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		err = fmt.Errorf("open %s failed,error:%s", fileName, err)
		return err
	}
	p.files[level].file = file

	return nil
}

func (p *fileOutput) closeFile(level uint8) {
	p.files[level].file.Close()
}

func (p *fileOutput) run() {
	data := logQueues.queue
	for v := range data {
		fileStruct := &p.files[v.level]
		if fileStruct.creatDate != v.date {
			p.closeFile(v.level)
			fileStruct.number = 0
			p.newFile(v.level)
		} else if fileStruct.size > defaultFileSize {
			p.closeFile(v.level)
			p.newFile(v.level)
		}

		writeInfo := fmt.Sprintf("(%s) %s:%s (%s in %s at line %d)\n", v.time,
			getLevelText(v.level), v.msg, v.funcName, v.filenName, v.lineNum)
		fileStruct.size += len(writeInfo)
		fmt.Fprintf(fileStruct.file, writeInfo)
	}
}

func (p *fileOutput) closeAll() {
	for i := 0; i < 6; i++ {
		p.closeFile(uint8(i))
	}
}
