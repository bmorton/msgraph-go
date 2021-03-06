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

// CalendarPermission struct for CalendarPermission
type CalendarPermission struct {
	// List of allowed sharing or delegating permission levels for the calendar. Possible values are: none, freeBusyRead, limitedRead, read, write, delegateWithoutPrivateEventAccess, delegateWithPrivateEventAccess, custom.
	AllowedRoles *[]*AnyOfmicrosoftGraphCalendarRoleType `json:"allowedRoles,omitempty"`
	// Represents a sharee or delegate who has access to the calendar. For the 'My Organization' sharee, the address property is null. Read-only.
	EmailAddress AnyOfmicrosoftGraphEmailAddress `json:"emailAddress,omitempty"`
	// True if the user in context (sharee or delegate) is inside the same organization as the calendar owner.
	IsInsideOrganization NullableBool `json:"isInsideOrganization,omitempty"`
	// True if the user can be removed from the list of sharees or delegates for the specified calendar, false otherwise. The 'My organization' user determines the permissions other people within your organization have to the given calendar. You cannot remove 'My organization' as a sharee to a calendar.
	IsRemovable NullableBool `json:"isRemovable,omitempty"`
	// Current permission level of the calendar sharee or delegate.
	Role AnyOfmicrosoftGraphCalendarRoleType `json:"role,omitempty"`
}

// NewCalendarPermission instantiates a new CalendarPermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCalendarPermission() *CalendarPermission {
	this := CalendarPermission{}
	return &this
}

// NewCalendarPermissionWithDefaults instantiates a new CalendarPermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCalendarPermissionWithDefaults() *CalendarPermission {
	this := CalendarPermission{}
	return &this
}

// GetAllowedRoles returns the AllowedRoles field value if set, zero value otherwise.
func (o *CalendarPermission) GetAllowedRoles() []*AnyOfmicrosoftGraphCalendarRoleType {
	if o == nil || o.AllowedRoles == nil {
		var ret []*AnyOfmicrosoftGraphCalendarRoleType
		return ret
	}
	return *o.AllowedRoles
}

// GetAllowedRolesOk returns a tuple with the AllowedRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CalendarPermission) GetAllowedRolesOk() (*[]*AnyOfmicrosoftGraphCalendarRoleType, bool) {
	if o == nil || o.AllowedRoles == nil {
		return nil, false
	}
	return o.AllowedRoles, true
}

// HasAllowedRoles returns a boolean if a field has been set.
func (o *CalendarPermission) HasAllowedRoles() bool {
	if o != nil && o.AllowedRoles != nil {
		return true
	}

	return false
}

// SetAllowedRoles gets a reference to the given []*AnyOfmicrosoftGraphCalendarRoleType and assigns it to the AllowedRoles field.
func (o *CalendarPermission) SetAllowedRoles(v []*AnyOfmicrosoftGraphCalendarRoleType) {
	o.AllowedRoles = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CalendarPermission) GetEmailAddress() AnyOfmicrosoftGraphEmailAddress {
	if o == nil  {
		var ret AnyOfmicrosoftGraphEmailAddress
		return ret
	}
	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CalendarPermission) GetEmailAddressOk() (*AnyOfmicrosoftGraphEmailAddress, bool) {
	if o == nil || o.EmailAddress == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *CalendarPermission) HasEmailAddress() bool {
	if o != nil && o.EmailAddress != nil {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given AnyOfmicrosoftGraphEmailAddress and assigns it to the EmailAddress field.
func (o *CalendarPermission) SetEmailAddress(v AnyOfmicrosoftGraphEmailAddress) {
	o.EmailAddress = v
}

// GetIsInsideOrganization returns the IsInsideOrganization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CalendarPermission) GetIsInsideOrganization() bool {
	if o == nil || o.IsInsideOrganization.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsInsideOrganization.Get()
}

// GetIsInsideOrganizationOk returns a tuple with the IsInsideOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CalendarPermission) GetIsInsideOrganizationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsInsideOrganization.Get(), o.IsInsideOrganization.IsSet()
}

