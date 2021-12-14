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

// MicrosoftGraphAppConsentApprovalRoute struct for MicrosoftGraphAppConsentApprovalRoute
type MicrosoftGraphAppConsentApprovalRoute struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	AppConsentRequests *[]MicrosoftGraphAppConsentRequest `json:"appConsentRequests,omitempty"`
}

// NewMicrosoftGraphAppConsentApprovalRoute instantiates a new MicrosoftGraphAppConsentApprovalRoute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphAppConsentApprovalRoute() *MicrosoftGraphAppConsentApprovalRoute {
	this := MicrosoftGraphAppConsentApprovalRoute{}
	return &this
}

// NewMicrosoftGraphAppConsentApprovalRouteWithDefaults instantiates a new MicrosoftGraphAppConsentApprovalRoute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphAppConsentApprovalRouteWithDefaults() *MicrosoftGraphAppConsentApprovalRoute {
	this := MicrosoftGraphAppConsentApprovalRoute{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphAppConsentApprovalRoute) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppConsentApprovalRoute) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphAppConsentApprovalRoute) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphAppConsentApprovalRoute) SetId(v string) {
	o.Id = &v
}

// GetAppConsentRequests returns the AppConsentRequests field value if set, zero value otherwise.
func (o *MicrosoftGraphAppConsentApprovalRoute) GetAppConsentRequests() []MicrosoftGraphAppConsentRequest {
	if o == nil || o.AppConsentRequests == nil {
		var ret []MicrosoftGraphAppConsentRequest
		return ret
	}
	return *o.AppConsentRequests
}

// GetAppConsentRequestsOk returns a tuple with the AppConsentRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphAppConsentApprovalRoute) GetAppConsentRequestsOk() (*[]MicrosoftGraphAppConsentRequest, bool) {
	if o == nil || o.AppConsentRequests == nil {
		return nil, false
	}
	return o.AppConsentRequests, true
}

// HasAppConsentRequests returns a boolean if a field has been set.
func (o *MicrosoftGraphAppConsentApprovalRoute) HasAppConsentRequests() bool {
	if o != nil && o.AppConsentRequests != nil {
		return true
	}

	return false
}

// SetAppConsentRequests gets a reference to the given []MicrosoftGraphAppConsentRequest and assigns it to the AppConsentRequests field.
func (o *MicrosoftGraphAppConsentApprovalRoute) SetAppConsentRequests(v []MicrosoftGraphAppConsentRequest) {
	o.AppConsentRequests = &v
}

func (o MicrosoftGraphAppConsentApprovalRoute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.AppConsentRequests != nil {
		toSerialize["appConsentRequests"] = o.AppConsentRequests
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphAppConsentApprovalRoute struct {
	value *MicrosoftGraphAppConsentApprovalRoute
	isSet bool
}

func (v NullableMicrosoftGraphAppConsentApprovalRoute) Get() *MicrosoftGraphAppConsentApprovalRoute {
	return v.value
}

func (v *NullableMicrosoftGraphAppConsentApprovalRoute) Set(val *MicrosoftGraphAppConsentApprovalRoute) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphAppConsentApprovalRoute) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphAppConsentApprovalRoute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphAppConsentApprovalRoute(val *MicrosoftGraphAppConsentApprovalRoute) *NullableMicrosoftGraphAppConsentApprovalRoute {
	return &NullableMicrosoftGraphAppConsentApprovalRoute{value: val, isSet: true}
}

func (v NullableMicrosoftGraphAppConsentApprovalRoute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphAppConsentApprovalRoute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


