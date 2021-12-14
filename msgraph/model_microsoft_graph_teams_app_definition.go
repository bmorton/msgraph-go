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

// MicrosoftGraphTeamsAppDefinition struct for MicrosoftGraphTeamsAppDefinition
type MicrosoftGraphTeamsAppDefinition struct {
	// Read-only.
	Id *string `json:"id,omitempty"`
	CreatedBy AnyOfmicrosoftGraphIdentitySet `json:"createdBy,omitempty"`
	// Verbose description of the application.
	Description NullableString `json:"description,omitempty"`
	// The name of the app provided by the app developer.
	DisplayName NullableString `json:"displayName,omitempty"`
	LastModifiedDateTime NullableTime `json:"lastModifiedDateTime,omitempty"`
	// The published status of a specific version of a Teams app. Possible values are:submitted — The specific version of the Teams app has been submitted and is under review. published  — The request to publish the specific version of the Teams app has been approved by the admin and the app is published.  rejected — The request to publish the specific version of the Teams app was rejected by the admin.
	PublishingState AnyOfmicrosoftGraphTeamsAppPublishingState `json:"publishingState,omitempty"`
	// Short description of the application.
	ShortDescription NullableString `json:"shortDescription,omitempty"`
	// The ID from the Teams app manifest.
	TeamsAppId NullableString `json:"teamsAppId,omitempty"`
	// The version number of the application.
	Version NullableString `json:"version,omitempty"`
	// The details of the bot specified in the Teams app manifest.
	Bot AnyOfmicrosoftGraphTeamworkBot `json:"bot,omitempty"`
}

// NewMicrosoftGraphTeamsAppDefinition instantiates a new MicrosoftGraphTeamsAppDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMicrosoftGraphTeamsAppDefinition() *MicrosoftGraphTeamsAppDefinition {
	this := MicrosoftGraphTeamsAppDefinition{}
	return &this
}

// NewMicrosoftGraphTeamsAppDefinitionWithDefaults instantiates a new MicrosoftGraphTeamsAppDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMicrosoftGraphTeamsAppDefinitionWithDefaults() *MicrosoftGraphTeamsAppDefinition {
	this := MicrosoftGraphTeamsAppDefinition{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MicrosoftGraphTeamsAppDefinition) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MicrosoftGraphTeamsAppDefinition) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MicrosoftGraphTeamsAppDefinition) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet {
	if o == nil  {
		var ret AnyOfmicrosoftGraphIdentitySet
		return ret
	}
	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given AnyOfmicrosoftGraphIdentitySet and assigns it to the CreatedBy field.
func (o *MicrosoftGraphTeamsAppDefinition) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet) {
	o.CreatedBy = v
}

// GetDescription returns the Description field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}
	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// HasDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasDescription() bool {
	if o != nil && o.Description.IsSet() {
		return true
	}

	return false
}

// SetDescription gets a reference to the given NullableString and assigns it to the Description field.
func (o *MicrosoftGraphTeamsAppDefinition) SetDescription(v string) {
	o.Description.Set(&v)
}
// SetDescriptionNil sets the value for Description to be an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) SetDescriptionNil() {
	o.Description.Set(nil)
}

// UnsetDescription ensures that no value is present for Description, not even an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) UnsetDescription() {
	o.Description.Unset()
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetDisplayName() string {
	if o == nil || o.DisplayName.Get() == nil {
		var ret string
		return ret
	}
	return *o.DisplayName.Get()
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetDisplayNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DisplayName.Get(), o.DisplayName.IsSet()
}

// HasDisplayName returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasDisplayName() bool {
	if o != nil && o.DisplayName.IsSet() {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given NullableString and assigns it to the DisplayName field.
func (o *MicrosoftGraphTeamsAppDefinition) SetDisplayName(v string) {
	o.DisplayName.Set(&v)
}
// SetDisplayNameNil sets the value for DisplayName to be an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) SetDisplayNameNil() {
	o.DisplayName.Set(nil)
}

// UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) UnsetDisplayName() {
	o.DisplayName.Unset()
}

// GetLastModifiedDateTime returns the LastModifiedDateTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetLastModifiedDateTime() time.Time {
	if o == nil || o.LastModifiedDateTime.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedDateTime.Get()
}

// GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetLastModifiedDateTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.LastModifiedDateTime.Get(), o.LastModifiedDateTime.IsSet()
}

// HasLastModifiedDateTime returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasLastModifiedDateTime() bool {
	if o != nil && o.LastModifiedDateTime.IsSet() {
		return true
	}

	return false
}

// SetLastModifiedDateTime gets a reference to the given NullableTime and assigns it to the LastModifiedDateTime field.
func (o *MicrosoftGraphTeamsAppDefinition) SetLastModifiedDateTime(v time.Time) {
	o.LastModifiedDateTime.Set(&v)
}
// SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) SetLastModifiedDateTimeNil() {
	o.LastModifiedDateTime.Set(nil)
}

// UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) UnsetLastModifiedDateTime() {
	o.LastModifiedDateTime.Unset()
}

// GetPublishingState returns the PublishingState field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetPublishingState() AnyOfmicrosoftGraphTeamsAppPublishingState {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamsAppPublishingState
		return ret
	}
	return o.PublishingState
}

// GetPublishingStateOk returns a tuple with the PublishingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetPublishingStateOk() (*AnyOfmicrosoftGraphTeamsAppPublishingState, bool) {
	if o == nil || o.PublishingState == nil {
		return nil, false
	}
	return &o.PublishingState, true
}

