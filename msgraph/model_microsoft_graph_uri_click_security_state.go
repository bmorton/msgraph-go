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

// MicrosoftGraphUriClickSecurityState struct for MicrosoftGraphUriClickSecurityState
type MicrosoftGraphUriClickSecurityState struct {
	ClickAction NullableString `json:"clickAction,omitempty"`
	ClickDateTime NullableTime `json:"clickDateTime,omitempty"`
	Id NullableString `json:"id,omitempty"`
	SourceId NullableString `json:"sourceId,omitempty"`
	UriDomain NullableString `json:"uriDomain,omitempty"`
	Verdict NullableString `json:"verdict,omitempty"`
}

// NewMicrosoftGraphUriClickSecurityState instantiates a new MicrosoftGraphUriClickSecurityState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUriClickSecurityState() *MicrosoftGraphUriClickSecurityState {
	this := MicrosoftGraphUriClickSecurityState{}
	return &this
}

// NewMicrosoftGraphUriClickSecurityStateWithDefaults instantiates a new MicrosoftGraphUriClickSecurityState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUriClickSecurityStateWithDefaults() *MicrosoftGraphUriClickSecurityState {
	this := MicrosoftGraphUriClickSecurityState{}
	return &this
}

// GetClickAction returns the ClickAction field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUriClickSecurityState) GetClickAction() string {
	if o == nil || o.ClickAction.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClickAction.Get()
}

// GetClickActionOk returns a tuple with the ClickAction field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUriClickSecurityState) GetClickActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClickAction.Get(), o.ClickAction.IsSet()
}

// HasClickAction returns a boolean if a field has been set.
func (o *MicrosoftGraphUriClickSecurityState) HasClickAction() bool {
	if o != nil && o.ClickAction.IsSet() {
		return true
	}

	return false
}

// SetClickAction gets a reference to the given NullableString and assigns it to the ClickAction field.
func (o *MicrosoftGraphUriClickSecurityState) SetClickAction(v string) {
	o.ClickAction.Set(&v)
}
// SetClickActionNil sets the value for ClickAction to be an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) SetClickActionNil() {
	o.ClickAction.Set(nil)
}

// UnsetClickAction ensures that no value is present for ClickAction, not even an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) UnsetClickAction() {
	o.ClickAction.Unset()
}

// GetClickDateTime returns the ClickDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUriClickSecurityState) GetClickDateTime() time.Time {
	if o == nil || o.ClickDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ClickDateTime.Get()
}

// GetClickDateTimeOk returns a tuple with the ClickDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUriClickSecurityState) GetClickDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClickDateTime.Get(), o.ClickDateTime.IsSet()
}

// HasClickDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphUriClickSecurityState) HasClickDateTime() bool {
	if o != nil && o.ClickDateTime.IsSet() {
		return true
	}

	return false
}

// SetClickDateTime gets a reference to the given NullableTime and assigns it to the ClickDateTime field.
func (o *MicrosoftGraphUriClickSecurityState) SetClickDateTime(v time.Time) {
	o.ClickDateTime.Set(&v)
}
// SetClickDateTimeNil sets the value for ClickDateTime to be an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) SetClickDateTimeNil() {
	o.ClickDateTime.Set(nil)
}

// UnsetClickDateTime ensures that no value is present for ClickDateTime, not even an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) UnsetClickDateTime() {
	o.ClickDateTime.Unset()
}

// GetId returns the Id field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUriClickSecurityState) GetId() string {
	if o == nil || o.Id.Get() == nil {
		var ret string
		return ret
	}
	return *o.Id.Get()
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUriClickSecurityState) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Id.Get(), o.Id.IsSet()
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUriClickSecurityState) HasId() bool {
	if o != nil && o.Id.IsSet() {
		return true
	}

	return false
}

// SetId gets a reference to the given NullableString and assigns it to the Id field.
func (o *MicrosoftGraphUriClickSecurityState) SetId(v string) {
	o.Id.Set(&v)
}
// SetIdNil sets the value for Id to be an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) SetIdNil() {
	o.Id.Set(nil)
}

