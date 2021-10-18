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

// TransactionBase A representation of a transaction
type TransactionBase struct {
	// Please use the `payment_channel` field, `transaction_type` will be deprecated in the future.  `digital:` transactions that took place online.  `place:` transactions that were made at a physical location.  `special:` transactions that relate to banks, e.g. fees or deposits.  `unresolved:` transactions that do not fit into the other three types. 
	TransactionType *string `json:"transaction_type,omitempty"`
	// The ID of a posted transaction's associated pending transaction, where applicable.
	PendingTransactionId NullableString `json:"pending_transaction_id,omitempty"`
	// The ID of the category to which this transaction belongs. See [Categories](https://plaid.com/docs/#category-overview).  If the `transactions` object was returned by an Assets endpoint such as `/asset_report/get/` or `/asset_report/pdf/get`, this field will only appear in an Asset Report with Insights.
	CategoryId NullableString `json:"category_id,omitempty"`
	// A hierarchical array of the categories to which this transaction belongs. See [Categories](https://plaid.com/docs/#category-overview).  If the `transactions` object was returned by an Assets endpoint such as `/asset_report/get/` or `/asset_report/pdf/get`, this field will only appear in an Asset Report with Insights.
	Category []string `json:"category,omitempty"`
	Location *Location `json:"location,omitempty"`
	PaymentMeta *PaymentMeta `json:"payment_meta,omitempty"`
	// The name of the account owner. This field is not typically populated and only relevant when dealing with sub-accounts.
	AccountOwner NullableString `json:"account_owner,omitempty"`
	// The merchant name or transaction description.  If the `transactions` object was returned by a Transactions endpoint such as `/transactions/get`, this field will always appear. If the `transactions` object was returned by an Assets endpoint such as `/asset_report/get/` or `/asset_report/pdf/get`, this field will only appear in an Asset Report with Insights.
	Name *string `json:"name,omitempty"`
	// The string returned by the financial institution to describe the transaction. For transactions returned by `/transactions/get`, this field is in beta and will be omitted unless the client is both enrolled in the closed beta program and has set `options.include_original_description` to `true`.
	OriginalDescription NullableString `json:"original_description,omitempty"`
	// The ID of the account in which this transaction occurred.
	AccountId string `json:"account_id"`
	// The settled value of the transaction, denominated in the account's currency, as stated in `iso_currency_code` or `unofficial_currency_code`. Positive values when money moves out of the account; negative values when money moves in. For example, debit card purchases are positive; credit card payments, direct deposits, and refunds are negative.
	Amount float32 `json:"amount"`
	// The ISO-4217 currency code of the transaction. Always `null` if `unofficial_currency_code` is non-null.
	IsoCurrencyCode NullableString `json:"iso_currency_code"`
	// The unofficial currency code associated with the transaction. Always `null` if `iso_currency_code` is non-`null`. Unofficial currency codes are used for currencies that do not have official ISO currency codes, such as cryptocurrencies and the currencies of certain countries.  See the [currency code schema](https://plaid.com/docs/api/accounts#currency-code-schema) for a full listing of supported `iso_currency_code`s.
	UnofficialCurrencyCode NullableString `json:"unofficial_currency_code"`
	// For pending transactions, the date that the transaction occurred; for posted transactions, the date that the transaction posted. Both dates are returned in an [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format ( `YYYY-MM-DD` ).
	Date string `json:"date"`
	// When `true`, identifies the transaction as pending or unsettled. Pending transaction details (name, type, amount, category ID) may change before they are settled.
	Pending bool `json:"pending"`
	// The unique ID of the transaction. Like all Plaid identifiers, the `transaction_id` is case sensitive.
	TransactionId string `json:"transaction_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionBase TransactionBase

// NewTransactionBase instantiates a new TransactionBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionBase(accountId string, amount float32, isoCurrencyCode NullableString, unofficialCurrencyCode NullableString, date string, pending bool, transactionId string) *TransactionBase {
	this := TransactionBase{}
	this.AccountId = accountId
	this.Amount = amount
	this.IsoCurrencyCode = isoCurrencyCode
	this.UnofficialCurrencyCode = unofficialCurrencyCode
	this.Date = date
	this.Pending = pending
	this.TransactionId = transactionId
	return &this
}

// NewTransactionBaseWithDefaults instantiates a new TransactionBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionBaseWithDefaults() *TransactionBase {
	this := TransactionBase{}
	return &this
}

// GetTransactionType returns the TransactionType field value if set, zero value otherwise.
func (o *TransactionBase) GetTransactionType() string {
	if o == nil || o.TransactionType == nil {
		var ret string
		return ret
	}
	return *o.TransactionType
}

// GetTransactionTypeOk returns a tuple with the TransactionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetTransactionTypeOk() (*string, bool) {
	if o == nil || o.TransactionType == nil {
		return nil, false
	}
	return o.TransactionType, true
}

// HasTransactionType returns a boolean if a field has been set.
func (o *TransactionBase) HasTransactionType() bool {
	if o != nil && o.TransactionType != nil {
		return true
	}

	return false
}

// SetTransactionType gets a reference to the given string and assigns it to the TransactionType field.
func (o *TransactionBase) SetTransactionType(v string) {
	o.TransactionType = &v
}

// GetPendingTransactionId returns the PendingTransactionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionBase) GetPendingTransactionId() string {
	if o == nil || o.PendingTransactionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PendingTransactionId.Get()
}

// GetPendingTransactionIdOk returns a tuple with the PendingTransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionBase) GetPendingTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PendingTransactionId.Get(), o.PendingTransactionId.IsSet()
}

// HasPendingTransactionId returns a boolean if a field has been set.
func (o *TransactionBase) HasPendingTransactionId() bool {
	if o != nil && o.PendingTransactionId.IsSet() {
		return true
	}

	return false
}

// SetPendingTransactionId gets a reference to the given NullableString and assigns it to the PendingTransactionId field.
func (o *TransactionBase) SetPendingTransactionId(v string) {
	o.PendingTransactionId.Set(&v)
}
// SetPendingTransactionIdNil sets the value for PendingTransactionId to be an explicit nil
func (o *TransactionBase) SetPendingTransactionIdNil() {
	o.PendingTransactionId.Set(nil)
}

// UnsetPendingTransactionId ensures that no value is present for PendingTransactionId, not even an explicit nil
func (o *TransactionBase) UnsetPendingTransactionId() {
	o.PendingTransactionId.Unset()
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionBase) GetCategoryId() string {
	if o == nil || o.CategoryId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CategoryId.Get()
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionBase) GetCategoryIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CategoryId.Get(), o.CategoryId.IsSet()
}

// HasCategoryId returns a boolean if a field has been set.
func (o *TransactionBase) HasCategoryId() bool {
	if o != nil && o.CategoryId.IsSet() {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given NullableString and assigns it to the CategoryId field.
func (o *TransactionBase) SetCategoryId(v string) {
	o.CategoryId.Set(&v)
}
// SetCategoryIdNil sets the value for CategoryId to be an explicit nil
func (o *TransactionBase) SetCategoryIdNil() {
	o.CategoryId.Set(nil)
}

// UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
func (o *TransactionBase) UnsetCategoryId() {
	o.CategoryId.Unset()
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionBase) GetCategory() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionBase) GetCategoryOk() (*[]string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return &o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *TransactionBase) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given []string and assigns it to the Category field.
func (o *TransactionBase) SetCategory(v []string) {
	o.Category = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *TransactionBase) GetLocation() Location {
	if o == nil || o.Location == nil {
		var ret Location
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetLocationOk() (*Location, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *TransactionBase) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given Location and assigns it to the Location field.
func (o *TransactionBase) SetLocation(v Location) {
	o.Location = &v
}

// GetPaymentMeta returns the PaymentMeta field value if set, zero value otherwise.
func (o *TransactionBase) GetPaymentMeta() PaymentMeta {
	if o == nil || o.PaymentMeta == nil {
		var ret PaymentMeta
		return ret
	}
	return *o.PaymentMeta
}

// GetPaymentMetaOk returns a tuple with the PaymentMeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetPaymentMetaOk() (*PaymentMeta, bool) {
	if o == nil || o.PaymentMeta == nil {
		return nil, false
	}
	return o.PaymentMeta, true
}

// HasPaymentMeta returns a boolean if a field has been set.
func (o *TransactionBase) HasPaymentMeta() bool {
	if o != nil && o.PaymentMeta != nil {
		return true
	}

	return false
}

// SetPaymentMeta gets a reference to the given PaymentMeta and assigns it to the PaymentMeta field.
func (o *TransactionBase) SetPaymentMeta(v PaymentMeta) {
	o.PaymentMeta = &v
}

// GetAccountOwner returns the AccountOwner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionBase) GetAccountOwner() string {
	if o == nil || o.AccountOwner.Get() == nil {
		var ret string
		return ret
	}
	return *o.AccountOwner.Get()
}

// GetAccountOwnerOk returns a tuple with the AccountOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionBase) GetAccountOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AccountOwner.Get(), o.AccountOwner.IsSet()
}

// HasAccountOwner returns a boolean if a field has been set.
func (o *TransactionBase) HasAccountOwner() bool {
	if o != nil && o.AccountOwner.IsSet() {
		return true
	}

	return false
}

// SetAccountOwner gets a reference to the given NullableString and assigns it to the AccountOwner field.
func (o *TransactionBase) SetAccountOwner(v string) {
	o.AccountOwner.Set(&v)
}
// SetAccountOwnerNil sets the value for AccountOwner to be an explicit nil
func (o *TransactionBase) SetAccountOwnerNil() {
	o.AccountOwner.Set(nil)
}

// UnsetAccountOwner ensures that no value is present for AccountOwner, not even an explicit nil
func (o *TransactionBase) UnsetAccountOwner() {
	o.AccountOwner.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransactionBase) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransactionBase) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransactionBase) SetName(v string) {
	o.Name = &v
}

// GetOriginalDescription returns the OriginalDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionBase) GetOriginalDescription() string {
	if o == nil || o.OriginalDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.OriginalDescription.Get()
}

// GetOriginalDescriptionOk returns a tuple with the OriginalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionBase) GetOriginalDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginalDescription.Get(), o.OriginalDescription.IsSet()
}

// HasOriginalDescription returns a boolean if a field has been set.
func (o *TransactionBase) HasOriginalDescription() bool {
	if o != nil && o.OriginalDescription.IsSet() {
		return true
	}

	return false
}

// SetOriginalDescription gets a reference to the given NullableString and assigns it to the OriginalDescription field.
func (o *TransactionBase) SetOriginalDescription(v string) {
	o.OriginalDescription.Set(&v)
}
// SetOriginalDescriptionNil sets the value for OriginalDescription to be an explicit nil
func (o *TransactionBase) SetOriginalDescriptionNil() {
	o.OriginalDescription.Set(nil)
}

// UnsetOriginalDescription ensures that no value is present for OriginalDescription, not even an explicit nil
func (o *TransactionBase) UnsetOriginalDescription() {
	o.OriginalDescription.Unset()
}

// GetAccountId returns the AccountId field value
func (o *TransactionBase) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransactionBase) SetAccountId(v string) {
	o.AccountId = v
}

// GetAmount returns the Amount field value
func (o *TransactionBase) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TransactionBase) SetAmount(v float32) {
	o.Amount = v
}

// GetIsoCurrencyCode returns the IsoCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionBase) GetIsoCurrencyCode() string {
	if o == nil || o.IsoCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.IsoCurrencyCode.Get()
}

// GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionBase) GetIsoCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsoCurrencyCode.Get(), o.IsoCurrencyCode.IsSet()
}

// SetIsoCurrencyCode sets field value
func (o *TransactionBase) SetIsoCurrencyCode(v string) {
	o.IsoCurrencyCode.Set(&v)
}

// GetUnofficialCurrencyCode returns the UnofficialCurrencyCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransactionBase) GetUnofficialCurrencyCode() string {
	if o == nil || o.UnofficialCurrencyCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.UnofficialCurrencyCode.Get()
}

// GetUnofficialCurrencyCodeOk returns a tuple with the UnofficialCurrencyCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionBase) GetUnofficialCurrencyCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnofficialCurrencyCode.Get(), o.UnofficialCurrencyCode.IsSet()
}

// SetUnofficialCurrencyCode sets field value
func (o *TransactionBase) SetUnofficialCurrencyCode(v string) {
	o.UnofficialCurrencyCode.Set(&v)
}

// GetDate returns the Date field value
func (o *TransactionBase) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *TransactionBase) SetDate(v string) {
	o.Date = v
}

// GetPending returns the Pending field value
func (o *TransactionBase) GetPending() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Pending
}

// GetPendingOk returns a tuple with the Pending field value
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetPendingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Pending, true
}

// SetPending sets field value
func (o *TransactionBase) SetPending(v bool) {
	o.Pending = v
}

// GetTransactionId returns the TransactionId field value
func (o *TransactionBase) GetTransactionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *TransactionBase) GetTransactionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *TransactionBase) SetTransactionId(v string) {
	o.TransactionId = v
}

func (o TransactionBase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransactionType != nil {
		toSerialize["transaction_type"] = o.TransactionType
	}
	if o.PendingTransactionId.IsSet() {
		toSerialize["pending_transaction_id"] = o.PendingTransactionId.Get()
	}
	if o.CategoryId.IsSet() {
		toSerialize["category_id"] = o.CategoryId.Get()
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.PaymentMeta != nil {
		toSerialize["payment_meta"] = o.PaymentMeta
	}
	if o.AccountOwner.IsSet() {
		toSerialize["account_owner"] = o.AccountOwner.Get()
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OriginalDescription.IsSet() {
		toSerialize["original_description"] = o.OriginalDescription.Get()
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["iso_currency_code"] = o.IsoCurrencyCode.Get()
	}
	if true {
		toSerialize["unofficial_currency_code"] = o.UnofficialCurrencyCode.Get()
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if true {
		toSerialize["pending"] = o.Pending
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionBase) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionBase := _TransactionBase{}

	if err = json.Unmarshal(bytes, &varTransactionBase); err == nil {
		*o = TransactionBase(varTransactionBase)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transaction_type")
		delete(additionalProperties, "pending_transaction_id")
		delete(additionalProperties, "category_id")
		delete(additionalProperties, "category")
		delete(additionalProperties, "location")
		delete(additionalProperties, "payment_meta")
		delete(additionalProperties, "account_owner")
		delete(additionalProperties, "name")
		delete(additionalProperties, "original_description")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "amount")
		delete(additionalProperties, "iso_currency_code")
		delete(additionalProperties, "unofficial_currency_code")
		delete(additionalProperties, "date")
		delete(additionalProperties, "pending")
		delete(additionalProperties, "transaction_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionBase struct {
	value *TransactionBase
	isSet bool
}

func (v NullableTransactionBase) Get() *TransactionBase {
	return v.value
}

func (v *NullableTransactionBase) Set(val *TransactionBase) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionBase) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionBase(val *TransactionBase) *NullableTransactionBase {
	return &NullableTransactionBase{value: val, isSet: true}
}

func (v NullableTransactionBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

