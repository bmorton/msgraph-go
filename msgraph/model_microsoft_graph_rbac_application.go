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

// MicrosoftGraphRbacApplication struct for MicrosoftGraphRbacApplication
type MicrosoftGraphRbacApplication struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Resource to grant access to users or groups.
	RoleAssignments *[]MicrosoftGraphUnifiedRoleAssignment `json:"roleAssignments,omitempty"`
	// Resource representing the roles allowed by RBAC providers and the permissions assigned to the roles.
	RoleDefinitions *[]MicrosoftGraphUnifiedRoleDefinition `json:"roleDefinitions,omitempty"`
}

// NewMicrosoftGraphRbacApplication instantiates a new MicrosoftGraphRbacApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphRbacApplication() *MicrosoftGraphRbacApplication {
	this := MicrosoftGraphRbacApplication{}
	return &this
}

// NewMicrosoftGraphRbacApplicationWithDefaults instantiates a new MicrosoftGraphRbacApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphRbacApplicationWithDefaults() *MicrosoftGraphRbacApplication {
	this := MicrosoftGraphRbacApplication{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphRbacApplication) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRbacApplication) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphRbacApplication) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphRbacApplication) SetId(v string) {
	o.Id = &v
}

// GetRoleAssignments returns the RoleAssignments field value if set, zero value otherwise.
func (o *MicrosoftGraphRbacApplication) GetRoleAssignments() []MicrosoftGraphUnifiedRoleAssignment {
	if o == nil || o.RoleAssignments == nil {
		var ret []MicrosoftGraphUnifiedRoleAssignment
		return ret
	}
	return *o.RoleAssignments
}

// GetRoleAssignmentsOk returns a tuple with the RoleAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRbacApplication) GetRoleAssignmentsOk() (*[]MicrosoftGraphUnifiedRoleAssignment, bool) {
	if o == nil || o.RoleAssignments == nil {
		return nil, false
	}
	return o.RoleAssignments, true
}

// HasRoleAssignments returns a boolean if a field has been set.
func (o *MicrosoftGraphRbacApplication) HasRoleAssignments() bool {
	if o != nil && o.RoleAssignments != nil {
		return true
	}

	return false
}

// SetRoleAssignments gets a reference to the given []MicrosoftGraphUnifiedRoleAssignment and assigns it to the RoleAssignments field.
func (o *MicrosoftGraphRbacApplication) SetRoleAssignments(v []MicrosoftGraphUnifiedRoleAssignment) {
	o.RoleAssignments = &v
}

// GetRoleDefinitions returns the RoleDefinitions field value if set, zero value otherwise.
func (o *MicrosoftGraphRbacApplication) GetRoleDefinitions() []MicrosoftGraphUnifiedRoleDefinition {
	if o == nil || o.RoleDefinitions == nil {
		var ret []MicrosoftGraphUnifiedRoleDefinition
		return ret
	}
	return *o.RoleDefinitions
}

// GetRoleDefinitionsOk returns a tuple with the RoleDefinitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphRbacApplication) GetRoleDefinitionsOk() (*[]MicrosoftGraphUnifiedRoleDefinition, bool) {
	if o == nil || o.RoleDefinitions == nil {
		return nil, false
	}
	return o.RoleDefinitions, true
}

// HasRoleDefinitions returns a boolean if a field has been set.
func (o *MicrosoftGraphRbacApplication) HasRoleDefinitions() bool {
	if o != nil && o.RoleDefinitions != nil {
		return true
	}

	return false
}

// SetRoleDefinitions gets a reference to the given []MicrosoftGraphUnifiedRoleDefinition and assigns it to the RoleDefinitions field.
func (o *MicrosoftGraphRbacApplication) SetRoleDefinitions(v []MicrosoftGraphUnifiedRoleDefinition) {
	o.RoleDefinitions = &v
}

func (o MicrosoftGraphRbacApplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RoleAssignments != nil {
		toSerialize["roleAssignments"] = o.RoleAssignments
	}
	if o.RoleDefinitions != nil {
		toSerialize["roleDefinitions"] = o.RoleDefinitions
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphRbacApplication struct {
	value *MicrosoftGraphRbacApplication
	isSet bool
}

func (v NullableMicrosoftGraphRbacApplication) Get() *MicrosoftGraphRbacApplication {
	return v.value
}

func (v *NullableMicrosoftGraphRbacApplication) Set(val *MicrosoftGraphRbacApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphRbacApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphRbacApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphRbacApplication(val *MicrosoftGraphRbacApplication) *NullableMicrosoftGraphRbacApplication {
	return &NullableMicrosoftGraphRbacApplication{value: val, isSet: true}
}

func (v NullableMicrosoftGraphRbacApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphRbacApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


