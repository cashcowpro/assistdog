package defaults

import (
	"fmt"
	"strconv"
	"time"
)

var supportedTimeLayouts = []string{
	time.RFC822,
	time.RFC3339,
	time.RFC3339Nano,
}

// NilString used to identify "nil" values on the table
var NilString string

// IntPtr variable
var IntPtr *int

// StringPtr variable
var StringPtr *string

// TimePtr variable
var TimePtr *time.Time

// Float32Ptr variable
var Float32Ptr *float32

// ParseString function
func ParseString(raw string) (interface{}, error) {
	return raw, nil
}

// ParseInt function
func ParseInt(raw string) (interface{}, error) {
	return strconv.Atoi(raw)
}

// ParseFloat32 function
func ParseFloat32(raw string) (interface{}, error) {
	ei, err := strconv.ParseFloat(raw, 32)
	return float32(ei), err
}

// ParseTime function
func ParseTime(raw string) (interface{}, error) {
	var fieldTime time.Time
	var err error
	for _, layout := range supportedTimeLayouts {
		fieldTime, err = time.Parse(layout, raw)
		if err != nil {
			continue
		}

		break
	}

	if err != nil {
		return nil, fmt.Errorf("unrecognized time format %v", raw)
	}

	return fieldTime, nil
}

// ParseStringPointer function
func ParseStringPointer(raw string) (interface{}, error) {
	if raw == NilString {
		return StringPtr, nil
	}
	return &raw, nil
}

// ParseIntPointer function
func ParseIntPointer(raw string) (interface{}, error) {
	if raw == NilString {
		return IntPtr, nil
	}

	parsedInt, err := ParseInt(raw)
	if err != nil {
		return nil, err
	}
	time := parsedInt.(int)
	return &time, nil
}

// ParseFloat32Pointer function
func ParseFloat32Pointer(raw string) (interface{}, error) {
	if raw == NilString {
		return Float32Ptr, nil
	}

	parsedFloat32, err := ParseFloat32(raw)
	if err != nil {
		return nil, err
	}
	float32 := parsedFloat32.(float32)
	return &float32, nil
}

// ParseTimePointer function
func ParseTimePointer(raw string) (interface{}, error) {
	if raw == NilString {
		return TimePtr, nil
	}

	parsedTime, err := ParseTime(raw)
	if err != nil {
		return nil, err
	}
	time := parsedTime.(time.Time)
	return &time, err
}
