# # TemplateInstance
Information about a specific template.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | The individual id for the template.  | [optional] [readonly]
**Type**| **string** | The type of template (user, base, or gallery).  | [optional] [readonly]
**Name**| **string** | The name of the template.  | [optional]
**DragAndDrop**| **bool** | Whether the template uses the drag and drop editor.  | [optional] [readonly]
**Responsive**| **bool** | Whether the template contains media queries to make it responsive.  | [optional] [readonly]
**Category**| **string** | If available, the category the template is listed in.  | [optional] [readonly]
**DateCreated**| [**time.Time**](time.Time.md) | The date and time the template was created in ISO 8601 format.  | [optional] [readonly]
**DateEdited**| [**time.Time**](time.Time.md) | The date and time the template was edited in ISO 8601 format.  | [optional] [readonly]
**CreatedBy**| **string** | The login name for template&#39;s creator.  | [optional] [readonly]
**EditedBy**| **string** | The login name who last edited the template.  | [optional] [readonly]
**Active**| **bool** | User templates are not &#39;deleted,&#39; but rather marked as &#39;inactive.&#39; Returns whether the template is still active.  | [optional] [readonly]
**FolderId**| **string** | The id of the folder the template is currently in.  | [optional]
**Thumbnail**| **string** | If available, the URL for a thumbnail of the template.  | [optional] [readonly]
**ShareUrl**| **string** | The URL used for [template sharing](https://mailchimp.com/help/share-a-template/).  | [optional] [readonly]
**ContentType**| **string** | How the template&#39;s content is put together. for more information please, see Model/string.php  | [optional] [readonly]
**Links**| [**[]ResourceLink**](ResourceLink.md) | A list of link types and descriptions for the API schema documents.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

