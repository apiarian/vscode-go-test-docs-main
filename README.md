# vscode-go test docs main package

This is part of a test to see how documentation tools behave. This was created
for use in vscode-go development, but it could certainly be used in other go
documentation experiments.

This package imports three others:
1. `vscode-go-test-docs-gopathonly` only exists in the `$GOPATH`.
2. `vscode-go-test-docs-vendoroverride` exists in both the `$GOPATH` and the `vendor/` directory.
3. `vscode-go-test-docs-vendoronly` only exists in the `vendor/` directory. This repository does not exist in github.

## Setup

```bash
cd $GOPATH/github.com/apiarian
git clone https://github.com/apiarian/vscode-go-test-docs-main.git
git clone https://github.com/apiarian/vscode-go-test-docs-gopathonly.git
git clone https://github.com/apiarian/vscode-go-test-docs-vendoroverride.git
```
