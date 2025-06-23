// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tailsamplingprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"

import (
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"
	"go.opentelemetry.io/collector/component"
)

func getNewDropPolicy(settings component.TelemetrySettings, config *DropCfg) (sampling.PolicyEvaluator, error) {
	subPolicyEvaluators := make([]sampling.PolicyEvaluator, len(config.SubPolicyCfg))
	for i := range config.SubPolicyCfg {
		policyCfg := &config.SubPolicyCfg[i]
		policy, err := getDropSubPolicyEvaluator(settings, policyCfg)
		if err != nil {
			return nil, err
		}
		subPolicyEvaluators[i] = policy
	}
	return sampling.NewDrop(settings.Logger, subPolicyEvaluators), nil
}

// Return instance of and sub-policy
func getDropSubPolicyEvaluator(settings component.TelemetrySettings, cfg *AndSubPolicyCfg) (sampling.PolicyEvaluator, error) {
	return getSharedPolicyEvaluator(settings, &cfg.sharedPolicyCfg)
}
