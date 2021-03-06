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

// MicrosoftGraphAccessPackageSubject struct for MicrosoftGraphAccessPackageSubject
type MicrosoftGraphAccessPackageSubject struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// The display name of the subject.
	DisplayName NullableString `json:"displayName,omitempty"`
	// The email address of the subject.
	Email NullableString `json:"email,omitempty"`
	// The object identifier of the subject. null if the subject is not yet a user in the tenant.
	ObjectId NullableString `json:"objectId,omitempty"`
	// A string representation of the principal's security identifier, if known, or null if the subject does not have a security identifier.
	OnPremisesSecurityIdentifier NullableString `json:"onPremisesSecurityIdentifier,omitempty"`
	// The principal name, if known, of the subject.
	PrincipalName NullableString `json:"principalName,omitempty"`
	// The resource type of the subject. The possible values are: notSpecified, user, servicePrincipal, unknownFutureValue.
	SubjectType AnyOfmicrosoftGraphAccessPackageSubjectType `json:"subjectType,omitempty"`
	// The connected organization of the subject. Read-only. Nullable.
	ConnectedOrganization AnyOfmicrosoftGraphConnectedOrganization `json:"connectedOrganization,omitempty"`
}

// NewMicrosoftGraphAccessPackageSubject instantiates a new MicrosoftGraphAccessPackageSubject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAccessPackageSubject() *MicrosoftGraphAccessPackageSubject {
	this := MicrosoftGraphAccessPackageSubject{}
	return &this
}

// NewMicrosoftGraphAccessPackageSubjectWithDefaults instantiates a new MicrosoftGraphAccessPackageSubject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAccessPackageSubjectWithDefaults() *MicrosoftGraphAccessPackageSubject {
	this := MicrosoftGraphAccessPackageSubject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphAccessPackageSubject) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAccessPackageSubject) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageSubject) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAccessPackageSubject) SetId(v string) {
	o.Id = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageSubject) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageSubject) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageSubject) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphAccessPackageSubject) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageSubject) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageSubject) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageSubject) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *MicrosoftGraphAccessPackageSubject) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) UnsetEmail() {
	o.Email.Unset()
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageSubject) GetObjectId() string {
	if o == nil || o.ObjectId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ObjectId.Get()
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageSubject) GetObjectIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ObjectId.Get(), o.ObjectId.IsSet()
}

// HasObjectId returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageSubject) HasObjectId() bool {
	if o != nil && o.ObjectId.IsSet() {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given NullableString and assigns it to the ObjectId field.
func (o *MicrosoftGraphAccessPackageSubject) SetObjectId(v string) {
	o.ObjectId.Set(&v)
}
// SetObjectIdNil sets the value for ObjectId to be an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) SetObjectIdNil() {
	o.ObjectId.Set(nil)
}

// UnsetObjectId ensures that no value is present for ObjectId, not even an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) UnsetObjectId() {
	o.ObjectId.Unset()
}

// GetOnPremisesSecurityIdentifier returns the OnPremisesSecurityIdentifier field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageSubject) GetOnPremisesSecurityIdentifier() string {
	if o == nil || o.OnPremisesSecurityIdentifier.Get() == nil {
		var ret string
		return ret
	}
	return *o.OnPremisesSecurityIdentifier.Get()
}

// GetOnPremisesSecurityIdentifierOk returns a tuple with the OnPremisesSecurityIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageSubject) GetOnPremisesSecurityIdentifierOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OnPremisesSecurityIdentifier.Get(), o.OnPremisesSecurityIdentifier.IsSet()
}

// HasOnPremisesSecurityIdentifier returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageSubject) HasOnPremisesSecurityIdentifier() bool {
	if o != nil && o.OnPremisesSecurityIdentifier.IsSet() {
		return true
	}

	return false
}

// SetOnPremisesSecurityIdentifier gets a reference to the given NullableString and assigns it to the OnPremisesSecurityIdentifier field.
func (o *MicrosoftGraphAccessPackageSubject) SetOnPremisesSecurityIdentifier(v string) {
	o.OnPremisesSecurityIdentifier.Set(&v)
}
// SetOnPremisesSecurityIdentifierNil sets the value for OnPremisesSecurityIdentifier to be an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) SetOnPremisesSecurityIdentifierNil() {
	o.OnPremisesSecurityIdentifier.Set(nil)
}

// UnsetOnPremisesSecurityIdentifier ensures that no value is present for OnPremisesSecurityIdentifier, not even an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) UnsetOnPremisesSecurityIdentifier() {
	o.OnPremisesSecurityIdentifier.Unset()
}

