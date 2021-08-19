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

// InvestmentsDefaultUpdateWebhook Fired when new or canceled transactions have been detected on an investment account.
type InvestmentsDefaultUpdateWebhook struct {
	// `INVESTMENTS_TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `DEFAULT_UPDATE`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Error NullableError `json:"error,omitempty"`
	// The number of new transactions reported since the last time this webhook was fired.
	NewInvestmentsTransactions float32 `json:"new_investments_transactions"`
	// The number of canceled transactions reported since the last time this webhook was fired.
	CanceledInvestmentsTransactions float32 `json:"canceled_investments_transactions"`
	AdditionalProperties map[string]interface{}
}

type _InvestmentsDefaultUpdateWebhook InvestmentsDefaultUpdateWebhook

// NewInvestmentsDefaultUpdateWebhook instantiates a new InvestmentsDefaultUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvestmentsDefaultUpdateWebhook(webhookType string, webhookCode string, itemId string, newInvestmentsTransactions float32, canceledInvestmentsTransactions float32) *InvestmentsDefaultUpdateWebhook {
	this := InvestmentsDefaultUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.NewInvestmentsTransactions = newInvestmentsTransactions
	this.CanceledInvestmentsTransactions = canceledInvestmentsTransactions
	return &this
}

// NewInvestmentsDefaultUpdateWebhookWithDefaults instantiates a new InvestmentsDefaultUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInvestmentsDefaultUpdateWebhookWithDefaults() *InvestmentsDefaultUpdateWebhook {
	this := InvestmentsDefaultUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *InvestmentsDefaultUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *InvestmentsDefaultUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *InvestmentsDefaultUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *InvestmentsDefaultUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *InvestmentsDefaultUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *InvestmentsDefaultUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *InvestmentsDefaultUpdateWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *InvestmentsDefaultUpdateWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *InvestmentsDefaultUpdateWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InvestmentsDefaultUpdateWebhook) GetError() Error {
	if o == nil || o.Error.Get() == nil {
		var ret Error
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InvestmentsDefaultUpdateWebhook) GetErrorOk() (*Error, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *InvestmentsDefaultUpdateWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableError and assigns it to the Error field.
func (o *InvestmentsDefaultUpdateWebhook) SetError(v Error) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *InvestmentsDefaultUpdateWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *InvestmentsDefaultUpdateWebhook) UnsetError() {
	o.Error.Unset()
}

// GetNewInvestmentsTransactions returns the NewInvestmentsTransactions field value
func (o *InvestmentsDefaultUpdateWebhook) GetNewInvestmentsTransactions() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewInvestmentsTransactions
}

// GetNewInvestmentsTransactionsOk returns a tuple with the NewInvestmentsTransactions field value
// and a boolean to check if the value has been set.
func (o *InvestmentsDefaultUpdateWebhook) GetNewInvestmentsTransactionsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewInvestmentsTransactions, true
}

// SetNewInvestmentsTransactions sets field value
func (o *InvestmentsDefaultUpdateWebhook) SetNewInvestmentsTransactions(v float32) {
	o.NewInvestmentsTransactions = v
}

// GetCanceledInvestmentsTransactions returns the CanceledInvestmentsTransactions field value
func (o *InvestmentsDefaultUpdateWebhook) GetCanceledInvestmentsTransactions() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.CanceledInvestmentsTransactions
}

// GetCanceledInvestmentsTransactionsOk returns a tuple with the CanceledInvestmentsTransactions field value
// and a boolean to check if the value has been set.
func (o *InvestmentsDefaultUpdateWebhook) GetCanceledInvestmentsTransactionsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CanceledInvestmentsTransactions, true
}

// SetCanceledInvestmentsTransactions sets field value
func (o *InvestmentsDefaultUpdateWebhook) SetCanceledInvestmentsTransactions(v float32) {
	o.CanceledInvestmentsTransactions = v
}

func (o InvestmentsDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["new_investments_transactions"] = o.NewInvestmentsTransactions
	}
	if true {
		toSerialize["canceled_investments_transactions"] = o.CanceledInvestmentsTransactions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InvestmentsDefaultUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varInvestmentsDefaultUpdateWebhook := _InvestmentsDefaultUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varInvestmentsDefaultUpdateWebhook); err == nil {
		*o = InvestmentsDefaultUpdateWebhook(varInvestmentsDefaultUpdateWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "error")
		delete(additionalProperties, "new_investments_transactions")
		delete(additionalProperties, "canceled_investments_transactions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInvestmentsDefaultUpdateWebhook struct {
	value *InvestmentsDefaultUpdateWebhook
	isSet bool
}

func (v NullableInvestmentsDefaultUpdateWebhook) Get() *InvestmentsDefaultUpdateWebhook {
	return v.value
}

func (v *NullableInvestmentsDefaultUpdateWebhook) Set(val *InvestmentsDefaultUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableInvestmentsDefaultUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableInvestmentsDefaultUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvestmentsDefaultUpdateWebhook(val *InvestmentsDefaultUpdateWebhook) *NullableInvestmentsDefaultUpdateWebhook {
	return &NullableInvestmentsDefaultUpdateWebhook{value: val, isSet: true}
}

func (v NullableInvestmentsDefaultUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvestmentsDefaultUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


