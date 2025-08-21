package golog

import (
	"strings"
	"time"
)

type logMessagePlaceholder int

const (
	PLACEHOLDER_APP_NAME  = 10
	PLACEHOLDER_LOG_LEVEL = 20
	PLACEHOLDER_TIMESTAMP = 30
	PLACEHOLDER_MESSAGE   = 40
)

var placeholderStringMap = map[logMessagePlaceholder]string{
	PLACEHOLDER_APP_NAME:  "{{APP_NAME}}",
	PLACEHOLDER_LOG_LEVEL: "{{LOG_LEVEL}}",
	PLACEHOLDER_TIMESTAMP: "{{TIMESTAMP}}",
	PLACEHOLDER_MESSAGE:   "{{MESSAGE}}",
}

var placeholderValueMap = map[logMessagePlaceholder]func(config Config, level logLevel, message string) string{
	PLACEHOLDER_APP_NAME: func(config Config, level logLevel, message string) string {
		return applicationLogConfig.ApplicationName
	},
	PLACEHOLDER_LOG_LEVEL: func(config Config, level logLevel, message string) string { return config.LogLabelMap[level] },
	PLACEHOLDER_TIMESTAMP: func(config Config, level logLevel, message string) string { return time.Now().Format(time.RFC3339) },
	PLACEHOLDER_MESSAGE:   func(config Config, level logLevel, message string) string { return message },
}

// Replace placeholders in the format string with actual values
func formatMessage(config Config, level logLevel, message string) string {
	format := config.MessageFormat
	for placeholder, placeholderStr := range placeholderStringMap {
		if placeholderValueFunc, exists := placeholderValueMap[placeholder]; exists {
			value := placeholderValueFunc(config, level, message)
			format = strings.ReplaceAll(format, placeholderStr, value)
		}
	}
	return format
}
