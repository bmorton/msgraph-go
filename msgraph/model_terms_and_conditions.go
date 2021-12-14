/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// TermsAndConditions A termsAndConditions entity represents the metadata and contents of a given Terms and Conditions (T&C) policy. T&C policies’ contents are presented to users upon their first attempt to enroll into Intune and subsequently upon edits where an administrator has required re-acceptance. They enable administrators to communicate the provisions to which a user must agree in order to have devices enrolled into Intune.
type TermsAndConditions struct {
	// Administrator-supplied explanation of the terms and conditions, typically describing what it means to accept the terms and conditions set out in the T&C policy. This is shown to the user on prompts to accept the T&C policy.
	AcceptanceStatement NullableString `json:"acceptanceStatement,omitempty"`
	// Administrator-supplied body text of the terms and conditions, typically the terms themselves. This is shown to the user on prompts to accept the T&C policy.
	BodyText NullableString `json:"bodyText,omitempty"`
	// DateTime the object was created.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Administrator-supplied description of the T&C policy.
	Description NullableString `json:"description,omitempty"`
	// Administrator-supplied name for the T&C policy.
	DisplayName *string `json:"displayName,omitempty"`
	// DateTime the object was last modified.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Administrator-supplied title of the terms and conditions. This is shown to the user on prompts to accept the T&C policy.
	Title NullableString `json:"title,omitempty"`
	// Integer indicating the current version of the terms. Incremented when an administrator makes a change to the terms and wishes to require users to re-accept the modified T&C policy.
	Version *int32 `json:"version,omitempty"`
	// The list of acceptance statuses for this T&C policy.
	AcceptanceStatuses *[]MicrosoftGraphTermsAndConditionsAcceptanceStatus `json:"acceptanceStatuses,omitempty"`
	// The list of assignments for this T&C policy.
	Assignments *[]MicrosoftGraphTermsAndConditionsAssignment `json:"assignments,omitempty"`
}

// NewTermsAndConditions instantiates a new TermsAndConditions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTermsAndConditions() *TermsAndConditions {
	this := TermsAndConditions{}
	return &this
}

// NewTermsAndConditionsWithDefaults instantiates a new TermsAndConditions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTermsAndConditionsWithDefaults() *TermsAndConditions {
	this := TermsAndConditions{}
	return &this
}

// GetAcceptanceStatement returns the AcceptanceStatement field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TermsAndConditions) GetAcceptanceStatement() string {
	if o == nil || o.AcceptanceStatement.Get() == nil {
		var ret string
		return ret
	}
	return *o.AcceptanceStatement.Get()
}

// GetAcceptanceStatementOk returns a tuple with the AcceptanceStatement field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TermsAndConditions) GetAcceptanceStatementOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AcceptanceStatement.Get(), o.AcceptanceStatement.IsSet()
}

// HasAcceptanceStatement returns a boolean if a field has been set.
func (o *TermsAndConditions) HasAcceptanceStatement() bool {
	if o != nil && o.AcceptanceStatement.IsSet() {
		return true
	}

	return false
}

// SetAcceptanceStatement gets a reference to the given NullableString and assigns it to the AcceptanceStatement field.
func (o *TermsAndConditions) SetAcceptanceStatement(v string) {
	o.AcceptanceStatement.Set(&v)
}
// SetAcceptanceStatementNil sets the value for AcceptanceStatement to be an explicit nil
func (o *TermsAndConditions) SetAcceptanceStatementNil() {
	o.AcceptanceStatement.Set(nil)
}

// UnsetAcceptanceStatement ensures that no value is present for AcceptanceStatement, not even an explicit nil
func (o *TermsAndConditions) UnsetAcceptanceStatement() {
	o.AcceptanceStatement.Unset()
}

// GetBodyText returns the BodyText field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TermsAndConditions) GetBodyText() string {
	if o == nil || o.BodyText.Get() == nil {
		var ret string
		return ret
	}
	return *o.BodyText.Get()
}

// GetBodyTextOk returns a tuple with the BodyText field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TermsAndConditions) GetBodyTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.BodyText.Get(), o.BodyText.IsSet()
}

// HasBodyText returns a boolean if a field has been set.
func (o *TermsAndConditions) HasBodyText() bool {
	if o != nil && o.BodyText.IsSet() {
		return true
	}

	return false
}

// SetBodyText gets a reference to the given NullableString and assigns it to the BodyText field.
func (o *TermsAndConditions) SetBodyText(v string) {
	o.BodyText.Set(&v)
}
// SetBodyTextNil sets the value for BodyText to be an explicit nil
func (o *TermsAndConditions) SetBodyTextNil() {
	o.BodyText.Set(nil)
}

