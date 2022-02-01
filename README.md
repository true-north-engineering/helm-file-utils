# Helm File Utils plugin

A Helm downloader plugin that supports different file manipulations, conversions, encoders and decoders.

![Plugin Tests](https://github.com/true-north-engineering/helm-file-utils/actions/file-utils-release.yml/badge.svg)

## Installation

After installing Helm, simply run the following:
```bash
helm plugin install https://github.com/true-north-engineering/helm-file-utils
```

## Usage

This is only applicable to the `-f` or `--values` option of a Helm
command (e.g. `install`, `upgrade` or `template`).  The basic usage
is to reference a directory (either absolutely, or relative to the
PWD) from which to collect all non-hidden files with the extension
`.yaml` or `.yml`, not including sub-directories:

```bash
helm upgrade -f base64enc://path/to/values path/to/chart
```
