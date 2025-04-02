package config

import "sample-proj/config"

// function that calls also the other functions to initialize the application
// such as db connections and port and debug level
func LoadConfig(devopsConfig config.RemoteConfig) {
    config.GetConfig(devopsConfig)
}
