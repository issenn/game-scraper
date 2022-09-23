package zap

import (
	// "go.uber.org/zap"
	// "go.uber.org/zap/zapcore"
)


type BaseLoggingConfig struct {
	// Level            AtomicLevel `mapstructure:"level" json:"level" yaml:"level" toml:"level"`
	// LevelCmp         string          `mapstructure:"levelCmp" json:"levelCmp" yaml:"levelCmp" toml:"levelCmp"`

	// // Log format, one of json or plain-text.
	// Encoding         string          `mapstructure:"encoding" json:"encoding" yaml:"encoding" toml:"encoding"`
}

type baseLoggerConfig struct {
}

type baseCoreConfig struct {
}



// type Config zap.Config


// func NewProductionConfig() zap.Config {
// 	return zap.Config{
// 		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
// 		Development: false,
// 		Sampling: &zap.SamplingConfig{
// 			Initial:    100,
// 			Thereafter: 100,
// 		},
// 		Encoding:         "json",
// 		EncoderConfig:    zap.NewProductionEncoderConfig(),
// 		OutputPaths:      []string{"stderr"},
// 		ErrorOutputPaths: []string{"stderr"},
// 	}
// 	// return zap.NewProductionConfig()
// }

// func NewDevelopmentConfig() zap.Config {
// 	return zap.NewDevelopmentConfig()
// }
