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

// MicrosoftGraphPrinterShare struct for MicrosoftGraphPrinterShare
type MicrosoftGraphPrinterShare struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The capabilities of the printer/printerShare.
	Capabilities AnyOfmicrosoftGraphPrinterCapabilities `json:"capabilities,omitempty"`
	// The default print settings of printer/printerShare.
	Defaults AnyOfmicrosoftGraphPrinterDefaults `json:"defaults,omitempty"`
	// The name of the printer/printerShare.
	DisplayName *string `json:"displayName,omitempty"`
	// Whether the printer/printerShare is currently accepting new print jobs.
	IsAcceptingJobs NullableBool `json:"isAcceptingJobs,omitempty"`
	// The physical and/or organizational location of the printer/printerShare.
	Location AnyOfmicrosoftGraphPrinterLocation `json:"location,omitempty"`
	// The manufacturer of the printer/printerShare.
	Manufacturer NullableString `json:"manufacturer,omitempty"`
	// The model name of the printer/printerShare.
	Model NullableString `json:"model,omitempty"`
	Status *MicrosoftGraphPrinterStatus `json:"status,omitempty"`
	// The list of jobs that are queued for printing by the printer/printerShare.
	Jobs *[]MicrosoftGraphPrintJob `json:"jobs,omitempty"`
	// If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined by the allowedUsers and allowedGroups navigation properties.
	AllowAllUsers *bool `json:"allowAllUsers,omitempty"`
	// The DateTimeOffset when the printer share was created. Read-only.
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// The groups whose users have access to print using the printer.
	AllowedGroups *[]MicrosoftGraphGroup `json:"allowedGroups,omitempty"`
	// The users who have access to print using the printer.
	AllowedUsers *[]MicrosoftGraphUser `json:"allowedUsers,omitempty"`
	// The printer that this printer share is related to.
	Printer AnyOfmicrosoftGraphPrinter `json:"printer,omitempty"`
}

// NewMicrosoftGraphPrinterShare instantiates a new MicrosoftGraphPrinterShare object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrinterShare() *MicrosoftGraphPrinterShare {
	this := MicrosoftGraphPrinterShare{}
	return &this
}

// NewMicrosoftGraphPrinterShareWithDefaults instantiates a new MicrosoftGraphPrinterShare object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrinterShareWithDefaults() *MicrosoftGraphPrinterShare {
	this := MicrosoftGraphPrinterShare{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterShare) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterShare) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPrinterShare) SetId(v string) {
	o.Id = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterShare) GetCapabilities() AnyOfmicrosoftGraphPrinterCapabilities {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrinterCapabilities
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterShare) GetCapabilitiesOk() (*AnyOfmicrosoftGraphPrinterCapabilities, bool) {
	if o == nil || o.Capabilities == nil {
		return nil, false
	}
	return &o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasCapabilities() bool {
	if o != nil && o.Capabilities != nil {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given AnyOfmicrosoftGraphPrinterCapabilities and assigns it to the Capabilities field.
func (o *MicrosoftGraphPrinterShare) SetCapabilities(v AnyOfmicrosoftGraphPrinterCapabilities) {
	o.Capabilities = v
}

// GetDefaults returns the Defaults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterShare) GetDefaults() AnyOfmicrosoftGraphPrinterDefaults {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrinterDefaults
		return ret
	}
	return o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterShare) GetDefaultsOk() (*AnyOfmicrosoftGraphPrinterDefaults, bool) {
	if o == nil || o.Defaults == nil {
		return nil, false
	}
	return &o.Defaults, true
}

// HasDefaults returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasDefaults() bool {
	if o != nil && o.Defaults != nil {
		return true
	}

	return false
}

// SetDefaults gets a reference to the given AnyOfmicrosoftGraphPrinterDefaults and assigns it to the Defaults field.
func (o *MicrosoftGraphPrinterShare) SetDefaults(v AnyOfmicrosoftGraphPrinterDefaults) {
	o.Defaults = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterShare) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterShare) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *MicrosoftGraphPrinterShare) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetIsAcceptingJobs returns the IsAcceptingJobs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterShare) GetIsAcceptingJobs() bool {
	if o == nil || o.IsAcceptingJobs.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsAcceptingJobs.Get()
}

// GetIsAcceptingJobsOk returns a tuple with the IsAcceptingJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterShare) GetIsAcceptingJobsOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsAcceptingJobs.Get(), o.IsAcceptingJobs.IsSet()
}

// HasIsAcceptingJobs returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasIsAcceptingJobs() bool {
	if o != nil && o.IsAcceptingJobs.IsSet() {
		return true
	}

	return false
}

// SetIsAcceptingJobs gets a reference to the given NullableBool and assigns it to the IsAcceptingJobs field.
func (o *MicrosoftGraphPrinterShare) SetIsAcceptingJobs(v bool) {
	o.IsAcceptingJobs.Set(&v)
}
// SetIsAcceptingJobsNil sets the value for IsAcceptingJobs to be an explicit nil
func (o *MicrosoftGraphPrinterShare) SetIsAcceptingJobsNil() {
	o.IsAcceptingJobs.Set(nil)
}

