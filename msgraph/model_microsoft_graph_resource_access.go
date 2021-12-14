/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MicrosoftGraphResourceAccess struct for MicrosoftGraphResourceAccess
type MicrosoftGraphResourceAccess struct {
	// The unique identifier for one of the oauth2PermissionScopes or appRole instances that the resource application exposes.
	Id *string `json:"id,omitempty"`
	// Specifies whether the id property references an oauth2PermissionScopes or an appRole. The possible values are: Scope (for OAuth 2.0 permission scopes) or Role (for app roles).
	Type NullableString `json:"type,omitempty"`
}

// NewMicrosoftGraphResourceAccess instantiates a new MicrosoftGraphResourceAccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphResourceAccess() *MicrosoftGraphResourceAccess {
	this := MicrosoftGraphResourceAccess{}
	return &this
}

// NewMicrosoftGraphResourceAccessWithDefaults instantiates a new MicrosoftGraphResourceAccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphResourceAccessWithDefaults() *MicrosoftGraphResourceAccess {
	this := MicrosoftGraphResourceAccess{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphResourceAccess) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphResourceAccess) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceAccess) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphResourceAccess) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceAccess) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceAccess) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceAccess) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *MicrosoftGraphResourceAccess) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *MicrosoftGraphResourceAccess) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *MicrosoftGraphResourceAccess) UnsetType() {
	o.Type.Unset()
}

func (o MicrosoftGraphResourceAccess) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphResourceAccess struct {
	value *MicrosoftGraphResourceAccess
	isSet bool
}

func (v NullableMicrosoftGraphResourceAccess) Get() *MicrosoftGraphResourceAccess {
	return v.value
}

func (v *NullableMicrosoftGraphResourceAccess) Set(val *MicrosoftGraphResourceAccess) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphResourceAccess) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphResourceAccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphResourceAccess(val *MicrosoftGraphResourceAccess) *NullableMicrosoftGraphResourceAccess {
	return &NullableMicrosoftGraphResourceAccess{value: val, isSet: true}
}

func (v NullableMicrosoftGraphResourceAccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphResourceAccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


