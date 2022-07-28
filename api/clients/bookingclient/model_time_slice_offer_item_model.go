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

// TimeSliceOfferItemModel struct for TimeSliceOfferItemModel
type TimeSliceOfferItemModel struct {
	UnitGroup EmbeddedUnitGroupModel `json:"unitGroup"`
	// The minimum guarantee type for the offer
	MinGuaranteeType *string `json:"minGuaranteeType,omitempty"`
	MinAdvance *PeriodModel `json:"minAdvance,omitempty"`
	MaxAdvance *PeriodModel `json:"maxAdvance,omitempty"`
	// The number of available units for the offer
	Available int32 `json:"available"`
	Restrictions *RateRestrictionsModel `json:"restrictions,omitempty"`
	// The prices for this offer
	Prices *[]PerOccupancyPriceItemModel `json:"prices,omitempty"`
}

// NewTimeSliceOfferItemModel instantiates a new TimeSliceOfferItemModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeSliceOfferItemModel(unitGroup EmbeddedUnitGroupModel, available int32) *TimeSliceOfferItemModel {
	this := TimeSliceOfferItemModel{}
	this.UnitGroup = unitGroup
	this.Available = available
	return &this
}

// NewTimeSliceOfferItemModelWithDefaults instantiates a new TimeSliceOfferItemModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeSliceOfferItemModelWithDefaults() *TimeSliceOfferItemModel {
	this := TimeSliceOfferItemModel{}
	return &this
}

// GetUnitGroup returns the UnitGroup field value
func (o *TimeSliceOfferItemModel) GetUnitGroup() EmbeddedUnitGroupModel {
	if o == nil {
		var ret EmbeddedUnitGroupModel
		return ret
	}

	return o.UnitGroup
}

// GetUnitGroupOk returns a tuple with the UnitGroup field value
// and a boolean to check if the value has been set.
func (o *TimeSliceOfferItemModel) GetUnitGroupOk() (*EmbeddedUnitGroupModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UnitGroup, true
}

// SetUnitGroup sets field value
func (o *TimeSliceOfferItemModel) SetUnitGroup(v EmbeddedUnitGroupModel) {
	o.UnitGroup = v
}

// GetMinGuaranteeType returns the MinGuaranteeType field value if set, zero value otherwise.
func (o *TimeSliceOfferItemModel) GetMinGuaranteeType() string {
	if o == nil || o.MinGuaranteeType == nil {
		var ret string
		return ret
	}
	return *o.MinGuaranteeType
}

// GetMinGuaranteeTypeOk returns a tuple with the MinGuaranteeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSliceOfferItemModel) GetMinGuaranteeTypeOk() (*string, bool) {
	if o == nil || o.MinGuaranteeType == nil {
		return nil, false
	}
	return o.MinGuaranteeType, true
}

// HasMinGuaranteeType returns a boolean if a field has been set.
func (o *TimeSliceOfferItemModel) HasMinGuaranteeType() bool {
	if o != nil && o.MinGuaranteeType != nil {
		return true
	}

	return false
}

// SetMinGuaranteeType gets a reference to the given string and assigns it to the MinGuaranteeType field.
func (o *TimeSliceOfferItemModel) SetMinGuaranteeType(v string) {
	o.MinGuaranteeType = &v
}

// GetMinAdvance returns the MinAdvance field value if set, zero value otherwise.
func (o *TimeSliceOfferItemModel) GetMinAdvance() PeriodModel {
	if o == nil || o.MinAdvance == nil {
		var ret PeriodModel
		return ret
	}
	return *o.MinAdvance
}

// GetMinAdvanceOk returns a tuple with the MinAdvance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSliceOfferItemModel) GetMinAdvanceOk() (*PeriodModel, bool) {
	if o == nil || o.MinAdvance == nil {
		return nil, false
	}
	return o.MinAdvance, true
}

// HasMinAdvance returns a boolean if a field has been set.
func (o *TimeSliceOfferItemModel) HasMinAdvance() bool {
	if o != nil && o.MinAdvance != nil {
		return true
	}

	return false
}

// SetMinAdvance gets a reference to the given PeriodModel and assigns it to the MinAdvance field.
func (o *TimeSliceOfferItemModel) SetMinAdvance(v PeriodModel) {
	o.MinAdvance = &v
}

// GetMaxAdvance returns the MaxAdvance field value if set, zero value otherwise.
func (o *TimeSliceOfferItemModel) GetMaxAdvance() PeriodModel {
	if o == nil || o.MaxAdvance == nil {
		var ret PeriodModel
		return ret
	}
	return *o.MaxAdvance
}

