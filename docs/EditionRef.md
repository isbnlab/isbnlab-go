# EditionRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Isbn10** | Pointer to **NullableString** |  | [optional] 
**Isbn13** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Published** | Pointer to **NullableString** |  | [optional] 
**PhysicalFormat** | Pointer to **NullableString** |  | [optional] 
**Pages** | Pointer to **NullableInt64** |  | [optional] 
**Imprint** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewEditionRef

`func NewEditionRef() *EditionRef`

NewEditionRef instantiates a new EditionRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditionRefWithDefaults

`func NewEditionRefWithDefaults() *EditionRef`

NewEditionRefWithDefaults instantiates a new EditionRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsbn10

`func (o *EditionRef) GetIsbn10() string`

GetIsbn10 returns the Isbn10 field if non-nil, zero value otherwise.

### GetIsbn10Ok

`func (o *EditionRef) GetIsbn10Ok() (*string, bool)`

GetIsbn10Ok returns a tuple with the Isbn10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn10

`func (o *EditionRef) SetIsbn10(v string)`

SetIsbn10 sets Isbn10 field to given value.

### HasIsbn10

`func (o *EditionRef) HasIsbn10() bool`

HasIsbn10 returns a boolean if a field has been set.

### SetIsbn10Nil

`func (o *EditionRef) SetIsbn10Nil(b bool)`

 SetIsbn10Nil sets the value for Isbn10 to be an explicit nil

### UnsetIsbn10
`func (o *EditionRef) UnsetIsbn10()`

UnsetIsbn10 ensures that no value is present for Isbn10, not even an explicit nil
### GetIsbn13

`func (o *EditionRef) GetIsbn13() string`

GetIsbn13 returns the Isbn13 field if non-nil, zero value otherwise.

### GetIsbn13Ok

`func (o *EditionRef) GetIsbn13Ok() (*string, bool)`

GetIsbn13Ok returns a tuple with the Isbn13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsbn13

`func (o *EditionRef) SetIsbn13(v string)`

SetIsbn13 sets Isbn13 field to given value.

### HasIsbn13

`func (o *EditionRef) HasIsbn13() bool`

HasIsbn13 returns a boolean if a field has been set.

### SetIsbn13Nil

`func (o *EditionRef) SetIsbn13Nil(b bool)`

 SetIsbn13Nil sets the value for Isbn13 to be an explicit nil

### UnsetIsbn13
`func (o *EditionRef) UnsetIsbn13()`

UnsetIsbn13 ensures that no value is present for Isbn13, not even an explicit nil
### GetTitle

`func (o *EditionRef) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EditionRef) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EditionRef) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EditionRef) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *EditionRef) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *EditionRef) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetSlug

`func (o *EditionRef) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *EditionRef) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *EditionRef) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *EditionRef) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *EditionRef) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *EditionRef) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetPublished

`func (o *EditionRef) GetPublished() string`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *EditionRef) GetPublishedOk() (*string, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *EditionRef) SetPublished(v string)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *EditionRef) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *EditionRef) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *EditionRef) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil
### GetPhysicalFormat

`func (o *EditionRef) GetPhysicalFormat() string`

GetPhysicalFormat returns the PhysicalFormat field if non-nil, zero value otherwise.

### GetPhysicalFormatOk

`func (o *EditionRef) GetPhysicalFormatOk() (*string, bool)`

GetPhysicalFormatOk returns a tuple with the PhysicalFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalFormat

`func (o *EditionRef) SetPhysicalFormat(v string)`

SetPhysicalFormat sets PhysicalFormat field to given value.

### HasPhysicalFormat

`func (o *EditionRef) HasPhysicalFormat() bool`

HasPhysicalFormat returns a boolean if a field has been set.

### SetPhysicalFormatNil

`func (o *EditionRef) SetPhysicalFormatNil(b bool)`

 SetPhysicalFormatNil sets the value for PhysicalFormat to be an explicit nil

### UnsetPhysicalFormat
`func (o *EditionRef) UnsetPhysicalFormat()`

UnsetPhysicalFormat ensures that no value is present for PhysicalFormat, not even an explicit nil
### GetPages

`func (o *EditionRef) GetPages() int64`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *EditionRef) GetPagesOk() (*int64, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *EditionRef) SetPages(v int64)`

SetPages sets Pages field to given value.

### HasPages

`func (o *EditionRef) HasPages() bool`

HasPages returns a boolean if a field has been set.

### SetPagesNil

`func (o *EditionRef) SetPagesNil(b bool)`

 SetPagesNil sets the value for Pages to be an explicit nil

### UnsetPages
`func (o *EditionRef) UnsetPages()`

UnsetPages ensures that no value is present for Pages, not even an explicit nil
### GetImprint

`func (o *EditionRef) GetImprint() string`

GetImprint returns the Imprint field if non-nil, zero value otherwise.

### GetImprintOk

`func (o *EditionRef) GetImprintOk() (*string, bool)`

GetImprintOk returns a tuple with the Imprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImprint

`func (o *EditionRef) SetImprint(v string)`

SetImprint sets Imprint field to given value.

### HasImprint

`func (o *EditionRef) HasImprint() bool`

HasImprint returns a boolean if a field has been set.

### SetImprintNil

`func (o *EditionRef) SetImprintNil(b bool)`

 SetImprintNil sets the value for Imprint to be an explicit nil

### UnsetImprint
`func (o *EditionRef) UnsetImprint()`

UnsetImprint ensures that no value is present for Imprint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


