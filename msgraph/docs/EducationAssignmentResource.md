# EducationAssignmentResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DistributeForStudentWork** | Pointer to **NullableBool** | Indicates whether this resource should be copied to each student submission for modification and submission. Required | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEducationResource**](anyOf&lt;microsoft.graph.educationResource&gt;.md) | Resource object that has been associated with this assignment. | [optional] 

## Methods

### NewEducationAssignmentResource

`func NewEducationAssignmentResource() *EducationAssignmentResource`

NewEducationAssignmentResource instantiates a new EducationAssignmentResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationAssignmentResourceWithDefaults

`func NewEducationAssignmentResourceWithDefaults() *EducationAssignmentResource`

NewEducationAssignmentResourceWithDefaults instantiates a new EducationAssignmentResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDistributeForStudentWork

`func (o *EducationAssignmentResource) GetDistributeForStudentWork() bool`

GetDistributeForStudentWork returns the DistributeForStudentWork field if non-nil, zero value otherwise.

### GetDistributeForStudentWorkOk

`func (o *EducationAssignmentResource) GetDistributeForStudentWorkOk() (*bool, bool)`

GetDistributeForStudentWorkOk returns a tuple with the DistributeForStudentWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeForStudentWork

`func (o *EducationAssignmentResource) SetDistributeForStudentWork(v bool)`

SetDistributeForStudentWork sets DistributeForStudentWork field to given value.

### HasDistributeForStudentWork

`func (o *EducationAssignmentResource) HasDistributeForStudentWork() bool`

HasDistributeForStudentWork returns a boolean if a field has been set.

### SetDistributeForStudentWorkNil

`func (o *EducationAssignmentResource) SetDistributeForStudentWorkNil(b bool)`

 SetDistributeForStudentWorkNil sets the value for DistributeForStudentWork to be an explicit nil

### UnsetDistributeForStudentWork
`func (o *EducationAssignmentResource) UnsetDistributeForStudentWork()`

UnsetDistributeForStudentWork ensures that no value is present for DistributeForStudentWork, not even an explicit nil
### GetResource

`func (o *EducationAssignmentResource) GetResource() AnyOfmicrosoftGraphEducationResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EducationAssignmentResource) GetResourceOk() (*AnyOfmicrosoftGraphEducationResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EducationAssignmentResource) SetResource(v AnyOfmicrosoftGraphEducationResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *EducationAssignmentResource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *EducationAssignmentResource) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *EducationAssignmentResource) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


