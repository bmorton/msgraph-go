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

// MicrosoftGraphBucketAggregationDefinition struct for MicrosoftGraphBucketAggregationDefinition
type MicrosoftGraphBucketAggregationDefinition struct {
	// True to specify the sort order as descending. The default is false, with the sort order as ascending. Optional.
	IsDescending NullableBool `json:"isDescending,omitempty"`
	// The minimum number of items that should be present in the aggregation to be returned in a bucket. Optional.
	MinimumCount NullableInt32 `json:"minimumCount,omitempty"`
	// A filter to define a matching criteria. The key should start with the specified prefix to be returned in the response. Optional.
	PrefixFilter NullableString `json:"prefixFilter,omitempty"`
	// Specifies the manual ranges to compute the aggregations. This is only valid for non-string refiners of date or numeric type. Optional.
	Ranges *[]*AnyOfmicrosoftGraphBucketAggregationRange `json:"ranges,omitempty"`
	// The possible values are count to sort by the number of matches in the aggregation, keyAsStringto sort alphabeticaly based on the key in the aggregation, keyAsNumber for numerical sorting based on the key in the aggregation. Required.
	SortBy AnyOfmicrosoftGraphBucketAggregationSortProperty `json:"sortBy,omitempty"`
}

// NewMicrosoftGraphBucketAggregationDefinition instantiates a new MicrosoftGraphBucketAggregationDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphBucketAggregationDefinition() *MicrosoftGraphBucketAggregationDefinition {
	this := MicrosoftGraphBucketAggregationDefinition{}
	return &this
}

// NewMicrosoftGraphBucketAggregationDefinitionWithDefaults instantiates a new MicrosoftGraphBucketAggregationDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphBucketAggregationDefinitionWithDefaults() *MicrosoftGraphBucketAggregationDefinition {
	this := MicrosoftGraphBucketAggregationDefinition{}
	return &this
}

// GetIsDescending returns the IsDescending field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBucketAggregationDefinition) GetIsDescending() bool {
	if o == nil || o.IsDescending.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsDescending.Get()
}

// GetIsDescendingOk returns a tuple with the IsDescending field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBucketAggregationDefinition) GetIsDescendingOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsDescending.Get(), o.IsDescending.IsSet()
}

// HasIsDescending returns a boolean if a field has been set.
func (o *MicrosoftGraphBucketAggregationDefinition) HasIsDescending() bool {
	if o != nil && o.IsDescending.IsSet() {
		return true
	}

	return false
}

// SetIsDescending gets a reference to the given NullableBool and assigns it to the IsDescending field.
func (o *MicrosoftGraphBucketAggregationDefinition) SetIsDescending(v bool) {
	o.IsDescending.Set(&v)
}
// SetIsDescendingNil sets the value for IsDescending to be an explicit nil
func (o *MicrosoftGraphBucketAggregationDefinition) SetIsDescendingNil() {
	o.IsDescending.Set(nil)
}

// UnsetIsDescending ensures that no value is present for IsDescending, not even an explicit nil
func (o *MicrosoftGraphBucketAggregationDefinition) UnsetIsDescending() {
	o.IsDescending.Unset()
}

// GetMinimumCount returns the MinimumCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBucketAggregationDefinition) GetMinimumCount() int32 {
	if o == nil || o.MinimumCount.Get() == nil {
		var ret int32
		return ret
	}
	return *o.MinimumCount.Get()
}

// GetMinimumCountOk returns a tuple with the MinimumCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBucketAggregationDefinition) GetMinimumCountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.MinimumCount.Get(), o.MinimumCount.IsSet()
}

// HasMinimumCount returns a boolean if a field has been set.
func (o *MicrosoftGraphBucketAggregationDefinition) HasMinimumCount() bool {
	if o != nil && o.MinimumCount.IsSet() {
		return true
	}

	return false
}

// SetMinimumCount gets a reference to the given NullableInt32 and assigns it to the MinimumCount field.
func (o *MicrosoftGraphBucketAggregationDefinition) SetMinimumCount(v int32) {
	o.MinimumCount.Set(&v)
}
// SetMinimumCountNil sets the value for MinimumCount to be an explicit nil
func (o *MicrosoftGraphBucketAggregationDefinition) SetMinimumCountNil() {
	o.MinimumCount.Set(nil)
}

// UnsetMinimumCount ensures that no value is present for MinimumCount, not even an explicit nil
func (o *MicrosoftGraphBucketAggregationDefinition) UnsetMinimumCount() {
	o.MinimumCount.Unset()
}

// GetPrefixFilter returns the PrefixFilter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBucketAggregationDefinition) GetPrefixFilter() string {
	if o == nil || o.PrefixFilter.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrefixFilter.Get()
}

