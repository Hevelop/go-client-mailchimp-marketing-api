/*
Mailchimp Marketing API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0.55
Contact: apihelp@mailchimp.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mailchimpmarketingapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type TemplatesAPI interface {

	/*
	DeleteTemplatesId Delete template

	Delete a specific template.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param templateId The unique id for the template.
	@return TemplatesAPIDeleteTemplatesIdRequest
	*/
	DeleteTemplatesId(ctx context.Context, templateId string) TemplatesAPIDeleteTemplatesIdRequest

	// DeleteTemplatesIdExecute executes the request
	DeleteTemplatesIdExecute(r TemplatesAPIDeleteTemplatesIdRequest) (*http.Response, error)

	/*
	GetTemplates List templates

	Get a list of an account's available templates.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TemplatesAPIGetTemplatesRequest
	*/
	GetTemplates(ctx context.Context) TemplatesAPIGetTemplatesRequest

	// GetTemplatesExecute executes the request
	//  @return Templates
	GetTemplatesExecute(r TemplatesAPIGetTemplatesRequest) (*Templates, *http.Response, error)

	/*
	GetTemplatesId Get template info

	Get information about a specific template.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param templateId The unique id for the template.
	@return TemplatesAPIGetTemplatesIdRequest
	*/
	GetTemplatesId(ctx context.Context, templateId string) TemplatesAPIGetTemplatesIdRequest

	// GetTemplatesIdExecute executes the request
	//  @return TemplateInstance
	GetTemplatesIdExecute(r TemplatesAPIGetTemplatesIdRequest) (*TemplateInstance, *http.Response, error)

	/*
	GetTemplatesIdDefaultContent View default content

	Get the sections that you can edit in a template, including each section's default content.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param templateId The unique id for the template.
	@return TemplatesAPIGetTemplatesIdDefaultContentRequest
	*/
	GetTemplatesIdDefaultContent(ctx context.Context, templateId string) TemplatesAPIGetTemplatesIdDefaultContentRequest

	// GetTemplatesIdDefaultContentExecute executes the request
	//  @return TemplateDefaultContent
	GetTemplatesIdDefaultContentExecute(r TemplatesAPIGetTemplatesIdDefaultContentRequest) (*TemplateDefaultContent, *http.Response, error)

	/*
	PatchTemplatesId Update template

	Update the name, HTML, or `folder_id` of an existing template.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param templateId The unique id for the template.
	@return TemplatesAPIPatchTemplatesIdRequest
	*/
	PatchTemplatesId(ctx context.Context, templateId string) TemplatesAPIPatchTemplatesIdRequest

	// PatchTemplatesIdExecute executes the request
	//  @return TemplateInstance
	PatchTemplatesIdExecute(r TemplatesAPIPatchTemplatesIdRequest) (*TemplateInstance, *http.Response, error)

	/*
	PostTemplates Add template

	Create a new template for the account. Only Classic templates are supported.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TemplatesAPIPostTemplatesRequest
	*/
	PostTemplates(ctx context.Context) TemplatesAPIPostTemplatesRequest

	// PostTemplatesExecute executes the request
	//  @return TemplateInstance
	PostTemplatesExecute(r TemplatesAPIPostTemplatesRequest) (*TemplateInstance, *http.Response, error)
}

// TemplatesAPIService TemplatesAPI service
type TemplatesAPIService service

type TemplatesAPIDeleteTemplatesIdRequest struct {
	ctx context.Context
	ApiService TemplatesAPI
	templateId string
}

func (r TemplatesAPIDeleteTemplatesIdRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteTemplatesIdExecute(r)
}

/*
DeleteTemplatesId Delete template

Delete a specific template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId The unique id for the template.
 @return TemplatesAPIDeleteTemplatesIdRequest
*/
func (a *TemplatesAPIService) DeleteTemplatesId(ctx context.Context, templateId string) TemplatesAPIDeleteTemplatesIdRequest {
	return TemplatesAPIDeleteTemplatesIdRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
func (a *TemplatesAPIService) DeleteTemplatesIdExecute(r TemplatesAPIDeleteTemplatesIdRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesAPIService.DeleteTemplatesId")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/templates/{template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"template_id"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TemplatesAPIGetTemplatesRequest struct {
	ctx context.Context
	ApiService TemplatesAPI
	fields *[]string
	excludeFields *[]string
	count *int32
	offset *int32
	createdBy *string
	sinceDateCreated *string
	beforeDateCreated *string
	type_ *string
	category *string
	folderId *string
	sortField *string
	contentType *string
	sortDir *string
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r TemplatesAPIGetTemplatesRequest) Fields(fields []string) TemplatesAPIGetTemplatesRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r TemplatesAPIGetTemplatesRequest) ExcludeFields(excludeFields []string) TemplatesAPIGetTemplatesRequest {
	r.excludeFields = &excludeFields
	return r
}

