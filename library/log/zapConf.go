package log

import (
	"fmt"
	"time"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

// nolint 此函数未使用
func formatEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("%d%02d%02d_%02d%02d%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second()))
}

func FormateLog(args []interface{}) *zap.Logger {
	log := logger.With(ToJsonData(args))
	return log
}

func Debug(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Debugf(msg)
}

func Info(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Infof(msg)
}

func Warn(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Warnf(msg)
}

func Error(msg string, args ...interface{}) {
	FormateLog(args).Sugar().Errorf(msg)
}

func ToJsonData(args []interface{}) zap.Field {
	det := make([]string, 0)
	if len(args) > 0 {
		for _, v := range args {
			det = append(det, fmt.Sprintf("%+v", v))
		}
	}
	zap := zap.Any("detail", det)
	return zap
}

/*
ZapInit 初始化
https://github.com/natefinch/lumberjack
*/
func ZapInit() {
	hook := lumberjack.Logger{
		Filename:   viper.GetString("zap.filename"),
		MaxSize:    viper.GetInt("zap.maxSize"),
		MaxBackups: viper.GetInt("zap.maxBackups"),
		MaxAge:     viper.GetInt("zap.maxAge"),
		Compress:   viper.GetBool("zap.compress"),
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        viper.GetString("zap.timeKey"),
		LevelKey:       viper.GetString("zap.levelKey"),
		NameKey:        viper.GetString("zap.nameKey"),
		CallerKey:      viper.GetString("zap.callerKey"),
		MessageKey:     viper.GetString("zap.messageKey"),
		StacktraceKey:  viper.GetString("zap.stacktraceKey"),
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()

	atomicLevel.SetLevel(zap.InfoLevel)

	var core zapcore.Core = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
		atomicLevel,
	)
	logger = zap.New(core)
}

/*
ZapInstance
https://godoc.org/go.uber.org/zap
*/
/*
var Zapcfg zap.Config

func ZapInstance () (*zap.SugaredLogger)  {
	// For some users, the presets offered by the NewProduction, NewDevelopment,
	// and NewExample constructors won't be appropriate. For most of those
	// users, the bundled Config struct offers the right balance of flexibility
	// and convenience. (For more complex needs, see the AdvancedConfiguration
	// example.)
	//
	// See the documentation for Config and zapcore.EncoderConfig for all the
	// available options.
	rawJSON := []byte(`{
	  "level": "debug",
	  "encoding": "json",
	  "outputPaths": ["stdout", "/tmp/logs"],
	  "errorOutputPaths": ["stderr"],
	  "encoderConfig": {
	    "messageKey": "message",
	    "levelKey": "level",
	    "levelEncoder": "lowercase"
	  }
	}`)

	if err := json.Unmarshal(rawJSON, &Zapcfg); err != nil {
		panic(err)
	}
	logger, err := Zapcfg.Build()
	if err != nil {
		panic(err)
	}
	sugar := logger.Sugar()
	//return logger
	return sugar
}
*/

// ZapInit 初始化
/*
func ZapInit1() {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: viper.GetBool("zap.development"),
		Encoding:    viper.GetString("zap.encoding"),
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        viper.GetString("zap.timeKey"),
			LevelKey:       viper.GetString("zap.levelKey"),
			NameKey:        viper.GetString("zap.nameKey"),
			CallerKey:      viper.GetString("zap.callerKey"),
			MessageKey:     viper.GetString("zap.messageKey"),
			StacktraceKey:  viper.GetString("zap.stacktraceKey"),
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     formatEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      viper.GetStringSlice("zap.outputPaths"),
		ErrorOutputPaths: viper.GetStringSlice("zap.errorOutputPaths"),
		InitialFields: map[string]interface{}{
			"app": "test",
		},
	}
	//fmt.Println(&cfg.OutputPaths)
	var err error
	logger, err = cfg.Build()
	if err != nil {
		panic("log init fail:" + err.Error())
	}
}
*/
