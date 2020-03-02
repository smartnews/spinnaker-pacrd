package main

import (
	"github.com/spf13/viper"
)

// PacrdConfig represents the valid configuration options for PaCRD
type PacrdConfig struct {
	SpinnakerServices
}

// SpinnakerServices represent the set of services that PaCRD must interface with.
type SpinnakerServices struct {
	Front50 string
	Orca    string
}

// InitConfig initializes configuration for PaCRD.
func InitConfig() (PacrdConfig, error) {

	var config PacrdConfig

	viper.SetConfigName("pacrd")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/opt/pacrd")
	viper.SetDefault("SpinnakerServices", SpinnakerServices{
		Front50: "http://spin-front50:8080",
		Orca:    "http://spin-orca:8083",
	})

	if err := viper.ReadInConfig(); err != nil {
		// Ignore not found errors; we'll use the defaults in this case.
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return PacrdConfig{}, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return PacrdConfig{}, err
	}

	return config, nil
}