// The number of records to return. Default value is 10. Maximum value is 1000
func (r TemplatesAPIGetTemplatesRequest) Count(count int32) TemplatesAPIGetTemplatesRequest {
	r.count = &count
	return r
}

// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0.
func (r TemplatesAPIGetTemplatesRequest) Offset(offset int32) TemplatesAPIGetTemplatesRequest {
	r.offset = &offset
	return r
}

// The Mailchimp account user who created the template.
func (r TemplatesAPIGetTemplatesRequest) CreatedBy(createdBy string) TemplatesAPIGetTemplatesRequest {
	r.createdBy = &createdBy
	return r
}

// Restrict the response to templates created after the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
func (r TemplatesAPIGetTemplatesRequest) SinceDateCreated(sinceDateCreated string) TemplatesAPIGetTemplatesRequest {
	r.sinceDateCreated = &sinceDateCreated
	return r
}

// Restrict the response to templates created before the set date. Uses ISO 8601 time format: 2015-10-21T15:41:36+00:00.
func (r TemplatesAPIGetTemplatesRequest) BeforeDateCreated(beforeDateCreated string) TemplatesAPIGetTemplatesRequest {
	r.beforeDateCreated = &beforeDateCreated
	return r
}

// Limit results based on template type.
func (r TemplatesAPIGetTemplatesRequest) Type_(type_ string) TemplatesAPIGetTemplatesRequest {
	r.type_ = &type_
	return r
}

// Limit results based on category.
func (r TemplatesAPIGetTemplatesRequest) Category(category string) TemplatesAPIGetTemplatesRequest {
	r.category = &category
	return r
}

// The unique folder id.
func (r TemplatesAPIGetTemplatesRequest) FolderId(folderId string) TemplatesAPIGetTemplatesRequest {
	r.folderId = &folderId
	return r
}

// Returns user templates sorted by the specified field.
func (r TemplatesAPIGetTemplatesRequest) SortField(sortField string) TemplatesAPIGetTemplatesRequest {
	r.sortField = &sortField
	return r
}

// Limit results based on how the template&#39;s content is put together. Only templates of type &#x60;user&#x60; can be filtered by &#x60;content_type&#x60;. If you want to retrieve saved templates created with the legacy email editor, then filter &#x60;content_type&#x60; to &#x60;template&#x60;. If you&#39;d rather pull your saved templates for the new editor, filter to &#x60;multichannel&#x60;. For code your own templates, filter to &#x60;html&#x60;.
func (r TemplatesAPIGetTemplatesRequest) ContentType(contentType string) TemplatesAPIGetTemplatesRequest {
	r.contentType = &contentType
	return r
}

// Determines the order direction for sorted results.
func (r TemplatesAPIGetTemplatesRequest) SortDir(sortDir string) TemplatesAPIGetTemplatesRequest {
	r.sortDir = &sortDir
	return r
}

func (r TemplatesAPIGetTemplatesRequest) Execute() (*Templates, *http.Response, error) {
	return r.ApiService.GetTemplatesExecute(r)
}

/*
GetTemplates List templates

Get a list of an account's available templates.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TemplatesAPIGetTemplatesRequest
*/
func (a *TemplatesAPIService) GetTemplates(ctx context.Context) TemplatesAPIGetTemplatesRequest {
	return TemplatesAPIGetTemplatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Templates
func (a *TemplatesAPIService) GetTemplatesExecute(r TemplatesAPIGetTemplatesRequest) (*Templates, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Templates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesAPIService.GetTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	if r.excludeFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", r.excludeFields, "csv")
	}
	if r.count != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "count", r.count, "")
	} else {
		var defaultValue int32 = 10
		r.count = &defaultValue
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
	} else {
		var defaultValue int32 = 0
		r.offset = &defaultValue
	}
	if r.createdBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "created_by", r.createdBy, "")
	}
	if r.sinceDateCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since_date_created", r.sinceDateCreated, "")
	}
	if r.beforeDateCreated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before_date_created", r.beforeDateCreated, "")
	}
	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.category != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "category", r.category, "")
	}
	if r.folderId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "folder_id", r.folderId, "")
	}
	if r.sortField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_field", r.sortField, "")
	}
	if r.contentType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "content_type", r.contentType, "")
	}
	if r.sortDir != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_dir", r.sortDir, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TemplatesAPIGetTemplatesIdRequest struct {
	ctx context.Context
	ApiService TemplatesAPI
	templateId string
	fields *[]string
	excludeFields *[]string
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r TemplatesAPIGetTemplatesIdRequest) Fields(fields []string) TemplatesAPIGetTemplatesIdRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r TemplatesAPIGetTemplatesIdRequest) ExcludeFields(excludeFields []string) TemplatesAPIGetTemplatesIdRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TemplatesAPIGetTemplatesIdRequest) Execute() (*TemplateInstance, *http.Response, error) {
	return r.ApiService.GetTemplatesIdExecute(r)
}

