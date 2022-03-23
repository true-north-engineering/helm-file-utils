# Helm File Utils plugin

A Helm downloader plugin that supports different file manipulations, conversions, encoders and decoders.

![Plugin Tests](https://github.com/true-north-engineering/helm-file-utils/actions/workflows/file-utils-release.yml/badge.svg)


## Table of contents

* [Installation](#install)
* [Usage and examples](#usage-and-examples)
    * [File transformations](#file-transformations)
    * [Examples](#examples)
  
## Installation

After installing Helm, simply run the following:

<<<<<<< Updated upstream
=======
```bash
helm plugin install https://github.com/true-north-engineering/helm-file-utils.git
```

For installing a specific release version (e.g. v0.1.0) please use following syntax:

>>>>>>> Stashed changes
```bash
helm plugin install https://github.com/true-north-engineering/helm-file-utils.git --version 0.1.0
```

```bash
https://github.com/true-north-engineering/helm-file-utils/releases/download/v0.1.3/helm-file-utils_0.1.3_linux_amd64.tar.gz
```
## Usage and examples

Helm File Utils allows user to do multiple transformations over given file. This plugin is only applicable to the `-f`
or `--values` option of a Helm command (e.g. `install`, `upgrade` or `template`). The basic usage is to reference a
directory (either absolutely, or relative to the PWD) from which to collect all non-hidden files with the extension
`.yaml` or `.yml`, not including sub-directories. Keyword used to associate plugin with given file directory is **futl**
.

Basic usage of plugin is as it follows:

````bash
helm install [NAME] [CHART] [flags] -f futl://path/to/values.yaml
````

## File transformations

In given `.yaml` or `.yml` file, multiple file transformations are possible. Transformations are classified in two
categories - Transformers(**T**) and Readers(**R**). Every command needs to consist of **at most** one Reader and
optional Transformers separated with **+** sign. Order of transformation evaluation is from right to left, which forces
Reader to always execute first.

### Reader
Used to read the file from given destination. If none is provided, **file** is considered as default.\
<<<<<<< Updated upstream
Available Readers: **file, dir
**Transformer**\
Transformers are used to do various transformations over the file.\
Available Transformers: **base64enc, base64dec**
=======
Available Readers: **file, dir, https, git_https, ssh**

**File**

**Dir**

**Https**

**Git_https**

**Ssh**

### Transformer
Transformers are used to do various transformations over the file.\
Available Transformers: **base64enc, base64dec, yaml2json, json2yaml**

**base64enc**

**base64dec**

**yaml2json**

**json2yaml**
>>>>>>> Stashed changes

Template for chaining file transformations is:

````bash
!futl T+T+T+R://path/to/file
````

<<<<<<< Updated upstream
## Example usage
=======

### Examples
>>>>>>> Stashed changes

````bash
helm install [NAME] [CHART] [flags] -f futl://home/usr/files 
````

```bash
* home
  * usr
    * files
      * values.yaml
    * charts
      * chart.yaml
```

```yaml
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

```yaml
#Chart.yaml
#this is just an example of a simple Chart file that is provided
apiVersion: v1
appVersion: "1.0"
description: Deploy a basic Chart Config Map
home: https://helm.sh/helm
name: example_chart
sources:
- https://github.com/helm/helm
version: 0.1.0
```
<<<<<<< Updated upstream
=======

For more examples please visit [this](EXAMPLES.md) page or check [tests](tests/).


## Issues

## 

## Contribution guide

>>>>>>> Stashed changes
