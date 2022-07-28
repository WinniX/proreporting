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

// GroupListModel struct for GroupListModel
type GroupListModel struct {
	Groups []GroupItemModel `json:"groups"`
	// Total count of items
	Count int64 `json:"count"`
}

// NewGroupListModel instantiates a new GroupListModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupListModel(groups []GroupItemModel, count int64) *GroupListModel {
	this := GroupListModel{}
	this.Groups = groups
	this.Count = count
	return &this
}

// NewGroupListModelWithDefaults instantiates a new GroupListModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupListModelWithDefaults() *GroupListModel {
	this := GroupListModel{}
	return &this
}

// GetGroups returns the Groups field value
func (o *GroupListModel) GetGroups() []GroupItemModel {
	if o == nil {
		var ret []GroupItemModel
		return ret
	}

	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value
// and a boolean to check if the value has been set.
func (o *GroupListModel) GetGroupsOk() (*[]GroupItemModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Groups, true
}

// SetGroups sets field value
func (o *GroupListModel) SetGroups(v []GroupItemModel) {
	o.Groups = v
}

// GetCount returns the Count field value
func (o *GroupListModel) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *GroupListModel) GetCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *GroupListModel) SetCount(v int64) {
	o.Count = v
}

func (o GroupListModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["groups"] = o.Groups
	}
	if true {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableGroupListModel struct {
	value *GroupListModel
	isSet bool
}

func (v NullableGroupListModel) Get() *GroupListModel {
	return v.value
}

func (v *NullableGroupListModel) Set(val *GroupListModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupListModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupListModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupListModel(val *GroupListModel) *NullableGroupListModel {
	return &NullableGroupListModel{value: val, isSet: true}
}

func (v NullableGroupListModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupListModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

