# # GalleryFile
An individual file listed in the File Manager.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | The unique id of the file.  | [optional] [readonly]
**FolderId**| **int32** | The id of the folder.  | [optional]
**Type**| **string** | The type of file in the File Manager. for more information please, see Model/string.php  | [optional] [readonly]
**Name**| **string** | The name of the file.  | [optional]
**FullSizeUrl**| **string** | The url of the full-size file.  | [optional] [readonly]
**ThumbnailUrl**| **string** | The url of the thumbnail preview.  | [optional] [readonly]
**Size**| **int32** | The size of the file in bytes.  | [optional] [readonly]
**CreatedAt**| [**time.Time**](time.Time.md) | The date and time a file was added to the File Manager in ISO 8601 format.  | [optional] [readonly]
**CreatedBy**| **string** | The username of the profile that uploaded the file.  | [optional] [readonly]
**Width**| **int32** | The width of the image.  | [optional] [readonly]
**Height**| **int32** | The height of an image.  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

