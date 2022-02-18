// Code generated by go-swagger; DO NOT EDIT.

package campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/graphql-editor/woodpecker/models"
)

// GetCampaignListReader is a Reader for the GetCampaignList structure.
type GetCampaignListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCampaignListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCampaignListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetCampaignListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCampaignListOK creates a GetCampaignListOK with default headers values
func NewGetCampaignListOK() *GetCampaignListOK {
	return &GetCampaignListOK{}
}

/* GetCampaignListOK describes a response with status code 200, with default header values.

An array of campaigns
*/
type GetCampaignListOK struct {
	Payload []*models.Campaign
}

func (o *GetCampaignListOK) Error() string {
	return fmt.Sprintf("[GET /campaign_list][%d] getCampaignListOK  %+v", 200, o.Payload)
}
func (o *GetCampaignListOK) GetPayload() []*models.Campaign {
	return o.Payload
}

func (o *GetCampaignListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCampaignListDefault creates a GetCampaignListDefault with default headers values
func NewGetCampaignListDefault(code int) *GetCampaignListDefault {
	return &GetCampaignListDefault{
		_statusCode: code,
	}
}

/* GetCampaignListDefault describes a response with status code -1, with default header values.

Unexpected error
*/
type GetCampaignListDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get campaign list default response
func (o *GetCampaignListDefault) Code() int {
	return o._statusCode
}

func (o *GetCampaignListDefault) Error() string {
	return fmt.Sprintf("[GET /campaign_list][%d] GetCampaignList default  %+v", o._statusCode, o.Payload)
}
func (o *GetCampaignListDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCampaignListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}