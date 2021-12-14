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

// MicrosoftGraphPrintUsageByPrinter struct for MicrosoftGraphPrintUsageByPrinter
type MicrosoftGraphPrintUsageByPrinter struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	CompletedBlackAndWhiteJobCount *int64 `json:"completedBlackAndWhiteJobCount,omitempty"`
	CompletedColorJobCount *int64 `json:"completedColorJobCount,omitempty"`
	IncompleteJobCount *int64 `json:"incompleteJobCount,omitempty"`
	UsageDate *string `json:"usageDate,omitempty"`
	PrinterId *string `json:"printerId,omitempty"`
}

// NewMicrosoftGraphPrintUsageByPrinter instantiates a new MicrosoftGraphPrintUsageByPrinter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphPrintUsageByPrinter() *MicrosoftGraphPrintUsageByPrinter {
	this := MicrosoftGraphPrintUsageByPrinter{}
	return &this
}

// NewMicrosoftGraphPrintUsageByPrinterWithDefaults instantiates a new MicrosoftGraphPrintUsageByPrinter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphPrintUsageByPrinterWithDefaults() *MicrosoftGraphPrintUsageByPrinter {
	this := MicrosoftGraphPrintUsageByPrinter{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintUsageByPrinter) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphPrintUsageByPrinter) SetId(v string) {
	o.Id = &v
}

// GetCompletedBlackAndWhiteJobCount returns the CompletedBlackAndWhiteJobCount field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintUsageByPrinter) GetCompletedBlackAndWhiteJobCount() int64 {
	if o == nil || o.CompletedBlackAndWhiteJobCount == nil {
		var ret int64
		return ret
	}
	return *o.CompletedBlackAndWhiteJobCount
}

// GetCompletedBlackAndWhiteJobCountOk returns a tuple with the CompletedBlackAndWhiteJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) GetCompletedBlackAndWhiteJobCountOk() (*int64, bool) {
	if o == nil || o.CompletedBlackAndWhiteJobCount == nil {
		return nil, false
	}
	return o.CompletedBlackAndWhiteJobCount, true
}

// HasCompletedBlackAndWhiteJobCount returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) HasCompletedBlackAndWhiteJobCount() bool {
	if o != nil && o.CompletedBlackAndWhiteJobCount != nil {
		return true
	}

	return false
}

// SetCompletedBlackAndWhiteJobCount gets a reference to the given int64 and assigns it to the CompletedBlackAndWhiteJobCount field.
func (o *MicrosoftGraphPrintUsageByPrinter) SetCompletedBlackAndWhiteJobCount(v int64) {
	o.CompletedBlackAndWhiteJobCount = &v
}

// GetCompletedColorJobCount returns the CompletedColorJobCount field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintUsageByPrinter) GetCompletedColorJobCount() int64 {
	if o == nil || o.CompletedColorJobCount == nil {
		var ret int64
		return ret
	}
	return *o.CompletedColorJobCount
}

// GetCompletedColorJobCountOk returns a tuple with the CompletedColorJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) GetCompletedColorJobCountOk() (*int64, bool) {
	if o == nil || o.CompletedColorJobCount == nil {
		return nil, false
	}
	return o.CompletedColorJobCount, true
}

// HasCompletedColorJobCount returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) HasCompletedColorJobCount() bool {
	if o != nil && o.CompletedColorJobCount != nil {
		return true
	}

	return false
}

// SetCompletedColorJobCount gets a reference to the given int64 and assigns it to the CompletedColorJobCount field.
func (o *MicrosoftGraphPrintUsageByPrinter) SetCompletedColorJobCount(v int64) {
	o.CompletedColorJobCount = &v
}

// GetIncompleteJobCount returns the IncompleteJobCount field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintUsageByPrinter) GetIncompleteJobCount() int64 {
	if o == nil || o.IncompleteJobCount == nil {
		var ret int64
		return ret
	}
	return *o.IncompleteJobCount
}