// UnsetBodyText ensures that no value is present for BodyText, not even an explicit nil
func (o *TermsAndConditions) UnsetBodyText() {
	o.BodyText.Unset()
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *TermsAndConditions) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *TermsAndConditions) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *TermsAndConditions) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TermsAndConditions) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TermsAndConditions) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *TermsAndConditions) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *TermsAndConditions) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *TermsAndConditions) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *TermsAndConditions) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *TermsAndConditions) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *TermsAndConditions) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *TermsAndConditions) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *TermsAndConditions) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *TermsAndConditions) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *TermsAndConditions) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TermsAndConditions) GetTitle() string {
	if o == nil || o.Title.Get() == nil {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TermsAndConditions) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *TermsAndConditions) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *TermsAndConditions) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *TermsAndConditions) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *TermsAndConditions) UnsetTitle() {
	o.Title.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *TermsAndConditions) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *TermsAndConditions) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *TermsAndConditions) SetVersion(v int32) {
	o.Version = &v
}

// GetAcceptanceStatuses returns the AcceptanceStatuses field value if set, zero value otherwise.
func (o *TermsAndConditions) GetAcceptanceStatuses() []MicrosoftGraphTermsAndConditionsAcceptanceStatus {
	if o == nil || o.AcceptanceStatuses == nil {
		var ret []MicrosoftGraphTermsAndConditionsAcceptanceStatus
		return ret
	}
	return *o.AcceptanceStatuses
}

// GetAcceptanceStatusesOk returns a tuple with the AcceptanceStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetAcceptanceStatusesOk() (*[]MicrosoftGraphTermsAndConditionsAcceptanceStatus, bool) {
	if o == nil || o.AcceptanceStatuses == nil {
		return nil, false
	}
	return o.AcceptanceStatuses, true
}

// HasAcceptanceStatuses returns a boolean if a field has been set.
func (o *TermsAndConditions) HasAcceptanceStatuses() bool {
	if o != nil && o.AcceptanceStatuses != nil {
		return true
	}

	return false
}

// SetAcceptanceStatuses gets a reference to the given []MicrosoftGraphTermsAndConditionsAcceptanceStatus and assigns it to the AcceptanceStatuses field.
func (o *TermsAndConditions) SetAcceptanceStatuses(v []MicrosoftGraphTermsAndConditionsAcceptanceStatus) {
	o.AcceptanceStatuses = &v
}

// GetAssignments returns the Assignments field value if set, zero value otherwise.
func (o *TermsAndConditions) GetAssignments() []MicrosoftGraphTermsAndConditionsAssignment {
	if o == nil || o.Assignments == nil {
		var ret []MicrosoftGraphTermsAndConditionsAssignment
		return ret
	}
	return *o.Assignments
}

// GetAssignmentsOk returns a tuple with the Assignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TermsAndConditions) GetAssignmentsOk() (*[]MicrosoftGraphTermsAndConditionsAssignment, bool) {
	if o == nil || o.Assignments == nil {
		return nil, false
	}
	return o.Assignments, true
}

// HasAssignments returns a boolean if a field has been set.
func (o *TermsAndConditions) HasAssignments() bool {
	if o != nil && o.Assignments != nil {
		return true
	}

	return false
}

// SetAssignments gets a reference to the given []MicrosoftGraphTermsAndConditionsAssignment and assigns it to the Assignments field.
func (o *TermsAndConditions) SetAssignments(v []MicrosoftGraphTermsAndConditionsAssignment) {
	o.Assignments = &v
}

func (o TermsAndConditions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AcceptanceStatement.IsSet() {
		toSerialize["acceptanceStatement"] = o.AcceptanceStatement.Get()
	}
	if o.BodyText.IsSet() {
		toSerialize["bodyText"] = o.BodyText.Get()
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.AcceptanceStatuses != nil {
		toSerialize["acceptanceStatuses"] = o.AcceptanceStatuses
	}
	if o.Assignments != nil {
		toSerialize["assignments"] = o.Assignments
	}
	return json.Marshal(toSerialize)
}

type NullableTermsAndConditions struct {
	value *TermsAndConditions
	isSet bool
}

func (v NullableTermsAndConditions) Get() *TermsAndConditions {
	return v.value
}

func (v *NullableTermsAndConditions) Set(val *TermsAndConditions) {
	v.value = val
	v.isSet = true
}

func (v NullableTermsAndConditions) IsSet() bool {
	return v.isSet
}

func (v *NullableTermsAndConditions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermsAndConditions(val *TermsAndConditions) *NullableTermsAndConditions {
	return &NullableTermsAndConditions{value: val, isSet: true}
}

func (v NullableTermsAndConditions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermsAndConditions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

