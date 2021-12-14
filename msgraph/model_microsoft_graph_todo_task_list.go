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

// MicrosoftGraphTodoTaskList struct for MicrosoftGraphTodoTaskList
type MicrosoftGraphTodoTaskList struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The name of the task list.
	DisplayName NullableString `json:"displayName,omitempty"`
	// True if the user is owner of the given task list.
	IsOwner *bool `json:"isOwner,omitempty"`
	// True if the task list is shared with other users
	IsShared *bool `json:"isShared,omitempty"`
	// Property indicating the list name if the given list is a well-known list. Possible values are: none, defaultList, flaggedEmails, unknownFutureValue.
	WellknownListName AnyOfmicrosoftGraphWellknownListName `json:"wellknownListName,omitempty"`
	// The collection of open extensions defined for the task list. Nullable.
	Extensions *[]MicrosoftGraphExtension `json:"extensions,omitempty"`
	// The tasks in this task list. Read-only. Nullable.
	Tasks *[]MicrosoftGraphTodoTask `json:"tasks,omitempty"`
}

// NewMicrosoftGraphTodoTaskList instantiates a new MicrosoftGraphTodoTaskList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTodoTaskList() *MicrosoftGraphTodoTaskList {
	this := MicrosoftGraphTodoTaskList{}
	return &this
}

// NewMicrosoftGraphTodoTaskListWithDefaults instantiates a new MicrosoftGraphTodoTaskList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTodoTaskListWithDefaults() *MicrosoftGraphTodoTaskList {
	this := MicrosoftGraphTodoTaskList{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphTodoTaskList) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTodoTaskList) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphTodoTaskList) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphTodoTaskList) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTodoTaskList) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTodoTaskList) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphTodoTaskList) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphTodoTaskList) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphTodoTaskList) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphTodoTaskList) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetIsOwner returns the IsOwner field value if set, zero value otherwise.
func (o *MicrosoftGraphTodoTaskList) GetIsOwner() bool {
	if o == nil || o.IsOwner == nil {
		var ret bool
		return ret
	}
	return *o.IsOwner
}

// GetIsOwnerOk returns a tuple with the IsOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTodoTaskList) GetIsOwnerOk() (*bool, bool) {
	if o == nil || o.IsOwner == nil {
		return nil, false
	}
	return o.IsOwner, true
}

// HasIsOwner returns a boolean if a field has been set.
func (o *MicrosoftGraphTodoTaskList) HasIsOwner() bool {
	if o != nil && o.IsOwner != nil {
		return true
	}

	return false
}

// SetIsOwner gets a reference to the given bool and assigns it to the IsOwner field.
func (o *MicrosoftGraphTodoTaskList) SetIsOwner(v bool) {
	o.IsOwner = &v
}

// GetIsShared returns the IsShared field value if set, zero value otherwise.
func (o *MicrosoftGraphTodoTaskList) GetIsShared() bool {
	if o == nil || o.IsShared == nil {
		var ret bool
		return ret
	}
	return *o.IsShared
}

// GetIsSharedOk returns a tuple with the IsShared field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTodoTaskList) GetIsSharedOk() (*bool, bool) {
	if o == nil || o.IsShared == nil {
		return nil, false
	}
	return o.IsShared, true
}

// HasIsShared returns a boolean if a field has been set.
func (o *MicrosoftGraphTodoTaskList) HasIsShared() bool {
	if o != nil && o.IsShared != nil {
		return true
	}

	return false
}

// SetIsShared gets a reference to the given bool and assigns it to the IsShared field.
func (o *MicrosoftGraphTodoTaskList) SetIsShared(v bool) {
	o.IsShared = &v
}

// GetWellknownListName returns the WellknownListName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTodoTaskList) GetWellknownListName() AnyOfmicrosoftGraphWellknownListName {
	if o == nil  {
		var ret AnyOfmicrosoftGraphWellknownListName
		return ret
	}
	return o.WellknownListName
}

// GetWellknownListNameOk returns a tuple with the WellknownListName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTodoTaskList) GetWellknownListNameOk() (*AnyOfmicrosoftGraphWellknownListName, bool) {
	if o == nil || o.WellknownListName == nil {
		return nil, false
	}
	return &o.WellknownListName, true
}

// HasWellknownListName returns a boolean if a field has been set.
func (o *MicrosoftGraphTodoTaskList) HasWellknownListName() bool {
	if o != nil && o.WellknownListName != nil {
		return true
	}

	return false
}

// SetWellknownListName gets a reference to the given AnyOfmicrosoftGraphWellknownListName and assigns it to the WellknownListName field.
func (o *MicrosoftGraphTodoTaskList) SetWellknownListName(v AnyOfmicrosoftGraphWellknownListName) {
	o.WellknownListName = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *MicrosoftGraphTodoTaskList) GetExtensions() []MicrosoftGraphExtension {
	if o == nil || o.Extensions == nil {
		var ret []MicrosoftGraphExtension
		return ret
	}
	return *o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTodoTaskList) GetExtensionsOk() (*[]MicrosoftGraphExtension, bool) {
	if o == nil || o.Extensions == nil {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *MicrosoftGraphTodoTaskList) HasExtensions() bool {
	if o != nil && o.Extensions != nil {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []MicrosoftGraphExtension and assigns it to the Extensions field.
func (o *MicrosoftGraphTodoTaskList) SetExtensions(v []MicrosoftGraphExtension) {
	o.Extensions = &v
}

// GetTasks returns the Tasks field value if set, zero value otherwise.
func (o *MicrosoftGraphTodoTaskList) GetTasks() []MicrosoftGraphTodoTask {
	if o == nil || o.Tasks == nil {
		var ret []MicrosoftGraphTodoTask
		return ret
	}
	return *o.Tasks
}

// GetTasksOk returns a tuple with the Tasks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTodoTaskList) GetTasksOk() (*[]MicrosoftGraphTodoTask, bool) {
	if o == nil || o.Tasks == nil {
		return nil, false
	}
	return o.Tasks, true
}

// HasTasks returns a boolean if a field has been set.
func (o *MicrosoftGraphTodoTaskList) HasTasks() bool {
	if o != nil && o.Tasks != nil {
		return true
	}

	return false
}

// SetTasks gets a reference to the given []MicrosoftGraphTodoTask and assigns it to the Tasks field.
func (o *MicrosoftGraphTodoTaskList) SetTasks(v []MicrosoftGraphTodoTask) {
	o.Tasks = &v
}

func (o MicrosoftGraphTodoTaskList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.IsOwner != nil {
		toSerialize["isOwner"] = o.IsOwner
	}
	if o.IsShared != nil {
		toSerialize["isShared"] = o.IsShared
	}
	if o.WellknownListName != nil {
		toSerialize["wellknownListName"] = o.WellknownListName
	}
	if o.Extensions != nil {
		toSerialize["extensions"] = o.Extensions
	}
	if o.Tasks != nil {
		toSerialize["tasks"] = o.Tasks
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTodoTaskList struct {
	value *MicrosoftGraphTodoTaskList
	isSet bool
}

func (v NullableMicrosoftGraphTodoTaskList) Get() *MicrosoftGraphTodoTaskList {
	return v.value
}

func (v *NullableMicrosoftGraphTodoTaskList) Set(val *MicrosoftGraphTodoTaskList) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTodoTaskList) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTodoTaskList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTodoTaskList(val *MicrosoftGraphTodoTaskList) *NullableMicrosoftGraphTodoTaskList {
	return &NullableMicrosoftGraphTodoTaskList{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTodoTaskList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTodoTaskList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


