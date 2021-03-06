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

// MicrosoftGraphBaseItem struct for MicrosoftGraphBaseItem
type MicrosoftGraphBaseItem struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Identity of the user, device, or application which created the item. Read-only.
	CreatedBy AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	// Date and time of item creation. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Provides a user-visible description of the item. Optional.
	Description NullableString `json:"description,omitempty"`
	// ETag for the item. Read-only.
	ETag NullableString `json:"eTag,omitempty"`
	// Identity of the user, device, and application which last modified the item. Read-only.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// Date and time the item was last modified. Read-only.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The name of the item. Read-write.
	Name NullableString `json:"name,omitempty"`
	// Parent information, if the item has a parent. Read-write.
	ParentReference AnyOfmicrosoftGraphItemReference `json:"parentReference,omitempty"`
	// URL that displays the resource in the browser. Read-only.
	WebUrl NullableString `json:"webUrl,omitempty"`
	// Identity of the user who created the item. Read-only.
	CreatedByUser AnyOfmicrosoftGraphUser `json:"createdByUser,omitempty"`
	// Identity of the user who last modified the item. Read-only.
	LastModifiedByUser AnyOfmicrosoftGraphUser `json:"lastModifiedByUser,omitempty"`
}

// NewMicrosoftGraphBaseItem instantiates a new MicrosoftGraphBaseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphBaseItem() *MicrosoftGraphBaseItem {
	this := MicrosoftGraphBaseItem{}
	return &this
}

// NewMicrosoftGraphBaseItemWithDefaults instantiates a new MicrosoftGraphBaseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphBaseItemWithDefaults() *MicrosoftGraphBaseItem {
	this := MicrosoftGraphBaseItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphBaseItem) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphBaseItem) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphBaseItem) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphBaseItem) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphBaseItem) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphBaseItem) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphBaseItem) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphBaseItem) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphBaseItem) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphBaseItem) UnsetDescription() {
	o.Description.Unset()
}

// GetETag returns the ETag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetETag() string {
	if o == nil || o.ETag.Get() == nil {
		var ret string
		return ret
	}
	return *o.ETag.Get()
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetETagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ETag.Get(), o.ETag.IsSet()
}

// HasETag returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasETag() bool {
	if o != nil && o.ETag.IsSet() {
		return true
	}

	return false
}

// SetETag gets a reference to the given NullableString and assigns it to the ETag field.
func (o *MicrosoftGraphBaseItem) SetETag(v string) {
	o.ETag.Set(&v)
}
// SetETagNil sets the value for ETag to be an explicit nil
func (o *MicrosoftGraphBaseItem) SetETagNil() {
	o.ETag.Set(nil)
}

// UnsetETag ensures that no value is present for ETag, not even an explicit nil
func (o *MicrosoftGraphBaseItem) UnsetETag() {
	o.ETag.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphBaseItem) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphBaseItem) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphBaseItem) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphBaseItem) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphBaseItem) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphBaseItem) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphBaseItem) UnsetName() {
	o.Name.Unset()
}

// GetParentReference returns the ParentReference field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetParentReference() AnyOfmicrosoftGraphItemReference {
	if o == nil  {
		var ret AnyOfmicrosoftGraphItemReference
		return ret
	}
	return o.ParentReference
}

// GetParentReferenceOk returns a tuple with the ParentReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetParentReferenceOk() (*AnyOfmicrosoftGraphItemReference, bool) {
	if o == nil || o.ParentReference == nil {
		return nil, false
	}
	return &o.ParentReference, true
}

// HasParentReference returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasParentReference() bool {
	if o != nil && o.ParentReference != nil {
		return true
	}

	return false
}

// SetParentReference gets a reference to the given AnyOfmicrosoftGraphItemReference and assigns it to the ParentReference field.
func (o *MicrosoftGraphBaseItem) SetParentReference(v AnyOfmicrosoftGraphItemReference) {
	o.ParentReference = v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetWebUrl() string {
	if o == nil || o.WebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasWebUrl() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *MicrosoftGraphBaseItem) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *MicrosoftGraphBaseItem) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *MicrosoftGraphBaseItem) UnsetWebUrl() {
	o.WebUrl.Unset()
}

// GetCreatedByUser returns the CreatedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetCreatedByUser() AnyOfmicrosoftGraphUser {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return o.CreatedByUser
}

// GetCreatedByUserOk returns a tuple with the CreatedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetCreatedByUserOk() (*AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.CreatedByUser == nil {
		return nil, false
	}
	return &o.CreatedByUser, true
}

// HasCreatedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasCreatedByUser() bool {
	if o != nil && o.CreatedByUser != nil {
		return true
	}

	return false
}

// SetCreatedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the CreatedByUser field.
func (o *MicrosoftGraphBaseItem) SetCreatedByUser(v AnyOfmicrosoftGraphUser) {
	o.CreatedByUser = v
}

// GetLastModifiedByUser returns the LastModifiedByUser field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBaseItem) GetLastModifiedByUser() AnyOfmicrosoftGraphUser {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUser
		return ret
	}
	return o.LastModifiedByUser
}

// GetLastModifiedByUserOk returns a tuple with the LastModifiedByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBaseItem) GetLastModifiedByUserOk() (*AnyOfmicrosoftGraphUser, bool) {
	if o == nil || o.LastModifiedByUser == nil {
		return nil, false
	}
	return &o.LastModifiedByUser, true
}

// HasLastModifiedByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphBaseItem) HasLastModifiedByUser() bool {
	if o != nil && o.LastModifiedByUser != nil {
		return true
	}

	return false
}

// SetLastModifiedByUser gets a reference to the given AnyOfmicrosoftGraphUser and assigns it to the LastModifiedByUser field.
func (o *MicrosoftGraphBaseItem) SetLastModifiedByUser(v AnyOfmicrosoftGraphUser) {
	o.LastModifiedByUser = v
}

func (o MicrosoftGraphBaseItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.ETag.IsSet() {
		toSerialize["eTag"] = o.ETag.Get()
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.ParentReference != nil {
		toSerialize["parentReference"] = o.ParentReference
	}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	if o.CreatedByUser != nil {
		toSerialize["createdByUser"] = o.CreatedByUser
	}
	if o.LastModifiedByUser != nil {
		toSerialize["lastModifiedByUser"] = o.LastModifiedByUser
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphBaseItem struct {
	value *MicrosoftGraphBaseItem
	isSet bool
}

func (v NullableMicrosoftGraphBaseItem) Get() *MicrosoftGraphBaseItem {
	return v.value
}

func (v *NullableMicrosoftGraphBaseItem) Set(val *MicrosoftGraphBaseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphBaseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphBaseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphBaseItem(val *MicrosoftGraphBaseItem) *NullableMicrosoftGraphBaseItem {
	return &NullableMicrosoftGraphBaseItem{value: val, isSet: true}
}

func (v NullableMicrosoftGraphBaseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphBaseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


