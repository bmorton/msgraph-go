# MicrosoftGraphFolder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildCount** | Pointer to **NullableInt32** | Number of children contained immediately within this container. | [optional] 
**View** | Pointer to [**AnyOfmicrosoftGraphFolderView**](anyOf&lt;microsoft.graph.folderView&gt;.md) | A collection of properties defining the recommended view for the folder. | [optional] 

## Methods

### NewMicrosoftGraphFolder

`func NewMicrosoftGraphFolder() *MicrosoftGraphFolder`

NewMicrosoftGraphFolder instantiates a new MicrosoftGraphFolder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphFolderWithDefaults

`func NewMicrosoftGraphFolderWithDefaults() *MicrosoftGraphFolder`

NewMicrosoftGraphFolderWithDefaults instantiates a new MicrosoftGraphFolder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildCount

`func (o *MicrosoftGraphFolder) GetChildCount() int32`

GetChildCount returns the ChildCount field if non-nil, zero value otherwise.

### GetChildCountOk

`func (o *MicrosoftGraphFolder) GetChildCountOk() (*int32, bool)`

GetChildCountOk returns a tuple with the ChildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildCount

`func (o *MicrosoftGraphFolder) SetChildCount(v int32)`

SetChildCount sets ChildCount field to given value.

### HasChildCount

`func (o *MicrosoftGraphFolder) HasChildCount() bool`

HasChildCount returns a boolean if a field has been set.

### SetChildCountNil

`func (o *MicrosoftGraphFolder) SetChildCountNil(b bool)`

 SetChildCountNil sets the value for ChildCount to be an explicit nil

### UnsetChildCount
`func (o *MicrosoftGraphFolder) UnsetChildCount()`

UnsetChildCount ensures that no value is present for ChildCount, not even an explicit nil
### GetView

`func (o *MicrosoftGraphFolder) GetView() AnyOfmicrosoftGraphFolderView`

GetView returns the View field if non-nil, zero value otherwise.

### GetViewOk

`func (o *MicrosoftGraphFolder) GetViewOk() (*AnyOfmicrosoftGraphFolderView, bool)`

GetViewOk returns a tuple with the View field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetView

`func (o *MicrosoftGraphFolder) SetView(v AnyOfmicrosoftGraphFolderView)`

SetView sets View field to given value.

### HasView

`func (o *MicrosoftGraphFolder) HasView() bool`

HasView returns a boolean if a field has been set.

### SetViewNil

`func (o *MicrosoftGraphFolder) SetViewNil(b bool)`

 SetViewNil sets the value for View to be an explicit nil

### UnsetView
`func (o *MicrosoftGraphFolder) UnsetView()`

UnsetView ensures that no value is present for View, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


