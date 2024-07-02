# # CreateAnAccountExport
Creates an account export with the given parameters.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeStages**| **[]string** | The stages of an account export to include. for more information please, see Model/array.php  |
**SinceTimestamp**| [**time.Time**](time.Time.md) | An ISO 8601 date that will limit the export to only records created after a given time. For instance, the reports stage will contain any campaign sent after the given timestamp. Audiences, however, are excluded from this limit.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

