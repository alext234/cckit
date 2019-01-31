// Package convert for transforming  between json serialized  []byte and go structs
package convert

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/pkg/errors"
)

var (
	// ErrUnableToConvertNilToStruct - nil cannot be converted to struct
	ErrUnableToConvertNilToStruct = errors.New(`unable to convert nil to [struct,array,slice,ptr]`)
	// ErrUnableToConvertValueToStruct - value  cannot be converted to struct
	ErrUnableToConvertValueToStruct = errors.New(`unable to convert value to [struct,array,slice,ptr]`)
)

const TypeInt = 1
const TypeString = ``
const TypeBool = true

// FromByter interface supports FromBytes func for converting to structure
type (
	FromByter interface {
		FromBytes([]byte) (interface{}, error)
	}

	// ToByter interface supports ToBytes func, marshalling to []byte (json.Marshall)
	ToByter interface {
		ToBytes() ([]byte, error)
	}
)

// TimestampToTime converts timestamp to time.Time
func TimestampToTime(ts *timestamp.Timestamp) time.Time {
	return time.Unix(ts.GetSeconds(), int64(ts.GetNanos()))
}
