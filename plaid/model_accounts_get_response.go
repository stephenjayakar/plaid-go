/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.39.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// AccountsGetResponse AccountsGetResponse defines the response schema for `/accounts/get` and `/accounts/balance/get`.
type AccountsGetResponse struct {
	// An array of financial institution accounts associated with the Item. If `/accounts/balance/get` was called, each account will include real-time balance information.
	Accounts []AccountBase `json:"accounts"`
	Item Item `json:"item"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _AccountsGetResponse AccountsGetResponse

// NewAccountsGetResponse instantiates a new AccountsGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountsGetResponse(accounts []AccountBase, item Item, requestId string) *AccountsGetResponse {
	this := AccountsGetResponse{}
	this.Accounts = accounts
	this.Item = item
	this.RequestId = requestId
	return &this
}

// NewAccountsGetResponseWithDefaults instantiates a new AccountsGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountsGetResponseWithDefaults() *AccountsGetResponse {
	this := AccountsGetResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *AccountsGetResponse) GetAccounts() []AccountBase {
	if o == nil {
		var ret []AccountBase
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *AccountsGetResponse) GetAccountsOk() (*[]AccountBase, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *AccountsGetResponse) SetAccounts(v []AccountBase) {
	o.Accounts = v
}

// GetItem returns the Item field value
func (o *AccountsGetResponse) GetItem() Item {
	if o == nil {
		var ret Item
		return ret
	}

	return o.Item
}

// GetItemOk returns a tuple with the Item field value
// and a boolean to check if the value has been set.
func (o *AccountsGetResponse) GetItemOk() (*Item, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Item, true
}

// SetItem sets field value
func (o *AccountsGetResponse) SetItem(v Item) {
	o.Item = v
}

// GetRequestId returns the RequestId field value
func (o *AccountsGetResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *AccountsGetResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *AccountsGetResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o AccountsGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	if true {
		toSerialize["item"] = o.Item
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountsGetResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAccountsGetResponse := _AccountsGetResponse{}

	if err = json.Unmarshal(bytes, &varAccountsGetResponse); err == nil {
		*o = AccountsGetResponse(varAccountsGetResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accounts")
		delete(additionalProperties, "item")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountsGetResponse struct {
	value *AccountsGetResponse
	isSet bool
}

func (v NullableAccountsGetResponse) Get() *AccountsGetResponse {
	return v.value
}

func (v *NullableAccountsGetResponse) Set(val *AccountsGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountsGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountsGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountsGetResponse(val *AccountsGetResponse) *NullableAccountsGetResponse {
	return &NullableAccountsGetResponse{value: val, isSet: true}
}

func (v NullableAccountsGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountsGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

