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

// Channel struct for Channel
type Channel struct {
	// Read only. Timestamp at which the channel was created.
	CreatedDateTime NullableTime `json:"createdDateTime,omitempty"`
	// Optional textual description for the channel.
	Description NullableString `json:"description,omitempty"`
	// Channel name as it will appear to the user in Microsoft Teams.
	DisplayName *string `json:"displayName,omitempty"`
	// The email address for sending messages to the channel. Read-only.
	Email NullableString `json:"email,omitempty"`
	// Indicates whether the channel should automatically be marked 'favorite' for all members of the team. Can only be set programmatically with Create team. Default: false.
	IsFavoriteByDefault NullableBool `json:"isFavoriteByDefault,omitempty"`
	// The type of the channel. Can be set during creation and can't be changed. Possible values are: standard - Channel inherits the list of members of the parent team; private - Channel can have members that are a subset of all the members on the parent team.
	MembershipType AnyOfmicrosoftGraphChannelMembershipType `json:"membershipType,omitempty"`
	// A hyperlink that will go to the channel in Microsoft Teams. This is the URL that you get when you right-click a channel in Microsoft Teams and select Get link to channel. This URL should be treated as an opaque blob, and not parsed. Read-only.
	WebUrl NullableString `json:"webUrl,omitempty"`
	// Metadata for the location where the channel's files are stored.
	FilesFolder AnyOfmicrosoftGraphDriveItem `json:"filesFolder,omitempty"`
	// A collection of membership records associated with the channel.
	Members *[]MicrosoftGraphConversationMember `json:"members,omitempty"`
	// A collection of all the messages in the channel. A navigation property. Nullable.
	Messages *[]MicrosoftGraphChatMessage `json:"messages,omitempty"`
	// A collection of all the tabs in the channel. A navigation property.
	Tabs *[]MicrosoftGraphTeamsTab `json:"tabs,omitempty"`
}

// NewChannel instantiates a new Channel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannel() *Channel {
	this := Channel{}
	return &this
}

// NewChannelWithDefaults instantiates a new Channel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelWithDefaults() *Channel {
	this := Channel{}
	return &this
}

// GetCreatedDateTime returns the CreatedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Channel) GetCreatedDateTime() time.Time {
	if o == nil || o.CreatedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedDateTime.Get()
}

// GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Channel) GetCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreatedDateTime.Get(), o.CreatedDateTime.IsSet()
}

// HasCreatedDateTime returns a boolean if a field has been set.
func (o *Channel) HasCreatedDateTime() bool {
	if o != nil && o.CreatedDateTime.IsSet() {
		return true
	}

	return false
}

// SetCreatedDateTime gets a reference to the given NullableTime and assigns it to the CreatedDateTime field.
func (o *Channel) SetCreatedDateTime(v time.Time) {
	o.CreatedDateTime.Set(&v)
}
// SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil
func (o *Channel) SetCreatedDateTimeNil() {
	o.CreatedDateTime.Set(nil)
}

// UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
func (o *Channel) UnsetCreatedDateTime() {
	o.CreatedDateTime.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Channel) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Channel) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *Channel) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *Channel) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *Channel) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *Channel) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Channel) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Channel) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Channel) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Channel) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Channel) GetEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *Channel) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *Channel) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *Channel) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *Channel) UnsetEmail() {
	o.Email.Unset()
}

// GetIsFavoriteByDefault returns the IsFavoriteByDefault field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Channel) GetIsFavoriteByDefault() bool {
	if o == nil || o.IsFavoriteByDefault.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IsFavoriteByDefault.Get()
}

// GetIsFavoriteByDefaultOk returns a tuple with the IsFavoriteByDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Channel) GetIsFavoriteByDefaultOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IsFavoriteByDefault.Get(), o.IsFavoriteByDefault.IsSet()
}

// HasIsFavoriteByDefault returns a boolean if a field has been set.
func (o *Channel) HasIsFavoriteByDefault() bool {
	if o != nil && o.IsFavoriteByDefault.IsSet() {
		return true
	}

	return false
}

// SetIsFavoriteByDefault gets a reference to the given NullableBool and assigns it to the IsFavoriteByDefault field.
func (o *Channel) SetIsFavoriteByDefault(v bool) {
	o.IsFavoriteByDefault.Set(&v)
}
// SetIsFavoriteByDefaultNil sets the value for IsFavoriteByDefault to be an explicit nil
func (o *Channel) SetIsFavoriteByDefaultNil() {
	o.IsFavoriteByDefault.Set(nil)
}

// UnsetIsFavoriteByDefault ensures that no value is present for IsFavoriteByDefault, not even an explicit nil
func (o *Channel) UnsetIsFavoriteByDefault() {
	o.IsFavoriteByDefault.Unset()
}

// GetMembershipType returns the MembershipType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Channel) GetMembershipType() AnyOfmicrosoftGraphChannelMembershipType {
	if o == nil  {
		var ret AnyOfmicrosoftGraphChannelMembershipType
		return ret
	}
	return o.MembershipType
}

// GetMembershipTypeOk returns a tuple with the MembershipType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Channel) GetMembershipTypeOk() (*AnyOfmicrosoftGraphChannelMembershipType, bool) {
	if o == nil || o.MembershipType == nil {
		return nil, false
	}
	return &o.MembershipType, true
}

