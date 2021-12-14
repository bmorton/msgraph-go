# MicrosoftGraphDocumentSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedContentTypes** | Pointer to [**[]AnyOfmicrosoftGraphContentTypeInfo**](AnyOfmicrosoftGraphContentTypeInfo.md) | Content types allowed in document set. | [optional] 
**DefaultContents** | Pointer to [**[]AnyOfmicrosoftGraphDocumentSetContent**](AnyOfmicrosoftGraphDocumentSetContent.md) | Default contents of document set. | [optional] 
**PropagateWelcomePageChanges** | Pointer to **NullableBool** | Specifies whether to push welcome page changes to inherited content types. | [optional] 
**ShouldPrefixNameToFile** | Pointer to **NullableBool** | Add the name of the document set to each file name. | [optional] 
**WelcomePageUrl** | Pointer to **NullableString** | Welcome page absolute URL. | [optional] 
**SharedColumns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) |  | [optional] 
**WelcomePageColumns** | Pointer to [**[]MicrosoftGraphColumnDefinition**](MicrosoftGraphColumnDefinition.md) |  | [optional] 

## Methods

### NewMicrosoftGraphDocumentSet

`func NewMicrosoftGraphDocumentSet() *MicrosoftGraphDocumentSet`

NewMicrosoftGraphDocumentSet instantiates a new MicrosoftGraphDocumentSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphDocumentSetWithDefaults

`func NewMicrosoftGraphDocumentSetWithDefaults() *MicrosoftGraphDocumentSet`

NewMicrosoftGraphDocumentSetWithDefaults instantiates a new MicrosoftGraphDocumentSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedContentTypes

`func (o *MicrosoftGraphDocumentSet) GetAllowedContentTypes() []*AnyOfmicrosoftGraphContentTypeInfo`

GetAllowedContentTypes returns the AllowedContentTypes field if non-nil, zero value otherwise.

### GetAllowedContentTypesOk

`func (o *MicrosoftGraphDocumentSet) GetAllowedContentTypesOk() (*[]*AnyOfmicrosoftGraphContentTypeInfo, bool)`

GetAllowedContentTypesOk returns a tuple with the AllowedContentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedContentTypes

`func (o *MicrosoftGraphDocumentSet) SetAllowedContentTypes(v []*AnyOfmicrosoftGraphContentTypeInfo)`

SetAllowedContentTypes sets AllowedContentTypes field to given value.

### HasAllowedContentTypes

`func (o *MicrosoftGraphDocumentSet) HasAllowedContentTypes() bool`

HasAllowedContentTypes returns a boolean if a field has been set.

### GetDefaultContents

`func (o *MicrosoftGraphDocumentSet) GetDefaultContents() []*AnyOfmicrosoftGraphDocumentSetContent`

GetDefaultContents returns the DefaultContents field if non-nil, zero value otherwise.

### GetDefaultContentsOk

`func (o *MicrosoftGraphDocumentSet) GetDefaultContentsOk() (*[]*AnyOfmicrosoftGraphDocumentSetContent, bool)`

GetDefaultContentsOk returns a tuple with the DefaultContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultContents

`func (o *MicrosoftGraphDocumentSet) SetDefaultContents(v []*AnyOfmicrosoftGraphDocumentSetContent)`

SetDefaultContents sets DefaultContents field to given value.

### HasDefaultContents

`func (o *MicrosoftGraphDocumentSet) HasDefaultContents() bool`

HasDefaultContents returns a boolean if a field has been set.

### GetPropagateWelcomePageChanges

`func (o *MicrosoftGraphDocumentSet) GetPropagateWelcomePageChanges() bool`

GetPropagateWelcomePageChanges returns the PropagateWelcomePageChanges field if non-nil, zero value otherwise.

### GetPropagateWelcomePageChangesOk

`func (o *MicrosoftGraphDocumentSet) GetPropagateWelcomePageChangesOk() (*bool, bool)`

GetPropagateWelcomePageChangesOk returns a tuple with the PropagateWelcomePageChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropagateWelcomePageChanges

`func (o *MicrosoftGraphDocumentSet) SetPropagateWelcomePageChanges(v bool)`

SetPropagateWelcomePageChanges sets PropagateWelcomePageChanges field to given value.

### HasPropagateWelcomePageChanges

`func (o *MicrosoftGraphDocumentSet) HasPropagateWelcomePageChanges() bool`

HasPropagateWelcomePageChanges returns a boolean if a field has been set.

### SetPropagateWelcomePageChangesNil

`func (o *MicrosoftGraphDocumentSet) SetPropagateWelcomePageChangesNil(b bool)`

 SetPropagateWelcomePageChangesNil sets the value for PropagateWelcomePageChanges to be an explicit nil

