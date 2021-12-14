# MicrosoftGraphPlannerTaskDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**Checklist** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The collection of checklist items on the task. | [optional] 
**Description** | Pointer to **NullableString** | Description of the task | [optional] 
**PreviewType** | Pointer to [**AnyOfmicrosoftGraphPlannerPreviewType**](anyOf&lt;microsoft.graph.plannerPreviewType&gt;.md) | This sets the type of preview that shows up on the task. The possible values are: automatic, noPreview, checklist, description, reference. When set to automatic the displayed preview is chosen by the app viewing the task. | [optional] 
**References** | Pointer to [**AnyOfobject**](anyOf&lt;object&gt;.md) | The collection of references on the task. | [optional] 

## Methods

### NewMicrosoftGraphPlannerTaskDetails

`func NewMicrosoftGraphPlannerTaskDetails() *MicrosoftGraphPlannerTaskDetails`

NewMicrosoftGraphPlannerTaskDetails instantiates a new MicrosoftGraphPlannerTaskDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphPlannerTaskDetailsWithDefaults

`func NewMicrosoftGraphPlannerTaskDetailsWithDefaults() *MicrosoftGraphPlannerTaskDetails`

NewMicrosoftGraphPlannerTaskDetailsWithDefaults instantiates a new MicrosoftGraphPlannerTaskDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphPlannerTaskDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphPlannerTaskDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphPlannerTaskDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetChecklist

`func (o *MicrosoftGraphPlannerTaskDetails) GetChecklist() AnyOfobject`

GetChecklist returns the Checklist field if non-nil, zero value otherwise.

### GetChecklistOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetChecklistOk() (*AnyOfobject, bool)`

GetChecklistOk returns a tuple with the Checklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklist

`func (o *MicrosoftGraphPlannerTaskDetails) SetChecklist(v AnyOfobject)`

SetChecklist sets Checklist field to given value.

### HasChecklist

`func (o *MicrosoftGraphPlannerTaskDetails) HasChecklist() bool`

HasChecklist returns a boolean if a field has been set.

### SetChecklistNil

`func (o *MicrosoftGraphPlannerTaskDetails) SetChecklistNil(b bool)`

 SetChecklistNil sets the value for Checklist to be an explicit nil

### UnsetChecklist
`func (o *MicrosoftGraphPlannerTaskDetails) UnsetChecklist()`

UnsetChecklist ensures that no value is present for Checklist, not even an explicit nil
### GetDescription

`func (o *MicrosoftGraphPlannerTaskDetails) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MicrosoftGraphPlannerTaskDetails) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MicrosoftGraphPlannerTaskDetails) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *MicrosoftGraphPlannerTaskDetails) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *MicrosoftGraphPlannerTaskDetails) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPreviewType

`func (o *MicrosoftGraphPlannerTaskDetails) GetPreviewType() AnyOfmicrosoftGraphPlannerPreviewType`

GetPreviewType returns the PreviewType field if non-nil, zero value otherwise.

### GetPreviewTypeOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetPreviewTypeOk() (*AnyOfmicrosoftGraphPlannerPreviewType, bool)`

GetPreviewTypeOk returns a tuple with the PreviewType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewType

`func (o *MicrosoftGraphPlannerTaskDetails) SetPreviewType(v AnyOfmicrosoftGraphPlannerPreviewType)`

SetPreviewType sets PreviewType field to given value.

### HasPreviewType

`func (o *MicrosoftGraphPlannerTaskDetails) HasPreviewType() bool`

HasPreviewType returns a boolean if a field has been set.

### SetPreviewTypeNil

`func (o *MicrosoftGraphPlannerTaskDetails) SetPreviewTypeNil(b bool)`

 SetPreviewTypeNil sets the value for PreviewType to be an explicit nil

### UnsetPreviewType
`func (o *MicrosoftGraphPlannerTaskDetails) UnsetPreviewType()`

UnsetPreviewType ensures that no value is present for PreviewType, not even an explicit nil
### GetReferences

`func (o *MicrosoftGraphPlannerTaskDetails) GetReferences() AnyOfobject`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *MicrosoftGraphPlannerTaskDetails) GetReferencesOk() (*AnyOfobject, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *MicrosoftGraphPlannerTaskDetails) SetReferences(v AnyOfobject)`

SetReferences sets References field to given value.

### HasReferences

`func (o *MicrosoftGraphPlannerTaskDetails) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### SetReferencesNil

`func (o *MicrosoftGraphPlannerTaskDetails) SetReferencesNil(b bool)`

 SetReferencesNil sets the value for References to be an explicit nil

### UnsetReferences
`func (o *MicrosoftGraphPlannerTaskDetails) UnsetReferences()`

UnsetReferences ensures that no value is present for References, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


