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

// MicrosoftGraphConnectedOrganization struct for MicrosoftGraphConnectedOrganization
type MicrosoftGraphConnectedOrganization struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// The description of the connected organization.
	Description NullableString `json:"description,omitempty"`
	// The display name of the connected organization.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The identity sources in this connected organization, one of azureActiveDirectoryTenant, domainIdentitySource or externalDomainFederation. Nullable.
	IdentitySources *[]*AnyOfobject `json:"identitySources,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	ModifiedDateTime NullableTime `json:"modifiedDateTime,omitempty"`
	// The state of a connected organization defines whether assignment policies with requestor scope type AllConfiguredConnectedOrganizationSubjects are applicable or not.  The possible values are: configured, proposed, unknownFutureValue.
	State AnyOfmicrosoftGraphConnectedOrganizationState `json:"state,omitempty"`
	// Nullable.
	ExternalSponsors *[]MicrosoftGraphDirectoryObject `json:"externalSponsors,omitempty"`
	// Nullable.
	InternalSponsors *[]MicrosoftGraphDirectoryObject `json:"internalSponsors,omitempty"`
}

// NewMicrosoftGraphConnectedOrganization instantiates a new MicrosoftGraphConnectedOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphConnectedOrganization() *MicrosoftGraphConnectedOrganization {
	this := MicrosoftGraphConnectedOrganization{}
	return &this
}

// NewMicrosoftGraphConnectedOrganizationWithDefaults instantiates a new MicrosoftGraphConnectedOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphConnectedOrganizationWithDefaults() *MicrosoftGraphConnectedOrganization {
	this := MicrosoftGraphConnectedOrganization{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphConnectedOrganization) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConnectedOrganization) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphConnectedOrganization) SetId(v string) {
	o.Id = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConnectedOrganization) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConnectedOrganization) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphConnectedOrganization) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphConnectedOrganization) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphConnectedOrganization) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConnectedOrganization) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConnectedOrganization) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphConnectedOrganization) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphConnectedOrganization) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphConnectedOrganization) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConnectedOrganization) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConnectedOrganization) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphConnectedOrganization) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphConnectedOrganization) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphConnectedOrganization) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetIdentitySources returns the IdentitySources field value if set, zero value otherwise.
func (o *MicrosoftGraphConnectedOrganization) GetIdentitySources() []*AnyOfobject {
	if o == nil || o.IdentitySources == nil {
		var ret []*AnyOfobject
		return ret
	}
	return *o.IdentitySources
}

// GetIdentitySourcesOk returns a tuple with the IdentitySources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConnectedOrganization) GetIdentitySourcesOk() (*[]*AnyOfobject, bool) {
	if o == nil || o.IdentitySources == nil {
		return nil, false
	}
	return o.IdentitySources, true
}

// HasIdentitySources returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasIdentitySources() bool {
	if o != nil && o.IdentitySources != nil {
		return true
	}

	return false
}

// SetIdentitySources gets a reference to the given []*AnyOfobject and assigns it to the IdentitySources field.
func (o *MicrosoftGraphConnectedOrganization) SetIdentitySources(v []*AnyOfobject) {
	o.IdentitySources = &v
}

// GetModifiedDateTime returns the ModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConnectedOrganization) GetModifiedDateTime() time.Time {
	if o == nil || o.ModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ModifiedDateTime.Get()
}

// GetModifiedDateTimeOk returns a tuple with the ModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConnectedOrganization) GetModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ModifiedDateTime.Get(), o.ModifiedDateTime.IsSet()
}

// HasModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasModifiedDateTime() bool {
	if o != nil && o.ModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetModifiedDateTime gets a reference to the given NullableTime and assigns it to the ModifiedDateTime field.
func (o *MicrosoftGraphConnectedOrganization) SetModifiedDateTime(v time.Time) {
	o.ModifiedDateTime.Set(&v)
}
// SetModifiedDateTimeNil sets the value for ModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphConnectedOrganization) SetModifiedDateTimeNil() {
	o.ModifiedDateTime.Set(nil)
}

// UnsetModifiedDateTime ensures that no value is present for ModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphConnectedOrganization) UnsetModifiedDateTime() {
	o.ModifiedDateTime.Unset()
}

// GetState returns the State field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphConnectedOrganization) GetState() AnyOfmicrosoftGraphConnectedOrganizationState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConnectedOrganizationState
		return ret
	}
	return o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphConnectedOrganization) GetStateOk() (*AnyOfmicrosoftGraphConnectedOrganizationState, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return &o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given AnyOfmicrosoftGraphConnectedOrganizationState and assigns it to the State field.
func (o *MicrosoftGraphConnectedOrganization) SetState(v AnyOfmicrosoftGraphConnectedOrganizationState) {
	o.State = v
}

// GetExternalSponsors returns the ExternalSponsors field value if set, zero value otherwise.
func (o *MicrosoftGraphConnectedOrganization) GetExternalSponsors() []MicrosoftGraphDirectoryObject {
	if o == nil || o.ExternalSponsors == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.ExternalSponsors
}

// GetExternalSponsorsOk returns a tuple with the ExternalSponsors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConnectedOrganization) GetExternalSponsorsOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.ExternalSponsors == nil {
		return nil, false
	}
	return o.ExternalSponsors, true
}

// HasExternalSponsors returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasExternalSponsors() bool {
	if o != nil && o.ExternalSponsors != nil {
		return true
	}

	return false
}

// SetExternalSponsors gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the ExternalSponsors field.
func (o *MicrosoftGraphConnectedOrganization) SetExternalSponsors(v []MicrosoftGraphDirectoryObject) {
	o.ExternalSponsors = &v
}

// GetInternalSponsors returns the InternalSponsors field value if set, zero value otherwise.
func (o *MicrosoftGraphConnectedOrganization) GetInternalSponsors() []MicrosoftGraphDirectoryObject {
	if o == nil || o.InternalSponsors == nil {
		var ret []MicrosoftGraphDirectoryObject
		return ret
	}
	return *o.InternalSponsors
}

// GetInternalSponsorsOk returns a tuple with the InternalSponsors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphConnectedOrganization) GetInternalSponsorsOk() (*[]MicrosoftGraphDirectoryObject, bool) {
	if o == nil || o.InternalSponsors == nil {
		return nil, false
	}
	return o.InternalSponsors, true
}

// HasInternalSponsors returns a boolean if a field has been set.
func (o *MicrosoftGraphConnectedOrganization) HasInternalSponsors() bool {
	if o != nil && o.InternalSponsors != nil {
		return true
	}

	return false
}

// SetInternalSponsors gets a reference to the given []MicrosoftGraphDirectoryObject and assigns it to the InternalSponsors field.
func (o *MicrosoftGraphConnectedOrganization) SetInternalSponsors(v []MicrosoftGraphDirectoryObject) {
	o.InternalSponsors = &v
}

func (o MicrosoftGraphConnectedOrganization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.IdentitySources != nil {
		toSerialize["identitySources"] = o.IdentitySources
	}
	if o.ModifiedDateTime.IsSet() {
		toSerialize["modifiedDateTime"] = o.ModifiedDateTime.Get()
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.ExternalSponsors != nil {
		toSerialize["externalSponsors"] = o.ExternalSponsors
	}
	if o.InternalSponsors != nil {
		toSerialize["internalSponsors"] = o.InternalSponsors
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphConnectedOrganization struct {
	value *MicrosoftGraphConnectedOrganization
	isSet bool
}

func (v NullableMicrosoftGraphConnectedOrganization) Get() *MicrosoftGraphConnectedOrganization {
	return v.value
}

func (v *NullableMicrosoftGraphConnectedOrganization) Set(val *MicrosoftGraphConnectedOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphConnectedOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphConnectedOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphConnectedOrganization(val *MicrosoftGraphConnectedOrganization) *NullableMicrosoftGraphConnectedOrganization {
	return &NullableMicrosoftGraphConnectedOrganization{value: val, isSet: true}
}

func (v NullableMicrosoftGraphConnectedOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphConnectedOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


