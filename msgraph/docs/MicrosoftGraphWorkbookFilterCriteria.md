# MicrosoftGraphWorkbookFilterCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | Pointer to **NullableString** |  | [optional] 
**Criterion1** | Pointer to **NullableString** |  | [optional] 
**Criterion2** | Pointer to **NullableString** |  | [optional] 
**DynamicCriteria** | Pointer to **string** |  | [optional] 
**FilterOn** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to [**AnyOfmicrosoftGraphWorkbookIcon**](anyOf&lt;microsoft.graph.workbookIcon&gt;.md) |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Values** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) |  | [optional] 

## Methods

### NewMicrosoftGraphWorkbookFilterCriteria

`func NewMicrosoftGraphWorkbookFilterCriteria() *MicrosoftGraphWorkbookFilterCriteria`

NewMicrosoftGraphWorkbookFilterCriteria instantiates a new MicrosoftGraphWorkbookFilterCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphWorkbookFilterCriteriaWithDefaults

`func NewMicrosoftGraphWorkbookFilterCriteriaWithDefaults() *MicrosoftGraphWorkbookFilterCriteria`

NewMicrosoftGraphWorkbookFilterCriteriaWithDefaults instantiates a new MicrosoftGraphWorkbookFilterCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *MicrosoftGraphWorkbookFilterCriteria) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *MicrosoftGraphWorkbookFilterCriteria) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCriterion1

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetCriterion1() string`

GetCriterion1 returns the Criterion1 field if non-nil, zero value otherwise.

### GetCriterion1Ok

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetCriterion1Ok() (*string, bool)`

GetCriterion1Ok returns a tuple with the Criterion1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriterion1

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetCriterion1(v string)`

SetCriterion1 sets Criterion1 field to given value.

### HasCriterion1

`func (o *MicrosoftGraphWorkbookFilterCriteria) HasCriterion1() bool`

HasCriterion1 returns a boolean if a field has been set.

### SetCriterion1Nil

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetCriterion1Nil(b bool)`

 SetCriterion1Nil sets the value for Criterion1 to be an explicit nil

### UnsetCriterion1
`func (o *MicrosoftGraphWorkbookFilterCriteria) UnsetCriterion1()`

UnsetCriterion1 ensures that no value is present for Criterion1, not even an explicit nil
### GetCriterion2

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetCriterion2() string`

GetCriterion2 returns the Criterion2 field if non-nil, zero value otherwise.

### GetCriterion2Ok

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetCriterion2Ok() (*string, bool)`

GetCriterion2Ok returns a tuple with the Criterion2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriterion2

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetCriterion2(v string)`

SetCriterion2 sets Criterion2 field to given value.

### HasCriterion2

`func (o *MicrosoftGraphWorkbookFilterCriteria) HasCriterion2() bool`

HasCriterion2 returns a boolean if a field has been set.

### SetCriterion2Nil

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetCriterion2Nil(b bool)`

 SetCriterion2Nil sets the value for Criterion2 to be an explicit nil

### UnsetCriterion2
`func (o *MicrosoftGraphWorkbookFilterCriteria) UnsetCriterion2()`

UnsetCriterion2 ensures that no value is present for Criterion2, not even an explicit nil
### GetDynamicCriteria

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetDynamicCriteria() string`

GetDynamicCriteria returns the DynamicCriteria field if non-nil, zero value otherwise.

### GetDynamicCriteriaOk

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetDynamicCriteriaOk() (*string, bool)`

GetDynamicCriteriaOk returns a tuple with the DynamicCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCriteria

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetDynamicCriteria(v string)`

SetDynamicCriteria sets DynamicCriteria field to given value.

### HasDynamicCriteria

`func (o *MicrosoftGraphWorkbookFilterCriteria) HasDynamicCriteria() bool`

HasDynamicCriteria returns a boolean if a field has been set.

### GetFilterOn

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetFilterOn() string`

GetFilterOn returns the FilterOn field if non-nil, zero value otherwise.

### GetFilterOnOk

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetFilterOnOk() (*string, bool)`

GetFilterOnOk returns a tuple with the FilterOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterOn

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetFilterOn(v string)`

SetFilterOn sets FilterOn field to given value.

### HasFilterOn

`func (o *MicrosoftGraphWorkbookFilterCriteria) HasFilterOn() bool`

HasFilterOn returns a boolean if a field has been set.

### GetIcon

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetIcon() AnyOfmicrosoftGraphWorkbookIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetIconOk() (*AnyOfmicrosoftGraphWorkbookIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetIcon(v AnyOfmicrosoftGraphWorkbookIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *MicrosoftGraphWorkbookFilterCriteria) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### SetIconNil

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetIconNil(b bool)`

 SetIconNil sets the value for Icon to be an explicit nil

### UnsetIcon
`func (o *MicrosoftGraphWorkbookFilterCriteria) UnsetIcon()`

UnsetIcon ensures that no value is present for Icon, not even an explicit nil
### GetOperator

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *MicrosoftGraphWorkbookFilterCriteria) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValues

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetValues() AnyOfobject`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *MicrosoftGraphWorkbookFilterCriteria) GetValuesOk() (*AnyOfobject, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetValues(v AnyOfobject)`

SetValues sets Values field to given value.

### HasValues

`func (o *MicrosoftGraphWorkbookFilterCriteria) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *MicrosoftGraphWorkbookFilterCriteria) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *MicrosoftGraphWorkbookFilterCriteria) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


