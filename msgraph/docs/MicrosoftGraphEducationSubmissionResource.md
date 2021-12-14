# MicrosoftGraphEducationSubmissionResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**AssignmentResourceUrl** | Pointer to **NullableString** | Pointer to the assignment from which this resource was copied. If this is null, the student uploaded the resource. | [optional] 
**Resource** | Pointer to [**AnyOfmicrosoftGraphEducationResource**](anyOf&lt;microsoft.graph.educationResource&gt;.md) | Resource object. | [optional] 

## Methods

### NewMicrosoftGraphEducationSubmissionResource

`func NewMicrosoftGraphEducationSubmissionResource() *MicrosoftGraphEducationSubmissionResource`

NewMicrosoftGraphEducationSubmissionResource instantiates a new MicrosoftGraphEducationSubmissionResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphEducationSubmissionResourceWithDefaults

`func NewMicrosoftGraphEducationSubmissionResourceWithDefaults() *MicrosoftGraphEducationSubmissionResource`

NewMicrosoftGraphEducationSubmissionResourceWithDefaults instantiates a new MicrosoftGraphEducationSubmissionResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphEducationSubmissionResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphEducationSubmissionResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphEducationSubmissionResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphEducationSubmissionResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAssignmentResourceUrl

`func (o *MicrosoftGraphEducationSubmissionResource) GetAssignmentResourceUrl() string`

GetAssignmentResourceUrl returns the AssignmentResourceUrl field if non-nil, zero value otherwise.

### GetAssignmentResourceUrlOk

`func (o *MicrosoftGraphEducationSubmissionResource) GetAssignmentResourceUrlOk() (*string, bool)`

GetAssignmentResourceUrlOk returns a tuple with the AssignmentResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignmentResourceUrl

`func (o *MicrosoftGraphEducationSubmissionResource) SetAssignmentResourceUrl(v string)`

SetAssignmentResourceUrl sets AssignmentResourceUrl field to given value.

### HasAssignmentResourceUrl

`func (o *MicrosoftGraphEducationSubmissionResource) HasAssignmentResourceUrl() bool`

HasAssignmentResourceUrl returns a boolean if a field has been set.

### SetAssignmentResourceUrlNil

`func (o *MicrosoftGraphEducationSubmissionResource) SetAssignmentResourceUrlNil(b bool)`

 SetAssignmentResourceUrlNil sets the value for AssignmentResourceUrl to be an explicit nil

### UnsetAssignmentResourceUrl
`func (o *MicrosoftGraphEducationSubmissionResource) UnsetAssignmentResourceUrl()`

UnsetAssignmentResourceUrl ensures that no value is present for AssignmentResourceUrl, not even an explicit nil
### GetResource

`func (o *MicrosoftGraphEducationSubmissionResource) GetResource() AnyOfmicrosoftGraphEducationResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *MicrosoftGraphEducationSubmissionResource) GetResourceOk() (*AnyOfmicrosoftGraphEducationResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *MicrosoftGraphEducationSubmissionResource) SetResource(v AnyOfmicrosoftGraphEducationResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *MicrosoftGraphEducationSubmissionResource) HasResource() bool`

HasResource returns a boolean if a field has been set.

### SetResourceNil

`func (o *MicrosoftGraphEducationSubmissionResource) SetResourceNil(b bool)`

 SetResourceNil sets the value for Resource to be an explicit nil

### UnsetResource
`func (o *MicrosoftGraphEducationSubmissionResource) UnsetResource()`

UnsetResource ensures that no value is present for Resource, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


