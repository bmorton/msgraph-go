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

// MicrosoftGraphReportRoot struct for MicrosoftGraphReportRoot
type MicrosoftGraphReportRoot struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	DailyPrintUsageByPrinter *[]MicrosoftGraphPrintUsageByPrinter `json:"dailyPrintUsageByPrinter,omitempty"`
	DailyPrintUsageByUser *[]MicrosoftGraphPrintUsageByUser `json:"dailyPrintUsageByUser,omitempty"`
	MonthlyPrintUsageByPrinter *[]MicrosoftGraphPrintUsageByPrinter `json:"monthlyPrintUsageByPrinter,omitempty"`
	MonthlyPrintUsageByUser *[]MicrosoftGraphPrintUsageByUser `json:"monthlyPrintUsageByUser,omitempty"`
}

// NewMicrosoftGraphReportRoot instantiates a new MicrosoftGraphReportRoot object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphReportRoot() *MicrosoftGraphReportRoot {
	this := MicrosoftGraphReportRoot{}
	return &this
}

// NewMicrosoftGraphReportRootWithDefaults instantiates a new MicrosoftGraphReportRoot object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphReportRootWithDefaults() *MicrosoftGraphReportRoot {
	this := MicrosoftGraphReportRoot{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphReportRoot) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphReportRoot) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphReportRoot) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphReportRoot) SetId(v string) {
	o.Id = &v
}

// GetDailyPrintUsageByPrinter returns the DailyPrintUsageByPrinter field value if set, zero value otherwise.
func (o *MicrosoftGraphReportRoot) GetDailyPrintUsageByPrinter() []MicrosoftGraphPrintUsageByPrinter {
	if o == nil || o.DailyPrintUsageByPrinter == nil {
		var ret []MicrosoftGraphPrintUsageByPrinter
		return ret
	}
	return *o.DailyPrintUsageByPrinter
}

// GetDailyPrintUsageByPrinterOk returns a tuple with the DailyPrintUsageByPrinter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphReportRoot) GetDailyPrintUsageByPrinterOk() (*[]MicrosoftGraphPrintUsageByPrinter, bool) {
	if o == nil || o.DailyPrintUsageByPrinter == nil {
		return nil, false
	}
	return o.DailyPrintUsageByPrinter, true
}

// HasDailyPrintUsageByPrinter returns a boolean if a field has been set.
func (o *MicrosoftGraphReportRoot) HasDailyPrintUsageByPrinter() bool {
	if o != nil && o.DailyPrintUsageByPrinter != nil {
		return true
	}

	return false
}

// SetDailyPrintUsageByPrinter gets a reference to the given []MicrosoftGraphPrintUsageByPrinter and assigns it to the DailyPrintUsageByPrinter field.
func (o *MicrosoftGraphReportRoot) SetDailyPrintUsageByPrinter(v []MicrosoftGraphPrintUsageByPrinter) {
	o.DailyPrintUsageByPrinter = &v
}

// GetDailyPrintUsageByUser returns the DailyPrintUsageByUser field value if set, zero value otherwise.
func (o *MicrosoftGraphReportRoot) GetDailyPrintUsageByUser() []MicrosoftGraphPrintUsageByUser {
	if o == nil || o.DailyPrintUsageByUser == nil {
		var ret []MicrosoftGraphPrintUsageByUser
		return ret
	}
	return *o.DailyPrintUsageByUser
}

// GetDailyPrintUsageByUserOk returns a tuple with the DailyPrintUsageByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphReportRoot) GetDailyPrintUsageByUserOk() (*[]MicrosoftGraphPrintUsageByUser, bool) {
	if o == nil || o.DailyPrintUsageByUser == nil {
		return nil, false
	}
	return o.DailyPrintUsageByUser, true
}

// HasDailyPrintUsageByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphReportRoot) HasDailyPrintUsageByUser() bool {
	if o != nil && o.DailyPrintUsageByUser != nil {
		return true
	}

	return false
}

// SetDailyPrintUsageByUser gets a reference to the given []MicrosoftGraphPrintUsageByUser and assigns it to the DailyPrintUsageByUser field.
func (o *MicrosoftGraphReportRoot) SetDailyPrintUsageByUser(v []MicrosoftGraphPrintUsageByUser) {
	o.DailyPrintUsageByUser = &v
}

