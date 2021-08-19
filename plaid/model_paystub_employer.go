/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.21.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// PaystubEmployer struct for PaystubEmployer
type PaystubEmployer struct {
	// The name of the employer on the paystub.
	Name NullableString `json:"name"`
	AdditionalProperties map[string]interface{}
}

type _PaystubEmployer PaystubEmployer

// NewPaystubEmployer instantiates a new PaystubEmployer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaystubEmployer(name NullableString) *PaystubEmployer {
	this := PaystubEmployer{}
	this.Name = name
	return &this
}

// NewPaystubEmployerWithDefaults instantiates a new PaystubEmployer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaystubEmployerWithDefaults() *PaystubEmployer {
	this := PaystubEmployer{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaystubEmployer) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaystubEmployer) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *PaystubEmployer) SetName(v string) {
	o.Name.Set(&v)
}

func (o PaystubEmployer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaystubEmployer) UnmarshalJSON(bytes []byte) (err error) {
	varPaystubEmployer := _PaystubEmployer{}

	if err = json.Unmarshal(bytes, &varPaystubEmployer); err == nil {
		*o = PaystubEmployer(varPaystubEmployer)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaystubEmployer struct {
	value *PaystubEmployer
	isSet bool
}

func (v NullablePaystubEmployer) Get() *PaystubEmployer {
	return v.value
}

func (v *NullablePaystubEmployer) Set(val *PaystubEmployer) {
	v.value = val
	v.isSet = true
}

func (v NullablePaystubEmployer) IsSet() bool {
	return v.isSet
}

func (v *NullablePaystubEmployer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaystubEmployer(val *PaystubEmployer) *NullablePaystubEmployer {
	return &NullablePaystubEmployer{value: val, isSet: true}
}

func (v NullablePaystubEmployer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaystubEmployer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