// GetPrefixFilterOk returns a tuple with the PrefixFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBucketAggregationDefinition) GetPrefixFilterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrefixFilter.Get(), o.PrefixFilter.IsSet()
}

// HasPrefixFilter returns a boolean if a field has been set.
func (o *MicrosoftGraphBucketAggregationDefinition) HasPrefixFilter() bool {
	if o != nil && o.PrefixFilter.IsSet() {
		return true
	}

	return false
}

// SetPrefixFilter gets a reference to the given NullableString and assigns it to the PrefixFilter field.
func (o *MicrosoftGraphBucketAggregationDefinition) SetPrefixFilter(v string) {
	o.PrefixFilter.Set(&v)
}
// SetPrefixFilterNil sets the value for PrefixFilter to be an explicit nil
func (o *MicrosoftGraphBucketAggregationDefinition) SetPrefixFilterNil() {
	o.PrefixFilter.Set(nil)
}

// UnsetPrefixFilter ensures that no value is present for PrefixFilter, not even an explicit nil
func (o *MicrosoftGraphBucketAggregationDefinition) UnsetPrefixFilter() {
	o.PrefixFilter.Unset()
}

// GetRanges returns the Ranges field value if set, zero value otherwise.
func (o *MicrosoftGraphBucketAggregationDefinition) GetRanges() []*AnyOfmicrosoftGraphBucketAggregationRange {
	if o == nil || o.Ranges == nil {
		var ret []*AnyOfmicrosoftGraphBucketAggregationRange
		return ret
	}
	return *o.Ranges
}

// GetRangesOk returns a tuple with the Ranges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphBucketAggregationDefinition) GetRangesOk() (*[]*AnyOfmicrosoftGraphBucketAggregationRange, bool) {
	if o == nil || o.Ranges == nil {
		return nil, false
	}
	return o.Ranges, true
}

// HasRanges returns a boolean if a field has been set.
func (o *MicrosoftGraphBucketAggregationDefinition) HasRanges() bool {
	if o != nil && o.Ranges != nil {
		return true
	}

	return false
}

// SetRanges gets a reference to the given []*AnyOfmicrosoftGraphBucketAggregationRange and assigns it to the Ranges field.
func (o *MicrosoftGraphBucketAggregationDefinition) SetRanges(v []*AnyOfmicrosoftGraphBucketAggregationRange) {
	o.Ranges = &v
}

// GetSortBy returns the SortBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphBucketAggregationDefinition) GetSortBy() AnyOfmicrosoftGraphBucketAggregationSortProperty {
	if o == nil  {
		var ret AnyOfmicrosoftGraphBucketAggregationSortProperty
		return ret
	}
	return o.SortBy
}

// GetSortByOk returns a tuple with the SortBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphBucketAggregationDefinition) GetSortByOk() (*AnyOfmicrosoftGraphBucketAggregationSortProperty, bool) {
	if o == nil || o.SortBy == nil {
		return nil, false
	}
	return &o.SortBy, true
}

// HasSortBy returns a boolean if a field has been set.
func (o *MicrosoftGraphBucketAggregationDefinition) HasSortBy() bool {
	if o != nil && o.SortBy != nil {
		return true
	}

	return false
}

// SetSortBy gets a reference to the given AnyOfmicrosoftGraphBucketAggregationSortProperty and assigns it to the SortBy field.
func (o *MicrosoftGraphBucketAggregationDefinition) SetSortBy(v AnyOfmicrosoftGraphBucketAggregationSortProperty) {
	o.SortBy = v
}

func (o MicrosoftGraphBucketAggregationDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsDescending.IsSet() {
		toSerialize["isDescending"] = o.IsDescending.Get()
	}
	if o.MinimumCount.IsSet() {
		toSerialize["minimumCount"] = o.MinimumCount.Get()
	}
	if o.PrefixFilter.IsSet() {
		toSerialize["prefixFilter"] = o.PrefixFilter.Get()
	}
	if o.Ranges != nil {
		toSerialize["ranges"] = o.Ranges
	}
	if o.SortBy != nil {
		toSerialize["sortBy"] = o.SortBy
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphBucketAggregationDefinition struct {
	value *MicrosoftGraphBucketAggregationDefinition
	isSet bool
}

func (v NullableMicrosoftGraphBucketAggregationDefinition) Get() *MicrosoftGraphBucketAggregationDefinition {
	return v.value
}

func (v *NullableMicrosoftGraphBucketAggregationDefinition) Set(val *MicrosoftGraphBucketAggregationDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphBucketAggregationDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphBucketAggregationDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphBucketAggregationDefinition(val *MicrosoftGraphBucketAggregationDefinition) *NullableMicrosoftGraphBucketAggregationDefinition {
	return &NullableMicrosoftGraphBucketAggregationDefinition{value: val, isSet: true}
}

func (v NullableMicrosoftGraphBucketAggregationDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphBucketAggregationDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


