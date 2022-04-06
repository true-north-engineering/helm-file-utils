# How to write tests?

If you are either adding new feature to the plugin or just adding more tests, please follow the instructions below.

## Adding new feature

When adding a new Reader or Transformer, you should also cover tests for it. Let's say you've developed 
Transformer called **JSON2YAML**. Here are the basic principles to follow:

- As it is Transformer, it should be placed in [transformer](file-utils/transformer) and named [json2yaml.go](file-utils/transformer/json2yaml.go).
- Test file for json2yaml Transformer should be placed in [tests](file-utils/tests) and named [json2yaml_test.go](file-utils/tests/jsom2yaml_test.go)*

### Test file

Test file should look like this:
```go
func TestJson2YamlTransform(t *testing.T) {
    var testCase = NewCmdTestCase("json2yaml", "json2yaml test",
    "futl://../../tests/yaml_json/json2yaml/input/values_json2yaml.yaml",
    "../../tests/yaml_json/json2yaml/output/json2yaml_output.txt", "", false, )
    ExecuteTests(t, *testCase)
}
```
Where constructor for creating a test is:
```go
NewCmdTestCase(name string, description string, input string, golden string, protocol string, wantError bool)
```

### How does it work

File transformations are done parsing the `!futl` command found in input file. Parsing is done from right to left which forces
Reader to always execute first. After all the transformations are done, file is compared to the golden file. Test is 

***NOTE** - every file **must** follow the ``[reader|trasnformer]_test`` naming convention otherwise Go won't recognize file as test
and program may not work as expected


## Adding test examples

As mentioned above, every test file refers to input and output(golden). 

Input file should be a `.yaml` file where the real magic happens. When creating an input file you can follow the template given below,
although it's not mandatory to do so.

```yaml
tests:
  name: json2yaml
  test:
    file: !futl json2yaml://../../tests/filetest/inputfile_json.json
```

where [inputfile_json.json](tests/filetest/inputfile_json.json) is

```json
{
  "json": [
    "simpletest",
    "to see if this works"
  ]
}
```

Output file should be a `.txt` file containing expected output once transformation is done.

```yaml
tests:
    name: json2yaml
    test:
        file: |
            json:
                - simpletest
                - to see if this works
```