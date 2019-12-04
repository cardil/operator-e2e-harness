package config

import "time"

var OperatorsNamespace = "openshift-operators"
var Channel = "stable"

type PollingConfig struct {
	// Interval specifies the time between two polls.
	Interval time.Duration
	// Timeout specifies the timeout for the function PollImmediate to reach a certain status.
	Timeout  time.Duration
}

var Polling = PollingConfig{
	Interval: 10 * time.Second,
	Timeout:  5 * time.Minute,
}