// UnsetIsAcceptingJobs ensures that no value is present for IsAcceptingJobs, not even an explicit nil
func (o *MicrosoftGraphPrinterShare) UnsetIsAcceptingJobs() {
	o.IsAcceptingJobs.Unset()
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterShare) GetLocation() AnyOfmicrosoftGraphPrinterLocation {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrinterLocation
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterShare) GetLocationOk() (*AnyOfmicrosoftGraphPrinterLocation, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return &o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given AnyOfmicrosoftGraphPrinterLocation and assigns it to the Location field.
func (o *MicrosoftGraphPrinterShare) SetLocation(v AnyOfmicrosoftGraphPrinterLocation) {
	o.Location = v
}

// GetManufacturer returns the Manufacturer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterShare) GetManufacturer() string {
	if o == nil || o.Manufacturer.Get() == nil {
		var ret string
		return ret
	}
	return *o.Manufacturer.Get()
}

// GetManufacturerOk returns a tuple with the Manufacturer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterShare) GetManufacturerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Manufacturer.Get(), o.Manufacturer.IsSet()
}

// HasManufacturer returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasManufacturer() bool {
	if o != nil && o.Manufacturer.IsSet() {
		return true
	}

	return false
}

// SetManufacturer gets a reference to the given NullableString and assigns it to the Manufacturer field.
func (o *MicrosoftGraphPrinterShare) SetManufacturer(v string) {
	o.Manufacturer.Set(&v)
}
// SetManufacturerNil sets the value for Manufacturer to be an explicit nil
func (o *MicrosoftGraphPrinterShare) SetManufacturerNil() {
	o.Manufacturer.Set(nil)
}

// UnsetManufacturer ensures that no value is present for Manufacturer, not even an explicit nil
func (o *MicrosoftGraphPrinterShare) UnsetManufacturer() {
	o.Manufacturer.Unset()
}

// GetModel returns the Model field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterShare) GetModel() string {
	if o == nil || o.Model.Get() == nil {
		var ret string
		return ret
	}
	return *o.Model.Get()
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterShare) GetModelOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Model.Get(), o.Model.IsSet()
}

// HasModel returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasModel() bool {
	if o != nil && o.Model.IsSet() {
		return true
	}

	return false
}

// SetModel gets a reference to the given NullableString and assigns it to the Model field.
func (o *MicrosoftGraphPrinterShare) SetModel(v string) {
	o.Model.Set(&v)
}
// SetModelNil sets the value for Model to be an explicit nil
func (o *MicrosoftGraphPrinterShare) SetModelNil() {
	o.Model.Set(nil)
}

// UnsetModel ensures that no value is present for Model, not even an explicit nil
func (o *MicrosoftGraphPrinterShare) UnsetModel() {
	o.Model.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterShare) GetStatus() MicrosoftGraphPrinterStatus {
	if o == nil || o.Status == nil {
		var ret MicrosoftGraphPrinterStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterShare) GetStatusOk() (*MicrosoftGraphPrinterStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MicrosoftGraphPrinterStatus and assigns it to the Status field.
func (o *MicrosoftGraphPrinterShare) SetStatus(v MicrosoftGraphPrinterStatus) {
	o.Status = &v
}

// GetJobs returns the Jobs field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterShare) GetJobs() []MicrosoftGraphPrintJob {
	if o == nil || o.Jobs == nil {
		var ret []MicrosoftGraphPrintJob
		return ret
	}
	return *o.Jobs
}

// GetJobsOk returns a tuple with the Jobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterShare) GetJobsOk() (*[]MicrosoftGraphPrintJob, bool) {
	if o == nil || o.Jobs == nil {
		return nil, false
	}
	return o.Jobs, true
}

// HasJobs returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasJobs() bool {
	if o != nil && o.Jobs != nil {
		return true
	}

	return false
}

// SetJobs gets a reference to the given []MicrosoftGraphPrintJob and assigns it to the Jobs field.
func (o *MicrosoftGraphPrinterShare) SetJobs(v []MicrosoftGraphPrintJob) {
	o.Jobs = &v
}

// GetAllowAllUsers returns the AllowAllUsers field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterShare) GetAllowAllUsers() bool {
	if o == nil || o.AllowAllUsers == nil {
		var ret bool
		return ret
	}
	return *o.AllowAllUsers
}

// GetAllowAllUsersOk returns a tuple with the AllowAllUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterShare) GetAllowAllUsersOk() (*bool, bool) {
	if o == nil || o.AllowAllUsers == nil {
		return nil, false
	}
	return o.AllowAllUsers, true
}

// HasAllowAllUsers returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasAllowAllUsers() bool {
	if o != nil && o.AllowAllUsers != nil {
		return true
	}

	return false
}

