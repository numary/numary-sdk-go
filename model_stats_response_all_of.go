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

// StatsResponseAllOf struct for StatsResponseAllOf
type StatsResponseAllOf struct {
	Data Stats `json:"data"`
}

// NewStatsResponseAllOf instantiates a new StatsResponseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatsResponseAllOf(data Stats) *StatsResponseAllOf {
	this := StatsResponseAllOf{}
	this.Data = data
	return &this
}

// NewStatsResponseAllOfWithDefaults instantiates a new StatsResponseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatsResponseAllOfWithDefaults() *StatsResponseAllOf {
	this := StatsResponseAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *StatsResponseAllOf) GetData() Stats {
	if o == nil {
		var ret Stats
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *StatsResponseAllOf) GetDataOk() (*Stats, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *StatsResponseAllOf) SetData(v Stats) {
	o.Data = v
}

func (o StatsResponseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableStatsResponseAllOf struct {
	value *StatsResponseAllOf
	isSet bool
}

func (v NullableStatsResponseAllOf) Get() *StatsResponseAllOf {
	return v.value
}

func (v *NullableStatsResponseAllOf) Set(val *StatsResponseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStatsResponseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStatsResponseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatsResponseAllOf(val *StatsResponseAllOf) *NullableStatsResponseAllOf {
	return &NullableStatsResponseAllOf{value: val, isSet: true}
}

func (v NullableStatsResponseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatsResponseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


