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

// MappingResponseAllOf struct for MappingResponseAllOf
type MappingResponseAllOf struct {
	Data Mapping `json:"data"`
}

// NewMappingResponseAllOf instantiates a new MappingResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMappingResponseAllOf(data Mapping) *MappingResponseAllOf {
	this := MappingResponseAllOf{}
	this.Data = data
	return &this
}

// NewMappingResponseAllOfWithDefaults instantiates a new MappingResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingResponseAllOfWithDefaults() *MappingResponseAllOf {
	this := MappingResponseAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *MappingResponseAllOf) GetData() Mapping {
	if o == nil {
		var ret Mapping
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *MappingResponseAllOf) GetDataOk() (*Mapping, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *MappingResponseAllOf) SetData(v Mapping) {
	o.Data = v
}

func (o MappingResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableMappingResponseAllOf struct {
	value *MappingResponseAllOf
	isSet bool
}

func (v NullableMappingResponseAllOf) Get() *MappingResponseAllOf {
	return v.value
}

func (v *NullableMappingResponseAllOf) Set(val *MappingResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMappingResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMappingResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMappingResponseAllOf(val *MappingResponseAllOf) *NullableMappingResponseAllOf {
	return &NullableMappingResponseAllOf{value: val, isSet: true}
}

func (v NullableMappingResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMappingResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

