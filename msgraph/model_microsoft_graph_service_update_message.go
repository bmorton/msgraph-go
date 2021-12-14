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

// MicrosoftGraphServiceUpdateMessage struct for MicrosoftGraphServiceUpdateMessage
type MicrosoftGraphServiceUpdateMessage struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Additional details about service event. This property doesn't support filters.
	Details *[]*AnyOfmicrosoftGraphKeyValuePair `json:"details,omitempty"`
	// The end time of the service event.
	EndDateTime NullableTime `json:"endDateTime,omitempty"`
	// The last modified time of the service event.
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// The start time of the service event.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// The title of the service event.
	Title *string `json:"title,omitempty"`
	// The expected deadline of the action for the message.
	ActionRequiredByDateTime NullableTime `json:"actionRequiredByDateTime,omitempty"`
	Body *MicrosoftGraphItemBody `json:"body,omitempty"`
	// The service message category. Possible values are: preventOrFixIssue, planForChange, stayInformed, unknownFutureValue.
	Category AnyOfmicrosoftGraphServiceUpdateCategory `json:"category,omitempty"`
	// Indicates whether the message describes a major update for the service.
	IsMajorChange NullableBool `json:"isMajorChange,omitempty"`
	// The affected services by the service message.
	Services *[]*string `json:"services,omitempty"`
	// The severity of the service message. Possible values are: normal, high, critical, unknownFutureValue.
	Severity AnyOfmicrosoftGraphServiceUpdateSeverity `json:"severity,omitempty"`
	// A collection of tags for the service message.
	Tags *[]*string `json:"tags,omitempty"`
	// Represents user view points data of the service message. This data includes message status such as whether the user has archived, read, or marked the message as favorite. This property is null when accessed with application permissions.
	ViewPoint AnyOfmicrosoftGraphServiceUpdateMessageViewpoint `json:"viewPoint,omitempty"`
}

// NewMicrosoftGraphServiceUpdateMessage instantiates a new MicrosoftGraphServiceUpdateMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphServiceUpdateMessage() *MicrosoftGraphServiceUpdateMessage {
	this := MicrosoftGraphServiceUpdateMessage{}
	return &this
}

// NewMicrosoftGraphServiceUpdateMessageWithDefaults instantiates a new MicrosoftGraphServiceUpdateMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphServiceUpdateMessageWithDefaults() *MicrosoftGraphServiceUpdateMessage {
	this := MicrosoftGraphServiceUpdateMessage{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceUpdateMessage) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceUpdateMessage) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphServiceUpdateMessage) SetId(v string) {
	o.Id = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceUpdateMessage) GetDetails() []*AnyOfmicrosoftGraphKeyValuePair {
	if o == nil || o.Details == nil {
		var ret []*AnyOfmicrosoftGraphKeyValuePair
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceUpdateMessage) GetDetailsOk() (*[]*AnyOfmicrosoftGraphKeyValuePair, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []*AnyOfmicrosoftGraphKeyValuePair and assigns it to the Details field.
func (o *MicrosoftGraphServiceUpdateMessage) SetDetails(v []*AnyOfmicrosoftGraphKeyValuePair) {
	o.Details = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphServiceUpdateMessage) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime.Get()
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphServiceUpdateMessage) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EndDateTime.Get(), o.EndDateTime.IsSet()
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasEndDateTime() bool {
	if o != nil && o.EndDateTime.IsSet() {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given NullableTime and assigns it to the EndDateTime field.
func (o *MicrosoftGraphServiceUpdateMessage) SetEndDateTime(v time.Time) {
	o.EndDateTime.Set(&v)
}
// SetEndDateTimeNil sets the value for EndDateTime to be an explicit nil
func (o *MicrosoftGraphServiceUpdateMessage) SetEndDateTimeNil() {
	o.EndDateTime.Set(nil)
}

// UnsetEndDateTime ensures that no value is present for EndDateTime, not even an explicit nil
func (o *MicrosoftGraphServiceUpdateMessage) UnsetEndDateTime() {
	o.EndDateTime.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceUpdateMessage) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceUpdateMessage) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedDateTime == nil {
		return nil, false
	}
	return o.LastModifiedDateTime, true
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime != nil {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given time.Time and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphServiceUpdateMessage) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime = &v
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceUpdateMessage) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceUpdateMessage) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MicrosoftGraphServiceUpdateMessage) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceUpdateMessage) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceUpdateMessage) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *MicrosoftGraphServiceUpdateMessage) SetTitle(v string) {
	o.Title = &v
}

