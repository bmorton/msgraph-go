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

// InlineObject162 struct for InlineObject162
type InlineObject162 struct {
	Comment NullableString `json:"Comment,omitempty"`
	SendResponse NullableBool `json:"SendResponse,omitempty"`
	ProposedNewTime AnyOfmicrosoftGraphTimeSlot `json:"ProposedNewTime,omitempty"`
}

// NewInlineObject162 instantiates a new InlineObject162 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject162() *InlineObject162 {
	this := InlineObject162{}
	var sendResponse bool = false
	this.SendResponse = *NewNullableBool(&sendResponse)
	return &this
}

// NewInlineObject162WithDefaults instantiates a new InlineObject162 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject162WithDefaults() *InlineObject162 {
	this := InlineObject162{}
	var sendResponse bool = false
	this.SendResponse = *NewNullableBool(&sendResponse)
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject162) GetComment() string {
	if o == nil || o.Comment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Comment.Get()
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject162) GetCommentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Comment.Get(), o.Comment.IsSet()
}

// HasComment returns a boolean if a field has been set.
func (o *InlineObject162) HasComment() bool {
	if o != nil && o.Comment.IsSet() {
		return true
	}

	return false
}

// SetComment gets a reference to the given NullableString and assigns it to the Comment field.
func (o *InlineObject162) SetComment(v string) {
	o.Comment.Set(&v)
}
// SetCommentNil sets the value for Comment to be an explicit nil
func (o *InlineObject162) SetCommentNil() {
	o.Comment.Set(nil)
}

// UnsetComment ensures that no value is present for Comment, not even an explicit nil
func (o *InlineObject162) UnsetComment() {
	o.Comment.Unset()
}

// GetSendResponse returns the SendResponse field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject162) GetSendResponse() bool {
	if o == nil || o.SendResponse.Get() == nil {
		var ret bool
		return ret
	}
	return *o.SendResponse.Get()
}

// GetSendResponseOk returns a tuple with the SendResponse field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject162) GetSendResponseOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SendResponse.Get(), o.SendResponse.IsSet()
}

// HasSendResponse returns a boolean if a field has been set.
func (o *InlineObject162) HasSendResponse() bool {
	if o != nil && o.SendResponse.IsSet() {
		return true
	}

	return false
}

// SetSendResponse gets a reference to the given NullableBool and assigns it to the SendResponse field.
func (o *InlineObject162) SetSendResponse(v bool) {
	o.SendResponse.Set(&v)
}
// SetSendResponseNil sets the value for SendResponse to be an explicit nil
func (o *InlineObject162) SetSendResponseNil() {
	o.SendResponse.Set(nil)
}

// UnsetSendResponse ensures that no value is present for SendResponse, not even an explicit nil
func (o *InlineObject162) UnsetSendResponse() {
	o.SendResponse.Unset()
}

// GetProposedNewTime returns the ProposedNewTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject162) GetProposedNewTime() AnyOfmicrosoftGraphTimeSlot {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTimeSlot
		return ret
	}
	return o.ProposedNewTime
}

// GetProposedNewTimeOk returns a tuple with the ProposedNewTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject162) GetProposedNewTimeOk() (*AnyOfmicrosoftGraphTimeSlot, bool) {
	if o == nil || o.ProposedNewTime == nil {
		return nil, false
	}
	return &o.ProposedNewTime, true
}

// HasProposedNewTime returns a boolean if a field has been set.
func (o *InlineObject162) HasProposedNewTime() bool {
	if o != nil && o.ProposedNewTime != nil {
		return true
	}

	return false
}

// SetProposedNewTime gets a reference to the given AnyOfmicrosoftGraphTimeSlot and assigns it to the ProposedNewTime field.
func (o *InlineObject162) SetProposedNewTime(v AnyOfmicrosoftGraphTimeSlot) {
	o.ProposedNewTime = v
}

func (o InlineObject162) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Comment.IsSet() {
		toSerialize["Comment"] = o.Comment.Get()
	}
	if o.SendResponse.IsSet() {
		toSerialize["SendResponse"] = o.SendResponse.Get()
	}
	if o.ProposedNewTime != nil {
		toSerialize["ProposedNewTime"] = o.ProposedNewTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject162 struct {
	value *InlineObject162
	isSet bool
}

func (v NullableInlineObject162) Get() *InlineObject162 {
	return v.value
}

func (v *NullableInlineObject162) Set(val *InlineObject162) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject162) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject162) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject162(val *InlineObject162) *NullableInlineObject162 {
	return &NullableInlineObject162{value: val, isSet: true}
}

func (v NullableInlineObject162) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject162) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

