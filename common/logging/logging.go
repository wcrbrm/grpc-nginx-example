package logging

import (
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Level converts both numeric and string level format into enum
func Level(s string) log.Level {
	switch s {
	case "1", "error":
		return log.ErrorLevel
	case "2", "warn":
		return log.WarnLevel
	case "3", "info":
		return log.InfoLevel
	case "4", "debug":
		return log.DebugLevel
	default:
		return log.FatalLevel
	}
}

// DefaultFormatter for JSON
func DefaultFormatter() *log.JSONFormatter {
	return &log.JSONFormatter{
		FieldMap: log.FieldMap{
			log.FieldKeyTime:  "@timestamp",
			log.FieldKeyLevel: "@level",
			log.FieldKeyMsg:   "@message",
			log.FieldKeyFunc:  "@caller",
		},
	}
}

// WithFn adds map of key values for logger
func WithFn(fields ...log.Fields) log.Fields {
	if len(fields) > 0 && fields[0] != nil {
		result := copyFields(fields[0])
		result["fn"] = getCallerName()
		return result
	}
	return log.Fields{
		"fn": getCallerName(),
	}
}

// WithTraceID adds trace ID to the fields included in logger
func WithTraceID(traceID string, fields ...log.Fields) log.Fields {
	if len(traceID) == 0 {
		if len(fields) > 0 {
			return fields[0]
		}
		return log.Fields{}
	} else if len(fields) > 0 && fields[0] != nil {
		result := copyFields(fields[0])
		result["trace_id"] = traceID
		return result
	}
	return log.Fields{
		"trace_id": traceID,
	}
}

// WithMore adds combines 2 maps of fields for logging
func WithMore(fields log.Fields, add log.Fields) log.Fields {
	fields = copyFields(fields)
	for k, v := range add {
		fields[k] = v
	}
	return fields
}

func copyFields(fields log.Fields) log.Fields {
	ff := make(log.Fields, len(fields))
	for k, v := range fields {
		ff[k] = v
	}
	return ff
}

// FnName returns name of the caller function
func FnName() string {
	return getCallerName()
}

func getCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	fullName := runtime.FuncForPC(pc).Name()
	parts := strings.Split(fullName, "/")
	nameParts := strings.Split(parts[len(parts)-1], ".")
	return nameParts[len(nameParts)-1]
}