// GetPrincipalName returns the PrincipalName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageSubject) GetPrincipalName() string {
	if o == nil || o.PrincipalName.Get() == nil {
		var ret string
		return ret
	}
	return *o.PrincipalName.Get()
}

// GetPrincipalNameOk returns a tuple with the PrincipalName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageSubject) GetPrincipalNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PrincipalName.Get(), o.PrincipalName.IsSet()
}

// HasPrincipalName returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageSubject) HasPrincipalName() bool {
	if o != nil && o.PrincipalName.IsSet() {
		return true
	}

	return false
}

// SetPrincipalName gets a reference to the given NullableString and assigns it to the PrincipalName field.
func (o *MicrosoftGraphAccessPackageSubject) SetPrincipalName(v string) {
	o.PrincipalName.Set(&v)
}
// SetPrincipalNameNil sets the value for PrincipalName to be an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) SetPrincipalNameNil() {
	o.PrincipalName.Set(nil)
}

// UnsetPrincipalName ensures that no value is present for PrincipalName, not even an explicit nil
func (o *MicrosoftGraphAccessPackageSubject) UnsetPrincipalName() {
	o.PrincipalName.Unset()
}

// GetSubjectType returns the SubjectType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageSubject) GetSubjectType() AnyOfmicrosoftGraphAccessPackageSubjectType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphAccessPackageSubjectType
		return ret
	}
	return o.SubjectType
}

// GetSubjectTypeOk returns a tuple with the SubjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageSubject) GetSubjectTypeOk() (*AnyOfmicrosoftGraphAccessPackageSubjectType, bool) {
	if o == nil || o.SubjectType == nil {
		return nil, false
	}
	return &o.SubjectType, true
}

// HasSubjectType returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageSubject) HasSubjectType() bool {
	if o != nil && o.SubjectType != nil {
		return true
	}

	return false
}

// SetSubjectType gets a reference to the given AnyOfmicrosoftGraphAccessPackageSubjectType and assigns it to the SubjectType field.
func (o *MicrosoftGraphAccessPackageSubject) SetSubjectType(v AnyOfmicrosoftGraphAccessPackageSubjectType) {
	o.SubjectType = v
}

// GetConnectedOrganization returns the ConnectedOrganization field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphAccessPackageSubject) GetConnectedOrganization() AnyOfmicrosoftGraphConnectedOrganization {
	if o == nil  {
		var ret AnyOfmicrosoftGraphConnectedOrganization
		return ret
	}
	return o.ConnectedOrganization
}

// GetConnectedOrganizationOk returns a tuple with the ConnectedOrganization field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphAccessPackageSubject) GetConnectedOrganizationOk() (*AnyOfmicrosoftGraphConnectedOrganization, bool) {
	if o == nil || o.ConnectedOrganization == nil {
		return nil, false
	}
	return &o.ConnectedOrganization, true
}

// HasConnectedOrganization returns a boolean if a field has been set.
func (o *MicrosoftGraphAccessPackageSubject) HasConnectedOrganization() bool {
	if o != nil && o.ConnectedOrganization != nil {
		return true
	}

	return false
}

// SetConnectedOrganization gets a reference to the given AnyOfmicrosoftGraphConnectedOrganization and assigns it to the ConnectedOrganization field.
func (o *MicrosoftGraphAccessPackageSubject) SetConnectedOrganization(v AnyOfmicrosoftGraphConnectedOrganization) {
	o.ConnectedOrganization = v
}

func (o MicrosoftGraphAccessPackageSubject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.ObjectId.IsSet() {
		toSerialize["objectId"] = o.ObjectId.Get()
	}
	if o.OnPremisesSecurityIdentifier.IsSet() {
		toSerialize["onPremisesSecurityIdentifier"] = o.OnPremisesSecurityIdentifier.Get()
	}
	if o.PrincipalName.IsSet() {
		toSerialize["principalName"] = o.PrincipalName.Get()
	}
	if o.SubjectType != nil {
		toSerialize["subjectType"] = o.SubjectType
	}
	if o.ConnectedOrganization != nil {
		toSerialize["connectedOrganization"] = o.ConnectedOrganization
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAccessPackageSubject struct {
	value *MicrosoftGraphAccessPackageSubject
	isSet bool
}

func (v NullableMicrosoftGraphAccessPackageSubject) Get() *MicrosoftGraphAccessPackageSubject {
	return v.value
}

func (v *NullableMicrosoftGraphAccessPackageSubject) Set(val *MicrosoftGraphAccessPackageSubject) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAccessPackageSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAccessPackageSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAccessPackageSubject(val *MicrosoftGraphAccessPackageSubject) *NullableMicrosoftGraphAccessPackageSubject {
	return &NullableMicrosoftGraphAccessPackageSubject{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAccessPackageSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAccessPackageSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


