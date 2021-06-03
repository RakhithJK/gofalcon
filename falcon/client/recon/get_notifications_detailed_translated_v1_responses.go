// Code generated by go-swagger; DO NOT EDIT.

package recon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// GetNotificationsDetailedTranslatedV1Reader is a Reader for the GetNotificationsDetailedTranslatedV1 structure.
type GetNotificationsDetailedTranslatedV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNotificationsDetailedTranslatedV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNotificationsDetailedTranslatedV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNotificationsDetailedTranslatedV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNotificationsDetailedTranslatedV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNotificationsDetailedTranslatedV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetNotificationsDetailedTranslatedV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNotificationsDetailedTranslatedV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetNotificationsDetailedTranslatedV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNotificationsDetailedTranslatedV1OK creates a GetNotificationsDetailedTranslatedV1OK with default headers values
func NewGetNotificationsDetailedTranslatedV1OK() *GetNotificationsDetailedTranslatedV1OK {
	return &GetNotificationsDetailedTranslatedV1OK{}
}

/* GetNotificationsDetailedTranslatedV1OK describes a response with status code 200, with default header values.

OK
*/
type GetNotificationsDetailedTranslatedV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainNotificationDetailsResponseV1
}

func (o *GetNotificationsDetailedTranslatedV1OK) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed-translated/v1][%d] getNotificationsDetailedTranslatedV1OK  %+v", 200, o.Payload)
}
func (o *GetNotificationsDetailedTranslatedV1OK) GetPayload() *models.DomainNotificationDetailsResponseV1 {
	return o.Payload
}

func (o *GetNotificationsDetailedTranslatedV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainNotificationDetailsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedTranslatedV1BadRequest creates a GetNotificationsDetailedTranslatedV1BadRequest with default headers values
func NewGetNotificationsDetailedTranslatedV1BadRequest() *GetNotificationsDetailedTranslatedV1BadRequest {
	return &GetNotificationsDetailedTranslatedV1BadRequest{}
}

/* GetNotificationsDetailedTranslatedV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetNotificationsDetailedTranslatedV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

func (o *GetNotificationsDetailedTranslatedV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed-translated/v1][%d] getNotificationsDetailedTranslatedV1BadRequest  %+v", 400, o.Payload)
}
func (o *GetNotificationsDetailedTranslatedV1BadRequest) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedTranslatedV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedTranslatedV1Unauthorized creates a GetNotificationsDetailedTranslatedV1Unauthorized with default headers values
func NewGetNotificationsDetailedTranslatedV1Unauthorized() *GetNotificationsDetailedTranslatedV1Unauthorized {
	return &GetNotificationsDetailedTranslatedV1Unauthorized{}
}

/* GetNotificationsDetailedTranslatedV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetNotificationsDetailedTranslatedV1Unauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

func (o *GetNotificationsDetailedTranslatedV1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed-translated/v1][%d] getNotificationsDetailedTranslatedV1Unauthorized  %+v", 401, o.Payload)
}
func (o *GetNotificationsDetailedTranslatedV1Unauthorized) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedTranslatedV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedTranslatedV1Forbidden creates a GetNotificationsDetailedTranslatedV1Forbidden with default headers values
func NewGetNotificationsDetailedTranslatedV1Forbidden() *GetNotificationsDetailedTranslatedV1Forbidden {
	return &GetNotificationsDetailedTranslatedV1Forbidden{}
}

/* GetNotificationsDetailedTranslatedV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetNotificationsDetailedTranslatedV1Forbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

func (o *GetNotificationsDetailedTranslatedV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed-translated/v1][%d] getNotificationsDetailedTranslatedV1Forbidden  %+v", 403, o.Payload)
}
func (o *GetNotificationsDetailedTranslatedV1Forbidden) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedTranslatedV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedTranslatedV1TooManyRequests creates a GetNotificationsDetailedTranslatedV1TooManyRequests with default headers values
func NewGetNotificationsDetailedTranslatedV1TooManyRequests() *GetNotificationsDetailedTranslatedV1TooManyRequests {
	return &GetNotificationsDetailedTranslatedV1TooManyRequests{}
}

/* GetNotificationsDetailedTranslatedV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetNotificationsDetailedTranslatedV1TooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *GetNotificationsDetailedTranslatedV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed-translated/v1][%d] getNotificationsDetailedTranslatedV1TooManyRequests  %+v", 429, o.Payload)
}
func (o *GetNotificationsDetailedTranslatedV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedTranslatedV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedTranslatedV1InternalServerError creates a GetNotificationsDetailedTranslatedV1InternalServerError with default headers values
func NewGetNotificationsDetailedTranslatedV1InternalServerError() *GetNotificationsDetailedTranslatedV1InternalServerError {
	return &GetNotificationsDetailedTranslatedV1InternalServerError{}
}

/* GetNotificationsDetailedTranslatedV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetNotificationsDetailedTranslatedV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainErrorsOnly
}

func (o *GetNotificationsDetailedTranslatedV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed-translated/v1][%d] getNotificationsDetailedTranslatedV1InternalServerError  %+v", 500, o.Payload)
}
func (o *GetNotificationsDetailedTranslatedV1InternalServerError) GetPayload() *models.DomainErrorsOnly {
	return o.Payload
}

func (o *GetNotificationsDetailedTranslatedV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsDetailedTranslatedV1Default creates a GetNotificationsDetailedTranslatedV1Default with default headers values
func NewGetNotificationsDetailedTranslatedV1Default(code int) *GetNotificationsDetailedTranslatedV1Default {
	return &GetNotificationsDetailedTranslatedV1Default{
		_statusCode: code,
	}
}

/* GetNotificationsDetailedTranslatedV1Default describes a response with status code -1, with default header values.

OK
*/
type GetNotificationsDetailedTranslatedV1Default struct {
	_statusCode int

	Payload *models.DomainNotificationDetailsResponseV1
}

// Code gets the status code for the get notifications detailed translated v1 default response
func (o *GetNotificationsDetailedTranslatedV1Default) Code() int {
	return o._statusCode
}

func (o *GetNotificationsDetailedTranslatedV1Default) Error() string {
	return fmt.Sprintf("[GET /recon/entities/notifications-detailed-translated/v1][%d] GetNotificationsDetailedTranslatedV1 default  %+v", o._statusCode, o.Payload)
}
func (o *GetNotificationsDetailedTranslatedV1Default) GetPayload() *models.DomainNotificationDetailsResponseV1 {
	return o.Payload
}

func (o *GetNotificationsDetailedTranslatedV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainNotificationDetailsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}