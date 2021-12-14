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

// MicrosoftGraphApplePushNotificationCertificate struct for MicrosoftGraphApplePushNotificationCertificate
type MicrosoftGraphApplePushNotificationCertificate struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Apple Id of the account used to create the MDM push certificate.
	AppleIdentifier NullableString `json:"appleIdentifier,omitempty"`
	// Not yet documented
	Certificate NullableString `json:"certificate,omitempty"`
	// Certificate serial number. This property is read-only.
	CertificateSerialNumber NullableString `json:"certificateSerialNumber,omitempty"`
	// The expiration date and time for Apple push notification certificate.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Last modified date and time for Apple push notification certificate.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Topic Id.
	TopicIdentifier NullableString `json:"topicIdentifier,omitempty"`
}

// NewMicrosoftGraphApplePushNotificationCertificate instantiates a new MicrosoftGraphApplePushNotificationCertificate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphApplePushNotificationCertificate() *MicrosoftGraphApplePushNotificationCertificate {
	this := MicrosoftGraphApplePushNotificationCertificate{}
	return &this
}

// NewMicrosoftGraphApplePushNotificationCertificateWithDefaults instantiates a new MicrosoftGraphApplePushNotificationCertificate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphApplePushNotificationCertificateWithDefaults() *MicrosoftGraphApplePushNotificationCertificate {
	this := MicrosoftGraphApplePushNotificationCertificate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphApplePushNotificationCertificate) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphApplePushNotificationCertificate) SetId(v string) {
	o.Id = &v
}

// GetAppleIdentifier returns the AppleIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphApplePushNotificationCertificate) GetAppleIdentifier() string {
	if o == nil || o.AppleIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppleIdentifier.Get()
}

// GetAppleIdentifierOk returns a tuple with the AppleIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphApplePushNotificationCertificate) GetAppleIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppleIdentifier.Get(), o.AppleIdentifier.IsSet()
}

// HasAppleIdentifier returns a boolean if a field has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) HasAppleIdentifier() bool {
	if o != nil && o.AppleIdentifier.IsSet() {
		return true
	}

	return false
}

// SetAppleIdentifier gets a reference to the given NullableString and assigns it to the AppleIdentifier field.
func (o *MicrosoftGraphApplePushNotificationCertificate) SetAppleIdentifier(v string) {
	o.AppleIdentifier.Set(&v)
}
// SetAppleIdentifierNil sets the value for AppleIdentifier to be an explicit nil
func (o *MicrosoftGraphApplePushNotificationCertificate) SetAppleIdentifierNil() {
	o.AppleIdentifier.Set(nil)
}

// UnsetAppleIdentifier ensures that no value is present for AppleIdentifier, not even an explicit nil
func (o *MicrosoftGraphApplePushNotificationCertificate) UnsetAppleIdentifier() {
	o.AppleIdentifier.Unset()
}

// GetCertificate returns the Certificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphApplePushNotificationCertificate) GetCertificate() string {
	if o == nil || o.Certificate.Get() == nil {
		var ret string
		return ret
	}
	return *o.Certificate.Get()
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphApplePushNotificationCertificate) GetCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Certificate.Get(), o.Certificate.IsSet()
}

// HasCertificate returns a boolean if a field has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) HasCertificate() bool {
	if o != nil && o.Certificate.IsSet() {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given NullableString and assigns it to the Certificate field.
func (o *MicrosoftGraphApplePushNotificationCertificate) SetCertificate(v string) {
	o.Certificate.Set(&v)
}
// SetCertificateNil sets the value for Certificate to be an explicit nil
func (o *MicrosoftGraphApplePushNotificationCertificate) SetCertificateNil() {
	o.Certificate.Set(nil)
}

// UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
func (o *MicrosoftGraphApplePushNotificationCertificate) UnsetCertificate() {
	o.Certificate.Unset()
}

// GetCertificateSerialNumber returns the CertificateSerialNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphApplePushNotificationCertificate) GetCertificateSerialNumber() string {
	if o == nil || o.CertificateSerialNumber.Get() == nil {
		var ret string
		return ret
	}
	return *o.CertificateSerialNumber.Get()
}

