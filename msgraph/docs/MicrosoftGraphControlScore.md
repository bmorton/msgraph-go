# MicrosoftGraphControlScore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlCategory** | Pointer to **NullableString** | Control action category (Identity, Data, Device, Apps, Infrastructure). | [optional] 
**ControlName** | Pointer to **NullableString** | Control unique name. | [optional] 
**Description** | Pointer to **NullableString** | Description of the control. | [optional] 
**Score** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | Tenant achieved score for the control (it varies day by day depending on tenant operations on the control). | [optional] 

## Methods

### NewMicrosoftGraphControlScore

`func NewMicrosoftGraphControlScore() *MicrosoftGraphControlScore`

NewMicrosoftGraphControlScore instantiates a new MicrosoftGraphControlScore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphControlScoreWithDefaults

`func NewMicrosoftGraphControlScoreWithDefaults() *MicrosoftGraphControlScore`

NewMicrosoftGraphControlScoreWithDefaults instantiates a new MicrosoftGraphControlScore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlCategory

`func (o *MicrosoftGraphControlScore) GetControlCategory() string`

GetControlCategory returns the ControlCategory field if non-nil, zero value otherwise.

### GetControlCategoryOk

`func (o *MicrosoftGraphControlScore) GetControlCategoryOk() (*string, bool)`

GetControlCategoryOk returns a tuple with the ControlCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlCategory

`func (o *MicrosoftGraphControlScore) SetControlCategory(v string)`

SetControlCategory sets ControlCategory field to given value.

### HasControlCategory

`func (o *MicrosoftGraphControlScore) HasControlCategory() bool`

HasControlCategory returns a boolean if a field has been set.

### SetControlCategoryNil

`func (o *MicrosoftGraphControlScore) SetControlCategoryNil(b bool)`

 SetControlCategoryNil sets the value for ControlCategory to be an explicit nil

### UnsetControlCategory
`func (o *MicrosoftGraphControlScore) UnsetControlCategory()`

UnsetControlCategory ensures that no value is present for ControlCategory, not even an explicit nil
### GetControlName

`func (o *MicrosoftGraphControlScore) GetControlName() string`

GetControlName returns the ControlName field if non-nil, zero value otherwise.

### GetControlNameOk

`func (o *MicrosoftGraphControlScore) GetControlNameOk() (*string, bool)`

GetControlNameOk returns a tuple with the ControlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlName

`func (o *MicrosoftGraphControlScore) SetControlName(v string)`

SetControlName sets ControlName field to given value.

### HasControlName

`func (o *MicrosoftGraphControlScore) HasControlName() bool`

HasControlName returns a boolean if a field has been set.

### SetControlNameNil

`func (o *MicrosoftGraphControlScore) SetControlNameNil(b bool)`

 SetControlNameNil sets the value for ControlName to be an explicit nil

### UnsetControlName
`func (o *MicrosoftGraphControlScore) UnsetControlName()`

UnsetControlName ensures that no value is present for ControlName, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphControlScore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphControlScore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphControlScore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphControlScore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphControlScore) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphControlScore) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScore

`func (o *MicrosoftGraphControlScore) GetScore() AnyOfnumberstringstring`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *MicrosoftGraphControlScore) GetScoreOk() (*AnyOfnumberstringstring, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *MicrosoftGraphControlScore) SetScore(v AnyOfnumberstringstring)`

SetScore sets Score field to given value.

### HasScore

`func (o *MicrosoftGraphControlScore) HasScore() bool`

HasScore returns a boolean if a field has been set.

### SetScoreNil

`func (o *MicrosoftGraphControlScore) SetScoreNil(b bool)`

 SetScoreNil sets the value for Score to be an explicit nil

### UnsetScore
`func (o *MicrosoftGraphControlScore) UnsetScore()`

UnsetScore ensures that no value is present for Score, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


