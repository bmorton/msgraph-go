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

// ManagedAppRegistration The ManagedAppEntity is the base entity type for all other entity types under app management workflow.
type ManagedAppRegistration struct {
	// The app package Identifier
	AppIdentifier AnyOfobject `json:"appIdentifier,omitempty"`
	// App version
	ApplicationVersion NullableString `json:"applicationVersion,omitempty"`
	// Date and time of creation
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Host device name
	DeviceName NullableString `json:"deviceName,omitempty"`
	// App management SDK generated tag, which helps relate apps hosted on the same device. Not guaranteed to relate apps in all conditions.
	DeviceTag NullableString `json:"deviceTag,omitempty"`
	// Host device type
	DeviceType NullableString `json:"deviceType,omitempty"`
	// Zero or more reasons an app registration is flagged. E.g. app running on rooted device
	FlaggedReasons *[]AnyOfmicrosoftGraphManagedAppFlaggedReason `json:"flaggedReasons,omitempty"`
	// Date and time of last the app synced with management service.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// App management SDK version
	ManagementSdkVersion NullableString `json:"managementSdkVersion,omitempty"`
	// Operating System version
	PlatformVersion NullableString `json:"platformVersion,omitempty"`
	// The user Id to who this app registration belongs.
	UserId NullableString `json:"userId,omitempty"`
	// Version of the entity.
	Version NullableString `json:"version,omitempty"`
	// Zero or more policys already applied on the registered app when it last synchronized with managment service.
	AppliedPolicies *[]MicrosoftGraphManagedAppPolicy `json:"appliedPolicies,omitempty"`
	// Zero or more policies admin intended for the app as of now.
	IntendedPolicies *[]MicrosoftGraphManagedAppPolicy `json:"intendedPolicies,omitempty"`
	// Zero or more long running operations triggered on the app registration.
	Operations *[]MicrosoftGraphManagedAppOperation `json:"operations,omitempty"`
}

// NewManagedAppRegistration instantiates a new ManagedAppRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedAppRegistration() *ManagedAppRegistration {
	this := ManagedAppRegistration{}
	return &this
}

// NewManagedAppRegistrationWithDefaults instantiates a new ManagedAppRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedAppRegistrationWithDefaults() *ManagedAppRegistration {
	this := ManagedAppRegistration{}
	return &this
}

// GetAppIdentifier returns the AppIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetAppIdentifier() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.AppIdentifier
}

// GetAppIdentifierOk returns a tuple with the AppIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetAppIdentifierOk() (*AnyOfobject, bool) {
	if o == nil || o.AppIdentifier == nil {
		return nil, false
	}
	return &o.AppIdentifier, true
}

// HasAppIdentifier returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasAppIdentifier() bool {
	if o != nil && o.AppIdentifier != nil {
		return true
	}

	return false
}

// SetAppIdentifier gets a reference to the given AnyOfobject and assigns it to the AppIdentifier field.
func (o *ManagedAppRegistration) SetAppIdentifier(v AnyOfobject) {
	o.AppIdentifier = v
}

// GetApplicationVersion returns the ApplicationVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetApplicationVersion() string {
	if o == nil || o.ApplicationVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.ApplicationVersion.Get()
}

// GetApplicationVersionOk returns a tuple with the ApplicationVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetApplicationVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationVersion.Get(), o.ApplicationVersion.IsSet()
}

// HasApplicationVersion returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasApplicationVersion() bool {
	if o != nil && o.ApplicationVersion.IsSet() {
		return true
	}

	return false
}

// SetApplicationVersion gets a reference to the given NullableString and assigns it to the ApplicationVersion field.
func (o *ManagedAppRegistration) SetApplicationVersion(v string) {
	o.ApplicationVersion.Set(&v)
}
// SetApplicationVersionNil sets the value for ApplicationVersion to be an explicit nil
func (o *ManagedAppRegistration) SetApplicationVersionNil() {
	o.ApplicationVersion.Set(nil)
}

// UnsetApplicationVersion ensures that no value is present for ApplicationVersion, not even an explicit nil
func (o *ManagedAppRegistration) UnsetApplicationVersion() {
	o.ApplicationVersion.Unset()
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *ManagedAppRegistration) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *ManagedAppRegistration) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetDeviceName() string {
	if o == nil || o.DeviceName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceName.Get()
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetDeviceNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceName.Get(), o.DeviceName.IsSet()
}

