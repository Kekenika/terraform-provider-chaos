# Terraform Provider for Chaos Engineering

This is a [Terraform](https://www.terraform.io/) provider for [Chaos Engineering](https://en.wikipedia.org/wiki/Chaos_engineering).

## Maintainers

This provider plugin is maintained by:

* [@KeisukeYamashita](https;//github.com/KeisukeYamashita)

## Support

- Simple failures
    - Terraform Apply
    - Terraform Plan

## Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= v0.12.0 (v0.11.x may work but not supported actively)

## Building The Provider

Clone repository to: `$GOPATH/src/github.com/Kekenika/terraform-provider-chaos`

```console
$ mkdir -p $GOPATH/src/github.com/Kekenika; cd $GOPATH/src/github.com/Kekenika
$ git clone git@github.com:Kekenika/terraform-provider-chaos
Enter the provider directory and build the provider

$ cd $GOPATH/src/github.com/Kekenika/terraform-provider-chaos
$ make build
```
