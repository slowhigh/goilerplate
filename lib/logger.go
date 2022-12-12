package lib

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	gorm_logger "gorm.io/gorm/logger"
)

type Logger struct {
	*zap.SugaredLogger
}

type GinLogger struct {
	*Logger
}

type FxLogger struct {
	*Logger
}

type GormLogger struct {
	*Logger
	gorm_logger.Config
}

var (
	globalLogger *Logger
	zapLogger    *zap.Logger
)

// GetLogger get the logger
func GetLogger() Logger {
	if globalLogger == nil {
		logger := newLogger(NewEnv())
		globalLogger = &logger
	}

	return *globalLogger
}

// GetGinLogger get the gin logger
func (l Logger) GetGinLogger() GinLogger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

// GetFxLogger get the fx logger
func (l *Logger) GetFxLogger() fxevent.Logger {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)

	return &FxLogger{
		Logger: newSugaredLogger(logger),
	}
}

// GetGormLogger get the gorm logger
func (l Logger) GetGormLogger() *GormLogger {
	logger := zapLogger.WithOptions(
		zap.AddCaller(),
		zap.AddCallerSkip(3),
	)

	return &GormLogger{
		Logger: newSugaredLogger(logger),
		Config: gorm_logger.Config{
			LogLevel: gorm_logger.Info,
		},
	}
}

// LogEvent log event for fx logger
func (l *FxLogger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Logger.Debug("OnStart hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Logger.Debug("OnStart hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Logger.Debug("OnStart hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Logger.Debug("OnStop hook executing: ",
			zap.String("callee", e.FunctionName),
			zap.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Logger.Debug("OnStop hook failed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.Error(e.Err),
			)
		} else {
			l.Logger.Debug("OnStop hook executed: ",
				zap.String("callee", e.FunctionName),
				zap.String("caller", e.CallerName),
				zap.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		l.Logger.Debug("supplied: ", zap.String("type", e.TypeName), zap.Error(e.Err))
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("provided: ", e.ConstructorName, " => ", rtype)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Logger.Debug("decorated: ",
				zap.String("decorator", e.DecoratorName),
				zap.String("type", rtype),
			)
		}
	case *fxevent.Invoking:
		l.Logger.Debug("invoking: ", e.FunctionName)
	case *fxevent.Started:
		if e.Err == nil {
			l.Logger.Debug("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err == nil {
			l.Logger.Debug("initialized: custom fxevent.Logger -> ", e.ConstructorName)
		}
	}
}

// Write interface implementation for gin-framework
func (gl GinLogger) Write(p []byte) (n int, err error) {
	gl.Info(string(p))
	return len(p), nil
}

// Printf prits go-fx logs
func (l FxLogger) Printf(str string, args ...interface{}) {
	if len(args) > 0 {
		l.Debugf(str, args)
	}
	l.Debug(str)
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

func newLogger(env Env) Logger {
	config := zap.NewDevelopmentConfig()
	logOutput := env.LogOutput

	if env.Environment == "development" {
		fmt.Println("encode level")
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	}

	if env.Environment == "production" && logOutput != "" {
		config.OutputPaths = []string{logOutput}
	}

	logLevel := env.LogLevel
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zap.PanicLevel
	}
	config.Level.SetLevel(level)

	zapLogger, _ = config.Build()
	logger := newSugaredLogger(zapLogger)

	return *logger
}

// LogMode log mode
func (gl *GormLogger) LogMode(level gorm_logger.LogLevel) gorm_logger.Interface {
	newLogger := *gl
	newLogger.LogLevel = level
	return &newLogger
}

// Info print info messages
func (gl GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= gorm_logger.Info {
		gl.Debugf(msg, data...)
	}
}

// Warn print warn messages
func (gl GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= gorm_logger.Warn {
		gl.Warnf(msg, data...)
	}
}

// Error print error messages
func (gl GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if gl.LogLevel >= gorm_logger.Error {
		gl.Errorf(msg, data...)
	}
}

// Trace print sql message
func (gl GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if gl.LogLevel < gorm_logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	if gl.LogLevel >= gorm_logger.Info {
		gl.Debugf("[%d ms, %d rows] sql -> %s", elapsed.Milliseconds(), rows, sql)
	} else if gl.LogLevel >= gorm_logger.Warn {
		gl.SugaredLogger.Warnf("[%d ms, %d rows] sql -> %s", elapsed.Milliseconds(), rows, sql)
	} else if gl.LogLevel >= gorm_logger.Error {
		gl.SugaredLogger.Errorf("[%d ms, %d rows] sql -> %s", elapsed.Milliseconds(), rows, sql)
	}
}
