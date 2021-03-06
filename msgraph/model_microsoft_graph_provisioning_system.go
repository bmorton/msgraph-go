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

// MicrosoftGraphProvisioningSystem struct for MicrosoftGraphProvisioningSystem
type MicrosoftGraphProvisioningSystem struct {
	// The identity's display name. Note that this may not always be available or up to date. For example, if a user changes their display name, the API may show the new value in a future response, but the items associated with the user won't show up as having changed when using delta.
	DisplayName NullableString `json:"displayName,omitempty"`
	// Unique identifier for the identity.
	Id NullableString `json:"id,omitempty"`
	// Details of the system.
	Details AnyOfobject `json:"details,omitempty"`
}

// NewMicrosoftGraphProvisioningSystem instantiates a new MicrosoftGraphProvisioningSystem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphProvisioningSystem() *MicrosoftGraphProvisioningSystem {
	this := MicrosoftGraphProvisioningSystem{}
	return &this
}

// NewMicrosoftGraphProvisioningSystemWithDefaults instantiates a new MicrosoftGraphProvisioningSystem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphProvisioningSystemWithDefaults() *MicrosoftGraphProvisioningSystem {
	this := MicrosoftGraphProvisioningSystem{}
	return &this
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphProvisioningSystem) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphProvisioningSystem) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphProvisioningSystem) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphProvisioningSystem) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphProvisioningSystem) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphProvisioningSystem) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphProvisioningSystem) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphProvisioningSystem) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphProvisioningSystem) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphProvisioningSystem) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphProvisioningSystem) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphProvisioningSystem) UnsetId() {
	o.Id.Unset()
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphProvisioningSystem) GetDetails() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphProvisioningSystem) GetDetailsOk() (*AnyOfobject, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MicrosoftGraphProvisioningSystem) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given AnyOfobject and assigns it to the Details field.
func (o *MicrosoftGraphProvisioningSystem) SetDetails(v AnyOfobject) {
	o.Details = v
}

func (o MicrosoftGraphProvisioningSystem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphProvisioningSystem struct {
	value *MicrosoftGraphProvisioningSystem
	isSet bool
}

func (v NullableMicrosoftGraphProvisioningSystem) Get() *MicrosoftGraphProvisioningSystem {
	return v.value
}

func (v *NullableMicrosoftGraphProvisioningSystem) Set(val *MicrosoftGraphProvisioningSystem) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphProvisioningSystem) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphProvisioningSystem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphProvisioningSystem(val *MicrosoftGraphProvisioningSystem) *NullableMicrosoftGraphProvisioningSystem {
	return &NullableMicrosoftGraphProvisioningSystem{value: val, isSet: true}
}

func (v NullableMicrosoftGraphProvisioningSystem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphProvisioningSystem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


