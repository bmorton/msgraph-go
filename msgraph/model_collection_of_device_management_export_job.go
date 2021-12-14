/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// CollectionOfDeviceManagementExportJob struct for CollectionOfDeviceManagementExportJob
type CollectionOfDeviceManagementExportJob struct {
	Value *[]MicrosoftGraphDeviceManagementExportJob `json:"value,omitempty"`
	OdataNextLink *string `json:"@odata.nextLink,omitempty"`
}

// NewCollectionOfDeviceManagementExportJob instantiates a new CollectionOfDeviceManagementExportJob object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionOfDeviceManagementExportJob() *CollectionOfDeviceManagementExportJob {
	this := CollectionOfDeviceManagementExportJob{}
	return &this
}

// NewCollectionOfDeviceManagementExportJobWithDefaults instantiates a new CollectionOfDeviceManagementExportJob object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionOfDeviceManagementExportJobWithDefaults() *CollectionOfDeviceManagementExportJob {
	this := CollectionOfDeviceManagementExportJob{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionOfDeviceManagementExportJob) GetValue() []MicrosoftGraphDeviceManagementExportJob {
	if o == nil || o.Value == nil {
		var ret []MicrosoftGraphDeviceManagementExportJob
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceManagementExportJob) GetValueOk() (*[]MicrosoftGraphDeviceManagementExportJob, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionOfDeviceManagementExportJob) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []MicrosoftGraphDeviceManagementExportJob and assigns it to the Value field.
func (o *CollectionOfDeviceManagementExportJob) SetValue(v []MicrosoftGraphDeviceManagementExportJob) {
	o.Value = &v
}

// GetOdataNextLink returns the OdataNextLink field value if set, zero value otherwise.
func (o *CollectionOfDeviceManagementExportJob) GetOdataNextLink() string {
	if o == nil || o.OdataNextLink == nil {
		var ret string
		return ret
	}
	return *o.OdataNextLink
}

// GetOdataNextLinkOk returns a tuple with the OdataNextLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionOfDeviceManagementExportJob) GetOdataNextLinkOk() (*string, bool) {
	if o == nil || o.OdataNextLink == nil {
		return nil, false
	}
	return o.OdataNextLink, true
}

// HasOdataNextLink returns a boolean if a field has been set.
func (o *CollectionOfDeviceManagementExportJob) HasOdataNextLink() bool {
	if o != nil && o.OdataNextLink != nil {
		return true
	}

	return false
}

// SetOdataNextLink gets a reference to the given string and assigns it to the OdataNextLink field.
func (o *CollectionOfDeviceManagementExportJob) SetOdataNextLink(v string) {
	o.OdataNextLink = &v
}

func (o CollectionOfDeviceManagementExportJob) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.OdataNextLink != nil {
		toSerialize["@odata.nextLink"] = o.OdataNextLink
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionOfDeviceManagementExportJob struct {
	value *CollectionOfDeviceManagementExportJob
	isSet bool
}

func (v NullableCollectionOfDeviceManagementExportJob) Get() *CollectionOfDeviceManagementExportJob {
	return v.value
}

func (v *NullableCollectionOfDeviceManagementExportJob) Set(val *CollectionOfDeviceManagementExportJob) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionOfDeviceManagementExportJob) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionOfDeviceManagementExportJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionOfDeviceManagementExportJob(val *CollectionOfDeviceManagementExportJob) *NullableCollectionOfDeviceManagementExportJob {
	return &NullableCollectionOfDeviceManagementExportJob{value: val, isSet: true}
}

func (v NullableCollectionOfDeviceManagementExportJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionOfDeviceManagementExportJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


