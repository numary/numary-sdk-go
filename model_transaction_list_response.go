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

// TransactionListResponse struct for TransactionListResponse
type TransactionListResponse struct {
	Cursor *Cursor `json:"cursor,omitempty"`
	Data []Transaction `json:"data"`
}

// NewTransactionListResponse instantiates a new TransactionListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionListResponse(data []Transaction) *TransactionListResponse {
	this := TransactionListResponse{}
	this.Data = data
	return &this
}

// NewTransactionListResponseWithDefaults instantiates a new TransactionListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionListResponseWithDefaults() *TransactionListResponse {
	this := TransactionListResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *TransactionListResponse) GetCursor() Cursor {
	if o == nil || o.Cursor == nil {
		var ret Cursor
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionListResponse) GetCursorOk() (*Cursor, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *TransactionListResponse) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given Cursor and assigns it to the Cursor field.
func (o *TransactionListResponse) SetCursor(v Cursor) {
	o.Cursor = &v
}

// GetData returns the Data field value
func (o *TransactionListResponse) GetData() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransactionListResponse) GetDataOk() (*[]Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransactionListResponse) SetData(v []Transaction) {
	o.Data = v
}

func (o TransactionListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionListResponse struct {
	value *TransactionListResponse
	isSet bool
}

func (v NullableTransactionListResponse) Get() *TransactionListResponse {
	return v.value
}

func (v *NullableTransactionListResponse) Set(val *TransactionListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionListResponse(val *TransactionListResponse) *NullableTransactionListResponse {
	return &NullableTransactionListResponse{value: val, isSet: true}
}

func (v NullableTransactionListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


