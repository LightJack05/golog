package golog

func Debug(message string) {
	writeMessage(LOG_LEVEL_DEBUG, message)
}

func Info(message string) {
	writeMessage(LOG_LEVEL_INFO, message)
}

func Warning(message string) {
	writeMessage(LOG_LEVEL_WARNING, message)
}

func Error(message string) {
	writeMessage(LOG_LEVEL_ERROR, message)
}

func Fatal(message string, panicProgram bool) {
	writeMessage(LOG_LEVEL_FATAL, message)
	if panicProgram {
		panic("Program terminated due to fatal error: " + message)
	}
}