// UnsetId ensures that no value is present for Id, not even an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) UnsetId() {
	o.Id.Unset()
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUriClickSecurityState) GetSourceId() string {
	if o == nil || o.SourceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUriClickSecurityState) GetSourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *MicrosoftGraphUriClickSecurityState) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableString and assigns it to the SourceId field.
func (o *MicrosoftGraphUriClickSecurityState) SetSourceId(v string) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetUriDomain returns the UriDomain field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUriClickSecurityState) GetUriDomain() string {
	if o == nil || o.UriDomain.Get() == nil {
		var ret string
		return ret
	}
	return *o.UriDomain.Get()
}

// GetUriDomainOk returns a tuple with the UriDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUriClickSecurityState) GetUriDomainOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UriDomain.Get(), o.UriDomain.IsSet()
}

// HasUriDomain returns a boolean if a field has been set.
func (o *MicrosoftGraphUriClickSecurityState) HasUriDomain() bool {
	if o != nil && o.UriDomain.IsSet() {
		return true
	}

	return false
}

// SetUriDomain gets a reference to the given NullableString and assigns it to the UriDomain field.
func (o *MicrosoftGraphUriClickSecurityState) SetUriDomain(v string) {
	o.UriDomain.Set(&v)
}
// SetUriDomainNil sets the value for UriDomain to be an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) SetUriDomainNil() {
	o.UriDomain.Set(nil)
}

// UnsetUriDomain ensures that no value is present for UriDomain, not even an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) UnsetUriDomain() {
	o.UriDomain.Unset()
}

// GetVerdict returns the Verdict field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUriClickSecurityState) GetVerdict() string {
	if o == nil || o.Verdict.Get() == nil {
		var ret string
		return ret
	}
	return *o.Verdict.Get()
}

// GetVerdictOk returns a tuple with the Verdict field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUriClickSecurityState) GetVerdictOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Verdict.Get(), o.Verdict.IsSet()
}

// HasVerdict returns a boolean if a field has been set.
func (o *MicrosoftGraphUriClickSecurityState) HasVerdict() bool {
	if o != nil && o.Verdict.IsSet() {
		return true
	}

	return false
}

// SetVerdict gets a reference to the given NullableString and assigns it to the Verdict field.
func (o *MicrosoftGraphUriClickSecurityState) SetVerdict(v string) {
	o.Verdict.Set(&v)
}
// SetVerdictNil sets the value for Verdict to be an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) SetVerdictNil() {
	o.Verdict.Set(nil)
}

// UnsetVerdict ensures that no value is present for Verdict, not even an explicit nil
func (o *MicrosoftGraphUriClickSecurityState) UnsetVerdict() {
	o.Verdict.Unset()
}

func (o MicrosoftGraphUriClickSecurityState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClickAction.IsSet() {
		toSerialize["clickAction"] = o.ClickAction.Get()
	}
	if o.ClickDateTime.IsSet() {
		toSerialize["clickDateTime"] = o.ClickDateTime.Get()
	}
	if o.Id.IsSet() {
		toSerialize["id"] = o.Id.Get()
	}
	if o.SourceId.IsSet() {
		toSerialize["sourceId"] = o.SourceId.Get()
	}
	if o.UriDomain.IsSet() {
		toSerialize["uriDomain"] = o.UriDomain.Get()
	}
	if o.Verdict.IsSet() {
		toSerialize["verdict"] = o.Verdict.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUriClickSecurityState struct {
	value *MicrosoftGraphUriClickSecurityState
	isSet bool
}

func (v NullableMicrosoftGraphUriClickSecurityState) Get() *MicrosoftGraphUriClickSecurityState {
	return v.value
}

func (v *NullableMicrosoftGraphUriClickSecurityState) Set(val *MicrosoftGraphUriClickSecurityState) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUriClickSecurityState) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUriClickSecurityState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUriClickSecurityState(val *MicrosoftGraphUriClickSecurityState) *NullableMicrosoftGraphUriClickSecurityState {
	return &NullableMicrosoftGraphUriClickSecurityState{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUriClickSecurityState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUriClickSecurityState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


