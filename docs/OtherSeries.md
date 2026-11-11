# OtherSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt64** |  | [optional] 
**Identifier** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Number** | Pointer to **NullableInt64** |  | [optional] 
**Kind** | Pointer to **NullableString** |  | [optional] 
**Parent** | Pointer to [**NullableSeriesParent**](SeriesParent.md) |  | [optional] 

## Methods

### NewOtherSeries

`func NewOtherSeries() *OtherSeries`

NewOtherSeries instantiates a new OtherSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtherSeriesWithDefaults

`func NewOtherSeriesWithDefaults() *OtherSeries`

NewOtherSeriesWithDefaults instantiates a new OtherSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OtherSeries) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OtherSeries) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OtherSeries) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OtherSeries) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OtherSeries) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OtherSeries) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetIdentifier

`func (o *OtherSeries) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *OtherSeries) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *OtherSeries) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *OtherSeries) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *OtherSeries) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *OtherSeries) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetSlug

`func (o *OtherSeries) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *OtherSeries) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *OtherSeries) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *OtherSeries) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *OtherSeries) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *OtherSeries) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *OtherSeries) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OtherSeries) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OtherSeries) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OtherSeries) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OtherSeries) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OtherSeries) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNumber

`func (o *OtherSeries) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OtherSeries) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OtherSeries) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OtherSeries) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *OtherSeries) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *OtherSeries) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetKind

`func (o *OtherSeries) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *OtherSeries) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *OtherSeries) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *OtherSeries) HasKind() bool`

HasKind returns a boolean if a field has been set.

### SetKindNil

`func (o *OtherSeries) SetKindNil(b bool)`

 SetKindNil sets the value for Kind to be an explicit nil

### UnsetKind
`func (o *OtherSeries) UnsetKind()`

UnsetKind ensures that no value is present for Kind, not even an explicit nil
### GetParent

`func (o *OtherSeries) GetParent() SeriesParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *OtherSeries) GetParentOk() (*SeriesParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *OtherSeries) SetParent(v SeriesParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *OtherSeries) HasParent() bool`

HasParent returns a boolean if a field has been set.

### SetParentNil

`func (o *OtherSeries) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *OtherSeries) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


