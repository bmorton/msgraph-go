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

// MicrosoftGraphTeamworkActivityTopic struct for MicrosoftGraphTeamworkActivityTopic
type MicrosoftGraphTeamworkActivityTopic struct {
	// Type of source. Possible values are: entityUrl, text. For supported Microsoft Graph URLs, use entityUrl. For custom text, use text.
	Source AnyOfmicrosoftGraphTeamworkActivityTopicSource `json:"source,omitempty"`
	// The topic value. If the value of the source property is entityUrl, this must be a Microsoft Graph URL. If the vaule is text, this must be a plain text value.
	Value *string `json:"value,omitempty"`
	// The link the user clicks when they select the notification. Optional when source is entityUrl; required when source is text.
	WebUrl NullableString `json:"webUrl,omitempty"`
}

// NewMicrosoftGraphTeamworkActivityTopic instantiates a new MicrosoftGraphTeamworkActivityTopic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTeamworkActivityTopic() *MicrosoftGraphTeamworkActivityTopic {
	this := MicrosoftGraphTeamworkActivityTopic{}
	return &this
}

// NewMicrosoftGraphTeamworkActivityTopicWithDefaults instantiates a new MicrosoftGraphTeamworkActivityTopic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTeamworkActivityTopicWithDefaults() *MicrosoftGraphTeamworkActivityTopic {
	this := MicrosoftGraphTeamworkActivityTopic{}
	return &this
}

// GetSource returns the Source field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamworkActivityTopic) GetSource() AnyOfmicrosoftGraphTeamworkActivityTopicSource {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamworkActivityTopicSource
		return ret
	}
	return o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamworkActivityTopic) GetSourceOk() (*AnyOfmicrosoftGraphTeamworkActivityTopicSource, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return &o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamworkActivityTopic) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given AnyOfmicrosoftGraphTeamworkActivityTopicSource and assigns it to the Source field.
func (o *MicrosoftGraphTeamworkActivityTopic) SetSource(v AnyOfmicrosoftGraphTeamworkActivityTopicSource) {
	o.Source = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *MicrosoftGraphTeamworkActivityTopic) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamworkActivityTopic) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamworkActivityTopic) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *MicrosoftGraphTeamworkActivityTopic) SetValue(v string) {
	o.Value = &v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamworkActivityTopic) GetWebUrl() string {
	if o == nil || o.WebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamworkActivityTopic) GetWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamworkActivityTopic) HasWebUrl() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *MicrosoftGraphTeamworkActivityTopic) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *MicrosoftGraphTeamworkActivityTopic) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *MicrosoftGraphTeamworkActivityTopic) UnsetWebUrl() {
	o.WebUrl.Unset()
}

func (o MicrosoftGraphTeamworkActivityTopic) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTeamworkActivityTopic struct {
	value *MicrosoftGraphTeamworkActivityTopic
	isSet bool
}

func (v NullableMicrosoftGraphTeamworkActivityTopic) Get() *MicrosoftGraphTeamworkActivityTopic {
	return v.value
}

func (v *NullableMicrosoftGraphTeamworkActivityTopic) Set(val *MicrosoftGraphTeamworkActivityTopic) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamworkActivityTopic) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamworkActivityTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamworkActivityTopic(val *MicrosoftGraphTeamworkActivityTopic) *NullableMicrosoftGraphTeamworkActivityTopic {
	return &NullableMicrosoftGraphTeamworkActivityTopic{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamworkActivityTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamworkActivityTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


