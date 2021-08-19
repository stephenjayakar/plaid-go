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

// InflowModel The `inflow_model` allows you to model a test account that receives regular income or make regular payments on a loan. Any transactions generated by the `inflow_model` will appear in addition to randomly generated test data or transactions specified by `override_accounts`.
type InflowModel struct {
	// Inflow model. One of the following:  `none`: No income `monthly-income`: Income occurs once per month `monthly-balance-payment`: Pays off the balance on a liability account at the given statement day of month `monthly-interest-only-payment`: Makes an interest-only payment on a liability account at the given statement day of month. Note that account types supported by Liabilities will accrue interest in the Sandbox. The types are account type `credit` with subtype `credit` or `paypal`, and account type `loan` with subtype `student` or `mortgage`.
	Type string `json:"type"`
	// Amount of income per month. This value is required if `type` is `monthly-income`.
	IncomeAmount float32 `json:"income_amount"`
	// Number between 1 and 28, or `last` meaning the last day of the month. The day of the month on which the income transaction will appear. The name of the income transaction. This field is required if `type` is `monthly-income`, `monthly-balance-payment` or `monthly-interest-only-payment`.
	PaymentDayOfMonth float32 `json:"payment_day_of_month"`
	// The name of the income transaction. This field is required if `type` is `monthly-income`, `monthly-balance-payment` or `monthly-interest-only-payment`.
	TransactionName string `json:"transaction_name"`
	// Number between 1 and 28, or `last` meaning the last day of the month. The day of the month on which the balance is calculated for the next payment. The name of the income transaction. This field is required if `type` is `monthly-balance-payment` or `monthly-interest-only-payment`.
	StatementDayOfMonth string `json:"statement_day_of_month"`
	AdditionalProperties map[string]interface{}
}

type _InflowModel InflowModel

// NewInflowModel instantiates a new InflowModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInflowModel(type_ string, incomeAmount float32, paymentDayOfMonth float32, transactionName string, statementDayOfMonth string) *InflowModel {
	this := InflowModel{}
	this.Type = type_
	this.IncomeAmount = incomeAmount
	this.PaymentDayOfMonth = paymentDayOfMonth
	this.TransactionName = transactionName
	this.StatementDayOfMonth = statementDayOfMonth
	return &this
}

// NewInflowModelWithDefaults instantiates a new InflowModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInflowModelWithDefaults() *InflowModel {
	this := InflowModel{}
	return &this
}

// GetType returns the Type field value
func (o *InflowModel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InflowModel) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InflowModel) SetType(v string) {
	o.Type = v
}

// GetIncomeAmount returns the IncomeAmount field value
func (o *InflowModel) GetIncomeAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.IncomeAmount
}

// GetIncomeAmountOk returns a tuple with the IncomeAmount field value
// and a boolean to check if the value has been set.
func (o *InflowModel) GetIncomeAmountOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IncomeAmount, true
}

// SetIncomeAmount sets field value
func (o *InflowModel) SetIncomeAmount(v float32) {
	o.IncomeAmount = v
}

// GetPaymentDayOfMonth returns the PaymentDayOfMonth field value
func (o *InflowModel) GetPaymentDayOfMonth() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PaymentDayOfMonth
}

// GetPaymentDayOfMonthOk returns a tuple with the PaymentDayOfMonth field value
// and a boolean to check if the value has been set.
func (o *InflowModel) GetPaymentDayOfMonthOk() (*float32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentDayOfMonth, true
}

// SetPaymentDayOfMonth sets field value
func (o *InflowModel) SetPaymentDayOfMonth(v float32) {
	o.PaymentDayOfMonth = v
}

// GetTransactionName returns the TransactionName field value
func (o *InflowModel) GetTransactionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionName
}

// GetTransactionNameOk returns a tuple with the TransactionName field value
// and a boolean to check if the value has been set.
func (o *InflowModel) GetTransactionNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionName, true
}

// SetTransactionName sets field value
func (o *InflowModel) SetTransactionName(v string) {
	o.TransactionName = v
}

// GetStatementDayOfMonth returns the StatementDayOfMonth field value
func (o *InflowModel) GetStatementDayOfMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatementDayOfMonth
}

// GetStatementDayOfMonthOk returns a tuple with the StatementDayOfMonth field value
// and a boolean to check if the value has been set.
func (o *InflowModel) GetStatementDayOfMonthOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StatementDayOfMonth, true
}

// SetStatementDayOfMonth sets field value
func (o *InflowModel) SetStatementDayOfMonth(v string) {
	o.StatementDayOfMonth = v
}

func (o InflowModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["income_amount"] = o.IncomeAmount
	}
	if true {
		toSerialize["payment_day_of_month"] = o.PaymentDayOfMonth
	}
	if true {
		toSerialize["transaction_name"] = o.TransactionName
	}
	if true {
		toSerialize["statement_day_of_month"] = o.StatementDayOfMonth
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *InflowModel) UnmarshalJSON(bytes []byte) (err error) {
	varInflowModel := _InflowModel{}

	if err = json.Unmarshal(bytes, &varInflowModel); err == nil {
		*o = InflowModel(varInflowModel)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "income_amount")
		delete(additionalProperties, "payment_day_of_month")
		delete(additionalProperties, "transaction_name")
		delete(additionalProperties, "statement_day_of_month")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInflowModel struct {
	value *InflowModel
	isSet bool
}

func (v NullableInflowModel) Get() *InflowModel {
	return v.value
}

func (v *NullableInflowModel) Set(val *InflowModel) {
	v.value = val
	v.isSet = true
}

func (v NullableInflowModel) IsSet() bool {
	return v.isSet
}

func (v *NullableInflowModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInflowModel(val *InflowModel) *NullableInflowModel {
	return &NullableInflowModel{value: val, isSet: true}
}

func (v NullableInflowModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInflowModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


