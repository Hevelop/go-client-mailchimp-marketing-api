# # GalleryFolder
An individual folder listed in the File Manager.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | The unique id for the folder.  | [optional] [readonly]
**Name**| **string** | The name of the folder.  | [optional]
**FileCount**| **int32** | The number of files in the folder.  | [optional] [readonly]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time a file was added to the File Manager in ISO 8601 format.  | [optional] [readonly]
**CreatedBy**| **string** | The username of the profile that created the folder.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

