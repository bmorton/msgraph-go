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

// InlineObject597 struct for InlineObject597
type InlineObject597 struct {
	GroupId NullableString `json:"groupId,omitempty"`
	RenameAs NullableString `json:"renameAs,omitempty"`
	NotebookFolder NullableString `json:"notebookFolder,omitempty"`
	SiteCollectionId NullableString `json:"siteCollectionId,omitempty"`
	SiteId NullableString `json:"siteId,omitempty"`
}

// NewInlineObject597 instantiates a new InlineObject597 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject597() *InlineObject597 {
	this := InlineObject597{}
	return &this
}

// NewInlineObject597WithDefaults instantiates a new InlineObject597 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject597WithDefaults() *InlineObject597 {
	this := InlineObject597{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject597) GetGroupId() string {
	if o == nil || o.GroupId.Get() == nil {
		var ret string
		return ret
	}
	return *o.GroupId.Get()
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject597) GetGroupIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GroupId.Get(), o.GroupId.IsSet()
}

// HasGroupId returns a boolean if a field has been set.
func (o *InlineObject597) HasGroupId() bool {
	if o != nil && o.GroupId.IsSet() {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given NullableString and assigns it to the GroupId field.
func (o *InlineObject597) SetGroupId(v string) {
	o.GroupId.Set(&v)
}
// SetGroupIdNil sets the value for GroupId to be an explicit nil
func (o *InlineObject597) SetGroupIdNil() {
	o.GroupId.Set(nil)
}

// UnsetGroupId ensures that no value is present for GroupId, not even an explicit nil
func (o *InlineObject597) UnsetGroupId() {
	o.GroupId.Unset()
}

// GetRenameAs returns the RenameAs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject597) GetRenameAs() string {
	if o == nil || o.RenameAs.Get() == nil {
		var ret string
		return ret
	}
	return *o.RenameAs.Get()
}

// GetRenameAsOk returns a tuple with the RenameAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject597) GetRenameAsOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.RenameAs.Get(), o.RenameAs.IsSet()
}

// HasRenameAs returns a boolean if a field has been set.
func (o *InlineObject597) HasRenameAs() bool {
	if o != nil && o.RenameAs.IsSet() {
		return true
	}

	return false
}

// SetRenameAs gets a reference to the given NullableString and assigns it to the RenameAs field.
func (o *InlineObject597) SetRenameAs(v string) {
	o.RenameAs.Set(&v)
}
// SetRenameAsNil sets the value for RenameAs to be an explicit nil
func (o *InlineObject597) SetRenameAsNil() {
	o.RenameAs.Set(nil)
}

// UnsetRenameAs ensures that no value is present for RenameAs, not even an explicit nil
func (o *InlineObject597) UnsetRenameAs() {
	o.RenameAs.Unset()
}

// GetNotebookFolder returns the NotebookFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject597) GetNotebookFolder() string {
	if o == nil || o.NotebookFolder.Get() == nil {
		var ret string
		return ret
	}
	return *o.NotebookFolder.Get()
}

// GetNotebookFolderOk returns a tuple with the NotebookFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject597) GetNotebookFolderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.NotebookFolder.Get(), o.NotebookFolder.IsSet()
}

// HasNotebookFolder returns a boolean if a field has been set.
func (o *InlineObject597) HasNotebookFolder() bool {
	if o != nil && o.NotebookFolder.IsSet() {
		return true
	}

	return false
}

// SetNotebookFolder gets a reference to the given NullableString and assigns it to the NotebookFolder field.
func (o *InlineObject597) SetNotebookFolder(v string) {
	o.NotebookFolder.Set(&v)
}
// SetNotebookFolderNil sets the value for NotebookFolder to be an explicit nil
func (o *InlineObject597) SetNotebookFolderNil() {
	o.NotebookFolder.Set(nil)
}

// UnsetNotebookFolder ensures that no value is present for NotebookFolder, not even an explicit nil
func (o *InlineObject597) UnsetNotebookFolder() {
	o.NotebookFolder.Unset()
}

// GetSiteCollectionId returns the SiteCollectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject597) GetSiteCollectionId() string {
	if o == nil || o.SiteCollectionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteCollectionId.Get()
}

// GetSiteCollectionIdOk returns a tuple with the SiteCollectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject597) GetSiteCollectionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteCollectionId.Get(), o.SiteCollectionId.IsSet()
}

// HasSiteCollectionId returns a boolean if a field has been set.
func (o *InlineObject597) HasSiteCollectionId() bool {
	if o != nil && o.SiteCollectionId.IsSet() {
		return true
	}

	return false
}

// SetSiteCollectionId gets a reference to the given NullableString and assigns it to the SiteCollectionId field.
func (o *InlineObject597) SetSiteCollectionId(v string) {
	o.SiteCollectionId.Set(&v)
}
// SetSiteCollectionIdNil sets the value for SiteCollectionId to be an explicit nil
func (o *InlineObject597) SetSiteCollectionIdNil() {
	o.SiteCollectionId.Set(nil)
}

// UnsetSiteCollectionId ensures that no value is present for SiteCollectionId, not even an explicit nil
func (o *InlineObject597) UnsetSiteCollectionId() {
	o.SiteCollectionId.Unset()
}

// GetSiteId returns the SiteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject597) GetSiteId() string {
	if o == nil || o.SiteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteId.Get()
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject597) GetSiteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteId.Get(), o.SiteId.IsSet()
}

// HasSiteId returns a boolean if a field has been set.
func (o *InlineObject597) HasSiteId() bool {
	if o != nil && o.SiteId.IsSet() {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given NullableString and assigns it to the SiteId field.
func (o *InlineObject597) SetSiteId(v string) {
	o.SiteId.Set(&v)
}
// SetSiteIdNil sets the value for SiteId to be an explicit nil
func (o *InlineObject597) SetSiteIdNil() {
	o.SiteId.Set(nil)
}

// UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
func (o *InlineObject597) UnsetSiteId() {
	o.SiteId.Unset()
}

func (o InlineObject597) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GroupId.IsSet() {
		toSerialize["groupId"] = o.GroupId.Get()
	}
	if o.RenameAs.IsSet() {
		toSerialize["renameAs"] = o.RenameAs.Get()
	}
	if o.NotebookFolder.IsSet() {
		toSerialize["notebookFolder"] = o.NotebookFolder.Get()
	}
	if o.SiteCollectionId.IsSet() {
		toSerialize["siteCollectionId"] = o.SiteCollectionId.Get()
	}
	if o.SiteId.IsSet() {
		toSerialize["siteId"] = o.SiteId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject597 struct {
	value *InlineObject597
	isSet bool
}

func (v NullableInlineObject597) Get() *InlineObject597 {
	return v.value
}

func (v *NullableInlineObject597) Set(val *InlineObject597) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject597) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject597) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject597(val *InlineObject597) *NullableInlineObject597 {
	return &NullableInlineObject597{value: val, isSet: true}
}

func (v NullableInlineObject597) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject597) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


