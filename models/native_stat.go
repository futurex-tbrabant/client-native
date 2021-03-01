// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NativeStat Stats
//
// Current stats for one object.
//
// swagger:model native_stat
type NativeStat struct {

	// backend name
	BackendName string `json:"backend_name,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// stats
	Stats *NativeStatStats `json:"stats,omitempty"`

	// type
	// Enum: [backend server frontend]
	Type string `json:"type,omitempty"`
}

// Validate validates this native stat
func (m *NativeStat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStats(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NativeStat) validateStats(formats strfmt.Registry) error {

	if swag.IsZero(m.Stats) { // not required
		return nil
	}

	if m.Stats != nil {
		if err := m.Stats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stats")
			}
			return err
		}
	}

	return nil
}

var nativeStatTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["backend","server","frontend"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatTypeTypePropEnum = append(nativeStatTypeTypePropEnum, v)
	}
}

const (

	// NativeStatTypeBackend captures enum value "backend"
	NativeStatTypeBackend string = "backend"

	// NativeStatTypeServer captures enum value "server"
	NativeStatTypeServer string = "server"

	// NativeStatTypeFrontend captures enum value "frontend"
	NativeStatTypeFrontend string = "frontend"
)

// prop value enum
func (m *NativeStat) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nativeStatTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NativeStat) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NativeStat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NativeStat) UnmarshalBinary(b []byte) error {
	var res NativeStat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