// GetMonthlyPrintUsageByPrinter returns the MonthlyPrintUsageByPrinter field value if set, zero value otherwise.
func (o *MicrosoftGraphReportRoot) GetMonthlyPrintUsageByPrinter() []MicrosoftGraphPrintUsageByPrinter {
	if o == nil || o.MonthlyPrintUsageByPrinter == nil {
		var ret []MicrosoftGraphPrintUsageByPrinter
		return ret
	}
	return *o.MonthlyPrintUsageByPrinter
}

// GetMonthlyPrintUsageByPrinterOk returns a tuple with the MonthlyPrintUsageByPrinter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphReportRoot) GetMonthlyPrintUsageByPrinterOk() (*[]MicrosoftGraphPrintUsageByPrinter, bool) {
	if o == nil || o.MonthlyPrintUsageByPrinter == nil {
		return nil, false
	}
	return o.MonthlyPrintUsageByPrinter, true
}

// HasMonthlyPrintUsageByPrinter returns a boolean if a field has been set.
func (o *MicrosoftGraphReportRoot) HasMonthlyPrintUsageByPrinter() bool {
	if o != nil && o.MonthlyPrintUsageByPrinter != nil {
		return true
	}

	return false
}

// SetMonthlyPrintUsageByPrinter gets a reference to the given []MicrosoftGraphPrintUsageByPrinter and assigns it to the MonthlyPrintUsageByPrinter field.
func (o *MicrosoftGraphReportRoot) SetMonthlyPrintUsageByPrinter(v []MicrosoftGraphPrintUsageByPrinter) {
	o.MonthlyPrintUsageByPrinter = &v
}

// GetMonthlyPrintUsageByUser returns the MonthlyPrintUsageByUser field value if set, zero value otherwise.
func (o *MicrosoftGraphReportRoot) GetMonthlyPrintUsageByUser() []MicrosoftGraphPrintUsageByUser {
	if o == nil || o.MonthlyPrintUsageByUser == nil {
		var ret []MicrosoftGraphPrintUsageByUser
		return ret
	}
	return *o.MonthlyPrintUsageByUser
}

// GetMonthlyPrintUsageByUserOk returns a tuple with the MonthlyPrintUsageByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphReportRoot) GetMonthlyPrintUsageByUserOk() (*[]MicrosoftGraphPrintUsageByUser, bool) {
	if o == nil || o.MonthlyPrintUsageByUser == nil {
		return nil, false
	}
	return o.MonthlyPrintUsageByUser, true
}

// HasMonthlyPrintUsageByUser returns a boolean if a field has been set.
func (o *MicrosoftGraphReportRoot) HasMonthlyPrintUsageByUser() bool {
	if o != nil && o.MonthlyPrintUsageByUser != nil {
		return true
	}

	return false
}

// SetMonthlyPrintUsageByUser gets a reference to the given []MicrosoftGraphPrintUsageByUser and assigns it to the MonthlyPrintUsageByUser field.
func (o *MicrosoftGraphReportRoot) SetMonthlyPrintUsageByUser(v []MicrosoftGraphPrintUsageByUser) {
	o.MonthlyPrintUsageByUser = &v
}

func (o MicrosoftGraphReportRoot) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DailyPrintUsageByPrinter != nil {
		toSerialize["dailyPrintUsageByPrinter"] = o.DailyPrintUsageByPrinter
	}
	if o.DailyPrintUsageByUser != nil {
		toSerialize["dailyPrintUsageByUser"] = o.DailyPrintUsageByUser
	}
	if o.MonthlyPrintUsageByPrinter != nil {
		toSerialize["monthlyPrintUsageByPrinter"] = o.MonthlyPrintUsageByPrinter
	}
	if o.MonthlyPrintUsageByUser != nil {
		toSerialize["monthlyPrintUsageByUser"] = o.MonthlyPrintUsageByUser
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphReportRoot struct {
	value *MicrosoftGraphReportRoot
	isSet bool
}

func (v NullableMicrosoftGraphReportRoot) Get() *MicrosoftGraphReportRoot {
	return v.value
}

func (v *NullableMicrosoftGraphReportRoot) Set(val *MicrosoftGraphReportRoot) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphReportRoot) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphReportRoot) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphReportRoot(val *MicrosoftGraphReportRoot) *NullableMicrosoftGraphReportRoot {
	return &NullableMicrosoftGraphReportRoot{value: val, isSet: true}
}

func (v NullableMicrosoftGraphReportRoot) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphReportRoot) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


