/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// SingleObjectResponse struct for SingleObjectResponse
type SingleObjectResponse struct {
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewSingleObjectResponse instantiates a new SingleObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleObjectResponse() *SingleObjectResponse {
	this := SingleObjectResponse{}
	return &this
}

// NewSingleObjectResponseWithDefaults instantiates a new SingleObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleObjectResponseWithDefaults() *SingleObjectResponse {
	this := SingleObjectResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SingleObjectResponse) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleObjectResponse) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SingleObjectResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *SingleObjectResponse) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o SingleObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableSingleObjectResponse struct {
	value *SingleObjectResponse
	isSet bool
}

func (v NullableSingleObjectResponse) Get() *SingleObjectResponse {
	return v.value
}

func (v *NullableSingleObjectResponse) Set(val *SingleObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleObjectResponse(val *SingleObjectResponse) *NullableSingleObjectResponse {
	return &NullableSingleObjectResponse{value: val, isSet: true}
}

func (v NullableSingleObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

