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

// AppRoleAssignment struct for AppRoleAssignment
type AppRoleAssignment struct {
	// The identifier (id) for the app role which is assigned to the principal. This app role must be exposed in the appRoles property on the resource application's service principal (resourceId). If the resource application has not declared any app roles, a default app role ID of 00000000-0000-0000-0000-000000000000 can be specified to signal that the principal is assigned to the resource app without any specific app roles. Required on create.
	AppRoleId *string `json:"appRoleId,omitempty"`
	// The time when the app role assignment was created.The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// The display name of the user, group, or service principal that was granted the app role assignment. Read-only. Supports $filter (eq and startswith).
	PrincipalDisplayName NullableString `json:"principalDisplayName,omitempty"`
	// The unique identifier (id) for the user, group or service principal being granted the app role. Required on create.
	PrincipalId NullableString `json:"principalId,omitempty"`
	// The type of the assigned principal. This can either be User, Group or ServicePrincipal. Read-only.
	PrincipalType NullableString `json:"principalType,omitempty"`
	// The display name of the resource app's service principal to which the assignment is made.
	ResourceDisplayName NullableString `json:"resourceDisplayName,omitempty"`
	// The unique identifier (id) for the resource service principal for which the assignment is made. Required on create. Supports $filter (eq only).
	ResourceId NullableString `json:"resourceId,omitempty"`
}

// NewAppRoleAssignment instantiates a new AppRoleAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRoleAssignment() *AppRoleAssignment {
	this := AppRoleAssignment{}
	return &this
}

// NewAppRoleAssignmentWithDefaults instantiates a new AppRoleAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRoleAssignmentWithDefaults() *AppRoleAssignment {
	this := AppRoleAssignment{}
	return &this
}

// GetAppRoleId returns the AppRoleId field value if set, zero value otherwise.
func (o *AppRoleAssignment) GetAppRoleId() string {
	if o == nil || o.AppRoleId == nil {
		var ret string
		return ret
	}
	return *o.AppRoleId
}

// GetAppRoleIdOk returns a tuple with the AppRoleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppRoleAssignment) GetAppRoleIdOk() (*string, bool) {
	if o == nil || o.AppRoleId == nil {
		return nil, false
	}
	return o.AppRoleId, true
}

// HasAppRoleId returns a boolean if a field has been set.
func (o *AppRoleAssignment) HasAppRoleId() bool {
	if o != nil && o.AppRoleId != nil {
		return true
	}

	return false
}

// SetAppRoleId gets a reference to the given string and assigns it to the AppRoleId field.
func (o *AppRoleAssignment) SetAppRoleId(v string) {
	o.AppRoleId = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRoleAssignment) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRoleAssignment) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *AppRoleAssignment) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *AppRoleAssignment) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *AppRoleAssignment) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *AppRoleAssignment) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetPrincipalDisplayName returns the PrincipalDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRoleAssignment) GetPrincipalDisplayName() string {
	if o == nil || o.PrincipalDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalDisplayName.Get()
}

// GetPrincipalDisplayNameOk returns a tuple with the PrincipalDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRoleAssignment) GetPrincipalDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalDisplayName.Get(), o.PrincipalDisplayName.IsSet()
}

// HasPrincipalDisplayName returns a boolean if a field has been set.
func (o *AppRoleAssignment) HasPrincipalDisplayName() bool {
	if o != nil && o.PrincipalDisplayName.IsSet() {
		return true
	}

	return false
}

// SetPrincipalDisplayName gets a reference to the given NullableString and assigns it to the PrincipalDisplayName field.
func (o *AppRoleAssignment) SetPrincipalDisplayName(v string) {
	o.PrincipalDisplayName.Set(&v)
}
// SetPrincipalDisplayNameNil sets the value for PrincipalDisplayName to be an explicit nil
func (o *AppRoleAssignment) SetPrincipalDisplayNameNil() {
	o.PrincipalDisplayName.Set(nil)
}

// UnsetPrincipalDisplayName ensures that no value is present for PrincipalDisplayName, not even an explicit nil
func (o *AppRoleAssignment) UnsetPrincipalDisplayName() {
	o.PrincipalDisplayName.Unset()
}

// GetPrincipalId returns the PrincipalId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRoleAssignment) GetPrincipalId() string {
	if o == nil || o.PrincipalId.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalId.Get()
}

// GetPrincipalIdOk returns a tuple with the PrincipalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRoleAssignment) GetPrincipalIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalId.Get(), o.PrincipalId.IsSet()
}

// HasPrincipalId returns a boolean if a field has been set.
func (o *AppRoleAssignment) HasPrincipalId() bool {
	if o != nil && o.PrincipalId.IsSet() {
		return true
	}

	return false
}

// SetPrincipalId gets a reference to the given NullableString and assigns it to the PrincipalId field.
func (o *AppRoleAssignment) SetPrincipalId(v string) {
	o.PrincipalId.Set(&v)
}
// SetPrincipalIdNil sets the value for PrincipalId to be an explicit nil
func (o *AppRoleAssignment) SetPrincipalIdNil() {
	o.PrincipalId.Set(nil)
}

