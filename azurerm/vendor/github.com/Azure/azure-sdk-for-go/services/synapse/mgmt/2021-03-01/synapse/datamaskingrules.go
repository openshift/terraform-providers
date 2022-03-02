package synapse

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DataMaskingRulesClient is the azure Synapse Analytics Management Client
type DataMaskingRulesClient struct {
	BaseClient
}

// NewDataMaskingRulesClient creates an instance of the DataMaskingRulesClient client.
func NewDataMaskingRulesClient(subscriptionID string) DataMaskingRulesClient {
	return NewDataMaskingRulesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataMaskingRulesClientWithBaseURI creates an instance of the DataMaskingRulesClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewDataMaskingRulesClientWithBaseURI(baseURI string, subscriptionID string) DataMaskingRulesClient {
	return DataMaskingRulesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a Sql pool data masking rule.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// dataMaskingRuleName - the name of the data masking rule.
// parameters - the required parameters for creating or updating a data masking rule.
func (client DataMaskingRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, dataMaskingRuleName string, parameters DataMaskingRule) (result DataMaskingRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataMaskingRulesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.DataMaskingRuleProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.DataMaskingRuleProperties.SchemaName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.DataMaskingRuleProperties.TableName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.DataMaskingRuleProperties.ColumnName", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("synapse.DataMaskingRulesClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, dataMaskingRuleName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DataMaskingRulesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, dataMaskingRuleName string, parameters DataMaskingRule) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataMaskingPolicyName": autorest.Encode("path", "Default"),
		"dataMaskingRuleName":   autorest.Encode("path", dataMaskingRuleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlPoolName":           autorest.Encode("path", SQLPoolName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"workspaceName":         autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.Location = nil
	parameters.Kind = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules/{dataMaskingRuleName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DataMaskingRulesClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DataMaskingRulesClient) CreateOrUpdateResponder(resp *http.Response) (result DataMaskingRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the specific Sql pool data masking rule.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
// dataMaskingRuleName - the name of the data masking rule.
func (client DataMaskingRulesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, dataMaskingRuleName string) (result DataMaskingRule, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataMaskingRulesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.DataMaskingRulesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName, dataMaskingRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DataMaskingRulesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string, dataMaskingRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataMaskingPolicyName": autorest.Encode("path", "Default"),
		"dataMaskingRuleName":   autorest.Encode("path", dataMaskingRuleName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlPoolName":           autorest.Encode("path", SQLPoolName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"workspaceName":         autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules/{dataMaskingRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DataMaskingRulesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DataMaskingRulesClient) GetResponder(resp *http.Response) (result DataMaskingRule, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySQLPool gets a list of Sql pool data masking rules.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// workspaceName - the name of the workspace
// SQLPoolName - SQL pool name
func (client DataMaskingRulesClient) ListBySQLPool(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (result DataMaskingRuleListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataMaskingRulesClient.ListBySQLPool")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("synapse.DataMaskingRulesClient", "ListBySQLPool", err.Error())
	}

	req, err := client.ListBySQLPoolPreparer(ctx, resourceGroupName, workspaceName, SQLPoolName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "ListBySQLPool", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySQLPoolSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "ListBySQLPool", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySQLPoolResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "synapse.DataMaskingRulesClient", "ListBySQLPool", resp, "Failure responding to request")
		return
	}

	return
}

// ListBySQLPoolPreparer prepares the ListBySQLPool request.
func (client DataMaskingRulesClient) ListBySQLPoolPreparer(ctx context.Context, resourceGroupName string, workspaceName string, SQLPoolName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"dataMaskingPolicyName": autorest.Encode("path", "Default"),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"sqlPoolName":           autorest.Encode("path", SQLPoolName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
		"workspaceName":         autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2021-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sqlPools/{sqlPoolName}/dataMaskingPolicies/{dataMaskingPolicyName}/rules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySQLPoolSender sends the ListBySQLPool request. The method will close the
// http.Response Body if it receives an error.
func (client DataMaskingRulesClient) ListBySQLPoolSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySQLPoolResponder handles the response to the ListBySQLPool request. The method always
// closes the http.Response Body.
func (client DataMaskingRulesClient) ListBySQLPoolResponder(resp *http.Response) (result DataMaskingRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
