// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package csvtoprotoparse contains runtime functionality needed by code
// generated by the go/csv-to-proto tool.
//
// These functions are not intended to be used outside of generated code "unless
// you know what you're doing."
package csvtoprotoparse

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes"
	ts "google.golang.org/protobuf/types/known/timestamppb"
)

// ParseFloat returns a float from a CSV field.
func ParseFloat(rawValue string) (float32, error) {
	v, err := strconv.ParseFloat(rawValue, 32)
	return float32(v), err
}

// ParseDouble returns a double from a CSV field.
func ParseDouble(rawValue string) (float64, error) {
	v, err := strconv.ParseFloat(rawValue, 64)
	return v, err
}

// ParseInt32 returns an int32 parsed from a CSV field.
func ParseInt32(rawValue string) (int32, error) {
	v, err := strconv.ParseInt(rawValue, 10, 32)
	return int32(v), err
}

// ParseInt64 returns an int64 from a CSV field.
func ParseInt64(rawValue string) (int64, error) {
	v, err := strconv.ParseInt(rawValue, 10, 64)
	return int64(v), err
}

// ParseString parses a string value from a CSV field.
//
// This function has a strange signature for the convenience of the generated
// code. It always returns its first argument and never returns an error.
func ParseString(rawValue string) (string, error) {
	return rawValue, nil
}

// ParseTimestamp returns a proto version of a timestamp using a given layout.
func ParseTimestamp(rawValue, layout string) (*ts.Timestamp, error) {
	t, err := time.Parse(rawValue, layout)
	if err != nil {
		return nil, err
	}
	tt, err := ptypes.TimestampProto(t)
	if err != nil {
		return nil, err
	}
	return tt, nil
}

// TimeToTimestamp returns a Timestamp proto from a time value that may be nil.
func TimeToTimestamp(t time.Time) (*ts.Timestamp, error) {
	return ptypes.TimestampProto(t)
}

// ReaderOption is used to specify a custom argument to csvtoproto readers at construction time.
type ReaderOption interface{}

// MustLoadLocation returns a time.Location or panics.
func MustLoadLocation(name string) *time.Location {
	tz, err := time.LoadLocation(name)
	if err != nil {
		panic(fmt.Errorf("error parsing timezone for lastModifiedTime: %w", err))
	}
	return tz
}
