---
subcategory: "VPN (Site-to-Site)"
layout: "aws"
page_title: "AWS: aws_customer_gateway"
description: |-
  Get an existing AWS Customer Gateway.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_customer_gateway

Get an existing AWS Customer Gateway.

## Example Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsCustomerGateway } from "./.gen/providers/aws/data-aws-customer-gateway";
import { VpnConnection } from "./.gen/providers/aws/vpn-connection";
import { VpnGateway } from "./.gen/providers/aws/vpn-gateway";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const main = new VpnGateway(this, "main", {
      amazonSideAsn: Token.asString(7224),
      vpcId: Token.asString(awsVpcMain.id),
    });
    const foo = new DataAwsCustomerGateway(this, "foo", {
      filter: [
        {
          name: "tag:Name",
          values: ["foo-prod"],
        },
      ],
    });
    new VpnConnection(this, "transit", {
      customerGatewayId: Token.asString(foo.id),
      staticRoutesOnly: false,
      type: Token.asString(foo.type),
      vpnGatewayId: main.id,
    });
  }
}

```

## Argument Reference

This data source supports the following arguments:

* `id` - (Optional) ID of the gateway.
* `filter` - (Optional) One or more [name-value pairs][dcg-filters] to filter by.

[dcg-filters]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeCustomerGateways.html

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the customer gateway.
* `bgpAsn` - Gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
* `bgpAsnExtended` - Gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
* `certificateArn` - ARN for the customer gateway certificate.
* `deviceName` - Name for the customer gateway device.
* `ipAddress` - IP address of the gateway's Internet-routable external interface.
* `tags` - Map of key-value pairs assigned to the gateway.
* `type` - Type of customer gateway. The only type AWS supports at this time is "ipsec.1".

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

- `read` - (Default `20m`)

<!-- cache-key: cdktf-0.20.8 input-323e8a0aaa7a35d5108570e2a9e6c0bf67943d1d1bebc7ad0e2935acaed0f304 -->