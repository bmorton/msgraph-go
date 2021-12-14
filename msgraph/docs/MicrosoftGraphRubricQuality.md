# MicrosoftGraphRubricQuality

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Criteria** | Pointer to [**[]AnyOfmicrosoftGraphRubricCriterion**](AnyOfmicrosoftGraphRubricCriterion.md) | The collection of criteria for this rubric quality. | [optional] 
**Description** | Pointer to [**AnyOfmicrosoftGraphEducationItemBody**](anyOf&lt;microsoft.graph.educationItemBody&gt;.md) | The description of this rubric quality. | [optional] 
**DisplayName** | Pointer to **NullableString** | The name of this rubric quality. | [optional] 
**QualityId** | Pointer to **NullableString** | The ID of this resource. | [optional] 
**Weight** | Pointer to [**AnyOfnumberstringstring**](anyOf&lt;number,string,string&gt;.md) | If present, a numerical weight for this quality.  Weights must add up to 100. | [optional] 

## Methods

### NewMicrosoftGraphRubricQuality

`func NewMicrosoftGraphRubricQuality() *MicrosoftGraphRubricQuality`

NewMicrosoftGraphRubricQuality instantiates a new MicrosoftGraphRubricQuality object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphRubricQualityWithDefaults

`func NewMicrosoftGraphRubricQualityWithDefaults() *MicrosoftGraphRubricQuality`

NewMicrosoftGraphRubricQualityWithDefaults instantiates a new MicrosoftGraphRubricQuality object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriteria

`func (o *MicrosoftGraphRubricQuality) GetCriteria() []*AnyOfmicrosoftGraphRubricCriterion`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *MicrosoftGraphRubricQuality) GetCriteriaOk() (*[]*AnyOfmicrosoftGraphRubricCriterion, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *MicrosoftGraphRubricQuality) SetCriteria(v []*AnyOfmicrosoftGraphRubricCriterion)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *MicrosoftGraphRubricQuality) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.

### GetDescription

`func (o *MicrosoftGraphRubricQuality) GetDescription() AnyOfmicrosoftGraphEducationItemBody`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphRubricQuality) GetDescriptionOk() (*AnyOfmicrosoftGraphEducationItemBody, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphRubricQuality) SetDescription(v AnyOfmicrosoftGraphEducationItemBody)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphRubricQuality) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphRubricQuality) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphRubricQuality) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDisplayName

`func (o *MicrosoftGraphRubricQuality) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *MicrosoftGraphRubricQuality) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *MicrosoftGraphRubricQuality) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *MicrosoftGraphRubricQuality) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *MicrosoftGraphRubricQuality) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *MicrosoftGraphRubricQuality) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetQualityId

`func (o *MicrosoftGraphRubricQuality) GetQualityId() string`

GetQualityId returns the QualityId field if non-nil, zero value otherwise.

### GetQualityIdOk

`func (o *MicrosoftGraphRubricQuality) GetQualityIdOk() (*string, bool)`

GetQualityIdOk returns a tuple with the QualityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityId

`func (o *MicrosoftGraphRubricQuality) SetQualityId(v string)`

SetQualityId sets QualityId field to given value.

### HasQualityId

`func (o *MicrosoftGraphRubricQuality) HasQualityId() bool`

HasQualityId returns a boolean if a field has been set.

### SetQualityIdNil

`func (o *MicrosoftGraphRubricQuality) SetQualityIdNil(b bool)`

 SetQualityIdNil sets the value for QualityId to be an explicit nil

### UnsetQualityId
`func (o *MicrosoftGraphRubricQuality) UnsetQualityId()`

UnsetQualityId ensures that no value is present for QualityId, not even an explicit nil
### GetWeight

`func (o *MicrosoftGraphRubricQuality) GetWeight() AnyOfnumberstringstring`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *MicrosoftGraphRubricQuality) GetWeightOk() (*AnyOfnumberstringstring, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *MicrosoftGraphRubricQuality) SetWeight(v AnyOfnumberstringstring)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *MicrosoftGraphRubricQuality) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *MicrosoftGraphRubricQuality) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *MicrosoftGraphRubricQuality) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


