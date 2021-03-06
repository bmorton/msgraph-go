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

// MicrosoftGraphContactFolder struct for MicrosoftGraphContactFolder
type MicrosoftGraphContactFolder struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The folder's display name.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The ID of the folder's parent folder.
	ParentFolderId NullableString `json:"parentFolderId,omitempty"`
	// The collection of child folders in the folder. Navigation property. Read-only. Nullable.
	ChildFolders *[]MicrosoftGraphContactFolder `json:"childFolders,omitempty"`
	// The contacts in the folder. Navigation property. Read-only. Nullable.
	Contacts *[]MicrosoftGraphContact `json:"contacts,omitempty"`
	// The collection of multi-value extended properties defined for the contactFolder. Read-only. Nullable.
	MultiValueExtendedProperties *[]MicrosoftGraphMultiValueLegacyExtendedProperty `json:"multiValueExtendedProperties,omitempty"`
	// The collection of single-value extended properties defined for the contactFolder. Read-only. Nullable.
	SingleValueExtendedProperties *[]MicrosoftGraphSingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}

// NewMicrosoftGraphContactFolder instantiates a new MicrosoftGraphContactFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphContactFolder() *MicrosoftGraphContactFolder {
	this := MicrosoftGraphContactFolder{}
	return &this
}

// NewMicrosoftGraphContactFolderWithDefaults instantiates a new MicrosoftGraphContactFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphContactFolderWithDefaults() *MicrosoftGraphContactFolder {
	this := MicrosoftGraphContactFolder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphContactFolder) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphContactFolder) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphContactFolder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphContactFolder) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphContactFolder) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphContactFolder) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphContactFolder) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphContactFolder) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphContactFolder) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphContactFolder) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetParentFolderId returns the ParentFolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphContactFolder) GetParentFolderId() string {
	if o == nil || o.ParentFolderId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentFolderId.Get()
}

// GetParentFolderIdOk returns a tuple with the ParentFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphContactFolder) GetParentFolderIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentFolderId.Get(), o.ParentFolderId.IsSet()
}

// HasParentFolderId returns a boolean if a field has been set.
func (o *MicrosoftGraphContactFolder) HasParentFolderId() bool {
	if o != nil && o.ParentFolderId.IsSet() {
		return true
	}

	return false
}

// SetParentFolderId gets a reference to the given NullableString and assigns it to the ParentFolderId field.
func (o *MicrosoftGraphContactFolder) SetParentFolderId(v string) {
	o.ParentFolderId.Set(&v)
}
// SetParentFolderIdNil sets the value for ParentFolderId to be an explicit nil
func (o *MicrosoftGraphContactFolder) SetParentFolderIdNil() {
	o.ParentFolderId.Set(nil)
}

// UnsetParentFolderId ensures that no value is present for ParentFolderId, not even an explicit nil
func (o *MicrosoftGraphContactFolder) UnsetParentFolderId() {
	o.ParentFolderId.Unset()
}

// GetChildFolders returns the ChildFolders field value if set, zero value otherwise.
func (o *MicrosoftGraphContactFolder) GetChildFolders() []MicrosoftGraphContactFolder {
	if o == nil || o.ChildFolders == nil {
		var ret []MicrosoftGraphContactFolder
		return ret
	}
	return *o.ChildFolders
}

// GetChildFoldersOk returns a tuple with the ChildFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphContactFolder) GetChildFoldersOk() (*[]MicrosoftGraphContactFolder, bool) {
	if o == nil || o.ChildFolders == nil {
		return nil, false
	}
	return o.ChildFolders, true
}

// HasChildFolders returns a boolean if a field has been set.
func (o *MicrosoftGraphContactFolder) HasChildFolders() bool {
	if o != nil && o.ChildFolders != nil {
		return true
	}

	return false
}

// SetChildFolders gets a reference to the given []MicrosoftGraphContactFolder and assigns it to the ChildFolders field.
func (o *MicrosoftGraphContactFolder) SetChildFolders(v []MicrosoftGraphContactFolder) {
	o.ChildFolders = &v
}

// GetContacts returns the Contacts field value if set, zero value otherwise.
func (o *MicrosoftGraphContactFolder) GetContacts() []MicrosoftGraphContact {
	if o == nil || o.Contacts == nil {
		var ret []MicrosoftGraphContact
		return ret
	}
	return *o.Contacts
}

// GetContactsOk returns a tuple with the Contacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphContactFolder) GetContactsOk() (*[]MicrosoftGraphContact, bool) {
	if o == nil || o.Contacts == nil {
		return nil, false
	}
	return o.Contacts, true
}

