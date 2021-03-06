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

// MicrosoftGraphEducationSubmission struct for MicrosoftGraphEducationSubmission
type MicrosoftGraphEducationSubmission struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Who this submission is assigned to.
	Recipient AnyOfobject `json:"recipient,omitempty"`
	// Folder where all file resources for this submission need to be stored.
	ResourcesFolderUrl NullableString `json:"resourcesFolderUrl,omitempty"`
	// User who moved the status of this submission to returned.
	ReturnedBy AnyOfmicrosoftGraphIdentitySet `json:"returnedBy,omitempty"`
	// Moment in time when the submission was returned. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ReturnedDateTime NullableTime `json:"returnedDateTime,omitempty"`
	// Read-Only. Possible values are: working, submitted, released, returned.
	Status AnyOfmicrosoftGraphEducationSubmissionStatus `json:"status,omitempty"`
	// User who moved the resource into the submitted state.
	SubmittedBy AnyOfmicrosoftGraphIdentitySet `json:"submittedBy,omitempty"`
	// Moment in time when the submission was moved into the submitted state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	SubmittedDateTime NullableTime `json:"submittedDateTime,omitempty"`
	// User who moved the resource from submitted into the working state.
	UnsubmittedBy AnyOfmicrosoftGraphIdentitySet `json:"unsubmittedBy,omitempty"`
	// Moment in time when the submission was moved from submitted into the working state. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	UnsubmittedDateTime NullableTime `json:"unsubmittedDateTime,omitempty"`
	// Read-Write. Nullable.
	Outcomes *[]MicrosoftGraphEducationOutcome `json:"outcomes,omitempty"`
	// Nullable.
	Resources *[]MicrosoftGraphEducationSubmissionResource `json:"resources,omitempty"`
	// Read-only. Nullable.
	SubmittedResources *[]MicrosoftGraphEducationSubmissionResource `json:"submittedResources,omitempty"`
}

// NewMicrosoftGraphEducationSubmission instantiates a new MicrosoftGraphEducationSubmission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphEducationSubmission() *MicrosoftGraphEducationSubmission {
	this := MicrosoftGraphEducationSubmission{}
	return &this
}

// NewMicrosoftGraphEducationSubmissionWithDefaults instantiates a new MicrosoftGraphEducationSubmission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphEducationSubmissionWithDefaults() *MicrosoftGraphEducationSubmission {
	this := MicrosoftGraphEducationSubmission{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationSubmission) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationSubmission) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphEducationSubmission) SetId(v string) {
	o.Id = &v
}

// GetRecipient returns the Recipient field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetRecipient() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetRecipientOk() (*AnyOfobject, bool) {
	if o == nil || o.Recipient == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// HasRecipient returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasRecipient() bool {
	if o != nil && o.Recipient != nil {
		return true
	}

	return false
}

// SetRecipient gets a reference to the given AnyOfobject and assigns it to the Recipient field.
func (o *MicrosoftGraphEducationSubmission) SetRecipient(v AnyOfobject) {
	o.Recipient = v
}

// GetResourcesFolderUrl returns the ResourcesFolderUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetResourcesFolderUrl() string {
	if o == nil || o.ResourcesFolderUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourcesFolderUrl.Get()
}

// GetResourcesFolderUrlOk returns a tuple with the ResourcesFolderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetResourcesFolderUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourcesFolderUrl.Get(), o.ResourcesFolderUrl.IsSet()
}

// HasResourcesFolderUrl returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasResourcesFolderUrl() bool {
	if o != nil && o.ResourcesFolderUrl.IsSet() {
		return true
	}

	return false
}

// SetResourcesFolderUrl gets a reference to the given NullableString and assigns it to the ResourcesFolderUrl field.
func (o *MicrosoftGraphEducationSubmission) SetResourcesFolderUrl(v string) {
	o.ResourcesFolderUrl.Set(&v)
}
// SetResourcesFolderUrlNil sets the value for ResourcesFolderUrl to be an explicit nil
func (o *MicrosoftGraphEducationSubmission) SetResourcesFolderUrlNil() {
	o.ResourcesFolderUrl.Set(nil)
}

// UnsetResourcesFolderUrl ensures that no value is present for ResourcesFolderUrl, not even an explicit nil
func (o *MicrosoftGraphEducationSubmission) UnsetResourcesFolderUrl() {
	o.ResourcesFolderUrl.Unset()
}

// GetReturnedBy returns the ReturnedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetReturnedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.ReturnedBy
}

// GetReturnedByOk returns a tuple with the ReturnedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetReturnedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.ReturnedBy == nil {
		return nil, false
	}
	return &o.ReturnedBy, true
}

// HasReturnedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasReturnedBy() bool {
	if o != nil && o.ReturnedBy != nil {
		return true
	}

	return false
}

// SetReturnedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the ReturnedBy field.
func (o *MicrosoftGraphEducationSubmission) SetReturnedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.ReturnedBy = v
}

// GetReturnedDateTime returns the ReturnedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetReturnedDateTime() time.Time {
	if o == nil || o.ReturnedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ReturnedDateTime.Get()
}

