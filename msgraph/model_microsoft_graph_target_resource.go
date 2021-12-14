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

// MicrosoftGraphTargetResource struct for MicrosoftGraphTargetResource
type MicrosoftGraphTargetResource struct {
	// Indicates the visible name defined for the resource. Typically specified when the resource is created.
	DisplayName NullableString `json:"displayName,omitempty"`
	// When type is set to Group, this indicates the group type. Possible values are: unifiedGroups, azureAD, and unknownFutureValue
	GroupType AnyOfmicrosoftGraphGroupType `json:"groupType,omitempty"`
	// Indicates the unique ID of the resource.
	Id NullableString `json:"id,omitempty"`
	// Indicates name, old value and new value of each attribute that changed. Property values depend on the operation type.
	ModifiedProperties *[]*AnyOfmicrosoftGraphModifiedProperty `json:"modifiedProperties,omitempty"`
	// Describes the resource type.  Example values include Application, Group, ServicePrincipal, and User.
	Type NullableString `json:"type,omitempty"`
	// When type is set to User, this includes the user name that initiated the action; null for other types.
	UserPrincipalName NullableString `json:"userPrincipalName,omitempty"`
}

// NewMicrosoftGraphTargetResource instantiates a new MicrosoftGraphTargetResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTargetResource() *MicrosoftGraphTargetResource {
	this := MicrosoftGraphTargetResource{}
	return &this
}

// NewMicrosoftGraphTargetResourceWithDefaults instantiates a new MicrosoftGraphTargetResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTargetResourceWithDefaults() *MicrosoftGraphTargetResource {
	this := MicrosoftGraphTargetResource{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTargetResource) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTargetResource) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphTargetResource) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphTargetResource) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphTargetResource) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphTargetResource) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetGroupType returns the GroupType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTargetResource) GetGroupType() AnyOfmicrosoftGraphGroupType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphGroupType
		return ret
	}
	return o.GroupType
}

// GetGroupTypeOk returns a tuple with the GroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTargetResource) GetGroupTypeOk() (*AnyOfmicrosoftGraphGroupType, bool) {
	if o == nil || o.GroupType == nil {
		return nil, false
	}
	return &o.GroupType, true
}

// HasGroupType returns a boolean if a field has been set.
func (o *MicrosoftGraphTargetResource) HasGroupType() bool {
	if o != nil && o.GroupType != nil {
		return true
	}

	return false
}

// SetGroupType gets a reference to the given AnyOfmicrosoftGraphGroupType and assigns it to the GroupType field.
func (o *MicrosoftGraphTargetResource) SetGroupType(v AnyOfmicrosoftGraphGroupType) {
	o.GroupType = v
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTargetResource) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTargetResource) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphTargetResource) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphTargetResource) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphTargetResource) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphTargetResource) UnsetId() {
	o.Id.Unset()
}

// GetModifiedProperties returns the ModifiedProperties field value if set, zero value otherwise.
func (o *MicrosoftGraphTargetResource) GetModifiedProperties() []*AnyOfmicrosoftGraphModifiedProperty {
	if o == nil || o.ModifiedProperties == nil {
		var ret []*AnyOfmicrosoftGraphModifiedProperty
		return ret
	}
	return *o.ModifiedProperties
}

// GetModifiedPropertiesOk returns a tuple with the ModifiedProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTargetResource) GetModifiedPropertiesOk() (*[]*AnyOfmicrosoftGraphModifiedProperty, bool) {
	if o == nil || o.ModifiedProperties == nil {
		return nil, false
	}
	return o.ModifiedProperties, true
}

// HasModifiedProperties returns a boolean if a field has been set.
func (o *MicrosoftGraphTargetResource) HasModifiedProperties() bool {
	if o != nil && o.ModifiedProperties != nil {
		return true
	}

	return false
}

// SetModifiedProperties gets a reference to the given []*AnyOfmicrosoftGraphModifiedProperty and assigns it to the ModifiedProperties field.
func (o *MicrosoftGraphTargetResource) SetModifiedProperties(v []*AnyOfmicrosoftGraphModifiedProperty) {
	o.ModifiedProperties = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTargetResource) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTargetResource) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphTargetResource) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *MicrosoftGraphTargetResource) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *MicrosoftGraphTargetResource) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *MicrosoftGraphTargetResource) UnsetType() {
	o.Type.Unset()
}

// GetUserPrincipalName returns the UserPrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTargetResource) GetUserPrincipalName() string {
	if o == nil || o.UserPrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserPrincipalName.Get()
}

// GetUserPrincipalNameOk returns a tuple with the UserPrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTargetResource) GetUserPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserPrincipalName.Get(), o.UserPrincipalName.IsSet()
}

// HasUserPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphTargetResource) HasUserPrincipalName() bool {
	if o != nil && o.UserPrincipalName.IsSet() {
		return true
	}

	return false
}

// SetUserPrincipalName gets a reference to the given NullableString and assigns it to the UserPrincipalName field.
func (o *MicrosoftGraphTargetResource) SetUserPrincipalName(v string) {
	o.UserPrincipalName.Set(&v)
}
// SetUserPrincipalNameNil sets the value for UserPrincipalName to be an explicit nil
func (o *MicrosoftGraphTargetResource) SetUserPrincipalNameNil() {
	o.UserPrincipalName.Set(nil)
}

// UnsetUserPrincipalName ensures that no value is present for UserPrincipalName, not even an explicit nil
func (o *MicrosoftGraphTargetResource) UnsetUserPrincipalName() {
	o.UserPrincipalName.Unset()
}

func (o MicrosoftGraphTargetResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.GroupType != nil {
		toSerialize["groupType"] = o.GroupType
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.ModifiedProperties != nil {
		toSerialize["modifiedProperties"] = o.ModifiedProperties
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.UserPrincipalName.IsSet() {
		toSerialize["userPrincipalName"] = o.UserPrincipalName.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTargetResource struct {
	value *MicrosoftGraphTargetResource
	isSet bool
}

func (v NullableMicrosoftGraphTargetResource) Get() *MicrosoftGraphTargetResource {
	return v.value
}

func (v *NullableMicrosoftGraphTargetResource) Set(val *MicrosoftGraphTargetResource) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTargetResource) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTargetResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTargetResource(val *MicrosoftGraphTargetResource) *NullableMicrosoftGraphTargetResource {
	return &NullableMicrosoftGraphTargetResource{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTargetResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTargetResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


