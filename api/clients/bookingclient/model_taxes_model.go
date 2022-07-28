/*
 * apaleo Booking API
 *
 * Resources and methods to manage guest journeys.
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bookingclient

import (
	"encoding/json"
)

// TaxesModel struct for TaxesModel
type TaxesModel struct {
	// The amount of taxes, which are VAT or Sales Taxes depending on the country of the property
	Tax float64 `json:"tax"`
	// The amount of City Tax including VAT
	CityTax float64 `json:"cityTax"`
}

// NewTaxesModel instantiates a new TaxesModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaxesModel(tax float64, cityTax float64) *TaxesModel {
	this := TaxesModel{}
	this.Tax = tax
	this.CityTax = cityTax
	return &this
}

// NewTaxesModelWithDefaults instantiates a new TaxesModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaxesModelWithDefaults() *TaxesModel {
	this := TaxesModel{}
	return &this
}

// GetTax returns the Tax field value
func (o *TaxesModel) GetTax() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Tax
}

// GetTaxOk returns a tuple with the Tax field value
// and a boolean to check if the value has been set.
func (o *TaxesModel) GetTaxOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Tax, true
}

// SetTax sets field value
func (o *TaxesModel) SetTax(v float64) {
	o.Tax = v
}

// GetCityTax returns the CityTax field value
func (o *TaxesModel) GetCityTax() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.CityTax
}

// GetCityTaxOk returns a tuple with the CityTax field value
// and a boolean to check if the value has been set.
func (o *TaxesModel) GetCityTaxOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CityTax, true
}

// SetCityTax sets field value
func (o *TaxesModel) SetCityTax(v float64) {
	o.CityTax = v
}

func (o TaxesModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tax"] = o.Tax
	}
	if true {
		toSerialize["cityTax"] = o.CityTax
	}
	return json.Marshal(toSerialize)
}

type NullableTaxesModel struct {
	value *TaxesModel
	isSet bool
}

func (v NullableTaxesModel) Get() *TaxesModel {
	return v.value
}

func (v *NullableTaxesModel) Set(val *TaxesModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTaxesModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTaxesModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaxesModel(val *TaxesModel) *NullableTaxesModel {
	return &NullableTaxesModel{value: val, isSet: true}
}

func (v NullableTaxesModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaxesModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