// SetAllowAllUsers gets a reference to the given bool and assigns it to the AllowAllUsers field.
func (o *MicrosoftGraphPrinterShare) SetAllowAllUsers(v bool) {
	o.AllowAllUsers = &v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterShare) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterShare) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil || o.CreatedDateTime == nil {
		return nil, false
	}
	return o.CreatedDateTime, true
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime != nil {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given time.Time and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphPrinterShare) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime = &v
}

// GetAllowedGroups returns the AllowedGroups field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterShare) GetAllowedGroups() []MicrosoftGraphGroup {
	if o == nil || o.AllowedGroups == nil {
		var ret []MicrosoftGraphGroup
		return ret
	}
	return *o.AllowedGroups
}

// GetAllowedGroupsOk returns a tuple with the AllowedGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterShare) GetAllowedGroupsOk() (*[]MicrosoftGraphGroup, bool) {
	if o == nil || o.AllowedGroups == nil {
		return nil, false
	}
	return o.AllowedGroups, true
}

// HasAllowedGroups returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasAllowedGroups() bool {
	if o != nil && o.AllowedGroups != nil {
		return true
	}

	return false
}

// SetAllowedGroups gets a reference to the given []MicrosoftGraphGroup and assigns it to the AllowedGroups field.
func (o *MicrosoftGraphPrinterShare) SetAllowedGroups(v []MicrosoftGraphGroup) {
	o.AllowedGroups = &v
}

// GetAllowedUsers returns the AllowedUsers field value if set, zero value otherwise.
func (o *MicrosoftGraphPrinterShare) GetAllowedUsers() []MicrosoftGraphUser {
	if o == nil || o.AllowedUsers == nil {
		var ret []MicrosoftGraphUser
		return ret
	}
	return *o.AllowedUsers
}

// GetAllowedUsersOk returns a tuple with the AllowedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrinterShare) GetAllowedUsersOk() (*[]MicrosoftGraphUser, bool) {
	if o == nil || o.AllowedUsers == nil {
		return nil, false
	}
	return o.AllowedUsers, true
}

// HasAllowedUsers returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasAllowedUsers() bool {
	if o != nil && o.AllowedUsers != nil {
		return true
	}

	return false
}

// SetAllowedUsers gets a reference to the given []MicrosoftGraphUser and assigns it to the AllowedUsers field.
func (o *MicrosoftGraphPrinterShare) SetAllowedUsers(v []MicrosoftGraphUser) {
	o.AllowedUsers = &v
}

// GetPrinter returns the Printer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphPrinterShare) GetPrinter() AnyOfmicrosoftGraphPrinter {
	if o == nil  {
		var ret AnyOfmicrosoftGraphPrinter
		return ret
	}
	return o.Printer
}

// GetPrinterOk returns a tuple with the Printer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphPrinterShare) GetPrinterOk() (*AnyOfmicrosoftGraphPrinter, bool) {
	if o == nil || o.Printer == nil {
		return nil, false
	}
	return &o.Printer, true
}

// HasPrinter returns a boolean if a field has been set.
func (o *MicrosoftGraphPrinterShare) HasPrinter() bool {
	if o != nil && o.Printer != nil {
		return true
	}

	return false
}

// SetPrinter gets a reference to the given AnyOfmicrosoftGraphPrinter and assigns it to the Printer field.
func (o *MicrosoftGraphPrinterShare) SetPrinter(v AnyOfmicrosoftGraphPrinter) {
	o.Printer = v
}

func (o MicrosoftGraphPrinterShare) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Capabilities != nil {
		toSerialize["capabilities"] = o.Capabilities
	}
	if o.Defaults != nil {
		toSerialize["defaults"] = o.Defaults
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.IsAcceptingJobs.IsSet() {
		toSerialize["isAcceptingJobs"] = o.IsAcceptingJobs.Get()
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Manufacturer.IsSet() {
		toSerialize["manufacturer"] = o.Manufacturer.Get()
	}
	if o.Model.IsSet() {
		toSerialize["model"] = o.Model.Get()
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Jobs != nil {
		toSerialize["jobs"] = o.Jobs
	}
	if o.AllowAllUsers != nil {
		toSerialize["allowAllUsers"] = o.AllowAllUsers
	}
	if o.CreatedDateTime != nil {
		toSerialize["createdDateTime"] = o.CreatedDateTime
	}
	if o.AllowedGroups != nil {
		toSerialize["allowedGroups"] = o.AllowedGroups
	}
	if o.AllowedUsers != nil {
		toSerialize["allowedUsers"] = o.AllowedUsers
	}
	if o.Printer != nil {
		toSerialize["printer"] = o.Printer
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrinterShare struct {
	value *MicrosoftGraphPrinterShare
	isSet bool
}

func (v NullableMicrosoftGraphPrinterShare) Get() *MicrosoftGraphPrinterShare {
	return v.value
}

func (v *NullableMicrosoftGraphPrinterShare) Set(val *MicrosoftGraphPrinterShare) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrinterShare) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrinterShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrinterShare(val *MicrosoftGraphPrinterShare) *NullableMicrosoftGraphPrinterShare {
	return &NullableMicrosoftGraphPrinterShare{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrinterShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrinterShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


