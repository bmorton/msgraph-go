# MicrosoftGraphRubricLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**AnyOfmicrosoftGraphEducationItemBody**](anyOf&lt;microsoft.graph.educationItemBody&gt;.md) | The description of this rubric level. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of this rubric level. | [optional] 
**Grading** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | Null if this is a no-points rubric; educationAssignmentPointsGradeType if it is a points rubric. | [optional] 
**LevelId** | Pointer to **NullableString** | The ID of this resource. | [optional] 

## Methods

### NewMicrosoftGraphRubricLevel

`func NewMicrosoftGraphRubricLevel() *MicrosoftGraphRubricLevel`

NewMicrosoftGraphRubricLevel instantiates a new MicrosoftGraphRubricLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRubricLevelWithDefaults

`func NewMicrosoftGraphRubricLevelWithDefaults() *MicrosoftGraphRubricLevel`

NewMicrosoftGraphRubricLevelWithDefaults instantiates a new MicrosoftGraphRubricLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MicrosoftGraphRubricLevel) GetDescription() AnyOfmicrosoftGraphEducationItemBody`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphRubricLevel) GetDescriptionOk() (*AnyOfmicrosoftGraphEducationItemBody, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphRubricLevel) SetDescription(v AnyOfmicrosoftGraphEducationItemBody)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphRubricLevel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphRubricLevel) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphRubricLevel) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphRubricLevel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphRubricLevel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphRubricLevel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphRubricLevel) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphRubricLevel) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphRubricLevel) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetGrading

`func (o *MicrosoftGraphRubricLevel) GetGrading() AnyOfobject`

GetGrading returns the Grading field if non-nil, zero value otherwise.

### GetGradingOk

`func (o *MicrosoftGraphRubricLevel) GetGradingOk() (*AnyOfobject, bool)`

GetGradingOk returns a tuple with the Grading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrading

`func (o *MicrosoftGraphRubricLevel) SetGrading(v AnyOfobject)`

SetGrading sets Grading field to given value.

### HasGrading

`func (o *MicrosoftGraphRubricLevel) HasGrading() bool`

HasGrading returns a boolean if a field has been set.

### SetGradingNil

`func (o *MicrosoftGraphRubricLevel) SetGradingNil(b bool)`

 SetGradingNil sets the value for Grading to be an explicit nil

### UnsetGrading
`func (o *MicrosoftGraphRubricLevel) UnsetGrading()`

UnsetGrading ensures that no value is present for Grading, not even an explicit nil
### GetLevelId

`func (o *MicrosoftGraphRubricLevel) GetLevelId() string`

GetLevelId returns the LevelId field if non-nil, zero value otherwise.

### GetLevelIdOk

`func (o *MicrosoftGraphRubricLevel) GetLevelIdOk() (*string, bool)`

GetLevelIdOk returns a tuple with the LevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelId

`func (o *MicrosoftGraphRubricLevel) SetLevelId(v string)`

SetLevelId sets LevelId field to given value.

### HasLevelId

`func (o *MicrosoftGraphRubricLevel) HasLevelId() bool`

HasLevelId returns a boolean if a field has been set.

### SetLevelIdNil

`func (o *MicrosoftGraphRubricLevel) SetLevelIdNil(b bool)`

 SetLevelIdNil sets the value for LevelId to be an explicit nil

### UnsetLevelId
`func (o *MicrosoftGraphRubricLevel) UnsetLevelId()`

UnsetLevelId ensures that no value is present for LevelId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


