package logger

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// EventLogger is the data structure that holds the required fields for
// event logging
/*// A single event show below.
{
    "@timestamp": "2016-05-23T08:05:34.853Z",
    "labels": {"environment": "prod"},
    "tags": ["security", "application"],
    "message": "User asomebody was created in domainfind application.",
    "event": {
      "kind": "event",
      "category": "iam",
      "type": ["creation", "creation"],
      "outcome":"success",
      "action": "user_create"
    },
    "user": {
      "name": "asomebody",
      "domain": "thenewdomain.com",
      "id": "12345678910"
    },
    "service": {
      "name": "domainfind"
    }
}*/
type EventInfo struct {
	timestamp		time				`json:"@timestamp"`
	labels struct{
		environment	string				`json:"environment"`
	} `json:"labels"`
	tags  			[]string			`json:"tags"`
	message			string				`json:"message"`
	event struct{
		kind		string				`json:"kind"`
		category	string				`json:"category"`
		eventType	[]string			`json:"type"`
		outcome		string				`json:"outcome"`
		action		string				`json:"action"`
	} `json:"event"`
	user struct{
		name		string				`json:"name"`
		domain		string				`json:"domain"`
		id			string				`json:"id"`
	} `json:"user"`   
	service struct{
		name		string				`json:"name"`	
	}`json:"service"`,
}

// NewEventInfoLogger returns a zap logger object with required config
func NewEventInfoLogger() *zap.Logger {
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	return logger
}

// LogEvent is a helper function thats used to log an event
func LogEvent(l *zap.Logger, env string, serviceName string, message string, outcome string, action string, username string, domainName string, userId string) {
	
	eventInfo := EventInfo{
		timestamp:    			time.Now(),
		labels:environment:     env,
		tags:  					["security", "application"],
		message:				message,
		event:kind:				"event",
		event:category:			"LKM TODO",
		event:eventType:		"LKM TODO: figure out what to put and if this can be an array or just string",
		event:outcome:			outcome,
		event.action:			action,
		user:name:				username,
		user:domain:			domainName,
		user:id:				userId,
		service:name:			serviceName
	}
	jsonEventInfo, _ := json.Marshal(eventInfo)
	// l.Info("", zap.Any("http-info", json.RawMessage(jsonHTTPInfo)))
	l.Info(string(jsonEventInfo))
}
