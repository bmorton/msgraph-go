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

// MicrosoftGraphDirectoryAudit struct for MicrosoftGraphDirectoryAudit
type MicrosoftGraphDirectoryAudit struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Indicates the date and time the activity was performed. The Timestamp type is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ActivityDateTime *time.Time `json:"activityDateTime,omitempty"`
	// Indicates the activity name or the operation name (examples: 'Create User' and 'Add member to group'). For full list, see Azure AD activity list.
	ActivityDisplayName *string `json:"activityDisplayName,omitempty"`
	// Indicates additional details on the activity.
	AdditionalDetails *[]*AnyOfmicrosoftGraphKeyValue `json:"additionalDetails,omitempty"`
	// Indicates which resource category that's targeted by the activity. (For example: User Management, Group Management etc..)
	Category *string `json:"category,omitempty"`
	// Indicates a unique ID that helps correlate activities that span across various services. Can be used to trace logs across services.
	CorrelationId NullableString `json:"correlationId,omitempty"`
	InitiatedBy *MicrosoftGraphAuditActivityInitiator `json:"initiatedBy,omitempty"`
	// Indicates information on which service initiated the activity (For example: Self-service Password Management, Core Directory, B2C, Invited Users, Microsoft Identity Manager, Privileged Identity Management.
	LoggedByService NullableString `json:"loggedByService,omitempty"`
	OperationType NullableString `json:"operationType,omitempty"`
	// Indicates the result of the activity. Possible values are: success, failure, timeout, unknownFutureValue.
	Result AnyOfmicrosoftGraphOperationResult `json:"result,omitempty"`
	// Indicates the reason for failure if the result is failure or timeout.
	ResultReason NullableString `json:"resultReason,omitempty"`
	// Indicates information on which resource was changed due to the activity. Target Resource Type can be User, Device, Directory, App, Role, Group, Policy or Other.
	TargetResources *[]*AnyOfmicrosoftGraphTargetResource `json:"targetResources,omitempty"`
}

// NewMicrosoftGraphDirectoryAudit instantiates a new MicrosoftGraphDirectoryAudit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphDirectoryAudit() *MicrosoftGraphDirectoryAudit {
	this := MicrosoftGraphDirectoryAudit{}
	return &this
}

// NewMicrosoftGraphDirectoryAuditWithDefaults instantiates a new MicrosoftGraphDirectoryAudit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphDirectoryAuditWithDefaults() *MicrosoftGraphDirectoryAudit {
	this := MicrosoftGraphDirectoryAudit{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryAudit) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryAudit) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphDirectoryAudit) SetId(v string) {
	o.Id = &v
}

// GetActivityDateTime returns the ActivityDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryAudit) GetActivityDateTime() time.Time {
	if o == nil || o.ActivityDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ActivityDateTime
}

// GetActivityDateTimeOk returns a tuple with the ActivityDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryAudit) GetActivityDateTimeOk() (*time.Time, bool) {
	if o == nil || o.ActivityDateTime == nil {
		return nil, false
	}
	return o.ActivityDateTime, true
}

// HasActivityDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasActivityDateTime() bool {
	if o != nil && o.ActivityDateTime != nil {
		return true
	}

	return false
}

// SetActivityDateTime gets a reference to the given time.Time and assigns it to the ActivityDateTime field.
func (o *MicrosoftGraphDirectoryAudit) SetActivityDateTime(v time.Time) {
	o.ActivityDateTime = &v
}

// GetActivityDisplayName returns the ActivityDisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryAudit) GetActivityDisplayName() string {
	if o == nil || o.ActivityDisplayName == nil {
		var ret string
		return ret
	}
	return *o.ActivityDisplayName
}

// GetActivityDisplayNameOk returns a tuple with the ActivityDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryAudit) GetActivityDisplayNameOk() (*string, bool) {
	if o == nil || o.ActivityDisplayName == nil {
		return nil, false
	}
	return o.ActivityDisplayName, true
}

// HasActivityDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasActivityDisplayName() bool {
	if o != nil && o.ActivityDisplayName != nil {
		return true
	}

	return false
}

