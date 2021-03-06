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

// MicrosoftGraphItemReference struct for MicrosoftGraphItemReference
type MicrosoftGraphItemReference struct {
	// Unique identifier of the drive instance that contains the item. Read-only.
	DriveId NullableString `json:"driveId,omitempty"`
	// Identifies the type of drive. See [drive][] resource for values.
	DriveType NullableString `json:"driveType,omitempty"`
	// Unique identifier of the item in the drive. Read-only.
	Id NullableString `json:"id,omitempty"`
	// The name of the item being referenced. Read-only.
	Name NullableString `json:"name,omitempty"`
	// Path that can be used to navigate to the item. Read-only.
	Path NullableString `json:"path,omitempty"`
	// A unique identifier for a shared resource that can be accessed via the [Shares][] API.
	ShareId NullableString `json:"shareId,omitempty"`
	// Returns identifiers useful for SharePoint REST compatibility. Read-only.
	SharepointIds AnyOfmicrosoftGraphSharepointIds `json:"sharepointIds,omitempty"`
	// For OneDrive for Business and SharePoint, this property represents the ID of the site that contains the parent document library of the driveItem resource. The value is the same as the id property of that [site][] resource. It is an opaque string that consists of three identifiers of the site. For OneDrive, this property is not populated.
	SiteId NullableString `json:"siteId,omitempty"`
}

// NewMicrosoftGraphItemReference instantiates a new MicrosoftGraphItemReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphItemReference() *MicrosoftGraphItemReference {
	this := MicrosoftGraphItemReference{}
	return &this
}

// NewMicrosoftGraphItemReferenceWithDefaults instantiates a new MicrosoftGraphItemReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphItemReferenceWithDefaults() *MicrosoftGraphItemReference {
	this := MicrosoftGraphItemReference{}
	return &this
}

// GetDriveId returns the DriveId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemReference) GetDriveId() string {
	if o == nil || o.DriveId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DriveId.Get()
}

// GetDriveIdOk returns a tuple with the DriveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemReference) GetDriveIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DriveId.Get(), o.DriveId.IsSet()
}

// HasDriveId returns a boolean if a field has been set.
func (o *MicrosoftGraphItemReference) HasDriveId() bool {
	if o != nil && o.DriveId.IsSet() {
		return true
	}

	return false
}

// SetDriveId gets a reference to the given NullableString and assigns it to the DriveId field.
func (o *MicrosoftGraphItemReference) SetDriveId(v string) {
	o.DriveId.Set(&v)
}
// SetDriveIdNil sets the value for DriveId to be an explicit nil
func (o *MicrosoftGraphItemReference) SetDriveIdNil() {
	o.DriveId.Set(nil)
}

// UnsetDriveId ensures that no value is present for DriveId, not even an explicit nil
func (o *MicrosoftGraphItemReference) UnsetDriveId() {
	o.DriveId.Unset()
}

// GetDriveType returns the DriveType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemReference) GetDriveType() string {
	if o == nil || o.DriveType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DriveType.Get()
}

// GetDriveTypeOk returns a tuple with the DriveType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemReference) GetDriveTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DriveType.Get(), o.DriveType.IsSet()
}

// HasDriveType returns a boolean if a field has been set.
func (o *MicrosoftGraphItemReference) HasDriveType() bool {
	if o != nil && o.DriveType.IsSet() {
		return true
	}

	return false
}

// SetDriveType gets a reference to the given NullableString and assigns it to the DriveType field.
func (o *MicrosoftGraphItemReference) SetDriveType(v string) {
	o.DriveType.Set(&v)
}
// SetDriveTypeNil sets the value for DriveType to be an explicit nil
func (o *MicrosoftGraphItemReference) SetDriveTypeNil() {
	o.DriveType.Set(nil)
}

// UnsetDriveType ensures that no value is present for DriveType, not even an explicit nil
func (o *MicrosoftGraphItemReference) UnsetDriveType() {
	o.DriveType.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemReference) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemReference) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphItemReference) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphItemReference) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphItemReference) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphItemReference) UnsetId() {
	o.Id.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemReference) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemReference) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphItemReference) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphItemReference) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphItemReference) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphItemReference) UnsetName() {
	o.Name.Unset()
}

// GetPath returns the Path field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemReference) GetPath() string {
	if o == nil || o.Path.Get() == nil {
		var ret string
		return ret
	}
	return *o.Path.Get()
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemReference) GetPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Path.Get(), o.Path.IsSet()
}

// HasPath returns a boolean if a field has been set.
func (o *MicrosoftGraphItemReference) HasPath() bool {
	if o != nil && o.Path.IsSet() {
		return true
	}

	return false
}

