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

// ResourceSpecificPermissionGrant struct for ResourceSpecificPermissionGrant
type ResourceSpecificPermissionGrant struct {
	// ID of the service principal of the Azure AD app that has been granted access. Read-only.
	ClientAppId NullableString `json:"clientAppId,omitempty"`
	// ID of the Azure AD app that has been granted access. Read-only.
	ClientId NullableString `json:"clientId,omitempty"`
	// The name of the resource-specific permission. Read-only.
	Permission NullableString `json:"permission,omitempty"`
	// The type of permission. Possible values are: Application, Delegated. Read-only.
	PermissionType NullableString `json:"permissionType,omitempty"`
	// ID of the Azure AD app that is hosting the resource. Read-only.
	ResourceAppId NullableString `json:"resourceAppId,omitempty"`
}

// NewResourceSpecificPermissionGrant instantiates a new ResourceSpecificPermissionGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceSpecificPermissionGrant() *ResourceSpecificPermissionGrant {
	this := ResourceSpecificPermissionGrant{}
	return &this
}

// NewResourceSpecificPermissionGrantWithDefaults instantiates a new ResourceSpecificPermissionGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSpecificPermissionGrantWithDefaults() *ResourceSpecificPermissionGrant {
	this := ResourceSpecificPermissionGrant{}
	return &this
}

// GetClientAppId returns the ClientAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceSpecificPermissionGrant) GetClientAppId() string {
	if o == nil || o.ClientAppId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientAppId.Get()
}

// GetClientAppIdOk returns a tuple with the ClientAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceSpecificPermissionGrant) GetClientAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientAppId.Get(), o.ClientAppId.IsSet()
}

// HasClientAppId returns a boolean if a field has been set.
func (o *ResourceSpecificPermissionGrant) HasClientAppId() bool {
	if o != nil && o.ClientAppId.IsSet() {
		return true
	}

	return false
}

// SetClientAppId gets a reference to the given NullableString and assigns it to the ClientAppId field.
func (o *ResourceSpecificPermissionGrant) SetClientAppId(v string) {
	o.ClientAppId.Set(&v)
}
// SetClientAppIdNil sets the value for ClientAppId to be an explicit nil
func (o *ResourceSpecificPermissionGrant) SetClientAppIdNil() {
	o.ClientAppId.Set(nil)
}

// UnsetClientAppId ensures that no value is present for ClientAppId, not even an explicit nil
func (o *ResourceSpecificPermissionGrant) UnsetClientAppId() {
	o.ClientAppId.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceSpecificPermissionGrant) GetClientId() string {
	if o == nil || o.ClientId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceSpecificPermissionGrant) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *ResourceSpecificPermissionGrant) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *ResourceSpecificPermissionGrant) SetClientId(v string) {
	o.ClientId.Set(&v)
}
// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *ResourceSpecificPermissionGrant) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *ResourceSpecificPermissionGrant) UnsetClientId() {
	o.ClientId.Unset()
}

// GetPermission returns the Permission field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceSpecificPermissionGrant) GetPermission() string {
	if o == nil || o.Permission.Get() == nil {
		var ret string
		return ret
	}
	return *o.Permission.Get()
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceSpecificPermissionGrant) GetPermissionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Permission.Get(), o.Permission.IsSet()
}

// HasPermission returns a boolean if a field has been set.
func (o *ResourceSpecificPermissionGrant) HasPermission() bool {
	if o != nil && o.Permission.IsSet() {
		return true
	}

	return false
}

// SetPermission gets a reference to the given NullableString and assigns it to the Permission field.
func (o *ResourceSpecificPermissionGrant) SetPermission(v string) {
	o.Permission.Set(&v)
}
// SetPermissionNil sets the value for Permission to be an explicit nil
func (o *ResourceSpecificPermissionGrant) SetPermissionNil() {
	o.Permission.Set(nil)
}

