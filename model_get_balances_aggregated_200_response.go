/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.0-beta.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// GetBalancesAggregated200Response struct for GetBalancesAggregated200Response
type GetBalancesAggregated200Response struct {
	Data map[string]int32 `json:"data"`
}

// NewGetBalancesAggregated200Response instantiates a new GetBalancesAggregated200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalancesAggregated200Response(data map[string]int32) *GetBalancesAggregated200Response {
	this := GetBalancesAggregated200Response{}
	this.Data = data
	return &this
}

// NewGetBalancesAggregated200ResponseWithDefaults instantiates a new GetBalancesAggregated200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalancesAggregated200ResponseWithDefaults() *GetBalancesAggregated200Response {
	this := GetBalancesAggregated200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetBalancesAggregated200Response) GetData() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetBalancesAggregated200Response) GetDataOk() (*map[string]int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetBalancesAggregated200Response) SetData(v map[string]int32) {
	o.Data = v
}

func (o GetBalancesAggregated200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetBalancesAggregated200Response struct {
	value *GetBalancesAggregated200Response
	isSet bool
}

func (v NullableGetBalancesAggregated200Response) Get() *GetBalancesAggregated200Response {
	return v.value
}

func (v *NullableGetBalancesAggregated200Response) Set(val *GetBalancesAggregated200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalancesAggregated200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalancesAggregated200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalancesAggregated200Response(val *GetBalancesAggregated200Response) *NullableGetBalancesAggregated200Response {
	return &NullableGetBalancesAggregated200Response{value: val, isSet: true}
}

func (v NullableGetBalancesAggregated200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalancesAggregated200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


