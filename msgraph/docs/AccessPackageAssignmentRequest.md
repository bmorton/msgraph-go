# AccessPackageAssignmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletedDateTime** | Pointer to **NullableTime** |  | [optional] 
**CreatedDateTime** | Pointer to **NullableTime** | The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Read-only. | [optional] 
**RequestType** | Pointer to [**AnyOfmicrosoftGraphAccessPackageRequestType**](anyOf&lt;microsoft.graph.accessPackageRequestType&gt;.md) | The type of the request. The possible values are: notSpecified, userAdd, userUpdate, userRemove, adminAdd, adminUpdate, adminRemove, systemAdd, systemUpdate, systemRemove, onBehalfAdd, unknownFutureValue. A request from the user themselves would have requestType of UserAdd or UserRemove. This property cannot be changed once set. | [optional] 
**Schedule** | Pointer to [**AnyOfmicrosoftGraphEntitlementManagementSchedule**](anyOf&lt;microsoft.graph.entitlementManagementSchedule&gt;.md) | The range of dates that access is to be assigned to the requestor. This property cannot be changed once set. | [optional] 
**State** | Pointer to [**AnyOfmicrosoftGraphAccessPackageRequestState**](anyOf&lt;microsoft.graph.accessPackageRequestState&gt;.md) | The state of the request. The possible values are: submitted, pendingApproval, delivering, delivered, deliveryFailed, denied, scheduled, canceled, partiallyDelivered, unknownFutureValue. Read-only. | [optional] 
**Status** | Pointer to **NullableString** | More information on the request processing status. Read-only. | [optional] 
**AccessPackage** | Pointer to [**AnyOfmicrosoftGraphAccessPackage**](anyOf&lt;microsoft.graph.accessPackage&gt;.md) | The access package associated with the accessPackageAssignmentRequest. An access package defines the collections of resource roles and the policies for how one or more users can get access to those resources. Read-only. Nullable.  Supports $expand. | [optional] 
**Assignment** | Pointer to [**AnyOfmicrosoftGraphAccessPackageAssignment**](anyOf&lt;microsoft.graph.accessPackageAssignment&gt;.md) | For a requestType of UserAdd or AdminAdd, this is an access package assignment requested to be created.  For a requestType of UserRemove, AdminRemove or SystemRemove, this has the id property of an existing assignment to be removed.   Supports $expand. | [optional] 
**Requestor** | Pointer to [**AnyOfmicrosoftGraphAccessPackageSubject**](anyOf&lt;microsoft.graph.accessPackageSubject&gt;.md) | The subject who requested or, if a direct assignment, was assigned. Read-only. Nullable. Supports $expand. | [optional] 

## Methods

### NewAccessPackageAssignmentRequest

`func NewAccessPackageAssignmentRequest() *AccessPackageAssignmentRequest`

NewAccessPackageAssignmentRequest instantiates a new AccessPackageAssignmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessPackageAssignmentRequestWithDefaults

`func NewAccessPackageAssignmentRequestWithDefaults() *AccessPackageAssignmentRequest`

NewAccessPackageAssignmentRequestWithDefaults instantiates a new AccessPackageAssignmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletedDateTime

`func (o *AccessPackageAssignmentRequest) GetCompletedDateTime() time.Time`

GetCompletedDateTime returns the CompletedDateTime field if non-nil, zero value otherwise.

### GetCompletedDateTimeOk

`func (o *AccessPackageAssignmentRequest) GetCompletedDateTimeOk() (*time.Time, bool)`

GetCompletedDateTimeOk returns a tuple with the CompletedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDateTime

`func (o *AccessPackageAssignmentRequest) SetCompletedDateTime(v time.Time)`

SetCompletedDateTime sets CompletedDateTime field to given value.

### HasCompletedDateTime

`func (o *AccessPackageAssignmentRequest) HasCompletedDateTime() bool`

