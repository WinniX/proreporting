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

// DesiredStayDetailsModel struct for DesiredStayDetailsModel
type DesiredStayDetailsModel struct {
	// Date and optional time of arrival<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	Arrival string `json:"arrival"`
	// Date and optional time of departure. Cannot be more than 5 years after arrival.<br />Specify either a pure date or a date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	Departure string `json:"departure"`
	// Number of adults
	Adults int32 `json:"adults"`
	// Ages of the children
	ChildrenAges *[]int32 `json:"childrenAges,omitempty"`
	// Whether the prices for time slices with no change to the rate plan should be re-quoted based on current prices, or if  only additions like change of number of adults should be calculated. Defaults to 'false'.
	Requote *bool `json:"requote,omitempty"`
	// The list of time slices
	TimeSlices []DesiredTimeSliceModel `json:"timeSlices"`
}

// NewDesiredStayDetailsModel instantiates a new DesiredStayDetailsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesiredStayDetailsModel(arrival string, departure string, adults int32, timeSlices []DesiredTimeSliceModel) *DesiredStayDetailsModel {
	this := DesiredStayDetailsModel{}
	this.Arrival = arrival
	this.Departure = departure
	this.Adults = adults
	this.TimeSlices = timeSlices
	return &this
}

// NewDesiredStayDetailsModelWithDefaults instantiates a new DesiredStayDetailsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesiredStayDetailsModelWithDefaults() *DesiredStayDetailsModel {
	this := DesiredStayDetailsModel{}
	return &this
}

// GetArrival returns the Arrival field value
func (o *DesiredStayDetailsModel) GetArrival() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Arrival
}

// GetArrivalOk returns a tuple with the Arrival field value
// and a boolean to check if the value has been set.
func (o *DesiredStayDetailsModel) GetArrivalOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Arrival, true
}

// SetArrival sets field value
func (o *DesiredStayDetailsModel) SetArrival(v string) {
	o.Arrival = v
}

// GetDeparture returns the Departure field value
func (o *DesiredStayDetailsModel) GetDeparture() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Departure
}

// GetDepartureOk returns a tuple with the Departure field value
// and a boolean to check if the value has been set.
func (o *DesiredStayDetailsModel) GetDepartureOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Departure, true
}

// SetDeparture sets field value
func (o *DesiredStayDetailsModel) SetDeparture(v string) {
	o.Departure = v
}

// GetAdults returns the Adults field value
func (o *DesiredStayDetailsModel) GetAdults() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Adults
}

// GetAdultsOk returns a tuple with the Adults field value
// and a boolean to check if the value has been set.
func (o *DesiredStayDetailsModel) GetAdultsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Adults, true
}

// SetAdults sets field value
func (o *DesiredStayDetailsModel) SetAdults(v int32) {
	o.Adults = v
}

// GetChildrenAges returns the ChildrenAges field value if set, zero value otherwise.
func (o *DesiredStayDetailsModel) GetChildrenAges() []int32 {
	if o == nil || o.ChildrenAges == nil {
		var ret []int32
		return ret
	}
	return *o.ChildrenAges
}

// GetChildrenAgesOk returns a tuple with the ChildrenAges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesiredStayDetailsModel) GetChildrenAgesOk() (*[]int32, bool) {
	if o == nil || o.ChildrenAges == nil {
		return nil, false
	}
	return o.ChildrenAges, true
}

// HasChildrenAges returns a boolean if a field has been set.
func (o *DesiredStayDetailsModel) HasChildrenAges() bool {
	if o != nil && o.ChildrenAges != nil {
		return true
	}

	return false
}

// SetChildrenAges gets a reference to the given []int32 and assigns it to the ChildrenAges field.
func (o *DesiredStayDetailsModel) SetChildrenAges(v []int32) {
	o.ChildrenAges = &v
}

// GetRequote returns the Requote field value if set, zero value otherwise.
func (o *DesiredStayDetailsModel) GetRequote() bool {
	if o == nil || o.Requote == nil {
		var ret bool
		return ret
	}
	return *o.Requote
}

// GetRequoteOk returns a tuple with the Requote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesiredStayDetailsModel) GetRequoteOk() (*bool, bool) {
	if o == nil || o.Requote == nil {
		return nil, false
	}
	return o.Requote, true
}

// HasRequote returns a boolean if a field has been set.
func (o *DesiredStayDetailsModel) HasRequote() bool {
	if o != nil && o.Requote != nil {
		return true
	}

	return false
}

// SetRequote gets a reference to the given bool and assigns it to the Requote field.
func (o *DesiredStayDetailsModel) SetRequote(v bool) {
	o.Requote = &v
}

// GetTimeSlices returns the TimeSlices field value
func (o *DesiredStayDetailsModel) GetTimeSlices() []DesiredTimeSliceModel {
	if o == nil {
		var ret []DesiredTimeSliceModel
		return ret
	}

	return o.TimeSlices
}

// GetTimeSlicesOk returns a tuple with the TimeSlices field value
// and a boolean to check if the value has been set.
func (o *DesiredStayDetailsModel) GetTimeSlicesOk() (*[]DesiredTimeSliceModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeSlices, true
}

// SetTimeSlices sets field value
func (o *DesiredStayDetailsModel) SetTimeSlices(v []DesiredTimeSliceModel) {
	o.TimeSlices = v
}

func (o DesiredStayDetailsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["arrival"] = o.Arrival
	}
	if true {
		toSerialize["departure"] = o.Departure
	}
	if true {
		toSerialize["adults"] = o.Adults
	}
	if o.ChildrenAges != nil {
		toSerialize["childrenAges"] = o.ChildrenAges
	}
	if o.Requote != nil {
		toSerialize["requote"] = o.Requote
	}
	if true {
		toSerialize["timeSlices"] = o.TimeSlices
	}
	return json.Marshal(toSerialize)
}

type NullableDesiredStayDetailsModel struct {
	value *DesiredStayDetailsModel
	isSet bool
}

func (v NullableDesiredStayDetailsModel) Get() *DesiredStayDetailsModel {
	return v.value
}

func (v *NullableDesiredStayDetailsModel) Set(val *DesiredStayDetailsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDesiredStayDetailsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDesiredStayDetailsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesiredStayDetailsModel(val *DesiredStayDetailsModel) *NullableDesiredStayDetailsModel {
	return &NullableDesiredStayDetailsModel{value: val, isSet: true}
}

func (v NullableDesiredStayDetailsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesiredStayDetailsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

