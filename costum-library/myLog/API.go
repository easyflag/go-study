package mylog

//OpenByFile xxx
func OpenByFile(path string, name string, level int, amount int) {
	logQueues = newLogQueue(amount)
	inputer = newInPut(level)
	outputer = newFileOutput(path, name)
	go outputer.run()
}

//OpenByConsole xxx
func OpenByConsole(level int, amount int) {
	logQueues = newLogQueue(amount)
	inputer = newInPut(level)
	outputer = newconsoleOutput()
	go outputer.run()
}

//Close xxx
func Close() {
	outputer.closeAll()
}

//SetLevel xxx
func SetLevel(level int) {
	inputer.setLevel(level)
}

//Debug xxx
func Debug(format string, args ...interface{}) {
	inputer.debug(format, args...)
}

//Trace xxx
func Trace(format string, args ...interface{}) {
	inputer.trace(format, args...)
}

//Info xxx
func Info(format string, args ...interface{}) {
	inputer.info(format, args...)
}

//Warn xxx
func Warn(format string, args ...interface{}) {
	inputer.warn(format, args...)
}

//Error xxx
func Error(format string, args ...interface{}) {
	inputer.error(format, args...)
}

//Fatal xxx
func Fatal(format string, args ...interface{}) {
	inputer.fatal(format, args...)
}
