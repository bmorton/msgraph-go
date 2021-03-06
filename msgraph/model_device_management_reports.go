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

// DeviceManagementReports Singleton entity that acts as a container for all reports functionality.
type DeviceManagementReports struct {
	// Entity representing a job to export a report
	ExportJobs *[]MicrosoftGraphDeviceManagementExportJob `json:"exportJobs,omitempty"`
}

// NewDeviceManagementReports instantiates a new DeviceManagementReports object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceManagementReports() *DeviceManagementReports {
	this := DeviceManagementReports{}
	return &this
}

// NewDeviceManagementReportsWithDefaults instantiates a new DeviceManagementReports object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceManagementReportsWithDefaults() *DeviceManagementReports {
	this := DeviceManagementReports{}
	return &this
}

// GetExportJobs returns the ExportJobs field value if set, zero value otherwise.
func (o *DeviceManagementReports) GetExportJobs() []MicrosoftGraphDeviceManagementExportJob {
	if o == nil || o.ExportJobs == nil {
		var ret []MicrosoftGraphDeviceManagementExportJob
		return ret
	}
	return *o.ExportJobs
}

// GetExportJobsOk returns a tuple with the ExportJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceManagementReports) GetExportJobsOk() (*[]MicrosoftGraphDeviceManagementExportJob, bool) {
	if o == nil || o.ExportJobs == nil {
		return nil, false
	}
	return o.ExportJobs, true
}

// HasExportJobs returns a boolean if a field has been set.
func (o *DeviceManagementReports) HasExportJobs() bool {
	if o != nil && o.ExportJobs != nil {
		return true
	}

	return false
}

// SetExportJobs gets a reference to the given []MicrosoftGraphDeviceManagementExportJob and assigns it to the ExportJobs field.
func (o *DeviceManagementReports) SetExportJobs(v []MicrosoftGraphDeviceManagementExportJob) {
	o.ExportJobs = &v
}

func (o DeviceManagementReports) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExportJobs != nil {
		toSerialize["exportJobs"] = o.ExportJobs
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceManagementReports struct {
	value *DeviceManagementReports
	isSet bool
}

func (v NullableDeviceManagementReports) Get() *DeviceManagementReports {
	return v.value
}

func (v *NullableDeviceManagementReports) Set(val *DeviceManagementReports) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceManagementReports) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceManagementReports) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceManagementReports(val *DeviceManagementReports) *NullableDeviceManagementReports {
	return &NullableDeviceManagementReports{value: val, isSet: true}
}

func (v NullableDeviceManagementReports) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceManagementReports) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


