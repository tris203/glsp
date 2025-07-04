package protocol

import (
	"encoding/json"
	"strconv"

	"github.com/tliron/glsp"
)

var True bool = true
var False bool = false

type Method = string

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#number

/**
 * Defines an integer number in the range of -2^31 to 2^31 - 1.
 */
type Integer = int32

/**
 * Defines an unsigned integer number in the range of 0 to 2^31 - 1.
 */
type UInteger = uint32

/**
 * Defines a decimal number. Since decimal numbers are very
 * rare in the language server specification we denote the
 * exact range with every decimal using the mathematics
 * interval notation (e.g. [0, 1] denotes all decimals d with
 * 0 <= d <= 1.
 */
type Decimal = float32

type IntegerOrString struct {
	Value any // Integer | string
}

// ([json.Marshaler] interface)
func (i *IntegerOrString) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}

// ([json.Unmarshaler] interface)
func (i *IntegerOrString) UnmarshalJSON(data []byte) error {
	var value Integer
	if err := json.Unmarshal(data, &value); err == nil {
		i.Value = value
		return nil
	} else {
		var value string
		if err := json.Unmarshal(data, &value); err == nil {
			i.Value = value
			return nil
		} else {
			return err
		}
	}
}

type BoolOrString struct {
	Value any // bool | string
}

// ([json.Marshaler] interface)
func (b BoolOrString) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.Value)
}

// ([json.Unmarshaler] interface)
func (b *BoolOrString) UnmarshalJSON(data []byte) error {
	var value bool
	if err := json.Unmarshal(data, &value); err == nil {
		b.Value = value
		return nil
	} else {
		var value string
		if err := json.Unmarshal(data, &value); err == nil {
			b.Value = value
			return nil
		} else {
			return err
		}
	}
}

// ([fmt.Stringer] interface)
func (b BoolOrString) String() string {
	if value, ok := b.Value.(bool); ok {
		return strconv.FormatBool(value)
	} else {
		return b.Value.(string)
	}
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#cancelRequest

const MethodCancelRequest = Method("$/cancelRequest")

type CancelParams struct {
	/**
	 * The request id to cancel.
	 */
	ID IntegerOrString `json:"id"`
}

// https://microsoft.github.io/language-server-protocol/specifications/specification-3-16#progress

const MethodProgress = Method("$/progress")

type RequestFunc[P any, R any] func(context *glsp.Context, params *P) (R, error)
type NotificationFunc[P any] func(context *glsp.Context, params *P) error
type ContextOnlyFunc func(context *glsp.Context) error

type ProgressParams struct {
	/**
	 * The progress token provided by the client or server.
	 */
	Token ProgressToken `json:"token"`

	/**
	 * The progress data.
	 */
	Value any `json:"value"`
}

type ProgressToken = IntegerOrString
