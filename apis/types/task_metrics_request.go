// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TaskMetricsRequest task metrics request
// swagger:model TaskMetricsRequest
type TaskMetricsRequest struct {

	// IP address which peer client carries
	IP string `json:"IP,omitempty"`

	// when registering, dfget will setup one uploader process.
	// This one acts as a server for peer pulling tasks.
	// This port is which this server listens on.
	//
	BacksourceReason string `json:"backsourceReason,omitempty"`

	// CID means the client ID. It maps to the specific dfget process.
	// When user wishes to download an image/file, user would start a dfget process to do this.
	// This dfget is treated a client and carries a client ID.
	// Thus, multiple dfget processes on the same peer have different CIDs.
	//
	CID string `json:"cID,omitempty"`

	// This attribute represents where the dfget requests come from. Dfget will pass
	// this field to supernode and supernode can do some checking and filtering via
	// black/white list mechanism to guarantee security, or some other purposes like debugging.
	//
	// Min Length: 1
	CallSystem string `json:"callSystem,omitempty"`

	// Duration for dfget task.
	//
	Duration float64 `json:"duration,omitempty"`

	// The length of the file dfget requests to download in bytes.
	FileLength int64 `json:"fileLength,omitempty"`

	// when registering, dfget will setup one uploader process.
	// This one acts as a server for peer pulling tasks.
	// This port is which this server listens on.
	//
	// Maximum: 65000
	// Minimum: 15000
	Port int32 `json:"port,omitempty"`

	// whether the download task success or not
	Success bool `json:"success,omitempty"`

	// IP address which peer client carries
	TaskID string `json:"taskId,omitempty"`
}

// Validate validates this task metrics request
func (m *TaskMetricsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallSystem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TaskMetricsRequest) validateCallSystem(formats strfmt.Registry) error {

	if swag.IsZero(m.CallSystem) { // not required
		return nil
	}

	if err := validate.MinLength("callSystem", "body", string(m.CallSystem), 1); err != nil {
		return err
	}

	return nil
}

func (m *TaskMetricsRequest) validatePort(formats strfmt.Registry) error {

	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if err := validate.MinimumInt("port", "body", int64(m.Port), 15000, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port", "body", int64(m.Port), 65000, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TaskMetricsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskMetricsRequest) UnmarshalBinary(b []byte) error {
	var res TaskMetricsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}