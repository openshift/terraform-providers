package backup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ItemLevelRecoveryConnectionsClient is the open API 2.0 Specs for Azure RecoveryServices Backup service
type ItemLevelRecoveryConnectionsClient struct {
	BaseClient
}

// NewItemLevelRecoveryConnectionsClient creates an instance of the ItemLevelRecoveryConnectionsClient client.
func NewItemLevelRecoveryConnectionsClient(subscriptionID string) ItemLevelRecoveryConnectionsClient {
	return NewItemLevelRecoveryConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewItemLevelRecoveryConnectionsClientWithBaseURI creates an instance of the ItemLevelRecoveryConnectionsClient
// client using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI
// (sovereign clouds, Azure stack).
func NewItemLevelRecoveryConnectionsClientWithBaseURI(baseURI string, subscriptionID string) ItemLevelRecoveryConnectionsClient {
	return ItemLevelRecoveryConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Provision provisions a script which invokes an iSCSI connection to the backup data. Executing this script opens a
// file
// explorer displaying all the recoverable files and folders. This is an asynchronous operation. To know the status of
// provisioning, call GetProtectedItemOperationResult API.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// fabricName - fabric name associated with the backed up items.
// containerName - container name associated with the backed up items.
// protectedItemName - backed up item name whose files/folders are to be restored.
// recoveryPointID - recovery point ID which represents backed up data. iSCSI connection will be provisioned
// for this backed up data.
// parameters - resource ILR request
func (client ItemLevelRecoveryConnectionsClient) Provision(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters ILRRequestResource) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ItemLevelRecoveryConnectionsClient.Provision")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ProvisionPreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Provision", nil, "Failure preparing request")
		return
	}

	resp, err := client.ProvisionSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Provision", resp, "Failure sending request")
		return
	}

	result, err = client.ProvisionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Provision", resp, "Failure responding to request")
		return
	}

	return
}

// ProvisionPreparer prepares the Provision request.
func (client ItemLevelRecoveryConnectionsClient) ProvisionPreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, parameters ILRRequestResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"recoveryPointId":   autorest.Encode("path", recoveryPointID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/provisionInstantItemRecovery", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ProvisionSender sends the Provision request. The method will close the
// http.Response Body if it receives an error.
func (client ItemLevelRecoveryConnectionsClient) ProvisionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ProvisionResponder handles the response to the Provision request. The method always
// closes the http.Response Body.
func (client ItemLevelRecoveryConnectionsClient) ProvisionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Revoke revokes an iSCSI connection which can be used to download a script. Executing this script opens a file
// explorer
// displaying all recoverable files and folders. This is an asynchronous operation.
// Parameters:
// vaultName - the name of the recovery services vault.
// resourceGroupName - the name of the resource group where the recovery services vault is present.
// fabricName - fabric name associated with the backed up items.
// containerName - container name associated with the backed up items.
// protectedItemName - backed up item name whose files/folders are to be restored.
// recoveryPointID - recovery point ID which represents backed up data. iSCSI connection will be revoked for
// this backed up data.
func (client ItemLevelRecoveryConnectionsClient) Revoke(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ItemLevelRecoveryConnectionsClient.Revoke")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RevokePreparer(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Revoke", nil, "Failure preparing request")
		return
	}

	resp, err := client.RevokeSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Revoke", resp, "Failure sending request")
		return
	}

	result, err = client.RevokeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "backup.ItemLevelRecoveryConnectionsClient", "Revoke", resp, "Failure responding to request")
		return
	}

	return
}

// RevokePreparer prepares the Revoke request.
func (client ItemLevelRecoveryConnectionsClient) RevokePreparer(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerName":     autorest.Encode("path", containerName),
		"fabricName":        autorest.Encode("path", fabricName),
		"protectedItemName": autorest.Encode("path", protectedItemName),
		"recoveryPointId":   autorest.Encode("path", recoveryPointID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"vaultName":         autorest.Encode("path", vaultName),
	}

	const APIVersion = "2016-12-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/Subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}/revokeInstantItemRecovery", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RevokeSender sends the Revoke request. The method will close the
// http.Response Body if it receives an error.
func (client ItemLevelRecoveryConnectionsClient) RevokeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RevokeResponder handles the response to the Revoke request. The method always
// closes the http.Response Body.
func (client ItemLevelRecoveryConnectionsClient) RevokeResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}
