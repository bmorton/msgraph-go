# ThreatAssessmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**AnyOfmicrosoftGraphThreatCategory**](anyOf&lt;microsoft.graph.threatCategory&gt;.md) | The threat category. Possible values are: spam, phishing, malware. | [optional] 
**ContentType** | Pointer to [**AnyOfmicrosoftGraphThreatAssessmentContentType**](anyOf&lt;microsoft.graph.threatAssessmentContentType&gt;.md) | The content type of threat assessment. Possible values are: mail, url, file. | [optional] 
**CreatedBy** | Pointer to [**AnyOfmicrosoftGraphIdentitySet**](anyOf&lt;microsoft.graph.identitySet&gt;.md) | The threat assessment request creator. | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. | [optional] 
**ExpectedAssessment** | Pointer to [**AnyOfmicrosoftGraphThreatExpectedAssessment**](anyOf&lt;microsoft.graph.threatExpectedAssessment&gt;.md) | The expected assessment from submitter. Possible values are: block, unblock. | [optional] 
**RequestSource** | Pointer to [**AnyOfmicrosoftGraphThreatAssessmentRequestSource**](anyOf&lt;microsoft.graph.threatAssessmentRequestSource&gt;.md) | The source of the threat assessment request. Possible values are: administrator. | [optional] 
**Status** | Pointer to [**AnyOfmicrosoftGraphThreatAssessmentStatus**](anyOf&lt;microsoft.graph.threatAssessmentStatus&gt;.md) | The assessment process status. Possible values are: pending, completed. | [optional] 
**Results** | Pointer to [**[]MicrosoftGraphThreatAssessmentResult**](MicrosoftGraphThreatAssessmentResult.md) | A collection of threat assessment results. Read-only. By default, a GET /threatAssessmentRequests/{id} does not return this property unless you apply $expand on it. | [optional] 

## Methods

### NewThreatAssessmentRequest

`func NewThreatAssessmentRequest() *ThreatAssessmentRequest`

NewThreatAssessmentRequest instantiates a new ThreatAssessmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThreatAssessmentRequestWithDefaults

`func NewThreatAssessmentRequestWithDefaults() *ThreatAssessmentRequest`

NewThreatAssessmentRequestWithDefaults instantiates a new ThreatAssessmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ThreatAssessmentRequest) GetCategory() AnyOfmicrosoftGraphThreatCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ThreatAssessmentRequest) GetCategoryOk() (*AnyOfmicrosoftGraphThreatCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ThreatAssessmentRequest) SetCategory(v AnyOfmicrosoftGraphThreatCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ThreatAssessmentRequest) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ThreatAssessmentRequest) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ThreatAssessmentRequest) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetContentType

`func (o *ThreatAssessmentRequest) GetContentType() AnyOfmicrosoftGraphThreatAssessmentContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ThreatAssessmentRequest) GetContentTypeOk() (*AnyOfmicrosoftGraphThreatAssessmentContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ThreatAssessmentRequest) SetContentType(v AnyOfmicrosoftGraphThreatAssessmentContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ThreatAssessmentRequest) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ThreatAssessmentRequest) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ThreatAssessmentRequest) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetCreatedBy

`func (o *ThreatAssessmentRequest) GetCreatedBy() AnyOfmicrosoftGraphIdentitySet`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ThreatAssessmentRequest) GetCreatedByOk() (*AnyOfmicrosoftGraphIdentitySet, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ThreatAssessmentRequest) SetCreatedBy(v AnyOfmicrosoftGraphIdentitySet)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ThreatAssessmentRequest) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ThreatAssessmentRequest) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ThreatAssessmentRequest) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedDateTime

`func (o *ThreatAssessmentRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *ThreatAssessmentRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *ThreatAssessmentRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *ThreatAssessmentRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *ThreatAssessmentRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *ThreatAssessmentRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetExpectedAssessment

`func (o *ThreatAssessmentRequest) GetExpectedAssessment() AnyOfmicrosoftGraphThreatExpectedAssessment`

GetExpectedAssessment returns the ExpectedAssessment field if non-nil, zero value otherwise.

### GetExpectedAssessmentOk

`func (o *ThreatAssessmentRequest) GetExpectedAssessmentOk() (*AnyOfmicrosoftGraphThreatExpectedAssessment, bool)`

GetExpectedAssessmentOk returns a tuple with the ExpectedAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedAssessment

`func (o *ThreatAssessmentRequest) SetExpectedAssessment(v AnyOfmicrosoftGraphThreatExpectedAssessment)`

SetExpectedAssessment sets ExpectedAssessment field to given value.

### HasExpectedAssessment

`func (o *ThreatAssessmentRequest) HasExpectedAssessment() bool`

HasExpectedAssessment returns a boolean if a field has been set.

### SetExpectedAssessmentNil

`func (o *ThreatAssessmentRequest) SetExpectedAssessmentNil(b bool)`

 SetExpectedAssessmentNil sets the value for ExpectedAssessment to be an explicit nil

### UnsetExpectedAssessment
`func (o *ThreatAssessmentRequest) UnsetExpectedAssessment()`

UnsetExpectedAssessment ensures that no value is present for ExpectedAssessment, not even an explicit nil
### GetRequestSource

`func (o *ThreatAssessmentRequest) GetRequestSource() AnyOfmicrosoftGraphThreatAssessmentRequestSource`

GetRequestSource returns the RequestSource field if non-nil, zero value otherwise.

### GetRequestSourceOk

`func (o *ThreatAssessmentRequest) GetRequestSourceOk() (*AnyOfmicrosoftGraphThreatAssessmentRequestSource, bool)`

GetRequestSourceOk returns a tuple with the RequestSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSource

`func (o *ThreatAssessmentRequest) SetRequestSource(v AnyOfmicrosoftGraphThreatAssessmentRequestSource)`

SetRequestSource sets RequestSource field to given value.

### HasRequestSource

`func (o *ThreatAssessmentRequest) HasRequestSource() bool`

HasRequestSource returns a boolean if a field has been set.

### SetRequestSourceNil

`func (o *ThreatAssessmentRequest) SetRequestSourceNil(b bool)`

 SetRequestSourceNil sets the value for RequestSource to be an explicit nil

### UnsetRequestSource
`func (o *ThreatAssessmentRequest) UnsetRequestSource()`

UnsetRequestSource ensures that no value is present for RequestSource, not even an explicit nil
### GetStatus

`func (o *ThreatAssessmentRequest) GetStatus() AnyOfmicrosoftGraphThreatAssessmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ThreatAssessmentRequest) GetStatusOk() (*AnyOfmicrosoftGraphThreatAssessmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ThreatAssessmentRequest) SetStatus(v AnyOfmicrosoftGraphThreatAssessmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ThreatAssessmentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *ThreatAssessmentRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *ThreatAssessmentRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetResults

`func (o *ThreatAssessmentRequest) GetResults() []MicrosoftGraphThreatAssessmentResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ThreatAssessmentRequest) GetResultsOk() (*[]MicrosoftGraphThreatAssessmentResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ThreatAssessmentRequest) SetResults(v []MicrosoftGraphThreatAssessmentResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *ThreatAssessmentRequest) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