// GetIncompleteJobCountOk returns a tuple with the IncompleteJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) GetIncompleteJobCountOk() (*int64, bool) {
	if o == nil || o.IncompleteJobCount == nil {
		return nil, false
	}
	return o.IncompleteJobCount, true
}

// HasIncompleteJobCount returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) HasIncompleteJobCount() bool {
	if o != nil && o.IncompleteJobCount != nil {
		return true
	}

	return false
}

// SetIncompleteJobCount gets a reference to the given int64 and assigns it to the IncompleteJobCount field.
func (o *MicrosoftGraphPrintUsageByPrinter) SetIncompleteJobCount(v int64) {
	o.IncompleteJobCount = &v
}

// GetUsageDate returns the UsageDate field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintUsageByPrinter) GetUsageDate() string {
	if o == nil || o.UsageDate == nil {
		var ret string
		return ret
	}
	return *o.UsageDate
}

// GetUsageDateOk returns a tuple with the UsageDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) GetUsageDateOk() (*string, bool) {
	if o == nil || o.UsageDate == nil {
		return nil, false
	}
	return o.UsageDate, true
}

// HasUsageDate returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) HasUsageDate() bool {
	if o != nil && o.UsageDate != nil {
		return true
	}

	return false
}

// SetUsageDate gets a reference to the given string and assigns it to the UsageDate field.
func (o *MicrosoftGraphPrintUsageByPrinter) SetUsageDate(v string) {
	o.UsageDate = &v
}

// GetPrinterId returns the PrinterId field value if set, zero value otherwise.
func (o *MicrosoftGraphPrintUsageByPrinter) GetPrinterId() string {
	if o == nil || o.PrinterId == nil {
		var ret string
		return ret
	}
	return *o.PrinterId
}

// GetPrinterIdOk returns a tuple with the PrinterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) GetPrinterIdOk() (*string, bool) {
	if o == nil || o.PrinterId == nil {
		return nil, false
	}
	return o.PrinterId, true
}

// HasPrinterId returns a boolean if a field has been set.
func (o *MicrosoftGraphPrintUsageByPrinter) HasPrinterId() bool {
	if o != nil && o.PrinterId != nil {
		return true
	}

	return false
}

// SetPrinterId gets a reference to the given string and assigns it to the PrinterId field.
func (o *MicrosoftGraphPrintUsageByPrinter) SetPrinterId(v string) {
	o.PrinterId = &v
}

func (o MicrosoftGraphPrintUsageByPrinter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CompletedBlackAndWhiteJobCount != nil {
		toSerialize["completedBlackAndWhiteJobCount"] = o.CompletedBlackAndWhiteJobCount
	}
	if o.CompletedColorJobCount != nil {
		toSerialize["completedColorJobCount"] = o.CompletedColorJobCount
	}
	if o.IncompleteJobCount != nil {
		toSerialize["incompleteJobCount"] = o.IncompleteJobCount
	}
	if o.UsageDate != nil {
		toSerialize["usageDate"] = o.UsageDate
	}
	if o.PrinterId != nil {
		toSerialize["printerId"] = o.PrinterId
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphPrintUsageByPrinter struct {
	value *MicrosoftGraphPrintUsageByPrinter
	isSet bool
}

func (v NullableMicrosoftGraphPrintUsageByPrinter) Get() *MicrosoftGraphPrintUsageByPrinter {
	return v.value
}

func (v *NullableMicrosoftGraphPrintUsageByPrinter) Set(val *MicrosoftGraphPrintUsageByPrinter) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphPrintUsageByPrinter) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphPrintUsageByPrinter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphPrintUsageByPrinter(val *MicrosoftGraphPrintUsageByPrinter) *NullableMicrosoftGraphPrintUsageByPrinter {
	return &NullableMicrosoftGraphPrintUsageByPrinter{value: val, isSet: true}
}

func (v NullableMicrosoftGraphPrintUsageByPrinter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphPrintUsageByPrinter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