// GetMaxAdvanceOk returns a tuple with the MaxAdvance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSliceOfferItemModel) GetMaxAdvanceOk() (*PeriodModel, bool) {
	if o == nil || o.MaxAdvance == nil {
		return nil, false
	}
	return o.MaxAdvance, true
}

// HasMaxAdvance returns a boolean if a field has been set.
func (o *TimeSliceOfferItemModel) HasMaxAdvance() bool {
	if o != nil && o.MaxAdvance != nil {
		return true
	}

	return false
}

// SetMaxAdvance gets a reference to the given PeriodModel and assigns it to the MaxAdvance field.
func (o *TimeSliceOfferItemModel) SetMaxAdvance(v PeriodModel) {
	o.MaxAdvance = &v
}

// GetAvailable returns the Available field value
func (o *TimeSliceOfferItemModel) GetAvailable() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value
// and a boolean to check if the value has been set.
func (o *TimeSliceOfferItemModel) GetAvailableOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Available, true
}

// SetAvailable sets field value
func (o *TimeSliceOfferItemModel) SetAvailable(v int32) {
	o.Available = v
}

// GetRestrictions returns the Restrictions field value if set, zero value otherwise.
func (o *TimeSliceOfferItemModel) GetRestrictions() RateRestrictionsModel {
	if o == nil || o.Restrictions == nil {
		var ret RateRestrictionsModel
		return ret
	}
	return *o.Restrictions
}

// GetRestrictionsOk returns a tuple with the Restrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSliceOfferItemModel) GetRestrictionsOk() (*RateRestrictionsModel, bool) {
	if o == nil || o.Restrictions == nil {
		return nil, false
	}
	return o.Restrictions, true
}

// HasRestrictions returns a boolean if a field has been set.
func (o *TimeSliceOfferItemModel) HasRestrictions() bool {
	if o != nil && o.Restrictions != nil {
		return true
	}

	return false
}

// SetRestrictions gets a reference to the given RateRestrictionsModel and assigns it to the Restrictions field.
func (o *TimeSliceOfferItemModel) SetRestrictions(v RateRestrictionsModel) {
	o.Restrictions = &v
}

// GetPrices returns the Prices field value if set, zero value otherwise.
func (o *TimeSliceOfferItemModel) GetPrices() []PerOccupancyPriceItemModel {
	if o == nil || o.Prices == nil {
		var ret []PerOccupancyPriceItemModel
		return ret
	}
	return *o.Prices
}

// GetPricesOk returns a tuple with the Prices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeSliceOfferItemModel) GetPricesOk() (*[]PerOccupancyPriceItemModel, bool) {
	if o == nil || o.Prices == nil {
		return nil, false
	}
	return o.Prices, true
}

// HasPrices returns a boolean if a field has been set.
func (o *TimeSliceOfferItemModel) HasPrices() bool {
	if o != nil && o.Prices != nil {
		return true
	}

	return false
}

// SetPrices gets a reference to the given []PerOccupancyPriceItemModel and assigns it to the Prices field.
func (o *TimeSliceOfferItemModel) SetPrices(v []PerOccupancyPriceItemModel) {
	o.Prices = &v
}

func (o TimeSliceOfferItemModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["unitGroup"] = o.UnitGroup
	}
	if o.MinGuaranteeType != nil {
		toSerialize["minGuaranteeType"] = o.MinGuaranteeType
	}
	if o.MinAdvance != nil {
		toSerialize["minAdvance"] = o.MinAdvance
	}
	if o.MaxAdvance != nil {
		toSerialize["maxAdvance"] = o.MaxAdvance
	}
	if true {
		toSerialize["available"] = o.Available
	}
	if o.Restrictions != nil {
		toSerialize["restrictions"] = o.Restrictions
	}
	if o.Prices != nil {
		toSerialize["prices"] = o.Prices
	}
	return json.Marshal(toSerialize)
}

type NullableTimeSliceOfferItemModel struct {
	value *TimeSliceOfferItemModel
	isSet bool
}

func (v NullableTimeSliceOfferItemModel) Get() *TimeSliceOfferItemModel {
	return v.value
}

func (v *NullableTimeSliceOfferItemModel) Set(val *TimeSliceOfferItemModel) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeSliceOfferItemModel) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeSliceOfferItemModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeSliceOfferItemModel(val *TimeSliceOfferItemModel) *NullableTimeSliceOfferItemModel {
	return &NullableTimeSliceOfferItemModel{value: val, isSet: true}
}

func (v NullableTimeSliceOfferItemModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeSliceOfferItemModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


