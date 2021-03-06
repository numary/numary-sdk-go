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

// TransactionCursorAllOf struct for TransactionCursorAllOf
type TransactionCursorAllOf struct {
	Data []Transaction `json:"data"`
}

// NewTransactionCursorAllOf instantiates a new TransactionCursorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionCursorAllOf(data []Transaction) *TransactionCursorAllOf {
	this := TransactionCursorAllOf{}
	this.Data = data
	return &this
}

// NewTransactionCursorAllOfWithDefaults instantiates a new TransactionCursorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionCursorAllOfWithDefaults() *TransactionCursorAllOf {
	this := TransactionCursorAllOf{}
	return &this
}

// GetData returns the Data field value
func (o *TransactionCursorAllOf) GetData() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionCursorAllOf) GetDataOk() (*[]Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionCursorAllOf) SetData(v []Transaction) {
	o.Data = v
}

func (o TransactionCursorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionCursorAllOf struct {
	value *TransactionCursorAllOf
	isSet bool
}

func (v NullableTransactionCursorAllOf) Get() *TransactionCursorAllOf {
	return v.value
}

func (v *NullableTransactionCursorAllOf) Set(val *TransactionCursorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionCursorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionCursorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionCursorAllOf(val *TransactionCursorAllOf) *NullableTransactionCursorAllOf {
	return &NullableTransactionCursorAllOf{value: val, isSet: true}
}

func (v NullableTransactionCursorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionCursorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


