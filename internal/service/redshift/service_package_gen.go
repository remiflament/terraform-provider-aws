// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package redshift

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  DataSourceCluster,
			TypeName: "aws_redshift_cluster",
		},
		{
			Factory:  DataSourceClusterCredentials,
			TypeName: "aws_redshift_cluster_credentials",
		},
		{
			Factory:  DataSourceOrderableCluster,
			TypeName: "aws_redshift_orderable_cluster",
		},
		{
			Factory:  DataSourceServiceAccount,
			TypeName: "aws_redshift_service_account",
		},
		{
			Factory:  DataSourceSubnetGroup,
			TypeName: "aws_redshift_subnet_group",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceAuthenticationProfile,
			TypeName: "aws_redshift_authentication_profile",
		},
		{
			Factory:  ResourceCluster,
			TypeName: "aws_redshift_cluster",
		},
		{
			Factory:  ResourceClusterIAMRoles,
			TypeName: "aws_redshift_cluster_iam_roles",
		},
		{
			Factory:  ResourceClusterSnapshot,
			TypeName: "aws_redshift_cluster_snapshot",
		},
		{
			Factory:  ResourceEndpointAccess,
			TypeName: "aws_redshift_endpoint_access",
		},
		{
			Factory:  ResourceEndpointAuthorization,
			TypeName: "aws_redshift_endpoint_authorization",
		},
		{
			Factory:  ResourceEventSubscription,
			TypeName: "aws_redshift_event_subscription",
		},
		{
			Factory:  ResourceHSMClientCertificate,
			TypeName: "aws_redshift_hsm_client_certificate",
		},
		{
			Factory:  ResourceHSMConfiguration,
			TypeName: "aws_redshift_hsm_configuration",
		},
		{
			Factory:  ResourceParameterGroup,
			TypeName: "aws_redshift_parameter_group",
		},
		{
			Factory:  ResourcePartner,
			TypeName: "aws_redshift_partner",
		},
		{
			Factory:  ResourceScheduledAction,
			TypeName: "aws_redshift_scheduled_action",
		},
		{
			Factory:  ResourceSecurityGroup,
			TypeName: "aws_redshift_security_group",
		},
		{
			Factory:  ResourceSnapshotCopyGrant,
			TypeName: "aws_redshift_snapshot_copy_grant",
		},
		{
			Factory:  ResourceSnapshotSchedule,
			TypeName: "aws_redshift_snapshot_schedule",
		},
		{
			Factory:  ResourceSnapshotScheduleAssociation,
			TypeName: "aws_redshift_snapshot_schedule_association",
		},
		{
			Factory:  ResourceSubnetGroup,
			TypeName: "aws_redshift_subnet_group",
		},
		{
			Factory:  ResourceUsageLimit,
			TypeName: "aws_redshift_usage_limit",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Redshift
}

var ServicePackage = &servicePackage{}