// HasMembershipType returns a boolean if a field has been set.
func (o *Channel) HasMembershipType() bool {
	if o != nil && o.MembershipType != nil {
		return true
	}

	return false
}

// SetMembershipType gets a reference to the given AnyOfmicrosoftGraphChannelMembershipType and assigns it to the MembershipType field.
func (o *Channel) SetMembershipType(v AnyOfmicrosoftGraphChannelMembershipType) {
	o.MembershipType = v
}

// GetWebUrl returns the WebUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Channel) GetWebUrl() string {
	if o == nil || o.WebUrl.Get() == nil {
		var ret string
		return ret
	}
	return *o.WebUrl.Get()
}

// GetWebUrlOk returns a tuple with the WebUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Channel) GetWebUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.WebUrl.Get(), o.WebUrl.IsSet()
}

// HasWebUrl returns a boolean if a field has been set.
func (o *Channel) HasWebUrl() bool {
	if o != nil && o.WebUrl.IsSet() {
		return true
	}

	return false
}

// SetWebUrl gets a reference to the given NullableString and assigns it to the WebUrl field.
func (o *Channel) SetWebUrl(v string) {
	o.WebUrl.Set(&v)
}
// SetWebUrlNil sets the value for WebUrl to be an explicit nil
func (o *Channel) SetWebUrlNil() {
	o.WebUrl.Set(nil)
}

// UnsetWebUrl ensures that no value is present for WebUrl, not even an explicit nil
func (o *Channel) UnsetWebUrl() {
	o.WebUrl.Unset()
}

// GetFilesFolder returns the FilesFolder field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Channel) GetFilesFolder() AnyOfmicrosoftGraphDriveItem {
	if o == nil  {
		var ret AnyOfmicrosoftGraphDriveItem
		return ret
	}
	return o.FilesFolder
}

// GetFilesFolderOk returns a tuple with the FilesFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Channel) GetFilesFolderOk() (*AnyOfmicrosoftGraphDriveItem, bool) {
	if o == nil || o.FilesFolder == nil {
		return nil, false
	}
	return &o.FilesFolder, true
}

// HasFilesFolder returns a boolean if a field has been set.
func (o *Channel) HasFilesFolder() bool {
	if o != nil && o.FilesFolder != nil {
		return true
	}

	return false
}

// SetFilesFolder gets a reference to the given AnyOfmicrosoftGraphDriveItem and assigns it to the FilesFolder field.
func (o *Channel) SetFilesFolder(v AnyOfmicrosoftGraphDriveItem) {
	o.FilesFolder = v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *Channel) GetMembers() []MicrosoftGraphConversationMember {
	if o == nil || o.Members == nil {
		var ret []MicrosoftGraphConversationMember
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetMembersOk() (*[]MicrosoftGraphConversationMember, bool) {
	if o == nil || o.Members == nil {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *Channel) HasMembers() bool {
	if o != nil && o.Members != nil {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []MicrosoftGraphConversationMember and assigns it to the Members field.
func (o *Channel) SetMembers(v []MicrosoftGraphConversationMember) {
	o.Members = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *Channel) GetMessages() []MicrosoftGraphChatMessage {
	if o == nil || o.Messages == nil {
		var ret []MicrosoftGraphChatMessage
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetMessagesOk() (*[]MicrosoftGraphChatMessage, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *Channel) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []MicrosoftGraphChatMessage and assigns it to the Messages field.
func (o *Channel) SetMessages(v []MicrosoftGraphChatMessage) {
	o.Messages = &v
}

// GetTabs returns the Tabs field value if set, zero value otherwise.
func (o *Channel) GetTabs() []MicrosoftGraphTeamsTab {
	if o == nil || o.Tabs == nil {
		var ret []MicrosoftGraphTeamsTab
		return ret
	}
	return *o.Tabs
}

// GetTabsOk returns a tuple with the Tabs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Channel) GetTabsOk() (*[]MicrosoftGraphTeamsTab, bool) {
	if o == nil || o.Tabs == nil {
		return nil, false
	}
	return o.Tabs, true
}

// HasTabs returns a boolean if a field has been set.
func (o *Channel) HasTabs() bool {
	if o != nil && o.Tabs != nil {
		return true
	}

	return false
}

// SetTabs gets a reference to the given []MicrosoftGraphTeamsTab and assigns it to the Tabs field.
func (o *Channel) SetTabs(v []MicrosoftGraphTeamsTab) {
	o.Tabs = &v
}

func (o Channel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedDateTime.IsSet() {
		toSerialize["createdDateTime"] = o.CreatedDateTime.Get()
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.IsFavoriteByDefault.IsSet() {
		toSerialize["isFavoriteByDefault"] = o.IsFavoriteByDefault.Get()
	}
	if o.MembershipType != nil {
		toSerialize["membershipType"] = o.MembershipType
	}
	if o.WebUrl.IsSet() {
		toSerialize["webUrl"] = o.WebUrl.Get()
	}
	if o.FilesFolder != nil {
		toSerialize["filesFolder"] = o.FilesFolder
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

type NullableChannel struct {
	value *Channel
	isSet bool
}

func (v NullableChannel) Get() *Channel {
	return v.value
}

func (v *NullableChannel) Set(val *Channel) {
	v.value = val
	v.isSet = true
}

func (v NullableChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannel(val *Channel) *NullableChannel {
	return &NullableChannel{value: val, isSet: true}
}

func (v NullableChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


