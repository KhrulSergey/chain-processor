package logger

import "github.com/stretchr/testify/mock"

// MockLogger mocks the logger.Logger interface
type MockLogger struct {
	mock.Mock
}

func (ml *MockLogger) Infof(format string, args ...interface{}) {
	ml.Called(format, args)
}

func (ml *MockLogger) Error(args ...interface{}) {
	ml.Called(args)
}

func (ml *MockLogger) Warning(args ...interface{}) {
	ml.Called(args)
}

func (ml *MockLogger) Warningf(format string, args ...interface{}) {
	ml.Called(format, args)
}

func (ml *MockLogger) Panic(args ...interface{}) {
	ml.Called(args)
}

func (ml *MockLogger) Panicf(format string, args ...interface{}) {
	ml.Called(format, args)
}

func (ml *MockLogger) Fatal(args ...interface{}) {
	ml.Called(args)
}

func (ml *MockLogger) Fatalf(template string, args ...interface{}) {
	ml.Called(template, args)
}

func (ml *MockLogger) Flush() error {
	// Simulate a mock call for Flush, ensuring the mock library tracks this function call.
	args := ml.Called()
	// Return the appropriate error (or nil) as per the mock's setup.
	return args.Error(0)
}

func (ml *MockLogger) Debug(args ...interface{}) {
	ml.Called(args)
}

func (ml *MockLogger) Debugf(format string, args ...interface{}) {
	ml.Called(format, args)
}

func (ml *MockLogger) Info(args ...interface{}) {
	ml.Called(args)
}

func (ml *MockLogger) Errorf(template string, args ...interface{}) {
	ml.Called(template, args)
}
