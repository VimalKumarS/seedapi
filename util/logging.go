package util

import (
	"seedapi/model"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger used for the whole application.
var Logger *zap.Logger

// func init() {
// 	Logger, _ = zap.NewProduction()
// }

// SetupLogging 
func SetupLogging(config *model.AppSetting){
	emptyLogging := model.Logging{}
	if emptyLogging == config.Logging  {
		Logger, _ = zap.NewProduction()
		return
	}

	w := zapcore.AddSync(&lumberjack.Logger{
		Filename : config.Logging.Filename,
		MaxSize : config.Logging.MaxSize,
	})
	
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	// Logger, _  := zap.NewProductionConfig().Build()
	// // Logger, _ = zap.NewProduction()
	// Logger.WithOptions( 
	// 	zap.WrapCore(
	// 		func(zapcore.Core) zapcore.Core{
	// 		return core
	// }))
	
	Logger = zap.New(core)
	//zap.AddStacktrace(zap.ErrorLevel)
}
