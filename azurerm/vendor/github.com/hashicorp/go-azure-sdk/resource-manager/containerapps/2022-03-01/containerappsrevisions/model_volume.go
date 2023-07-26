package containerappsrevisions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Volume struct {
	Name        *string      `json:"name,omitempty"`
	StorageName *string      `json:"storageName,omitempty"`
	StorageType *StorageType `json:"storageType,omitempty"`
}
