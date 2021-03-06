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

// MicrosoftGraphUserConsentRequest struct for MicrosoftGraphUserConsentRequest
type MicrosoftGraphUserConsentRequest struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The identifier of the approval of the request.
	ApprovalId NullableString `json:"approvalId,omitempty"`
	// The request completion date time.
	CompletedDateTime NullableTime `json:"completedDateTime,omitempty"`
	// The user who created this request.
	CreatedBy AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	// The request creation date time.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Free text field to define any custom data for the request. Not used.
	CustomData NullableString `json:"customData,omitempty"`
	// The status of the request. Not nullable. The possible values are: Canceled, Denied, Failed, Granted, PendingAdminDecision, PendingApproval, PendingProvisioning, PendingScheduleCreation, Provisioned, Revoked, and ScheduleCreated. Not nullable.
	Status *string `json:"status,omitempty"`
	// The user's justification for requiring access to the app. Supports $filter (eq only) and $orderby.
	Reason NullableString `json:"reason,omitempty"`
	// Approval decisions associated with a request.
	Approval AnyOfmicrosoftGraphApproval `json:"approval,omitempty"`
}

// NewMicrosoftGraphUserConsentRequest instantiates a new MicrosoftGraphUserConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphUserConsentRequest() *MicrosoftGraphUserConsentRequest {
	this := MicrosoftGraphUserConsentRequest{}
	return &this
}

// NewMicrosoftGraphUserConsentRequestWithDefaults instantiates a new MicrosoftGraphUserConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphUserConsentRequestWithDefaults() *MicrosoftGraphUserConsentRequest {
	this := MicrosoftGraphUserConsentRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphUserConsentRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserConsentRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphUserConsentRequest) SetId(v string) {
	o.Id = &v
}

// GetApprovalId returns the ApprovalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserConsentRequest) GetApprovalId() string {
	if o == nil || o.ApprovalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApprovalId.Get()
}

// GetApprovalIdOk returns a tuple with the ApprovalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserConsentRequest) GetApprovalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApprovalId.Get(), o.ApprovalId.IsSet()
}

// HasApprovalId returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasApprovalId() bool {
	if o != nil && o.ApprovalId.IsSet() {
		return true
	}

	return false
}

// SetApprovalId gets a reference to the given NullableString and assigns it to the ApprovalId field.
func (o *MicrosoftGraphUserConsentRequest) SetApprovalId(v string) {
	o.ApprovalId.Set(&v)
}
// SetApprovalIdNil sets the value for ApprovalId to be an explicit nil
func (o *MicrosoftGraphUserConsentRequest) SetApprovalIdNil() {
	o.ApprovalId.Set(nil)
}

// UnsetApprovalId ensures that no value is present for ApprovalId, not even an explicit nil
func (o *MicrosoftGraphUserConsentRequest) UnsetApprovalId() {
	o.ApprovalId.Unset()
}

// GetCompletedDateTime returns the CompletedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserConsentRequest) GetCompletedDateTime() time.Time {
	if o == nil || o.CompletedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CompletedDateTime.Get()
}

// GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserConsentRequest) GetCompletedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CompletedDateTime.Get(), o.CompletedDateTime.IsSet()
}

// HasCompletedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasCompletedDateTime() bool {
	if o != nil && o.CompletedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCompletedDateTime gets a reference to the given NullableTime and assigns it to the CompletedDateTime field.
func (o *MicrosoftGraphUserConsentRequest) SetCompletedDateTime(v time.Time) {
	o.CompletedDateTime.Set(&v)
}
// SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil
func (o *MicrosoftGraphUserConsentRequest) SetCompletedDateTimeNil() {
	o.CompletedDateTime.Set(nil)
}

// UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
func (o *MicrosoftGraphUserConsentRequest) UnsetCompletedDateTime() {
	o.CompletedDateTime.Unset()
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserConsentRequest) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserConsentRequest) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphUserConsentRequest) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserConsentRequest) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserConsentRequest) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphUserConsentRequest) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphUserConsentRequest) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphUserConsentRequest) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetCustomData returns the CustomData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserConsentRequest) GetCustomData() string {
	if o == nil || o.CustomData.Get() == nil {
		var ret string
		return ret
	}
	return *o.CustomData.Get()
}

