package linkedservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedServiceProperties struct {
	ProvisioningState     *LinkedServiceEntityStatus `json:"provisioningState,omitempty"`
	ResourceId            *string                    `json:"resourceId,omitempty"`
	WriteAccessResourceId *string                    `json:"writeAccessResourceId,omitempty"`
}
