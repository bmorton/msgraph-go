/*
Partial Graph API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// ConversationThread struct for ConversationThread
type ConversationThread struct {
	// The Cc: recipients for the thread. Returned only on $select.
	CcRecipients *[]MicrosoftGraphRecipient `json:"ccRecipients,omitempty"`
	// Indicates whether any of the posts within this thread has at least one attachment. Returned by default.
	HasAttachments *bool `json:"hasAttachments,omitempty"`
	// Indicates if the thread is locked. Returned by default.
	IsLocked *bool `json:"isLocked,omitempty"`
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default.
	LastDeliveredDateTime *time.Time `json:"lastDeliveredDateTime,omitempty"`
	// A short summary from the body of the latest post in this conversation. Returned by default.
	Preview *string `json:"preview,omitempty"`
	// The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. Returned by default.
	Topic *string `json:"topic,omitempty"`
	// The To: recipients for the thread. Returned only on $select.
	ToRecipients *[]MicrosoftGraphRecipient `json:"toRecipients,omitempty"`
	// All the users that sent a message to this thread. Returned by default.
	UniqueSenders *[]string `json:"uniqueSenders,omitempty"`
	// Read-only. Nullable.
	Posts *[]MicrosoftGraphPost `json:"posts,omitempty"`
}

// NewConversationThread instantiates a new ConversationThread object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConversationThread() *ConversationThread {
	this := ConversationThread{}
	return &this
}

// NewConversationThreadWithDefaults instantiates a new ConversationThread object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConversationThreadWithDefaults() *ConversationThread {
	this := ConversationThread{}
	return &this
}

// GetCcRecipients returns the CcRecipients field value if set, zero value otherwise.
func (o *ConversationThread) GetCcRecipients() []MicrosoftGraphRecipient {
	if o == nil || o.CcRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret
	}
	return *o.CcRecipients
}

// GetCcRecipientsOk returns a tuple with the CcRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetCcRecipientsOk() (*[]MicrosoftGraphRecipient, bool) {
	if o == nil || o.CcRecipients == nil {
		return nil, false
	}
	return o.CcRecipients, true
}

// HasCcRecipients returns a boolean if a field has been set.
func (o *ConversationThread) HasCcRecipients() bool {
	if o != nil && o.CcRecipients != nil {
		return true
	}

	return false
}

// SetCcRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the CcRecipients field.
func (o *ConversationThread) SetCcRecipients(v []MicrosoftGraphRecipient) {
	o.CcRecipients = &v
}

// GetHasAttachments returns the HasAttachments field value if set, zero value otherwise.
func (o *ConversationThread) GetHasAttachments() bool {
	if o == nil || o.HasAttachments == nil {
		var ret bool
		return ret
	}
	return *o.HasAttachments
}

// GetHasAttachmentsOk returns a tuple with the HasAttachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetHasAttachmentsOk() (*bool, bool) {
	if o == nil || o.HasAttachments == nil {
		return nil, false
	}
	return o.HasAttachments, true
}

// HasHasAttachments returns a boolean if a field has been set.
func (o *ConversationThread) HasHasAttachments() bool {
	if o != nil && o.HasAttachments != nil {
		return true
	}

	return false
}

// SetHasAttachments gets a reference to the given bool and assigns it to the HasAttachments field.
func (o *ConversationThread) SetHasAttachments(v bool) {
	o.HasAttachments = &v
}

// GetIsLocked returns the IsLocked field value if set, zero value otherwise.
func (o *ConversationThread) GetIsLocked() bool {
	if o == nil || o.IsLocked == nil {
		var ret bool
		return ret
	}
	return *o.IsLocked
}

// GetIsLockedOk returns a tuple with the IsLocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetIsLockedOk() (*bool, bool) {
	if o == nil || o.IsLocked == nil {
		return nil, false
	}
	return o.IsLocked, true
}

// HasIsLocked returns a boolean if a field has been set.
func (o *ConversationThread) HasIsLocked() bool {
	if o != nil && o.IsLocked != nil {
		return true
	}

	return false
}

// SetIsLocked gets a reference to the given bool and assigns it to the IsLocked field.
func (o *ConversationThread) SetIsLocked(v bool) {
	o.IsLocked = &v
}

// GetLastDeliveredDateTime returns the LastDeliveredDateTime field value if set, zero value otherwise.
func (o *ConversationThread) GetLastDeliveredDateTime() time.Time {
	if o == nil || o.LastDeliveredDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDeliveredDateTime
}

// GetLastDeliveredDateTimeOk returns a tuple with the LastDeliveredDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetLastDeliveredDateTimeOk() (*time.Time, bool) {
	if o == nil || o.LastDeliveredDateTime == nil {
		return nil, false
	}
	return o.LastDeliveredDateTime, true
}

// HasLastDeliveredDateTime returns a boolean if a field has been set.
func (o *ConversationThread) HasLastDeliveredDateTime() bool {
	if o != nil && o.LastDeliveredDateTime != nil {
		return true
	}

	return false
}

// SetLastDeliveredDateTime gets a reference to the given time.Time and assigns it to the LastDeliveredDateTime field.
func (o *ConversationThread) SetLastDeliveredDateTime(v time.Time) {
	o.LastDeliveredDateTime = &v
}

// GetPreview returns the Preview field value if set, zero value otherwise.
func (o *ConversationThread) GetPreview() string {
	if o == nil || o.Preview == nil {
		var ret string
		return ret
	}
	return *o.Preview
}

// GetPreviewOk returns a tuple with the Preview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetPreviewOk() (*string, bool) {
	if o == nil || o.Preview == nil {
		return nil, false
	}
	return o.Preview, true
}

// HasPreview returns a boolean if a field has been set.
func (o *ConversationThread) HasPreview() bool {
	if o != nil && o.Preview != nil {
		return true
	}

	return false
}

// SetPreview gets a reference to the given string and assigns it to the Preview field.
func (o *ConversationThread) SetPreview(v string) {
	o.Preview = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *ConversationThread) GetTopic() string {
	if o == nil || o.Topic == nil {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetTopicOk() (*string, bool) {
	if o == nil || o.Topic == nil {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ConversationThread) HasTopic() bool {
	if o != nil && o.Topic != nil {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *ConversationThread) SetTopic(v string) {
	o.Topic = &v
}

// GetToRecipients returns the ToRecipients field value if set, zero value otherwise.
func (o *ConversationThread) GetToRecipients() []MicrosoftGraphRecipient {
	if o == nil || o.ToRecipients == nil {
		var ret []MicrosoftGraphRecipient
		return ret
	}
	return *o.ToRecipients
}

// GetToRecipientsOk returns a tuple with the ToRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetToRecipientsOk() (*[]MicrosoftGraphRecipient, bool) {
	if o == nil || o.ToRecipients == nil {
		return nil, false
	}
	return o.ToRecipients, true
}

// HasToRecipients returns a boolean if a field has been set.
func (o *ConversationThread) HasToRecipients() bool {
	if o != nil && o.ToRecipients != nil {
		return true
	}

	return false
}

// SetToRecipients gets a reference to the given []MicrosoftGraphRecipient and assigns it to the ToRecipients field.
func (o *ConversationThread) SetToRecipients(v []MicrosoftGraphRecipient) {
	o.ToRecipients = &v
}

// GetUniqueSenders returns the UniqueSenders field value if set, zero value otherwise.
func (o *ConversationThread) GetUniqueSenders() []string {
	if o == nil || o.UniqueSenders == nil {
		var ret []string
		return ret
	}
	return *o.UniqueSenders
}

// GetUniqueSendersOk returns a tuple with the UniqueSenders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetUniqueSendersOk() (*[]string, bool) {
	if o == nil || o.UniqueSenders == nil {
		return nil, false
	}
	return o.UniqueSenders, true
}

// HasUniqueSenders returns a boolean if a field has been set.
func (o *ConversationThread) HasUniqueSenders() bool {
	if o != nil && o.UniqueSenders != nil {
		return true
	}

	return false
}

// SetUniqueSenders gets a reference to the given []string and assigns it to the UniqueSenders field.
func (o *ConversationThread) SetUniqueSenders(v []string) {
	o.UniqueSenders = &v
}

// GetPosts returns the Posts field value if set, zero value otherwise.
func (o *ConversationThread) GetPosts() []MicrosoftGraphPost {
	if o == nil || o.Posts == nil {
		var ret []MicrosoftGraphPost
		return ret
	}
	return *o.Posts
}

// GetPostsOk returns a tuple with the Posts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConversationThread) GetPostsOk() (*[]MicrosoftGraphPost, bool) {
	if o == nil || o.Posts == nil {
		return nil, false
	}
	return o.Posts, true
}

// HasPosts returns a boolean if a field has been set.
func (o *ConversationThread) HasPosts() bool {
	if o != nil && o.Posts != nil {
		return true
	}

	return false
}

// SetPosts gets a reference to the given []MicrosoftGraphPost and assigns it to the Posts field.
func (o *ConversationThread) SetPosts(v []MicrosoftGraphPost) {
	o.Posts = &v
}

func (o ConversationThread) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CcRecipients != nil {
		toSerialize["ccRecipients"] = o.CcRecipients
	}
	if o.HasAttachments != nil {
		toSerialize["hasAttachments"] = o.HasAttachments
	}
	if o.IsLocked != nil {
		toSerialize["isLocked"] = o.IsLocked
	}
	if o.LastDeliveredDateTime != nil {
		toSerialize["lastDeliveredDateTime"] = o.LastDeliveredDateTime
	}
	if o.Preview != nil {
		toSerialize["preview"] = o.Preview
	}
	if o.Topic != nil {
		toSerialize["topic"] = o.Topic
	}
	if o.ToRecipients != nil {
		toSerialize["toRecipients"] = o.ToRecipients
	}
	if o.UniqueSenders != nil {
		toSerialize["uniqueSenders"] = o.UniqueSenders
	}
	if o.Posts != nil {
		toSerialize["posts"] = o.Posts
	}
	return json.Marshal(toSerialize)
}

type NullableConversationThread struct {
	value *ConversationThread
	isSet bool
}

func (v NullableConversationThread) Get() *ConversationThread {
	return v.value
}

func (v *NullableConversationThread) Set(val *ConversationThread) {
	v.value = val
	v.isSet = true
}

func (v NullableConversationThread) IsSet() bool {
	return v.isSet
}

func (v *NullableConversationThread) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConversationThread(val *ConversationThread) *NullableConversationThread {
	return &NullableConversationThread{value: val, isSet: true}
}

func (v NullableConversationThread) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConversationThread) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


