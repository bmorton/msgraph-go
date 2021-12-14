# OutlookCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to [**AnyOfmicrosoftGraphCategoryColor**](anyOf&lt;microsoft.graph.categoryColor&gt;.md) | A pre-set color constant that characterizes a category, and that is mapped to one of 25 predefined colors. See the note below. | [optional] 
**DisplayName** | Pointer to **NullableString** | A unique name that identifies a category in the user&#39;s mailbox. After a category is created, the name cannot be changed. Read-only. | [optional] 

## Methods

### NewOutlookCategory

`func NewOutlookCategory() *OutlookCategory`

NewOutlookCategory instantiates a new OutlookCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlookCategoryWithDefaults

`func NewOutlookCategoryWithDefaults() *OutlookCategory`

NewOutlookCategoryWithDefaults instantiates a new OutlookCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *OutlookCategory) GetColor() AnyOfmicrosoftGraphCategoryColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *OutlookCategory) GetColorOk() (*AnyOfmicrosoftGraphCategoryColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *OutlookCategory) SetColor(v AnyOfmicrosoftGraphCategoryColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *OutlookCategory) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *OutlookCategory) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *OutlookCategory) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetDisplayName

`func (o *OutlookCategory) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *OutlookCategory) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *OutlookCategory) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *OutlookCategory) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *OutlookCategory) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *OutlookCategory) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


