package main

import "os"

func GetDefaultStreamMapConsole() ConfigStreamMap {
	return ConfigStreamMap{
		LOG_LEVEL_DEBUG:   os.Stderr,
		LOG_LEVEL_INFO:    os.Stderr,
		LOG_LEVEL_WARNING: os.Stderr,
		LOG_LEVEL_ERROR:   os.Stderr,
		LOG_LEVEL_FATAL:   os.Stderr,
	}
}

func GetDefaultLogLabelMapConsole() ConfigLabelMap {
	return ConfigLabelMap{
		LOG_LEVEL_DEBUG:   "DEBUG",
		LOG_LEVEL_INFO:    "INFO",
		LOG_LEVEL_WARNING: "WARNING",
		LOG_LEVEL_ERROR:   "ERROR",
		LOG_LEVEL_FATAL:   "FATAL",
	}
}

func GetDefaultConfigConsole() Config {
	return Config{
		MessageFormat:   "[{{LOG_LEVEL}}] {{MESSAGE}}",
		StreamMap:       GetDefaultStreamMapConsole(),
		LogLabelMap:     GetDefaultLogLabelMapConsole(),
	}
}

//TODO: Add Logfmt and FileStream default configs
