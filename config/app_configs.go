package config

import ()

const logPrefix = `go-dating-`

var AppConfig AppConfigStruct

// the struct that stores the app configuration a custom type by us
type AppConfigStruct struct {
    Port int `json:"port"` 
    Debug bool `json:"debug"`
    Timezone string `json:"timezone"`
    Location string `json:"location"`
}

func GetAppConfig(appConfig AppConfigStruct) {
    AppConfig = appConfig
    
    // check for mandatory fields or throw exception
    // in the go project it uses a custom error handler to handle the exceptions well use the next best thing
    if AppConfig.Port == 0 {
        panic("Port is mandatory")
    }
    
    if AppConfig.Timezone == "" {
        panic("Timezone is mandatory")
    }
}
