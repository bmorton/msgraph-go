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

// CloudAppSecuritySessionControl struct for CloudAppSecuritySessionControl
type CloudAppSecuritySessionControl struct {
	// Possible values are: mcasConfigured, monitorOnly, blockDownloads, unknownFutureValue. For more information, see Deploy Conditional Access App Control for featured apps.
	CloudAppSecurityType AnyOfmicrosoftGraphCloudAppSecuritySessionControlType `json:"cloudAppSecurityType,omitempty"`
}

// NewCloudAppSecuritySessionControl instantiates a new CloudAppSecuritySessionControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudAppSecuritySessionControl() *CloudAppSecuritySessionControl {
	this := CloudAppSecuritySessionControl{}
	return &this
}

// NewCloudAppSecuritySessionControlWithDefaults instantiates a new CloudAppSecuritySessionControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudAppSecuritySessionControlWithDefaults() *CloudAppSecuritySessionControl {
	this := CloudAppSecuritySessionControl{}
	return &this
}

// GetCloudAppSecurityType returns the CloudAppSecurityType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CloudAppSecuritySessionControl) GetCloudAppSecurityType() AnyOfmicrosoftGraphCloudAppSecuritySessionControlType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphCloudAppSecuritySessionControlType
		return ret
	}
	return o.CloudAppSecurityType
}

// GetCloudAppSecurityTypeOk returns a tuple with the CloudAppSecurityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CloudAppSecuritySessionControl) GetCloudAppSecurityTypeOk() (*AnyOfmicrosoftGraphCloudAppSecuritySessionControlType, bool) {
	if o == nil || o.CloudAppSecurityType == nil {
		return nil, false
	}
	return &o.CloudAppSecurityType, true
}

// HasCloudAppSecurityType returns a boolean if a field has been set.
func (o *CloudAppSecuritySessionControl) HasCloudAppSecurityType() bool {
	if o != nil && o.CloudAppSecurityType != nil {
		return true
	}

	return false
}

// SetCloudAppSecurityType gets a reference to the given AnyOfmicrosoftGraphCloudAppSecuritySessionControlType and assigns it to the CloudAppSecurityType field.
func (o *CloudAppSecuritySessionControl) SetCloudAppSecurityType(v AnyOfmicrosoftGraphCloudAppSecuritySessionControlType) {
	o.CloudAppSecurityType = v
}

func (o CloudAppSecuritySessionControl) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudAppSecurityType != nil {
		toSerialize["cloudAppSecurityType"] = o.CloudAppSecurityType
	}
	return json.Marshal(toSerialize)
}

type NullableCloudAppSecuritySessionControl struct {
	value *CloudAppSecuritySessionControl
	isSet bool
}

func (v NullableCloudAppSecuritySessionControl) Get() *CloudAppSecuritySessionControl {
	return v.value
}

func (v *NullableCloudAppSecuritySessionControl) Set(val *CloudAppSecuritySessionControl) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudAppSecuritySessionControl) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudAppSecuritySessionControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudAppSecuritySessionControl(val *CloudAppSecuritySessionControl) *NullableCloudAppSecuritySessionControl {
	return &NullableCloudAppSecuritySessionControl{value: val, isSet: true}
}

func (v NullableCloudAppSecuritySessionControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudAppSecuritySessionControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


