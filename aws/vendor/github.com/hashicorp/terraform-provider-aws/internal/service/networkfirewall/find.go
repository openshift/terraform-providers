package networkfirewall

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/networkfirewall"
)

// FindLoggingConfiguration returns the LoggingConfigurationOutput from a call to DescribeLoggingConfigurationWithContext
// given the context and FindFirewall ARN.
// Returns nil if the FindLoggingConfiguration is not found.
func FindLoggingConfiguration(ctx context.Context, conn *networkfirewall.NetworkFirewall, arn string) (*networkfirewall.DescribeLoggingConfigurationOutput, error) {
	input := &networkfirewall.DescribeLoggingConfigurationInput{
		FirewallArn: aws.String(arn),
	}
	output, err := conn.DescribeLoggingConfigurationWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// FindFirewallPolicyByNameAndARN returns the FirewallPolicyOutput from a call to DescribeFirewallPolicyWithContext
// given the context and at least one of FirewallPolicyArn and FirewallPolicyName.
func FindFirewallPolicyByNameAndARN(ctx context.Context, conn *networkfirewall.NetworkFirewall, arn string, name string) (*networkfirewall.DescribeFirewallPolicyOutput, error) {
	input := &networkfirewall.DescribeFirewallPolicyInput{}
	if arn != "" {
		input.FirewallPolicyArn = aws.String(arn)
	}
	if name != "" {
		input.FirewallPolicyName = aws.String(name)
	}

	output, err := conn.DescribeFirewallPolicyWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// FindResourcePolicy returns the Policy string from a call to DescribeResourcePolicyWithContext
// given the context and resource ARN.
// Returns nil if the FindResourcePolicy is not found.
func FindResourcePolicy(ctx context.Context, conn *networkfirewall.NetworkFirewall, arn string) (*string, error) {
	input := &networkfirewall.DescribeResourcePolicyInput{
		ResourceArn: aws.String(arn),
	}
	output, err := conn.DescribeResourcePolicyWithContext(ctx, input)
	if err != nil {
		return nil, err
	}
	if output == nil {
		return nil, nil
	}
	return output.Policy, nil
}
