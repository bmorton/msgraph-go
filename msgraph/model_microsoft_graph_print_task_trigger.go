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

// MicrosoftGraphPrintTaskTrigger struct for MicrosoftGraphPrintTaskTrigger
type MicrosoftGraphPrintTaskTrigger struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The Universal Print event that will cause a new printTask to be triggered. Valid values are described in the following table.
	Event AnyOfmicrosoftGraphPrintEvent `json:"event,omitempty"`
	Definition *MicrosoftGraphPrintTaskDefinition `json:"definition,omitempty"`
}

// NewMicrosoftGraphPrintTaskTrigger instantiates a new MicrosoftGraphPrintTaskTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintTaskTrigger() *MicrosoftGraphPrintTaskTrigger {
	this := MicrosoftGraphPrintTaskTrigger{}
	return &this
}

// NewMicrosoftGraphPrintTaskTriggerWithDefaults instantiates a new MicrosoftGraphPrintTaskTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintTaskTriggerWithDefaults() *MicrosoftGraphPrintTaskTrigger {
	this := MicrosoftGraphPrintTaskTrigger{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintTaskTrigger) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintTaskTrigger) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintTaskTrigger) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPrintTaskTrigger) SetId(v string) {
	o.Id = &v
}

// GetEvent returns the Event field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrintTaskTrigger) GetEvent() AnyOfmicrosoftGraphPrintEvent {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrintEvent
		return ret
	}
	return o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrintTaskTrigger) GetEventOk() (*AnyOfmicrosoftGraphPrintEvent, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return &o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintTaskTrigger) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given AnyOfmicrosoftGraphPrintEvent and assigns it to the Event field.
func (o *MicrosoftGraphPrintTaskTrigger) SetEvent(v AnyOfmicrosoftGraphPrintEvent) {
	o.Event = v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintTaskTrigger) GetDefinition() MicrosoftGraphPrintTaskDefinition {
	if o == nil || o.Definition == nil {
		var ret MicrosoftGraphPrintTaskDefinition
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintTaskTrigger) GetDefinitionOk() (*MicrosoftGraphPrintTaskDefinition, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintTaskTrigger) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given MicrosoftGraphPrintTaskDefinition and assigns it to the Definition field.
func (o *MicrosoftGraphPrintTaskTrigger) SetDefinition(v MicrosoftGraphPrintTaskDefinition) {
	o.Definition = &v
}

func (o MicrosoftGraphPrintTaskTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintTaskTrigger struct {
	value *MicrosoftGraphPrintTaskTrigger
	isSet bool
}

func (v NullableMicrosoftGraphPrintTaskTrigger) Get() *MicrosoftGraphPrintTaskTrigger {
	return v.value
}

func (v *NullableMicrosoftGraphPrintTaskTrigger) Set(val *MicrosoftGraphPrintTaskTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintTaskTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintTaskTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintTaskTrigger(val *MicrosoftGraphPrintTaskTrigger) *NullableMicrosoftGraphPrintTaskTrigger {
	return &NullableMicrosoftGraphPrintTaskTrigger{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintTaskTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintTaskTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

