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

// MicrosoftGraphVppToken struct for MicrosoftGraphVppToken
type MicrosoftGraphVppToken struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The apple Id associated with the given Apple Volume Purchase Program Token.
	AppleId NullableString `json:"appleId,omitempty"`
	// Whether or not apps for the VPP token will be automatically updated.
	AutomaticallyUpdateApps *bool `json:"automaticallyUpdateApps,omitempty"`
	// Whether or not apps for the VPP token will be automatically updated.
	CountryOrRegion NullableString `json:"countryOrRegion,omitempty"`
	// The expiration date time of the Apple Volume Purchase Program Token.
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// Last modification date time associated with the Apple Volume Purchase Program Token.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The last time when an application sync was done with the Apple volume purchase program service using the the Apple Volume Purchase Program Token.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// Current sync status of the last application sync which was triggered using the Apple Volume Purchase Program Token. Possible values are: none, inProgress, completed, failed. Possible values are: none, inProgress, completed, failed.
	LastSyncStatus AnyOfmicrosoftGraphVppTokenSyncStatus `json:"lastSyncStatus,omitempty"`
	// The organization associated with the Apple Volume Purchase Program Token
	OrganizationName NullableString `json:"organizationName,omitempty"`
	// Current state of the Apple Volume Purchase Program Token. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM. Possible values are: unknown, valid, expired, invalid, assignedToExternalMDM.
	State AnyOfmicrosoftGraphVppTokenState `json:"state,omitempty"`
	// The Apple Volume Purchase Program Token string downloaded from the Apple Volume Purchase Program.
	Token NullableString `json:"token,omitempty"`
	// The type of volume purchase program which the given Apple Volume Purchase Program Token is associated with. Possible values are: business, education. Possible values are: business, education.
	VppTokenAccountType AnyOfmicrosoftGraphVppTokenAccountType `json:"vppTokenAccountType,omitempty"`
}

// NewMicrosoftGraphVppToken instantiates a new MicrosoftGraphVppToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphVppToken() *MicrosoftGraphVppToken {
	this := MicrosoftGraphVppToken{}
	return &this
}

// NewMicrosoftGraphVppTokenWithDefaults instantiates a new MicrosoftGraphVppToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphVppTokenWithDefaults() *MicrosoftGraphVppToken {
	this := MicrosoftGraphVppToken{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphVppToken) SetId(v string) {
	o.Id = &v
}

// GetAppleId returns the AppleId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVppToken) GetAppleId() string {
	if o == nil || o.AppleId.Get() == nil {
		var ret string
		return ret
	}
	return *o.AppleId.Get()
}

// GetAppleIdOk returns a tuple with the AppleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVppToken) GetAppleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AppleId.Get(), o.AppleId.IsSet()
}

// HasAppleId returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasAppleId() bool {
	if o != nil && o.AppleId.IsSet() {
		return true
	}

	return false
}

// SetAppleId gets a reference to the given NullableString and assigns it to the AppleId field.
func (o *MicrosoftGraphVppToken) SetAppleId(v string) {
	o.AppleId.Set(&v)
}
// SetAppleIdNil sets the value for AppleId to be an explicit nil
func (o *MicrosoftGraphVppToken) SetAppleIdNil() {
	o.AppleId.Set(nil)
}

// UnsetAppleId ensures that no value is present for AppleId, not even an explicit nil
func (o *MicrosoftGraphVppToken) UnsetAppleId() {
	o.AppleId.Unset()
}

// GetAutomaticallyUpdateApps returns the AutomaticallyUpdateApps field value if set, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetAutomaticallyUpdateApps() bool {
	if o == nil || o.AutomaticallyUpdateApps == nil {
		var ret bool
		return ret
	}
	return *o.AutomaticallyUpdateApps
}

// GetAutomaticallyUpdateAppsOk returns a tuple with the AutomaticallyUpdateApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetAutomaticallyUpdateAppsOk() (*bool, bool) {
	if o == nil || o.AutomaticallyUpdateApps == nil {
		return nil, false
	}
	return o.AutomaticallyUpdateApps, true
}

// HasAutomaticallyUpdateApps returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasAutomaticallyUpdateApps() bool {
	if o != nil && o.AutomaticallyUpdateApps != nil {
		return true
	}

	return false
}

// SetAutomaticallyUpdateApps gets a reference to the given bool and assigns it to the AutomaticallyUpdateApps field.
func (o *MicrosoftGraphVppToken) SetAutomaticallyUpdateApps(v bool) {
	o.AutomaticallyUpdateApps = &v
}

