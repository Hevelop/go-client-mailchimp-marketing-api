# # CollectionAuthorization
Do particular authorization constraints around this collection limit creation of new instances?

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MayCreate**| **bool** | May the user create additional instances of this resource?  |
**MaxInstances**| **int32** | How many total instances of this resource are allowed? This is independent of any filter conditions applied to the query. As a special case, -1 indicates unlimited.  |
**CurrentTotalInstances**| **int32** | How many total instances of this resource are already in use? This is independent of any filter conditions applied to the query. Value may be larger than max_instances. As a special case, -1 is returned when access is unlimited.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

