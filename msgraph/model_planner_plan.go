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

// PlannerPlan struct for PlannerPlan
type PlannerPlan struct {
	// Read-only. The user who created the plan.
	CreatedBy AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	// Read-only. Date and time at which the plan is created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// ID of the Group that owns the plan. A valid group must exist before this field can be set. After it is set, this property can’t be updated.
	Owner NullableString `json:"owner,omitempty"`
	// Required. Title of the plan.
	Title *string `json:"title,omitempty"`
	// Read-only. Nullable. Collection of buckets in the plan.
	Buckets *[]MicrosoftGraphPlannerBucket `json:"buckets,omitempty"`
	// Read-only. Nullable. Additional details about the plan.
	Details AnyOfmicrosoftGraphPlannerPlanDetails `json:"details,omitempty"`
	// Read-only. Nullable. Collection of tasks in the plan.
	Tasks *[]MicrosoftGraphPlannerTask `json:"tasks,omitempty"`
}

// NewPlannerPlan instantiates a new PlannerPlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlannerPlan() *PlannerPlan {
	this := PlannerPlan{}
	return &this
}

// NewPlannerPlanWithDefaults instantiates a new PlannerPlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlannerPlanWithDefaults() *PlannerPlan {
	this := PlannerPlan{}
	return &this
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerPlan) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerPlan) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *PlannerPlan) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *PlannerPlan) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerPlan) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerPlan) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *PlannerPlan) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *PlannerPlan) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *PlannerPlan) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *PlannerPlan) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetOwner returns the Owner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerPlan) GetOwner() string {
	if o == nil || o.Owner.Get() == nil {
		var ret string
		return ret
	}
	return *o.Owner.Get()
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerPlan) GetOwnerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Owner.Get(), o.Owner.IsSet()
}

// HasOwner returns a boolean if a field has been set.
func (o *PlannerPlan) HasOwner() bool {
	if o != nil && o.Owner.IsSet() {
		return true
	}

	return false
}

// SetOwner gets a reference to the given NullableString and assigns it to the Owner field.
func (o *PlannerPlan) SetOwner(v string) {
	o.Owner.Set(&v)
}
// SetOwnerNil sets the value for Owner to be an explicit nil
func (o *PlannerPlan) SetOwnerNil() {
	o.Owner.Set(nil)
}

// UnsetOwner ensures that no value is present for Owner, not even an explicit nil
func (o *PlannerPlan) UnsetOwner() {
	o.Owner.Unset()
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *PlannerPlan) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerPlan) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *PlannerPlan) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *PlannerPlan) SetTitle(v string) {
	o.Title = &v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *PlannerPlan) GetBuckets() []MicrosoftGraphPlannerBucket {
	if o == nil || o.Buckets == nil {
		var ret []MicrosoftGraphPlannerBucket
		return ret
	}
	return *o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerPlan) GetBucketsOk() (*[]MicrosoftGraphPlannerBucket, bool) {
	if o == nil || o.Buckets == nil {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *PlannerPlan) HasBuckets() bool {
	if o != nil && o.Buckets != nil {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []MicrosoftGraphPlannerBucket and assigns it to the Buckets field.
func (o *PlannerPlan) SetBuckets(v []MicrosoftGraphPlannerBucket) {
	o.Buckets = &v
}

// GetDetails returns the Details field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PlannerPlan) GetDetails() AnyOfmicrosoftGraphPlannerPlanDetails {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPlannerPlanDetails
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PlannerPlan) GetDetailsOk() (*AnyOfmicrosoftGraphPlannerPlanDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return &o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *PlannerPlan) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given AnyOfmicrosoftGraphPlannerPlanDetails and assigns it to the Details field.
func (o *PlannerPlan) SetDetails(v AnyOfmicrosoftGraphPlannerPlanDetails) {
	o.Details = v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *PlannerPlan) GetTasks() []MicrosoftGraphPlannerTask {
	if o == nil || o.Tasks == nil {
		var ret []MicrosoftGraphPlannerTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlannerPlan) GetTasksOk() (*[]MicrosoftGraphPlannerTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *PlannerPlan) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []MicrosoftGraphPlannerTask and assigns it to the Tasks field.
func (o *PlannerPlan) SetTasks(v []MicrosoftGraphPlannerTask) {
	o.Tasks = &v
}

func (o PlannerPlan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.Owner.IsSet() {
		toSerialize["owner"] = o.Owner.Get()
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Buckets != nil {
		toSerialize["buckets"] = o.Buckets
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullablePlannerPlan struct {
	value *PlannerPlan
	isSet bool
}

func (v NullablePlannerPlan) Get() *PlannerPlan {
	return v.value
}

func (v *NullablePlannerPlan) Set(val *PlannerPlan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlannerPlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlannerPlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlannerPlan(val *PlannerPlan) *NullablePlannerPlan {
	return &NullablePlannerPlan{value: val, isSet: true}
}

func (v NullablePlannerPlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlannerPlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

