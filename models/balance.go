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

// Balance balance
//
// swagger:model balance
type Balance struct {

	// algorithm
	// Required: true
	// Enum: [roundrobin static-rr leastconn first source uri url_param hdr random rdp-cookie]
	Algorithm *string `json:"algorithm"`

	// hdr name
	HdrName string `json:"hdr_name,omitempty"`

	// hdr use domain only
	HdrUseDomainOnly bool `json:"hdr_use_domain_only,omitempty"`

	// random draws
	RandomDraws int64 `json:"random_draws,omitempty"`

	// rdp cookie name
	// Pattern: ^[^\s]+$
	RdpCookieName string `json:"rdp_cookie_name,omitempty"`

	// uri depth
	// Pattern: ^[^\d+$]
	URIDepth int64 `json:"uri_depth,omitempty"`

	// uri len
	// Pattern: ^[^\d+$]
	URILen int64 `json:"uri_len,omitempty"`

	// uri whole
	URIWhole bool `json:"uri_whole,omitempty"`

	// url param
	// Pattern: ^[^\s]+$
	URLParam string `json:"url_param,omitempty"`

	// url param check post
	URLParamCheckPost int64 `json:"url_param_check_post,omitempty"`

	// url param max wait
	// Pattern: ^[^\d+$]
	URLParamMaxWait int64 `json:"url_param_max_wait,omitempty"`
}

// Validate validates this balance
func (m *Balance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRdpCookieName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURIDepth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURILen(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLParam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURLParamMaxWait(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var balanceTypeAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["roundrobin","static-rr","leastconn","first","source","uri","url_param","hdr","random","rdp-cookie"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		balanceTypeAlgorithmPropEnum = append(balanceTypeAlgorithmPropEnum, v)
	}
}

const (

	// BalanceAlgorithmRoundrobin captures enum value "roundrobin"
	BalanceAlgorithmRoundrobin string = "roundrobin"

	// BalanceAlgorithmStaticRr captures enum value "static-rr"
	BalanceAlgorithmStaticRr string = "static-rr"

	// BalanceAlgorithmLeastconn captures enum value "leastconn"
	BalanceAlgorithmLeastconn string = "leastconn"

	// BalanceAlgorithmFirst captures enum value "first"
	BalanceAlgorithmFirst string = "first"

	// BalanceAlgorithmSource captures enum value "source"
	BalanceAlgorithmSource string = "source"

	// BalanceAlgorithmURI captures enum value "uri"
	BalanceAlgorithmURI string = "uri"

	// BalanceAlgorithmURLParam captures enum value "url_param"
	BalanceAlgorithmURLParam string = "url_param"

	// BalanceAlgorithmHdr captures enum value "hdr"
	BalanceAlgorithmHdr string = "hdr"

	// BalanceAlgorithmRandom captures enum value "random"
	BalanceAlgorithmRandom string = "random"

	// BalanceAlgorithmRdpCookie captures enum value "rdp-cookie"
	BalanceAlgorithmRdpCookie string = "rdp-cookie"
)

// prop value enum
func (m *Balance) validateAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, balanceTypeAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Balance) validateAlgorithm(formats strfmt.Registry) error {

	if err := validate.Required("algorithm", "body", m.Algorithm); err != nil {
		return err
	}

	// value enum
	if err := m.validateAlgorithmEnum("algorithm", "body", *m.Algorithm); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateRdpCookieName(formats strfmt.Registry) error {

	if swag.IsZero(m.RdpCookieName) { // not required
		return nil
	}

	if err := validate.Pattern("rdp_cookie_name", "body", string(m.RdpCookieName), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateURIDepth(formats strfmt.Registry) error {

	if swag.IsZero(m.URIDepth) { // not required
		return nil
	}

	if err := validate.Pattern("uri_depth", "body", string(m.URIDepth), `^[^\d+$]`); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateURILen(formats strfmt.Registry) error {

	if swag.IsZero(m.URILen) { // not required
		return nil
	}

	if err := validate.Pattern("uri_len", "body", string(m.URILen), `^[^\d+$]`); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateURLParam(formats strfmt.Registry) error {

	if swag.IsZero(m.URLParam) { // not required
		return nil
	}

	if err := validate.Pattern("url_param", "body", string(m.URLParam), `^[^\s]+$`); err != nil {
		return err
	}

	return nil
}

func (m *Balance) validateURLParamMaxWait(formats strfmt.Registry) error {

	if swag.IsZero(m.URLParamMaxWait) { // not required
		return nil
	}

	if err := validate.Pattern("url_param_max_wait", "body", string(m.URLParamMaxWait), `^[^\d+$]`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Balance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Balance) UnmarshalBinary(b []byte) error {
	var res Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
