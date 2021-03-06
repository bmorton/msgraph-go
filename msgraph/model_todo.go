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

// Todo struct for Todo
type Todo struct {
	// The task lists in the users mailbox.
	Lists *[]MicrosoftGraphTodoTaskList `json:"lists,omitempty"`
}

// NewTodo instantiates a new Todo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTodo() *Todo {
	this := Todo{}
	return &this
}

// NewTodoWithDefaults instantiates a new Todo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTodoWithDefaults() *Todo {
	this := Todo{}
	return &this
}

// GetLists returns the Lists field value if set, zero value otherwise.
func (o *Todo) GetLists() []MicrosoftGraphTodoTaskList {
	if o == nil || o.Lists == nil {
		var ret []MicrosoftGraphTodoTaskList
		return ret
	}
	return *o.Lists
}

// GetListsOk returns a tuple with the Lists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Todo) GetListsOk() (*[]MicrosoftGraphTodoTaskList, bool) {
	if o == nil || o.Lists == nil {
		return nil, false
	}
	return o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *Todo) HasLists() bool {
	if o != nil && o.Lists != nil {
		return true
	}

	return false
}

// SetLists gets a reference to the given []MicrosoftGraphTodoTaskList and assigns it to the Lists field.
func (o *Todo) SetLists(v []MicrosoftGraphTodoTaskList) {
	o.Lists = &v
}

func (o Todo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Lists != nil {
		toSerialize["lists"] = o.Lists
	}
	return json.Marshal(toSerialize)
}

type NullableTodo struct {
	value *Todo
	isSet bool
}

func (v NullableTodo) Get() *Todo {
	return v.value
}

func (v *NullableTodo) Set(val *Todo) {
	v.value = val
	v.isSet = true
}

func (v NullableTodo) IsSet() bool {
	return v.isSet
}

func (v *NullableTodo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTodo(val *Todo) *NullableTodo {
	return &NullableTodo{value: val, isSet: true}
}

func (v NullableTodo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTodo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


