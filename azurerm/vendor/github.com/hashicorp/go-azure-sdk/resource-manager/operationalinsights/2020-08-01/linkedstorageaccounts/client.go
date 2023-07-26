package linkedstorageaccounts

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedStorageAccountsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewLinkedStorageAccountsClientWithBaseURI(endpoint string) LinkedStorageAccountsClient {
	return LinkedStorageAccountsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
