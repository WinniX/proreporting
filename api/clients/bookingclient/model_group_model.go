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

// GroupModel struct for GroupModel
type GroupModel struct {
	// Group id
	Id string `json:"id"`
	// Name of the group
	Name string `json:"name"`
	// Start date and time of the earliest block for this group<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	From *time.Time `json:"from,omitempty"`
	// End date and time of the latest block for this group<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	To *time.Time `json:"to,omitempty"`
	Booker *BookerModel `json:"booker,omitempty"`
	// Additional information and comments
	Comment *string `json:"comment,omitempty"`
	// Additional information and comment by the booker
	BookerComment *string `json:"bookerComment,omitempty"`
	PaymentAccount *PaymentAccountModel `json:"paymentAccount,omitempty"`
	// Date of creation<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	Created time.Time `json:"created"`
	// Date of last modification<br />A date and time (without fractional second part) in UTC or with UTC offset as defined in <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO8601:2004</a>
	Modified time.Time `json:"modified"`
	// Blocks within this group
	Blocks *[]GroupBlockModel `json:"blocks,omitempty"`
	// The list of actions for this group
	Actions *[]ActionModelGroupActionNotAllowedGroupActionReason `json:"actions,omitempty"`
	// The list of property ids this group belongs to
	PropertyIds []string `json:"propertyIds"`
}

// NewGroupModel instantiates a new GroupModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupModel(id string, name string, created time.Time, modified time.Time, propertyIds []string) *GroupModel {
	this := GroupModel{}
	this.Id = id
	this.Name = name
	this.Created = created
	this.Modified = modified
	this.PropertyIds = propertyIds
	return &this
}

// NewGroupModelWithDefaults instantiates a new GroupModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupModelWithDefaults() *GroupModel {
	this := GroupModel{}
	return &this
}

// GetId returns the Id field value
func (o *GroupModel) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GroupModel) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GroupModel) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *GroupModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GroupModel) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GroupModel) SetName(v string) {
	o.Name = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *GroupModel) GetFrom() time.Time {
	if o == nil || o.From == nil {
		var ret time.Time
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupModel) GetFromOk() (*time.Time, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *GroupModel) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given time.Time and assigns it to the From field.
func (o *GroupModel) SetFrom(v time.Time) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *GroupModel) GetTo() time.Time {
	if o == nil || o.To == nil {
		var ret time.Time
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupModel) GetToOk() (*time.Time, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *GroupModel) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given time.Time and assigns it to the To field.
func (o *GroupModel) SetTo(v time.Time) {
	o.To = &v
}

// GetBooker returns the Booker field value if set, zero value otherwise.
func (o *GroupModel) GetBooker() BookerModel {
	if o == nil || o.Booker == nil {
		var ret BookerModel
		return ret
	}
	return *o.Booker
}

// GetBookerOk returns a tuple with the Booker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupModel) GetBookerOk() (*BookerModel, bool) {
	if o == nil || o.Booker == nil {
		return nil, false
	}
	return o.Booker, true
}

// HasBooker returns a boolean if a field has been set.
func (o *GroupModel) HasBooker() bool {
	if o != nil && o.Booker != nil {
		return true
	}

	return false
}

// SetBooker gets a reference to the given BookerModel and assigns it to the Booker field.
func (o *GroupModel) SetBooker(v BookerModel) {
	o.Booker = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *GroupModel) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupModel) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *GroupModel) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *GroupModel) SetComment(v string) {
	o.Comment = &v
}

// GetBookerComment returns the BookerComment field value if set, zero value otherwise.
func (o *GroupModel) GetBookerComment() string {
	if o == nil || o.BookerComment == nil {
		var ret string
		return ret
	}
	return *o.BookerComment
}

// GetBookerCommentOk returns a tuple with the BookerComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupModel) GetBookerCommentOk() (*string, bool) {
	if o == nil || o.BookerComment == nil {
		return nil, false
	}
	return o.BookerComment, true
}

// HasBookerComment returns a boolean if a field has been set.
func (o *GroupModel) HasBookerComment() bool {
	if o != nil && o.BookerComment != nil {
		return true
	}

	return false
}

