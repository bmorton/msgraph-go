/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
)

// InlineObject638 struct for InlineObject638
type InlineObject638 struct {
	Id NullableString `json:"id,omitempty"`
	GroupId NullableString `json:"groupId,omitempty"`
	SiteCollectionId NullableString `json:"siteCollectionId,omitempty"`
	SiteId NullableString `json:"siteId,omitempty"`
}

// NewInlineObject638 instantiates a new InlineObject638 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject638() *InlineObject638 {
	this := InlineObject638{}
	return &this
}

// NewInlineObject638WithDefaults instantiates a new InlineObject638 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject638WithDefaults() *InlineObject638 {
	this := InlineObject638{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject638) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject638) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject638) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *InlineObject638) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InlineObject638) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InlineObject638) UnsetId() {
	o.Id.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject638) GetGroupId() string {
	if o == nil || o.GroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject638) GetGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *InlineObject638) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableString and assigns it to the GroupId field.
func (o *InlineObject638) SetGroupId(v string) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *InlineObject638) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *InlineObject638) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetSiteCollectionId returns the SiteCollectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject638) GetSiteCollectionId() string {
	if o == nil || o.SiteCollectionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteCollectionId.Get()
}

// GetSiteCollectionIdOk returns a tuple with the SiteCollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject638) GetSiteCollectionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteCollectionId.Get(), o.SiteCollectionId.IsSet()
}

// HasSiteCollectionId returns a boolean if a field has been set.
func (o *InlineObject638) HasSiteCollectionId() bool {
	if o != nil && o.SiteCollectionId.IsSet() {
		return true
	}

	return false
}

// SetSiteCollectionId gets a reference to the given NullableString and assigns it to the SiteCollectionId field.
func (o *InlineObject638) SetSiteCollectionId(v string) {
	o.SiteCollectionId.Set(&v)
}
// SetSiteCollectionIdNil sets the value for SiteCollectionId to be an explicit nil
func (o *InlineObject638) SetSiteCollectionIdNil() {
	o.SiteCollectionId.Set(nil)
}

// UnsetSiteCollectionId ensures that no value is present for SiteCollectionId, not even an explicit nil
func (o *InlineObject638) UnsetSiteCollectionId() {
	o.SiteCollectionId.Unset()
}

// GetSiteId returns the SiteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject638) GetSiteId() string {
	if o == nil || o.SiteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteId.Get()
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject638) GetSiteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteId.Get(), o.SiteId.IsSet()
}

// HasSiteId returns a boolean if a field has been set.
func (o *InlineObject638) HasSiteId() bool {
	if o != nil && o.SiteId.IsSet() {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given NullableString and assigns it to the SiteId field.
func (o *InlineObject638) SetSiteId(v string) {
	o.SiteId.Set(&v)
}
// SetSiteIdNil sets the value for SiteId to be an explicit nil
func (o *InlineObject638) SetSiteIdNil() {
	o.SiteId.Set(nil)
}

// UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
func (o *InlineObject638) UnsetSiteId() {
	o.SiteId.Unset()
}

func (o InlineObject638) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.GroupId.IsSet() {
		toSerialize["groupId"] = o.GroupId.Get()
	}
	if o.SiteCollectionId.IsSet() {
		toSerialize["siteCollectionId"] = o.SiteCollectionId.Get()
	}
	if o.SiteId.IsSet() {
		toSerialize["siteId"] = o.SiteId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject638 struct {
	value *InlineObject638
	isSet bool
}

func (v NullableInlineObject638) Get() *InlineObject638 {
	return v.value
}

func (v *NullableInlineObject638) Set(val *InlineObject638) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject638) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject638) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject638(val *InlineObject638) *NullableInlineObject638 {
	return &NullableInlineObject638{value: val, isSet: true}
}

func (v NullableInlineObject638) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject638) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


