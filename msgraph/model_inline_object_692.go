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

// InlineObject692 struct for InlineObject692
type InlineObject692 struct {
	Id NullableString `json:"id,omitempty"`
	GroupId NullableString `json:"groupId,omitempty"`
	RenameAs NullableString `json:"renameAs,omitempty"`
	SiteCollectionId NullableString `json:"siteCollectionId,omitempty"`
	SiteId NullableString `json:"siteId,omitempty"`
}

// NewInlineObject692 instantiates a new InlineObject692 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject692() *InlineObject692 {
	this := InlineObject692{}
	return &this
}

// NewInlineObject692WithDefaults instantiates a new InlineObject692 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject692WithDefaults() *InlineObject692 {
	this := InlineObject692{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject692) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject692) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *InlineObject692) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *InlineObject692) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *InlineObject692) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *InlineObject692) UnsetId() {
	o.Id.Unset()
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject692) GetGroupId() string {
	if o == nil || o.GroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject692) GetGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *InlineObject692) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableString and assigns it to the GroupId field.
func (o *InlineObject692) SetGroupId(v string) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *InlineObject692) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *InlineObject692) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetRenameAs returns the RenameAs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject692) GetRenameAs() string {
	if o == nil || o.RenameAs.Get() == nil {
		var ret string
		return ret
	}
	return *o.RenameAs.Get()
}

// GetRenameAsOk returns a tuple with the RenameAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject692) GetRenameAsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenameAs.Get(), o.RenameAs.IsSet()
}

// HasRenameAs returns a boolean if a field has been set.
func (o *InlineObject692) HasRenameAs() bool {
	if o != nil && o.RenameAs.IsSet() {
		return true
	}

	return false
}

// SetRenameAs gets a reference to the given NullableString and assigns it to the RenameAs field.
func (o *InlineObject692) SetRenameAs(v string) {
	o.RenameAs.Set(&v)
}
// SetRenameAsNil sets the value for RenameAs to be an explicit nil
func (o *InlineObject692) SetRenameAsNil() {
	o.RenameAs.Set(nil)
}

// UnsetRenameAs ensures that no value is present for RenameAs, not even an explicit nil
func (o *InlineObject692) UnsetRenameAs() {
	o.RenameAs.Unset()
}

// GetSiteCollectionId returns the SiteCollectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject692) GetSiteCollectionId() string {
	if o == nil || o.SiteCollectionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteCollectionId.Get()
}

// GetSiteCollectionIdOk returns a tuple with the SiteCollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject692) GetSiteCollectionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteCollectionId.Get(), o.SiteCollectionId.IsSet()
}

// HasSiteCollectionId returns a boolean if a field has been set.
func (o *InlineObject692) HasSiteCollectionId() bool {
	if o != nil && o.SiteCollectionId.IsSet() {
		return true
	}

	return false
}

// SetSiteCollectionId gets a reference to the given NullableString and assigns it to the SiteCollectionId field.
func (o *InlineObject692) SetSiteCollectionId(v string) {
	o.SiteCollectionId.Set(&v)
}
// SetSiteCollectionIdNil sets the value for SiteCollectionId to be an explicit nil
func (o *InlineObject692) SetSiteCollectionIdNil() {
	o.SiteCollectionId.Set(nil)
}

// UnsetSiteCollectionId ensures that no value is present for SiteCollectionId, not even an explicit nil
func (o *InlineObject692) UnsetSiteCollectionId() {
	o.SiteCollectionId.Unset()
}

// GetSiteId returns the SiteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject692) GetSiteId() string {
	if o == nil || o.SiteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteId.Get()
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject692) GetSiteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteId.Get(), o.SiteId.IsSet()
}

// HasSiteId returns a boolean if a field has been set.
func (o *InlineObject692) HasSiteId() bool {
	if o != nil && o.SiteId.IsSet() {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given NullableString and assigns it to the SiteId field.
func (o *InlineObject692) SetSiteId(v string) {
	o.SiteId.Set(&v)
}
// SetSiteIdNil sets the value for SiteId to be an explicit nil
func (o *InlineObject692) SetSiteIdNil() {
	o.SiteId.Set(nil)
}

// UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
func (o *InlineObject692) UnsetSiteId() {
	o.SiteId.Unset()
}

func (o InlineObject692) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.GroupId.IsSet() {
		toSerialize["groupId"] = o.GroupId.Get()
	}
	if o.RenameAs.IsSet() {
		toSerialize["renameAs"] = o.RenameAs.Get()
	}
	if o.SiteCollectionId.IsSet() {
		toSerialize["siteCollectionId"] = o.SiteCollectionId.Get()
	}
	if o.SiteId.IsSet() {
		toSerialize["siteId"] = o.SiteId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject692 struct {
	value *InlineObject692
	isSet bool
}

func (v NullableInlineObject692) Get() *InlineObject692 {
	return v.value
}

func (v *NullableInlineObject692) Set(val *InlineObject692) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject692) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject692) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject692(val *InlineObject692) *NullableInlineObject692 {
	return &NullableInlineObject692{value: val, isSet: true}
}

func (v NullableInlineObject692) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject692) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


