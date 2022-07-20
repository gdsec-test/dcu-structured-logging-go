package logger

import (
	"encoding/json"
	//"fmt"

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

// UserStruct is a structure for the 'user' json object in the event message
type UserStruct struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
	ID     string `json:"id"`
}

// ServiceStruct is a structure for the 'service' json object in the event message
type ServiceStruct struct {
	Name string `json:"name"`
}

// EventInfo is the data structure that holds the required fields for event logging
type EventInfo struct {
	Tags    []string      `json:"tags"`
	Message string        `json:"message"`
	Event   EventStruct   `json:"event"`
	User    UserStruct    `json:"user"`
	Service ServiceStruct `json:"service"`
}

// NewEventInfoLogger returns a zap logger object with required config
func NewEventInfoLogger() *zap.Logger {
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"stdout"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:    "@timestamp",
			EncodeTime: zapcore.ISO8601TimeEncoder,
		},
	}.Build()
	return logger
}

// LogEvent is a helper function thats used to log an event
func LogEvent(l *zap.Logger, env string, serviceName string, message string, outcome string, action string, username string, domainName string, userID string) {
	tags := [...]string{"security", "appication"}

	eventInfo := EventInfo{
		Tags:    tags[:],
		Message: message,
		Event: EventStruct{
			Kind:      "event",
			Category:  "iam",
			EventType: "change",
			Outcome:   outcome,
			Action:    action,
		},
		User: UserStruct{
			Name:   username,
			Domain: domainName,
			ID:     userID,
		},
		Service: ServiceStruct{
			Name: serviceName,
		},
	}

	labelsData := LabelsStruct{
		Environment: env,
	}

	jsonEventInfo, _ := json.Marshal(eventInfo)
	jsonLabels, _ := json.Marshal(labelsData)

	l.Info("", zap.Any("labels", json.RawMessage(jsonLabels)), zap.Any("event-info", json.RawMessage(jsonEventInfo)))
}
