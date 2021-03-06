package j_stack_template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// JStackTemplateCloneReader is a Reader for the JStackTemplateClone structure.
type JStackTemplateCloneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JStackTemplateCloneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewJStackTemplateCloneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewJStackTemplateCloneOK creates a JStackTemplateCloneOK with default headers values
func NewJStackTemplateCloneOK() *JStackTemplateCloneOK {
	return &JStackTemplateCloneOK{}
}

/*JStackTemplateCloneOK handles this case with default header values.

OK
*/
type JStackTemplateCloneOK struct {
	Payload JStackTemplateCloneOKBody
}

func (o *JStackTemplateCloneOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JStackTemplate.clone/{id}][%d] jStackTemplateCloneOK  %+v", 200, o.Payload)
}

func (o *JStackTemplateCloneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*JStackTemplateCloneOKBody j stack template clone o k body
swagger:model JStackTemplateCloneOKBody
*/
type JStackTemplateCloneOKBody struct {
	models.JStackTemplate

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *JStackTemplateCloneOKBody) UnmarshalJSON(raw []byte) error {

	var jStackTemplateCloneOKBodyAO0 models.JStackTemplate
	if err := swag.ReadJSON(raw, &jStackTemplateCloneOKBodyAO0); err != nil {
		return err
	}
	o.JStackTemplate = jStackTemplateCloneOKBodyAO0

	var jStackTemplateCloneOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &jStackTemplateCloneOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = jStackTemplateCloneOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o JStackTemplateCloneOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	jStackTemplateCloneOKBodyAO0, err := swag.WriteJSON(o.JStackTemplate)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jStackTemplateCloneOKBodyAO0)

	jStackTemplateCloneOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, jStackTemplateCloneOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this j stack template clone o k body
func (o *JStackTemplateCloneOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JStackTemplate.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