// SetBookerComment gets a reference to the given string and assigns it to the BookerComment field.
func (o *GroupModel) SetBookerComment(v string) {
	o.BookerComment = &v
}

// GetPaymentAccount returns the PaymentAccount field value if set, zero value otherwise.
func (o *GroupModel) GetPaymentAccount() PaymentAccountModel {
	if o == nil || o.PaymentAccount == nil {
		var ret PaymentAccountModel
		return ret
	}
	return *o.PaymentAccount
}

// GetPaymentAccountOk returns a tuple with the PaymentAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupModel) GetPaymentAccountOk() (*PaymentAccountModel, bool) {
	if o == nil || o.PaymentAccount == nil {
		return nil, false
	}
	return o.PaymentAccount, true
}

// HasPaymentAccount returns a boolean if a field has been set.
func (o *GroupModel) HasPaymentAccount() bool {
	if o != nil && o.PaymentAccount != nil {
		return true
	}

	return false
}

// SetPaymentAccount gets a reference to the given PaymentAccountModel and assigns it to the PaymentAccount field.
func (o *GroupModel) SetPaymentAccount(v PaymentAccountModel) {
	o.PaymentAccount = &v
}

// GetCreated returns the Created field value
func (o *GroupModel) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GroupModel) GetCreatedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GroupModel) SetCreated(v time.Time) {
	o.Created = v
}

// GetModified returns the Modified field value
func (o *GroupModel) GetModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value
// and a boolean to check if the value has been set.
func (o *GroupModel) GetModifiedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Modified, true
}

// SetModified sets field value
func (o *GroupModel) SetModified(v time.Time) {
	o.Modified = v
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *GroupModel) GetBlocks() []GroupBlockModel {
	if o == nil || o.Blocks == nil {
		var ret []GroupBlockModel
		return ret
	}
	return *o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupModel) GetBlocksOk() (*[]GroupBlockModel, bool) {
	if o == nil || o.Blocks == nil {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *GroupModel) HasBlocks() bool {
	if o != nil && o.Blocks != nil {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given []GroupBlockModel and assigns it to the Blocks field.
func (o *GroupModel) SetBlocks(v []GroupBlockModel) {
	o.Blocks = &v
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *GroupModel) GetActions() []ActionModelGroupActionNotAllowedGroupActionReason {
	if o == nil || o.Actions == nil {
		var ret []ActionModelGroupActionNotAllowedGroupActionReason
		return ret
	}
	return *o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupModel) GetActionsOk() (*[]ActionModelGroupActionNotAllowedGroupActionReason, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *GroupModel) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given []ActionModelGroupActionNotAllowedGroupActionReason and assigns it to the Actions field.
func (o *GroupModel) SetActions(v []ActionModelGroupActionNotAllowedGroupActionReason) {
	o.Actions = &v
}

// GetPropertyIds returns the PropertyIds field value
func (o *GroupModel) GetPropertyIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PropertyIds
}

// GetPropertyIdsOk returns a tuple with the PropertyIds field value
// and a boolean to check if the value has been set.
func (o *GroupModel) GetPropertyIdsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PropertyIds, true
}

// SetPropertyIds sets field value
func (o *GroupModel) SetPropertyIds(v []string) {
	o.PropertyIds = v
}

func (o GroupModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.From != nil {
		toSerialize["from"] = o.From
	}
	if o.To != nil {
		toSerialize["to"] = o.To
	}
	if o.Booker != nil {
		toSerialize["booker"] = o.Booker
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.BookerComment != nil {
		toSerialize["bookerComment"] = o.BookerComment
	}
	if o.PaymentAccount != nil {
		toSerialize["paymentAccount"] = o.PaymentAccount
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["modified"] = o.Modified
	}
	if o.Blocks != nil {
		toSerialize["blocks"] = o.Blocks
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if true {
		toSerialize["propertyIds"] = o.PropertyIds
	}
	return json.Marshal(toSerialize)
}

type NullableGroupModel struct {
	value *GroupModel
	isSet bool
}

func (v NullableGroupModel) Get() *GroupModel {
	return v.value
}

func (v *NullableGroupModel) Set(val *GroupModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupModel(val *GroupModel) *NullableGroupModel {
	return &NullableGroupModel{value: val, isSet: true}
}

func (v NullableGroupModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


