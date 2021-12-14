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

// MicrosoftGraphResourceReference struct for MicrosoftGraphResourceReference
type MicrosoftGraphResourceReference struct {
	// The item's unique identifier.
	Id NullableString `json:"id,omitempty"`
	// A string value that can be used to classify the item, such as 'microsoft.graph.driveItem'
	Type NullableString `json:"type,omitempty"`
	// A URL leading to the referenced item.
	WebUrl NullableString `json:"webUrl,omitempty"`
}

// NewMicrosoftGraphResourceReference instantiates a new MicrosoftGraphResourceReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphResourceReference() *MicrosoftGraphResourceReference {
	this := MicrosoftGraphResourceReference{}
	return &this
}

// NewMicrosoftGraphResourceReferenceWithDefaults instantiates a new MicrosoftGraphResourceReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphResourceReferenceWithDefaults() *MicrosoftGraphResourceReference {
	this := MicrosoftGraphResourceReference{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceReference) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceReference) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceReference) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphResourceReference) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphResourceReference) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphResourceReference) UnsetId() {
	o.Id.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceReference) GetType() string {
	if o == nil || o.Type.Get() == nil {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceReference) GetTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceReference) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *MicrosoftGraphResourceReference) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *MicrosoftGraphResourceReference) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *MicrosoftGraphResourceReference) UnsetType() {
	o.Type.Unset()
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphResourceReference) GetWebUrl() string {
	if o == nil || o.WebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphResourceReference) GetWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphResourceReference) HasWebUrl() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *MicrosoftGraphResourceReference) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *MicrosoftGraphResourceReference) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *MicrosoftGraphResourceReference) UnsetWebUrl() {
	o.WebUrl.Unset()
}

func (o MicrosoftGraphResourceReference) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphResourceReference struct {
	value *MicrosoftGraphResourceReference
	isSet bool
}

func (v NullableMicrosoftGraphResourceReference) Get() *MicrosoftGraphResourceReference {
	return v.value
}

func (v *NullableMicrosoftGraphResourceReference) Set(val *MicrosoftGraphResourceReference) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphResourceReference) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphResourceReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphResourceReference(val *MicrosoftGraphResourceReference) *NullableMicrosoftGraphResourceReference {
	return &NullableMicrosoftGraphResourceReference{value: val, isSet: true}
}

func (v NullableMicrosoftGraphResourceReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphResourceReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

