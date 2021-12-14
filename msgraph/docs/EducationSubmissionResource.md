# EducationSubmissionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignmentResourceUrl** | Pointer to **NullableString** | Pointer to the assignment from which this resource was copied. If this is null, the student uploaded the resource. | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEducationResource**](anyOf&lt;microsoft.graph.educationResource&gt;.md) | Resource object. | [optional] 

## Methods

### NewEducationSubmissionResource

`func NewEducationSubmissionResource() *EducationSubmissionResource`

NewEducationSubmissionResource instantiates a new EducationSubmissionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEducationSubmissionResourceWithDefaults

`func NewEducationSubmissionResourceWithDefaults() *EducationSubmissionResource`

NewEducationSubmissionResourceWithDefaults instantiates a new EducationSubmissionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignmentResourceUrl

`func (o *EducationSubmissionResource) GetAssignmentResourceUrl() string`

GetAssignmentResourceUrl returns the AssignmentResourceUrl field if non-nil, zero value otherwise.

### GetAssignmentResourceUrlOk

`func (o *EducationSubmissionResource) GetAssignmentResourceUrlOk() (*string, bool)`

GetAssignmentResourceUrlOk returns a tuple with the AssignmentResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentResourceUrl

`func (o *EducationSubmissionResource) SetAssignmentResourceUrl(v string)`

SetAssignmentResourceUrl sets AssignmentResourceUrl field to given value.

### HasAssignmentResourceUrl

`func (o *EducationSubmissionResource) HasAssignmentResourceUrl() bool`

HasAssignmentResourceUrl returns a boolean if a field has been set.

### SetAssignmentResourceUrlNil

`func (o *EducationSubmissionResource) SetAssignmentResourceUrlNil(b bool)`

 SetAssignmentResourceUrlNil sets the value for AssignmentResourceUrl to be an explicit nil

### UnsetAssignmentResourceUrl
`func (o *EducationSubmissionResource) UnsetAssignmentResourceUrl()`

UnsetAssignmentResourceUrl ensures that no value is present for AssignmentResourceUrl, not even an explicit nil
### GetResource

`func (o *EducationSubmissionResource) GetResource() AnyOfmicrosoftGraphEducationResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EducationSubmissionResource) GetResourceOk() (*AnyOfmicrosoftGraphEducationResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EducationSubmissionResource) SetResource(v AnyOfmicrosoftGraphEducationResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *EducationSubmissionResource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *EducationSubmissionResource) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *EducationSubmissionResource) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


