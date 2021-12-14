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

// PrintUsage struct for PrintUsage
type PrintUsage struct {
	CompletedBlackAndWhiteJobCount *int64 `json:"completedBlackAndWhiteJobCount,omitempty"`
	CompletedColorJobCount *int64 `json:"completedColorJobCount,omitempty"`
	IncompleteJobCount *int64 `json:"incompleteJobCount,omitempty"`
	UsageDate *string `json:"usageDate,omitempty"`
}

// NewPrintUsage instantiates a new PrintUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrintUsage() *PrintUsage {
	this := PrintUsage{}
	return &this
}

// NewPrintUsageWithDefaults instantiates a new PrintUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrintUsageWithDefaults() *PrintUsage {
	this := PrintUsage{}
	return &this
}

// GetCompletedBlackAndWhiteJobCount returns the CompletedBlackAndWhiteJobCount field value if set, zero value otherwise.
func (o *PrintUsage) GetCompletedBlackAndWhiteJobCount() int64 {
	if o == nil || o.CompletedBlackAndWhiteJobCount == nil {
		var ret int64
		return ret
	}
	return *o.CompletedBlackAndWhiteJobCount
}

// GetCompletedBlackAndWhiteJobCountOk returns a tuple with the CompletedBlackAndWhiteJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintUsage) GetCompletedBlackAndWhiteJobCountOk() (*int64, bool) {
	if o == nil || o.CompletedBlackAndWhiteJobCount == nil {
		return nil, false
	}
	return o.CompletedBlackAndWhiteJobCount, true
}

// HasCompletedBlackAndWhiteJobCount returns a boolean if a field has been set.
func (o *PrintUsage) HasCompletedBlackAndWhiteJobCount() bool {
	if o != nil && o.CompletedBlackAndWhiteJobCount != nil {
		return true
	}

	return false
}

// SetCompletedBlackAndWhiteJobCount gets a reference to the given int64 and assigns it to the CompletedBlackAndWhiteJobCount field.
func (o *PrintUsage) SetCompletedBlackAndWhiteJobCount(v int64) {
	o.CompletedBlackAndWhiteJobCount = &v
}

// GetCompletedColorJobCount returns the CompletedColorJobCount field value if set, zero value otherwise.
func (o *PrintUsage) GetCompletedColorJobCount() int64 {
	if o == nil || o.CompletedColorJobCount == nil {
		var ret int64
		return ret
	}
	return *o.CompletedColorJobCount
}

// GetCompletedColorJobCountOk returns a tuple with the CompletedColorJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintUsage) GetCompletedColorJobCountOk() (*int64, bool) {
	if o == nil || o.CompletedColorJobCount == nil {
		return nil, false
	}
	return o.CompletedColorJobCount, true
}

// HasCompletedColorJobCount returns a boolean if a field has been set.
func (o *PrintUsage) HasCompletedColorJobCount() bool {
	if o != nil && o.CompletedColorJobCount != nil {
		return true
	}

	return false
}

// SetCompletedColorJobCount gets a reference to the given int64 and assigns it to the CompletedColorJobCount field.
func (o *PrintUsage) SetCompletedColorJobCount(v int64) {
	o.CompletedColorJobCount = &v
}

// GetIncompleteJobCount returns the IncompleteJobCount field value if set, zero value otherwise.
func (o *PrintUsage) GetIncompleteJobCount() int64 {
	if o == nil || o.IncompleteJobCount == nil {
		var ret int64
		return ret
	}
	return *o.IncompleteJobCount
}

// GetIncompleteJobCountOk returns a tuple with the IncompleteJobCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintUsage) GetIncompleteJobCountOk() (*int64, bool) {
	if o == nil || o.IncompleteJobCount == nil {
		return nil, false
	}
	return o.IncompleteJobCount, true
}

// HasIncompleteJobCount returns a boolean if a field has been set.
func (o *PrintUsage) HasIncompleteJobCount() bool {
	if o != nil && o.IncompleteJobCount != nil {
		return true
	}

	return false
}

// SetIncompleteJobCount gets a reference to the given int64 and assigns it to the IncompleteJobCount field.
func (o *PrintUsage) SetIncompleteJobCount(v int64) {
	o.IncompleteJobCount = &v
}

// GetUsageDate returns the UsageDate field value if set, zero value otherwise.
func (o *PrintUsage) GetUsageDate() string {
	if o == nil || o.UsageDate == nil {
		var ret string
		return ret
	}
	return *o.UsageDate
}

// GetUsageDateOk returns a tuple with the UsageDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrintUsage) GetUsageDateOk() (*string, bool) {
	if o == nil || o.UsageDate == nil {
		return nil, false
	}
	return o.UsageDate, true
}

// HasUsageDate returns a boolean if a field has been set.
func (o *PrintUsage) HasUsageDate() bool {
	if o != nil && o.UsageDate != nil {
		return true
	}

	return false
}

// SetUsageDate gets a reference to the given string and assigns it to the UsageDate field.
func (o *PrintUsage) SetUsageDate(v string) {
	o.UsageDate = &v
}

func (o PrintUsage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
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
	return json.Marshal(toSerialize)
}

type NullablePrintUsage struct {
	value *PrintUsage
	isSet bool
}

func (v NullablePrintUsage) Get() *PrintUsage {
	return v.value
}

func (v *NullablePrintUsage) Set(val *PrintUsage) {
	v.value = val
	v.isSet = true
}

func (v NullablePrintUsage) IsSet() bool {
	return v.isSet
}

func (v *NullablePrintUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrintUsage(val *PrintUsage) *NullablePrintUsage {
	return &NullablePrintUsage{value: val, isSet: true}
}

func (v NullablePrintUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrintUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


