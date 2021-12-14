/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// OutlookItem struct for OutlookItem
type OutlookItem struct {
	// The categories associated with the item
	Categories *[]*string `json:"categories,omitempty"`
	// Identifies the version of the item. Every time the item is changed, changeKey changes as well. This allows Exchange to apply changes to the correct version of the object. Read-only.
	ChangeKey NullableString `json:"changeKey,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
}

// NewOutlookItem instantiates a new OutlookItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutlookItem() *OutlookItem {
	this := OutlookItem{}
	return &this
}

// NewOutlookItemWithDefaults instantiates a new OutlookItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutlookItemWithDefaults() *OutlookItem {
	this := OutlookItem{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *OutlookItem) GetCategories() []*string {
	if o == nil || o.Categories == nil {
		var ret []*string
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutlookItem) GetCategoriesOk() (*[]*string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *OutlookItem) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []*string and assigns it to the Categories field.
func (o *OutlookItem) SetCategories(v []*string) {
	o.Categories = &v
}

// GetChangeKey returns the ChangeKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutlookItem) GetChangeKey() string {
	if o == nil || o.ChangeKey.Get() == nil {
		var ret string
		return ret
	}
	return *o.ChangeKey.Get()
}

// GetChangeKeyOk returns a tuple with the ChangeKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutlookItem) GetChangeKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ChangeKey.Get(), o.ChangeKey.IsSet()
}

// HasChangeKey returns a boolean if a field has been set.
func (o *OutlookItem) HasChangeKey() bool {
	if o != nil && o.ChangeKey.IsSet() {
		return true
	}

	return false
}

// SetChangeKey gets a reference to the given NullableString and assigns it to the ChangeKey field.
func (o *OutlookItem) SetChangeKey(v string) {
	o.ChangeKey.Set(&v)
}
// SetChangeKeyNil sets the value for ChangeKey to be an explicit nil
func (o *OutlookItem) SetChangeKeyNil() {
	o.ChangeKey.Set(nil)
}

// UnsetChangeKey ensures that no value is present for ChangeKey, not even an explicit nil
func (o *OutlookItem) UnsetChangeKey() {
	o.ChangeKey.Unset()
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutlookItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutlookItem) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *OutlookItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *OutlookItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *OutlookItem) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *OutlookItem) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OutlookItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OutlookItem) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *OutlookItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *OutlookItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *OutlookItem) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *OutlookItem) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

func (o OutlookItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if o.ChangeKey.IsSet() {
		toSerialize["changeKey"] = o.ChangeKey.Get()
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableOutlookItem struct {
	value *OutlookItem
	isSet bool
}

func (v NullableOutlookItem) Get() *OutlookItem {
	return v.value
}

func (v *NullableOutlookItem) Set(val *OutlookItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOutlookItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOutlookItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutlookItem(val *OutlookItem) *NullableOutlookItem {
	return &NullableOutlookItem{value: val, isSet: true}
}

func (v NullableOutlookItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutlookItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


