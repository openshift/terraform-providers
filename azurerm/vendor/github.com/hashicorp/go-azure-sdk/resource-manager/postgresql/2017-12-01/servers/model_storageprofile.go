package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageProfile struct {
	BackupRetentionDays *int64              `json:"backupRetentionDays,omitempty"`
	GeoRedundantBackup  *GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
	StorageAutogrow     *StorageAutogrow    `json:"storageAutogrow,omitempty"`
	StorageMB           *int64              `json:"storageMB,omitempty"`
}
