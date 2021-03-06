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

// MicrosoftGraphParticipantInfo struct for MicrosoftGraphParticipantInfo
type MicrosoftGraphParticipantInfo struct {
	// The ISO 3166-1 Alpha-2 country code of the participant's best estimated physical location at the start of the call. Read-only.
	CountryCode NullableString `json:"countryCode,omitempty"`
	// The type of endpoint the participant is using. Possible values are: default, skypeForBusiness, or skypeForBusinessVoipPhone. Read-only.
	EndpointType AnyOfmicrosoftGraphEndpointType `json:"endpointType,omitempty"`
	Identity *MicrosoftGraphIdentitySet `json:"identity,omitempty"`
	// The language culture string. Read-only.
	LanguageId NullableString `json:"languageId,omitempty"`
	// The participant ID of the participant. Read-only.
	ParticipantId NullableString `json:"participantId,omitempty"`
	// The home region of the participant. This can be a country, a continent, or a larger geographic region. This does not change based on the participant's current physical location. Read-only.
	Region NullableString `json:"region,omitempty"`
}

// NewMicrosoftGraphParticipantInfo instantiates a new MicrosoftGraphParticipantInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphParticipantInfo() *MicrosoftGraphParticipantInfo {
	this := MicrosoftGraphParticipantInfo{}
	return &this
}

// NewMicrosoftGraphParticipantInfoWithDefaults instantiates a new MicrosoftGraphParticipantInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphParticipantInfoWithDefaults() *MicrosoftGraphParticipantInfo {
	this := MicrosoftGraphParticipantInfo{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphParticipantInfo) GetCountryCode() string {
	if o == nil || o.CountryCode.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryCode.Get()
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphParticipantInfo) GetCountryCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountryCode.Get(), o.CountryCode.IsSet()
}

// HasCountryCode returns a boolean if a field has been set.
func (o *MicrosoftGraphParticipantInfo) HasCountryCode() bool {
	if o != nil && o.CountryCode.IsSet() {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given NullableString and assigns it to the CountryCode field.
func (o *MicrosoftGraphParticipantInfo) SetCountryCode(v string) {
	o.CountryCode.Set(&v)
}
// SetCountryCodeNil sets the value for CountryCode to be an explicit nil
func (o *MicrosoftGraphParticipantInfo) SetCountryCodeNil() {
	o.CountryCode.Set(nil)
}

// UnsetCountryCode ensures that no value is present for CountryCode, not even an explicit nil
func (o *MicrosoftGraphParticipantInfo) UnsetCountryCode() {
	o.CountryCode.Unset()
}

// GetEndpointType returns the EndpointType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphParticipantInfo) GetEndpointType() AnyOfmicrosoftGraphEndpointType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEndpointType
		return ret
	}
	return o.EndpointType
}

// GetEndpointTypeOk returns a tuple with the EndpointType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphParticipantInfo) GetEndpointTypeOk() (*AnyOfmicrosoftGraphEndpointType, bool) {
	if o == nil || o.EndpointType == nil {
		return nil, false
	}
	return &o.EndpointType, true
}

// HasEndpointType returns a boolean if a field has been set.
func (o *MicrosoftGraphParticipantInfo) HasEndpointType() bool {
	if o != nil && o.EndpointType != nil {
		return true
	}

	return false
}

// SetEndpointType gets a reference to the given AnyOfmicrosoftGraphEndpointType and assigns it to the EndpointType field.
func (o *MicrosoftGraphParticipantInfo) SetEndpointType(v AnyOfmicrosoftGraphEndpointType) {
	o.EndpointType = v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *MicrosoftGraphParticipantInfo) GetIdentity() MicrosoftGraphIdentitySet {
	if o == nil || o.Identity == nil {
		var ret MicrosoftGraphIdentitySet
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphParticipantInfo) GetIdentityOk() (*MicrosoftGraphIdentitySet, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *MicrosoftGraphParticipantInfo) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given MicrosoftGraphIdentitySet and assigns it to the Identity field.
func (o *MicrosoftGraphParticipantInfo) SetIdentity(v MicrosoftGraphIdentitySet) {
	o.Identity = &v
}

// GetLanguageId returns the LanguageId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphParticipantInfo) GetLanguageId() string {
	if o == nil || o.LanguageId.Get() == nil {
		var ret string
		return ret
	}
	return *o.LanguageId.Get()
}

