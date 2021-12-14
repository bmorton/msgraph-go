/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package msgraph

import (
	"encoding/json"
	"time"
)

// MicrosoftGraphChat struct for MicrosoftGraphChat
type MicrosoftGraphChat struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	// Specifies the type of chat. Possible values are: group, oneOnOne, meeting, unknownFutureValue.
	ChatType AnyOfmicrosoftGraphChatType `json:"chatType,omitempty"`
	// Date and time at which the chat was created. Read-only.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Date and time at which the chat was renamed or list of members were last changed. Read-only.
	LastUpdatedDateTime NullableTime `json:"lastUpdatedDateTime,omitempty"`
	// (Optional) Subject or topic for the chat. Only available for group chats.
	Topic NullableString `json:"topic,omitempty"`
	// A collection of all the apps in the chat. Nullable.
	InstalledApps *[]MicrosoftGraphTeamsAppInstallation `json:"installedApps,omitempty"`
	// A collection of all the members in the chat. Nullable.
	Members *[]MicrosoftGraphConversationMember `json:"members,omitempty"`
	// A collection of all the messages in the chat. Nullable.
	Messages *[]MicrosoftGraphChatMessage `json:"messages,omitempty"`
	Tabs *[]MicrosoftGraphTeamsTab `json:"tabs,omitempty"`
}

// NewMicrosoftGraphChat instantiates a new MicrosoftGraphChat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphChat() *MicrosoftGraphChat {
	this := MicrosoftGraphChat{}
	return &this
}

// NewMicrosoftGraphChatWithDefaults instantiates a new MicrosoftGraphChat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphChatWithDefaults() *MicrosoftGraphChat {
	this := MicrosoftGraphChat{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphChat) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChat) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphChat) SetId(v string) {
	o.Id = &v
}

// GetChatType returns the ChatType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChat) GetChatType() AnyOfmicrosoftGraphChatType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphChatType
		return ret
	}
	return o.ChatType
}

// GetChatTypeOk returns a tuple with the ChatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChat) GetChatTypeOk() (*AnyOfmicrosoftGraphChatType, bool) {
	if o == nil || o.ChatType == nil {
		return nil, false
	}
	return &o.ChatType, true
}

// HasChatType returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasChatType() bool {
	if o != nil && o.ChatType != nil {
		return true
	}

	return false
}

// SetChatType gets a reference to the given AnyOfmicrosoftGraphChatType and assigns it to the ChatType field.
func (o *MicrosoftGraphChat) SetChatType(v AnyOfmicrosoftGraphChatType) {
	o.ChatType = v
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChat) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChat) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *MicrosoftGraphChat) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *MicrosoftGraphChat) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *MicrosoftGraphChat) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetLastUpdatedDateTime returns the LastUpdatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChat) GetLastUpdatedDateTime() time.Time {
	if o == nil || o.LastUpdatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastUpdatedDateTime.Get()
}

// GetLastUpdatedDateTimeOk returns a tuple with the LastUpdatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChat) GetLastUpdatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastUpdatedDateTime.Get(), o.LastUpdatedDateTime.IsSet()
}

// HasLastUpdatedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasLastUpdatedDateTime() bool {
	if o != nil && o.LastUpdatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastUpdatedDateTime gets a reference to the given NullableTime and assigns it to the LastUpdatedDateTime field.
func (o *MicrosoftGraphChat) SetLastUpdatedDateTime(v time.Time) {
	o.LastUpdatedDateTime.Set(&v)
}
// SetLastUpdatedDateTimeNil sets the value for LastUpdatedDateTime to be an explicit nil
func (o *MicrosoftGraphChat) SetLastUpdatedDateTimeNil() {
	o.LastUpdatedDateTime.Set(nil)
}

// UnsetLastUpdatedDateTime ensures that no value is present for LastUpdatedDateTime, not even an explicit nil
func (o *MicrosoftGraphChat) UnsetLastUpdatedDateTime() {
	o.LastUpdatedDateTime.Unset()
}

// GetTopic returns the Topic field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphChat) GetTopic() string {
	if o == nil || o.Topic.Get() == nil {
		var ret string
		return ret
	}
	return *o.Topic.Get()
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphChat) GetTopicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Topic.Get(), o.Topic.IsSet()
}

// HasTopic returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasTopic() bool {
	if o != nil && o.Topic.IsSet() {
		return true
	}

	return false
}

