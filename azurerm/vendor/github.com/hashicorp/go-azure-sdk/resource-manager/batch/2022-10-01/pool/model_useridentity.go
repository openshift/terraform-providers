package pool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserIdentity struct {
	AutoUser *AutoUserSpecification `json:"autoUser,omitempty"`
	UserName *string                `json:"userName,omitempty"`
}
