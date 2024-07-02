# # AccountExportsInner
An account export.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportId**| **int32** | The ID for the export.  | [optional]
**Started**| [**time.Time**](time.Time.md) | Start time for the export.  | [optional]
**Finished**| [**time.Time**](time.Time.md) | If finished, the finish time for the export.  | [optional]
**SizeInBytes**| **int32** | The size of the uncompressed export in bytes.  | [optional]
**DownloadUrl**| **string** | If the export is finished, the download URL for an export. URLs are only valid for 90 days after the export completes.  | [optional]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

