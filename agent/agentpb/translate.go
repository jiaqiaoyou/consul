package agentpb

import (
	"reflect"
	"time"

	"github.com/gogo/protobuf/types"
)

var (
	tsType   = reflect.TypeOf((*types.Timestamp)(nil))
	timeType = reflect.TypeOf((*time.Time)(nil)).Elem()
)

// HookPBTimestampToTime is a mapstructure decode hook to translate a protobuf timestamp
// to a time.Time value
func HookPBTimestampToTime(from, to reflect.Type, data interface{}) (interface{}, error) {
	if to == timeType && from == tsType {
		ts := data.(*types.Timestamp)
		return time.Unix(ts.Seconds, int64(ts.Nanos)), nil
	}

	return data, nil
}

// HookTimeToPBtimestamp is a mapstructure decode hook to translate a time.Time value to
// a protobuf Timestamp value.
func HookTimeToPBTimestamp(from, to reflect.Type, data interface{}) (interface{}, error) {
	if from == timeType && to == tsType {
		ts := data.(time.Time)
		return &types.Timestamp{Seconds: ts.Unix(), Nanos: int32(ts.UnixNano())}, nil
	}

	return data, nil
}