// GetLanguageIdOk returns a tuple with the LanguageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphParticipantInfo) GetLanguageIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LanguageId.Get(), o.LanguageId.IsSet()
}

// HasLanguageId returns a boolean if a field has been set.
func (o *MicrosoftGraphParticipantInfo) HasLanguageId() bool {
	if o != nil && o.LanguageId.IsSet() {
		return true
	}

	return false
}

// SetLanguageId gets a reference to the given NullableString and assigns it to the LanguageId field.
func (o *MicrosoftGraphParticipantInfo) SetLanguageId(v string) {
	o.LanguageId.Set(&v)
}
// SetLanguageIdNil sets the value for LanguageId to be an explicit nil
func (o *MicrosoftGraphParticipantInfo) SetLanguageIdNil() {
	o.LanguageId.Set(nil)
}

// UnsetLanguageId ensures that no value is present for LanguageId, not even an explicit nil
func (o *MicrosoftGraphParticipantInfo) UnsetLanguageId() {
	o.LanguageId.Unset()
}

// GetParticipantId returns the ParticipantId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphParticipantInfo) GetParticipantId() string {
	if o == nil || o.ParticipantId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParticipantId.Get()
}

// GetParticipantIdOk returns a tuple with the ParticipantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphParticipantInfo) GetParticipantIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParticipantId.Get(), o.ParticipantId.IsSet()
}

// HasParticipantId returns a boolean if a field has been set.
func (o *MicrosoftGraphParticipantInfo) HasParticipantId() bool {
	if o != nil && o.ParticipantId.IsSet() {
		return true
	}

	return false
}

// SetParticipantId gets a reference to the given NullableString and assigns it to the ParticipantId field.
func (o *MicrosoftGraphParticipantInfo) SetParticipantId(v string) {
	o.ParticipantId.Set(&v)
}
// SetParticipantIdNil sets the value for ParticipantId to be an explicit nil
func (o *MicrosoftGraphParticipantInfo) SetParticipantIdNil() {
	o.ParticipantId.Set(nil)
}

// UnsetParticipantId ensures that no value is present for ParticipantId, not even an explicit nil
func (o *MicrosoftGraphParticipantInfo) UnsetParticipantId() {
	o.ParticipantId.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphParticipantInfo) GetRegion() string {
	if o == nil || o.Region.Get() == nil {
		var ret string
		return ret
	}
	return *o.Region.Get()
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphParticipantInfo) GetRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Region.Get(), o.Region.IsSet()
}

// HasRegion returns a boolean if a field has been set.
func (o *MicrosoftGraphParticipantInfo) HasRegion() bool {
	if o != nil && o.Region.IsSet() {
		return true
	}

	return false
}

// SetRegion gets a reference to the given NullableString and assigns it to the Region field.
func (o *MicrosoftGraphParticipantInfo) SetRegion(v string) {
	o.Region.Set(&v)
}
// SetRegionNil sets the value for Region to be an explicit nil
func (o *MicrosoftGraphParticipantInfo) SetRegionNil() {
	o.Region.Set(nil)
}

// UnsetRegion ensures that no value is present for Region, not even an explicit nil
func (o *MicrosoftGraphParticipantInfo) UnsetRegion() {
	o.Region.Unset()
}

func (o MicrosoftGraphParticipantInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CountryCode.IsSet() {
		toSerialize["countryCode"] = o.CountryCode.Get()
	}
	if o.EndpointType != nil {
		toSerialize["endpointType"] = o.EndpointType
	}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.LanguageId.IsSet() {
		toSerialize["languageId"] = o.LanguageId.Get()
	}
	if o.ParticipantId.IsSet() {
		toSerialize["participantId"] = o.ParticipantId.Get()
	}
	if o.Region.IsSet() {
		toSerialize["region"] = o.Region.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphParticipantInfo struct {
	value *MicrosoftGraphParticipantInfo
	isSet bool
}

func (v NullableMicrosoftGraphParticipantInfo) Get() *MicrosoftGraphParticipantInfo {
	return v.value
}

func (v *NullableMicrosoftGraphParticipantInfo) Set(val *MicrosoftGraphParticipantInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphParticipantInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphParticipantInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphParticipantInfo(val *MicrosoftGraphParticipantInfo) *NullableMicrosoftGraphParticipantInfo {
	return &NullableMicrosoftGraphParticipantInfo{value: val, isSet: true}
}

func (v NullableMicrosoftGraphParticipantInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphParticipantInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


