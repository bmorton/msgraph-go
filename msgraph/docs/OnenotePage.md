# OnenotePage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewOnenotePage

`func NewOnenotePage() *OnenotePage`

NewOnenotePage instantiates a new OnenotePage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnenotePageWithDefaults

`func NewOnenotePageWithDefaults() *OnenotePage`

NewOnenotePageWithDefaults instantiates a new OnenotePage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *OnenotePage) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *OnenotePage) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *OnenotePage) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *OnenotePage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *OnenotePage) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *OnenotePage) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetContentUrl

`func (o *OnenotePage) GetContentUrl() string`

GetContentUrl returns the ContentUrl field if non-nil, zero value otherwise.

### GetContentUrlOk

`func (o *OnenotePage) GetContentUrlOk() (*string, bool)`

GetContentUrlOk returns a tuple with the ContentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentUrl

`func (o *OnenotePage) SetContentUrl(v string)`

SetContentUrl sets ContentUrl field to given value.

### HasContentUrl

`func (o *OnenotePage) HasContentUrl() bool`

HasContentUrl returns a boolean if a field has been set.

### SetContentUrlNil

`func (o *OnenotePage) SetContentUrlNil(b bool)`

 SetContentUrlNil sets the value for ContentUrl to be an explicit nil

### UnsetContentUrl
`func (o *OnenotePage) UnsetContentUrl()`

UnsetContentUrl ensures that no value is present for ContentUrl, not even an explicit nil
### GetCreatedByAppId

`func (o *OnenotePage) GetCreatedByAppId() string`

GetCreatedByAppId returns the CreatedByAppId field if non-nil, zero value otherwise.

### GetCreatedByAppIdOk

`func (o *OnenotePage) GetCreatedByAppIdOk() (*string, bool)`

GetCreatedByAppIdOk returns a tuple with the CreatedByAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByAppId

`func (o *OnenotePage) SetCreatedByAppId(v string)`

SetCreatedByAppId sets CreatedByAppId field to given value.

### HasCreatedByAppId

`func (o *OnenotePage) HasCreatedByAppId() bool`

HasCreatedByAppId returns a boolean if a field has been set.

### SetCreatedByAppIdNil

`func (o *OnenotePage) SetCreatedByAppIdNil(b bool)`

 SetCreatedByAppIdNil sets the value for CreatedByAppId to be an explicit nil

### UnsetCreatedByAppId
`func (o *OnenotePage) UnsetCreatedByAppId()`

UnsetCreatedByAppId ensures that no value is present for CreatedByAppId, not even an explicit nil
### GetLastModifiedDateTime

`func (o *OnenotePage) GetLastModifiedDateTime() time.Time`

GetLastModifiedDateTime returns the LastModifiedDateTime field if non-nil, zero value otherwise.

### GetLastModifiedDateTimeOk

`func (o *OnenotePage) GetLastModifiedDateTimeOk() (*time.Time, bool)`

GetLastModifiedDateTimeOk returns a tuple with the LastModifiedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedDateTime

`func (o *OnenotePage) SetLastModifiedDateTime(v time.Time)`

SetLastModifiedDateTime sets LastModifiedDateTime field to given value.

### HasLastModifiedDateTime

`func (o *OnenotePage) HasLastModifiedDateTime() bool`

HasLastModifiedDateTime returns a boolean if a field has been set.

### SetLastModifiedDateTimeNil

`func (o *OnenotePage) SetLastModifiedDateTimeNil(b bool)`

 SetLastModifiedDateTimeNil sets the value for LastModifiedDateTime to be an explicit nil

### UnsetLastModifiedDateTime
`func (o *OnenotePage) UnsetLastModifiedDateTime()`

UnsetLastModifiedDateTime ensures that no value is present for LastModifiedDateTime, not even an explicit nil
### GetLevel

`func (o *OnenotePage) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *OnenotePage) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *OnenotePage) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *OnenotePage) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### SetLevelNil

`func (o *OnenotePage) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *OnenotePage) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetLinks

`func (o *OnenotePage) GetLinks() AnyOfmicrosoftGraphPageLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *OnenotePage) GetLinksOk() (*AnyOfmicrosoftGraphPageLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *OnenotePage) SetLinks(v AnyOfmicrosoftGraphPageLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *OnenotePage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### SetLinksNil

`func (o *OnenotePage) SetLinksNil(b bool)`

 SetLinksNil sets the value for Links to be an explicit nil

### UnsetLinks
`func (o *OnenotePage) UnsetLinks()`

UnsetLinks ensures that no value is present for Links, not even an explicit nil
### GetOrder

`func (o *OnenotePage) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *OnenotePage) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *OnenotePage) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *OnenotePage) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### SetOrderNil

`func (o *OnenotePage) SetOrderNil(b bool)`

 SetOrderNil sets the value for Order to be an explicit nil

### UnsetOrder
`func (o *OnenotePage) UnsetOrder()`

UnsetOrder ensures that no value is present for Order, not even an explicit nil
### GetTitle

`func (o *OnenotePage) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *OnenotePage) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *OnenotePage) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *OnenotePage) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *OnenotePage) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *OnenotePage) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUserTags

`func (o *OnenotePage) GetUserTags() []*string`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *OnenotePage) GetUserTagsOk() (*[]*string, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTags

`func (o *OnenotePage) SetUserTags(v []*string)`

SetUserTags sets UserTags field to given value.

### HasUserTags

`func (o *OnenotePage) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.

### GetParentNotebook

`func (o *OnenotePage) GetParentNotebook() AnyOfmicrosoftGraphNotebook`

GetParentNotebook returns the ParentNotebook field if non-nil, zero value otherwise.

### GetParentNotebookOk

`func (o *OnenotePage) GetParentNotebookOk() (*AnyOfmicrosoftGraphNotebook, bool)`

GetParentNotebookOk returns a tuple with the ParentNotebook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNotebook

`func (o *OnenotePage) SetParentNotebook(v AnyOfmicrosoftGraphNotebook)`

SetParentNotebook sets ParentNotebook field to given value.

### HasParentNotebook

`func (o *OnenotePage) HasParentNotebook() bool`

HasParentNotebook returns a boolean if a field has been set.

### SetParentNotebookNil

`func (o *OnenotePage) SetParentNotebookNil(b bool)`

 SetParentNotebookNil sets the value for ParentNotebook to be an explicit nil

### UnsetParentNotebook
`func (o *OnenotePage) UnsetParentNotebook()`

UnsetParentNotebook ensures that no value is present for ParentNotebook, not even an explicit nil
### GetParentSection

`func (o *OnenotePage) GetParentSection() AnyOfmicrosoftGraphOnenoteSection`

GetParentSection returns the ParentSection field if non-nil, zero value otherwise.

### GetParentSectionOk

`func (o *OnenotePage) GetParentSectionOk() (*AnyOfmicrosoftGraphOnenoteSection, bool)`

GetParentSectionOk returns a tuple with the ParentSection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSection

`func (o *OnenotePage) SetParentSection(v AnyOfmicrosoftGraphOnenoteSection)`

SetParentSection sets ParentSection field to given value.

### HasParentSection

`func (o *OnenotePage) HasParentSection() bool`

HasParentSection returns a boolean if a field has been set.

### SetParentSectionNil

`func (o *OnenotePage) SetParentSectionNil(b bool)`

 SetParentSectionNil sets the value for ParentSection to be an explicit nil

### UnsetParentSection
`func (o *OnenotePage) UnsetParentSection()`

UnsetParentSection ensures that no value is present for ParentSection, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


