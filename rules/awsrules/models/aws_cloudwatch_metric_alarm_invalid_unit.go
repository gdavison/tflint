// This file generated by `tools/model-rule-gen/main.go`. DO NOT EDIT

package models

import (
	"log"

	"github.com/hashicorp/hcl2/hcl"
	"github.com/wata727/tflint/issue"
	"github.com/wata727/tflint/tflint"
)

// AwsCloudwatchMetricAlarmInvalidUnitRule checks the pattern is valid
type AwsCloudwatchMetricAlarmInvalidUnitRule struct {
	resourceType  string
	attributeName string
	enum          []string
}

// NewAwsCloudwatchMetricAlarmInvalidUnitRule returns new rule with default attributes
func NewAwsCloudwatchMetricAlarmInvalidUnitRule() *AwsCloudwatchMetricAlarmInvalidUnitRule {
	return &AwsCloudwatchMetricAlarmInvalidUnitRule{
		resourceType:  "aws_cloudwatch_metric_alarm",
		attributeName: "unit",
		enum: []string{
			"Seconds",
			"Microseconds",
			"Milliseconds",
			"Bytes",
			"Kilobytes",
			"Megabytes",
			"Gigabytes",
			"Terabytes",
			"Bits",
			"Kilobits",
			"Megabits",
			"Gigabits",
			"Terabits",
			"Percent",
			"Count",
			"Bytes/Second",
			"Kilobytes/Second",
			"Megabytes/Second",
			"Gigabytes/Second",
			"Terabytes/Second",
			"Bits/Second",
			"Kilobits/Second",
			"Megabits/Second",
			"Gigabits/Second",
			"Terabits/Second",
			"Count/Second",
			"None",
		},
	}
}

// Name returns the rule name
func (r *AwsCloudwatchMetricAlarmInvalidUnitRule) Name() string {
	return "aws_cloudwatch_metric_alarm_invalid_unit"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsCloudwatchMetricAlarmInvalidUnitRule) Enabled() bool {
	return true
}

// Type returns the rule severity
func (r *AwsCloudwatchMetricAlarmInvalidUnitRule) Type() string {
	return issue.ERROR
}

// Link returns the rule reference link
func (r *AwsCloudwatchMetricAlarmInvalidUnitRule) Link() string {
	return ""
}

// Check checks the pattern is valid
func (r *AwsCloudwatchMetricAlarmInvalidUnitRule) Check(runner *tflint.Runner) error {
	log.Printf("[INFO] Check `%s` rule for `%s` runner", r.Name(), runner.TFConfigPath())

	return runner.WalkResourceAttributes(r.resourceType, r.attributeName, func(attribute *hcl.Attribute) error {
		var val string
		err := runner.EvaluateExpr(attribute.Expr, &val)

		return runner.EnsureNoError(err, func() error {
			found := false
			for _, item := range r.enum {
				if item == val {
					found = true
				}
			}
			if !found {
				runner.EmitIssue(
					r,
					`unit is not a valid value`,
					attribute.Expr.Range(),
				)
			}
			return nil
		})
	})
}