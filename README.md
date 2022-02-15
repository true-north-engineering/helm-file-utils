# Helm File Utils plugin

A Helm downloader plugin that supports different file manipulations, conversions, encoders and decoders.

![Plugin Tests](https://github.com/true-north-engineering/helm-file-utils/actions/workflows/file-utils-release.yml/badge.svg)

## Installation

After installing Helm, simply run the following:
```bash
helm plugin install https://github.com/true-north-engineering/helm-file-utils
```

## Usage

Helm File Utils allows user to do multiple transformations over given file. 
This plugin is only applicable to the `-f` or `--values` option of a Helm
command (e.g. `install`, `upgrade` or `template`).  The basic usage
is to reference a directory (either absolutely, or relative to the
PWD) from which to collect all non-hidden files with the extension
`.yaml` or `.yml`, not including sub-directories. Keyword used to 
associate plugin with given file directory is **futl**.

Basic usage of plugin is as it follows:
````bash
helm install -f futl://path/to/yaml/file /path/to/chart/file
````

### File transformations

In given `.yaml` or `.yml` file, multiple file transformations are possible.
Transformations are classified in two categories - Transformers(**T**) and Readers(**R**). 
Every command needs to consist of **at most** one Reader and **at least** one Transformer separated with **+** sign.
Order of transformation evaluation is from right to left, which forces Reader to always execute first.

**Reader**\
Used to read the file from given destination. If none is provided, **file** is considered as default.\
Available Readers: **file, dir, ~~http(s)~~, ~~git~~**

**Transformer**\
Transformers are used to do various transformations over the file.\
Available Transformers: **base64enc, ~~base64dec, xslt, custom~~**

Template for chaining file transformations is:
````bash
!futl T+T+T+R://path/to/file
````


##Example usage

````bash
helm install -f futl://home/usr/files /home/usr/charts
````

```bash
* home
  * usr
    * files
      * values.yaml
    * charts
      * chart.yaml
```

```bash
#values.yaml

#default reader is file so having "file" listed as reader is deprecated
example_file: 
  - name: example_file
    file: !futl base64enc://example_file.txt
    
#since we are iterating over dir, reader of type "dir" is needed
example_dir: 
  - name: example_dir
    file: !futl base64enc+dir://example_dir
```