[![CircleCI](https://circleci.com/gh/alexfalkowski/go-client-template.svg?style=svg)](https://circleci.com/gh/alexfalkowski/go-client-template)
[![codecov](https://codecov.io/gh/alexfalkowski/go-client-template/graph/badge.svg)](https://codecov.io/gh/alexfalkowski/go-client-template)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/go-client-template.svg)](https://pkg.go.dev/github.com/alexfalkowski/go-client-template)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# go-client-template

<!-- repository:description -->
A template for Go command-line clients with shared configuration and telemetry.
<!-- /repository:description -->

<!--
Client owners and documentation tools: expand the introduction with the intended
users, target service or API, and the task this client performs. Use repository
metadata for identity, internal/cmd/client/ for behavior, internal/config/ and
test/.config/client.yaml for configuration, and test/features/ for verified usage.
Add a working request, authentication requirements, and error, retry, timeout,
pagination, or rate limit behavior when supported by code or owner input. Explain
design decisions only when evidenced; do not invent integrations or rationale.
Remove this comment once the client docs are complete.
-->

<!-- template:setup -->
## Create a client

Create the repository through InfraOps using this template, with its final name
and description. Clone it into a directory matching the repository name, then run
from the clone's root:

```sh
gh auth login
tools/init
```

Initialization requires Bash 4+, GNU sed, Git, the GitHub CLI, and the development
tools below. It reads the name, owner, and description from the GitHub repository
at `origin`, then updates the module path, client names, README, and badges.
An empty description or a checkout directory that differs from the repository
name stops initialization before any files are changed.

The script installs dependencies, formats the project, and removes `tools/` and
this section. It then creates an initialization branch and runs `make ready`,
which commits, pushes, opens a pull request, and enables auto-merge.
<!-- /template:setup -->

## Development

Use the Go and Ruby toolchain provided by the
[CI image](.circleci/config.yml). The development workflow also uses betteralign,
golangci-lint, and RuboCop. See `make help` for supported commands and their
prerequisites.

Initialize the shared tooling before using Make, then install dependencies and
build the client:

```sh
git submodule sync && git submodule update --init
make dep
make build
```

## Client

Run the client with the local development configuration:

```sh
./go-client-template client -config file:test/.config/client.yaml
```

The starter logs `awesome client` and exits successfully. It does not yet call an
external API or configure authentication. Implement the client behavior in
[internal/cmd/client/](internal/cmd/client/) and document a working request here.
The [client feature](test/features/client/client.feature) verifies the initial
command's exit status and log output.

## Configuration

The local [configuration](test/.config/client.yaml) selects the development
environment, UUID identifiers, and text logging at info level:

```yaml
environment: development
id:
  kind: uuid
telemetry:
  logger:
    kind: text
    level: info
```

Add client-specific settings to [internal/config/](internal/config/) and keep the
sample configuration and usage documentation aligned with them.

## References

- [Go package documentation](https://pkg.go.dev/github.com/alexfalkowski/go-client-template)
- [Releases](https://github.com/alexfalkowski/go-client-template/releases)
- [Shared tooling](bin/README.md)
