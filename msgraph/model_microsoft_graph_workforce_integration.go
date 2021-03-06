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

// MicrosoftGraphWorkforceIntegration struct for MicrosoftGraphWorkforceIntegration
type MicrosoftGraphWorkforceIntegration struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Identity of the person who last modified the entity.
	LastModifiedBy AnyOfmicrosoftGraphIdentitySet `json:"lastModifiedBy,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// API version for the call back URL. Start with 1.
	ApiVersion NullableInt32 `json:"apiVersion,omitempty"`
	// Name of the workforce integration.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The workforce integration encryption resource.
	Encryption AnyOfmicrosoftGraphWorkforceIntegrationEncryption `json:"encryption,omitempty"`
	// Indicates whether this workforce integration is currently active and available.
	IsActive NullableBool `json:"isActive,omitempty"`
	// The Shifts entities supported for synchronous change notifications. Shifts will make a call back to the url provided on client changes on those entities added here. By default, no entities are supported for change notifications. Possible values are: none, shift, swapRequest, userShiftPreferences, openshift, openShiftRequest, offerShiftRequest, unknownFutureValue.
	SupportedEntities AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities `json:"supportedEntities,omitempty"`
	// Workforce Integration URL for callbacks from the Shifts service.
	Url NullableString `json:"url,omitempty"`
}

// NewMicrosoftGraphWorkforceIntegration instantiates a new MicrosoftGraphWorkforceIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkforceIntegration() *MicrosoftGraphWorkforceIntegration {
	this := MicrosoftGraphWorkforceIntegration{}
	return &this
}

// NewMicrosoftGraphWorkforceIntegrationWithDefaults instantiates a new MicrosoftGraphWorkforceIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkforceIntegrationWithDefaults() *MicrosoftGraphWorkforceIntegration {
	this := MicrosoftGraphWorkforceIntegration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkforceIntegration) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkforceIntegration) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkforceIntegration) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphWorkforceIntegration) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetLastModifiedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetLastModifiedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return &o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the LastModifiedBy field.
func (o *MicrosoftGraphWorkforceIntegration) SetLastModifiedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.LastModifiedBy = v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphWorkforceIntegration) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetApiVersion() int32 {
	if o == nil || o.ApiVersion.Get() == nil {
		var ret int32
		return ret
	}
	return *o.ApiVersion.Get()
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetApiVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApiVersion.Get(), o.ApiVersion.IsSet()
}

// HasApiVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasApiVersion() bool {
	if o != nil && o.ApiVersion.IsSet() {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given NullableInt32 and assigns it to the ApiVersion field.
func (o *MicrosoftGraphWorkforceIntegration) SetApiVersion(v int32) {
	o.ApiVersion.Set(&v)
}
// SetApiVersionNil sets the value for ApiVersion to be an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) SetApiVersionNil() {
	o.ApiVersion.Set(nil)
}

// UnsetApiVersion ensures that no value is present for ApiVersion, not even an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) UnsetApiVersion() {
	o.ApiVersion.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphWorkforceIntegration) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetEncryption returns the Encryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetEncryption() AnyOfmicrosoftGraphWorkforceIntegrationEncryption {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkforceIntegrationEncryption
		return ret
	}
	return o.Encryption
}

// GetEncryptionOk returns a tuple with the Encryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetEncryptionOk() (*AnyOfmicrosoftGraphWorkforceIntegrationEncryption, bool) {
	if o == nil || o.Encryption == nil {
		return nil, false
	}
	return &o.Encryption, true
}

// HasEncryption returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasEncryption() bool {
	if o != nil && o.Encryption != nil {
		return true
	}

	return false
}

// SetEncryption gets a reference to the given AnyOfmicrosoftGraphWorkforceIntegrationEncryption and assigns it to the Encryption field.
func (o *MicrosoftGraphWorkforceIntegration) SetEncryption(v AnyOfmicrosoftGraphWorkforceIntegrationEncryption) {
	o.Encryption = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetIsActive() bool {
	if o == nil || o.IsActive.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsActive.Get()
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetIsActiveOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsActive.Get(), o.IsActive.IsSet()
}

// HasIsActive returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasIsActive() bool {
	if o != nil && o.IsActive.IsSet() {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given NullableBool and assigns it to the IsActive field.
func (o *MicrosoftGraphWorkforceIntegration) SetIsActive(v bool) {
	o.IsActive.Set(&v)
}
// SetIsActiveNil sets the value for IsActive to be an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) SetIsActiveNil() {
	o.IsActive.Set(nil)
}

// UnsetIsActive ensures that no value is present for IsActive, not even an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) UnsetIsActive() {
	o.IsActive.Unset()
}

// GetSupportedEntities returns the SupportedEntities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetSupportedEntities() AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities
		return ret
	}
	return o.SupportedEntities
}

// GetSupportedEntitiesOk returns a tuple with the SupportedEntities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetSupportedEntitiesOk() (*AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities, bool) {
	if o == nil || o.SupportedEntities == nil {
		return nil, false
	}
	return &o.SupportedEntities, true
}

// HasSupportedEntities returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasSupportedEntities() bool {
	if o != nil && o.SupportedEntities != nil {
		return true
	}

	return false
}

// SetSupportedEntities gets a reference to the given AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities and assigns it to the SupportedEntities field.
func (o *MicrosoftGraphWorkforceIntegration) SetSupportedEntities(v AnyOfmicrosoftGraphWorkforceIntegrationSupportedEntities) {
	o.SupportedEntities = v
}

// GetUrl returns the Url field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkforceIntegration) GetUrl() string {
	if o == nil || o.Url.Get() == nil {
		var ret string
		return ret
	}
	return *o.Url.Get()
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkforceIntegration) GetUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Url.Get(), o.Url.IsSet()
}

// HasUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkforceIntegration) HasUrl() bool {
	if o != nil && o.Url.IsSet() {
		return true
	}

	return false
}

// SetUrl gets a reference to the given NullableString and assigns it to the Url field.
func (o *MicrosoftGraphWorkforceIntegration) SetUrl(v string) {
	o.Url.Set(&v)
}
// SetUrlNil sets the value for Url to be an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) SetUrlNil() {
	o.Url.Set(nil)
}

// UnsetUrl ensures that no value is present for Url, not even an explicit nil
func (o *MicrosoftGraphWorkforceIntegration) UnsetUrl() {
	o.Url.Unset()
}

func (o MicrosoftGraphWorkforceIntegration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.ApiVersion.IsSet() {
		toSerialize["apiVersion"] = o.ApiVersion.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Encryption != nil {
		toSerialize["encryption"] = o.Encryption
	}
	if o.IsActive.IsSet() {
		toSerialize["isActive"] = o.IsActive.Get()
	}
	if o.SupportedEntities != nil {
		toSerialize["supportedEntities"] = o.SupportedEntities
	}
	if o.Url.IsSet() {
		toSerialize["url"] = o.Url.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkforceIntegration struct {
	value *MicrosoftGraphWorkforceIntegration
	isSet bool
}

func (v NullableMicrosoftGraphWorkforceIntegration) Get() *MicrosoftGraphWorkforceIntegration {
	return v.value
}

func (v *NullableMicrosoftGraphWorkforceIntegration) Set(val *MicrosoftGraphWorkforceIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkforceIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkforceIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkforceIntegration(val *MicrosoftGraphWorkforceIntegration) *NullableMicrosoftGraphWorkforceIntegration {
	return &NullableMicrosoftGraphWorkforceIntegration{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkforceIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkforceIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


