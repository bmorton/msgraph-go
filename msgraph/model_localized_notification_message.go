/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// LocalizedNotificationMessage The text content of a Notification Message Template for the specified locale.
type LocalizedNotificationMessage struct {
	// Flag to indicate whether or not this is the default locale for language fallback. This flag can only be set. To unset, set this property to true on another Localized Notification Message.
	IsDefault *bool `json:"isDefault,omitempty"`
	// DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The Locale for which this message is destined.
	Locale *string `json:"locale,omitempty"`
	// The Message Template content.
	MessageTemplate *string `json:"messageTemplate,omitempty"`
	// The Message Template Subject.
	Subject *string `json:"subject,omitempty"`
}

// NewLocalizedNotificationMessage instantiates a new LocalizedNotificationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalizedNotificationMessage() *LocalizedNotificationMessage {
	this := LocalizedNotificationMessage{}
	return &this
}

// NewLocalizedNotificationMessageWithDefaults instantiates a new LocalizedNotificationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalizedNotificationMessageWithDefaults() *LocalizedNotificationMessage {
	this := LocalizedNotificationMessage{}
	return &this
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *LocalizedNotificationMessage) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *LocalizedNotificationMessage) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *LocalizedNotificationMessage) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *LocalizedNotificationMessage) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *LocalizedNotificationMessage) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *LocalizedNotificationMessage) SetLocale(v string) {
	o.Locale = &v
}

// GetMessageTemplate returns the MessageTemplate field value if set, zero value otherwise.
func (o *LocalizedNotificationMessage) GetMessageTemplate() string {
	if o == nil || o.MessageTemplate == nil {
		var ret string
		return ret
	}
	return *o.MessageTemplate
}

// GetMessageTemplateOk returns a tuple with the MessageTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetMessageTemplateOk() (*string, bool) {
	if o == nil || o.MessageTemplate == nil {
		return nil, false
	}
	return o.MessageTemplate, true
}

// HasMessageTemplate returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasMessageTemplate() bool {
	if o != nil && o.MessageTemplate != nil {
		return true
	}

	return false
}

// SetMessageTemplate gets a reference to the given string and assigns it to the MessageTemplate field.
func (o *LocalizedNotificationMessage) SetMessageTemplate(v string) {
	o.MessageTemplate = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *LocalizedNotificationMessage) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalizedNotificationMessage) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *LocalizedNotificationMessage) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *LocalizedNotificationMessage) SetSubject(v string) {
	o.Subject = &v
}

func (o LocalizedNotificationMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsDefault != nil {
		toSerialize["isDefault"] = o.IsDefault
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Locale != nil {
		toSerialize["locale"] = o.Locale
	}
	if o.MessageTemplate != nil {
		toSerialize["messageTemplate"] = o.MessageTemplate
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableLocalizedNotificationMessage struct {
	value *LocalizedNotificationMessage
	isSet bool
}

func (v NullableLocalizedNotificationMessage) Get() *LocalizedNotificationMessage {
	return v.value
}

func (v *NullableLocalizedNotificationMessage) Set(val *LocalizedNotificationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalizedNotificationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalizedNotificationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalizedNotificationMessage(val *LocalizedNotificationMessage) *NullableLocalizedNotificationMessage {
	return &NullableLocalizedNotificationMessage{value: val, isSet: true}
}

func (v NullableLocalizedNotificationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalizedNotificationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