// GetActionRequiredByDateTime returns the ActionRequiredByDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphServiceUpdateMessage) GetActionRequiredByDateTime() time.Time {
	if o == nil || o.ActionRequiredByDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.ActionRequiredByDateTime.Get()
}

// GetActionRequiredByDateTimeOk returns a tuple with the ActionRequiredByDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphServiceUpdateMessage) GetActionRequiredByDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ActionRequiredByDateTime.Get(), o.ActionRequiredByDateTime.IsSet()
}

// HasActionRequiredByDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasActionRequiredByDateTime() bool {
	if o != nil && o.ActionRequiredByDateTime.IsSet() {
		return true
	}

	return false
}

// SetActionRequiredByDateTime gets a reference to the given NullableTime and assigns it to the ActionRequiredByDateTime field.
func (o *MicrosoftGraphServiceUpdateMessage) SetActionRequiredByDateTime(v time.Time) {
	o.ActionRequiredByDateTime.Set(&v)
}
// SetActionRequiredByDateTimeNil sets the value for ActionRequiredByDateTime to be an explicit nil
func (o *MicrosoftGraphServiceUpdateMessage) SetActionRequiredByDateTimeNil() {
	o.ActionRequiredByDateTime.Set(nil)
}

// UnsetActionRequiredByDateTime ensures that no value is present for ActionRequiredByDateTime, not even an explicit nil
func (o *MicrosoftGraphServiceUpdateMessage) UnsetActionRequiredByDateTime() {
	o.ActionRequiredByDateTime.Unset()
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceUpdateMessage) GetBody() MicrosoftGraphItemBody {
	if o == nil || o.Body == nil {
		var ret MicrosoftGraphItemBody
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceUpdateMessage) GetBodyOk() (*MicrosoftGraphItemBody, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given MicrosoftGraphItemBody and assigns it to the Body field.
func (o *MicrosoftGraphServiceUpdateMessage) SetBody(v MicrosoftGraphItemBody) {
	o.Body = &v
}

// GetCategory returns the Category field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphServiceUpdateMessage) GetCategory() AnyOfmicrosoftGraphServiceUpdateCategory {
	if o == nil  {
		var ret AnyOfmicrosoftGraphServiceUpdateCategory
		return ret
	}
	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphServiceUpdateMessage) GetCategoryOk() (*AnyOfmicrosoftGraphServiceUpdateCategory, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return &o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given AnyOfmicrosoftGraphServiceUpdateCategory and assigns it to the Category field.
func (o *MicrosoftGraphServiceUpdateMessage) SetCategory(v AnyOfmicrosoftGraphServiceUpdateCategory) {
	o.Category = v
}

// GetIsMajorChange returns the IsMajorChange field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphServiceUpdateMessage) GetIsMajorChange() bool {
	if o == nil || o.IsMajorChange.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsMajorChange.Get()
}

// GetIsMajorChangeOk returns a tuple with the IsMajorChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphServiceUpdateMessage) GetIsMajorChangeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsMajorChange.Get(), o.IsMajorChange.IsSet()
}

// HasIsMajorChange returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasIsMajorChange() bool {
	if o != nil && o.IsMajorChange.IsSet() {
		return true
	}

	return false
}

// SetIsMajorChange gets a reference to the given NullableBool and assigns it to the IsMajorChange field.
func (o *MicrosoftGraphServiceUpdateMessage) SetIsMajorChange(v bool) {
	o.IsMajorChange.Set(&v)
}
// SetIsMajorChangeNil sets the value for IsMajorChange to be an explicit nil
func (o *MicrosoftGraphServiceUpdateMessage) SetIsMajorChangeNil() {
	o.IsMajorChange.Set(nil)
}

