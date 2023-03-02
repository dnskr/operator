// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Operator
// Copyright (c) 2023 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GemaltoConfigurationResponse gemalto configuration response
//
// swagger:model gemaltoConfigurationResponse
type GemaltoConfigurationResponse struct {

	// keysecure
	// Required: true
	Keysecure *GemaltoConfigurationResponseKeysecure `json:"keysecure"`
}

// Validate validates this gemalto configuration response
func (m *GemaltoConfigurationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKeysecure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GemaltoConfigurationResponse) validateKeysecure(formats strfmt.Registry) error {

	if err := validate.Required("keysecure", "body", m.Keysecure); err != nil {
		return err
	}

	if m.Keysecure != nil {
		if err := m.Keysecure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keysecure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keysecure")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gemalto configuration response based on the context it is used
func (m *GemaltoConfigurationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateKeysecure(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GemaltoConfigurationResponse) contextValidateKeysecure(ctx context.Context, formats strfmt.Registry) error {

	if m.Keysecure != nil {
		if err := m.Keysecure.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keysecure")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keysecure")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GemaltoConfigurationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GemaltoConfigurationResponse) UnmarshalBinary(b []byte) error {
	var res GemaltoConfigurationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GemaltoConfigurationResponseKeysecure gemalto configuration response keysecure
//
// swagger:model GemaltoConfigurationResponseKeysecure
type GemaltoConfigurationResponseKeysecure struct {

	// credentials
	// Required: true
	Credentials *GemaltoConfigurationResponseKeysecureCredentials `json:"credentials"`

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`
}

// Validate validates this gemalto configuration response keysecure
func (m *GemaltoConfigurationResponseKeysecure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GemaltoConfigurationResponseKeysecure) validateCredentials(formats strfmt.Registry) error {

	if err := validate.Required("keysecure"+"."+"credentials", "body", m.Credentials); err != nil {
		return err
	}

	if m.Credentials != nil {
		if err := m.Credentials.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keysecure" + "." + "credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keysecure" + "." + "credentials")
			}
			return err
		}
	}

	return nil
}

func (m *GemaltoConfigurationResponseKeysecure) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("keysecure"+"."+"endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this gemalto configuration response keysecure based on the context it is used
func (m *GemaltoConfigurationResponseKeysecure) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GemaltoConfigurationResponseKeysecure) contextValidateCredentials(ctx context.Context, formats strfmt.Registry) error {

	if m.Credentials != nil {
		if err := m.Credentials.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keysecure" + "." + "credentials")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keysecure" + "." + "credentials")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GemaltoConfigurationResponseKeysecure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GemaltoConfigurationResponseKeysecure) UnmarshalBinary(b []byte) error {
	var res GemaltoConfigurationResponseKeysecure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GemaltoConfigurationResponseKeysecureCredentials gemalto configuration response keysecure credentials
//
// swagger:model GemaltoConfigurationResponseKeysecureCredentials
type GemaltoConfigurationResponseKeysecureCredentials struct {

	// domain
	// Required: true
	Domain *string `json:"domain"`

	// retry
	Retry int64 `json:"retry,omitempty"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this gemalto configuration response keysecure credentials
func (m *GemaltoConfigurationResponseKeysecureCredentials) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GemaltoConfigurationResponseKeysecureCredentials) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("keysecure"+"."+"credentials"+"."+"domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *GemaltoConfigurationResponseKeysecureCredentials) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("keysecure"+"."+"credentials"+"."+"token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this gemalto configuration response keysecure credentials based on context it is used
func (m *GemaltoConfigurationResponseKeysecureCredentials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GemaltoConfigurationResponseKeysecureCredentials) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GemaltoConfigurationResponseKeysecureCredentials) UnmarshalBinary(b []byte) error {
	var res GemaltoConfigurationResponseKeysecureCredentials
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}