// UnsetPermission ensures that no value is present for Permission, not even an explicit nil
func (o *ResourceSpecificPermissionGrant) UnsetPermission() {
	o.Permission.Unset()
}

// GetPermissionType returns the PermissionType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceSpecificPermissionGrant) GetPermissionType() string {
	if o == nil || o.PermissionType.Get() == nil {
		var ret string
		return ret
	}
	return *o.PermissionType.Get()
}

// GetPermissionTypeOk returns a tuple with the PermissionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceSpecificPermissionGrant) GetPermissionTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.PermissionType.Get(), o.PermissionType.IsSet()
}

// HasPermissionType returns a boolean if a field has been set.
func (o *ResourceSpecificPermissionGrant) HasPermissionType() bool {
	if o != nil && o.PermissionType.IsSet() {
		return true
	}

	return false
}

// SetPermissionType gets a reference to the given NullableString and assigns it to the PermissionType field.
func (o *ResourceSpecificPermissionGrant) SetPermissionType(v string) {
	o.PermissionType.Set(&v)
}
// SetPermissionTypeNil sets the value for PermissionType to be an explicit nil
func (o *ResourceSpecificPermissionGrant) SetPermissionTypeNil() {
	o.PermissionType.Set(nil)
}

// UnsetPermissionType ensures that no value is present for PermissionType, not even an explicit nil
func (o *ResourceSpecificPermissionGrant) UnsetPermissionType() {
	o.PermissionType.Unset()
}

// GetResourceAppId returns the ResourceAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ResourceSpecificPermissionGrant) GetResourceAppId() string {
	if o == nil || o.ResourceAppId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ResourceAppId.Get()
}

// GetResourceAppIdOk returns a tuple with the ResourceAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ResourceSpecificPermissionGrant) GetResourceAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ResourceAppId.Get(), o.ResourceAppId.IsSet()
}

// HasResourceAppId returns a boolean if a field has been set.
func (o *ResourceSpecificPermissionGrant) HasResourceAppId() bool {
	if o != nil && o.ResourceAppId.IsSet() {
		return true
	}

	return false
}

// SetResourceAppId gets a reference to the given NullableString and assigns it to the ResourceAppId field.
func (o *ResourceSpecificPermissionGrant) SetResourceAppId(v string) {
	o.ResourceAppId.Set(&v)
}
// SetResourceAppIdNil sets the value for ResourceAppId to be an explicit nil
func (o *ResourceSpecificPermissionGrant) SetResourceAppIdNil() {
	o.ResourceAppId.Set(nil)
}

// UnsetResourceAppId ensures that no value is present for ResourceAppId, not even an explicit nil
func (o *ResourceSpecificPermissionGrant) UnsetResourceAppId() {
	o.ResourceAppId.Unset()
}

func (o ResourceSpecificPermissionGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ClientAppId.IsSet() {
		toSerialize["clientAppId"] = o.ClientAppId.Get()
	}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.Permission.IsSet() {
		toSerialize["permission"] = o.Permission.Get()
	}
	if o.PermissionType.IsSet() {
		toSerialize["permissionType"] = o.PermissionType.Get()
	}
	if o.ResourceAppId.IsSet() {
		toSerialize["resourceAppId"] = o.ResourceAppId.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableResourceSpecificPermissionGrant struct {
	value *ResourceSpecificPermissionGrant
	isSet bool
}

func (v NullableResourceSpecificPermissionGrant) Get() *ResourceSpecificPermissionGrant {
	return v.value
}

func (v *NullableResourceSpecificPermissionGrant) Set(val *ResourceSpecificPermissionGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceSpecificPermissionGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSpecificPermissionGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSpecificPermissionGrant(val *ResourceSpecificPermissionGrant) *NullableResourceSpecificPermissionGrant {
	return &NullableResourceSpecificPermissionGrant{value: val, isSet: true}
}

func (v NullableResourceSpecificPermissionGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSpecificPermissionGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


