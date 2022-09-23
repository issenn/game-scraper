package logger


import (
	"fmt"

	// "go.uber.org/zap"
	// "go.uber.org/zap/zapcore"

	"github.com/issenn/game-scraper/internal/config"
	// "github.com/issenn/game-scraper/pkg/logger"
)


func printConfig() {
	cfg := config.GetGlobalConfig()
	fmt.Printf("GetGlobalConfig: %#v\n", cfg)
	// cfg.Log.LevelCmp = "aaa"
	// fmt.Printf("GetGlobalConfig Level: %#v\n", cfg.Log.Level)
	// fmt.Printf("GetGlobalConfig Level: %#v\n", cfg.Log.LevelCmp)
	// config.StoreGlobalConfig(cfg)
	// fmt.Printf("GetGlobalConfig: %#v\n", config.GetGlobalConfig())
	// fmt.Printf("GetGlobalConfig Level: %#v\n", config.GetGlobalConfig().Log.Level)
	// fmt.Printf("GetGlobalConfig Level: %#v\n", config.GetGlobalConfig().Log.LevelCmp)
}

func testConfig() {
	printConfig()
}

func InitLogger() {
	// testConfig()
	if ok := config.LogEnabled(); ok {
		// // fmt.Printf("GetGlobalConfig in logger: %#v\n", config.GetGlobalConfig().Log.Logger[0].Core[0])
		// logConfig := config.GetGlobalConfig().Log
		// // level := logConfig.Level
		// // levelCmp := logConfig.LevelCmp
		// // development := logConfig.Development
		// // encoding := logConfig.Encoding
		// // logging := logConfig.Logging
		// loggersConfig := logConfig.Logger
		// for _, loggerConfig := range loggersConfig {
		// 	name := loggerConfig.Name
		// 	fmt.Println(name)
		// 	logging := loggerConfig.Logging
		// 	fmt.Printf("value: [%v]\n", logging)
		// 	// level := loggerConfig.Level
		// 	// levelCmp := loggerConfig.LevelCmp
		// 	coreConfigs := loggerConfig.Core
		// 	for _, coreConfig := range coreConfigs {
		// 		use := coreConfig.Use
		// 		fmt.Printf("value: [%v]\n", use)
		// 		break
		// 	}
		// 	break
		// }
		// zapConfig := zap.Config{
		// 	Level: loggerConfig.Level,
		// 	Development: loggerConfig.Development,
		// 	Encoding: loggerConfig.Encoding,
		// 	EncoderConfig: zapcore.EncoderConfig{
		// 		MessageKey: loggerConfig.Core[0].EncoderConfig.MessageKey,
		// 		LevelKey: loggerConfig.Core[0].EncoderConfig.LevelKey,
		// 		EncodeLevel: loggerConfig.Core[0].EncoderConfig.EncodeLevel,
		// 	},
		// 	OutputPaths: loggerConfig.Core[0].OutputPaths,
		// 	ErrorOutputPaths: loggerConfig.ErrorOutputPaths,
		// }
		// logger := zap.Must(zapConfig.Build())
		// defer logger.Sync()
		// logger.Info("logger construction succeeded")
	}
}
