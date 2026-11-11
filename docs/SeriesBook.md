# SeriesBook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isbn10** | Pointer to **NullableString** |  | [optional] 
**Isbn13** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Number** | Pointer to **NullableInt64** |  | [optional] 
**WorkId** | Pointer to **NullableInt64** |  | [optional] 
**WorkSlug** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSeriesBook

`func NewSeriesBook() *SeriesBook`

NewSeriesBook instantiates a new SeriesBook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesBookWithDefaults

`func NewSeriesBookWithDefaults() *SeriesBook`

NewSeriesBookWithDefaults instantiates a new SeriesBook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsbn10

`func (o *SeriesBook) GetIsbn10() string`

GetIsbn10 returns the Isbn10 field if non-nil, zero value otherwise.

### GetIsbn10Ok

`func (o *SeriesBook) GetIsbn10Ok() (*string, bool)`

GetIsbn10Ok returns a tuple with the Isbn10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn10

`func (o *SeriesBook) SetIsbn10(v string)`

SetIsbn10 sets Isbn10 field to given value.

### HasIsbn10

`func (o *SeriesBook) HasIsbn10() bool`

HasIsbn10 returns a boolean if a field has been set.

### SetIsbn10Nil

`func (o *SeriesBook) SetIsbn10Nil(b bool)`

 SetIsbn10Nil sets the value for Isbn10 to be an explicit nil

### UnsetIsbn10
`func (o *SeriesBook) UnsetIsbn10()`

UnsetIsbn10 ensures that no value is present for Isbn10, not even an explicit nil
### GetIsbn13

`func (o *SeriesBook) GetIsbn13() string`

GetIsbn13 returns the Isbn13 field if non-nil, zero value otherwise.

### GetIsbn13Ok

`func (o *SeriesBook) GetIsbn13Ok() (*string, bool)`

GetIsbn13Ok returns a tuple with the Isbn13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn13

`func (o *SeriesBook) SetIsbn13(v string)`

SetIsbn13 sets Isbn13 field to given value.

### HasIsbn13

`func (o *SeriesBook) HasIsbn13() bool`

HasIsbn13 returns a boolean if a field has been set.

### SetIsbn13Nil

`func (o *SeriesBook) SetIsbn13Nil(b bool)`

 SetIsbn13Nil sets the value for Isbn13 to be an explicit nil

### UnsetIsbn13
`func (o *SeriesBook) UnsetIsbn13()`

UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
### GetTitle

`func (o *SeriesBook) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SeriesBook) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SeriesBook) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SeriesBook) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SeriesBook) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SeriesBook) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSlug

`func (o *SeriesBook) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SeriesBook) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SeriesBook) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *SeriesBook) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *SeriesBook) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *SeriesBook) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetNumber

`func (o *SeriesBook) GetNumber() int64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SeriesBook) GetNumberOk() (*int64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SeriesBook) SetNumber(v int64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SeriesBook) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### SetNumberNil

`func (o *SeriesBook) SetNumberNil(b bool)`

 SetNumberNil sets the value for Number to be an explicit nil

### UnsetNumber
`func (o *SeriesBook) UnsetNumber()`

UnsetNumber ensures that no value is present for Number, not even an explicit nil
### GetWorkId

`func (o *SeriesBook) GetWorkId() int64`

GetWorkId returns the WorkId field if non-nil, zero value otherwise.

### GetWorkIdOk

`func (o *SeriesBook) GetWorkIdOk() (*int64, bool)`

GetWorkIdOk returns a tuple with the WorkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkId

`func (o *SeriesBook) SetWorkId(v int64)`

SetWorkId sets WorkId field to given value.

### HasWorkId

`func (o *SeriesBook) HasWorkId() bool`

HasWorkId returns a boolean if a field has been set.

### SetWorkIdNil

`func (o *SeriesBook) SetWorkIdNil(b bool)`

 SetWorkIdNil sets the value for WorkId to be an explicit nil

### UnsetWorkId
`func (o *SeriesBook) UnsetWorkId()`

UnsetWorkId ensures that no value is present for WorkId, not even an explicit nil
### GetWorkSlug

`func (o *SeriesBook) GetWorkSlug() string`

GetWorkSlug returns the WorkSlug field if non-nil, zero value otherwise.

### GetWorkSlugOk

`func (o *SeriesBook) GetWorkSlugOk() (*string, bool)`

GetWorkSlugOk returns a tuple with the WorkSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkSlug

`func (o *SeriesBook) SetWorkSlug(v string)`

SetWorkSlug sets WorkSlug field to given value.

### HasWorkSlug

`func (o *SeriesBook) HasWorkSlug() bool`

HasWorkSlug returns a boolean if a field has been set.

### SetWorkSlugNil

`func (o *SeriesBook) SetWorkSlugNil(b bool)`

 SetWorkSlugNil sets the value for WorkSlug to be an explicit nil

### UnsetWorkSlug
`func (o *SeriesBook) UnsetWorkSlug()`

UnsetWorkSlug ensures that no value is present for WorkSlug, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


