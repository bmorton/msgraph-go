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

// MicrosoftGraphVisualInfo struct for MicrosoftGraphVisualInfo
type MicrosoftGraphVisualInfo struct {
	// Optional. JSON object used to represent an icon which represents the application used to generate the activity
	Attribution AnyOfmicrosoftGraphImageInfo `json:"attribution,omitempty"`
	// Optional. Background color used to render the activity in the UI - brand color for the application source of the activity. Must be a valid hex color
	BackgroundColor NullableString `json:"backgroundColor,omitempty"`
	// Optional. Custom piece of data - JSON object used to provide custom content to render the activity in the Windows Shell UI
	Content AnyOfobject `json:"content,omitempty"`
	// Optional. Longer text description of the user's unique activity (example: document name, first sentence, and/or metadata)
	Description NullableString `json:"description,omitempty"`
	// Required. Short text description of the user's unique activity (for example, document name in cases where an activity refers to document creation)
	DisplayText *string `json:"displayText,omitempty"`
}

// NewMicrosoftGraphVisualInfo instantiates a new MicrosoftGraphVisualInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphVisualInfo() *MicrosoftGraphVisualInfo {
	this := MicrosoftGraphVisualInfo{}
	return &this
}

// NewMicrosoftGraphVisualInfoWithDefaults instantiates a new MicrosoftGraphVisualInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphVisualInfoWithDefaults() *MicrosoftGraphVisualInfo {
	this := MicrosoftGraphVisualInfo{}
	return &this
}

// GetAttribution returns the Attribution field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVisualInfo) GetAttribution() AnyOfmicrosoftGraphImageInfo {
	if o == nil  {
		var ret AnyOfmicrosoftGraphImageInfo
		return ret
	}
	return o.Attribution
}

// GetAttributionOk returns a tuple with the Attribution field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVisualInfo) GetAttributionOk() (*AnyOfmicrosoftGraphImageInfo, bool) {
	if o == nil || o.Attribution == nil {
		return nil, false
	}
	return &o.Attribution, true
}

// HasAttribution returns a boolean if a field has been set.
func (o *MicrosoftGraphVisualInfo) HasAttribution() bool {
	if o != nil && o.Attribution != nil {
		return true
	}

	return false
}

// SetAttribution gets a reference to the given AnyOfmicrosoftGraphImageInfo and assigns it to the Attribution field.
func (o *MicrosoftGraphVisualInfo) SetAttribution(v AnyOfmicrosoftGraphImageInfo) {
	o.Attribution = v
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVisualInfo) GetBackgroundColor() string {
	if o == nil || o.BackgroundColor.Get() == nil {
		var ret string
		return ret
	}
	return *o.BackgroundColor.Get()
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVisualInfo) GetBackgroundColorOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BackgroundColor.Get(), o.BackgroundColor.IsSet()
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *MicrosoftGraphVisualInfo) HasBackgroundColor() bool {
	if o != nil && o.BackgroundColor.IsSet() {
		return true
	}

	return false
}

// SetBackgroundColor gets a reference to the given NullableString and assigns it to the BackgroundColor field.
func (o *MicrosoftGraphVisualInfo) SetBackgroundColor(v string) {
	o.BackgroundColor.Set(&v)
}
// SetBackgroundColorNil sets the value for BackgroundColor to be an explicit nil
func (o *MicrosoftGraphVisualInfo) SetBackgroundColorNil() {
	o.BackgroundColor.Set(nil)
}

// UnsetBackgroundColor ensures that no value is present for BackgroundColor, not even an explicit nil
func (o *MicrosoftGraphVisualInfo) UnsetBackgroundColor() {
	o.BackgroundColor.Unset()
}

// GetContent returns the Content field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVisualInfo) GetContent() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVisualInfo) GetContentOk() (*AnyOfobject, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return &o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *MicrosoftGraphVisualInfo) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given AnyOfobject and assigns it to the Content field.
func (o *MicrosoftGraphVisualInfo) SetContent(v AnyOfobject) {
	o.Content = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVisualInfo) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVisualInfo) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphVisualInfo) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphVisualInfo) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphVisualInfo) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphVisualInfo) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayText returns the DisplayText field value if set, zero value otherwise.
func (o *MicrosoftGraphVisualInfo) GetDisplayText() string {
	if o == nil || o.DisplayText == nil {
		var ret string
		return ret
	}
	return *o.DisplayText
}

// GetDisplayTextOk returns a tuple with the DisplayText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVisualInfo) GetDisplayTextOk() (*string, bool) {
	if o == nil || o.DisplayText == nil {
		return nil, false
	}
	return o.DisplayText, true
}

// HasDisplayText returns a boolean if a field has been set.
func (o *MicrosoftGraphVisualInfo) HasDisplayText() bool {
	if o != nil && o.DisplayText != nil {
		return true
	}

	return false
}

// SetDisplayText gets a reference to the given string and assigns it to the DisplayText field.
func (o *MicrosoftGraphVisualInfo) SetDisplayText(v string) {
	o.DisplayText = &v
}

func (o MicrosoftGraphVisualInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attribution != nil {
		toSerialize["attribution"] = o.Attribution
	}
	if o.BackgroundColor.IsSet() {
		toSerialize["backgroundColor"] = o.BackgroundColor.Get()
	}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayText != nil {
		toSerialize["displayText"] = o.DisplayText
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphVisualInfo struct {
	value *MicrosoftGraphVisualInfo
	isSet bool
}

func (v NullableMicrosoftGraphVisualInfo) Get() *MicrosoftGraphVisualInfo {
	return v.value
}

func (v *NullableMicrosoftGraphVisualInfo) Set(val *MicrosoftGraphVisualInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphVisualInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphVisualInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphVisualInfo(val *MicrosoftGraphVisualInfo) *NullableMicrosoftGraphVisualInfo {
	return &NullableMicrosoftGraphVisualInfo{value: val, isSet: true}
}

func (v NullableMicrosoftGraphVisualInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphVisualInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


