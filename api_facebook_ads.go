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


type FacebookAdsAPI interface {

	/*
	GetAllFacebookAds List facebook ads

	Get list of Facebook ads.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FacebookAdsAPIGetAllFacebookAdsRequest
	*/
	GetAllFacebookAds(ctx context.Context) FacebookAdsAPIGetAllFacebookAdsRequest

	// GetAllFacebookAdsExecute executes the request
	//  @return GetAllFacebookAds200Response
	GetAllFacebookAdsExecute(r FacebookAdsAPIGetAllFacebookAdsRequest) (*GetAllFacebookAds200Response, *http.Response, error)

	/*
	GetFacebookAdsId Get facebook ad info

	Get details of a Facebook ad.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param outreachId The outreach id.
	@return FacebookAdsAPIGetFacebookAdsIdRequest
	*/
	GetFacebookAdsId(ctx context.Context, outreachId string) FacebookAdsAPIGetFacebookAdsIdRequest

	// GetFacebookAdsIdExecute executes the request
	//  @return GetFacebookAdsId200Response
	GetFacebookAdsIdExecute(r FacebookAdsAPIGetFacebookAdsIdRequest) (*GetFacebookAdsId200Response, *http.Response, error)
}

// FacebookAdsAPIService FacebookAdsAPI service
type FacebookAdsAPIService service

type FacebookAdsAPIGetAllFacebookAdsRequest struct {
	ctx context.Context
	ApiService FacebookAdsAPI
	fields *[]string
	excludeFields *[]string
	count *int32
	offset *int32
	sortField *string
	sortDir *string
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r FacebookAdsAPIGetAllFacebookAdsRequest) Fields(fields []string) FacebookAdsAPIGetAllFacebookAdsRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r FacebookAdsAPIGetAllFacebookAdsRequest) ExcludeFields(excludeFields []string) FacebookAdsAPIGetAllFacebookAdsRequest {
	r.excludeFields = &excludeFields
	return r
}

// The number of records to return. Default value is 10. Maximum value is 1000
func (r FacebookAdsAPIGetAllFacebookAdsRequest) Count(count int32) FacebookAdsAPIGetAllFacebookAdsRequest {
	r.count = &count
	return r
}

// Used for [pagination](https://mailchimp.com/developer/marketing/docs/methods-parameters/#pagination), this it the number of records from a collection to skip. Default value is 0.
func (r FacebookAdsAPIGetAllFacebookAdsRequest) Offset(offset int32) FacebookAdsAPIGetAllFacebookAdsRequest {
	r.offset = &offset
	return r
}

// Returns files sorted by the specified field.
func (r FacebookAdsAPIGetAllFacebookAdsRequest) SortField(sortField string) FacebookAdsAPIGetAllFacebookAdsRequest {
	r.sortField = &sortField
	return r
}

// Determines the order direction for sorted results.
func (r FacebookAdsAPIGetAllFacebookAdsRequest) SortDir(sortDir string) FacebookAdsAPIGetAllFacebookAdsRequest {
	r.sortDir = &sortDir
	return r
}

func (r FacebookAdsAPIGetAllFacebookAdsRequest) Execute() (*GetAllFacebookAds200Response, *http.Response, error) {
	return r.ApiService.GetAllFacebookAdsExecute(r)
}

/*
GetAllFacebookAds List facebook ads

Get list of Facebook ads.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FacebookAdsAPIGetAllFacebookAdsRequest
*/
func (a *FacebookAdsAPIService) GetAllFacebookAds(ctx context.Context) FacebookAdsAPIGetAllFacebookAdsRequest {
	return FacebookAdsAPIGetAllFacebookAdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAllFacebookAds200Response
func (a *FacebookAdsAPIService) GetAllFacebookAdsExecute(r FacebookAdsAPIGetAllFacebookAdsRequest) (*GetAllFacebookAds200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAllFacebookAds200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FacebookAdsAPIService.GetAllFacebookAds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/facebook-ads"

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
	if r.sortField != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_field", r.sortField, "")
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

type FacebookAdsAPIGetFacebookAdsIdRequest struct {
	ctx context.Context
	ApiService FacebookAdsAPI
	outreachId string
	fields *[]string
	excludeFields *[]string
}

// A comma-separated list of fields to return. Reference parameters of sub-objects with dot notation.
func (r FacebookAdsAPIGetFacebookAdsIdRequest) Fields(fields []string) FacebookAdsAPIGetFacebookAdsIdRequest {
	r.fields = &fields
	return r
}

// A comma-separated list of fields to exclude. Reference parameters of sub-objects with dot notation.
func (r FacebookAdsAPIGetFacebookAdsIdRequest) ExcludeFields(excludeFields []string) FacebookAdsAPIGetFacebookAdsIdRequest {
	r.excludeFields = &excludeFields
	return r
}

func (r FacebookAdsAPIGetFacebookAdsIdRequest) Execute() (*GetFacebookAdsId200Response, *http.Response, error) {
	return r.ApiService.GetFacebookAdsIdExecute(r)
}

/*
GetFacebookAdsId Get facebook ad info

Get details of a Facebook ad.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param outreachId The outreach id.
 @return FacebookAdsAPIGetFacebookAdsIdRequest
*/
func (a *FacebookAdsAPIService) GetFacebookAdsId(ctx context.Context, outreachId string) FacebookAdsAPIGetFacebookAdsIdRequest {
	return FacebookAdsAPIGetFacebookAdsIdRequest{
		ApiService: a,
		ctx: ctx,
		outreachId: outreachId,
	}
}

// Execute executes the request
//  @return GetFacebookAdsId200Response
func (a *FacebookAdsAPIService) GetFacebookAdsIdExecute(r FacebookAdsAPIGetFacebookAdsIdRequest) (*GetFacebookAdsId200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFacebookAdsId200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FacebookAdsAPIService.GetFacebookAdsId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/facebook-ads/{outreach_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"outreach_id"+"}", url.PathEscape(parameterValueToString(r.outreachId, "outreachId")), -1)

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