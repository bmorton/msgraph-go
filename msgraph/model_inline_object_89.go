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

// InlineObject89 struct for InlineObject89
type InlineObject89 struct {
	Name NullableString `json:"name,omitempty"`
	Select *[]*string `json:"select,omitempty"`
	Search NullableString `json:"search,omitempty"`
	GroupBy *[]*string `json:"groupBy,omitempty"`
	OrderBy *[]*string `json:"orderBy,omitempty"`
	Skip NullableInt32 `json:"skip,omitempty"`
	Top NullableInt32 `json:"top,omitempty"`
	SessionId NullableString `json:"sessionId,omitempty"`
	Filter NullableString `json:"filter,omitempty"`
}

// NewInlineObject89 instantiates a new InlineObject89 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject89() *InlineObject89 {
	this := InlineObject89{}
	return &this
}

// NewInlineObject89WithDefaults instantiates a new InlineObject89 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject89WithDefaults() *InlineObject89 {
	this := InlineObject89{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject89) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject89) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *InlineObject89) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *InlineObject89) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *InlineObject89) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *InlineObject89) UnsetName() {
	o.Name.Unset()
}

// GetSelect returns the Select field value if set, zero value otherwise.
func (o *InlineObject89) GetSelect() []*string {
	if o == nil || o.Select == nil {
		var ret []*string
		return ret
	}
	return *o.Select
}

// GetSelectOk returns a tuple with the Select field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject89) GetSelectOk() (*[]*string, bool) {
	if o == nil || o.Select == nil {
		return nil, false
	}
	return o.Select, true
}

// HasSelect returns a boolean if a field has been set.
func (o *InlineObject89) HasSelect() bool {
	if o != nil && o.Select != nil {
		return true
	}

	return false
}

// SetSelect gets a reference to the given []*string and assigns it to the Select field.
func (o *InlineObject89) SetSelect(v []*string) {
	o.Select = &v
}

// GetSearch returns the Search field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject89) GetSearch() string {
	if o == nil || o.Search.Get() == nil {
		var ret string
		return ret
	}
	return *o.Search.Get()
}

// GetSearchOk returns a tuple with the Search field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject89) GetSearchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Search.Get(), o.Search.IsSet()
}

// HasSearch returns a boolean if a field has been set.
func (o *InlineObject89) HasSearch() bool {
	if o != nil && o.Search.IsSet() {
		return true
	}

	return false
}

// SetSearch gets a reference to the given NullableString and assigns it to the Search field.
func (o *InlineObject89) SetSearch(v string) {
	o.Search.Set(&v)
}
// SetSearchNil sets the value for Search to be an explicit nil
func (o *InlineObject89) SetSearchNil() {
	o.Search.Set(nil)
}

// UnsetSearch ensures that no value is present for Search, not even an explicit nil
func (o *InlineObject89) UnsetSearch() {
	o.Search.Unset()
}

// GetGroupBy returns the GroupBy field value if set, zero value otherwise.
func (o *InlineObject89) GetGroupBy() []*string {
	if o == nil || o.GroupBy == nil {
		var ret []*string
		return ret
	}
	return *o.GroupBy
}

// GetGroupByOk returns a tuple with the GroupBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject89) GetGroupByOk() (*[]*string, bool) {
	if o == nil || o.GroupBy == nil {
		return nil, false
	}
	return o.GroupBy, true
}

// HasGroupBy returns a boolean if a field has been set.
func (o *InlineObject89) HasGroupBy() bool {
	if o != nil && o.GroupBy != nil {
		return true
	}

	return false
}

// SetGroupBy gets a reference to the given []*string and assigns it to the GroupBy field.
func (o *InlineObject89) SetGroupBy(v []*string) {
	o.GroupBy = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *InlineObject89) GetOrderBy() []*string {
	if o == nil || o.OrderBy == nil {
		var ret []*string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject89) GetOrderByOk() (*[]*string, bool) {
	if o == nil || o.OrderBy == nil {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *InlineObject89) HasOrderBy() bool {
	if o != nil && o.OrderBy != nil {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given []*string and assigns it to the OrderBy field.
func (o *InlineObject89) SetOrderBy(v []*string) {
	o.OrderBy = &v
}

// GetSkip returns the Skip field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject89) GetSkip() int32 {
	if o == nil || o.Skip.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Skip.Get()
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject89) GetSkipOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Skip.Get(), o.Skip.IsSet()
}

// HasSkip returns a boolean if a field has been set.
func (o *InlineObject89) HasSkip() bool {
	if o != nil && o.Skip.IsSet() {
		return true
	}

	return false
}

// SetSkip gets a reference to the given NullableInt32 and assigns it to the Skip field.
func (o *InlineObject89) SetSkip(v int32) {
	o.Skip.Set(&v)
}
// SetSkipNil sets the value for Skip to be an explicit nil
func (o *InlineObject89) SetSkipNil() {
	o.Skip.Set(nil)
}

// UnsetSkip ensures that no value is present for Skip, not even an explicit nil
func (o *InlineObject89) UnsetSkip() {
	o.Skip.Unset()
}

// GetTop returns the Top field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject89) GetTop() int32 {
	if o == nil || o.Top.Get() == nil {
		var ret int32
		return ret
	}
	return *o.Top.Get()
}

// GetTopOk returns a tuple with the Top field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject89) GetTopOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Top.Get(), o.Top.IsSet()
}