// SetActivityDisplayName gets a reference to the given string and assigns it to the ActivityDisplayName field.
func (o *MicrosoftGraphDirectoryAudit) SetActivityDisplayName(v string) {
	o.ActivityDisplayName = &v
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryAudit) GetAdditionalDetails() []*AnyOfmicrosoftGraphKeyValue {
	if o == nil || o.AdditionalDetails == nil {
		var ret []*AnyOfmicrosoftGraphKeyValue
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryAudit) GetAdditionalDetailsOk() (*[]*AnyOfmicrosoftGraphKeyValue, bool) {
	if o == nil || o.AdditionalDetails == nil {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasAdditionalDetails() bool {
	if o != nil && o.AdditionalDetails != nil {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given []*AnyOfmicrosoftGraphKeyValue and assigns it to the AdditionalDetails field.
func (o *MicrosoftGraphDirectoryAudit) SetAdditionalDetails(v []*AnyOfmicrosoftGraphKeyValue) {
	o.AdditionalDetails = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryAudit) GetCategory() string {
	if o == nil || o.Category == nil {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryAudit) GetCategoryOk() (*string, bool) {
	if o == nil || o.Category == nil {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasCategory() bool {
	if o != nil && o.Category != nil {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *MicrosoftGraphDirectoryAudit) SetCategory(v string) {
	o.Category = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryAudit) GetCorrelationId() string {
	if o == nil || o.CorrelationId.Get() == nil {
		var ret string
		return ret
	}
	return *o.CorrelationId.Get()
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryAudit) GetCorrelationIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CorrelationId.Get(), o.CorrelationId.IsSet()
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasCorrelationId() bool {
	if o != nil && o.CorrelationId.IsSet() {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given NullableString and assigns it to the CorrelationId field.
func (o *MicrosoftGraphDirectoryAudit) SetCorrelationId(v string) {
	o.CorrelationId.Set(&v)
}
// SetCorrelationIdNil sets the value for CorrelationId to be an explicit nil
func (o *MicrosoftGraphDirectoryAudit) SetCorrelationIdNil() {
	o.CorrelationId.Set(nil)
}

// UnsetCorrelationId ensures that no value is present for CorrelationId, not even an explicit nil
func (o *MicrosoftGraphDirectoryAudit) UnsetCorrelationId() {
	o.CorrelationId.Unset()
}

// GetInitiatedBy returns the InitiatedBy field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryAudit) GetInitiatedBy() MicrosoftGraphAuditActivityInitiator {
	if o == nil || o.InitiatedBy == nil {
		var ret MicrosoftGraphAuditActivityInitiator
		return ret
	}
	return *o.InitiatedBy
}

// GetInitiatedByOk returns a tuple with the InitiatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryAudit) GetInitiatedByOk() (*MicrosoftGraphAuditActivityInitiator, bool) {
	if o == nil || o.InitiatedBy == nil {
		return nil, false
	}
	return o.InitiatedBy, true
}

// HasInitiatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasInitiatedBy() bool {
	if o != nil && o.InitiatedBy != nil {
		return true
	}

	return false
}

// SetInitiatedBy gets a reference to the given MicrosoftGraphAuditActivityInitiator and assigns it to the InitiatedBy field.
func (o *MicrosoftGraphDirectoryAudit) SetInitiatedBy(v MicrosoftGraphAuditActivityInitiator) {
	o.InitiatedBy = &v
}

// GetLoggedByService returns the LoggedByService field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryAudit) GetLoggedByService() string {
	if o == nil || o.LoggedByService.Get() == nil {
		var ret string
		return ret
	}
	return *o.LoggedByService.Get()
}

// GetLoggedByServiceOk returns a tuple with the LoggedByService field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryAudit) GetLoggedByServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LoggedByService.Get(), o.LoggedByService.IsSet()
}

// HasLoggedByService returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasLoggedByService() bool {
	if o != nil && o.LoggedByService.IsSet() {
		return true
	}

	return false
}

// SetLoggedByService gets a reference to the given NullableString and assigns it to the LoggedByService field.
func (o *MicrosoftGraphDirectoryAudit) SetLoggedByService(v string) {
	o.LoggedByService.Set(&v)
}
// SetLoggedByServiceNil sets the value for LoggedByService to be an explicit nil
func (o *MicrosoftGraphDirectoryAudit) SetLoggedByServiceNil() {
	o.LoggedByService.Set(nil)
}

// UnsetLoggedByService ensures that no value is present for LoggedByService, not even an explicit nil
func (o *MicrosoftGraphDirectoryAudit) UnsetLoggedByService() {
	o.LoggedByService.Unset()
}

// GetOperationType returns the OperationType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryAudit) GetOperationType() string {
	if o == nil || o.OperationType.Get() == nil {
		var ret string
		return ret
	}
	return *o.OperationType.Get()
}

// GetOperationTypeOk returns a tuple with the OperationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryAudit) GetOperationTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OperationType.Get(), o.OperationType.IsSet()
}

// HasOperationType returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasOperationType() bool {
	if o != nil && o.OperationType.IsSet() {
		return true
	}

	return false
}

// SetOperationType gets a reference to the given NullableString and assigns it to the OperationType field.
func (o *MicrosoftGraphDirectoryAudit) SetOperationType(v string) {
	o.OperationType.Set(&v)
}
// SetOperationTypeNil sets the value for OperationType to be an explicit nil
func (o *MicrosoftGraphDirectoryAudit) SetOperationTypeNil() {
	o.OperationType.Set(nil)
}

// UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
func (o *MicrosoftGraphDirectoryAudit) UnsetOperationType() {
	o.OperationType.Unset()
}

// GetResult returns the Result field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryAudit) GetResult() AnyOfmicrosoftGraphOperationResult {
	if o == nil  {
		var ret AnyOfmicrosoftGraphOperationResult
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryAudit) GetResultOk() (*AnyOfmicrosoftGraphOperationResult, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return &o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given AnyOfmicrosoftGraphOperationResult and assigns it to the Result field.
func (o *MicrosoftGraphDirectoryAudit) SetResult(v AnyOfmicrosoftGraphOperationResult) {
	o.Result = v
}

// GetResultReason returns the ResultReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphDirectoryAudit) GetResultReason() string {
	if o == nil || o.ResultReason.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResultReason.Get()
}

// GetResultReasonOk returns a tuple with the ResultReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphDirectoryAudit) GetResultReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResultReason.Get(), o.ResultReason.IsSet()
}

// HasResultReason returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasResultReason() bool {
	if o != nil && o.ResultReason.IsSet() {
		return true
	}

	return false
}

