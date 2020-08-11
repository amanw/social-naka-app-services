// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostResponse post response
//
// swagger:model postResponse
type PostResponse struct {

	// post
	Post *Post `json:"post,omitempty"`

	// posts
	Posts []*Post `json:"posts"`

	// response
	Response *Response `json:"response,omitempty"`
}

// Validate validates this post response
func (m *PostResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostResponse) validatePost(formats strfmt.Registry) error {

	if swag.IsZero(m.Post) { // not required
		return nil
	}

	if m.Post != nil {
		if err := m.Post.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post")
			}
			return err
		}
	}

	return nil
}

func (m *PostResponse) validatePosts(formats strfmt.Registry) error {

	if swag.IsZero(m.Posts) { // not required
		return nil
	}

	for i := 0; i < len(m.Posts); i++ {
		if swag.IsZero(m.Posts[i]) { // not required
			continue
		}

		if m.Posts[i] != nil {
			if err := m.Posts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("posts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostResponse) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostResponse) UnmarshalBinary(b []byte) error {
	var res PostResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