HasCompletedDateTime returns a boolean if a field has been set.

### SetCompletedDateTimeNil

`func (o *AccessPackageAssignmentRequest) SetCompletedDateTimeNil(b bool)`

 SetCompletedDateTimeNil sets the value for CompletedDateTime to be an explicit nil

### UnsetCompletedDateTime
`func (o *AccessPackageAssignmentRequest) UnsetCompletedDateTime()`

UnsetCompletedDateTime ensures that no value is present for CompletedDateTime, not even an explicit nil
### GetCreatedDateTime

`func (o *AccessPackageAssignmentRequest) GetCreatedDateTime() time.Time`

GetCreatedDateTime returns the CreatedDateTime field if non-nil, zero value otherwise.

### GetCreatedDateTimeOk

`func (o *AccessPackageAssignmentRequest) GetCreatedDateTimeOk() (*time.Time, bool)`

GetCreatedDateTimeOk returns a tuple with the CreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDateTime

`func (o *AccessPackageAssignmentRequest) SetCreatedDateTime(v time.Time)`

SetCreatedDateTime sets CreatedDateTime field to given value.

### HasCreatedDateTime

`func (o *AccessPackageAssignmentRequest) HasCreatedDateTime() bool`

HasCreatedDateTime returns a boolean if a field has been set.

### SetCreatedDateTimeNil

`func (o *AccessPackageAssignmentRequest) SetCreatedDateTimeNil(b bool)`

 SetCreatedDateTimeNil sets the value for CreatedDateTime to be an explicit nil

### UnsetCreatedDateTime
`func (o *AccessPackageAssignmentRequest) UnsetCreatedDateTime()`

UnsetCreatedDateTime ensures that no value is present for CreatedDateTime, not even an explicit nil
### GetRequestType

`func (o *AccessPackageAssignmentRequest) GetRequestType() AnyOfmicrosoftGraphAccessPackageRequestType`

GetRequestType returns the RequestType field if non-nil, zero value otherwise.

### GetRequestTypeOk

`func (o *AccessPackageAssignmentRequest) GetRequestTypeOk() (*AnyOfmicrosoftGraphAccessPackageRequestType, bool)`

GetRequestTypeOk returns a tuple with the RequestType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestType

`func (o *AccessPackageAssignmentRequest) SetRequestType(v AnyOfmicrosoftGraphAccessPackageRequestType)`

SetRequestType sets RequestType field to given value.

### HasRequestType

`func (o *AccessPackageAssignmentRequest) HasRequestType() bool`

HasRequestType returns a boolean if a field has been set.

### SetRequestTypeNil

`func (o *AccessPackageAssignmentRequest) SetRequestTypeNil(b bool)`

 SetRequestTypeNil sets the value for RequestType to be an explicit nil

### UnsetRequestType
`func (o *AccessPackageAssignmentRequest) UnsetRequestType()`

UnsetRequestType ensures that no value is present for RequestType, not even an explicit nil
### GetSchedule

`func (o *AccessPackageAssignmentRequest) GetSchedule() AnyOfmicrosoftGraphEntitlementManagementSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AccessPackageAssignmentRequest) GetScheduleOk() (*AnyOfmicrosoftGraphEntitlementManagementSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AccessPackageAssignmentRequest) SetSchedule(v AnyOfmicrosoftGraphEntitlementManagementSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AccessPackageAssignmentRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *AccessPackageAssignmentRequest) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *AccessPackageAssignmentRequest) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetState

`func (o *AccessPackageAssignmentRequest) GetState() AnyOfmicrosoftGraphAccessPackageRequestState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AccessPackageAssignmentRequest) GetStateOk() (*AnyOfmicrosoftGraphAccessPackageRequestState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AccessPackageAssignmentRequest) SetState(v AnyOfmicrosoftGraphAccessPackageRequestState)`

SetState sets State field to given value.

### HasState

`func (o *AccessPackageAssignmentRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *AccessPackageAssignmentRequest) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AccessPackageAssignmentRequest) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil
### GetStatus

