# MicrosoftGraphAuthoredNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Author** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Identity information about the note&#39;s author. | [optional] 
**Content** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) | The content of the note. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 

## Methods

### NewMicrosoftGraphAuthoredNote

`func NewMicrosoftGraphAuthoredNote() *MicrosoftGraphAuthoredNote`

NewMicrosoftGraphAuthoredNote instantiates a new MicrosoftGraphAuthoredNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAuthoredNoteWithDefaults

`func NewMicrosoftGraphAuthoredNoteWithDefaults() *MicrosoftGraphAuthoredNote`

NewMicrosoftGraphAuthoredNoteWithDefaults instantiates a new MicrosoftGraphAuthoredNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphAuthoredNote) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphAuthoredNote) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphAuthoredNote) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphAuthoredNote) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthor

`func (o *MicrosoftGraphAuthoredNote) GetAuthor() AnyOfmicrosoftGraphIdentity`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *MicrosoftGraphAuthoredNote) GetAuthorOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *MicrosoftGraphAuthoredNote) SetAuthor(v AnyOfmicrosoftGraphIdentity)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *MicrosoftGraphAuthoredNote) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *MicrosoftGraphAuthoredNote) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *MicrosoftGraphAuthoredNote) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetContent

`func (o *MicrosoftGraphAuthoredNote) GetContent() AnyOfmicrosoftGraphItemBody`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphAuthoredNote) GetContentOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphAuthoredNote) SetContent(v AnyOfmicrosoftGraphItemBody)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphAuthoredNote) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphAuthoredNote) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphAuthoredNote) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphAuthoredNote) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphAuthoredNote) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphAuthoredNote) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphAuthoredNote) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphAuthoredNote) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphAuthoredNote) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


