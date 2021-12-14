# MicrosoftGraphSubjectRightsRequestDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExcludedItemCount** | Pointer to **NullableInt64** | Count of items that are excluded from the request. | [optional] 
**InsightCounts** | Pointer to [**[]AnyOfmicrosoftGraphKeyValuePair**](AnyOfmicrosoftGraphKeyValuePair.md) | Count of items per insight. | [optional] 
**ItemCount** | Pointer to **NullableInt64** | Count of items found. | [optional] 
**ItemNeedReview** | Pointer to **NullableInt64** | Count of item that need review. | [optional] 
**ProductItemCounts** | Pointer to [**[]AnyOfmicrosoftGraphKeyValuePair**](AnyOfmicrosoftGraphKeyValuePair.md) | Count of items per product, such as Exchange, SharePoint, OneDrive, and Teams. | [optional] 
**SignedOffItemCount** | Pointer to **NullableInt64** | Count of items signed off by the administrator. | [optional] 
**TotalItemSize** | Pointer to **NullableInt64** | Total item size in bytes. | [optional] 

## Methods

### NewMicrosoftGraphSubjectRightsRequestDetail

`func NewMicrosoftGraphSubjectRightsRequestDetail() *MicrosoftGraphSubjectRightsRequestDetail`

NewMicrosoftGraphSubjectRightsRequestDetail instantiates a new MicrosoftGraphSubjectRightsRequestDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphSubjectRightsRequestDetailWithDefaults

`func NewMicrosoftGraphSubjectRightsRequestDetailWithDefaults() *MicrosoftGraphSubjectRightsRequestDetail`

NewMicrosoftGraphSubjectRightsRequestDetailWithDefaults instantiates a new MicrosoftGraphSubjectRightsRequestDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExcludedItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetExcludedItemCount() int64`

GetExcludedItemCount returns the ExcludedItemCount field if non-nil, zero value otherwise.

### GetExcludedItemCountOk

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetExcludedItemCountOk() (*int64, bool)`

GetExcludedItemCountOk returns a tuple with the ExcludedItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetExcludedItemCount(v int64)`

SetExcludedItemCount sets ExcludedItemCount field to given value.

### HasExcludedItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) HasExcludedItemCount() bool`

HasExcludedItemCount returns a boolean if a field has been set.

### SetExcludedItemCountNil

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetExcludedItemCountNil(b bool)`

 SetExcludedItemCountNil sets the value for ExcludedItemCount to be an explicit nil

### UnsetExcludedItemCount
`func (o *MicrosoftGraphSubjectRightsRequestDetail) UnsetExcludedItemCount()`

UnsetExcludedItemCount ensures that no value is present for ExcludedItemCount, not even an explicit nil
### GetInsightCounts

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetInsightCounts() []*AnyOfmicrosoftGraphKeyValuePair`

GetInsightCounts returns the InsightCounts field if non-nil, zero value otherwise.

### GetInsightCountsOk

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetInsightCountsOk() (*[]*AnyOfmicrosoftGraphKeyValuePair, bool)`

GetInsightCountsOk returns a tuple with the InsightCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCounts

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetInsightCounts(v []*AnyOfmicrosoftGraphKeyValuePair)`

SetInsightCounts sets InsightCounts field to given value.

### HasInsightCounts

`func (o *MicrosoftGraphSubjectRightsRequestDetail) HasInsightCounts() bool`

HasInsightCounts returns a boolean if a field has been set.

### GetItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetItemCount() int64`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetItemCountOk() (*int64, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetItemCount(v int64)`

SetItemCount sets ItemCount field to given value.

### HasItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) HasItemCount() bool`

HasItemCount returns a boolean if a field has been set.

### SetItemCountNil

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetItemCountNil(b bool)`

 SetItemCountNil sets the value for ItemCount to be an explicit nil

### UnsetItemCount
`func (o *MicrosoftGraphSubjectRightsRequestDetail) UnsetItemCount()`

UnsetItemCount ensures that no value is present for ItemCount, not even an explicit nil
### GetItemNeedReview

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetItemNeedReview() int64`