// HasContacts returns a boolean if a field has been set.
func (o *MicrosoftGraphContactFolder) HasContacts() bool {
	if o != nil && o.Contacts != nil {
		return true
	}

	return false
}

// SetContacts gets a reference to the given []MicrosoftGraphContact and assigns it to the Contacts field.
func (o *MicrosoftGraphContactFolder) SetContacts(v []MicrosoftGraphContact) {
	o.Contacts = &v
}

// GetMultiValueExtendedProperties returns the MultiValueExtendedProperties field value if set, zero value otherwise.
func (o *MicrosoftGraphContactFolder) GetMultiValueExtendedProperties() []MicrosoftGraphMultiValueLegacyExtendedProperty {
	if o == nil || o.MultiValueExtendedProperties == nil {
		var ret []MicrosoftGraphMultiValueLegacyExtendedProperty
		return ret
	}
	return *o.MultiValueExtendedProperties
}

// GetMultiValueExtendedPropertiesOk returns a tuple with the MultiValueExtendedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphContactFolder) GetMultiValueExtendedPropertiesOk() (*[]MicrosoftGraphMultiValueLegacyExtendedProperty, bool) {
	if o == nil || o.MultiValueExtendedProperties == nil {
		return nil, false
	}
	return o.MultiValueExtendedProperties, true
}

// HasMultiValueExtendedProperties returns a boolean if a field has been set.
func (o *MicrosoftGraphContactFolder) HasMultiValueExtendedProperties() bool {
	if o != nil && o.MultiValueExtendedProperties != nil {
		return true
	}

	return false
}

// SetMultiValueExtendedProperties gets a reference to the given []MicrosoftGraphMultiValueLegacyExtendedProperty and assigns it to the MultiValueExtendedProperties field.
func (o *MicrosoftGraphContactFolder) SetMultiValueExtendedProperties(v []MicrosoftGraphMultiValueLegacyExtendedProperty) {
	o.MultiValueExtendedProperties = &v
}

// GetSingleValueExtendedProperties returns the SingleValueExtendedProperties field value if set, zero value otherwise.
func (o *MicrosoftGraphContactFolder) GetSingleValueExtendedProperties() []MicrosoftGraphSingleValueLegacyExtendedProperty {
	if o == nil || o.SingleValueExtendedProperties == nil {
		var ret []MicrosoftGraphSingleValueLegacyExtendedProperty
		return ret
	}
	return *o.SingleValueExtendedProperties
}

// GetSingleValueExtendedPropertiesOk returns a tuple with the SingleValueExtendedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphContactFolder) GetSingleValueExtendedPropertiesOk() (*[]MicrosoftGraphSingleValueLegacyExtendedProperty, bool) {
	if o == nil || o.SingleValueExtendedProperties == nil {
		return nil, false
	}
	return o.SingleValueExtendedProperties, true
}

// HasSingleValueExtendedProperties returns a boolean if a field has been set.
func (o *MicrosoftGraphContactFolder) HasSingleValueExtendedProperties() bool {
	if o != nil && o.SingleValueExtendedProperties != nil {
		return true
	}

	return false
}

// SetSingleValueExtendedProperties gets a reference to the given []MicrosoftGraphSingleValueLegacyExtendedProperty and assigns it to the SingleValueExtendedProperties field.
func (o *MicrosoftGraphContactFolder) SetSingleValueExtendedProperties(v []MicrosoftGraphSingleValueLegacyExtendedProperty) {
	o.SingleValueExtendedProperties = &v
}

func (o MicrosoftGraphContactFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.ParentFolderId.IsSet() {
		toSerialize["parentFolderId"] = o.ParentFolderId.Get()
	}
	if o.ChildFolders != nil {
		toSerialize["childFolders"] = o.ChildFolders
	}
	if o.Contacts != nil {
		toSerialize["contacts"] = o.Contacts
	}
	if o.MultiValueExtendedProperties != nil {
		toSerialize["multiValueExtendedProperties"] = o.MultiValueExtendedProperties
	}
	if o.SingleValueExtendedProperties != nil {
		toSerialize["singleValueExtendedProperties"] = o.SingleValueExtendedProperties
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphContactFolder struct {
	value *MicrosoftGraphContactFolder
	isSet bool
}

func (v NullableMicrosoftGraphContactFolder) Get() *MicrosoftGraphContactFolder {
	return v.value
}

func (v *NullableMicrosoftGraphContactFolder) Set(val *MicrosoftGraphContactFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphContactFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphContactFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphContactFolder(val *MicrosoftGraphContactFolder) *NullableMicrosoftGraphContactFolder {
	return &NullableMicrosoftGraphContactFolder{value: val, isSet: true}
}

func (v NullableMicrosoftGraphContactFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphContactFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


