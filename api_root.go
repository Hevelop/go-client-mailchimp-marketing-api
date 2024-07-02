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


type RootAPI interface {

	/*
	GetRoot List api root resources

	Get links to all other resources available in the API.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return RootAPIGetRootRequest
	*/
	GetRoot(ctx context.Context) RootAPIGetRootRequest

	// GetRootExecute executes the request
	//  @return APIRoot
	GetRootExecute(r RootAPIGetRootRequest) (*APIRoot, *http.Response, error)
}

// RootAPIService RootAPI service
type RootAPIService service

type RootAPIGetRootRequest struct {
	ctx context.Context
	ApiService RootAPI
	fields *[]string
	excludeFields *[]string
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r RootAPIGetRootRequest) Fields(fields []string) RootAPIGetRootRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r RootAPIGetRootRequest) ExcludeFields(excludeFields []string) RootAPIGetRootRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r RootAPIGetRootRequest) Execute() (*APIRoot, *http.Response, error) {
	return r.ApiService.GetRootExecute(r)
}

/*
GetRoot List api root resources

Get links to all other resources available in the API.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return RootAPIGetRootRequest
*/
func (a *RootAPIService) GetRoot(ctx context.Context) RootAPIGetRootRequest {
	return RootAPIGetRootRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return APIRoot
func (a *RootAPIService) GetRootExecute(r RootAPIGetRootRequest) (*APIRoot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *APIRoot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RootAPIService.GetRoot")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/"

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
