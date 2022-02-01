/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.61.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// WalletTransactionCounterpartyNumbers The counterparty's bank account numbers
type WalletTransactionCounterpartyNumbers struct {
	Bacs WalletTransactionCounterpartyBACS `json:"bacs"`
}

// NewWalletTransactionCounterpartyNumbers instantiates a new WalletTransactionCounterpartyNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletTransactionCounterpartyNumbers(bacs WalletTransactionCounterpartyBACS) *WalletTransactionCounterpartyNumbers {
	this := WalletTransactionCounterpartyNumbers{}
	this.Bacs = bacs
	return &this
}

// NewWalletTransactionCounterpartyNumbersWithDefaults instantiates a new WalletTransactionCounterpartyNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletTransactionCounterpartyNumbersWithDefaults() *WalletTransactionCounterpartyNumbers {
	this := WalletTransactionCounterpartyNumbers{}
	return &this
}

// GetBacs returns the Bacs field value
func (o *WalletTransactionCounterpartyNumbers) GetBacs() WalletTransactionCounterpartyBACS {
	if o == nil {
		var ret WalletTransactionCounterpartyBACS
		return ret
	}

	return o.Bacs
}

// GetBacsOk returns a tuple with the Bacs field value
// and a boolean to check if the value has been set.
func (o *WalletTransactionCounterpartyNumbers) GetBacsOk() (*WalletTransactionCounterpartyBACS, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bacs, true
}

// SetBacs sets field value
func (o *WalletTransactionCounterpartyNumbers) SetBacs(v WalletTransactionCounterpartyBACS) {
	o.Bacs = v
}

func (o WalletTransactionCounterpartyNumbers) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bacs"] = o.Bacs
	}
	return json.Marshal(toSerialize)
}

type NullableWalletTransactionCounterpartyNumbers struct {
	value *WalletTransactionCounterpartyNumbers
	isSet bool
}

func (v NullableWalletTransactionCounterpartyNumbers) Get() *WalletTransactionCounterpartyNumbers {
	return v.value
}

func (v *NullableWalletTransactionCounterpartyNumbers) Set(val *WalletTransactionCounterpartyNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletTransactionCounterpartyNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletTransactionCounterpartyNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletTransactionCounterpartyNumbers(val *WalletTransactionCounterpartyNumbers) *NullableWalletTransactionCounterpartyNumbers {
	return &NullableWalletTransactionCounterpartyNumbers{value: val, isSet: true}
}

func (v NullableWalletTransactionCounterpartyNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletTransactionCounterpartyNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

