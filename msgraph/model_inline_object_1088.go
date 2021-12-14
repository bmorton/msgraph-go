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

// InlineObject1088 struct for InlineObject1088
type InlineObject1088 struct {
	Message *MicrosoftGraphMessage `json:"Message,omitempty"`
	SaveToSentItems NullableBool `json:"SaveToSentItems,omitempty"`
}

// NewInlineObject1088 instantiates a new InlineObject1088 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1088() *InlineObject1088 {
	this := InlineObject1088{}
	var saveToSentItems bool = false
	this.SaveToSentItems = *NewNullableBool(&saveToSentItems)
	return &this
}

// NewInlineObject1088WithDefaults instantiates a new InlineObject1088 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1088WithDefaults() *InlineObject1088 {
	this := InlineObject1088{}
	var saveToSentItems bool = false
	this.SaveToSentItems = *NewNullableBool(&saveToSentItems)
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineObject1088) GetMessage() MicrosoftGraphMessage {
	if o == nil || o.Message == nil {
		var ret MicrosoftGraphMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1088) GetMessageOk() (*MicrosoftGraphMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineObject1088) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given MicrosoftGraphMessage and assigns it to the Message field.
func (o *InlineObject1088) SetMessage(v MicrosoftGraphMessage) {
	o.Message = &v
}

// GetSaveToSentItems returns the SaveToSentItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject1088) GetSaveToSentItems() bool {
	if o == nil || o.SaveToSentItems.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SaveToSentItems.Get()
}

// GetSaveToSentItemsOk returns a tuple with the SaveToSentItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject1088) GetSaveToSentItemsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SaveToSentItems.Get(), o.SaveToSentItems.IsSet()
}

// HasSaveToSentItems returns a boolean if a field has been set.
func (o *InlineObject1088) HasSaveToSentItems() bool {
	if o != nil && o.SaveToSentItems.IsSet() {
		return true
	}

	return false
}

// SetSaveToSentItems gets a reference to the given NullableBool and assigns it to the SaveToSentItems field.
func (o *InlineObject1088) SetSaveToSentItems(v bool) {
	o.SaveToSentItems.Set(&v)
}
// SetSaveToSentItemsNil sets the value for SaveToSentItems to be an explicit nil
func (o *InlineObject1088) SetSaveToSentItemsNil() {
	o.SaveToSentItems.Set(nil)
}

// UnsetSaveToSentItems ensures that no value is present for SaveToSentItems, not even an explicit nil
func (o *InlineObject1088) UnsetSaveToSentItems() {
	o.SaveToSentItems.Unset()
}

func (o InlineObject1088) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.SaveToSentItems.IsSet() {
		toSerialize["SaveToSentItems"] = o.SaveToSentItems.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject1088 struct {
	value *InlineObject1088
	isSet bool
}

func (v NullableInlineObject1088) Get() *InlineObject1088 {
	return v.value
}

func (v *NullableInlineObject1088) Set(val *InlineObject1088) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1088) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1088) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1088(val *InlineObject1088) *NullableInlineObject1088 {
	return &NullableInlineObject1088{value: val, isSet: true}
}

func (v NullableInlineObject1088) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1088) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