// GetReturnedDateTimeOk returns a tuple with the ReturnedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetReturnedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ReturnedDateTime.Get(), o.ReturnedDateTime.IsSet()
}

// HasReturnedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasReturnedDateTime() bool {
	if o != nil && o.ReturnedDateTime.IsSet() {
		return true
	}

	return false
}

// SetReturnedDateTime gets a reference to the given NullableTime and assigns it to the ReturnedDateTime field.
func (o *MicrosoftGraphEducationSubmission) SetReturnedDateTime(v time.Time) {
	o.ReturnedDateTime.Set(&v)
}
// SetReturnedDateTimeNil sets the value for ReturnedDateTime to be an explicit nil
func (o *MicrosoftGraphEducationSubmission) SetReturnedDateTimeNil() {
	o.ReturnedDateTime.Set(nil)
}

// UnsetReturnedDateTime ensures that no value is present for ReturnedDateTime, not even an explicit nil
func (o *MicrosoftGraphEducationSubmission) UnsetReturnedDateTime() {
	o.ReturnedDateTime.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetStatus() AnyOfmicrosoftGraphEducationSubmissionStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEducationSubmissionStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetStatusOk() (*AnyOfmicrosoftGraphEducationSubmissionStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphEducationSubmissionStatus and assigns it to the Status field.
func (o *MicrosoftGraphEducationSubmission) SetStatus(v AnyOfmicrosoftGraphEducationSubmissionStatus) {
	o.Status = v
}

// GetSubmittedBy returns the SubmittedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetSubmittedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.SubmittedBy
}

// GetSubmittedByOk returns a tuple with the SubmittedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetSubmittedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.SubmittedBy == nil {
		return nil, false
	}
	return &o.SubmittedBy, true
}

// HasSubmittedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasSubmittedBy() bool {
	if o != nil && o.SubmittedBy != nil {
		return true
	}

	return false
}

// SetSubmittedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the SubmittedBy field.
func (o *MicrosoftGraphEducationSubmission) SetSubmittedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.SubmittedBy = v
}

// GetSubmittedDateTime returns the SubmittedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetSubmittedDateTime() time.Time {
	if o == nil || o.SubmittedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.SubmittedDateTime.Get()
}

// GetSubmittedDateTimeOk returns a tuple with the SubmittedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetSubmittedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SubmittedDateTime.Get(), o.SubmittedDateTime.IsSet()
}

// HasSubmittedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasSubmittedDateTime() bool {
	if o != nil && o.SubmittedDateTime.IsSet() {
		return true
	}

	return false
}

// SetSubmittedDateTime gets a reference to the given NullableTime and assigns it to the SubmittedDateTime field.
func (o *MicrosoftGraphEducationSubmission) SetSubmittedDateTime(v time.Time) {
	o.SubmittedDateTime.Set(&v)
}
// SetSubmittedDateTimeNil sets the value for SubmittedDateTime to be an explicit nil
func (o *MicrosoftGraphEducationSubmission) SetSubmittedDateTimeNil() {
	o.SubmittedDateTime.Set(nil)
}

// UnsetSubmittedDateTime ensures that no value is present for SubmittedDateTime, not even an explicit nil
func (o *MicrosoftGraphEducationSubmission) UnsetSubmittedDateTime() {
	o.SubmittedDateTime.Unset()
}

// GetUnsubmittedBy returns the UnsubmittedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetUnsubmittedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.UnsubmittedBy
}

// GetUnsubmittedByOk returns a tuple with the UnsubmittedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetUnsubmittedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.UnsubmittedBy == nil {
		return nil, false
	}
	return &o.UnsubmittedBy, true
}

// HasUnsubmittedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasUnsubmittedBy() bool {
	if o != nil && o.UnsubmittedBy != nil {
		return true
	}

	return false
}

// SetUnsubmittedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the UnsubmittedBy field.
func (o *MicrosoftGraphEducationSubmission) SetUnsubmittedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.UnsubmittedBy = v
}

// GetUnsubmittedDateTime returns the UnsubmittedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphEducationSubmission) GetUnsubmittedDateTime() time.Time {
	if o == nil || o.UnsubmittedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.UnsubmittedDateTime.Get()
}

// GetUnsubmittedDateTimeOk returns a tuple with the UnsubmittedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphEducationSubmission) GetUnsubmittedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UnsubmittedDateTime.Get(), o.UnsubmittedDateTime.IsSet()
}

// HasUnsubmittedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasUnsubmittedDateTime() bool {
	if o != nil && o.UnsubmittedDateTime.IsSet() {
		return true
	}

	return false
}

// SetUnsubmittedDateTime gets a reference to the given NullableTime and assigns it to the UnsubmittedDateTime field.
func (o *MicrosoftGraphEducationSubmission) SetUnsubmittedDateTime(v time.Time) {
	o.UnsubmittedDateTime.Set(&v)
}
// SetUnsubmittedDateTimeNil sets the value for UnsubmittedDateTime to be an explicit nil
func (o *MicrosoftGraphEducationSubmission) SetUnsubmittedDateTimeNil() {
	o.UnsubmittedDateTime.Set(nil)
}

