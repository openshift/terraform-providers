package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerAppProperties struct {
	Configuration              *Configuration                 `json:"configuration,omitempty"`
	CustomDomainVerificationId *string                        `json:"customDomainVerificationId,omitempty"`
	LatestRevisionFqdn         *string                        `json:"latestRevisionFqdn,omitempty"`
	LatestRevisionName         *string                        `json:"latestRevisionName,omitempty"`
	ManagedEnvironmentId       *string                        `json:"managedEnvironmentId,omitempty"`
	OutboundIPAddresses        *[]string                      `json:"outboundIpAddresses,omitempty"`
	ProvisioningState          *ContainerAppProvisioningState `json:"provisioningState,omitempty"`
	Template                   *Template                      `json:"template,omitempty"`
}
