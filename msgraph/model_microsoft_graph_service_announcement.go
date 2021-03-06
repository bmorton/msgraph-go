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

// MicrosoftGraphServiceAnnouncement struct for MicrosoftGraphServiceAnnouncement
type MicrosoftGraphServiceAnnouncement struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly.
	HealthOverviews *[]MicrosoftGraphServiceHealth `json:"healthOverviews,omitempty"`
	// A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly.
	Issues *[]MicrosoftGraphServiceHealthIssue `json:"issues,omitempty"`
	// A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly.
	Messages *[]MicrosoftGraphServiceUpdateMessage `json:"messages,omitempty"`
}

// NewMicrosoftGraphServiceAnnouncement instantiates a new MicrosoftGraphServiceAnnouncement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphServiceAnnouncement() *MicrosoftGraphServiceAnnouncement {
	this := MicrosoftGraphServiceAnnouncement{}
	return &this
}

// NewMicrosoftGraphServiceAnnouncementWithDefaults instantiates a new MicrosoftGraphServiceAnnouncement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphServiceAnnouncementWithDefaults() *MicrosoftGraphServiceAnnouncement {
	this := MicrosoftGraphServiceAnnouncement{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceAnnouncement) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceAnnouncement) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceAnnouncement) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphServiceAnnouncement) SetId(v string) {
	o.Id = &v
}

// GetHealthOverviews returns the HealthOverviews field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceAnnouncement) GetHealthOverviews() []MicrosoftGraphServiceHealth {
	if o == nil || o.HealthOverviews == nil {
		var ret []MicrosoftGraphServiceHealth
		return ret
	}
	return *o.HealthOverviews
}

// GetHealthOverviewsOk returns a tuple with the HealthOverviews field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceAnnouncement) GetHealthOverviewsOk() (*[]MicrosoftGraphServiceHealth, bool) {
	if o == nil || o.HealthOverviews == nil {
		return nil, false
	}
	return o.HealthOverviews, true
}

// HasHealthOverviews returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceAnnouncement) HasHealthOverviews() bool {
	if o != nil && o.HealthOverviews != nil {
		return true
	}

	return false
}

// SetHealthOverviews gets a reference to the given []MicrosoftGraphServiceHealth and assigns it to the HealthOverviews field.
func (o *MicrosoftGraphServiceAnnouncement) SetHealthOverviews(v []MicrosoftGraphServiceHealth) {
	o.HealthOverviews = &v
}

// GetIssues returns the Issues field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceAnnouncement) GetIssues() []MicrosoftGraphServiceHealthIssue {
	if o == nil || o.Issues == nil {
		var ret []MicrosoftGraphServiceHealthIssue
		return ret
	}
	return *o.Issues
}

// GetIssuesOk returns a tuple with the Issues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceAnnouncement) GetIssuesOk() (*[]MicrosoftGraphServiceHealthIssue, bool) {
	if o == nil || o.Issues == nil {
		return nil, false
	}
	return o.Issues, true
}

// HasIssues returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceAnnouncement) HasIssues() bool {
	if o != nil && o.Issues != nil {
		return true
	}

	return false
}

// SetIssues gets a reference to the given []MicrosoftGraphServiceHealthIssue and assigns it to the Issues field.
func (o *MicrosoftGraphServiceAnnouncement) SetIssues(v []MicrosoftGraphServiceHealthIssue) {
	o.Issues = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *MicrosoftGraphServiceAnnouncement) GetMessages() []MicrosoftGraphServiceUpdateMessage {
	if o == nil || o.Messages == nil {
		var ret []MicrosoftGraphServiceUpdateMessage
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphServiceAnnouncement) GetMessagesOk() (*[]MicrosoftGraphServiceUpdateMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *MicrosoftGraphServiceAnnouncement) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []MicrosoftGraphServiceUpdateMessage and assigns it to the Messages field.
func (o *MicrosoftGraphServiceAnnouncement) SetMessages(v []MicrosoftGraphServiceUpdateMessage) {
	o.Messages = &v
}

func (o MicrosoftGraphServiceAnnouncement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.HealthOverviews != nil {
		toSerialize["healthOverviews"] = o.HealthOverviews
	}
	if o.Issues != nil {
		toSerialize["issues"] = o.Issues
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphServiceAnnouncement struct {
	value *MicrosoftGraphServiceAnnouncement
	isSet bool
}

func (v NullableMicrosoftGraphServiceAnnouncement) Get() *MicrosoftGraphServiceAnnouncement {
	return v.value
}

func (v *NullableMicrosoftGraphServiceAnnouncement) Set(val *MicrosoftGraphServiceAnnouncement) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphServiceAnnouncement) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphServiceAnnouncement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphServiceAnnouncement(val *MicrosoftGraphServiceAnnouncement) *NullableMicrosoftGraphServiceAnnouncement {
	return &NullableMicrosoftGraphServiceAnnouncement{value: val, isSet: true}
}

func (v NullableMicrosoftGraphServiceAnnouncement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphServiceAnnouncement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