// GetCountryOrRegion returns the CountryOrRegion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVppToken) GetCountryOrRegion() string {
	if o == nil || o.CountryOrRegion.Get() == nil {
		var ret string
		return ret
	}
	return *o.CountryOrRegion.Get()
}

// GetCountryOrRegionOk returns a tuple with the CountryOrRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVppToken) GetCountryOrRegionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CountryOrRegion.Get(), o.CountryOrRegion.IsSet()
}

// HasCountryOrRegion returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasCountryOrRegion() bool {
	if o != nil && o.CountryOrRegion.IsSet() {
		return true
	}

	return false
}

// SetCountryOrRegion gets a reference to the given NullableString and assigns it to the CountryOrRegion field.
func (o *MicrosoftGraphVppToken) SetCountryOrRegion(v string) {
	o.CountryOrRegion.Set(&v)
}
// SetCountryOrRegionNil sets the value for CountryOrRegion to be an explicit nil
func (o *MicrosoftGraphVppToken) SetCountryOrRegionNil() {
	o.CountryOrRegion.Set(nil)
}

// UnsetCountryOrRegion ensures that no value is present for CountryOrRegion, not even an explicit nil
func (o *MicrosoftGraphVppToken) UnsetCountryOrRegion() {
	o.CountryOrRegion.Unset()
}

// GetExpirationDateTime returns the ExpirationDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetExpirationDateTime() time.Time {
	if o == nil || o.ExpirationDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDateTime
}

// GetExpirationDateTimeOk returns a tuple with the ExpirationDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetExpirationDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ExpirationDateTime == nil {
		return nil, false
	}
	return o.ExpirationDateTime, true
}

// HasExpirationDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasExpirationDateTime() bool {
	if o != nil && o.ExpirationDateTime != nil {
		return true
	}

	return false
}

// SetExpirationDateTime gets a reference to the given time.Time and assigns it to the ExpirationDateTime field.
func (o *MicrosoftGraphVppToken) SetExpirationDateTime(v time.Time) {
	o.ExpirationDateTime = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphVppToken) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetLastSyncDateTime returns the LastSyncDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphVppToken) GetLastSyncDateTime() time.Time {
	if o == nil || o.LastSyncDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSyncDateTime
}

// GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphVppToken) GetLastSyncDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastSyncDateTime == nil {
		return nil, false
	}
	return o.LastSyncDateTime, true
}

// HasLastSyncDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasLastSyncDateTime() bool {
	if o != nil && o.LastSyncDateTime != nil {
		return true
	}

	return false
}

// SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.
func (o *MicrosoftGraphVppToken) SetLastSyncDateTime(v time.Time) {
	o.LastSyncDateTime = &v
}

// GetLastSyncStatus returns the LastSyncStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVppToken) GetLastSyncStatus() AnyOfmicrosoftGraphVppTokenSyncStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphVppTokenSyncStatus
		return ret
	}
	return o.LastSyncStatus
}

// GetLastSyncStatusOk returns a tuple with the LastSyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVppToken) GetLastSyncStatusOk() (*AnyOfmicrosoftGraphVppTokenSyncStatus, bool) {
	if o == nil || o.LastSyncStatus == nil {
		return nil, false
	}
	return &o.LastSyncStatus, true
}

// HasLastSyncStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasLastSyncStatus() bool {
	if o != nil && o.LastSyncStatus != nil {
		return true
	}

	return false
}

// SetLastSyncStatus gets a reference to the given AnyOfmicrosoftGraphVppTokenSyncStatus and assigns it to the LastSyncStatus field.
func (o *MicrosoftGraphVppToken) SetLastSyncStatus(v AnyOfmicrosoftGraphVppTokenSyncStatus) {
	o.LastSyncStatus = v
}

// GetOrganizationName returns the OrganizationName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVppToken) GetOrganizationName() string {
	if o == nil || o.OrganizationName.Get() == nil {
		var ret string
		return ret
	}
	return *o.OrganizationName.Get()
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVppToken) GetOrganizationNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OrganizationName.Get(), o.OrganizationName.IsSet()
}

// HasOrganizationName returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasOrganizationName() bool {
	if o != nil && o.OrganizationName.IsSet() {
		return true
	}

	return false
}

