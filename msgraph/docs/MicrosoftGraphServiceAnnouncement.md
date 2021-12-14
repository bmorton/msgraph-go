# MicrosoftGraphServiceAnnouncement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Read-only. | [optional] 
**HealthOverviews** | Pointer to [**[]MicrosoftGraphServiceHealth**](MicrosoftGraphServiceHealth.md) | A collection of service health information for tenant. This property is a contained navigation property, it is nullable and readonly. | [optional] 
**Issues** | Pointer to [**[]MicrosoftGraphServiceHealthIssue**](MicrosoftGraphServiceHealthIssue.md) | A collection of service issues for tenant. This property is a contained navigation property, it is nullable and readonly. | [optional] 
**Messages** | Pointer to [**[]MicrosoftGraphServiceUpdateMessage**](MicrosoftGraphServiceUpdateMessage.md) | A collection of service messages for tenant. This property is a contained navigation property, it is nullable and readonly. | [optional] 

## Methods

### NewMicrosoftGraphServiceAnnouncement

`func NewMicrosoftGraphServiceAnnouncement() *MicrosoftGraphServiceAnnouncement`

NewMicrosoftGraphServiceAnnouncement instantiates a new MicrosoftGraphServiceAnnouncement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphServiceAnnouncementWithDefaults

`func NewMicrosoftGraphServiceAnnouncementWithDefaults() *MicrosoftGraphServiceAnnouncement`

NewMicrosoftGraphServiceAnnouncementWithDefaults instantiates a new MicrosoftGraphServiceAnnouncement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MicrosoftGraphServiceAnnouncement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphServiceAnnouncement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphServiceAnnouncement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphServiceAnnouncement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHealthOverviews

`func (o *MicrosoftGraphServiceAnnouncement) GetHealthOverviews() []MicrosoftGraphServiceHealth`

GetHealthOverviews returns the HealthOverviews field if non-nil, zero value otherwise.

### GetHealthOverviewsOk

`func (o *MicrosoftGraphServiceAnnouncement) GetHealthOverviewsOk() (*[]MicrosoftGraphServiceHealth, bool)`

GetHealthOverviewsOk returns a tuple with the HealthOverviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthOverviews

`func (o *MicrosoftGraphServiceAnnouncement) SetHealthOverviews(v []MicrosoftGraphServiceHealth)`

SetHealthOverviews sets HealthOverviews field to given value.

### HasHealthOverviews

`func (o *MicrosoftGraphServiceAnnouncement) HasHealthOverviews() bool`

HasHealthOverviews returns a boolean if a field has been set.

### GetIssues

`func (o *MicrosoftGraphServiceAnnouncement) GetIssues() []MicrosoftGraphServiceHealthIssue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *MicrosoftGraphServiceAnnouncement) GetIssuesOk() (*[]MicrosoftGraphServiceHealthIssue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *MicrosoftGraphServiceAnnouncement) SetIssues(v []MicrosoftGraphServiceHealthIssue)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *MicrosoftGraphServiceAnnouncement) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetMessages

`func (o *MicrosoftGraphServiceAnnouncement) GetMessages() []MicrosoftGraphServiceUpdateMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *MicrosoftGraphServiceAnnouncement) GetMessagesOk() (*[]MicrosoftGraphServiceUpdateMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *MicrosoftGraphServiceAnnouncement) SetMessages(v []MicrosoftGraphServiceUpdateMessage)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *MicrosoftGraphServiceAnnouncement) HasMessages() bool`

HasMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


