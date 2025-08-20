package main

import "os"

type Config struct {
	ApplicationName string
	MessageFormat   string

	StreamMap   map[LOG_LEVEL]*os.File
	LogLabelMap map[LOG_LEVEL]string
}

func GetDefaultStreamMap() map[LOG_LEVEL]*os.File {
	return map[LOG_LEVEL]*os.File{
		LOG_LEVEL_DEBUG:   os.Stderr,
		LOG_LEVEL_INFO:    os.Stderr,
		LOG_LEVEL_WARNING: os.Stderr,
		LOG_LEVEL_ERROR:   os.Stderr,
		LOG_LEVEL_FATAL:   os.Stderr,
	}
}

func GetDefaultLogLabelMap() map[LOG_LEVEL]string {
	return map[LOG_LEVEL]string{
		LOG_LEVEL_DEBUG:   "DEBUG",
		LOG_LEVEL_INFO:    "INFO",
		LOG_LEVEL_WARNING: "WARNING",
		LOG_LEVEL_ERROR:   "ERROR",
		LOG_LEVEL_FATAL:   "FATAL",
	}
}

func GetDefaultConfig() Config {
	return Config{
		ApplicationName: "",
		MessageFormat:   "[{{LOG_LEVEL}}] {{MESSAGE}}",
		StreamMap:       GetDefaultStreamMap(),
		LogLabelMap:     GetDefaultLogLabelMap(),
	}
}