// HasIsInsideOrganization returns a boolean if a field has been set.
func (o *CalendarPermission) HasIsInsideOrganization() bool {
	if o != nil && o.IsInsideOrganization.IsSet() {
		return true
	}

	return false
}

// SetIsInsideOrganization gets a reference to the given NullableBool and assigns it to the IsInsideOrganization field.
func (o *CalendarPermission) SetIsInsideOrganization(v bool) {
	o.IsInsideOrganization.Set(&v)
}
// SetIsInsideOrganizationNil sets the value for IsInsideOrganization to be an explicit nil
func (o *CalendarPermission) SetIsInsideOrganizationNil() {
	o.IsInsideOrganization.Set(nil)
}

// UnsetIsInsideOrganization ensures that no value is present for IsInsideOrganization, not even an explicit nil
func (o *CalendarPermission) UnsetIsInsideOrganization() {
	o.IsInsideOrganization.Unset()
}

// GetIsRemovable returns the IsRemovable field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CalendarPermission) GetIsRemovable() bool {
	if o == nil || o.IsRemovable.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsRemovable.Get()
}

// GetIsRemovableOk returns a tuple with the IsRemovable field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CalendarPermission) GetIsRemovableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsRemovable.Get(), o.IsRemovable.IsSet()
}

// HasIsRemovable returns a boolean if a field has been set.
func (o *CalendarPermission) HasIsRemovable() bool {
	if o != nil && o.IsRemovable.IsSet() {
		return true
	}

	return false
}

// SetIsRemovable gets a reference to the given NullableBool and assigns it to the IsRemovable field.
func (o *CalendarPermission) SetIsRemovable(v bool) {
	o.IsRemovable.Set(&v)
}
// SetIsRemovableNil sets the value for IsRemovable to be an explicit nil
func (o *CalendarPermission) SetIsRemovableNil() {
	o.IsRemovable.Set(nil)
}

// UnsetIsRemovable ensures that no value is present for IsRemovable, not even an explicit nil
func (o *CalendarPermission) UnsetIsRemovable() {
	o.IsRemovable.Unset()
}

// GetRole returns the Role field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CalendarPermission) GetRole() AnyOfmicrosoftGraphCalendarRoleType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCalendarRoleType
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CalendarPermission) GetRoleOk() (*AnyOfmicrosoftGraphCalendarRoleType, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return &o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *CalendarPermission) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given AnyOfmicrosoftGraphCalendarRoleType and assigns it to the Role field.
func (o *CalendarPermission) SetRole(v AnyOfmicrosoftGraphCalendarRoleType) {
	o.Role = v
}

func (o CalendarPermission) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowedRoles != nil {
		toSerialize["allowedRoles"] = o.AllowedRoles
	}
	if o.EmailAddress != nil {
		toSerialize["emailAddress"] = o.EmailAddress
	}
	if o.IsInsideOrganization.IsSet() {
		toSerialize["isInsideOrganization"] = o.IsInsideOrganization.Get()
	}
	if o.IsRemovable.IsSet() {
		toSerialize["isRemovable"] = o.IsRemovable.Get()
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableCalendarPermission struct {
	value *CalendarPermission
	isSet bool
}

func (v NullableCalendarPermission) Get() *CalendarPermission {
	return v.value
}

func (v *NullableCalendarPermission) Set(val *CalendarPermission) {
	v.value = val
	v.isSet = true
}

func (v NullableCalendarPermission) IsSet() bool {
	return v.isSet
}

func (v *NullableCalendarPermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalendarPermission(val *CalendarPermission) *NullableCalendarPermission {
	return &NullableCalendarPermission{value: val, isSet: true}
}

func (v NullableCalendarPermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalendarPermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