// HasTop returns a boolean if a field has been set.
func (o *InlineObject89) HasTop() bool {
	if o != nil && o.Top.IsSet() {
		return true
	}

	return false
}

// SetTop gets a reference to the given NullableInt32 and assigns it to the Top field.
func (o *InlineObject89) SetTop(v int32) {
	o.Top.Set(&v)
}
// SetTopNil sets the value for Top to be an explicit nil
func (o *InlineObject89) SetTopNil() {
	o.Top.Set(nil)
}

// UnsetTop ensures that no value is present for Top, not even an explicit nil
func (o *InlineObject89) UnsetTop() {
	o.Top.Unset()
}

// GetSessionId returns the SessionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject89) GetSessionId() string {
	if o == nil || o.SessionId.Get() == nil {
		var ret string
		return ret
	}
	return *o.SessionId.Get()
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject89) GetSessionIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SessionId.Get(), o.SessionId.IsSet()
}

// HasSessionId returns a boolean if a field has been set.
func (o *InlineObject89) HasSessionId() bool {
	if o != nil && o.SessionId.IsSet() {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given NullableString and assigns it to the SessionId field.
func (o *InlineObject89) SetSessionId(v string) {
	o.SessionId.Set(&v)
}
// SetSessionIdNil sets the value for SessionId to be an explicit nil
func (o *InlineObject89) SetSessionIdNil() {
	o.SessionId.Set(nil)
}

// UnsetSessionId ensures that no value is present for SessionId, not even an explicit nil
func (o *InlineObject89) UnsetSessionId() {
	o.SessionId.Unset()
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *InlineObject89) GetFilter() string {
	if o == nil || o.Filter.Get() == nil {
		var ret string
		return ret
	}
	return *o.Filter.Get()
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *InlineObject89) GetFilterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Filter.Get(), o.Filter.IsSet()
}

// HasFilter returns a boolean if a field has been set.
func (o *InlineObject89) HasFilter() bool {
	if o != nil && o.Filter.IsSet() {
		return true
	}

	return false
}

// SetFilter gets a reference to the given NullableString and assigns it to the Filter field.
func (o *InlineObject89) SetFilter(v string) {
	o.Filter.Set(&v)
}
// SetFilterNil sets the value for Filter to be an explicit nil
func (o *InlineObject89) SetFilterNil() {
	o.Filter.Set(nil)
}

// UnsetFilter ensures that no value is present for Filter, not even an explicit nil
func (o *InlineObject89) UnsetFilter() {
	o.Filter.Unset()
}

func (o InlineObject89) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Select != nil {
		toSerialize["select"] = o.Select
	}
	if o.Search.IsSet() {
		toSerialize["search"] = o.Search.Get()
	}
	if o.GroupBy != nil {
		toSerialize["groupBy"] = o.GroupBy
	}
	if o.OrderBy != nil {
		toSerialize["orderBy"] = o.OrderBy
	}
	if o.Skip.IsSet() {
		toSerialize["skip"] = o.Skip.Get()
	}
	if o.Top.IsSet() {
		toSerialize["top"] = o.Top.Get()
	}
	if o.SessionId.IsSet() {
		toSerialize["sessionId"] = o.SessionId.Get()
	}
	if o.Filter.IsSet() {
		toSerialize["filter"] = o.Filter.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject89 struct {
	value *InlineObject89
	isSet bool
}

func (v NullableInlineObject89) Get() *InlineObject89 {
	return v.value
}

func (v *NullableInlineObject89) Set(val *InlineObject89) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject89) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject89) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject89(val *InlineObject89) *NullableInlineObject89 {
	return &NullableInlineObject89{value: val, isSet: true}
}

func (v NullableInlineObject89) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject89) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


