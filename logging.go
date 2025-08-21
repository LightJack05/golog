package golog

func writeMessage(level logLevel, message string) {
	if applicationLogConfig.Configs == nil || applicationLogConfig.Configs.Len() == 0 {
		panic("No logging configuration set up. Please call Setup or AddConfig before logging messages.")
	}

	for e := applicationLogConfig.Configs.Front(); e != nil; e = e.Next() {
		config := e.Value.(Config)
		if stream, ok := config.StreamMap[level]; ok {
			formattedMessage := formatMessage(config, level, message)
			stream.WriteString(formattedMessage + "\n")
		}
	}
}