// HasPublishingState returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasPublishingState() bool {
	if o != nil && o.PublishingState != nil {
		return true
	}

	return false
}

// SetPublishingState gets a reference to the given AnyOfmicrosoftGraphTeamsAppPublishingState and assigns it to the PublishingState field.
func (o *MicrosoftGraphTeamsAppDefinition) SetPublishingState(v AnyOfmicrosoftGraphTeamsAppPublishingState) {
	o.PublishingState = v
}

// GetShortDescription returns the ShortDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetShortDescription() string {
	if o == nil || o.ShortDescription.Get() == nil {
		var ret string
		return ret
	}
	return *o.ShortDescription.Get()
}

// GetShortDescriptionOk returns a tuple with the ShortDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetShortDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ShortDescription.Get(), o.ShortDescription.IsSet()
}

// HasShortDescription returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasShortDescription() bool {
	if o != nil && o.ShortDescription.IsSet() {
		return true
	}

	return false
}

// SetShortDescription gets a reference to the given NullableString and assigns it to the ShortDescription field.
func (o *MicrosoftGraphTeamsAppDefinition) SetShortDescription(v string) {
	o.ShortDescription.Set(&v)
}
// SetShortDescriptionNil sets the value for ShortDescription to be an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) SetShortDescriptionNil() {
	o.ShortDescription.Set(nil)
}

// UnsetShortDescription ensures that no value is present for ShortDescription, not even an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) UnsetShortDescription() {
	o.ShortDescription.Unset()
}

// GetTeamsAppId returns the TeamsAppId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetTeamsAppId() string {
	if o == nil || o.TeamsAppId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TeamsAppId.Get()
}

// GetTeamsAppIdOk returns a tuple with the TeamsAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetTeamsAppIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TeamsAppId.Get(), o.TeamsAppId.IsSet()
}

// HasTeamsAppId returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasTeamsAppId() bool {
	if o != nil && o.TeamsAppId.IsSet() {
		return true
	}

	return false
}

// SetTeamsAppId gets a reference to the given NullableString and assigns it to the TeamsAppId field.
func (o *MicrosoftGraphTeamsAppDefinition) SetTeamsAppId(v string) {
	o.TeamsAppId.Set(&v)
}
// SetTeamsAppIdNil sets the value for TeamsAppId to be an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) SetTeamsAppIdNil() {
	o.TeamsAppId.Set(nil)
}

// UnsetTeamsAppId ensures that no value is present for TeamsAppId, not even an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) UnsetTeamsAppId() {
	o.TeamsAppId.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetVersion() string {
	if o == nil || o.Version.Get() == nil {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *MicrosoftGraphTeamsAppDefinition) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *MicrosoftGraphTeamsAppDefinition) UnsetVersion() {
	o.Version.Unset()
}

// GetBot returns the Bot field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MicrosoftGraphTeamsAppDefinition) GetBot() AnyOfmicrosoftGraphTeamworkBot {
	if o == nil  {
		var ret AnyOfmicrosoftGraphTeamworkBot
		return ret
	}
	return o.Bot
}

// GetBotOk returns a tuple with the Bot field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MicrosoftGraphTeamsAppDefinition) GetBotOk() (*AnyOfmicrosoftGraphTeamworkBot, bool) {
	if o == nil || o.Bot == nil {
		return nil, false
	}
	return &o.Bot, true
}

// HasBot returns a boolean if a field has been set.
func (o *MicrosoftGraphTeamsAppDefinition) HasBot() bool {
	if o != nil && o.Bot != nil {
		return true
	}

	return false
}

// SetBot gets a reference to the given AnyOfmicrosoftGraphTeamworkBot and assigns it to the Bot field.
func (o *MicrosoftGraphTeamsAppDefinition) SetBot(v AnyOfmicrosoftGraphTeamworkBot) {
	o.Bot = v
}

func (o MicrosoftGraphTeamsAppDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.Description.IsSet() {
		toSerialize["description"] = o.Description.Get()
	}
	if o.DisplayName.IsSet() {
		toSerialize["displayName"] = o.DisplayName.Get()
	}
	if o.LastModifiedDateTime.IsSet() {
		toSerialize["lastModifiedDateTime"] = o.LastModifiedDateTime.Get()
	}
	if o.PublishingState != nil {
		toSerialize["publishingState"] = o.PublishingState
	}
	if o.ShortDescription.IsSet() {
		toSerialize["shortDescription"] = o.ShortDescription.Get()
	}
	if o.TeamsAppId.IsSet() {
		toSerialize["teamsAppId"] = o.TeamsAppId.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if o.Bot != nil {
		toSerialize["bot"] = o.Bot
	}
	return json.Marshal(toSerialize)
}

type NullableMicrosoftGraphTeamsAppDefinition struct {
	value *MicrosoftGraphTeamsAppDefinition
	isSet bool
}

func (v NullableMicrosoftGraphTeamsAppDefinition) Get() *MicrosoftGraphTeamsAppDefinition {
	return v.value
}

func (v *NullableMicrosoftGraphTeamsAppDefinition) Set(val *MicrosoftGraphTeamsAppDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableMicrosoftGraphTeamsAppDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableMicrosoftGraphTeamsAppDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMicrosoftGraphTeamsAppDefinition(val *MicrosoftGraphTeamsAppDefinition) *NullableMicrosoftGraphTeamsAppDefinition {
	return &NullableMicrosoftGraphTeamsAppDefinition{value: val, isSet: true}
}

func (v NullableMicrosoftGraphTeamsAppDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMicrosoftGraphTeamsAppDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


