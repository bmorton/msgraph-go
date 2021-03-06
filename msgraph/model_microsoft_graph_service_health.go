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

// MicrosoftGraphServiceHealth struct for MicrosoftGraphServiceHealth
type MicrosoftGraphServiceHealth struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The service name. Use the list healthOverviews operation to get exact string names for services subscribed by the tenant.
	Service *string `json:"service,omitempty"`
	// Show the overral service health status. Possible values are: serviceOperational, investigating, restoringService, verifyingService, serviceRestored, postIncidentReviewPublished, serviceDegradation, serviceInterruption, extendedRecovery, falsePositive, investigationSuspended, resolved, mitigatedExternal, mitigated, resolvedExternal, confirmed, reported, unknownFutureValue.
	Status AnyOfmicrosoftGraphServiceHealthStatus `json:"status,omitempty"`
	// A collection of issues happened on the service, with detailed information for each issue.
	Issues *[]MicrosoftGraphServiceHealthIssue `json:"issues,omitempty"`
}

// NewMicrosoftGraphServiceHealth instantiates a new MicrosoftGraphServiceHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphServiceHealth() *MicrosoftGraphServiceHealth {
	this := MicrosoftGraphServiceHealth{}
	return &this
}

// NewMicrosoftGraphServiceHealthWithDefaults instantiates a new MicrosoftGraphServiceHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphServiceHealthWithDefaults() *MicrosoftGraphServiceHealth {
	this := MicrosoftGraphServiceHealth{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceHealth) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceHealth) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceHealth) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphServiceHealth) SetId(v string) {
	o.Id = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceHealth) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceHealth) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceHealth) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *MicrosoftGraphServiceHealth) SetService(v string) {
	o.Service = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphServiceHealth) GetStatus() AnyOfmicrosoftGraphServiceHealthStatus {
	if o == nil  {
		var ret AnyOfmicrosoftGraphServiceHealthStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphServiceHealth) GetStatusOk() (*AnyOfmicrosoftGraphServiceHealthStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceHealth) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AnyOfmicrosoftGraphServiceHealthStatus and assigns it to the Status field.
func (o *MicrosoftGraphServiceHealth) SetStatus(v AnyOfmicrosoftGraphServiceHealthStatus) {
	o.Status = v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceHealth) GetIssues() []MicrosoftGraphServiceHealthIssue {
	if o == nil || o.Issues == nil {
		var ret []MicrosoftGraphServiceHealthIssue
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceHealth) GetIssuesOk() (*[]MicrosoftGraphServiceHealthIssue, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceHealth) HasIssues() bool {
	if o != nil && o.Issues != nil {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []MicrosoftGraphServiceHealthIssue and assigns it to the Issues field.
func (o *MicrosoftGraphServiceHealth) SetIssues(v []MicrosoftGraphServiceHealthIssue) {
	o.Issues = &v
}

func (o MicrosoftGraphServiceHealth) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Issues != nil {
		toSerialize["issues"] = o.Issues
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphServiceHealth struct {
	value *MicrosoftGraphServiceHealth
	isSet bool
}

func (v NullableMicrosoftGraphServiceHealth) Get() *MicrosoftGraphServiceHealth {
	return v.value
}

func (v *NullableMicrosoftGraphServiceHealth) Set(val *MicrosoftGraphServiceHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphServiceHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphServiceHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphServiceHealth(val *MicrosoftGraphServiceHealth) *NullableMicrosoftGraphServiceHealth {
	return &NullableMicrosoftGraphServiceHealth{value: val, isSet: true}
}

func (v NullableMicrosoftGraphServiceHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphServiceHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