GetItemNeedReview returns the ItemNeedReview field if non-nil, zero value otherwise.

### GetItemNeedReviewOk

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetItemNeedReviewOk() (*int64, bool)`

GetItemNeedReviewOk returns a tuple with the ItemNeedReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemNeedReview

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetItemNeedReview(v int64)`

SetItemNeedReview sets ItemNeedReview field to given value.

### HasItemNeedReview

`func (o *MicrosoftGraphSubjectRightsRequestDetail) HasItemNeedReview() bool`

HasItemNeedReview returns a boolean if a field has been set.

### SetItemNeedReviewNil

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetItemNeedReviewNil(b bool)`

 SetItemNeedReviewNil sets the value for ItemNeedReview to be an explicit nil

### UnsetItemNeedReview
`func (o *MicrosoftGraphSubjectRightsRequestDetail) UnsetItemNeedReview()`

UnsetItemNeedReview ensures that no value is present for ItemNeedReview, not even an explicit nil
### GetProductItemCounts

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetProductItemCounts() []*AnyOfmicrosoftGraphKeyValuePair`

GetProductItemCounts returns the ProductItemCounts field if non-nil, zero value otherwise.

### GetProductItemCountsOk

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetProductItemCountsOk() (*[]*AnyOfmicrosoftGraphKeyValuePair, bool)`

GetProductItemCountsOk returns a tuple with the ProductItemCounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductItemCounts

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetProductItemCounts(v []*AnyOfmicrosoftGraphKeyValuePair)`

SetProductItemCounts sets ProductItemCounts field to given value.

### HasProductItemCounts

`func (o *MicrosoftGraphSubjectRightsRequestDetail) HasProductItemCounts() bool`

HasProductItemCounts returns a boolean if a field has been set.

### GetSignedOffItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetSignedOffItemCount() int64`

GetSignedOffItemCount returns the SignedOffItemCount field if non-nil, zero value otherwise.

### GetSignedOffItemCountOk

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetSignedOffItemCountOk() (*int64, bool)`

GetSignedOffItemCountOk returns a tuple with the SignedOffItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedOffItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetSignedOffItemCount(v int64)`

SetSignedOffItemCount sets SignedOffItemCount field to given value.

### HasSignedOffItemCount

`func (o *MicrosoftGraphSubjectRightsRequestDetail) HasSignedOffItemCount() bool`

HasSignedOffItemCount returns a boolean if a field has been set.

### SetSignedOffItemCountNil

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetSignedOffItemCountNil(b bool)`

 SetSignedOffItemCountNil sets the value for SignedOffItemCount to be an explicit nil

### UnsetSignedOffItemCount
`func (o *MicrosoftGraphSubjectRightsRequestDetail) UnsetSignedOffItemCount()`

UnsetSignedOffItemCount ensures that no value is present for SignedOffItemCount, not even an explicit nil
### GetTotalItemSize

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetTotalItemSize() int64`

GetTotalItemSize returns the TotalItemSize field if non-nil, zero value otherwise.

### GetTotalItemSizeOk

`func (o *MicrosoftGraphSubjectRightsRequestDetail) GetTotalItemSizeOk() (*int64, bool)`

GetTotalItemSizeOk returns a tuple with the TotalItemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalItemSize

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetTotalItemSize(v int64)`

SetTotalItemSize sets TotalItemSize field to given value.

### HasTotalItemSize

`func (o *MicrosoftGraphSubjectRightsRequestDetail) HasTotalItemSize() bool`

HasTotalItemSize returns a boolean if a field has been set.

### SetTotalItemSizeNil

`func (o *MicrosoftGraphSubjectRightsRequestDetail) SetTotalItemSizeNil(b bool)`

 SetTotalItemSizeNil sets the value for TotalItemSize to be an explicit nil

### UnsetTotalItemSize
`func (o *MicrosoftGraphSubjectRightsRequestDetail) UnsetTotalItemSize()`

UnsetTotalItemSize ensures that no value is present for TotalItemSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


