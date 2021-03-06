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

// MicrosoftGraphLocalizedNotificationMessage struct for MicrosoftGraphLocalizedNotificationMessage
type MicrosoftGraphLocalizedNotificationMessage struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
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

// NewMicrosoftGraphLocalizedNotificationMessage instantiates a new MicrosoftGraphLocalizedNotificationMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphLocalizedNotificationMessage() *MicrosoftGraphLocalizedNotificationMessage {
	this := MicrosoftGraphLocalizedNotificationMessage{}
	return &this
}

// NewMicrosoftGraphLocalizedNotificationMessageWithDefaults instantiates a new MicrosoftGraphLocalizedNotificationMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphLocalizedNotificationMessageWithDefaults() *MicrosoftGraphLocalizedNotificationMessage {
	this := MicrosoftGraphLocalizedNotificationMessage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphLocalizedNotificationMessage) SetId(v string) {
	o.Id = &v
}

// GetIsDefault returns the IsDefault field value if set, zero value otherwise.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetIsDefault() bool {
	if o == nil || o.IsDefault == nil {
		var ret bool
		return ret
	}
	return *o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetIsDefaultOk() (*bool, bool) {
	if o == nil || o.IsDefault == nil {
		return nil, false
	}
	return o.IsDefault, true
}

// HasIsDefault returns a boolean if a field has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) HasIsDefault() bool {
	if o != nil && o.IsDefault != nil {
		return true
	}

	return false
}

// SetIsDefault gets a reference to the given bool and assigns it to the IsDefault field.
func (o *MicrosoftGraphLocalizedNotificationMessage) SetIsDefault(v bool) {
	o.IsDefault = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphLocalizedNotificationMessage) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetLocale() string {
	if o == nil || o.Locale == nil {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetLocaleOk() (*string, bool) {
	if o == nil || o.Locale == nil {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) HasLocale() bool {
	if o != nil && o.Locale != nil {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *MicrosoftGraphLocalizedNotificationMessage) SetLocale(v string) {
	o.Locale = &v
}

// GetMessageTemplate returns the MessageTemplate field value if set, zero value otherwise.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetMessageTemplate() string {
	if o == nil || o.MessageTemplate == nil {
		var ret string
		return ret
	}
	return *o.MessageTemplate
}

// GetMessageTemplateOk returns a tuple with the MessageTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetMessageTemplateOk() (*string, bool) {
	if o == nil || o.MessageTemplate == nil {
		return nil, false
	}
	return o.MessageTemplate, true
}

// HasMessageTemplate returns a boolean if a field has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) HasMessageTemplate() bool {
	if o != nil && o.MessageTemplate != nil {
		return true
	}

	return false
}

// SetMessageTemplate gets a reference to the given string and assigns it to the MessageTemplate field.
func (o *MicrosoftGraphLocalizedNotificationMessage) SetMessageTemplate(v string) {
	o.MessageTemplate = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *MicrosoftGraphLocalizedNotificationMessage) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *MicrosoftGraphLocalizedNotificationMessage) SetSubject(v string) {
	o.Subject = &v
}

func (o MicrosoftGraphLocalizedNotificationMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
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

type NullableMicrosoftGraphLocalizedNotificationMessage struct {
	value *MicrosoftGraphLocalizedNotificationMessage
	isSet bool
}

func (v NullableMicrosoftGraphLocalizedNotificationMessage) Get() *MicrosoftGraphLocalizedNotificationMessage {
	return v.value
}

func (v *NullableMicrosoftGraphLocalizedNotificationMessage) Set(val *MicrosoftGraphLocalizedNotificationMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphLocalizedNotificationMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphLocalizedNotificationMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphLocalizedNotificationMessage(val *MicrosoftGraphLocalizedNotificationMessage) *NullableMicrosoftGraphLocalizedNotificationMessage {
	return &NullableMicrosoftGraphLocalizedNotificationMessage{value: val, isSet: true}
}

func (v NullableMicrosoftGraphLocalizedNotificationMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphLocalizedNotificationMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


