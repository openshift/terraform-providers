// Package filesystem implements the Azure ARM Filesystem service API version 2016-11-01.
//
// Creates an Azure Data Lake Store filesystem client.
package filesystem

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultAdlsFileSystemDNSSuffix is the default value for adls file system dns suffix
	DefaultAdlsFileSystemDNSSuffix = "azuredatalakestore.net"
)

// BaseClient is the base client for Filesystem.
type BaseClient struct {
	autorest.Client
	AdlsFileSystemDNSSuffix string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithoutDefaults(DefaultAdlsFileSystemDNSSuffix)
}

// NewWithoutDefaults creates an instance of the BaseClient client.
func NewWithoutDefaults(adlsFileSystemDNSSuffix string) BaseClient {
	return BaseClient{
		Client:                  autorest.NewClientWithUserAgent(UserAgent()),
		AdlsFileSystemDNSSuffix: adlsFileSystemDNSSuffix,
	}
}
