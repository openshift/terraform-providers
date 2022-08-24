package serverrestart

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FailoverMode string

const (
	FailoverModeForcedFailover    FailoverMode = "ForcedFailover"
	FailoverModeForcedSwitchover  FailoverMode = "ForcedSwitchover"
	FailoverModePlannedFailover   FailoverMode = "PlannedFailover"
	FailoverModePlannedSwitchover FailoverMode = "PlannedSwitchover"
)

func PossibleValuesForFailoverMode() []string {
	return []string{
		string(FailoverModeForcedFailover),
		string(FailoverModeForcedSwitchover),
		string(FailoverModePlannedFailover),
		string(FailoverModePlannedSwitchover),
	}
}

func parseFailoverMode(input string) (*FailoverMode, error) {
	vals := map[string]FailoverMode{
		"forcedfailover":    FailoverModeForcedFailover,
		"forcedswitchover":  FailoverModeForcedSwitchover,
		"plannedfailover":   FailoverModePlannedFailover,
		"plannedswitchover": FailoverModePlannedSwitchover,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FailoverMode(input)
	return &out, nil
}
