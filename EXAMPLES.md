# Examples

Here are more examples of plugin usage.

Below you can find just few examples of chaining files transformations using Readers and Transformers. All the examples are taken from
[tests](tests/). If you wonder how those tests are created and how they actually work,
please check [How To Write Tests](TESTS.md).

Input `.yaml` file
```yaml
tests:
  name: multiple
  test_file:
    json2yaml_yaml2json: !futl yaml2json+json2yaml+dir://../../tests/dirtest/json
    base64_yamljson: !futl yaml2json+base64dec+base64enc+json2yaml://../../tests/filetest/inputfile_json.json
    yaml2json: !futl base64enc+yaml2json://../../tests/filetest/inputfile_yaml.yaml
    git_https_yaml2json: !futl base64dec+base64enc+yaml2json+git_https://github.com/true-north-engineering/helm-file-utils.git tests/filetest/inputfile_yaml.yaml
```

Expected output file
```yaml
tests:
    name: multiple
    test_file:
        base64_yamljson: |-
            {
              "json": [
                "simpletest",
                "to see if this works"
              ]
            }
        git_https_yaml2json: |-
            {
              "json": [
                "simpletest",
                "to see if this works"
              ]
            }
        json2yaml_yaml2json:
            json_input1.json: |-
                {
                  "json": [
                    "json test 1",
                    "to see if this works"
                  ]
                }
            json_input2.json: |-
                {
                  "json": [
                    "json test 2",
                    "this should also work"
                  ]
                }
        yaml2json: ewogICJqc29uIjogWwogICAgInNpbXBsZXRlc3QiLAogICAgInRvIHNlZSBpZiB0aGlzIHdvcmtzIgogIF0KfQ==
```