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

// MicrosoftGraphIncomingContext struct for MicrosoftGraphIncomingContext
type MicrosoftGraphIncomingContext struct {
	// The ID of the participant that is under observation. Read-only.
	ObservedParticipantId NullableString `json:"observedParticipantId,omitempty"`
	// The identity that the call is happening on behalf of.
	OnBehalfOf AnyOfmicrosoftGraphIdentitySet `json:"onBehalfOf,omitempty"`
	// The ID of the participant that triggered the incoming call. Read-only.
	SourceParticipantId NullableString `json:"sourceParticipantId,omitempty"`
	// The identity that transferred the call.
	Transferor AnyOfmicrosoftGraphIdentitySet `json:"transferor,omitempty"`
}

// NewMicrosoftGraphIncomingContext instantiates a new MicrosoftGraphIncomingContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphIncomingContext() *MicrosoftGraphIncomingContext {
	this := MicrosoftGraphIncomingContext{}
	return &this
}

// NewMicrosoftGraphIncomingContextWithDefaults instantiates a new MicrosoftGraphIncomingContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphIncomingContextWithDefaults() *MicrosoftGraphIncomingContext {
	this := MicrosoftGraphIncomingContext{}
	return &this
}

// GetObservedParticipantId returns the ObservedParticipantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIncomingContext) GetObservedParticipantId() string {
	if o == nil || o.ObservedParticipantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObservedParticipantId.Get()
}

// GetObservedParticipantIdOk returns a tuple with the ObservedParticipantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIncomingContext) GetObservedParticipantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObservedParticipantId.Get(), o.ObservedParticipantId.IsSet()
}

// HasObservedParticipantId returns a boolean if a field has been set.
func (o *MicrosoftGraphIncomingContext) HasObservedParticipantId() bool {
	if o != nil && o.ObservedParticipantId.IsSet() {
		return true
	}

	return false
}

// SetObservedParticipantId gets a reference to the given NullableString and assigns it to the ObservedParticipantId field.
func (o *MicrosoftGraphIncomingContext) SetObservedParticipantId(v string) {
	o.ObservedParticipantId.Set(&v)
}
// SetObservedParticipantIdNil sets the value for ObservedParticipantId to be an explicit nil
func (o *MicrosoftGraphIncomingContext) SetObservedParticipantIdNil() {
	o.ObservedParticipantId.Set(nil)
}

// UnsetObservedParticipantId ensures that no value is present for ObservedParticipantId, not even an explicit nil
func (o *MicrosoftGraphIncomingContext) UnsetObservedParticipantId() {
	o.ObservedParticipantId.Unset()
}

// GetOnBehalfOf returns the OnBehalfOf field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIncomingContext) GetOnBehalfOf() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.OnBehalfOf
}

// GetOnBehalfOfOk returns a tuple with the OnBehalfOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIncomingContext) GetOnBehalfOfOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.OnBehalfOf == nil {
		return nil, false
	}
	return &o.OnBehalfOf, true
}

// HasOnBehalfOf returns a boolean if a field has been set.
func (o *MicrosoftGraphIncomingContext) HasOnBehalfOf() bool {
	if o != nil && o.OnBehalfOf != nil {
		return true
	}

	return false
}

// SetOnBehalfOf gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the OnBehalfOf field.
func (o *MicrosoftGraphIncomingContext) SetOnBehalfOf(v AnyOfmicrosoftGraphIdentitySet) {
	o.OnBehalfOf = v
}

// GetSourceParticipantId returns the SourceParticipantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIncomingContext) GetSourceParticipantId() string {
	if o == nil || o.SourceParticipantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceParticipantId.Get()
}

// GetSourceParticipantIdOk returns a tuple with the SourceParticipantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIncomingContext) GetSourceParticipantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceParticipantId.Get(), o.SourceParticipantId.IsSet()
}

// HasSourceParticipantId returns a boolean if a field has been set.
func (o *MicrosoftGraphIncomingContext) HasSourceParticipantId() bool {
	if o != nil && o.SourceParticipantId.IsSet() {
		return true
	}

	return false
}

// SetSourceParticipantId gets a reference to the given NullableString and assigns it to the SourceParticipantId field.
func (o *MicrosoftGraphIncomingContext) SetSourceParticipantId(v string) {
	o.SourceParticipantId.Set(&v)
}
// SetSourceParticipantIdNil sets the value for SourceParticipantId to be an explicit nil
func (o *MicrosoftGraphIncomingContext) SetSourceParticipantIdNil() {
	o.SourceParticipantId.Set(nil)
}

// UnsetSourceParticipantId ensures that no value is present for SourceParticipantId, not even an explicit nil
func (o *MicrosoftGraphIncomingContext) UnsetSourceParticipantId() {
	o.SourceParticipantId.Unset()
}

// GetTransferor returns the Transferor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphIncomingContext) GetTransferor() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.Transferor
}

// GetTransferorOk returns a tuple with the Transferor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphIncomingContext) GetTransferorOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.Transferor == nil {
		return nil, false
	}
	return &o.Transferor, true
}

// HasTransferor returns a boolean if a field has been set.
func (o *MicrosoftGraphIncomingContext) HasTransferor() bool {
	if o != nil && o.Transferor != nil {
		return true
	}

	return false
}

// SetTransferor gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the Transferor field.
func (o *MicrosoftGraphIncomingContext) SetTransferor(v AnyOfmicrosoftGraphIdentitySet) {
	o.Transferor = v
}

func (o MicrosoftGraphIncomingContext) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObservedParticipantId.IsSet() {
		toSerialize["observedParticipantId"] = o.ObservedParticipantId.Get()
	}
	if o.OnBehalfOf != nil {
		toSerialize["onBehalfOf"] = o.OnBehalfOf
	}
	if o.SourceParticipantId.IsSet() {
		toSerialize["sourceParticipantId"] = o.SourceParticipantId.Get()
	}
	if o.Transferor != nil {
		toSerialize["transferor"] = o.Transferor
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphIncomingContext struct {
	value *MicrosoftGraphIncomingContext
	isSet bool
}

func (v NullableMicrosoftGraphIncomingContext) Get() *MicrosoftGraphIncomingContext {
	return v.value
}

func (v *NullableMicrosoftGraphIncomingContext) Set(val *MicrosoftGraphIncomingContext) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphIncomingContext) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphIncomingContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphIncomingContext(val *MicrosoftGraphIncomingContext) *NullableMicrosoftGraphIncomingContext {
	return &NullableMicrosoftGraphIncomingContext{value: val, isSet: true}
}

func (v NullableMicrosoftGraphIncomingContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphIncomingContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


