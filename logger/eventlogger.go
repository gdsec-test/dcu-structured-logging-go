package logger

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LabelsStruct is a structure for the 'labels' json object in the event message
type LabelsStruct struct {
	Environment string `json:"environment"`
}

// SourceStruct is a structure for the 'source' json object in the event message
type SourceStruct struct {
	IP string `json:"ip"`
}

// ServiceStruct is a structure for the 'service' json object in the event message
type ServiceStruct struct {
	Name string `json:"name"`
}

// NewEventInfoLogger returns a zap logger object with required config
func NewEventInfoLogger() *zap.Logger {
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message",

			TimeKey:    "@timestamp",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}.Build()
	return logger
}

// LogEvent is a helper function thats used to log an event
func LogEvent(l *zap.Logger, env string, serviceName string, message string, sourceIP string, userData interface{}, eventData interface{}) {
	tags := [...]string{"security", "application"}
	labelsData := LabelsStruct{
		Environment: env,
	}
	sourceData := SourceStruct{
		IP: sourceIP,
	}
	serviceData := ServiceStruct{
		Name: serviceName,
	}

	jsonLabels, _ := json.Marshal(labelsData)
	jsonSourceData, _ := json.Marshal(sourceData)
	jsonServiceData, _ := json.Marshal(serviceData)

	l.Info(message, zap.Any("labels", json.RawMessage(jsonLabels)), zap.Any("tags", tags[:]), zap.Any("event", eventData), zap.Any("source", json.RawMessage(jsonSourceData)), zap.Any("service", json.RawMessage(jsonServiceData)), zap.Any("user", userData))
}
