# PlannerTaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checklist** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The collection of checklist items on the task. | [optional] 
**Description** | Pointer to **NullableString** | Description of the task | [optional] 
**PreviewType** | Pointer to [**AnyOfmicrosoftGraphPlannerPreviewType**](anyOf&lt;microsoft.graph.plannerPreviewType&gt;.md) | This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task. | [optional] 
**References** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The collection of references on the task. | [optional] 

## Methods

### NewPlannerTaskDetails

`func NewPlannerTaskDetails() *PlannerTaskDetails`

NewPlannerTaskDetails instantiates a new PlannerTaskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlannerTaskDetailsWithDefaults

`func NewPlannerTaskDetailsWithDefaults() *PlannerTaskDetails`

NewPlannerTaskDetailsWithDefaults instantiates a new PlannerTaskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecklist

`func (o *PlannerTaskDetails) GetChecklist() AnyOfobject`

GetChecklist returns the Checklist field if non-nil, zero value otherwise.

### GetChecklistOk

`func (o *PlannerTaskDetails) GetChecklistOk() (*AnyOfobject, bool)`

GetChecklistOk returns a tuple with the Checklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklist

`func (o *PlannerTaskDetails) SetChecklist(v AnyOfobject)`

SetChecklist sets Checklist field to given value.

### HasChecklist

`func (o *PlannerTaskDetails) HasChecklist() bool`

HasChecklist returns a boolean if a field has been set.

### SetChecklistNil

`func (o *PlannerTaskDetails) SetChecklistNil(b bool)`

 SetChecklistNil sets the value for Checklist to be an explicit nil

### UnsetChecklist
`func (o *PlannerTaskDetails) UnsetChecklist()`

UnsetChecklist ensures that no value is present for Checklist, not even an explicit nil
### GetDescription

`func (o *PlannerTaskDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlannerTaskDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlannerTaskDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlannerTaskDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *PlannerTaskDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *PlannerTaskDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPreviewType

`func (o *PlannerTaskDetails) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType`

GetPreviewType returns the PreviewType field if non-nil, zero value otherwise.

### GetPreviewTypeOk

`func (o *PlannerTaskDetails) GetPreviewTypeOk() (*AnyOfmicrosoftGraphPlannerPreviewType, bool)`

GetPreviewTypeOk returns a tuple with the PreviewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewType

`func (o *PlannerTaskDetails) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType)`

SetPreviewType sets PreviewType field to given value.

### HasPreviewType

`func (o *PlannerTaskDetails) HasPreviewType() bool`

HasPreviewType returns a boolean if a field has been set.

### SetPreviewTypeNil

`func (o *PlannerTaskDetails) SetPreviewTypeNil(b bool)`

 SetPreviewTypeNil sets the value for PreviewType to be an explicit nil

### UnsetPreviewType
`func (o *PlannerTaskDetails) UnsetPreviewType()`

UnsetPreviewType ensures that no value is present for PreviewType, not even an explicit nil
### GetReferences

`func (o *PlannerTaskDetails) GetReferences() AnyOfobject`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *PlannerTaskDetails) GetReferencesOk() (*AnyOfobject, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *PlannerTaskDetails) SetReferences(v AnyOfobject)`

SetReferences sets References field to given value.

### HasReferences

`func (o *PlannerTaskDetails) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *PlannerTaskDetails) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *PlannerTaskDetails) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


