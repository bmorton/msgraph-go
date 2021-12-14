# MicrosoftGraphItemReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DriveId** | Pointer to **NullableString** | Unique identifier of the drive instance that contains the item. Read-only. | [optional] 
**DriveType** | Pointer to **NullableString** | Identifies the type of drive. See [drive][] resource for values. | [optional] 
**Id** | Pointer to **NullableString** | Unique identifier of the item in the drive. Read-only. | [optional] 
**Name** | Pointer to **NullableString** | The name of the item being referenced. Read-only. | [optional] 
**Path** | Pointer to **NullableString** | Path that can be used to navigate to the item. Read-only. | [optional] 
**ShareId** | Pointer to **NullableString** | A unique identifier for a shared resource that can be accessed via the [Shares][] API. | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) | Returns identifiers useful for SharePoint REST compatibility. Read-only. | [optional] 
**SiteId** | Pointer to **NullableString** | For OneDrive for Business and SharePoint, this property represents the ID of the site that contains the parent document library of the driveItem resource. The value is the same as the id property of that [site][] resource. It is an opaque string that consists of three identifiers of the site. For OneDrive, this property is not populated. | [optional] 

## Methods

### NewMicrosoftGraphItemReference

`func NewMicrosoftGraphItemReference() *MicrosoftGraphItemReference`

NewMicrosoftGraphItemReference instantiates a new MicrosoftGraphItemReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphItemReferenceWithDefaults

`func NewMicrosoftGraphItemReferenceWithDefaults() *MicrosoftGraphItemReference`

NewMicrosoftGraphItemReferenceWithDefaults instantiates a new MicrosoftGraphItemReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDriveId

`func (o *MicrosoftGraphItemReference) GetDriveId() string`

GetDriveId returns the DriveId field if non-nil, zero value otherwise.

### GetDriveIdOk

`func (o *MicrosoftGraphItemReference) GetDriveIdOk() (*string, bool)`

GetDriveIdOk returns a tuple with the DriveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveId

`func (o *MicrosoftGraphItemReference) SetDriveId(v string)`

SetDriveId sets DriveId field to given value.

### HasDriveId

`func (o *MicrosoftGraphItemReference) HasDriveId() bool`

HasDriveId returns a boolean if a field has been set.

### SetDriveIdNil

`func (o *MicrosoftGraphItemReference) SetDriveIdNil(b bool)`

 SetDriveIdNil sets the value for DriveId to be an explicit nil

### UnsetDriveId
`func (o *MicrosoftGraphItemReference) UnsetDriveId()`

UnsetDriveId ensures that no value is present for DriveId, not even an explicit nil
### GetDriveType

`func (o *MicrosoftGraphItemReference) GetDriveType() string`

GetDriveType returns the DriveType field if non-nil, zero value otherwise.

### GetDriveTypeOk

`func (o *MicrosoftGraphItemReference) GetDriveTypeOk() (*string, bool)`

GetDriveTypeOk returns a tuple with the DriveType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveType

`func (o *MicrosoftGraphItemReference) SetDriveType(v string)`

SetDriveType sets DriveType field to given value.

### HasDriveType

`func (o *MicrosoftGraphItemReference) HasDriveType() bool`

HasDriveType returns a boolean if a field has been set.

### SetDriveTypeNil

`func (o *MicrosoftGraphItemReference) SetDriveTypeNil(b bool)`

 SetDriveTypeNil sets the value for DriveType to be an explicit nil

### UnsetDriveType
`func (o *MicrosoftGraphItemReference) UnsetDriveType()`

UnsetDriveType ensures that no value is present for DriveType, not even an explicit nil
### GetId

`func (o *MicrosoftGraphItemReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MicrosoftGraphItemReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MicrosoftGraphItemReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MicrosoftGraphItemReference) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *MicrosoftGraphItemReference) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *MicrosoftGraphItemReference) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *MicrosoftGraphItemReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MicrosoftGraphItemReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MicrosoftGraphItemReference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MicrosoftGraphItemReference) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *MicrosoftGraphItemReference) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *MicrosoftGraphItemReference) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPath

`func (o *MicrosoftGraphItemReference) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MicrosoftGraphItemReference) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MicrosoftGraphItemReference) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MicrosoftGraphItemReference) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *MicrosoftGraphItemReference) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *MicrosoftGraphItemReference) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil
### GetShareId

`func (o *MicrosoftGraphItemReference) GetShareId() string`

GetShareId returns the ShareId field if non-nil, zero value otherwise.

### GetShareIdOk

`func (o *MicrosoftGraphItemReference) GetShareIdOk() (*string, bool)`

GetShareIdOk returns a tuple with the ShareId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareId

`func (o *MicrosoftGraphItemReference) SetShareId(v string)`

SetShareId sets ShareId field to given value.

### HasShareId

`func (o *MicrosoftGraphItemReference) HasShareId() bool`

HasShareId returns a boolean if a field has been set.

### SetShareIdNil

`func (o *MicrosoftGraphItemReference) SetShareIdNil(b bool)`

 SetShareIdNil sets the value for ShareId to be an explicit nil

### UnsetShareId
`func (o *MicrosoftGraphItemReference) UnsetShareId()`

UnsetShareId ensures that no value is present for ShareId, not even an explicit nil
### GetSharepointIds

`func (o *MicrosoftGraphItemReference) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *MicrosoftGraphItemReference) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointIds

`func (o *MicrosoftGraphItemReference) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds sets SharepointIds field to given value.

### HasSharepointIds

`func (o *MicrosoftGraphItemReference) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIdsNil

`func (o *MicrosoftGraphItemReference) SetSharepointIdsNil(b bool)`

 SetSharepointIdsNil sets the value for SharepointIds to be an explicit nil

### UnsetSharepointIds
`func (o *MicrosoftGraphItemReference) UnsetSharepointIds()`

UnsetSharepointIds ensures that no value is present for SharepointIds, not even an explicit nil
### GetSiteId

`func (o *MicrosoftGraphItemReference) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *MicrosoftGraphItemReference) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *MicrosoftGraphItemReference) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *MicrosoftGraphItemReference) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *MicrosoftGraphItemReference) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *MicrosoftGraphItemReference) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