// UnsetPrincipalId ensures that no value is present for PrincipalId, not even an explicit nil
func (o *AppRoleAssignment) UnsetPrincipalId() {
	o.PrincipalId.Unset()
}

// GetPrincipalType returns the PrincipalType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRoleAssignment) GetPrincipalType() string {
	if o == nil || o.PrincipalType.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalType.Get()
}

// GetPrincipalTypeOk returns a tuple with the PrincipalType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRoleAssignment) GetPrincipalTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalType.Get(), o.PrincipalType.IsSet()
}

// HasPrincipalType returns a boolean if a field has been set.
func (o *AppRoleAssignment) HasPrincipalType() bool {
	if o != nil && o.PrincipalType.IsSet() {
		return true
	}

	return false
}

// SetPrincipalType gets a reference to the given NullableString and assigns it to the PrincipalType field.
func (o *AppRoleAssignment) SetPrincipalType(v string) {
	o.PrincipalType.Set(&v)
}
// SetPrincipalTypeNil sets the value for PrincipalType to be an explicit nil
func (o *AppRoleAssignment) SetPrincipalTypeNil() {
	o.PrincipalType.Set(nil)
}

// UnsetPrincipalType ensures that no value is present for PrincipalType, not even an explicit nil
func (o *AppRoleAssignment) UnsetPrincipalType() {
	o.PrincipalType.Unset()
}

// GetResourceDisplayName returns the ResourceDisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRoleAssignment) GetResourceDisplayName() string {
	if o == nil || o.ResourceDisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceDisplayName.Get()
}

// GetResourceDisplayNameOk returns a tuple with the ResourceDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRoleAssignment) GetResourceDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceDisplayName.Get(), o.ResourceDisplayName.IsSet()
}

// HasResourceDisplayName returns a boolean if a field has been set.
func (o *AppRoleAssignment) HasResourceDisplayName() bool {
	if o != nil && o.ResourceDisplayName.IsSet() {
		return true
	}

	return false
}

// SetResourceDisplayName gets a reference to the given NullableString and assigns it to the ResourceDisplayName field.
func (o *AppRoleAssignment) SetResourceDisplayName(v string) {
	o.ResourceDisplayName.Set(&v)
}
// SetResourceDisplayNameNil sets the value for ResourceDisplayName to be an explicit nil
func (o *AppRoleAssignment) SetResourceDisplayNameNil() {
	o.ResourceDisplayName.Set(nil)
}

// UnsetResourceDisplayName ensures that no value is present for ResourceDisplayName, not even an explicit nil
func (o *AppRoleAssignment) UnsetResourceDisplayName() {
	o.ResourceDisplayName.Unset()
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AppRoleAssignment) GetResourceId() string {
	if o == nil || o.ResourceId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceId.Get()
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AppRoleAssignment) GetResourceIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceId.Get(), o.ResourceId.IsSet()
}

// HasResourceId returns a boolean if a field has been set.
func (o *AppRoleAssignment) HasResourceId() bool {
	if o != nil && o.ResourceId.IsSet() {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given NullableString and assigns it to the ResourceId field.
func (o *AppRoleAssignment) SetResourceId(v string) {
	o.ResourceId.Set(&v)
}
// SetResourceIdNil sets the value for ResourceId to be an explicit nil
func (o *AppRoleAssignment) SetResourceIdNil() {
	o.ResourceId.Set(nil)
}

// UnsetResourceId ensures that no value is present for ResourceId, not even an explicit nil
func (o *AppRoleAssignment) UnsetResourceId() {
	o.ResourceId.Unset()
}

func (o AppRoleAssignment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppRoleId != nil {
		toSerialize["appRoleId"] = o.AppRoleId
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.PrincipalDisplayName.IsSet() {
		toSerialize["principalDisplayName"] = o.PrincipalDisplayName.Get()
	}
	if o.PrincipalId.IsSet() {
		toSerialize["principalId"] = o.PrincipalId.Get()
	}
	if o.PrincipalType.IsSet() {
		toSerialize["principalType"] = o.PrincipalType.Get()
	}
	if o.ResourceDisplayName.IsSet() {
		toSerialize["resourceDisplayName"] = o.ResourceDisplayName.Get()
	}
	if o.ResourceId.IsSet() {
		toSerialize["resourceId"] = o.ResourceId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableAppRoleAssignment struct {
	value *AppRoleAssignment
	isSet bool
}

func (v NullableAppRoleAssignment) Get() *AppRoleAssignment {
	return v.value
}

func (v *NullableAppRoleAssignment) Set(val *AppRoleAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRoleAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRoleAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRoleAssignment(val *AppRoleAssignment) *NullableAppRoleAssignment {
	return &NullableAppRoleAssignment{value: val, isSet: true}
}

func (v NullableAppRoleAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRoleAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


