# AWS Vault Exec

Wrapper to help run aws-vault exec managed via ENV

Expects AWS_VAULT_PROFILE to be set

If its not set, any following args will be executed on their own


## Releasing

Needs `goreleaser` (https://goreleaser.com/)

`git tag -a VERSION -m VERSION`

eg: `git tag -a 1.0.0 -m 1.0.0`

`goreleaser release --rm-dist`
