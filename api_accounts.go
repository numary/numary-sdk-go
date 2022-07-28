/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type AccountsApi interface {

	/*
	AddMetadataToAccount Add metadata to an account.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @param address Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
	 @return ApiAddMetadataToAccountRequest
	*/
	AddMetadataToAccount(ctx _context.Context, ledger string, address string) ApiAddMetadataToAccountRequest

	// AddMetadataToAccountExecute executes the request
	AddMetadataToAccountExecute(r ApiAddMetadataToAccountRequest) (*_nethttp.Response, error)

	/*
	CountAccounts Count the accounts from a ledger.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiCountAccountsRequest
	*/
	CountAccounts(ctx _context.Context, ledger string) ApiCountAccountsRequest

	// CountAccountsExecute executes the request
	CountAccountsExecute(r ApiCountAccountsRequest) (*_nethttp.Response, error)

	/*
	GetAccount Get account by its address.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @param address Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
	 @return ApiGetAccountRequest
	*/
	GetAccount(ctx _context.Context, ledger string, address string) ApiGetAccountRequest

	// GetAccountExecute executes the request
	//  @return GetAccount200Response
	GetAccountExecute(r ApiGetAccountRequest) (GetAccount200Response, *_nethttp.Response, error)

	/*
	ListAccounts List accounts from a ledger.

	List accounts from a ledger, sorted by address in descending order.

	 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 @param ledger Name of the ledger.
	 @return ApiListAccountsRequest
	*/
	ListAccounts(ctx _context.Context, ledger string) ApiListAccountsRequest

	// ListAccountsExecute executes the request
	//  @return ListAccounts200Response
	ListAccountsExecute(r ApiListAccountsRequest) (ListAccounts200Response, *_nethttp.Response, error)
}

// AccountsApiService AccountsApi service
type AccountsApiService service

type ApiAddMetadataToAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	ledger string
	address string
	requestBody *map[string]interface{}
}

// metadata
func (r ApiAddMetadataToAccountRequest) RequestBody(requestBody map[string]interface{}) ApiAddMetadataToAccountRequest {
	r.requestBody = &requestBody
	return r
}

func (r ApiAddMetadataToAccountRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.AddMetadataToAccountExecute(r)
}

/*
AddMetadataToAccount Add metadata to an account.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @param address Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
 @return ApiAddMetadataToAccountRequest
*/
func (a *AccountsApiService) AddMetadataToAccount(ctx _context.Context, ledger string, address string) ApiAddMetadataToAccountRequest {
	return ApiAddMetadataToAccountRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
		address: address,
	}
}