### UnsetPropagateWelcomePageChanges
`func (o *MicrosoftGraphDocumentSet) UnsetPropagateWelcomePageChanges()`

UnsetPropagateWelcomePageChanges ensures that no value is present for PropagateWelcomePageChanges, not even an explicit nil
### GetShouldPrefixNameToFile

`func (o *MicrosoftGraphDocumentSet) GetShouldPrefixNameToFile() bool`

GetShouldPrefixNameToFile returns the ShouldPrefixNameToFile field if non-nil, zero value otherwise.

### GetShouldPrefixNameToFileOk

`func (o *MicrosoftGraphDocumentSet) GetShouldPrefixNameToFileOk() (*bool, bool)`

GetShouldPrefixNameToFileOk returns a tuple with the ShouldPrefixNameToFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShouldPrefixNameToFile

`func (o *MicrosoftGraphDocumentSet) SetShouldPrefixNameToFile(v bool)`

SetShouldPrefixNameToFile sets ShouldPrefixNameToFile field to given value.

### HasShouldPrefixNameToFile

`func (o *MicrosoftGraphDocumentSet) HasShouldPrefixNameToFile() bool`

HasShouldPrefixNameToFile returns a boolean if a field has been set.

### SetShouldPrefixNameToFileNil

`func (o *MicrosoftGraphDocumentSet) SetShouldPrefixNameToFileNil(b bool)`

 SetShouldPrefixNameToFileNil sets the value for ShouldPrefixNameToFile to be an explicit nil

### UnsetShouldPrefixNameToFile
`func (o *MicrosoftGraphDocumentSet) UnsetShouldPrefixNameToFile()`

UnsetShouldPrefixNameToFile ensures that no value is present for ShouldPrefixNameToFile, not even an explicit nil
### GetWelcomePageUrl

`func (o *MicrosoftGraphDocumentSet) GetWelcomePageUrl() string`

GetWelcomePageUrl returns the WelcomePageUrl field if non-nil, zero value otherwise.

### GetWelcomePageUrlOk

`func (o *MicrosoftGraphDocumentSet) GetWelcomePageUrlOk() (*string, bool)`

GetWelcomePageUrlOk returns a tuple with the WelcomePageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomePageUrl

`func (o *MicrosoftGraphDocumentSet) SetWelcomePageUrl(v string)`

SetWelcomePageUrl sets WelcomePageUrl field to given value.

### HasWelcomePageUrl

`func (o *MicrosoftGraphDocumentSet) HasWelcomePageUrl() bool`

HasWelcomePageUrl returns a boolean if a field has been set.

### SetWelcomePageUrlNil

`func (o *MicrosoftGraphDocumentSet) SetWelcomePageUrlNil(b bool)`

 SetWelcomePageUrlNil sets the value for WelcomePageUrl to be an explicit nil

### UnsetWelcomePageUrl
`func (o *MicrosoftGraphDocumentSet) UnsetWelcomePageUrl()`

UnsetWelcomePageUrl ensures that no value is present for WelcomePageUrl, not even an explicit nil
### GetSharedColumns

`func (o *MicrosoftGraphDocumentSet) GetSharedColumns() []MicrosoftGraphColumnDefinition`

GetSharedColumns returns the SharedColumns field if non-nil, zero value otherwise.

### GetSharedColumnsOk

`func (o *MicrosoftGraphDocumentSet) GetSharedColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetSharedColumnsOk returns a tuple with the SharedColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedColumns

`func (o *MicrosoftGraphDocumentSet) SetSharedColumns(v []MicrosoftGraphColumnDefinition)`

SetSharedColumns sets SharedColumns field to given value.

### HasSharedColumns

`func (o *MicrosoftGraphDocumentSet) HasSharedColumns() bool`

HasSharedColumns returns a boolean if a field has been set.

### GetWelcomePageColumns

`func (o *MicrosoftGraphDocumentSet) GetWelcomePageColumns() []MicrosoftGraphColumnDefinition`

GetWelcomePageColumns returns the WelcomePageColumns field if non-nil, zero value otherwise.

### GetWelcomePageColumnsOk

`func (o *MicrosoftGraphDocumentSet) GetWelcomePageColumnsOk() (*[]MicrosoftGraphColumnDefinition, bool)`

GetWelcomePageColumnsOk returns a tuple with the WelcomePageColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomePageColumns

`func (o *MicrosoftGraphDocumentSet) SetWelcomePageColumns(v []MicrosoftGraphColumnDefinition)`

SetWelcomePageColumns sets WelcomePageColumns field to given value.

### HasWelcomePageColumns

`func (o *MicrosoftGraphDocumentSet) HasWelcomePageColumns() bool`

HasWelcomePageColumns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


