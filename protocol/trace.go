package protocol

import (
	"fmt"
	"sync"

	"github.com/tliron/glsp"
)

var traceValue TraceValue = Off
var traceValueLock sync.Mutex

func GetTraceValue() TraceValue {
	traceValueLock.Lock()
	defer traceValueLock.Unlock()
	return traceValue
}

func SetTraceValue(value TraceValue) {
	traceValueLock.Lock()
	defer traceValueLock.Unlock()

	// The spec clearly says "message", but some implementations use "messages" instead
	if value == "messages" {
		value = Messages
	}

	traceValue = value
}

func HasTraceLevel(value TraceValue) bool {
	value_ := GetTraceValue()
	switch value_ {
	case Off:
		return false

	case Messages:
		return value == Messages

	case Verbose:
		return true

	default:
		panic(fmt.Sprintf("unsupported trace level: %s", value_))
	}
}

func HasTraceMessageType(type_ MessageType) bool {
	switch type_ {
	case Error, Warning, Info:
		return HasTraceLevel(Messages)

	case Log:
		return HasTraceLevel(Verbose)

	default:
		panic(fmt.Sprintf("unsupported message type: %d", type_))
	}
}

func Trace(context *glsp.Context, type_ MessageType, message string) error {
	if HasTraceMessageType(type_) {
		go context.Notify(ServerWindowLogMessage, &LogMessageParams{
			Type:    type_,
			Message: message,
		})
	}
	return nil
}