// UnsetIsMajorChange ensures that no value is present for IsMajorChange, not even an explicit nil
func (o *MicrosoftGraphServiceUpdateMessage) UnsetIsMajorChange() {
	o.IsMajorChange.Unset()
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceUpdateMessage) GetServices() []*string {
	if o == nil || o.Services == nil {
		var ret []*string
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceUpdateMessage) GetServicesOk() (*[]*string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []*string and assigns it to the Services field.
func (o *MicrosoftGraphServiceUpdateMessage) SetServices(v []*string) {
	o.Services = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphServiceUpdateMessage) GetSeverity() AnyOfmicrosoftGraphServiceUpdateSeverity {
	if o == nil  {
		var ret AnyOfmicrosoftGraphServiceUpdateSeverity
		return ret
	}
	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphServiceUpdateMessage) GetSeverityOk() (*AnyOfmicrosoftGraphServiceUpdateSeverity, bool) {
	if o == nil || o.Severity == nil {
		return nil, false
	}
	return &o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasSeverity() bool {
	if o != nil && o.Severity != nil {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given AnyOfmicrosoftGraphServiceUpdateSeverity and assigns it to the Severity field.
func (o *MicrosoftGraphServiceUpdateMessage) SetSeverity(v AnyOfmicrosoftGraphServiceUpdateSeverity) {
	o.Severity = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceUpdateMessage) GetTags() []*string {
	if o == nil || o.Tags == nil {
		var ret []*string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceUpdateMessage) GetTagsOk() (*[]*string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []*string and assigns it to the Tags field.
func (o *MicrosoftGraphServiceUpdateMessage) SetTags(v []*string) {
	o.Tags = &v
}

// GetViewPoint returns the ViewPoint field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphServiceUpdateMessage) GetViewPoint() AnyOfmicrosoftGraphServiceUpdateMessageViewpoint {
	if o == nil  {
		var ret AnyOfmicrosoftGraphServiceUpdateMessageViewpoint
		return ret
	}
	return o.ViewPoint
}

// GetViewPointOk returns a tuple with the ViewPoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphServiceUpdateMessage) GetViewPointOk() (*AnyOfmicrosoftGraphServiceUpdateMessageViewpoint, bool) {
	if o == nil || o.ViewPoint == nil {
		return nil, false
	}
	return &o.ViewPoint, true
}

// HasViewPoint returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceUpdateMessage) HasViewPoint() bool {
	if o != nil && o.ViewPoint != nil {
		return true
	}

	return false
}

// SetViewPoint gets a reference to the given AnyOfmicrosoftGraphServiceUpdateMessageViewpoint and assigns it to the ViewPoint field.
func (o *MicrosoftGraphServiceUpdateMessage) SetViewPoint(v AnyOfmicrosoftGraphServiceUpdateMessageViewpoint) {
	o.ViewPoint = v
}

func (o MicrosoftGraphServiceUpdateMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.EndDateTime.IsSet() {
		toSerialize["endDateTime"] = o.EndDateTime.Get()
	}
	if o.LastModifiedDateTime != nil {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime
	}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.ActionRequiredByDateTime.IsSet() {
		toSerialize["actionRequiredByDateTime"] = o.ActionRequiredByDateTime.Get()
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.IsMajorChange.IsSet() {
		toSerialize["isMajorChange"] = o.IsMajorChange.Get()
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Severity != nil {
		toSerialize["severity"] = o.Severity
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ViewPoint != nil {
		toSerialize["viewPoint"] = o.ViewPoint
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphServiceUpdateMessage struct {
	value *MicrosoftGraphServiceUpdateMessage
	isSet bool
}

func (v NullableMicrosoftGraphServiceUpdateMessage) Get() *MicrosoftGraphServiceUpdateMessage {
	return v.value
}

func (v *NullableMicrosoftGraphServiceUpdateMessage) Set(val *MicrosoftGraphServiceUpdateMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphServiceUpdateMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphServiceUpdateMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphServiceUpdateMessage(val *MicrosoftGraphServiceUpdateMessage) *NullableMicrosoftGraphServiceUpdateMessage {
	return &NullableMicrosoftGraphServiceUpdateMessage{value: val, isSet: true}
}

func (v NullableMicrosoftGraphServiceUpdateMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphServiceUpdateMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

