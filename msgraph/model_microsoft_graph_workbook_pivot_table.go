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

// MicrosoftGraphWorkbookPivotTable struct for MicrosoftGraphWorkbookPivotTable
type MicrosoftGraphWorkbookPivotTable struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Name of the PivotTable.
	Name NullableString `json:"name,omitempty"`
	// The worksheet containing the current PivotTable. Read-only.
	Worksheet AnyOfmicrosoftGraphWorkbookWorksheet `json:"worksheet,omitempty"`
}

// NewMicrosoftGraphWorkbookPivotTable instantiates a new MicrosoftGraphWorkbookPivotTable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphWorkbookPivotTable() *MicrosoftGraphWorkbookPivotTable {
	this := MicrosoftGraphWorkbookPivotTable{}
	return &this
}

// NewMicrosoftGraphWorkbookPivotTableWithDefaults instantiates a new MicrosoftGraphWorkbookPivotTable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphWorkbookPivotTableWithDefaults() *MicrosoftGraphWorkbookPivotTable {
	this := MicrosoftGraphWorkbookPivotTable{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphWorkbookPivotTable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphWorkbookPivotTable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookPivotTable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphWorkbookPivotTable) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookPivotTable) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookPivotTable) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookPivotTable) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *MicrosoftGraphWorkbookPivotTable) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *MicrosoftGraphWorkbookPivotTable) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *MicrosoftGraphWorkbookPivotTable) UnsetName() {
	o.Name.Unset()
}

// GetWorksheet returns the Worksheet field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphWorkbookPivotTable) GetWorksheet() AnyOfmicrosoftGraphWorkbookWorksheet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWorkbookWorksheet
		return ret
	}
	return o.Worksheet
}

// GetWorksheetOk returns a tuple with the Worksheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphWorkbookPivotTable) GetWorksheetOk() (*AnyOfmicrosoftGraphWorkbookWorksheet, bool) {
	if o == nil || o.Worksheet == nil {
		return nil, false
	}
	return &o.Worksheet, true
}

// HasWorksheet returns a boolean if a field has been set.
func (o *MicrosoftGraphWorkbookPivotTable) HasWorksheet() bool {
	if o != nil && o.Worksheet != nil {
		return true
	}

	return false
}

// SetWorksheet gets a reference to the given AnyOfmicrosoftGraphWorkbookWorksheet and assigns it to the Worksheet field.
func (o *MicrosoftGraphWorkbookPivotTable) SetWorksheet(v AnyOfmicrosoftGraphWorkbookWorksheet) {
	o.Worksheet = v
}

func (o MicrosoftGraphWorkbookPivotTable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.Worksheet != nil {
		toSerialize["worksheet"] = o.Worksheet
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphWorkbookPivotTable struct {
	value *MicrosoftGraphWorkbookPivotTable
	isSet bool
}

func (v NullableMicrosoftGraphWorkbookPivotTable) Get() *MicrosoftGraphWorkbookPivotTable {
	return v.value
}

func (v *NullableMicrosoftGraphWorkbookPivotTable) Set(val *MicrosoftGraphWorkbookPivotTable) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphWorkbookPivotTable) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphWorkbookPivotTable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphWorkbookPivotTable(val *MicrosoftGraphWorkbookPivotTable) *NullableMicrosoftGraphWorkbookPivotTable {
	return &NullableMicrosoftGraphWorkbookPivotTable{value: val, isSet: true}
}

func (v NullableMicrosoftGraphWorkbookPivotTable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphWorkbookPivotTable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


