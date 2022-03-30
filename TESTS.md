# How to write tests?

If you are either adding new feature to the plugin or just adding more tests, please follow the instructions below.

## Adding new feature

When adding a new Reader or Transformer, you should also cover tests for it. Let's say you've developed 
Transformer called **XSLT**. Here are the basic principles to follow:

- As it is Transformer, it should be placed in [transformer](file-utils/transformer) and named [xslt.go](file-utils/transformer/xslt.go).
- Test file for xslt Transformer should be placed in [tests](file-utils/tests) and named [xslt_test.go](file-utils/tests/xslt_test.go)*

***NOTE** - every file must follow the ``[reader|trasnformer]_test`` naming convention otherwise Go won't recognize file as test

## Adding test examples