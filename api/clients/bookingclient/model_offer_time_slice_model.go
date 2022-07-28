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
	"time"
)

// OfferTimeSliceModel struct for OfferTimeSliceModel
type OfferTimeSliceModel struct {
	// The start date and time for this time slice<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	From time.Time `json:"from"`
	// The end date and time for this time slice<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	To time.Time `json:"to"`
	// The number of available units for that time slice
	AvailableUnits int32 `json:"availableUnits"`
	BaseAmount AmountModel `json:"baseAmount"`
	TotalGrossAmount MonetaryValueModel `json:"totalGrossAmount"`
	// The breakdown for services included in the offer
	IncludedServices *[]OfferServiceModel `json:"includedServices,omitempty"`
}

// NewOfferTimeSliceModel instantiates a new OfferTimeSliceModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfferTimeSliceModel(from time.Time, to time.Time, availableUnits int32, baseAmount AmountModel, totalGrossAmount MonetaryValueModel) *OfferTimeSliceModel {
	this := OfferTimeSliceModel{}
	this.From = from
	this.To = to
	this.AvailableUnits = availableUnits
	this.BaseAmount = baseAmount
	this.TotalGrossAmount = totalGrossAmount
	return &this
}

// NewOfferTimeSliceModelWithDefaults instantiates a new OfferTimeSliceModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferTimeSliceModelWithDefaults() *OfferTimeSliceModel {
	this := OfferTimeSliceModel{}
	return &this
}

// GetFrom returns the From field value
func (o *OfferTimeSliceModel) GetFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *OfferTimeSliceModel) GetFromOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *OfferTimeSliceModel) SetFrom(v time.Time) {
	o.From = v
}

// GetTo returns the To field value
func (o *OfferTimeSliceModel) GetTo() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *OfferTimeSliceModel) GetToOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *OfferTimeSliceModel) SetTo(v time.Time) {
	o.To = v
}

// GetAvailableUnits returns the AvailableUnits field value
func (o *OfferTimeSliceModel) GetAvailableUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvailableUnits
}

// GetAvailableUnitsOk returns a tuple with the AvailableUnits field value
// and a boolean to check if the value has been set.
func (o *OfferTimeSliceModel) GetAvailableUnitsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableUnits, true
}

// SetAvailableUnits sets field value
func (o *OfferTimeSliceModel) SetAvailableUnits(v int32) {
	o.AvailableUnits = v
}

// GetBaseAmount returns the BaseAmount field value
func (o *OfferTimeSliceModel) GetBaseAmount() AmountModel {
	if o == nil {
		var ret AmountModel
		return ret
	}

	return o.BaseAmount
}

// GetBaseAmountOk returns a tuple with the BaseAmount field value
// and a boolean to check if the value has been set.
func (o *OfferTimeSliceModel) GetBaseAmountOk() (*AmountModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BaseAmount, true
}

// SetBaseAmount sets field value
func (o *OfferTimeSliceModel) SetBaseAmount(v AmountModel) {
	o.BaseAmount = v
}

// GetTotalGrossAmount returns the TotalGrossAmount field value
func (o *OfferTimeSliceModel) GetTotalGrossAmount() MonetaryValueModel {
	if o == nil {
		var ret MonetaryValueModel
		return ret
	}

	return o.TotalGrossAmount
}

// GetTotalGrossAmountOk returns a tuple with the TotalGrossAmount field value
// and a boolean to check if the value has been set.
func (o *OfferTimeSliceModel) GetTotalGrossAmountOk() (*MonetaryValueModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalGrossAmount, true
}

// SetTotalGrossAmount sets field value
func (o *OfferTimeSliceModel) SetTotalGrossAmount(v MonetaryValueModel) {
	o.TotalGrossAmount = v
}

// GetIncludedServices returns the IncludedServices field value if set, zero value otherwise.
func (o *OfferTimeSliceModel) GetIncludedServices() []OfferServiceModel {
	if o == nil || o.IncludedServices == nil {
		var ret []OfferServiceModel
		return ret
	}
	return *o.IncludedServices
}

// GetIncludedServicesOk returns a tuple with the IncludedServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfferTimeSliceModel) GetIncludedServicesOk() (*[]OfferServiceModel, bool) {
	if o == nil || o.IncludedServices == nil {
		return nil, false
	}
	return o.IncludedServices, true
}

// HasIncludedServices returns a boolean if a field has been set.
func (o *OfferTimeSliceModel) HasIncludedServices() bool {
	if o != nil && o.IncludedServices != nil {
		return true
	}

	return false
}

// SetIncludedServices gets a reference to the given []OfferServiceModel and assigns it to the IncludedServices field.
func (o *OfferTimeSliceModel) SetIncludedServices(v []OfferServiceModel) {
	o.IncludedServices = &v
}

func (o OfferTimeSliceModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["availableUnits"] = o.AvailableUnits
	}
	if true {
		toSerialize["baseAmount"] = o.BaseAmount
	}
	if true {
		toSerialize["totalGrossAmount"] = o.TotalGrossAmount
	}
	if o.IncludedServices != nil {
		toSerialize["includedServices"] = o.IncludedServices
	}
	return json.Marshal(toSerialize)
}

type NullableOfferTimeSliceModel struct {
	value *OfferTimeSliceModel
	isSet bool
}

func (v NullableOfferTimeSliceModel) Get() *OfferTimeSliceModel {
	return v.value
}

func (v *NullableOfferTimeSliceModel) Set(val *OfferTimeSliceModel) {
	v.value = val
	v.isSet = true
}

func (v NullableOfferTimeSliceModel) IsSet() bool {
	return v.isSet
}

func (v *NullableOfferTimeSliceModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfferTimeSliceModel(val *OfferTimeSliceModel) *NullableOfferTimeSliceModel {
	return &NullableOfferTimeSliceModel{value: val, isSet: true}
}

func (v NullableOfferTimeSliceModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfferTimeSliceModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

