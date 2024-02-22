// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DataUser data user
//
// swagger:model data.User
type DataUser struct {

	// age
	// Required: true
	Age *int64 `json:"age"`

	// email
	// Required: true
	Email *string `json:"email"`

	// id
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	// Min Length: 3
	Name *string `json:"name"`

	// password
	// Required: true
	// Min Length: 7
	Password *string `json:"password"`

	// phone
	// Required: true
	// Minimum: 10
	Phone *int64 `json:"phone"`

	// role
	// Required: true
	// Enum: ["admin","manager","user"]
	Role *string `json:"role"`
}

// Validate validates this data user
func (m *DataUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataUser) validateAge(formats strfmt.Registry) error {

	if err := validate.Required("age", "body", m.Age); err != nil {
		return err
	}

	return nil
}

func (m *DataUser) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *DataUser) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	return nil
}

func (m *DataUser) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", *m.Password, 7); err != nil {
		return err
	}

	return nil
}

func (m *DataUser) validatePhone(formats strfmt.Registry) error {

	if err := validate.Required("phone", "body", m.Phone); err != nil {
		return err
	}

	if err := validate.MinimumInt("phone", "body", *m.Phone, 10, false); err != nil {
		return err
	}

	return nil
}

var dataUserTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["admin","manager","user"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dataUserTypeRolePropEnum = append(dataUserTypeRolePropEnum, v)
	}
}

const (

	// DataUserRoleAdmin captures enum value "admin"
	DataUserRoleAdmin string = "admin"

	// DataUserRoleManager captures enum value "manager"
	DataUserRoleManager string = "manager"

	// DataUserRoleUser captures enum value "user"
	DataUserRoleUser string = "user"
)

// prop value enum
func (m *DataUser) validateRoleEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dataUserTypeRolePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DataUser) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this data user based on context it is used
func (m *DataUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataUser) UnmarshalBinary(b []byte) error {
	var res DataUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}