// HasDeviceName returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasDeviceName() bool {
	if o != nil && o.DeviceName.IsSet() {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given NullableString and assigns it to the DeviceName field.
func (o *ManagedAppRegistration) SetDeviceName(v string) {
	o.DeviceName.Set(&v)
}
// SetDeviceNameNil sets the value for DeviceName to be an explicit nil
func (o *ManagedAppRegistration) SetDeviceNameNil() {
	o.DeviceName.Set(nil)
}

// UnsetDeviceName ensures that no value is present for DeviceName, not even an explicit nil
func (o *ManagedAppRegistration) UnsetDeviceName() {
	o.DeviceName.Unset()
}

// GetDeviceTag returns the DeviceTag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetDeviceTag() string {
	if o == nil || o.DeviceTag.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceTag.Get()
}

// GetDeviceTagOk returns a tuple with the DeviceTag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetDeviceTagOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceTag.Get(), o.DeviceTag.IsSet()
}

// HasDeviceTag returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasDeviceTag() bool {
	if o != nil && o.DeviceTag.IsSet() {
		return true
	}

	return false
}

// SetDeviceTag gets a reference to the given NullableString and assigns it to the DeviceTag field.
func (o *ManagedAppRegistration) SetDeviceTag(v string) {
	o.DeviceTag.Set(&v)
}
// SetDeviceTagNil sets the value for DeviceTag to be an explicit nil
func (o *ManagedAppRegistration) SetDeviceTagNil() {
	o.DeviceTag.Set(nil)
}

// UnsetDeviceTag ensures that no value is present for DeviceTag, not even an explicit nil
func (o *ManagedAppRegistration) UnsetDeviceTag() {
	o.DeviceTag.Unset()
}

// GetDeviceType returns the DeviceType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetDeviceType() string {
	if o == nil || o.DeviceType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DeviceType.Get()
}

// GetDeviceTypeOk returns a tuple with the DeviceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetDeviceTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DeviceType.Get(), o.DeviceType.IsSet()
}

// HasDeviceType returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasDeviceType() bool {
	if o != nil && o.DeviceType.IsSet() {
		return true
	}

	return false
}

// SetDeviceType gets a reference to the given NullableString and assigns it to the DeviceType field.
func (o *ManagedAppRegistration) SetDeviceType(v string) {
	o.DeviceType.Set(&v)
}
// SetDeviceTypeNil sets the value for DeviceType to be an explicit nil
func (o *ManagedAppRegistration) SetDeviceTypeNil() {
	o.DeviceType.Set(nil)
}

// UnsetDeviceType ensures that no value is present for DeviceType, not even an explicit nil
func (o *ManagedAppRegistration) UnsetDeviceType() {
	o.DeviceType.Unset()
}

// GetFlaggedReasons returns the FlaggedReasons field value if set, zero value otherwise.
func (o *ManagedAppRegistration) GetFlaggedReasons() []AnyOfmicrosoftGraphManagedAppFlaggedReason {
	if o == nil || o.FlaggedReasons == nil {
		var ret []AnyOfmicrosoftGraphManagedAppFlaggedReason
		return ret
	}
	return *o.FlaggedReasons
}

// GetFlaggedReasonsOk returns a tuple with the FlaggedReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetFlaggedReasonsOk() (*[]AnyOfmicrosoftGraphManagedAppFlaggedReason, bool) {
	if o == nil || o.FlaggedReasons == nil {
		return nil, false
	}
	return o.FlaggedReasons, true
}

// HasFlaggedReasons returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasFlaggedReasons() bool {
	if o != nil && o.FlaggedReasons != nil {
		return true
	}

	return false
}

// SetFlaggedReasons gets a reference to the given []AnyOfmicrosoftGraphManagedAppFlaggedReason and assigns it to the FlaggedReasons field.
func (o *ManagedAppRegistration) SetFlaggedReasons(v []AnyOfmicrosoftGraphManagedAppFlaggedReason) {
	o.FlaggedReasons = &v
}

// GetLastSyncDateTime returns the LastSyncDateTime field value if set, zero value otherwise.
func (o *ManagedAppRegistration) GetLastSyncDateTime() time.Time {
	if o == nil || o.LastSyncDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSyncDateTime
}

// GetLastSyncDateTimeOk returns a tuple with the LastSyncDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetLastSyncDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastSyncDateTime == nil {
		return nil, false
	}
	return o.LastSyncDateTime, true
}

