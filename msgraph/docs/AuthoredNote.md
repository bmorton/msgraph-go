# AuthoredNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to [**AnyOfmicrosoftGraphIdentity**](anyOf&lt;microsoft.graph.identity&gt;.md) | Identity information about the note&#39;s author. | [optional] 
**Content** | Pointer to [**AnyOfmicrosoftGraphItemBody**](anyOf&lt;microsoft.graph.itemBody&gt;.md) | The content of the note. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 

## Methods

### NewAuthoredNote

`func NewAuthoredNote() *AuthoredNote`

NewAuthoredNote instantiates a new AuthoredNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthoredNoteWithDefaults

`func NewAuthoredNoteWithDefaults() *AuthoredNote`

NewAuthoredNoteWithDefaults instantiates a new AuthoredNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *AuthoredNote) GetAuthor() AnyOfmicrosoftGraphIdentity`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AuthoredNote) GetAuthorOk() (*AnyOfmicrosoftGraphIdentity, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AuthoredNote) SetAuthor(v AnyOfmicrosoftGraphIdentity)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AuthoredNote) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *AuthoredNote) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *AuthoredNote) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetContent

`func (o *AuthoredNote) GetContent() AnyOfmicrosoftGraphItemBody`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AuthoredNote) GetContentOk() (*AnyOfmicrosoftGraphItemBody, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AuthoredNote) SetContent(v AnyOfmicrosoftGraphItemBody)`

SetContent sets Content field to given value.

### HasContent

`func (o *AuthoredNote) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *AuthoredNote) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *AuthoredNote) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetCreatedDateTime

`func (o *AuthoredNote) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AuthoredNote) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AuthoredNote) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AuthoredNote) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *AuthoredNote) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *AuthoredNote) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