/*
GetTemplatesId Get template info

Get information about a specific template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId The unique id for the template.
 @return TemplatesAPIGetTemplatesIdRequest
*/
func (a *TemplatesAPIService) GetTemplatesId(ctx context.Context, templateId string) TemplatesAPIGetTemplatesIdRequest {
	return TemplatesAPIGetTemplatesIdRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return TemplateInstance
func (a *TemplatesAPIService) GetTemplatesIdExecute(r TemplatesAPIGetTemplatesIdRequest) (*TemplateInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesAPIService.GetTemplatesId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/templates/{template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"template_id"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	if r.excludeFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", r.excludeFields, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TemplatesAPIGetTemplatesIdDefaultContentRequest struct {
	ctx context.Context
	ApiService TemplatesAPI
	templateId string
	fields *[]string
	excludeFields *[]string
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r TemplatesAPIGetTemplatesIdDefaultContentRequest) Fields(fields []string) TemplatesAPIGetTemplatesIdDefaultContentRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r TemplatesAPIGetTemplatesIdDefaultContentRequest) ExcludeFields(excludeFields []string) TemplatesAPIGetTemplatesIdDefaultContentRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r TemplatesAPIGetTemplatesIdDefaultContentRequest) Execute() (*TemplateDefaultContent, *http.Response, error) {
	return r.ApiService.GetTemplatesIdDefaultContentExecute(r)
}

/*
GetTemplatesIdDefaultContent View default content

Get the sections that you can edit in a template, including each section's default content.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId The unique id for the template.
 @return TemplatesAPIGetTemplatesIdDefaultContentRequest
*/
func (a *TemplatesAPIService) GetTemplatesIdDefaultContent(ctx context.Context, templateId string) TemplatesAPIGetTemplatesIdDefaultContentRequest {
	return TemplatesAPIGetTemplatesIdDefaultContentRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return TemplateDefaultContent
func (a *TemplatesAPIService) GetTemplatesIdDefaultContentExecute(r TemplatesAPIGetTemplatesIdDefaultContentRequest) (*TemplateDefaultContent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateDefaultContent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesAPIService.GetTemplatesIdDefaultContent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/templates/{template_id}/default-content"
	localVarPath = strings.Replace(localVarPath, "{"+"template_id"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	if r.excludeFields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exclude_fields", r.excludeFields, "csv")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TemplatesAPIPatchTemplatesIdRequest struct {
	ctx context.Context
	ApiService TemplatesAPI
	templateId string
	body *TemplateInstance1
}

func (r TemplatesAPIPatchTemplatesIdRequest) Body(body TemplateInstance1) TemplatesAPIPatchTemplatesIdRequest {
	r.body = &body
	return r
}

func (r TemplatesAPIPatchTemplatesIdRequest) Execute() (*TemplateInstance, *http.Response, error) {
	return r.ApiService.PatchTemplatesIdExecute(r)
}

/*
PatchTemplatesId Update template

Update the name, HTML, or `folder_id` of an existing template.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param templateId The unique id for the template.
 @return TemplatesAPIPatchTemplatesIdRequest
*/
func (a *TemplatesAPIService) PatchTemplatesId(ctx context.Context, templateId string) TemplatesAPIPatchTemplatesIdRequest {
	return TemplatesAPIPatchTemplatesIdRequest{
		ApiService: a,
		ctx: ctx,
		templateId: templateId,
	}
}

// Execute executes the request
//  @return TemplateInstance
func (a *TemplatesAPIService) PatchTemplatesIdExecute(r TemplatesAPIPatchTemplatesIdRequest) (*TemplateInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesAPIService.PatchTemplatesId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/templates/{template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"template_id"+"}", url.PathEscape(parameterValueToString(r.templateId, "templateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type TemplatesAPIPostTemplatesRequest struct {
	ctx context.Context
	ApiService TemplatesAPI
	body *TemplateInstance1
}

func (r TemplatesAPIPostTemplatesRequest) Body(body TemplateInstance1) TemplatesAPIPostTemplatesRequest {
	r.body = &body
	return r
}

func (r TemplatesAPIPostTemplatesRequest) Execute() (*TemplateInstance, *http.Response, error) {
	return r.ApiService.PostTemplatesExecute(r)
}

/*
PostTemplates Add template

Create a new template for the account. Only Classic templates are supported.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TemplatesAPIPostTemplatesRequest
*/
func (a *TemplatesAPIService) PostTemplates(ctx context.Context) TemplatesAPIPostTemplatesRequest {
	return TemplatesAPIPostTemplatesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TemplateInstance
func (a *TemplatesAPIService) PostTemplatesExecute(r TemplatesAPIPostTemplatesRequest) (*TemplateInstance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TemplateInstance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TemplatesAPIService.PostTemplates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
			var v ProblemDetailDocument
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
