# Contribution guide

Helm file utils project is [Apache 2.0 licensed](LICENSE) and accepts
contributions via GitHub pull requests. This document outlines some of the
conventions on development workflow, branch formatting, commit message formatting 
and other resources to make it easier to get your contribution accepted.

## Support Channels

The official support channels, for both users and contributors, are:

- ??? for user questions.
- GitHub [Issues](https://github.com/true-north-engineering/helm-file-utils/issues)* for bug reports and feature requests.

*Before opening a new issue or submitting a new pull request, it's helpful to
search the project - it's likely that another user has already reported the
issue you're facing, or it's a known issue that we're already planning to fix in upcoming release.

## How to Contribute

Pull Requests (PRs) are the main and exclusive way to contribute to the official Helm File Utils project.
In order for a PR to be accepted it needs to pass a list of requirements:

- You should be able to run the same query using `git`. We don't accept features that are not implemented in the official git implementation
- The expected behavior must match the [official git implementation](https://github.com/git/git)
- The actual behavior must be correctly explained with natural language and providing a minimum working example in Go that reproduces it
- All PRs must be written in idiomatic Go, formatted according to [gofmt](https://golang.org/cmd/gofmt/), and without any warnings from [go vet](https://golang.org/cmd/vet/)
- PRs must follow main Go principles (e.g. package names, code commenting, ...), more can be found in the [Effective Go](https://go.dev/doc/effective_go) or [CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments)
- Before requesting PR, all tests shall pass
- If the PR is either a bug fix or a new feature, it has to include a suite of unit tests, for more please refer to [How To Write Tests](TESTS.md)
- In any case, all the PRs have to pass the personal evaluation of at least one of the maintainers of Helm File Utils plugin

### Branches and format of commit message

Every branch should follow branch naming conventions. If branch refers to a specific issue or bug, please refer to it by starting 
the name of the branch with issue ID followed by issue name. Otherwise, try to name it as describable as possible.

e.g. for issue -  #100 Add yaml2json transformer :
```text
100-yaml2json-transformer
```

Every commit message should describe what was changed, under which context and, if applicable, the GitHub issue it relates to.
Please also try to keep commits as 