// HasLastSyncDateTime returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasLastSyncDateTime() bool {
	if o != nil && o.LastSyncDateTime != nil {
		return true
	}

	return false
}

// SetLastSyncDateTime gets a reference to the given time.Time and assigns it to the LastSyncDateTime field.
func (o *ManagedAppRegistration) SetLastSyncDateTime(v time.Time) {
	o.LastSyncDateTime = &v
}

// GetManagementSdkVersion returns the ManagementSdkVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetManagementSdkVersion() string {
	if o == nil || o.ManagementSdkVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.ManagementSdkVersion.Get()
}

// GetManagementSdkVersionOk returns a tuple with the ManagementSdkVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetManagementSdkVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ManagementSdkVersion.Get(), o.ManagementSdkVersion.IsSet()
}

// HasManagementSdkVersion returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasManagementSdkVersion() bool {
	if o != nil && o.ManagementSdkVersion.IsSet() {
		return true
	}

	return false
}

// SetManagementSdkVersion gets a reference to the given NullableString and assigns it to the ManagementSdkVersion field.
func (o *ManagedAppRegistration) SetManagementSdkVersion(v string) {
	o.ManagementSdkVersion.Set(&v)
}
// SetManagementSdkVersionNil sets the value for ManagementSdkVersion to be an explicit nil
func (o *ManagedAppRegistration) SetManagementSdkVersionNil() {
	o.ManagementSdkVersion.Set(nil)
}

// UnsetManagementSdkVersion ensures that no value is present for ManagementSdkVersion, not even an explicit nil
func (o *ManagedAppRegistration) UnsetManagementSdkVersion() {
	o.ManagementSdkVersion.Unset()
}

// GetPlatformVersion returns the PlatformVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetPlatformVersion() string {
	if o == nil || o.PlatformVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.PlatformVersion.Get()
}

// GetPlatformVersionOk returns a tuple with the PlatformVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetPlatformVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PlatformVersion.Get(), o.PlatformVersion.IsSet()
}

// HasPlatformVersion returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasPlatformVersion() bool {
	if o != nil && o.PlatformVersion.IsSet() {
		return true
	}

	return false
}

// SetPlatformVersion gets a reference to the given NullableString and assigns it to the PlatformVersion field.
func (o *ManagedAppRegistration) SetPlatformVersion(v string) {
	o.PlatformVersion.Set(&v)
}
// SetPlatformVersionNil sets the value for PlatformVersion to be an explicit nil
func (o *ManagedAppRegistration) SetPlatformVersionNil() {
	o.PlatformVersion.Set(nil)
}

// UnsetPlatformVersion ensures that no value is present for PlatformVersion, not even an explicit nil
func (o *ManagedAppRegistration) UnsetPlatformVersion() {
	o.PlatformVersion.Unset()
}

// GetUserId returns the UserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetUserId() string {
	if o == nil || o.UserId.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserId.Get()
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserId.Get(), o.UserId.IsSet()
}

// HasUserId returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasUserId() bool {
	if o != nil && o.UserId.IsSet() {
		return true
	}

	return false
}

// SetUserId gets a reference to the given NullableString and assigns it to the UserId field.
func (o *ManagedAppRegistration) SetUserId(v string) {
	o.UserId.Set(&v)
}
// SetUserIdNil sets the value for UserId to be an explicit nil
func (o *ManagedAppRegistration) SetUserIdNil() {
	o.UserId.Set(nil)
}

// UnsetUserId ensures that no value is present for UserId, not even an explicit nil
func (o *ManagedAppRegistration) UnsetUserId() {
	o.UserId.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ManagedAppRegistration) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ManagedAppRegistration) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *ManagedAppRegistration) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *ManagedAppRegistration) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *ManagedAppRegistration) UnsetVersion() {
	o.Version.Unset()
}

// GetAppliedPolicies returns the AppliedPolicies field value if set, zero value otherwise.
func (o *ManagedAppRegistration) GetAppliedPolicies() []MicrosoftGraphManagedAppPolicy {
	if o == nil || o.AppliedPolicies == nil {
		var ret []MicrosoftGraphManagedAppPolicy
		return ret
	}
	return *o.AppliedPolicies
}

// GetAppliedPoliciesOk returns a tuple with the AppliedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetAppliedPoliciesOk() (*[]MicrosoftGraphManagedAppPolicy, bool) {
	if o == nil || o.AppliedPolicies == nil {
		return nil, false
	}
	return o.AppliedPolicies, true
}