// SetTopic gets a reference to the given NullableString and assigns it to the Topic field.
func (o *MicrosoftGraphChat) SetTopic(v string) {
	o.Topic.Set(&v)
}
// SetTopicNil sets the value for Topic to be an explicit nil
func (o *MicrosoftGraphChat) SetTopicNil() {
	o.Topic.Set(nil)
}

// UnsetTopic ensures that no value is present for Topic, not even an explicit nil
func (o *MicrosoftGraphChat) UnsetTopic() {
	o.Topic.Unset()
}

// GetInstalledApps returns the InstalledApps field value if set, zero value otherwise.
func (o *MicrosoftGraphChat) GetInstalledApps() []MicrosoftGraphTeamsAppInstallation {
	if o == nil || o.InstalledApps == nil {
		var ret []MicrosoftGraphTeamsAppInstallation
		return ret
	}
	return *o.InstalledApps
}

// GetInstalledAppsOk returns a tuple with the InstalledApps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChat) GetInstalledAppsOk() (*[]MicrosoftGraphTeamsAppInstallation, bool) {
	if o == nil || o.InstalledApps == nil {
		return nil, false
	}
	return o.InstalledApps, true
}

// HasInstalledApps returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasInstalledApps() bool {
	if o != nil && o.InstalledApps != nil {
		return true
	}

	return false
}

// SetInstalledApps gets a reference to the given []MicrosoftGraphTeamsAppInstallation and assigns it to the InstalledApps field.
func (o *MicrosoftGraphChat) SetInstalledApps(v []MicrosoftGraphTeamsAppInstallation) {
	o.InstalledApps = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *MicrosoftGraphChat) GetMembers() []MicrosoftGraphConversationMember {
	if o == nil || o.Members == nil {
		var ret []MicrosoftGraphConversationMember
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChat) GetMembersOk() (*[]MicrosoftGraphConversationMember, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []MicrosoftGraphConversationMember and assigns it to the Members field.
func (o *MicrosoftGraphChat) SetMembers(v []MicrosoftGraphConversationMember) {
	o.Members = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *MicrosoftGraphChat) GetMessages() []MicrosoftGraphChatMessage {
	if o == nil || o.Messages == nil {
		var ret []MicrosoftGraphChatMessage
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChat) GetMessagesOk() (*[]MicrosoftGraphChatMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []MicrosoftGraphChatMessage and assigns it to the Messages field.
func (o *MicrosoftGraphChat) SetMessages(v []MicrosoftGraphChatMessage) {
	o.Messages = &v
}

// GetTabs returns the Tabs field value if set, zero value otherwise.
func (o *MicrosoftGraphChat) GetTabs() []MicrosoftGraphTeamsTab {
	if o == nil || o.Tabs == nil {
		var ret []MicrosoftGraphTeamsTab
		return ret
	}
	return *o.Tabs
}

// GetTabsOk returns a tuple with the Tabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphChat) GetTabsOk() (*[]MicrosoftGraphTeamsTab, bool) {
	if o == nil || o.Tabs == nil {
		return nil, false
	}
	return o.Tabs, true
}

// HasTabs returns a boolean if a field has been set.
func (o *MicrosoftGraphChat) HasTabs() bool {
	if o != nil && o.Tabs != nil {
		return true
	}

	return false
}

// SetTabs gets a reference to the given []MicrosoftGraphTeamsTab and assigns it to the Tabs field.
func (o *MicrosoftGraphChat) SetTabs(v []MicrosoftGraphTeamsTab) {
	o.Tabs = &v
}

func (o MicrosoftGraphChat) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ChatType != nil {
		toSerialize["chatType"] = o.ChatType
	}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.LastUpdatedDateTime.IsSet() {
		toSerialize["lastUpdatedDateTime"] = o.LastUpdatedDateTime.Get()
	}
	if o.Topic.IsSet() {
		toSerialize["topic"] = o.Topic.Get()
	}
	if o.InstalledApps != nil {
		toSerialize["installedApps"] = o.InstalledApps
	}
	if o.Members != nil {
		toSerialize["members"] = o.Members
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Tabs != nil {
		toSerialize["tabs"] = o.Tabs
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphChat struct {
	value *MicrosoftGraphChat
	isSet bool
}

func (v NullableMicrosoftGraphChat) Get() *MicrosoftGraphChat {
	return v.value
}

func (v *NullableMicrosoftGraphChat) Set(val *MicrosoftGraphChat) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphChat) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphChat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphChat(val *MicrosoftGraphChat) *NullableMicrosoftGraphChat {
	return &NullableMicrosoftGraphChat{value: val, isSet: true}
}

func (v NullableMicrosoftGraphChat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphChat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


