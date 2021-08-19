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

// InvestmentsHoldingsGetRequest InvestmentsHoldingsGetRequest defines the request schema for `/investments/holdings/get`
type InvestmentsHoldingsGetRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// The access token associated with the Item data is being requested for.
	AccessToken string `json:"access_token"`
	Options *InvestmentHoldingsGetRequestOptions `json:"options,omitempty"`
}

// NewInvestmentsHoldingsGetRequest instantiates a new InvestmentsHoldingsGetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsHoldingsGetRequest(accessToken string) *InvestmentsHoldingsGetRequest {
	this := InvestmentsHoldingsGetRequest{}
	this.AccessToken = accessToken
	return &this
}

// NewInvestmentsHoldingsGetRequestWithDefaults instantiates a new InvestmentsHoldingsGetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsHoldingsGetRequestWithDefaults() *InvestmentsHoldingsGetRequest {
	this := InvestmentsHoldingsGetRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *InvestmentsHoldingsGetRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *InvestmentsHoldingsGetRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *InvestmentsHoldingsGetRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *InvestmentsHoldingsGetRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *InvestmentsHoldingsGetRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *InvestmentsHoldingsGetRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessToken returns the AccessToken field value
func (o *InvestmentsHoldingsGetRequest) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetRequest) GetAccessTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *InvestmentsHoldingsGetRequest) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *InvestmentsHoldingsGetRequest) GetOptions() InvestmentHoldingsGetRequestOptions {
	if o == nil || o.Options == nil {
		var ret InvestmentHoldingsGetRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InvestmentsHoldingsGetRequest) GetOptionsOk() (*InvestmentHoldingsGetRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *InvestmentsHoldingsGetRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given InvestmentHoldingsGetRequestOptions and assigns it to the Options field.
func (o *InvestmentsHoldingsGetRequest) SetOptions(v InvestmentHoldingsGetRequestOptions) {
	o.Options = &v
}

func (o InvestmentsHoldingsGetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableInvestmentsHoldingsGetRequest struct {
	value *InvestmentsHoldingsGetRequest
	isSet bool
}

func (v NullableInvestmentsHoldingsGetRequest) Get() *InvestmentsHoldingsGetRequest {
	return v.value
}

func (v *NullableInvestmentsHoldingsGetRequest) Set(val *InvestmentsHoldingsGetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsHoldingsGetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsHoldingsGetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsHoldingsGetRequest(val *InvestmentsHoldingsGetRequest) *NullableInvestmentsHoldingsGetRequest {
	return &NullableInvestmentsHoldingsGetRequest{value: val, isSet: true}
}

func (v NullableInvestmentsHoldingsGetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsHoldingsGetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