// HasAppliedPolicies returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasAppliedPolicies() bool {
	if o != nil && o.AppliedPolicies != nil {
		return true
	}

	return false
}

// SetAppliedPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the AppliedPolicies field.
func (o *ManagedAppRegistration) SetAppliedPolicies(v []MicrosoftGraphManagedAppPolicy) {
	o.AppliedPolicies = &v
}

// GetIntendedPolicies returns the IntendedPolicies field value if set, zero value otherwise.
func (o *ManagedAppRegistration) GetIntendedPolicies() []MicrosoftGraphManagedAppPolicy {
	if o == nil || o.IntendedPolicies == nil {
		var ret []MicrosoftGraphManagedAppPolicy
		return ret
	}
	return *o.IntendedPolicies
}

// GetIntendedPoliciesOk returns a tuple with the IntendedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetIntendedPoliciesOk() (*[]MicrosoftGraphManagedAppPolicy, bool) {
	if o == nil || o.IntendedPolicies == nil {
		return nil, false
	}
	return o.IntendedPolicies, true
}

// HasIntendedPolicies returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasIntendedPolicies() bool {
	if o != nil && o.IntendedPolicies != nil {
		return true
	}

	return false
}

// SetIntendedPolicies gets a reference to the given []MicrosoftGraphManagedAppPolicy and assigns it to the IntendedPolicies field.
func (o *ManagedAppRegistration) SetIntendedPolicies(v []MicrosoftGraphManagedAppPolicy) {
	o.IntendedPolicies = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *ManagedAppRegistration) GetOperations() []MicrosoftGraphManagedAppOperation {
	if o == nil || o.Operations == nil {
		var ret []MicrosoftGraphManagedAppOperation
		return ret
	}
	return *o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedAppRegistration) GetOperationsOk() (*[]MicrosoftGraphManagedAppOperation, bool) {
	if o == nil || o.Operations == nil {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *ManagedAppRegistration) HasOperations() bool {
	if o != nil && o.Operations != nil {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []MicrosoftGraphManagedAppOperation and assigns it to the Operations field.
func (o *ManagedAppRegistration) SetOperations(v []MicrosoftGraphManagedAppOperation) {
	o.Operations = &v
}

func (o ManagedAppRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppIdentifier != nil {
		toSerialize["appIdentifier"] = o.AppIdentifier
	}
	if o.ApplicationVersion.IsSet() {
		toSerialize["applicationVersion"] = o.ApplicationVersion.Get()
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.DeviceName.IsSet() {
		toSerialize["deviceName"] = o.DeviceName.Get()
	}
	if o.DeviceTag.IsSet() {
		toSerialize["deviceTag"] = o.DeviceTag.Get()
	}
	if o.DeviceType.IsSet() {
		toSerialize["deviceType"] = o.DeviceType.Get()
	}
	if o.FlaggedReasons != nil {
		toSerialize["flaggedReasons"] = o.FlaggedReasons
	}
	if o.LastSyncDateTime != nil {
		toSerialize["lastSyncDateTime"] = o.LastSyncDateTime
	}
	if o.ManagementSdkVersion.IsSet() {
		toSerialize["managementSdkVersion"] = o.ManagementSdkVersion.Get()
	}
	if o.PlatformVersion.IsSet() {
		toSerialize["platformVersion"] = o.PlatformVersion.Get()
	}
	if o.UserId.IsSet() {
		toSerialize["userId"] = o.UserId.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.AppliedPolicies != nil {
		toSerialize["appliedPolicies"] = o.AppliedPolicies
	}
	if o.IntendedPolicies != nil {
		toSerialize["intendedPolicies"] = o.IntendedPolicies
	}
	if o.Operations != nil {
		toSerialize["operations"] = o.Operations
	}
	return json.Marshal(toSerialize)
}

type NullableManagedAppRegistration struct {
	value *ManagedAppRegistration
	isSet bool
}

func (v NullableManagedAppRegistration) Get() *ManagedAppRegistration {
	return v.value
}

func (v *NullableManagedAppRegistration) Set(val *ManagedAppRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedAppRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedAppRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedAppRegistration(val *ManagedAppRegistration) *NullableManagedAppRegistration {
	return &NullableManagedAppRegistration{value: val, isSet: true}
}

func (v NullableManagedAppRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedAppRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


