# Terraform Provider for Confidant [![Build Status](https://travis-ci.org/stripe/terraform-provider-confidant.svg?branch=master)](https://travis-ci.org/stripe/terraform-provider-confidant)

This terraform provider is used to manage resources supported by [Confidant](https://lyft.github.io/confidant/).

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) 0.10+
- Go 1.11.0 or higher
- [Confidant](https://github.com/lyft/confidant) 1.1.19 or higher

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/stripe/terraform-provider-confidant`

```sh
$ mkdir -p $GOPATH/src/github.com/stripe; cd $GOPATH/src/github.com/stripe
$ git clone git@github.com:stripe/terraform-provider-confidant
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/tripe/terraform-provider-confidant
$ go install
```

Or using bazel:

```sh
bazel build //:terraform-provider-confidant
```

## Using the provider

If you're building the provider, follow the instructions to [install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin)
After placing it into your plugins directory, run `terraform init` to initialize it.

### Configuring the provider

```sh
provider "confidant" {
  # The key name used for KMS auth (see Confidant's docs for more details)
  authkey   = "alias/for/your/kms/key"

  # The AWS region your KMS keys are located in
  region    = "us-west-2"

  # The URL to access confidant's API
  url       = "https://your_confidant_server.com"

  # The to/from to use for KMS authentication (see Confidant's docs for examples)
  from      = "your_aws_user_id"
  to        = "ConfidantServer"
}
```

### Resource: service

The provider allows you to create and manage confidant services. This includes
managing the secrets that are assigned to that service. It does not include
creating prerequisite's outside of the Confidant API such as IAM Roles.

```sh
resource "confidant_service" "example-service" {
  name        = "example-service"
  credentials = [
      "name_of_credential_1",
      "name_of_credential_2",
  ]
}
```

## Contributing

For guidance on contributing, please see [contribution guidelines](https://github.com/stripe/terraform-provider-confidant/blob/master/.github/CONTRIBUTING.md).
If you have other development questions we don't cover, please file an issue!


