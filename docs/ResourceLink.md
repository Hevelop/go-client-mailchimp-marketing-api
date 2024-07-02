# # ResourceLink
This object represents a link from the resource where it is found to another resource or action that may be performed.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rel**| **string** | As with an HTML &#39;rel&#39; attribute, this describes the type of link.  | [optional] [readonly]
**Href**| **string** | This property contains a fully-qualified URL that can be called to retrieve the linked resource or perform the linked action.  | [optional] [readonly]
**Method**| **string** | The HTTP method that should be used when accessing the URL defined in &#39;href&#39;. for more information please, see Model/string.php  | [optional] [readonly]
**TargetSchema**| **string** | For GETs, this is a URL representing the schema that the response should conform to.  | [optional] [readonly]
**Schema**| **string** | For HTTP methods that can receive bodies (POST and PUT), this is a URL representing the schema that the body should conform to.  | [optional] [readonly]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

