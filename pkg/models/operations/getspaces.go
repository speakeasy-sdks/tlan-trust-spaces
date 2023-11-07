// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"tlan-trust-spaces/v2/pkg/models/shared"
)

type GetSpacesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// ok
	SpaceList *shared.SpaceList
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetSpacesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSpacesResponse) GetSpaceList() *shared.SpaceList {
	if o == nil {
		return nil
	}
	return o.SpaceList
}

func (o *GetSpacesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSpacesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
