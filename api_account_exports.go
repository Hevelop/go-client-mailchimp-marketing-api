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
)


type AccountExportsAPI interface {

	/*
	GetAccountExports List account exports

	Get a list of account exports for a given account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AccountExportsAPIGetAccountExportsRequest
	*/
	GetAccountExports(ctx context.Context) AccountExportsAPIGetAccountExportsRequest

	// GetAccountExportsExecute executes the request
	//  @return GetAccountExports200Response
	GetAccountExportsExecute(r AccountExportsAPIGetAccountExportsRequest) (*GetAccountExports200Response, *http.Response, error)

	/*
	PostAccountExport Add export

	Create a new account export in your Mailchimp account.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AccountExportsAPIPostAccountExportRequest
	*/
	PostAccountExport(ctx context.Context) AccountExportsAPIPostAccountExportRequest

	// PostAccountExportExecute executes the request
	//  @return AccountExportsInner
	PostAccountExportExecute(r AccountExportsAPIPostAccountExportRequest) (*AccountExportsInner, *http.Response, error)
}

// AccountExportsAPIService AccountExportsAPI service
type AccountExportsAPIService service

type AccountExportsAPIGetAccountExportsRequest struct {
	ctx context.Context
	ApiService AccountExportsAPI
	fields *[]string
	excludeFields *[]string
	count *int32
	offset *int32
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r AccountExportsAPIGetAccountExportsRequest) Fields(fields []string) AccountExportsAPIGetAccountExportsRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r AccountExportsAPIGetAccountExportsRequest) ExcludeFields(excludeFields []string) AccountExportsAPIGetAccountExportsRequest {
	r.excludeFields = &excludeFields
	return r
}

// The number of records to return. Default value is 10. Maximum value is 1000
func (r AccountExportsAPIGetAccountExportsRequest) Count(count int32) AccountExportsAPIGetAccountExportsRequest {
	r.count = &count
	return r
}

// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0.
func (r AccountExportsAPIGetAccountExportsRequest) Offset(offset int32) AccountExportsAPIGetAccountExportsRequest {
	r.offset = &offset
	return r
}

func (r AccountExportsAPIGetAccountExportsRequest) Execute() (*GetAccountExports200Response, *http.Response, error) {
	return r.ApiService.GetAccountExportsExecute(r)
}

/*
GetAccountExports List account exports

Get a list of account exports for a given account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AccountExportsAPIGetAccountExportsRequest
*/
func (a *AccountExportsAPIService) GetAccountExports(ctx context.Context) AccountExportsAPIGetAccountExportsRequest {
	return AccountExportsAPIGetAccountExportsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAccountExports200Response
func (a *AccountExportsAPIService) GetAccountExportsExecute(r AccountExportsAPIGetAccountExportsRequest) (*GetAccountExports200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAccountExports200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountExportsAPIService.GetAccountExports")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-exports"

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

type AccountExportsAPIPostAccountExportRequest struct {
	ctx context.Context
	ApiService AccountExportsAPI
	body *CreateAnAccountExport
}

func (r AccountExportsAPIPostAccountExportRequest) Body(body CreateAnAccountExport) AccountExportsAPIPostAccountExportRequest {
	r.body = &body
	return r
}

func (r AccountExportsAPIPostAccountExportRequest) Execute() (*AccountExportsInner, *http.Response, error) {
	return r.ApiService.PostAccountExportExecute(r)
}

/*
PostAccountExport Add export

Create a new account export in your Mailchimp account.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return AccountExportsAPIPostAccountExportRequest
*/
func (a *AccountExportsAPIService) PostAccountExport(ctx context.Context) AccountExportsAPIPostAccountExportRequest {
	return AccountExportsAPIPostAccountExportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AccountExportsInner
func (a *AccountExportsAPIService) PostAccountExportExecute(r AccountExportsAPIPostAccountExportRequest) (*AccountExportsInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountExportsInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountExportsAPIService.PostAccountExport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/account-exports"

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
