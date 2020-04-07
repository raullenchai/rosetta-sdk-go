// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Generated by: OpenAPI Generator (https://openapi-generator.tech)

package gen

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
)

// Linger please
var (
	_ _context.Context
)

// AccountAPIService AccountAPI service
type AccountAPIService service

// AccountBalance Get an array of all Account Balances for an Account Identifier and the Block
// Identifier at which the balance lookup was performed.  Some consumers of account balance data
// need to know at which block the balance was calculated to reconcile account balance changes.  To
// get all balances associated with an account, it may be necessary to perform multiple balance
// requests with unique Account Identifiers.  If the client supports it, passing nil
// AccountIdentifier metadata to the request should fetch all balances (if applicable).  It is also
// possible to perform a historical balance lookup (if the server supports it) by passing in an
// optional BlockIdentifier.
func (a *AccountAPIService) AccountBalance(
	ctx _context.Context,
	accountBalanceRequest AccountBalanceRequest,
) (*AccountBalanceResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod = _nethttp.MethodPost
		localVarPostBody   interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/account/balance"
	localVarHeaderParams := make(map[string]string)

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
	localVarPostBody = &accountBalanceRequest

	r, err := a.client.prepareRequest(
		ctx,
		localVarPath,
		localVarHTTPMethod,
		localVarPostBody,
		localVarHeaderParams,
	)
	if err != nil {
		return nil, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(ctx, r)
	if err != nil || localVarHTTPResponse == nil {
		return nil, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	defer localVarHTTPResponse.Body.Close()
	if err != nil {
		return nil, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode != 200 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return nil, localVarHTTPResponse, newErr
	}

	var v AccountBalanceResponse
	err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return nil, localVarHTTPResponse, newErr
	}

	return &v, localVarHTTPResponse, nil
}
