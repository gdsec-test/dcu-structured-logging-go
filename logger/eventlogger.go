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

//EventStruct is a structure for the 'event' json object in the event message
type EventStruct struct {
	Kind      string `json:"kind"`
	Category  string `json:"category"`
	EventType string `json:"type"`
	Outcome   string `json:"outcome"`
	Action    string `json:"action"`
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
func LogEvent(l *zap.Logger, env string, serviceName string, message string, outcome string, action string, sourceIP string, extras []byte) {
	tags := [...]string{"security", "appication"}
	labelsData := LabelsStruct{
		Environment: env,
	}
	eventData := EventStruct{
		Kind:      "event",
		Category:  "iam",
		EventType: "change",
		Outcome:   outcome,
		Action:    action,
	}
	sourceData := SourceStruct{
		IP: sourceIP,
	}
	serviceData := ServiceStruct{
		Name: serviceName,
	}

	jsonLabels, _ := json.Marshal(labelsData)
	jsonEventData, _ := json.Marshal(eventData)
	jsonSourceData, _ := json.Marshal(sourceData)
	jsonServiceData, _ := json.Marshal(serviceData)

	l.Info(message, zap.Any("labels", json.RawMessage(jsonLabels)), zap.Any("tags", tags[:]), zap.Any("event", json.RawMessage(jsonEventData)), zap.Any("source", json.RawMessage(jsonSourceData)), zap.Any("service", json.RawMessage(jsonServiceData)), zap.Any("extra", json.RawMessage(extras)))
}