`func (o *AccessPackageAssignmentRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessPackageAssignmentRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessPackageAssignmentRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccessPackageAssignmentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AccessPackageAssignmentRequest) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AccessPackageAssignmentRequest) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetAccessPackage

`func (o *AccessPackageAssignmentRequest) GetAccessPackage() AnyOfmicrosoftGraphAccessPackage`

GetAccessPackage returns the AccessPackage field if non-nil, zero value otherwise.

### GetAccessPackageOk

`func (o *AccessPackageAssignmentRequest) GetAccessPackageOk() (*AnyOfmicrosoftGraphAccessPackage, bool)`

GetAccessPackageOk returns a tuple with the AccessPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPackage

`func (o *AccessPackageAssignmentRequest) SetAccessPackage(v AnyOfmicrosoftGraphAccessPackage)`

SetAccessPackage sets AccessPackage field to given value.

### HasAccessPackage

`func (o *AccessPackageAssignmentRequest) HasAccessPackage() bool`

HasAccessPackage returns a boolean if a field has been set.

### SetAccessPackageNil

`func (o *AccessPackageAssignmentRequest) SetAccessPackageNil(b bool)`

 SetAccessPackageNil sets the value for AccessPackage to be an explicit nil

### UnsetAccessPackage
`func (o *AccessPackageAssignmentRequest) UnsetAccessPackage()`

UnsetAccessPackage ensures that no value is present for AccessPackage, not even an explicit nil
### GetAssignment

`func (o *AccessPackageAssignmentRequest) GetAssignment() AnyOfmicrosoftGraphAccessPackageAssignment`

GetAssignment returns the Assignment field if non-nil, zero value otherwise.

### GetAssignmentOk

`func (o *AccessPackageAssignmentRequest) GetAssignmentOk() (*AnyOfmicrosoftGraphAccessPackageAssignment, bool)`

GetAssignmentOk returns a tuple with the Assignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignment

`func (o *AccessPackageAssignmentRequest) SetAssignment(v AnyOfmicrosoftGraphAccessPackageAssignment)`

SetAssignment sets Assignment field to given value.

### HasAssignment

`func (o *AccessPackageAssignmentRequest) HasAssignment() bool`

HasAssignment returns a boolean if a field has been set.

### SetAssignmentNil

`func (o *AccessPackageAssignmentRequest) SetAssignmentNil(b bool)`

 SetAssignmentNil sets the value for Assignment to be an explicit nil

### UnsetAssignment
`func (o *AccessPackageAssignmentRequest) UnsetAssignment()`

UnsetAssignment ensures that no value is present for Assignment, not even an explicit nil
### GetRequestor

`func (o *AccessPackageAssignmentRequest) GetRequestor() AnyOfmicrosoftGraphAccessPackageSubject`

GetRequestor returns the Requestor field if non-nil, zero value otherwise.

### GetRequestorOk

`func (o *AccessPackageAssignmentRequest) GetRequestorOk() (*AnyOfmicrosoftGraphAccessPackageSubject, bool)`

GetRequestorOk returns a tuple with the Requestor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestor

`func (o *AccessPackageAssignmentRequest) SetRequestor(v AnyOfmicrosoftGraphAccessPackageSubject)`

SetRequestor sets Requestor field to given value.

### HasRequestor

`func (o *AccessPackageAssignmentRequest) HasRequestor() bool`

HasRequestor returns a boolean if a field has been set.

### SetRequestorNil

`func (o *AccessPackageAssignmentRequest) SetRequestorNil(b bool)`

 SetRequestorNil sets the value for Requestor to be an explicit nil

### UnsetRequestor
`func (o *AccessPackageAssignmentRequest) UnsetRequestor()`

UnsetRequestor ensures that no value is present for Requestor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


