# VerifiedSCION

> **Warning**:
> This repository is currently outdated.
> We are currently in the process of updating our specification to more recent versions of SCION. More updates on this will be provided soon.

This package contains the **verified** implementation of the
[SCION](http://www.scion-architecture.net) protocol, a future Internet architecture.
SCION is the first
clean-slate Internet architecture designed to provide route control, failure
isolation, and explicit trust information for end-to-end communication.

![VerifiedSCION sticker](./logo.png)

To find out more about the project, please visit the [official project page](https://www.pm.inf.ethz.ch/research/verifiedscion.html).

## Methodology
We focus on verifying the main implementation of SCION, written in the *Go* programming language.

To that end, we have developed [Gobra](https://www.pm.inf.ethz.ch/research/gobra.html), a program verifier for Go. Gobra allows users to annotate Go code with specifications in the form of logical assertions establishing the behaviour of the program. 
It then automatically checks whether the implementation matches its given specification.

Initially, we aim at verifying the dataplane component of the SCION border router. We established the following milestones that we 
will achieve in the process:
1. verify memory safety, crash freedom, and race-freedom of the SCION dataplane code
2. prove progress properties of the dataplane code 
3. prove that the IO behaviour of the router matches the protocol description

When necessary, we make reasonable assumptions and explicitly state them.

## Repo Structure
The repository contains two main directories:
- `go/` contains the original code on which the verified version is based (commit `ae63a60fe8ade106230f20a6e6eb086529f7a2e0` from the [SCION repository](https://github.com/scionproto/scion))
- `gobra/` contains the verified version of the code using *Gobra*

**Observations**
- The package structure of `gobra/` directly mimics the one of `go/`
- The code available in `gobra/` does not contain a complete verified version of the one available in `go/`. Instead, it contains only the code required to verify the dataplane of the border router. This is to be expected, given that this is an ongoing project.

## License
[![License](https://img.shields.io/github/license/scionproto/scion.svg?maxAge=2592000)](https://github.com/scionproto/scion/blob/master/LICENSE)
