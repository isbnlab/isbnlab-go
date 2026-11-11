# Ratings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**NullableRatingSummary**](RatingSummary.md) |  | [optional] 
**Counts** | Pointer to [**NullableRatingCounts**](RatingCounts.md) |  | [optional] 

## Methods

### NewRatings

`func NewRatings() *Ratings`

NewRatings instantiates a new Ratings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatingsWithDefaults

`func NewRatingsWithDefaults() *Ratings`

NewRatingsWithDefaults instantiates a new Ratings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *Ratings) GetSummary() RatingSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *Ratings) GetSummaryOk() (*RatingSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *Ratings) SetSummary(v RatingSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *Ratings) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### SetSummaryNil

`func (o *Ratings) SetSummaryNil(b bool)`

 SetSummaryNil sets the value for Summary to be an explicit nil

### UnsetSummary
`func (o *Ratings) UnsetSummary()`

UnsetSummary ensures that no value is present for Summary, not even an explicit nil
### GetCounts

`func (o *Ratings) GetCounts() RatingCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *Ratings) GetCountsOk() (*RatingCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *Ratings) SetCounts(v RatingCounts)`

SetCounts sets Counts field to given value.

### HasCounts

`func (o *Ratings) HasCounts() bool`

HasCounts returns a boolean if a field has been set.

### SetCountsNil

`func (o *Ratings) SetCountsNil(b bool)`

 SetCountsNil sets the value for Counts to be an explicit nil

### UnsetCounts
`func (o *Ratings) UnsetCounts()`

UnsetCounts ensures that no value is present for Counts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


