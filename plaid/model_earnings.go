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

// Earnings An object representing both a breakdown of earnings on a paystub and the total earnings.
type Earnings struct {
	Subtotals *[]EarningsTotal `json:"subtotals,omitempty"`
	Totals *[]EarningsTotal `json:"totals,omitempty"`
	Breakdown *[]EarningsBreakdown `json:"breakdown,omitempty"`
	Total *EarningsTotal `json:"total,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Earnings Earnings

// NewEarnings instantiates a new Earnings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarnings() *Earnings {
	this := Earnings{}
	return &this
}

// NewEarningsWithDefaults instantiates a new Earnings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarningsWithDefaults() *Earnings {
	this := Earnings{}
	return &this
}

// GetSubtotals returns the Subtotals field value if set, zero value otherwise.
func (o *Earnings) GetSubtotals() []EarningsTotal {
	if o == nil || o.Subtotals == nil {
		var ret []EarningsTotal
		return ret
	}
	return *o.Subtotals
}

// GetSubtotalsOk returns a tuple with the Subtotals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Earnings) GetSubtotalsOk() (*[]EarningsTotal, bool) {
	if o == nil || o.Subtotals == nil {
		return nil, false
	}
	return o.Subtotals, true
}

// HasSubtotals returns a boolean if a field has been set.
func (o *Earnings) HasSubtotals() bool {
	if o != nil && o.Subtotals != nil {
		return true
	}

	return false
}

// SetSubtotals gets a reference to the given []EarningsTotal and assigns it to the Subtotals field.
func (o *Earnings) SetSubtotals(v []EarningsTotal) {
	o.Subtotals = &v
}

// GetTotals returns the Totals field value if set, zero value otherwise.
func (o *Earnings) GetTotals() []EarningsTotal {
	if o == nil || o.Totals == nil {
		var ret []EarningsTotal
		return ret
	}
	return *o.Totals
}

// GetTotalsOk returns a tuple with the Totals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Earnings) GetTotalsOk() (*[]EarningsTotal, bool) {
	if o == nil || o.Totals == nil {
		return nil, false
	}
	return o.Totals, true
}

// HasTotals returns a boolean if a field has been set.
func (o *Earnings) HasTotals() bool {
	if o != nil && o.Totals != nil {
		return true
	}

	return false
}

// SetTotals gets a reference to the given []EarningsTotal and assigns it to the Totals field.
func (o *Earnings) SetTotals(v []EarningsTotal) {
	o.Totals = &v
}

// GetBreakdown returns the Breakdown field value if set, zero value otherwise.
func (o *Earnings) GetBreakdown() []EarningsBreakdown {
	if o == nil || o.Breakdown == nil {
		var ret []EarningsBreakdown
		return ret
	}
	return *o.Breakdown
}

// GetBreakdownOk returns a tuple with the Breakdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Earnings) GetBreakdownOk() (*[]EarningsBreakdown, bool) {
	if o == nil || o.Breakdown == nil {
		return nil, false
	}
	return o.Breakdown, true
}

// HasBreakdown returns a boolean if a field has been set.
func (o *Earnings) HasBreakdown() bool {
	if o != nil && o.Breakdown != nil {
		return true
	}

	return false
}

// SetBreakdown gets a reference to the given []EarningsBreakdown and assigns it to the Breakdown field.
func (o *Earnings) SetBreakdown(v []EarningsBreakdown) {
	o.Breakdown = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Earnings) GetTotal() EarningsTotal {
	if o == nil || o.Total == nil {
		var ret EarningsTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Earnings) GetTotalOk() (*EarningsTotal, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Earnings) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given EarningsTotal and assigns it to the Total field.
func (o *Earnings) SetTotal(v EarningsTotal) {
	o.Total = &v
}

func (o Earnings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Subtotals != nil {
		toSerialize["subtotals"] = o.Subtotals
	}
	if o.Totals != nil {
		toSerialize["totals"] = o.Totals
	}
	if o.Breakdown != nil {
		toSerialize["breakdown"] = o.Breakdown
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Earnings) UnmarshalJSON(bytes []byte) (err error) {
	varEarnings := _Earnings{}

	if err = json.Unmarshal(bytes, &varEarnings); err == nil {
		*o = Earnings(varEarnings)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "subtotals")
		delete(additionalProperties, "totals")
		delete(additionalProperties, "breakdown")
		delete(additionalProperties, "total")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEarnings struct {
	value *Earnings
	isSet bool
}

func (v NullableEarnings) Get() *Earnings {
	return v.value
}

func (v *NullableEarnings) Set(val *Earnings) {
	v.value = val
	v.isSet = true
}

func (v NullableEarnings) IsSet() bool {
	return v.isSet
}

func (v *NullableEarnings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarnings(val *Earnings) *NullableEarnings {
	return &NullableEarnings{value: val, isSet: true}
}

func (v NullableEarnings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarnings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

