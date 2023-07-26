package devices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDetails struct {
	EstimatedInstallTimeInMins *int64                 `json:"estimatedInstallTimeInMins,omitempty"`
	FriendlyVersionNumber      *string                `json:"friendlyVersionNumber,omitempty"`
	InstallationImpact         *InstallationImpact    `json:"installationImpact,omitempty"`
	RebootBehavior             *InstallRebootBehavior `json:"rebootBehavior,omitempty"`
	Status                     *UpdateStatus          `json:"status,omitempty"`
	TargetVersion              *string                `json:"targetVersion,omitempty"`
	UpdateSize                 *float64               `json:"updateSize,omitempty"`
	UpdateTitle                *string                `json:"updateTitle,omitempty"`
	UpdateType                 *UpdateType            `json:"updateType,omitempty"`
}
