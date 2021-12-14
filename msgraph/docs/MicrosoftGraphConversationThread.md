# MicrosoftGraphConversationThread

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**CcRecipients** | Pointer to [**[]MicrosoftGraphRecipient**](MicrosoftGraphRecipient.md) | The Cc: recipients for the thread. Returned only on $select. | [optional] 
**HasAttachments** | Pointer to **bool** | Indicates whether any of the posts within this thread has at least one attachment. Returned by default. | [optional] 
**IsLocked** | Pointer to **bool** | Indicates if the thread is locked. Returned by default. | [optional] 
**LastDeliveredDateTime** | Pointer to **time.Time** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.Returned by default. | [optional] 
**Preview** | Pointer to **string** | A short summary from the body of the latest post in this conversation. Returned by default. | [optional] 
**Topic** | Pointer to **string** | The topic of the conversation. This property can be set when the conversation is created, but it cannot be updated. Returned by default. | [optional] 
**ToRecipients** | Pointer to [**[]MicrosoftGraphRecipient**](MicrosoftGraphRecipient.md) | The To: recipients for the thread. Returned only on $select. | [optional] 
**UniqueSenders** | Pointer to **[]string** | All the users that sent a message to this thread. Returned by default. | [optional] 
**Posts** | Pointer to [**[]MicrosoftGraphPost**](MicrosoftGraphPost.md) | Read-only. Nullable. | [optional] 

## Methods

### NewMicrosoftGraphConversationThread

`func NewMicrosoftGraphConversationThread() *MicrosoftGraphConversationThread`

NewMicrosoftGraphConversationThread instantiates a new MicrosoftGraphConversationThread object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConversationThreadWithDefaults

`func NewMicrosoftGraphConversationThreadWithDefaults() *MicrosoftGraphConversationThread`

NewMicrosoftGraphConversationThreadWithDefaults instantiates a new MicrosoftGraphConversationThread object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphConversationThread) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphConversationThread) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphConversationThread) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphConversationThread) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCcRecipients

`func (o *MicrosoftGraphConversationThread) GetCcRecipients() []MicrosoftGraphRecipient`

GetCcRecipients returns the CcRecipients field if non-nil, zero value otherwise.

### GetCcRecipientsOk

`func (o *MicrosoftGraphConversationThread) GetCcRecipientsOk() (*[]MicrosoftGraphRecipient, bool)`

GetCcRecipientsOk returns a tuple with the CcRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCcRecipients

`func (o *MicrosoftGraphConversationThread) SetCcRecipients(v []MicrosoftGraphRecipient)`

SetCcRecipients sets CcRecipients field to given value.

### HasCcRecipients

`func (o *MicrosoftGraphConversationThread) HasCcRecipients() bool`

HasCcRecipients returns a boolean if a field has been set.

### GetHasAttachments

`func (o *MicrosoftGraphConversationThread) GetHasAttachments() bool`

GetHasAttachments returns the HasAttachments field if non-nil, zero value otherwise.

### GetHasAttachmentsOk

`func (o *MicrosoftGraphConversationThread) GetHasAttachmentsOk() (*bool, bool)`

GetHasAttachmentsOk returns a tuple with the HasAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAttachments

`func (o *MicrosoftGraphConversationThread) SetHasAttachments(v bool)`

SetHasAttachments sets HasAttachments field to given value.

### HasHasAttachments

`func (o *MicrosoftGraphConversationThread) HasHasAttachments() bool`

HasHasAttachments returns a boolean if a field has been set.

### GetIsLocked

`func (o *MicrosoftGraphConversationThread) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *MicrosoftGraphConversationThread) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *MicrosoftGraphConversationThread) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *MicrosoftGraphConversationThread) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetLastDeliveredDateTime

`func (o *MicrosoftGraphConversationThread) GetLastDeliveredDateTime() time.Time`

GetLastDeliveredDateTime returns the LastDeliveredDateTime field if non-nil, zero value otherwise.

### GetLastDeliveredDateTimeOk

`func (o *MicrosoftGraphConversationThread) GetLastDeliveredDateTimeOk() (*time.Time, bool)`

GetLastDeliveredDateTimeOk returns a tuple with the LastDeliveredDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeliveredDateTime

`func (o *MicrosoftGraphConversationThread) SetLastDeliveredDateTime(v time.Time)`

SetLastDeliveredDateTime sets LastDeliveredDateTime field to given value.

### HasLastDeliveredDateTime

`func (o *MicrosoftGraphConversationThread) HasLastDeliveredDateTime() bool`

HasLastDeliveredDateTime returns a boolean if a field has been set.

### GetPreview

`func (o *MicrosoftGraphConversationThread) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *MicrosoftGraphConversationThread) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *MicrosoftGraphConversationThread) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *MicrosoftGraphConversationThread) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetTopic

`func (o *MicrosoftGraphConversationThread) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *MicrosoftGraphConversationThread) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *MicrosoftGraphConversationThread) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *MicrosoftGraphConversationThread) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetToRecipients

`func (o *MicrosoftGraphConversationThread) GetToRecipients() []MicrosoftGraphRecipient`

GetToRecipients returns the ToRecipients field if non-nil, zero value otherwise.

### GetToRecipientsOk

`func (o *MicrosoftGraphConversationThread) GetToRecipientsOk() (*[]MicrosoftGraphRecipient, bool)`

GetToRecipientsOk returns a tuple with the ToRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToRecipients

`func (o *MicrosoftGraphConversationThread) SetToRecipients(v []MicrosoftGraphRecipient)`

SetToRecipients sets ToRecipients field to given value.

### HasToRecipients

`func (o *MicrosoftGraphConversationThread) HasToRecipients() bool`

HasToRecipients returns a boolean if a field has been set.

### GetUniqueSenders

`func (o *MicrosoftGraphConversationThread) GetUniqueSenders() []string`

GetUniqueSenders returns the UniqueSenders field if non-nil, zero value otherwise.

### GetUniqueSendersOk

`func (o *MicrosoftGraphConversationThread) GetUniqueSendersOk() (*[]string, bool)`

GetUniqueSendersOk returns a tuple with the UniqueSenders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueSenders

`func (o *MicrosoftGraphConversationThread) SetUniqueSenders(v []string)`

SetUniqueSenders sets UniqueSenders field to given value.

### HasUniqueSenders

`func (o *MicrosoftGraphConversationThread) HasUniqueSenders() bool`

HasUniqueSenders returns a boolean if a field has been set.

### GetPosts

`func (o *MicrosoftGraphConversationThread) GetPosts() []MicrosoftGraphPost`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *MicrosoftGraphConversationThread) GetPostsOk() (*[]MicrosoftGraphPost, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *MicrosoftGraphConversationThread) SetPosts(v []MicrosoftGraphPost)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *MicrosoftGraphConversationThread) HasPosts() bool`

HasPosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