// GetCustomDataOk returns a tuple with the CustomData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserConsentRequest) GetCustomDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CustomData.Get(), o.CustomData.IsSet()
}

// HasCustomData returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasCustomData() bool {
	if o != nil && o.CustomData.IsSet() {
		return true
	}

	return false
}

// SetCustomData gets a reference to the given NullableString and assigns it to the CustomData field.
func (o *MicrosoftGraphUserConsentRequest) SetCustomData(v string) {
	o.CustomData.Set(&v)
}
// SetCustomDataNil sets the value for CustomData to be an explicit nil
func (o *MicrosoftGraphUserConsentRequest) SetCustomDataNil() {
	o.CustomData.Set(nil)
}

// UnsetCustomData ensures that no value is present for CustomData, not even an explicit nil
func (o *MicrosoftGraphUserConsentRequest) UnsetCustomData() {
	o.CustomData.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MicrosoftGraphUserConsentRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphUserConsentRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *MicrosoftGraphUserConsentRequest) SetStatus(v string) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserConsentRequest) GetReason() string {
	if o == nil || o.Reason.Get() == nil {
		var ret string
		return ret
	}
	return *o.Reason.Get()
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserConsentRequest) GetReasonOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Reason.Get(), o.Reason.IsSet()
}

// HasReason returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasReason() bool {
	if o != nil && o.Reason.IsSet() {
		return true
	}

	return false
}

// SetReason gets a reference to the given NullableString and assigns it to the Reason field.
func (o *MicrosoftGraphUserConsentRequest) SetReason(v string) {
	o.Reason.Set(&v)
}
// SetReasonNil sets the value for Reason to be an explicit nil
func (o *MicrosoftGraphUserConsentRequest) SetReasonNil() {
	o.Reason.Set(nil)
}

// UnsetReason ensures that no value is present for Reason, not even an explicit nil
func (o *MicrosoftGraphUserConsentRequest) UnsetReason() {
	o.Reason.Unset()
}

// GetApproval returns the Approval field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphUserConsentRequest) GetApproval() AnyOfmicrosoftGraphApproval {
	if o == nil  {
		var ret AnyOfmicrosoftGraphApproval
		return ret
	}
	return o.Approval
}

// GetApprovalOk returns a tuple with the Approval field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphUserConsentRequest) GetApprovalOk() (*AnyOfmicrosoftGraphApproval, bool) {
	if o == nil || o.Approval == nil {
		return nil, false
	}
	return &o.Approval, true
}

// HasApproval returns a boolean if a field has been set.
func (o *MicrosoftGraphUserConsentRequest) HasApproval() bool {
	if o != nil && o.Approval != nil {
		return true
	}

	return false
}

// SetApproval gets a reference to the given AnyOfmicrosoftGraphApproval and assigns it to the Approval field.
func (o *MicrosoftGraphUserConsentRequest) SetApproval(v AnyOfmicrosoftGraphApproval) {
	o.Approval = v
}

func (o MicrosoftGraphUserConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ApprovalId.IsSet() {
		toSerialize["approvalId"] = o.ApprovalId.Get()
	}
	if o.CompletedDateTime.IsSet() {
		toSerialize["completedDateTime"] = o.CompletedDateTime.Get()
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.CustomData.IsSet() {
		toSerialize["customData"] = o.CustomData.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Reason.IsSet() {
		toSerialize["reason"] = o.Reason.Get()
	}
	if o.Approval != nil {
		toSerialize["approval"] = o.Approval
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphUserConsentRequest struct {
	value *MicrosoftGraphUserConsentRequest
	isSet bool
}

func (v NullableMicrosoftGraphUserConsentRequest) Get() *MicrosoftGraphUserConsentRequest {
	return v.value
}

func (v *NullableMicrosoftGraphUserConsentRequest) Set(val *MicrosoftGraphUserConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphUserConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphUserConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphUserConsentRequest(val *MicrosoftGraphUserConsentRequest) *NullableMicrosoftGraphUserConsentRequest {
	return &NullableMicrosoftGraphUserConsentRequest{value: val, isSet: true}
}

func (v NullableMicrosoftGraphUserConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphUserConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


