# MicrosoftGraphEducationAssignmentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**DistributeForStudentWork** | Pointer to **NullableBool** | Indicates whether this resource should be copied to each student submission for modification and submission. Required | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEducationResource**](anyOf&lt;microsoft.graph.educationResource&gt;.md) | Resource object that has been associated with this assignment. | [optional] 

## Methods

### NewMicrosoftGraphEducationAssignmentResource

`func NewMicrosoftGraphEducationAssignmentResource() *MicrosoftGraphEducationAssignmentResource`

NewMicrosoftGraphEducationAssignmentResource instantiates a new MicrosoftGraphEducationAssignmentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationAssignmentResourceWithDefaults

`func NewMicrosoftGraphEducationAssignmentResourceWithDefaults() *MicrosoftGraphEducationAssignmentResource`

NewMicrosoftGraphEducationAssignmentResourceWithDefaults instantiates a new MicrosoftGraphEducationAssignmentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEducationAssignmentResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationAssignmentResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEducationAssignmentResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEducationAssignmentResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDistributeForStudentWork

`func (o *MicrosoftGraphEducationAssignmentResource) GetDistributeForStudentWork() bool`

GetDistributeForStudentWork returns the DistributeForStudentWork field if non-nil, zero value otherwise.

### GetDistributeForStudentWorkOk

`func (o *MicrosoftGraphEducationAssignmentResource) GetDistributeForStudentWorkOk() (*bool, bool)`

GetDistributeForStudentWorkOk returns a tuple with the DistributeForStudentWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeForStudentWork

`func (o *MicrosoftGraphEducationAssignmentResource) SetDistributeForStudentWork(v bool)`

SetDistributeForStudentWork sets DistributeForStudentWork field to given value.

### HasDistributeForStudentWork

`func (o *MicrosoftGraphEducationAssignmentResource) HasDistributeForStudentWork() bool`

HasDistributeForStudentWork returns a boolean if a field has been set.

### SetDistributeForStudentWorkNil

`func (o *MicrosoftGraphEducationAssignmentResource) SetDistributeForStudentWorkNil(b bool)`

 SetDistributeForStudentWorkNil sets the value for DistributeForStudentWork to be an explicit nil

### UnsetDistributeForStudentWork
`func (o *MicrosoftGraphEducationAssignmentResource) UnsetDistributeForStudentWork()`

UnsetDistributeForStudentWork ensures that no value is present for DistributeForStudentWork, not even an explicit nil
### GetResource

`func (o *MicrosoftGraphEducationAssignmentResource) GetResource() AnyOfmicrosoftGraphEducationResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphEducationAssignmentResource) GetResourceOk() (*AnyOfmicrosoftGraphEducationResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphEducationAssignmentResource) SetResource(v AnyOfmicrosoftGraphEducationResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphEducationAssignmentResource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MicrosoftGraphEducationAssignmentResource) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MicrosoftGraphEducationAssignmentResource) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


