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

// HistoricalUpdateWebhook Fired when an Item's historical transaction pull is completed and Plaid has prepared as much historical transaction data as possible for the Item. Once this webhook has been fired, transaction data beyond the most recent 30 days can be fetched for the Item.
type HistoricalUpdateWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `HISTORICAL_UPDATE`
	WebhookCode string `json:"webhook_code"`
	Error NullableError `json:"error,omitempty"`
	// The number of new, unfetched transactions available
	NewTransactions float32 `json:"new_transactions"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	AdditionalProperties map[string]interface{}
}

type _HistoricalUpdateWebhook HistoricalUpdateWebhook

// NewHistoricalUpdateWebhook instantiates a new HistoricalUpdateWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHistoricalUpdateWebhook(webhookType string, webhookCode string, newTransactions float32, itemId string) *HistoricalUpdateWebhook {
	this := HistoricalUpdateWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.NewTransactions = newTransactions
	this.ItemId = itemId
	return &this
}

// NewHistoricalUpdateWebhookWithDefaults instantiates a new HistoricalUpdateWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHistoricalUpdateWebhookWithDefaults() *HistoricalUpdateWebhook {
	this := HistoricalUpdateWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *HistoricalUpdateWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *HistoricalUpdateWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *HistoricalUpdateWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *HistoricalUpdateWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *HistoricalUpdateWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *HistoricalUpdateWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *HistoricalUpdateWebhook) GetError() Error {
	if o == nil || o.Error.Get() == nil {
		var ret Error
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *HistoricalUpdateWebhook) GetErrorOk() (*Error, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *HistoricalUpdateWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableError and assigns it to the Error field.
func (o *HistoricalUpdateWebhook) SetError(v Error) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *HistoricalUpdateWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *HistoricalUpdateWebhook) UnsetError() {
	o.Error.Unset()
}

// GetNewTransactions returns the NewTransactions field value
func (o *HistoricalUpdateWebhook) GetNewTransactions() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NewTransactions
}

// GetNewTransactionsOk returns a tuple with the NewTransactions field value
// and a boolean to check if the value has been set.
func (o *HistoricalUpdateWebhook) GetNewTransactionsOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NewTransactions, true
}

// SetNewTransactions sets field value
func (o *HistoricalUpdateWebhook) SetNewTransactions(v float32) {
	o.NewTransactions = v
}

// GetItemId returns the ItemId field value
func (o *HistoricalUpdateWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *HistoricalUpdateWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *HistoricalUpdateWebhook) SetItemId(v string) {
	o.ItemId = v
}

func (o HistoricalUpdateWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["new_transactions"] = o.NewTransactions
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HistoricalUpdateWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varHistoricalUpdateWebhook := _HistoricalUpdateWebhook{}

	if err = json.Unmarshal(bytes, &varHistoricalUpdateWebhook); err == nil {
		*o = HistoricalUpdateWebhook(varHistoricalUpdateWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "new_transactions")
		delete(additionalProperties, "item_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHistoricalUpdateWebhook struct {
	value *HistoricalUpdateWebhook
	isSet bool
}

func (v NullableHistoricalUpdateWebhook) Get() *HistoricalUpdateWebhook {
	return v.value
}

func (v *NullableHistoricalUpdateWebhook) Set(val *HistoricalUpdateWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableHistoricalUpdateWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableHistoricalUpdateWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHistoricalUpdateWebhook(val *HistoricalUpdateWebhook) *NullableHistoricalUpdateWebhook {
	return &NullableHistoricalUpdateWebhook{value: val, isSet: true}
}

func (v NullableHistoricalUpdateWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHistoricalUpdateWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


