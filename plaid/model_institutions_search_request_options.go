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

// InstitutionsSearchRequestOptions An optional object to filter `/institutions/search` results.
type InstitutionsSearchRequestOptions struct {
	// Limit results to institutions with or without OAuth login flows. This is primarily relevant to institutions with European country codes
	Oauth *bool `json:"oauth,omitempty"`
	// When true, return the institution's homepage URL, logo and primary brand color.
	IncludeOptionalMetadata *bool `json:"include_optional_metadata,omitempty"`
	// When `true`, returns metadata related to the Payment Initiation product indicating which payment configurations are supported.
	IncludePaymentInitiationMetadata *bool `json:"include_payment_initiation_metadata,omitempty"`
	PaymentInitiation NullableInstitutionsSearchPaymentInitiationOptions `json:"payment_initiation,omitempty"`
}

// NewInstitutionsSearchRequestOptions instantiates a new InstitutionsSearchRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstitutionsSearchRequestOptions() *InstitutionsSearchRequestOptions {
	this := InstitutionsSearchRequestOptions{}
	var includePaymentInitiationMetadata bool = false
	this.IncludePaymentInitiationMetadata = &includePaymentInitiationMetadata
	return &this
}

// NewInstitutionsSearchRequestOptionsWithDefaults instantiates a new InstitutionsSearchRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstitutionsSearchRequestOptionsWithDefaults() *InstitutionsSearchRequestOptions {
	this := InstitutionsSearchRequestOptions{}
	var includePaymentInitiationMetadata bool = false
	this.IncludePaymentInitiationMetadata = &includePaymentInitiationMetadata
	return &this
}

// GetOauth returns the Oauth field value if set, zero value otherwise.
func (o *InstitutionsSearchRequestOptions) GetOauth() bool {
	if o == nil || o.Oauth == nil {
		var ret bool
		return ret
	}
	return *o.Oauth
}

// GetOauthOk returns a tuple with the Oauth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequestOptions) GetOauthOk() (*bool, bool) {
	if o == nil || o.Oauth == nil {
		return nil, false
	}
	return o.Oauth, true
}

// HasOauth returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasOauth() bool {
	if o != nil && o.Oauth != nil {
		return true
	}

	return false
}

// SetOauth gets a reference to the given bool and assigns it to the Oauth field.
func (o *InstitutionsSearchRequestOptions) SetOauth(v bool) {
	o.Oauth = &v
}

// GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field value if set, zero value otherwise.
func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadata() bool {
	if o == nil || o.IncludeOptionalMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOptionalMetadata
}

// GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool) {
	if o == nil || o.IncludeOptionalMetadata == nil {
		return nil, false
	}
	return o.IncludeOptionalMetadata, true
}

// HasIncludeOptionalMetadata returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasIncludeOptionalMetadata() bool {
	if o != nil && o.IncludeOptionalMetadata != nil {
		return true
	}

	return false
}

// SetIncludeOptionalMetadata gets a reference to the given bool and assigns it to the IncludeOptionalMetadata field.
func (o *InstitutionsSearchRequestOptions) SetIncludeOptionalMetadata(v bool) {
	o.IncludeOptionalMetadata = &v
}

// GetIncludePaymentInitiationMetadata returns the IncludePaymentInitiationMetadata field value if set, zero value otherwise.
func (o *InstitutionsSearchRequestOptions) GetIncludePaymentInitiationMetadata() bool {
	if o == nil || o.IncludePaymentInitiationMetadata == nil {
		var ret bool
		return ret
	}
	return *o.IncludePaymentInitiationMetadata
}

// GetIncludePaymentInitiationMetadataOk returns a tuple with the IncludePaymentInitiationMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstitutionsSearchRequestOptions) GetIncludePaymentInitiationMetadataOk() (*bool, bool) {
	if o == nil || o.IncludePaymentInitiationMetadata == nil {
		return nil, false
	}
	return o.IncludePaymentInitiationMetadata, true
}

// HasIncludePaymentInitiationMetadata returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasIncludePaymentInitiationMetadata() bool {
	if o != nil && o.IncludePaymentInitiationMetadata != nil {
		return true
	}

	return false
}

// SetIncludePaymentInitiationMetadata gets a reference to the given bool and assigns it to the IncludePaymentInitiationMetadata field.
func (o *InstitutionsSearchRequestOptions) SetIncludePaymentInitiationMetadata(v bool) {
	o.IncludePaymentInitiationMetadata = &v
}

// GetPaymentInitiation returns the PaymentInitiation field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InstitutionsSearchRequestOptions) GetPaymentInitiation() InstitutionsSearchPaymentInitiationOptions {
	if o == nil || o.PaymentInitiation.Get() == nil {
		var ret InstitutionsSearchPaymentInitiationOptions
		return ret
	}
	return *o.PaymentInitiation.Get()
}

// GetPaymentInitiationOk returns a tuple with the PaymentInitiation field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InstitutionsSearchRequestOptions) GetPaymentInitiationOk() (*InstitutionsSearchPaymentInitiationOptions, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PaymentInitiation.Get(), o.PaymentInitiation.IsSet()
}

// HasPaymentInitiation returns a boolean if a field has been set.
func (o *InstitutionsSearchRequestOptions) HasPaymentInitiation() bool {
	if o != nil && o.PaymentInitiation.IsSet() {
		return true
	}

	return false
}

// SetPaymentInitiation gets a reference to the given NullableInstitutionsSearchPaymentInitiationOptions and assigns it to the PaymentInitiation field.
func (o *InstitutionsSearchRequestOptions) SetPaymentInitiation(v InstitutionsSearchPaymentInitiationOptions) {
	o.PaymentInitiation.Set(&v)
}
// SetPaymentInitiationNil sets the value for PaymentInitiation to be an explicit nil
func (o *InstitutionsSearchRequestOptions) SetPaymentInitiationNil() {
	o.PaymentInitiation.Set(nil)
}

// UnsetPaymentInitiation ensures that no value is present for PaymentInitiation, not even an explicit nil
func (o *InstitutionsSearchRequestOptions) UnsetPaymentInitiation() {
	o.PaymentInitiation.Unset()
}

func (o InstitutionsSearchRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Oauth != nil {
		toSerialize["oauth"] = o.Oauth
	}
	if o.IncludeOptionalMetadata != nil {
		toSerialize["include_optional_metadata"] = o.IncludeOptionalMetadata
	}
	if o.IncludePaymentInitiationMetadata != nil {
		toSerialize["include_payment_initiation_metadata"] = o.IncludePaymentInitiationMetadata
	}
	if o.PaymentInitiation.IsSet() {
		toSerialize["payment_initiation"] = o.PaymentInitiation.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInstitutionsSearchRequestOptions struct {
	value *InstitutionsSearchRequestOptions
	isSet bool
}

func (v NullableInstitutionsSearchRequestOptions) Get() *InstitutionsSearchRequestOptions {
	return v.value
}

func (v *NullableInstitutionsSearchRequestOptions) Set(val *InstitutionsSearchRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableInstitutionsSearchRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableInstitutionsSearchRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstitutionsSearchRequestOptions(val *InstitutionsSearchRequestOptions) *NullableInstitutionsSearchRequestOptions {
	return &NullableInstitutionsSearchRequestOptions{value: val, isSet: true}
}

func (v NullableInstitutionsSearchRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstitutionsSearchRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


