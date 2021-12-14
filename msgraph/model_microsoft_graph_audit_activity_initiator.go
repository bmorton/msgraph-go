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

// MicrosoftGraphAuditActivityInitiator struct for MicrosoftGraphAuditActivityInitiator
type MicrosoftGraphAuditActivityInitiator struct {
	// If the resource initiating the activity is an app, this property indicates all the app related information like appId, Name, servicePrincipalId, Name.
	App AnyOfmicrosoftGraphAppIdentity `json:"app,omitempty"`
	// If the resource initiating the activity is a user, this property Indicates all the user related information like userId, Name, UserPrinicpalName.
	User AnyOfmicrosoftGraphUserIdentity `json:"user,omitempty"`
}

// NewMicrosoftGraphAuditActivityInitiator instantiates a new MicrosoftGraphAuditActivityInitiator object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAuditActivityInitiator() *MicrosoftGraphAuditActivityInitiator {
	this := MicrosoftGraphAuditActivityInitiator{}
	return &this
}

// NewMicrosoftGraphAuditActivityInitiatorWithDefaults instantiates a new MicrosoftGraphAuditActivityInitiator object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAuditActivityInitiatorWithDefaults() *MicrosoftGraphAuditActivityInitiator {
	this := MicrosoftGraphAuditActivityInitiator{}
	return &this
}

// GetApp returns the App field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuditActivityInitiator) GetApp() AnyOfmicrosoftGraphAppIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAppIdentity
		return ret
	}
	return o.App
}

// GetAppOk returns a tuple with the App field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuditActivityInitiator) GetAppOk() (*AnyOfmicrosoftGraphAppIdentity, bool) {
	if o == nil || o.App == nil {
		return nil, false
	}
	return &o.App, true
}

// HasApp returns a boolean if a field has been set.
func (o *MicrosoftGraphAuditActivityInitiator) HasApp() bool {
	if o != nil && o.App != nil {
		return true
	}

	return false
}

// SetApp gets a reference to the given AnyOfmicrosoftGraphAppIdentity and assigns it to the App field.
func (o *MicrosoftGraphAuditActivityInitiator) SetApp(v AnyOfmicrosoftGraphAppIdentity) {
	o.App = v
}

// GetUser returns the User field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAuditActivityInitiator) GetUser() AnyOfmicrosoftGraphUserIdentity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphUserIdentity
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAuditActivityInitiator) GetUserOk() (*AnyOfmicrosoftGraphUserIdentity, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return &o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MicrosoftGraphAuditActivityInitiator) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given AnyOfmicrosoftGraphUserIdentity and assigns it to the User field.
func (o *MicrosoftGraphAuditActivityInitiator) SetUser(v AnyOfmicrosoftGraphUserIdentity) {
	o.User = v
}

func (o MicrosoftGraphAuditActivityInitiator) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.App != nil {
		toSerialize["app"] = o.App
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAuditActivityInitiator struct {
	value *MicrosoftGraphAuditActivityInitiator
	isSet bool
}

func (v NullableMicrosoftGraphAuditActivityInitiator) Get() *MicrosoftGraphAuditActivityInitiator {
	return v.value
}

func (v *NullableMicrosoftGraphAuditActivityInitiator) Set(val *MicrosoftGraphAuditActivityInitiator) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAuditActivityInitiator) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAuditActivityInitiator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAuditActivityInitiator(val *MicrosoftGraphAuditActivityInitiator) *NullableMicrosoftGraphAuditActivityInitiator {
	return &NullableMicrosoftGraphAuditActivityInitiator{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAuditActivityInitiator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAuditActivityInitiator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


