# MicrosoftGraphOnenotePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Self** | Pointer to **NullableString** | The endpoint where you can get details about the page. Read-only. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The date and time when the page was created. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**Content** | Pointer to **NullableString** | The page&#39;s HTML content. | [optional] 
**ContentUrl** | Pointer to **NullableString** | The URL for the page&#39;s HTML content.  Read-only. | [optional] 
**CreatedByAppId** | Pointer to **NullableString** | The unique identifier of the application that created the page. Read-only. | [optional] 
**LastModifiedDateTime** | Pointer to **NullableTime** | The date and time when the page was last modified. The timestamp represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**Level** | Pointer to **NullableInt32** | The indentation level of the page. Read-only. | [optional] 
**Links** | Pointer to [**AnyOfmicrosoftGraphPageLinks**](anyOf&lt;microsoft.graph.pageLinks&gt;.md) | Links for opening the page. The oneNoteClientURL link opens the page in the OneNote native client if it &#39;s installed. The oneNoteWebUrl link opens the page in OneNote on the web. Read-only. | [optional] 
**Order** | Pointer to **NullableInt32** | The order of the page within its parent section. Read-only. | [optional] 
**Title** | Pointer to **NullableString** | The title of the page. | [optional] 
**UserTags** | Pointer to **[]string** |  | [optional] 
**ParentNotebook** | Pointer to [**AnyOfmicrosoftGraphNotebook**](anyOf&lt;microsoft.graph.notebook&gt;.md) | The notebook that contains the page.  Read-only. | [optional] 
**ParentSection** | Pointer to [**AnyOfmicrosoftGraphOnenoteSection**](anyOf&lt;microsoft.graph.onenoteSection&gt;.md) | The section that contains the page. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphOnenotePage

`func NewMicrosoftGraphOnenotePage() *MicrosoftGraphOnenotePage`

NewMicrosoftGraphOnenotePage instantiates a new MicrosoftGraphOnenotePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOnenotePageWithDefaults

`func NewMicrosoftGraphOnenotePageWithDefaults() *MicrosoftGraphOnenotePage`

NewMicrosoftGraphOnenotePageWithDefaults instantiates a new MicrosoftGraphOnenotePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOnenotePage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOnenotePage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOnenotePage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOnenotePage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSelf

`func (o *MicrosoftGraphOnenotePage) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *MicrosoftGraphOnenotePage) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *MicrosoftGraphOnenotePage) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *MicrosoftGraphOnenotePage) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### SetSelfNil

`func (o *MicrosoftGraphOnenotePage) SetSelfNil(b bool)`

 SetSelfNil sets the value for Self to be an explicit nil

### UnsetSelf
`func (o *MicrosoftGraphOnenotePage) UnsetSelf()`

UnsetSelf ensures that no value is present for Self, not even an explicit nil
### GetCreatedDateTime

`func (o *MicrosoftGraphOnenotePage) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *MicrosoftGraphOnenotePage) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *MicrosoftGraphOnenotePage) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *MicrosoftGraphOnenotePage) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *MicrosoftGraphOnenotePage) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *MicrosoftGraphOnenotePage) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetContent

`func (o *MicrosoftGraphOnenotePage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MicrosoftGraphOnenotePage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MicrosoftGraphOnenotePage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MicrosoftGraphOnenotePage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MicrosoftGraphOnenotePage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MicrosoftGraphOnenotePage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentUrl

`func (o *MicrosoftGraphOnenotePage) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *MicrosoftGraphOnenotePage) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *MicrosoftGraphOnenotePage) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *MicrosoftGraphOnenotePage) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrlNil

`func (o *MicrosoftGraphOnenotePage) SetContentUrlNil(b bool)`

 SetContentUrlNil sets the value for ContentUrl to be an explicit nil

### UnsetContentUrl
`func (o *MicrosoftGraphOnenotePage) UnsetContentUrl()`

UnsetContentUrl ensures that no value is present for ContentUrl, not even an explicit nil
### GetCreatedByAppId

`func (o *MicrosoftGraphOnenotePage) GetCreatedByAppId() string`

GetCreatedByAppId returns the CreatedByAppId field if non-nil, zero value otherwise.

### GetCreatedByAppIdOk

`func (o *MicrosoftGraphOnenotePage) GetCreatedByAppIdOk() (*string, bool)`

GetCreatedByAppIdOk returns a tuple with the CreatedByAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAppId

`func (o *MicrosoftGraphOnenotePage) SetCreatedByAppId(v string)`

SetCreatedByAppId sets CreatedByAppId field to given value.

