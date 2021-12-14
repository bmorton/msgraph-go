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

// InlineObject30 struct for InlineObject30
type InlineObject30 struct {
	TransferTarget *MicrosoftGraphInvitationParticipantInfo `json:"transferTarget,omitempty"`
	Transferee AnyOfmicrosoftGraphParticipantInfo `json:"transferee,omitempty"`
}

// NewInlineObject30 instantiates a new InlineObject30 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject30() *InlineObject30 {
	this := InlineObject30{}
	return &this
}

// NewInlineObject30WithDefaults instantiates a new InlineObject30 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject30WithDefaults() *InlineObject30 {
	this := InlineObject30{}
	return &this
}

// GetTransferTarget returns the TransferTarget field value if set, zero value otherwise.
func (o *InlineObject30) GetTransferTarget() MicrosoftGraphInvitationParticipantInfo {
	if o == nil || o.TransferTarget == nil {
		var ret MicrosoftGraphInvitationParticipantInfo
		return ret
	}
	return *o.TransferTarget
}

// GetTransferTargetOk returns a tuple with the TransferTarget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject30) GetTransferTargetOk() (*MicrosoftGraphInvitationParticipantInfo, bool) {
	if o == nil || o.TransferTarget == nil {
		return nil, false
	}
	return o.TransferTarget, true
}

// HasTransferTarget returns a boolean if a field has been set.
func (o *InlineObject30) HasTransferTarget() bool {
	if o != nil && o.TransferTarget != nil {
		return true
	}

	return false
}

// SetTransferTarget gets a reference to the given MicrosoftGraphInvitationParticipantInfo and assigns it to the TransferTarget field.
func (o *InlineObject30) SetTransferTarget(v MicrosoftGraphInvitationParticipantInfo) {
	o.TransferTarget = &v
}

// GetTransferee returns the Transferee field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject30) GetTransferee() AnyOfmicrosoftGraphParticipantInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphParticipantInfo
		return ret
	}
	return o.Transferee
}

// GetTransfereeOk returns a tuple with the Transferee field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject30) GetTransfereeOk() (*AnyOfmicrosoftGraphParticipantInfo, bool) {
	if o == nil || o.Transferee == nil {
		return nil, false
	}
	return &o.Transferee, true
}

// HasTransferee returns a boolean if a field has been set.
func (o *InlineObject30) HasTransferee() bool {
	if o != nil && o.Transferee != nil {
		return true
	}

	return false
}

// SetTransferee gets a reference to the given AnyOfmicrosoftGraphParticipantInfo and assigns it to the Transferee field.
func (o *InlineObject30) SetTransferee(v AnyOfmicrosoftGraphParticipantInfo) {
	o.Transferee = v
}

func (o InlineObject30) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TransferTarget != nil {
		toSerialize["transferTarget"] = o.TransferTarget
	}
	if o.Transferee != nil {
		toSerialize["transferee"] = o.Transferee
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject30 struct {
	value *InlineObject30
	isSet bool
}

func (v NullableInlineObject30) Get() *InlineObject30 {
	return v.value
}

func (v *NullableInlineObject30) Set(val *InlineObject30) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject30) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject30) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject30(val *InlineObject30) *NullableInlineObject30 {
	return &NullableInlineObject30{value: val, isSet: true}
}

func (v NullableInlineObject30) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject30) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


