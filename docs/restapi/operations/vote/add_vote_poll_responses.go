// Code generated by go-swagger; DO NOT EDIT.

package vote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// AddVotePollNoContentCode is the HTTP code returned for type AddVotePollNoContent
const AddVotePollNoContentCode int = 204

/*AddVotePollNoContent Vote added

swagger:response addVotePollNoContent
*/
type AddVotePollNoContent struct {
}

// NewAddVotePollNoContent creates AddVotePollNoContent with default headers values
func NewAddVotePollNoContent() *AddVotePollNoContent {

	return &AddVotePollNoContent{}
}

// WriteResponse to the client
func (o *AddVotePollNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// AddVotePollNotFoundCode is the HTTP code returned for type AddVotePollNotFound
const AddVotePollNotFoundCode int = 404

/*AddVotePollNotFound Resource not found

swagger:response addVotePollNotFound
*/
type AddVotePollNotFound struct {
}

// NewAddVotePollNotFound creates AddVotePollNotFound with default headers values
func NewAddVotePollNotFound() *AddVotePollNotFound {

	return &AddVotePollNotFound{}
}

// WriteResponse to the client
func (o *AddVotePollNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
