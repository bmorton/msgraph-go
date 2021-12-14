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

// MicrosoftGraphObjectIdentity struct for MicrosoftGraphObjectIdentity
type MicrosoftGraphObjectIdentity struct {
	// Specifies the issuer of the identity, for example facebook.com.For local accounts (where signInType is not federated), this property is the local B2C tenant default domain name, for example contoso.onmicrosoft.com.For external users from other Azure AD organization, this will be the domain of the federated organization, for example contoso.com.Supports $filter. 512 character limit.
	Issuer NullableString `json:"issuer,omitempty"`
	// Specifies the unique identifier assigned to the user by the issuer. The combination of issuer and issuerAssignedId must be unique within the organization. Represents the sign-in name for the user, when signInType is set to emailAddress or userName (also known as local accounts).When signInType is set to: emailAddress, (or a custom string that starts with emailAddress like emailAddress1) issuerAssignedId must be a valid email addressuserName, issuerAssignedId must be a valid local part of an email addressSupports $filter. 100 character limit.
	IssuerAssignedId NullableString `json:"issuerAssignedId,omitempty"`
	// Specifies the user sign-in types in your directory, such as emailAddress, userName, federated, or userPrincipalName. federated represents a unique identifier for a user from an issuer, that can be in any format chosen by the issuer. Setting or updating a userPrincipalName identity will update the value of the userPrincipalName property on the user object. The validations performed on the userPrincipalName property on the user object, for example, verified domains and acceptable characters, will be performed when setting or updating a userPrincipalName identity. Additional validation is enforced on issuerAssignedId when the sign-in type is set to emailAddress or userName. This property can also be set to any custom string.
	SignInType NullableString `json:"signInType,omitempty"`
}

// NewMicrosoftGraphObjectIdentity instantiates a new MicrosoftGraphObjectIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphObjectIdentity() *MicrosoftGraphObjectIdentity {
	this := MicrosoftGraphObjectIdentity{}
	return &this
}

// NewMicrosoftGraphObjectIdentityWithDefaults instantiates a new MicrosoftGraphObjectIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphObjectIdentityWithDefaults() *MicrosoftGraphObjectIdentity {
	this := MicrosoftGraphObjectIdentity{}
	return &this
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphObjectIdentity) GetIssuer() string {
	if o == nil || o.Issuer.Get() == nil {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphObjectIdentity) GetIssuerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *MicrosoftGraphObjectIdentity) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *MicrosoftGraphObjectIdentity) SetIssuer(v string) {
	o.Issuer.Set(&v)
}
// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *MicrosoftGraphObjectIdentity) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *MicrosoftGraphObjectIdentity) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetIssuerAssignedId returns the IssuerAssignedId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphObjectIdentity) GetIssuerAssignedId() string {
	if o == nil || o.IssuerAssignedId.Get() == nil {
		var ret string
		return ret
	}
	return *o.IssuerAssignedId.Get()
}

// GetIssuerAssignedIdOk returns a tuple with the IssuerAssignedId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphObjectIdentity) GetIssuerAssignedIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IssuerAssignedId.Get(), o.IssuerAssignedId.IsSet()
}

// HasIssuerAssignedId returns a boolean if a field has been set.
func (o *MicrosoftGraphObjectIdentity) HasIssuerAssignedId() bool {
	if o != nil && o.IssuerAssignedId.IsSet() {
		return true
	}

	return false
}

// SetIssuerAssignedId gets a reference to the given NullableString and assigns it to the IssuerAssignedId field.
func (o *MicrosoftGraphObjectIdentity) SetIssuerAssignedId(v string) {
	o.IssuerAssignedId.Set(&v)
}
// SetIssuerAssignedIdNil sets the value for IssuerAssignedId to be an explicit nil
func (o *MicrosoftGraphObjectIdentity) SetIssuerAssignedIdNil() {
	o.IssuerAssignedId.Set(nil)
}

// UnsetIssuerAssignedId ensures that no value is present for IssuerAssignedId, not even an explicit nil
func (o *MicrosoftGraphObjectIdentity) UnsetIssuerAssignedId() {
	o.IssuerAssignedId.Unset()
}

// GetSignInType returns the SignInType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphObjectIdentity) GetSignInType() string {
	if o == nil || o.SignInType.Get() == nil {
		var ret string
		return ret
	}
	return *o.SignInType.Get()
}

// GetSignInTypeOk returns a tuple with the SignInType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphObjectIdentity) GetSignInTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SignInType.Get(), o.SignInType.IsSet()
}

// HasSignInType returns a boolean if a field has been set.
func (o *MicrosoftGraphObjectIdentity) HasSignInType() bool {
	if o != nil && o.SignInType.IsSet() {
		return true
	}

	return false
}

// SetSignInType gets a reference to the given NullableString and assigns it to the SignInType field.
func (o *MicrosoftGraphObjectIdentity) SetSignInType(v string) {
	o.SignInType.Set(&v)
}
// SetSignInTypeNil sets the value for SignInType to be an explicit nil
func (o *MicrosoftGraphObjectIdentity) SetSignInTypeNil() {
	o.SignInType.Set(nil)
}

// UnsetSignInType ensures that no value is present for SignInType, not even an explicit nil
func (o *MicrosoftGraphObjectIdentity) UnsetSignInType() {
	o.SignInType.Unset()
}

func (o MicrosoftGraphObjectIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if o.IssuerAssignedId.IsSet() {
		toSerialize["issuerAssignedId"] = o.IssuerAssignedId.Get()
	}
	if o.SignInType.IsSet() {
		toSerialize["signInType"] = o.SignInType.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphObjectIdentity struct {
	value *MicrosoftGraphObjectIdentity
	isSet bool
}

func (v NullableMicrosoftGraphObjectIdentity) Get() *MicrosoftGraphObjectIdentity {
	return v.value
}

func (v *NullableMicrosoftGraphObjectIdentity) Set(val *MicrosoftGraphObjectIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphObjectIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphObjectIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphObjectIdentity(val *MicrosoftGraphObjectIdentity) *NullableMicrosoftGraphObjectIdentity {
	return &NullableMicrosoftGraphObjectIdentity{value: val, isSet: true}
}

func (v NullableMicrosoftGraphObjectIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphObjectIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


