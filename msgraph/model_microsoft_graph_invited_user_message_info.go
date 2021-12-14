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

// MicrosoftGraphInvitedUserMessageInfo struct for MicrosoftGraphInvitedUserMessageInfo
type MicrosoftGraphInvitedUserMessageInfo struct {
	// Additional recipients the invitation message should be sent to. Currently only 1 additional recipient is supported.
	CcRecipients *[]*AnyOfmicrosoftGraphRecipient `json:"ccRecipients,omitempty"`
	// Customized message body you want to send if you don't want the default message.
	CustomizedMessageBody NullableString `json:"customizedMessageBody,omitempty"`
	// The language you want to send the default message in. If the customizedMessageBody is specified, this property is ignored, and the message is sent using the customizedMessageBody. The language format should be in ISO 639. The default is en-US.
	MessageLanguage NullableString `json:"messageLanguage,omitempty"`
}

// NewMicrosoftGraphInvitedUserMessageInfo instantiates a new MicrosoftGraphInvitedUserMessageInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphInvitedUserMessageInfo() *MicrosoftGraphInvitedUserMessageInfo {
	this := MicrosoftGraphInvitedUserMessageInfo{}
	return &this
}

// NewMicrosoftGraphInvitedUserMessageInfoWithDefaults instantiates a new MicrosoftGraphInvitedUserMessageInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphInvitedUserMessageInfoWithDefaults() *MicrosoftGraphInvitedUserMessageInfo {
	this := MicrosoftGraphInvitedUserMessageInfo{}
	return &this
}

// GetCcRecipients returns the CcRecipients field value if set, zero value otherwise.
func (o *MicrosoftGraphInvitedUserMessageInfo) GetCcRecipients() []*AnyOfmicrosoftGraphRecipient {
	if o == nil || o.CcRecipients == nil {
		var ret []*AnyOfmicrosoftGraphRecipient
		return ret
	}
	return *o.CcRecipients
}

// GetCcRecipientsOk returns a tuple with the CcRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphInvitedUserMessageInfo) GetCcRecipientsOk() (*[]*AnyOfmicrosoftGraphRecipient, bool) {
	if o == nil || o.CcRecipients == nil {
		return nil, false
	}
	return o.CcRecipients, true
}

// HasCcRecipients returns a boolean if a field has been set.
func (o *MicrosoftGraphInvitedUserMessageInfo) HasCcRecipients() bool {
	if o != nil && o.CcRecipients != nil {
		return true
	}

	return false
}

// SetCcRecipients gets a reference to the given []*AnyOfmicrosoftGraphRecipient and assigns it to the CcRecipients field.
func (o *MicrosoftGraphInvitedUserMessageInfo) SetCcRecipients(v []*AnyOfmicrosoftGraphRecipient) {
	o.CcRecipients = &v
}

// GetCustomizedMessageBody returns the CustomizedMessageBody field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInvitedUserMessageInfo) GetCustomizedMessageBody() string {
	if o == nil || o.CustomizedMessageBody.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomizedMessageBody.Get()
}

// GetCustomizedMessageBodyOk returns a tuple with the CustomizedMessageBody field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInvitedUserMessageInfo) GetCustomizedMessageBodyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomizedMessageBody.Get(), o.CustomizedMessageBody.IsSet()
}

// HasCustomizedMessageBody returns a boolean if a field has been set.
func (o *MicrosoftGraphInvitedUserMessageInfo) HasCustomizedMessageBody() bool {
	if o != nil && o.CustomizedMessageBody.IsSet() {
		return true
	}

	return false
}

// SetCustomizedMessageBody gets a reference to the given NullableString and assigns it to the CustomizedMessageBody field.
func (o *MicrosoftGraphInvitedUserMessageInfo) SetCustomizedMessageBody(v string) {
	o.CustomizedMessageBody.Set(&v)
}
// SetCustomizedMessageBodyNil sets the value for CustomizedMessageBody to be an explicit nil
func (o *MicrosoftGraphInvitedUserMessageInfo) SetCustomizedMessageBodyNil() {
	o.CustomizedMessageBody.Set(nil)
}

// UnsetCustomizedMessageBody ensures that no value is present for CustomizedMessageBody, not even an explicit nil
func (o *MicrosoftGraphInvitedUserMessageInfo) UnsetCustomizedMessageBody() {
	o.CustomizedMessageBody.Unset()
}

// GetMessageLanguage returns the MessageLanguage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphInvitedUserMessageInfo) GetMessageLanguage() string {
	if o == nil || o.MessageLanguage.Get() == nil {
		var ret string
		return ret
	}
	return *o.MessageLanguage.Get()
}

// GetMessageLanguageOk returns a tuple with the MessageLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphInvitedUserMessageInfo) GetMessageLanguageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MessageLanguage.Get(), o.MessageLanguage.IsSet()
}

// HasMessageLanguage returns a boolean if a field has been set.
func (o *MicrosoftGraphInvitedUserMessageInfo) HasMessageLanguage() bool {
	if o != nil && o.MessageLanguage.IsSet() {
		return true
	}

	return false
}

// SetMessageLanguage gets a reference to the given NullableString and assigns it to the MessageLanguage field.
func (o *MicrosoftGraphInvitedUserMessageInfo) SetMessageLanguage(v string) {
	o.MessageLanguage.Set(&v)
}
// SetMessageLanguageNil sets the value for MessageLanguage to be an explicit nil
func (o *MicrosoftGraphInvitedUserMessageInfo) SetMessageLanguageNil() {
	o.MessageLanguage.Set(nil)
}

// UnsetMessageLanguage ensures that no value is present for MessageLanguage, not even an explicit nil
func (o *MicrosoftGraphInvitedUserMessageInfo) UnsetMessageLanguage() {
	o.MessageLanguage.Unset()
}

func (o MicrosoftGraphInvitedUserMessageInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CcRecipients != nil {
		toSerialize["ccRecipients"] = o.CcRecipients
	}
	if o.CustomizedMessageBody.IsSet() {
		toSerialize["customizedMessageBody"] = o.CustomizedMessageBody.Get()
	}
	if o.MessageLanguage.IsSet() {
		toSerialize["messageLanguage"] = o.MessageLanguage.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphInvitedUserMessageInfo struct {
	value *MicrosoftGraphInvitedUserMessageInfo
	isSet bool
}

func (v NullableMicrosoftGraphInvitedUserMessageInfo) Get() *MicrosoftGraphInvitedUserMessageInfo {
	return v.value
}

func (v *NullableMicrosoftGraphInvitedUserMessageInfo) Set(val *MicrosoftGraphInvitedUserMessageInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphInvitedUserMessageInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphInvitedUserMessageInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphInvitedUserMessageInfo(val *MicrosoftGraphInvitedUserMessageInfo) *NullableMicrosoftGraphInvitedUserMessageInfo {
	return &NullableMicrosoftGraphInvitedUserMessageInfo{value: val, isSet: true}
}

func (v NullableMicrosoftGraphInvitedUserMessageInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphInvitedUserMessageInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