// SetPath gets a reference to the given NullableString and assigns it to the Path field.
func (o *MicrosoftGraphItemReference) SetPath(v string) {
	o.Path.Set(&v)
}
// SetPathNil sets the value for Path to be an explicit nil
func (o *MicrosoftGraphItemReference) SetPathNil() {
	o.Path.Set(nil)
}

// UnsetPath ensures that no value is present for Path, not even an explicit nil
func (o *MicrosoftGraphItemReference) UnsetPath() {
	o.Path.Unset()
}

// GetShareId returns the ShareId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemReference) GetShareId() string {
	if o == nil || o.ShareId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShareId.Get()
}

// GetShareIdOk returns a tuple with the ShareId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemReference) GetShareIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShareId.Get(), o.ShareId.IsSet()
}

// HasShareId returns a boolean if a field has been set.
func (o *MicrosoftGraphItemReference) HasShareId() bool {
	if o != nil && o.ShareId.IsSet() {
		return true
	}

	return false
}

// SetShareId gets a reference to the given NullableString and assigns it to the ShareId field.
func (o *MicrosoftGraphItemReference) SetShareId(v string) {
	o.ShareId.Set(&v)
}
// SetShareIdNil sets the value for ShareId to be an explicit nil
func (o *MicrosoftGraphItemReference) SetShareIdNil() {
	o.ShareId.Set(nil)
}

// UnsetShareId ensures that no value is present for ShareId, not even an explicit nil
func (o *MicrosoftGraphItemReference) UnsetShareId() {
	o.ShareId.Unset()
}

// GetSharepointIds returns the SharepointIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemReference) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds {
	if o == nil  {
		var ret AnyOfmicrosoftGraphSharepointIds
		return ret
	}
	return o.SharepointIds
}

// GetSharepointIdsOk returns a tuple with the SharepointIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemReference) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool) {
	if o == nil || o.SharepointIds == nil {
		return nil, false
	}
	return &o.SharepointIds, true
}

// HasSharepointIds returns a boolean if a field has been set.
func (o *MicrosoftGraphItemReference) HasSharepointIds() bool {
	if o != nil && o.SharepointIds != nil {
		return true
	}

	return false
}

// SetSharepointIds gets a reference to the given AnyOfmicrosoftGraphSharepointIds and assigns it to the SharepointIds field.
func (o *MicrosoftGraphItemReference) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds) {
	o.SharepointIds = v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphItemReference) GetSiteId() string {
	if o == nil || o.SiteId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SiteId.Get()
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphItemReference) GetSiteIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SiteId.Get(), o.SiteId.IsSet()
}

// HasSiteId returns a boolean if a field has been set.
func (o *MicrosoftGraphItemReference) HasSiteId() bool {
	if o != nil && o.SiteId.IsSet() {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given NullableString and assigns it to the SiteId field.
func (o *MicrosoftGraphItemReference) SetSiteId(v string) {
	o.SiteId.Set(&v)
}
// SetSiteIdNil sets the value for SiteId to be an explicit nil
func (o *MicrosoftGraphItemReference) SetSiteIdNil() {
	o.SiteId.Set(nil)
}

// UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
func (o *MicrosoftGraphItemReference) UnsetSiteId() {
	o.SiteId.Unset()
}

func (o MicrosoftGraphItemReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DriveId.IsSet() {
		toSerialize["driveId"] = o.DriveId.Get()
	}
	if o.DriveType.IsSet() {
		toSerialize["driveType"] = o.DriveType.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Path.IsSet() {
		toSerialize["path"] = o.Path.Get()
	}
	if o.ShareId.IsSet() {
		toSerialize["shareId"] = o.ShareId.Get()
	}
	if o.SharepointIds != nil {
		toSerialize["sharepointIds"] = o.SharepointIds
	}
	if o.SiteId.IsSet() {
		toSerialize["siteId"] = o.SiteId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphItemReference struct {
	value *MicrosoftGraphItemReference
	isSet bool
}

func (v NullableMicrosoftGraphItemReference) Get() *MicrosoftGraphItemReference {
	return v.value
}

func (v *NullableMicrosoftGraphItemReference) Set(val *MicrosoftGraphItemReference) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphItemReference) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphItemReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphItemReference(val *MicrosoftGraphItemReference) *NullableMicrosoftGraphItemReference {
	return &NullableMicrosoftGraphItemReference{value: val, isSet: true}
}

func (v NullableMicrosoftGraphItemReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphItemReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