// GetCertificateSerialNumberOk returns a tuple with the CertificateSerialNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphApplePushNotificationCertificate) GetCertificateSerialNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CertificateSerialNumber.Get(), o.CertificateSerialNumber.IsSet()
}

// HasCertificateSerialNumber returns a boolean if a field has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) HasCertificateSerialNumber() bool {
	if o != nil && o.CertificateSerialNumber.IsSet() {
		return true
	}

	return false
}

// SetCertificateSerialNumber gets a reference to the given NullableString and assigns it to the CertificateSerialNumber field.
func (o *MicrosoftGraphApplePushNotificationCertificate) SetCertificateSerialNumber(v string) {
	o.CertificateSerialNumber.Set(&v)
}
// SetCertificateSerialNumberNil sets the value for CertificateSerialNumber to be an explicit nil
func (o *MicrosoftGraphApplePushNotificationCertificate) SetCertificateSerialNumberNil() {
	o.CertificateSerialNumber.Set(nil)
}

// UnsetCertificateSerialNumber ensures that no value is present for CertificateSerialNumber, not even an explicit nil
func (o *MicrosoftGraphApplePushNotificationCertificate) UnsetCertificateSerialNumber() {
	o.CertificateSerialNumber.Unset()
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphApplePushNotificationCertificate) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDateTime == nil {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime != nil {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *MicrosoftGraphApplePushNotificationCertificate) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphApplePushNotificationCertificate) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphApplePushNotificationCertificate) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetTopicIdentifier returns the TopicIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphApplePushNotificationCertificate) GetTopicIdentifier() string {
	if o == nil || o.TopicIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.TopicIdentifier.Get()
}

// GetTopicIdentifierOk returns a tuple with the TopicIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphApplePushNotificationCertificate) GetTopicIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TopicIdentifier.Get(), o.TopicIdentifier.IsSet()
}

// HasTopicIdentifier returns a boolean if a field has been set.
func (o *MicrosoftGraphApplePushNotificationCertificate) HasTopicIdentifier() bool {
	if o != nil && o.TopicIdentifier.IsSet() {
		return true
	}

	return false
}

// SetTopicIdentifier gets a reference to the given NullableString and assigns it to the TopicIdentifier field.
func (o *MicrosoftGraphApplePushNotificationCertificate) SetTopicIdentifier(v string) {
	o.TopicIdentifier.Set(&v)
}
// SetTopicIdentifierNil sets the value for TopicIdentifier to be an explicit nil
func (o *MicrosoftGraphApplePushNotificationCertificate) SetTopicIdentifierNil() {
	o.TopicIdentifier.Set(nil)
}

// UnsetTopicIdentifier ensures that no value is present for TopicIdentifier, not even an explicit nil
func (o *MicrosoftGraphApplePushNotificationCertificate) UnsetTopicIdentifier() {
	o.TopicIdentifier.Unset()
}

func (o MicrosoftGraphApplePushNotificationCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AppleIdentifier.IsSet() {
		toSerialize["appleIdentifier"] = o.AppleIdentifier.Get()
	}
	if o.Certificate.IsSet() {
		toSerialize["certificate"] = o.Certificate.Get()
	}
	if o.CertificateSerialNumber.IsSet() {
		toSerialize["certificateSerialNumber"] = o.CertificateSerialNumber.Get()
	}
	if o.ExpirationDateTime != nil {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.TopicIdentifier.IsSet() {
		toSerialize["topicIdentifier"] = o.TopicIdentifier.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphApplePushNotificationCertificate struct {
	value *MicrosoftGraphApplePushNotificationCertificate
	isSet bool
}

func (v NullableMicrosoftGraphApplePushNotificationCertificate) Get() *MicrosoftGraphApplePushNotificationCertificate {
	return v.value
}

func (v *NullableMicrosoftGraphApplePushNotificationCertificate) Set(val *MicrosoftGraphApplePushNotificationCertificate) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphApplePushNotificationCertificate) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphApplePushNotificationCertificate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphApplePushNotificationCertificate(val *MicrosoftGraphApplePushNotificationCertificate) *NullableMicrosoftGraphApplePushNotificationCertificate {
	return &NullableMicrosoftGraphApplePushNotificationCertificate{value: val, isSet: true}
}

func (v NullableMicrosoftGraphApplePushNotificationCertificate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphApplePushNotificationCertificate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


