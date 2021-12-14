# ListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to [**AnyOfmicrosoftGraphContentTypeInfo**](anyOf&lt;microsoft.graph.contentTypeInfo&gt;.md) | The content type of this list item | [optional] 
**SharepointIds** | Pointer to [**AnyOfmicrosoftGraphSharepointIds**](anyOf&lt;microsoft.graph.sharepointIds&gt;.md) | Returns identifiers useful for SharePoint REST compatibility. Read-only. | [optional] 
**Analytics** | Pointer to [**AnyOfmicrosoftGraphItemAnalytics**](anyOf&lt;microsoft.graph.itemAnalytics&gt;.md) | Analytics about the view activities that took place on this item. | [optional] 
**DriveItem** | Pointer to [**AnyOfmicrosoftGraphDriveItem**](anyOf&lt;microsoft.graph.driveItem&gt;.md) | For document libraries, the driveItem relationship exposes the listItem as a [driveItem][] | [optional] 
**Fields** | Pointer to [**AnyOfmicrosoftGraphFieldValueSet**](anyOf&lt;microsoft.graph.fieldValueSet&gt;.md) | The values of the columns set on this list item. | [optional] 
**Versions** | Pointer to [**[]MicrosoftGraphListItemVersion**](MicrosoftGraphListItemVersion.md) | The list of previous versions of the list item. | [optional] 

## Methods

### NewListItem

`func NewListItem() *ListItem`

NewListItem instantiates a new ListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListItemWithDefaults

`func NewListItemWithDefaults() *ListItem`

NewListItemWithDefaults instantiates a new ListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *ListItem) GetContentType() AnyOfmicrosoftGraphContentTypeInfo`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ListItem) GetContentTypeOk() (*AnyOfmicrosoftGraphContentTypeInfo, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ListItem) SetContentType(v AnyOfmicrosoftGraphContentTypeInfo)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ListItem) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ListItem) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ListItem) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetSharepointIds

`func (o *ListItem) GetSharepointIds() AnyOfmicrosoftGraphSharepointIds`

GetSharepointIds returns the SharepointIds field if non-nil, zero value otherwise.

### GetSharepointIdsOk

`func (o *ListItem) GetSharepointIdsOk() (*AnyOfmicrosoftGraphSharepointIds, bool)`

GetSharepointIdsOk returns a tuple with the SharepointIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharepointIds

`func (o *ListItem) SetSharepointIds(v AnyOfmicrosoftGraphSharepointIds)`

SetSharepointIds sets SharepointIds field to given value.

### HasSharepointIds

`func (o *ListItem) HasSharepointIds() bool`

HasSharepointIds returns a boolean if a field has been set.

### SetSharepointIdsNil

`func (o *ListItem) SetSharepointIdsNil(b bool)`

 SetSharepointIdsNil sets the value for SharepointIds to be an explicit nil

### UnsetSharepointIds
`func (o *ListItem) UnsetSharepointIds()`

UnsetSharepointIds ensures that no value is present for SharepointIds, not even an explicit nil
### GetAnalytics

`func (o *ListItem) GetAnalytics() AnyOfmicrosoftGraphItemAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *ListItem) GetAnalyticsOk() (*AnyOfmicrosoftGraphItemAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *ListItem) SetAnalytics(v AnyOfmicrosoftGraphItemAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *ListItem) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalyticsNil

`func (o *ListItem) SetAnalyticsNil(b bool)`

 SetAnalyticsNil sets the value for Analytics to be an explicit nil

### UnsetAnalytics
`func (o *ListItem) UnsetAnalytics()`

UnsetAnalytics ensures that no value is present for Analytics, not even an explicit nil
### GetDriveItem

`func (o *ListItem) GetDriveItem() AnyOfmicrosoftGraphDriveItem`

GetDriveItem returns the DriveItem field if non-nil, zero value otherwise.

### GetDriveItemOk

`func (o *ListItem) GetDriveItemOk() (*AnyOfmicrosoftGraphDriveItem, bool)`

GetDriveItemOk returns a tuple with the DriveItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriveItem

`func (o *ListItem) SetDriveItem(v AnyOfmicrosoftGraphDriveItem)`

SetDriveItem sets DriveItem field to given value.

### HasDriveItem

`func (o *ListItem) HasDriveItem() bool`

HasDriveItem returns a boolean if a field has been set.

### SetDriveItemNil

`func (o *ListItem) SetDriveItemNil(b bool)`

 SetDriveItemNil sets the value for DriveItem to be an explicit nil

### UnsetDriveItem
`func (o *ListItem) UnsetDriveItem()`

UnsetDriveItem ensures that no value is present for DriveItem, not even an explicit nil
### GetFields

`func (o *ListItem) GetFields() AnyOfmicrosoftGraphFieldValueSet`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ListItem) GetFieldsOk() (*AnyOfmicrosoftGraphFieldValueSet, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ListItem) SetFields(v AnyOfmicrosoftGraphFieldValueSet)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ListItem) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *ListItem) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *ListItem) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetVersions

`func (o *ListItem) GetVersions() []MicrosoftGraphListItemVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListItem) GetVersionsOk() (*[]MicrosoftGraphListItemVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListItem) SetVersions(v []MicrosoftGraphListItemVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ListItem) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


