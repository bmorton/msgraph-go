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

// Security struct for Security
type Security struct {
	// Read-only. Nullable.
	Alerts *[]MicrosoftGraphAlert `json:"alerts,omitempty"`
	SecureScoreControlProfiles *[]MicrosoftGraphSecureScoreControlProfile `json:"secureScoreControlProfiles,omitempty"`
	SecureScores *[]MicrosoftGraphSecureScore `json:"secureScores,omitempty"`
}

// NewSecurity instantiates a new Security object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurity() *Security {
	this := Security{}
	return &this
}

// NewSecurityWithDefaults instantiates a new Security object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityWithDefaults() *Security {
	this := Security{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *Security) GetAlerts() []MicrosoftGraphAlert {
	if o == nil || o.Alerts == nil {
		var ret []MicrosoftGraphAlert
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetAlertsOk() (*[]MicrosoftGraphAlert, bool) {
	if o == nil || o.Alerts == nil {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *Security) HasAlerts() bool {
	if o != nil && o.Alerts != nil {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []MicrosoftGraphAlert and assigns it to the Alerts field.
func (o *Security) SetAlerts(v []MicrosoftGraphAlert) {
	o.Alerts = &v
}

// GetSecureScoreControlProfiles returns the SecureScoreControlProfiles field value if set, zero value otherwise.
func (o *Security) GetSecureScoreControlProfiles() []MicrosoftGraphSecureScoreControlProfile {
	if o == nil || o.SecureScoreControlProfiles == nil {
		var ret []MicrosoftGraphSecureScoreControlProfile
		return ret
	}
	return *o.SecureScoreControlProfiles
}

// GetSecureScoreControlProfilesOk returns a tuple with the SecureScoreControlProfiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetSecureScoreControlProfilesOk() (*[]MicrosoftGraphSecureScoreControlProfile, bool) {
	if o == nil || o.SecureScoreControlProfiles == nil {
		return nil, false
	}
	return o.SecureScoreControlProfiles, true
}

// HasSecureScoreControlProfiles returns a boolean if a field has been set.
func (o *Security) HasSecureScoreControlProfiles() bool {
	if o != nil && o.SecureScoreControlProfiles != nil {
		return true
	}

	return false
}

// SetSecureScoreControlProfiles gets a reference to the given []MicrosoftGraphSecureScoreControlProfile and assigns it to the SecureScoreControlProfiles field.
func (o *Security) SetSecureScoreControlProfiles(v []MicrosoftGraphSecureScoreControlProfile) {
	o.SecureScoreControlProfiles = &v
}

// GetSecureScores returns the SecureScores field value if set, zero value otherwise.
func (o *Security) GetSecureScores() []MicrosoftGraphSecureScore {
	if o == nil || o.SecureScores == nil {
		var ret []MicrosoftGraphSecureScore
		return ret
	}
	return *o.SecureScores
}

// GetSecureScoresOk returns a tuple with the SecureScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Security) GetSecureScoresOk() (*[]MicrosoftGraphSecureScore, bool) {
	if o == nil || o.SecureScores == nil {
		return nil, false
	}
	return o.SecureScores, true
}

// HasSecureScores returns a boolean if a field has been set.
func (o *Security) HasSecureScores() bool {
	if o != nil && o.SecureScores != nil {
		return true
	}

	return false
}

// SetSecureScores gets a reference to the given []MicrosoftGraphSecureScore and assigns it to the SecureScores field.
func (o *Security) SetSecureScores(v []MicrosoftGraphSecureScore) {
	o.SecureScores = &v
}

func (o Security) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alerts != nil {
		toSerialize["alerts"] = o.Alerts
	}
	if o.SecureScoreControlProfiles != nil {
		toSerialize["secureScoreControlProfiles"] = o.SecureScoreControlProfiles
	}
	if o.SecureScores != nil {
		toSerialize["secureScores"] = o.SecureScores
	}
	return json.Marshal(toSerialize)
}

type NullableSecurity struct {
	value *Security
	isSet bool
}

func (v NullableSecurity) Get() *Security {
	return v.value
}

func (v *NullableSecurity) Set(val *Security) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurity) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurity(val *Security) *NullableSecurity {
	return &NullableSecurity{value: val, isSet: true}
}

func (v NullableSecurity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