### HasCreatedByAppId

`func (o *MicrosoftGraphOnenotePage) HasCreatedByAppId() bool`

HasCreatedByAppId returns a boolean if a field has been set.

### SetCreatedByAppIdNil

`func (o *MicrosoftGraphOnenotePage) SetCreatedByAppIdNil(b bool)`

 SetCreatedByAppIdNil sets the value for CreatedByAppId to be an explicit nil

### UnsetCreatedByAppId
`func (o *MicrosoftGraphOnenotePage) UnsetCreatedByAppId()`

UnsetCreatedByAppId ensures that no value is present for CreatedByAppId, not even an explicit nil
### GetLastModifiedDateTime

`func (o *MicrosoftGraphOnenotePage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *MicrosoftGraphOnenotePage) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *MicrosoftGraphOnenotePage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *MicrosoftGraphOnenotePage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *MicrosoftGraphOnenotePage) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *MicrosoftGraphOnenotePage) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetLevel

`func (o *MicrosoftGraphOnenotePage) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *MicrosoftGraphOnenotePage) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *MicrosoftGraphOnenotePage) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *MicrosoftGraphOnenotePage) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *MicrosoftGraphOnenotePage) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *MicrosoftGraphOnenotePage) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetLinks

`func (o *MicrosoftGraphOnenotePage) GetLinks() AnyOfmicrosoftGraphPageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *MicrosoftGraphOnenotePage) GetLinksOk() (*AnyOfmicrosoftGraphPageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *MicrosoftGraphOnenotePage) SetLinks(v AnyOfmicrosoftGraphPageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *MicrosoftGraphOnenotePage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *MicrosoftGraphOnenotePage) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *MicrosoftGraphOnenotePage) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetOrder

`func (o *MicrosoftGraphOnenotePage) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MicrosoftGraphOnenotePage) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MicrosoftGraphOnenotePage) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MicrosoftGraphOnenotePage) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *MicrosoftGraphOnenotePage) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *MicrosoftGraphOnenotePage) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetTitle

`func (o *MicrosoftGraphOnenotePage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MicrosoftGraphOnenotePage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MicrosoftGraphOnenotePage) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MicrosoftGraphOnenotePage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MicrosoftGraphOnenotePage) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MicrosoftGraphOnenotePage) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserTags

`func (o *MicrosoftGraphOnenotePage) GetUserTags() []*string`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *MicrosoftGraphOnenotePage) GetUserTagsOk() (*[]*string, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTags

`func (o *MicrosoftGraphOnenotePage) SetUserTags(v []*string)`

SetUserTags sets UserTags field to given value.

### HasUserTags

`func (o *MicrosoftGraphOnenotePage) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.

### GetParentNotebook

`func (o *MicrosoftGraphOnenotePage) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *MicrosoftGraphOnenotePage) GetParentNotebookOk() (*AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNotebook

`func (o *MicrosoftGraphOnenotePage) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook sets ParentNotebook field to given value.

### HasParentNotebook

`func (o *MicrosoftGraphOnenotePage) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebookNil

`func (o *MicrosoftGraphOnenotePage) SetParentNotebookNil(b bool)`

 SetParentNotebookNil sets the value for ParentNotebook to be an explicit nil

### UnsetParentNotebook
`func (o *MicrosoftGraphOnenotePage) UnsetParentNotebook()`

UnsetParentNotebook ensures that no value is present for ParentNotebook, not even an explicit nil
### GetParentSection

`func (o *MicrosoftGraphOnenotePage) GetParentSection() AnyOfmicrosoftGraphOnenoteSection`

GetParentSection returns the ParentSection field if non-nil, zero value otherwise.

### GetParentSectionOk

`func (o *MicrosoftGraphOnenotePage) GetParentSectionOk() (*AnyOfmicrosoftGraphOnenoteSection, bool)`

GetParentSectionOk returns a tuple with the ParentSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSection

`func (o *MicrosoftGraphOnenotePage) SetParentSection(v AnyOfmicrosoftGraphOnenoteSection)`

SetParentSection sets ParentSection field to given value.

### HasParentSection

`func (o *MicrosoftGraphOnenotePage) HasParentSection() bool`

HasParentSection returns a boolean if a field has been set.

### SetParentSectionNil

`func (o *MicrosoftGraphOnenotePage) SetParentSectionNil(b bool)`

 SetParentSectionNil sets the value for ParentSection to be an explicit nil

### UnsetParentSection
`func (o *MicrosoftGraphOnenotePage) UnsetParentSection()`

UnsetParentSection ensures that no value is present for ParentSection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


