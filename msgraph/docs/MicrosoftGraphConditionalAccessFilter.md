# MicrosoftGraphConditionalAccessFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to [**AnyOfmicrosoftGraphFilterMode**](anyOf&lt;microsoft.graph.filterMode&gt;.md) | Mode to use for the filter. Possible values are include or exclude. | [optional] 
**Rule** | Pointer to **string** | Rule syntax is similar to that used for membership rules for groups in Azure Active Directory (Azure AD). For details, see rules with multiple expressions | [optional] 

## Methods

### NewMicrosoftGraphConditionalAccessFilter

`func NewMicrosoftGraphConditionalAccessFilter() *MicrosoftGraphConditionalAccessFilter`

NewMicrosoftGraphConditionalAccessFilter instantiates a new MicrosoftGraphConditionalAccessFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphConditionalAccessFilterWithDefaults

`func NewMicrosoftGraphConditionalAccessFilterWithDefaults() *MicrosoftGraphConditionalAccessFilter`

NewMicrosoftGraphConditionalAccessFilterWithDefaults instantiates a new MicrosoftGraphConditionalAccessFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *MicrosoftGraphConditionalAccessFilter) GetMode() AnyOfmicrosoftGraphFilterMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *MicrosoftGraphConditionalAccessFilter) GetModeOk() (*AnyOfmicrosoftGraphFilterMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *MicrosoftGraphConditionalAccessFilter) SetMode(v AnyOfmicrosoftGraphFilterMode)`

SetMode sets Mode field to given value.

### HasMode

`func (o *MicrosoftGraphConditionalAccessFilter) HasMode() bool`

HasMode returns a boolean if a field has been set.

### SetModeNil

`func (o *MicrosoftGraphConditionalAccessFilter) SetModeNil(b bool)`

 SetModeNil sets the value for Mode to be an explicit nil

### UnsetMode
`func (o *MicrosoftGraphConditionalAccessFilter) UnsetMode()`

UnsetMode ensures that no value is present for Mode, not even an explicit nil
### GetRule

`func (o *MicrosoftGraphConditionalAccessFilter) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *MicrosoftGraphConditionalAccessFilter) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *MicrosoftGraphConditionalAccessFilter) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *MicrosoftGraphConditionalAccessFilter) HasRule() bool`

HasRule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


