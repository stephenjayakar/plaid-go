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

// AssetReportCreateRequest AssetReportCreateRequest defines the request schema for `/asset_report/create`
type AssetReportCreateRequest struct {
	// Your Plaid API `client_id`. The `client_id` is required and may be provided either in the `PLAID-CLIENT-ID` header or as part of a request body.
	ClientId *string `json:"client_id,omitempty"`
	// Your Plaid API `secret`. The `secret` is required and may be provided either in the `PLAID-SECRET` header or as part of a request body.
	Secret *string `json:"secret,omitempty"`
	// An array of access tokens corresponding to the Items that will be included in the report. The `assets` product must have been initialized for the Items during link; the Assets product cannot be added after initialization.
	AccessTokens []string `json:"access_tokens"`
	// The maximum integer number of days of history to include in the Asset Report. If using Fannie Mae Day 1 Certainty, `days_requested` must be at least 61 for new originations or at least 31 for refinancings.
	DaysRequested int32 `json:"days_requested"`
	Options *AssetReportCreateRequestOptions `json:"options,omitempty"`
}

// NewAssetReportCreateRequest instantiates a new AssetReportCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetReportCreateRequest(accessTokens []string, daysRequested int32) *AssetReportCreateRequest {
	this := AssetReportCreateRequest{}
	this.AccessTokens = accessTokens
	this.DaysRequested = daysRequested
	return &this
}

// NewAssetReportCreateRequestWithDefaults instantiates a new AssetReportCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetReportCreateRequestWithDefaults() *AssetReportCreateRequest {
	this := AssetReportCreateRequest{}
	return &this
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AssetReportCreateRequest) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequest) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *AssetReportCreateRequest) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AssetReportCreateRequest) SetClientId(v string) {
	o.ClientId = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *AssetReportCreateRequest) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequest) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *AssetReportCreateRequest) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *AssetReportCreateRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetAccessTokens returns the AccessTokens field value
func (o *AssetReportCreateRequest) GetAccessTokens() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccessTokens
}

// GetAccessTokensOk returns a tuple with the AccessTokens field value
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequest) GetAccessTokensOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccessTokens, true
}

// SetAccessTokens sets field value
func (o *AssetReportCreateRequest) SetAccessTokens(v []string) {
	o.AccessTokens = v
}

// GetDaysRequested returns the DaysRequested field value
func (o *AssetReportCreateRequest) GetDaysRequested() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DaysRequested
}

// GetDaysRequestedOk returns a tuple with the DaysRequested field value
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequest) GetDaysRequestedOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DaysRequested, true
}

// SetDaysRequested sets field value
func (o *AssetReportCreateRequest) SetDaysRequested(v int32) {
	o.DaysRequested = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AssetReportCreateRequest) GetOptions() AssetReportCreateRequestOptions {
	if o == nil || o.Options == nil {
		var ret AssetReportCreateRequestOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetReportCreateRequest) GetOptionsOk() (*AssetReportCreateRequestOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AssetReportCreateRequest) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AssetReportCreateRequestOptions and assigns it to the Options field.
func (o *AssetReportCreateRequest) SetOptions(v AssetReportCreateRequestOptions) {
	o.Options = &v
}

func (o AssetReportCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientId != nil {
		toSerialize["client_id"] = o.ClientId
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if true {
		toSerialize["access_tokens"] = o.AccessTokens
	}
	if true {
		toSerialize["days_requested"] = o.DaysRequested
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableAssetReportCreateRequest struct {
	value *AssetReportCreateRequest
	isSet bool
}

func (v NullableAssetReportCreateRequest) Get() *AssetReportCreateRequest {
	return v.value
}

func (v *NullableAssetReportCreateRequest) Set(val *AssetReportCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetReportCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetReportCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetReportCreateRequest(val *AssetReportCreateRequest) *NullableAssetReportCreateRequest {
	return &NullableAssetReportCreateRequest{value: val, isSet: true}
}

func (v NullableAssetReportCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetReportCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


