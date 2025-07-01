# Singularity
[![codecov](https://codecov.io/github/data-preservation-programs/singularity/branch/main/graph/badge.svg?token=1A3BMQU3LM)](https://codecov.io/github/data-preservation-programs/singularity)
[![Go Report Card](https://goreportcard.com/badge/github.com/data-preservation-programs/singularity)](https://goreportcard.com/report/github.com/data-preservation-programs/singularity)
[![Go Reference](https://pkg.go.dev/badge/github.com/data-preservation-programs/singularity.svg)](https://pkg.go.dev/github.com/data-preservation-programs/singularity)
[![Build](https://github.com/data-preservation-programs/singularity/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/data-preservation-programs/singularity/actions/workflows/go.yml)

The new pure-go implementation of Singularity provides everything you need to onboard your, or your client's data to Filecoin network.

## Documentation
[Read the Doc](https://data-programs.gitbook.io/singularity/overview/readme)

## Related projects
- [js-singularity](https://github.com/tech-greedy/singularity) -
The predecessor that was implemented in Node.js
- [js-singularity-import-boost](https://github.com/tech-greedy/singularity-import) -
Automatically import deals to boost for Filecoin storage providers
- [js-singularity-browser](https://github.com/tech-greedy/singularity-browser) -
A next.js app for browsing singularity made deals
- [go-generate-car](https://github.com/tech-greedy/generate-car) -
The internal tool used by `js-singularity` to generate car files as well as commp
- [go-generate-ipld-car](https://github.com/tech-greedy/generate-car#generate-ipld-car) -
The internal tool used by `js-singularity` to regenerate the CAR that captures the unixfs dag of the dataset.

## License
Dual-licensed under [MIT](https://github.com/filecoin-project/lotus/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/lotus/blob/master/LICENSE-APACHE)
