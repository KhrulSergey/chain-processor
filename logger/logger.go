package logger

import "go.uber.org/zap"

type Logger interface {
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Warning(args ...interface{})
	Warningf(template string, args ...interface{})
	Panic(args ...interface{})
	Panicf(tempplate string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Flush() error
}

type logger struct {
	l *zap.SugaredLogger
}

func NewRelease() (Logger, error) {
	instance, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	return &logger{
		l: instance.Sugar(),
	}, nil
}

func NewDebug() (Logger, error) {
	instance, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &logger{
		l: instance.Sugar(),
	}, nil
}

func (i *logger) Debug(args ...interface{}) {
	i.l.Debug(args)
}

func (i *logger) Debugf(template string, args ...interface{}) {
	i.l.Debugf(template, args...)
}

func (i *logger) Info(args ...interface{}) {
	i.l.Info(args)
}

func (i *logger) Infof(template string, args ...interface{}) {
	i.l.Infof(template, args...)
}

func (i *logger) Error(args ...interface{}) {
	i.l.Error(args)
}

func (i *logger) Errorf(template string, args ...interface{}) {
	i.l.Errorf(template, args...)
}

func (i *logger) Warning(args ...interface{}) {
	i.l.Warn(args)
}

func (i *logger) Warningf(template string, args ...interface{}) {
	i.l.Warnf(template, args...)
}

func (i *logger) Panic(args ...interface{}) {
	i.l.Panic(args)
}

func (i *logger) Panicf(template string, args ...interface{}) {
	i.l.Panicf(template, args...)
}

func (i *logger) Fatal(args ...interface{}) {
	i.l.Fatal(args)
}

func (i *logger) Fatalf(template string, args ...interface{}) {
	i.l.Fatalf(template, args...)
}

func (i *logger) Flush() error {
	return i.l.Sync()
}
