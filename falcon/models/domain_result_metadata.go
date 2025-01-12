// Code generated by go-swagger; DO NOT EDIT.

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

// DomainResultMetadata domain result metadata
//
// swagger:model domain.ResultMetadata
type DomainResultMetadata struct {

	// execution duration
	// Required: true
	ExecutionDuration *int64 `json:"execution_duration"`

	// execution start
	// Required: true
	// Format: date-time
	ExecutionStart *strfmt.DateTime `json:"execution_start"`

	// report file name
	// Required: true
	ReportFileName *string `json:"report_file_name"`

	// result count
	// Required: true
	ResultCount *int32 `json:"result_count"`

	// result id
	// Required: true
	ResultID *string `json:"result_id"`

	// search window end
	// Required: true
	// Format: date-time
	SearchWindowEnd *strfmt.DateTime `json:"search_window_end"`

	// search window start
	// Required: true
	// Format: date-time
	SearchWindowStart *strfmt.DateTime `json:"search_window_start"`
}

// Validate validates this domain result metadata
func (m *DomainResultMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExecutionDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecutionStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResultID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchWindowEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchWindowStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainResultMetadata) validateExecutionDuration(formats strfmt.Registry) error {

	if err := validate.Required("execution_duration", "body", m.ExecutionDuration); err != nil {
		return err
	}

	return nil
}

func (m *DomainResultMetadata) validateExecutionStart(formats strfmt.Registry) error {

	if err := validate.Required("execution_start", "body", m.ExecutionStart); err != nil {
		return err
	}

	if err := validate.FormatOf("execution_start", "body", "date-time", m.ExecutionStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainResultMetadata) validateReportFileName(formats strfmt.Registry) error {

	if err := validate.Required("report_file_name", "body", m.ReportFileName); err != nil {
		return err
	}

	return nil
}

func (m *DomainResultMetadata) validateResultCount(formats strfmt.Registry) error {

	if err := validate.Required("result_count", "body", m.ResultCount); err != nil {
		return err
	}

	return nil
}

func (m *DomainResultMetadata) validateResultID(formats strfmt.Registry) error {

	if err := validate.Required("result_id", "body", m.ResultID); err != nil {
		return err
	}

	return nil
}

func (m *DomainResultMetadata) validateSearchWindowEnd(formats strfmt.Registry) error {

	if err := validate.Required("search_window_end", "body", m.SearchWindowEnd); err != nil {
		return err
	}

	if err := validate.FormatOf("search_window_end", "body", "date-time", m.SearchWindowEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainResultMetadata) validateSearchWindowStart(formats strfmt.Registry) error {

	if err := validate.Required("search_window_start", "body", m.SearchWindowStart); err != nil {
		return err
	}

	if err := validate.FormatOf("search_window_start", "body", "date-time", m.SearchWindowStart.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain result metadata based on context it is used
func (m *DomainResultMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainResultMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainResultMetadata) UnmarshalBinary(b []byte) error {
	var res DomainResultMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
