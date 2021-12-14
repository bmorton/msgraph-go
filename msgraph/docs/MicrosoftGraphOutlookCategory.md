# MicrosoftGraphOutlookCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Color** | Pointer to [**AnyOfmicrosoftGraphCategoryColor**](anyOf&lt;microsoft.graph.categoryColor&gt;.md) | A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below. | [optional] 
**DisplayName** | Pointer to **NullableString** | A unique name that identifies a category in the user&#39;s mailbox. After a category is created, the name cannot be changed. Read-only. | [optional] 

## Methods

### NewMicrosoftGraphOutlookCategory

`func NewMicrosoftGraphOutlookCategory() *MicrosoftGraphOutlookCategory`

NewMicrosoftGraphOutlookCategory instantiates a new MicrosoftGraphOutlookCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphOutlookCategoryWithDefaults

`func NewMicrosoftGraphOutlookCategoryWithDefaults() *MicrosoftGraphOutlookCategory`

NewMicrosoftGraphOutlookCategoryWithDefaults instantiates a new MicrosoftGraphOutlookCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphOutlookCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphOutlookCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphOutlookCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphOutlookCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetColor

`func (o *MicrosoftGraphOutlookCategory) GetColor() AnyOfmicrosoftGraphCategoryColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MicrosoftGraphOutlookCategory) GetColorOk() (*AnyOfmicrosoftGraphCategoryColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MicrosoftGraphOutlookCategory) SetColor(v AnyOfmicrosoftGraphCategoryColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *MicrosoftGraphOutlookCategory) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MicrosoftGraphOutlookCategory) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MicrosoftGraphOutlookCategory) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphOutlookCategory) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphOutlookCategory) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphOutlookCategory) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphOutlookCategory) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphOutlookCategory) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphOutlookCategory) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


