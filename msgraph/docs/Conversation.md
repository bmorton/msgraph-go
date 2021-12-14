# Conversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasAttachments** | Pointer to **bool** | Indicates whether any of the posts within this Conversation has at least one attachment. Supports $filter (eq, ne) and $search. | [optional] 
**LastDeliveredDateTime** | Pointer to **time.Time** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z | [optional] 
**Preview** | Pointer to **string** | A short summary from the body of the latest post in this conversation. Supports $filter (eq, ne, le, ge). | [optional] 
**Topic** | Pointer to **string** | The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. | [optional] 
**UniqueSenders** | Pointer to **[]string** | All the users that sent a message to this Conversation. | [optional] 
**Threads** | Pointer to [**[]MicrosoftGraphConversationThread**](MicrosoftGraphConversationThread.md) | A collection of all the conversation threads in the conversation. A navigation property. Read-only. Nullable. | [optional] 

## Methods

### NewConversation

`func NewConversation() *Conversation`

NewConversation instantiates a new Conversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversationWithDefaults

`func NewConversationWithDefaults() *Conversation`

NewConversationWithDefaults instantiates a new Conversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasAttachments

`func (o *Conversation) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *Conversation) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *Conversation) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *Conversation) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### GetLastDeliveredDateTime

`func (o *Conversation) GetLastDeliveredDateTime() time.Time`

GetLastDeliveredDateTime returns the LastDeliveredDateTime field if non-nil, zero value otherwise.

### GetLastDeliveredDateTimeOk

`func (o *Conversation) GetLastDeliveredDateTimeOk() (*time.Time, bool)`

GetLastDeliveredDateTimeOk returns a tuple with the LastDeliveredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeliveredDateTime

`func (o *Conversation) SetLastDeliveredDateTime(v time.Time)`

SetLastDeliveredDateTime sets LastDeliveredDateTime field to given value.

### HasLastDeliveredDateTime

`func (o *Conversation) HasLastDeliveredDateTime() bool`

HasLastDeliveredDateTime returns a boolean if a field has been set.

### GetPreview

`func (o *Conversation) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *Conversation) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *Conversation) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *Conversation) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetTopic

`func (o *Conversation) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *Conversation) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *Conversation) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *Conversation) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetUniqueSenders

`func (o *Conversation) GetUniqueSenders() []string`

GetUniqueSenders returns the UniqueSenders field if non-nil, zero value otherwise.

### GetUniqueSendersOk

`func (o *Conversation) GetUniqueSendersOk() (*[]string, bool)`

GetUniqueSendersOk returns a tuple with the UniqueSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSenders

`func (o *Conversation) SetUniqueSenders(v []string)`

SetUniqueSenders sets UniqueSenders field to given value.

### HasUniqueSenders

`func (o *Conversation) HasUniqueSenders() bool`

HasUniqueSenders returns a boolean if a field has been set.

### GetThreads

`func (o *Conversation) GetThreads() []MicrosoftGraphConversationThread`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *Conversation) GetThreadsOk() (*[]MicrosoftGraphConversationThread, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *Conversation) SetThreads(v []MicrosoftGraphConversationThread)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *Conversation) HasThreads() bool`

HasThreads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


