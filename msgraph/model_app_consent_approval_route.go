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

// AppConsentApprovalRoute struct for AppConsentApprovalRoute
type AppConsentApprovalRoute struct {
	AppConsentRequests *[]MicrosoftGraphAppConsentRequest `json:"appConsentRequests,omitempty"`
}

// NewAppConsentApprovalRoute instantiates a new AppConsentApprovalRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppConsentApprovalRoute() *AppConsentApprovalRoute {
	this := AppConsentApprovalRoute{}
	return &this
}

// NewAppConsentApprovalRouteWithDefaults instantiates a new AppConsentApprovalRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppConsentApprovalRouteWithDefaults() *AppConsentApprovalRoute {
	this := AppConsentApprovalRoute{}
	return &this
}

// GetAppConsentRequests returns the AppConsentRequests field value if set, zero value otherwise.
func (o *AppConsentApprovalRoute) GetAppConsentRequests() []MicrosoftGraphAppConsentRequest {
	if o == nil || o.AppConsentRequests == nil {
		var ret []MicrosoftGraphAppConsentRequest
		return ret
	}
	return *o.AppConsentRequests
}

// GetAppConsentRequestsOk returns a tuple with the AppConsentRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppConsentApprovalRoute) GetAppConsentRequestsOk() (*[]MicrosoftGraphAppConsentRequest, bool) {
	if o == nil || o.AppConsentRequests == nil {
		return nil, false
	}
	return o.AppConsentRequests, true
}

// HasAppConsentRequests returns a boolean if a field has been set.
func (o *AppConsentApprovalRoute) HasAppConsentRequests() bool {
	if o != nil && o.AppConsentRequests != nil {
		return true
	}

	return false
}

// SetAppConsentRequests gets a reference to the given []MicrosoftGraphAppConsentRequest and assigns it to the AppConsentRequests field.
func (o *AppConsentApprovalRoute) SetAppConsentRequests(v []MicrosoftGraphAppConsentRequest) {
	o.AppConsentRequests = &v
}

func (o AppConsentApprovalRoute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppConsentRequests != nil {
		toSerialize["appConsentRequests"] = o.AppConsentRequests
	}
	return json.Marshal(toSerialize)
}

type NullableAppConsentApprovalRoute struct {
	value *AppConsentApprovalRoute
	isSet bool
}

func (v NullableAppConsentApprovalRoute) Get() *AppConsentApprovalRoute {
	return v.value
}

func (v *NullableAppConsentApprovalRoute) Set(val *AppConsentApprovalRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableAppConsentApprovalRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableAppConsentApprovalRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppConsentApprovalRoute(val *AppConsentApprovalRoute) *NullableAppConsentApprovalRoute {
	return &NullableAppConsentApprovalRoute{value: val, isSet: true}
}

func (v NullableAppConsentApprovalRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppConsentApprovalRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


