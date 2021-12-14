# MicrosoftGraphAccessReviewReviewerScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **NullableString** | The query specifying who will be the reviewer. See table for examples. | [optional] 
**QueryRoot** | Pointer to **NullableString** | In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative source of the query. This property is only required if a relative query, for example, ./manager, is specified. Possible value: decisions. | [optional] 
**QueryType** | Pointer to **NullableString** | The type of query. Examples include MicrosoftGraph and ARM. | [optional] 

## Methods

### NewMicrosoftGraphAccessReviewReviewerScope

`func NewMicrosoftGraphAccessReviewReviewerScope() *MicrosoftGraphAccessReviewReviewerScope`

NewMicrosoftGraphAccessReviewReviewerScope instantiates a new MicrosoftGraphAccessReviewReviewerScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMicrosoftGraphAccessReviewReviewerScopeWithDefaults

`func NewMicrosoftGraphAccessReviewReviewerScopeWithDefaults() *MicrosoftGraphAccessReviewReviewerScope`

NewMicrosoftGraphAccessReviewReviewerScopeWithDefaults instantiates a new MicrosoftGraphAccessReviewReviewerScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *MicrosoftGraphAccessReviewReviewerScope) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MicrosoftGraphAccessReviewReviewerScope) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *MicrosoftGraphAccessReviewReviewerScope) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *MicrosoftGraphAccessReviewReviewerScope) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetQueryRoot

`func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryRoot() string`

GetQueryRoot returns the QueryRoot field if non-nil, zero value otherwise.

### GetQueryRootOk

`func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryRootOk() (*string, bool)`

GetQueryRootOk returns a tuple with the QueryRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryRoot

`func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryRoot(v string)`

SetQueryRoot sets QueryRoot field to given value.

### HasQueryRoot

`func (o *MicrosoftGraphAccessReviewReviewerScope) HasQueryRoot() bool`

HasQueryRoot returns a boolean if a field has been set.

### SetQueryRootNil

`func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryRootNil(b bool)`

 SetQueryRootNil sets the value for QueryRoot to be an explicit nil

### UnsetQueryRoot
`func (o *MicrosoftGraphAccessReviewReviewerScope) UnsetQueryRoot()`

UnsetQueryRoot ensures that no value is present for QueryRoot, not even an explicit nil
### GetQueryType

`func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryType() string`

GetQueryType returns the QueryType field if non-nil, zero value otherwise.

### GetQueryTypeOk

`func (o *MicrosoftGraphAccessReviewReviewerScope) GetQueryTypeOk() (*string, bool)`

GetQueryTypeOk returns a tuple with the QueryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryType

`func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryType(v string)`

SetQueryType sets QueryType field to given value.

### HasQueryType

`func (o *MicrosoftGraphAccessReviewReviewerScope) HasQueryType() bool`

HasQueryType returns a boolean if a field has been set.

### SetQueryTypeNil

`func (o *MicrosoftGraphAccessReviewReviewerScope) SetQueryTypeNil(b bool)`

 SetQueryTypeNil sets the value for QueryType to be an explicit nil

### UnsetQueryType
`func (o *MicrosoftGraphAccessReviewReviewerScope) UnsetQueryType()`

UnsetQueryType ensures that no value is present for QueryType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