// UnsetUnsubmittedDateTime ensures that no value is present for UnsubmittedDateTime, not even an explicit nil
func (o *MicrosoftGraphEducationSubmission) UnsetUnsubmittedDateTime() {
	o.UnsubmittedDateTime.Unset()
}

// GetOutcomes returns the Outcomes field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationSubmission) GetOutcomes() []MicrosoftGraphEducationOutcome {
	if o == nil || o.Outcomes == nil {
		var ret []MicrosoftGraphEducationOutcome
		return ret
	}
	return *o.Outcomes
}

// GetOutcomesOk returns a tuple with the Outcomes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationSubmission) GetOutcomesOk() (*[]MicrosoftGraphEducationOutcome, bool) {
	if o == nil || o.Outcomes == nil {
		return nil, false
	}
	return o.Outcomes, true
}

// HasOutcomes returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasOutcomes() bool {
	if o != nil && o.Outcomes != nil {
		return true
	}

	return false
}

// SetOutcomes gets a reference to the given []MicrosoftGraphEducationOutcome and assigns it to the Outcomes field.
func (o *MicrosoftGraphEducationSubmission) SetOutcomes(v []MicrosoftGraphEducationOutcome) {
	o.Outcomes = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationSubmission) GetResources() []MicrosoftGraphEducationSubmissionResource {
	if o == nil || o.Resources == nil {
		var ret []MicrosoftGraphEducationSubmissionResource
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationSubmission) GetResourcesOk() (*[]MicrosoftGraphEducationSubmissionResource, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []MicrosoftGraphEducationSubmissionResource and assigns it to the Resources field.
func (o *MicrosoftGraphEducationSubmission) SetResources(v []MicrosoftGraphEducationSubmissionResource) {
	o.Resources = &v
}

// GetSubmittedResources returns the SubmittedResources field value if set, zero value otherwise.
func (o *MicrosoftGraphEducationSubmission) GetSubmittedResources() []MicrosoftGraphEducationSubmissionResource {
	if o == nil || o.SubmittedResources == nil {
		var ret []MicrosoftGraphEducationSubmissionResource
		return ret
	}
	return *o.SubmittedResources
}

// GetSubmittedResourcesOk returns a tuple with the SubmittedResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphEducationSubmission) GetSubmittedResourcesOk() (*[]MicrosoftGraphEducationSubmissionResource, bool) {
	if o == nil || o.SubmittedResources == nil {
		return nil, false
	}
	return o.SubmittedResources, true
}

// HasSubmittedResources returns a boolean if a field has been set.
func (o *MicrosoftGraphEducationSubmission) HasSubmittedResources() bool {
	if o != nil && o.SubmittedResources != nil {
		return true
	}

	return false
}

// SetSubmittedResources gets a reference to the given []MicrosoftGraphEducationSubmissionResource and assigns it to the SubmittedResources field.
func (o *MicrosoftGraphEducationSubmission) SetSubmittedResources(v []MicrosoftGraphEducationSubmissionResource) {
	o.SubmittedResources = &v
}

func (o MicrosoftGraphEducationSubmission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Recipient != nil {
		toSerialize["recipient"] = o.Recipient
	}
	if o.ResourcesFolderUrl.IsSet() {
		toSerialize["resourcesFolderUrl"] = o.ResourcesFolderUrl.Get()
	}
	if o.ReturnedBy != nil {
		toSerialize["returnedBy"] = o.ReturnedBy
	}
	if o.ReturnedDateTime.IsSet() {
		toSerialize["returnedDateTime"] = o.ReturnedDateTime.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.SubmittedBy != nil {
		toSerialize["submittedBy"] = o.SubmittedBy
	}
	if o.SubmittedDateTime.IsSet() {
		toSerialize["submittedDateTime"] = o.SubmittedDateTime.Get()
	}
	if o.UnsubmittedBy != nil {
		toSerialize["unsubmittedBy"] = o.UnsubmittedBy
	}
	if o.UnsubmittedDateTime.IsSet() {
		toSerialize["unsubmittedDateTime"] = o.UnsubmittedDateTime.Get()
	}
	if o.Outcomes != nil {
		toSerialize["outcomes"] = o.Outcomes
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.SubmittedResources != nil {
		toSerialize["submittedResources"] = o.SubmittedResources
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphEducationSubmission struct {
	value *MicrosoftGraphEducationSubmission
	isSet bool
}

func (v NullableMicrosoftGraphEducationSubmission) Get() *MicrosoftGraphEducationSubmission {
	return v.value
}

func (v *NullableMicrosoftGraphEducationSubmission) Set(val *MicrosoftGraphEducationSubmission) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphEducationSubmission) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphEducationSubmission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphEducationSubmission(val *MicrosoftGraphEducationSubmission) *NullableMicrosoftGraphEducationSubmission {
	return &NullableMicrosoftGraphEducationSubmission{value: val, isSet: true}
}

func (v NullableMicrosoftGraphEducationSubmission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphEducationSubmission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