// SetOrganizationName gets a reference to the given NullableString and assigns it to the OrganizationName field.
func (o *MicrosoftGraphVppToken) SetOrganizationName(v string) {
	o.OrganizationName.Set(&v)
}
// SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil
func (o *MicrosoftGraphVppToken) SetOrganizationNameNil() {
	o.OrganizationName.Set(nil)
}

// UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
func (o *MicrosoftGraphVppToken) UnsetOrganizationName() {
	o.OrganizationName.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVppToken) GetState() AnyOfmicrosoftGraphVppTokenState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphVppTokenState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVppToken) GetStateOk() (*AnyOfmicrosoftGraphVppTokenState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphVppTokenState and assigns it to the State field.
func (o *MicrosoftGraphVppToken) SetState(v AnyOfmicrosoftGraphVppTokenState) {
	o.State = v
}

// GetToken returns the Token field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVppToken) GetToken() string {
	if o == nil || o.Token.Get() == nil {
		var ret string
		return ret
	}
	return *o.Token.Get()
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVppToken) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Token.Get(), o.Token.IsSet()
}

// HasToken returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasToken() bool {
	if o != nil && o.Token.IsSet() {
		return true
	}

	return false
}

// SetToken gets a reference to the given NullableString and assigns it to the Token field.
func (o *MicrosoftGraphVppToken) SetToken(v string) {
	o.Token.Set(&v)
}
// SetTokenNil sets the value for Token to be an explicit nil
func (o *MicrosoftGraphVppToken) SetTokenNil() {
	o.Token.Set(nil)
}

// UnsetToken ensures that no value is present for Token, not even an explicit nil
func (o *MicrosoftGraphVppToken) UnsetToken() {
	o.Token.Unset()
}

// GetVppTokenAccountType returns the VppTokenAccountType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphVppToken) GetVppTokenAccountType() AnyOfmicrosoftGraphVppTokenAccountType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphVppTokenAccountType
		return ret
	}
	return o.VppTokenAccountType
}

// GetVppTokenAccountTypeOk returns a tuple with the VppTokenAccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphVppToken) GetVppTokenAccountTypeOk() (*AnyOfmicrosoftGraphVppTokenAccountType, bool) {
	if o == nil || o.VppTokenAccountType == nil {
		return nil, false
	}
	return &o.VppTokenAccountType, true
}

// HasVppTokenAccountType returns a boolean if a field has been set.
func (o *MicrosoftGraphVppToken) HasVppTokenAccountType() bool {
	if o != nil && o.VppTokenAccountType != nil {
		return true
	}

	return false
}

// SetVppTokenAccountType gets a reference to the given AnyOfmicrosoftGraphVppTokenAccountType and assigns it to the VppTokenAccountType field.
func (o *MicrosoftGraphVppToken) SetVppTokenAccountType(v AnyOfmicrosoftGraphVppTokenAccountType) {
	o.VppTokenAccountType = v
}

func (o MicrosoftGraphVppToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AppleId.IsSet() {
		toSerialize["appleId"] = o.AppleId.Get()
	}
	if o.AutomaticallyUpdateApps != nil {
		toSerialize["automaticallyUpdateApps"] = o.AutomaticallyUpdateApps
	}
	if o.CountryOrRegion.IsSet() {
		toSerialize["countryOrRegion"] = o.CountryOrRegion.Get()
	}
	if o.ExpirationDateTime != nil {
		toSerialize["expirationDateTime"] = o.ExpirationDateTime
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.LastSyncDateTime != nil {
		toSerialize["lastSyncDateTime"] = o.LastSyncDateTime
	}
	if o.LastSyncStatus != nil {
		toSerialize["lastSyncStatus"] = o.LastSyncStatus
	}
	if o.OrganizationName.IsSet() {
		toSerialize["organizationName"] = o.OrganizationName.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Token.IsSet() {
		toSerialize["token"] = o.Token.Get()
	}
	if o.VppTokenAccountType != nil {
		toSerialize["vppTokenAccountType"] = o.VppTokenAccountType
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphVppToken struct {
	value *MicrosoftGraphVppToken
	isSet bool
}

func (v NullableMicrosoftGraphVppToken) Get() *MicrosoftGraphVppToken {
	return v.value
}

func (v *NullableMicrosoftGraphVppToken) Set(val *MicrosoftGraphVppToken) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphVppToken) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphVppToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphVppToken(val *MicrosoftGraphVppToken) *NullableMicrosoftGraphVppToken {
	return &NullableMicrosoftGraphVppToken{value: val, isSet: true}
}

func (v NullableMicrosoftGraphVppToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphVppToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


