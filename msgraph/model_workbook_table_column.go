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

// WorkbookTableColumn struct for WorkbookTableColumn
type WorkbookTableColumn struct {
	// Returns the index number of the column within the columns collection of the table. Zero-indexed. Read-only.
	Index *int32 `json:"index,omitempty"`
	// Returns the name of the table column.
	Name NullableString `json:"name,omitempty"`
	// Represents the raw values of the specified range. The data returned could be of type string, number, or a boolean. Cell that contain an error will return the error string.
	Values AnyOfobject `json:"values,omitempty"`
	// Retrieve the filter applied to the column. Read-only.
	Filter AnyOfmicrosoftGraphWorkbookFilter `json:"filter,omitempty"`
}

// NewWorkbookTableColumn instantiates a new WorkbookTableColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkbookTableColumn() *WorkbookTableColumn {
	this := WorkbookTableColumn{}
	return &this
}

// NewWorkbookTableColumnWithDefaults instantiates a new WorkbookTableColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkbookTableColumnWithDefaults() *WorkbookTableColumn {
	this := WorkbookTableColumn{}
	return &this
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *WorkbookTableColumn) GetIndex() int32 {
	if o == nil || o.Index == nil {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkbookTableColumn) GetIndexOk() (*int32, bool) {
	if o == nil || o.Index == nil {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *WorkbookTableColumn) HasIndex() bool {
	if o != nil && o.Index != nil {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *WorkbookTableColumn) SetIndex(v int32) {
	o.Index = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookTableColumn) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookTableColumn) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *WorkbookTableColumn) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *WorkbookTableColumn) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *WorkbookTableColumn) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *WorkbookTableColumn) UnsetName() {
	o.Name.Unset()
}

// GetValues returns the Values field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookTableColumn) GetValues() AnyOfobject {
	if o == nil  {
		var ret AnyOfobject
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookTableColumn) GetValuesOk() (*AnyOfobject, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return &o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *WorkbookTableColumn) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given AnyOfobject and assigns it to the Values field.
func (o *WorkbookTableColumn) SetValues(v AnyOfobject) {
	o.Values = v
}

// GetFilter returns the Filter field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkbookTableColumn) GetFilter() AnyOfmicrosoftGraphWorkbookFilter {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookFilter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkbookTableColumn) GetFilterOk() (*AnyOfmicrosoftGraphWorkbookFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return &o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *WorkbookTableColumn) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given AnyOfmicrosoftGraphWorkbookFilter and assigns it to the Filter field.
func (o *WorkbookTableColumn) SetFilter(v AnyOfmicrosoftGraphWorkbookFilter) {
	o.Filter = v
}

func (o WorkbookTableColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Index != nil {
		toSerialize["index"] = o.Index
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	return json.Marshal(toSerialize)
}

type NullableWorkbookTableColumn struct {
	value *WorkbookTableColumn
	isSet bool
}

func (v NullableWorkbookTableColumn) Get() *WorkbookTableColumn {
	return v.value
}

func (v *NullableWorkbookTableColumn) Set(val *WorkbookTableColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkbookTableColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkbookTableColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkbookTableColumn(val *WorkbookTableColumn) *NullableWorkbookTableColumn {
	return &NullableWorkbookTableColumn{value: val, isSet: true}
}

func (v NullableWorkbookTableColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkbookTableColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