// SetResultReason gets a reference to the given NullableString and assigns it to the ResultReason field.
func (o *MicrosoftGraphDirectoryAudit) SetResultReason(v string) {
	o.ResultReason.Set(&v)
}
// SetResultReasonNil sets the value for ResultReason to be an explicit nil
func (o *MicrosoftGraphDirectoryAudit) SetResultReasonNil() {
	o.ResultReason.Set(nil)
}

// UnsetResultReason ensures that no value is present for ResultReason, not even an explicit nil
func (o *MicrosoftGraphDirectoryAudit) UnsetResultReason() {
	o.ResultReason.Unset()
}

// GetTargetResources returns the TargetResources field value if set, zero value otherwise.
func (o *MicrosoftGraphDirectoryAudit) GetTargetResources() []*AnyOfmicrosoftGraphTargetResource {
	if o == nil || o.TargetResources == nil {
		var ret []*AnyOfmicrosoftGraphTargetResource
		return ret
	}
	return *o.TargetResources
}

// GetTargetResourcesOk returns a tuple with the TargetResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphDirectoryAudit) GetTargetResourcesOk() (*[]*AnyOfmicrosoftGraphTargetResource, bool) {
	if o == nil || o.TargetResources == nil {
		return nil, false
	}
	return o.TargetResources, true
}

// HasTargetResources returns a boolean if a field has been set.
func (o *MicrosoftGraphDirectoryAudit) HasTargetResources() bool {
	if o != nil && o.TargetResources != nil {
		return true
	}

	return false
}

// SetTargetResources gets a reference to the given []*AnyOfmicrosoftGraphTargetResource and assigns it to the TargetResources field.
func (o *MicrosoftGraphDirectoryAudit) SetTargetResources(v []*AnyOfmicrosoftGraphTargetResource) {
	o.TargetResources = &v
}

func (o MicrosoftGraphDirectoryAudit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ActivityDateTime != nil {
		toSerialize["activityDateTime"] = o.ActivityDateTime
	}
	if o.ActivityDisplayName != nil {
		toSerialize["activityDisplayName"] = o.ActivityDisplayName
	}
	if o.AdditionalDetails != nil {
		toSerialize["additionalDetails"] = o.AdditionalDetails
	}
	if o.Category != nil {
		toSerialize["category"] = o.Category
	}
	if o.CorrelationId.IsSet() {
		toSerialize["correlationId"] = o.CorrelationId.Get()
	}
	if o.InitiatedBy != nil {
		toSerialize["initiatedBy"] = o.InitiatedBy
	}
	if o.LoggedByService.IsSet() {
		toSerialize["loggedByService"] = o.LoggedByService.Get()
	}
	if o.OperationType.IsSet() {
		toSerialize["operationType"] = o.OperationType.Get()
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.ResultReason.IsSet() {
		toSerialize["resultReason"] = o.ResultReason.Get()
	}
	if o.TargetResources != nil {
		toSerialize["targetResources"] = o.TargetResources
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphDirectoryAudit struct {
	value *MicrosoftGraphDirectoryAudit
	isSet bool
}

func (v NullableMicrosoftGraphDirectoryAudit) Get() *MicrosoftGraphDirectoryAudit {
	return v.value
}

func (v *NullableMicrosoftGraphDirectoryAudit) Set(val *MicrosoftGraphDirectoryAudit) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphDirectoryAudit) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphDirectoryAudit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphDirectoryAudit(val *MicrosoftGraphDirectoryAudit) *NullableMicrosoftGraphDirectoryAudit {
	return &NullableMicrosoftGraphDirectoryAudit{value: val, isSet: true}
}

func (v NullableMicrosoftGraphDirectoryAudit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphDirectoryAudit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


