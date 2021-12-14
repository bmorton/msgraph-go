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

// MicrosoftGraphPlannerProgressTaskBoardTaskFormat struct for MicrosoftGraphPlannerProgressTaskBoardTaskFormat
type MicrosoftGraphPlannerProgressTaskBoardTaskFormat struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Hint value used to order the task on the Progress view of the Task Board. The format is defined as outlined here.
	OrderHint NullableString `json:"orderHint,omitempty"`
}

// NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat instantiates a new MicrosoftGraphPlannerProgressTaskBoardTaskFormat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPlannerProgressTaskBoardTaskFormat() *MicrosoftGraphPlannerProgressTaskBoardTaskFormat {
	this := MicrosoftGraphPlannerProgressTaskBoardTaskFormat{}
	return &this
}

// NewMicrosoftGraphPlannerProgressTaskBoardTaskFormatWithDefaults instantiates a new MicrosoftGraphPlannerProgressTaskBoardTaskFormat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPlannerProgressTaskBoardTaskFormatWithDefaults() *MicrosoftGraphPlannerProgressTaskBoardTaskFormat {
	this := MicrosoftGraphPlannerProgressTaskBoardTaskFormat{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) SetId(v string) {
	o.Id = &v
}

// GetOrderHint returns the OrderHint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) GetOrderHint() string {
	if o == nil || o.OrderHint.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrderHint.Get()
}

// GetOrderHintOk returns a tuple with the OrderHint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) GetOrderHintOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OrderHint.Get(), o.OrderHint.IsSet()
}

// HasOrderHint returns a boolean if a field has been set.
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) HasOrderHint() bool {
	if o != nil && o.OrderHint.IsSet() {
		return true
	}

	return false
}

// SetOrderHint gets a reference to the given NullableString and assigns it to the OrderHint field.
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) SetOrderHint(v string) {
	o.OrderHint.Set(&v)
}
// SetOrderHintNil sets the value for OrderHint to be an explicit nil
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) SetOrderHintNil() {
	o.OrderHint.Set(nil)
}

// UnsetOrderHint ensures that no value is present for OrderHint, not even an explicit nil
func (o *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) UnsetOrderHint() {
	o.OrderHint.Unset()
}

func (o MicrosoftGraphPlannerProgressTaskBoardTaskFormat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrderHint.IsSet() {
		toSerialize["orderHint"] = o.OrderHint.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat struct {
	value *MicrosoftGraphPlannerProgressTaskBoardTaskFormat
	isSet bool
}

func (v NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat) Get() *MicrosoftGraphPlannerProgressTaskBoardTaskFormat {
	return v.value
}

func (v *NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat) Set(val *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat(val *MicrosoftGraphPlannerProgressTaskBoardTaskFormat) *NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat {
	return &NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPlannerProgressTaskBoardTaskFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