// Execute executes the request
func (a *AccountsApiService) AddMetadataToAccountExecute(r ApiAddMetadataToAccountRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.AddMetadataToAccount")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/accounts/{address}/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", _neturl.PathEscape(parameterToString(r.address, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.requestBody == nil {
		return nil, reportError("requestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.requestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetAccount400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCountAccountsRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	ledger string
	address *string
	metadata *map[string]interface{}
}

// Filter accounts by address pattern (regular expression placed between ^ and $).
func (r ApiCountAccountsRequest) Address(address string) ApiCountAccountsRequest {
	r.address = &address
	return r
}
// Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
func (r ApiCountAccountsRequest) Metadata(metadata map[string]interface{}) ApiCountAccountsRequest {
	r.metadata = &metadata
	return r
}

func (r ApiCountAccountsRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.CountAccountsExecute(r)
}

/*
CountAccounts Count the accounts from a ledger.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiCountAccountsRequest
*/
func (a *AccountsApiService) CountAccounts(ctx _context.Context, ledger string) ApiCountAccountsRequest {
	return ApiCountAccountsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
func (a *AccountsApiService) CountAccountsExecute(r ApiCountAccountsRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodHead
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.CountAccounts")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.address != nil {
		localVarQueryParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.metadata != nil {
		localVarQueryParams.Add("metadata", parameterToString(*r.metadata, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetAccountRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	ledger string
	address string
}


func (r ApiGetAccountRequest) Execute() (GetAccount200Response, *_nethttp.Response, error) {
	return r.ApiService.GetAccountExecute(r)
}

/*
GetAccount Get account by its address.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @param address Exact address of the account. It must match the following regular expressions pattern: ``` ^\\w+(:\\w+)*$ ``` 
 @return ApiGetAccountRequest
*/
func (a *AccountsApiService) GetAccount(ctx _context.Context, ledger string, address string) ApiGetAccountRequest {
	return ApiGetAccountRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
		address: address,
	}
}

// Execute executes the request
//  @return GetAccount200Response
func (a *AccountsApiService) GetAccountExecute(r ApiGetAccountRequest) (GetAccount200Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  GetAccount200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.GetAccount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/accounts/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", _neturl.PathEscape(parameterToString(r.address, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v GetAccount400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAccountsRequest struct {
	ctx _context.Context
	ApiService AccountsApi
	ledger string
	pageSize *int32
	after *string
	address *string
	metadata *map[string]interface{}
	balance *int64
	balanceOperator *string
	paginationToken *string
}

// The maximum number of results to return per page
func (r ApiListAccountsRequest) PageSize(pageSize int32) ApiListAccountsRequest {
	r.pageSize = &pageSize
	return r
}
// Pagination cursor, will return accounts after given address, in descending order.
func (r ApiListAccountsRequest) After(after string) ApiListAccountsRequest {
	r.after = &after
	return r
}
// Filter accounts by address pattern (regular expression placed between ^ and $).
func (r ApiListAccountsRequest) Address(address string) ApiListAccountsRequest {
	r.address = &address
	return r
}
// Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
func (r ApiListAccountsRequest) Metadata(metadata map[string]interface{}) ApiListAccountsRequest {
	r.metadata = &metadata
	return r
}
// Filter accounts by their balance (default operator is gte)
func (r ApiListAccountsRequest) Balance(balance int64) ApiListAccountsRequest {
	r.balance = &balance
	return r
}
// Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, or equal
func (r ApiListAccountsRequest) BalanceOperator(balanceOperator string) ApiListAccountsRequest {
	r.balanceOperator = &balanceOperator
	return r
}
// Parameter used in pagination requests. Maximum page size is set to 15. Set to the value of next for the next page of results.  Set to the value of previous for the previous page of results. No other parameters can be set when the pagination token is set. 
func (r ApiListAccountsRequest) PaginationToken(paginationToken string) ApiListAccountsRequest {
	r.paginationToken = &paginationToken
	return r
}

func (r ApiListAccountsRequest) Execute() (ListAccounts200Response, *_nethttp.Response, error) {
	return r.ApiService.ListAccountsExecute(r)
}

/*
ListAccounts List accounts from a ledger.

List accounts from a ledger, sorted by address in descending order.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ledger Name of the ledger.
 @return ApiListAccountsRequest
*/
func (a *AccountsApiService) ListAccounts(ctx _context.Context, ledger string) ApiListAccountsRequest {
	return ApiListAccountsRequest{
		ApiService: a,
		ctx: ctx,
		ledger: ledger,
	}
}

// Execute executes the request
//  @return ListAccounts200Response
func (a *AccountsApiService) ListAccountsExecute(r ApiListAccountsRequest) (ListAccounts200Response, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  ListAccounts200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AccountsApiService.ListAccounts")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ledger}/accounts"
	localVarPath = strings.Replace(localVarPath, "{"+"ledger"+"}", _neturl.PathEscape(parameterToString(r.ledger, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.after != nil {
		localVarQueryParams.Add("after", parameterToString(*r.after, ""))
	}
	if r.address != nil {
		localVarQueryParams.Add("address", parameterToString(*r.address, ""))
	}
	if r.metadata != nil {
		localVarQueryParams.Add("metadata", parameterToString(*r.metadata, ""))
	}
	if r.balance != nil {
		localVarQueryParams.Add("balance", parameterToString(*r.balance, ""))
	}
	if r.balanceOperator != nil {
		localVarQueryParams.Add("balance_operator", parameterToString(*r.balanceOperator, ""))
	}
	if r.paginationToken != nil {
		localVarQueryParams.Add("pagination_token", parameterToString(*r.paginationToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ListAccounts400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
