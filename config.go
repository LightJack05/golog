package golog

import "os"
import "container/list"

type ConfigStreamMap map[logLevel]*os.File
type ConfigLabelMap map[logLevel]string

type ApplicationLogConfig struct {
	ApplicationName string
	Configs         *list.List
}

type Config struct {
	MessageFormat string

	StreamMap   map[logLevel]*os.File
	LogLabelMap map[logLevel]string
}

var applicationLogConfig ApplicationLogConfig

// checkForDoubleSetup Check if the application log configuration has already been set up. If it has, panic to prevent double setup.
func checkForDoubleSetup() {
	if applicationLogConfig.Configs != nil && applicationLogConfig.Configs.Len() > 0 {
		panic("Application log configuration has already been set up. Use AddConfig to add additional configurations.")
	}
}

func isLoggingConfigured() bool {
	return applicationLogConfig.Configs != nil && applicationLogConfig.Configs.Len() > 0
}

// Setup Setup the application with a single configuration and an application name.
func Setup(applicationName string, config Config) {
	checkForDoubleSetup()
	applicationLogConfig.ApplicationName = applicationName
	applicationLogConfig.Configs = list.New()
	applicationLogConfig.Configs.PushBack(config)
}

// SetupAnonymous Setup the application with a single configuration without an application name.
func SetupAnonymous(config Config) {
	checkForDoubleSetup()
	applicationLogConfig.ApplicationName = ""
	applicationLogConfig.Configs = list.New()
	applicationLogConfig.Configs.PushBack(config)
}

// SetupDefaultConsole Setup the application with the default console configuration.
func SetupDefaultConsole() {
	checkForDoubleSetup()
	SetupAnonymous(GetDefaultConfigConsole())
}

// AddConfig Add a new configuration to the application log configuration. This allows for multiple configurations to be used, which can be useful for different logging contexts or environments.
func AddConfig(config Config) {
	if applicationLogConfig.Configs == nil {
		applicationLogConfig.Configs = list.New()
	}
	applicationLogConfig.Configs.PushBack(config)
}
