# Go Coding Guidelines

## 1. Overview
This is a coding guideline for Go programming language with the purpose of helping to write more readable and maintainable code.

### Recommended reading
- [Effective GO](https://golang.org/doc/effective_go.html)
- [Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Package Names](https://blog.golang.org/package-names)

## 2. Guidelines
### 2.1. gofmt and goimports commands
- You should make use of _gofmt_ and _goimports_ commands. This will automatically fix the majority of mechanical style issues, goimports additionally add, remove and organize imports lines as necessary.
- Most IDEs have plugins to run this commands automatically on save.

### 2.1. Comments

#### 2.1.1. General

- All exported names, non-trivial unexported declarations and complex code logic should have comments.
- Comments should be full sentences, even if it seems redundant.
- Comments should be brief and concise.

```go
// GetOrderByID fetch order for a given order ID.
func (service *orderService) GetOrderByID(ctx context.Context, orderID uint) (string, error) {
    ...
}
```

#### 2.1.2. Functions Comments

- Exported functions must have comments.
- Function comments must start with the name of the function and end in a period.

#### 2.1.3. Constant and Variable Comments

- Exported variables and constants should have comments.
- If there is a comment, it must begin with the name of what is being described and end in a period.
- If constants/variables are related and grouped, you should use a single comment for the group:

```go
// Error codes returned by failures to parse an expression.
var (
  ErrInternal = errors.New("regexp: internal error")
  ErrUnmatchedLpar = errors.New("regexp: unmatched '('")
  ErrUnmatchedRpar = errors.New("regexp: unmatched ')'")
  ...
)
```

### 2.2. Naming

#### 2.2.1. General

- Multiword names must use MixedCaps or mixedCaps rather than underscores.
- Declarations that are used only locally must not be exported.
- Words in names that are initialisms or acronyms (e.g. "URL" or "ID") must have a consistent case. For example, "URL" should appear as "URL" or "url", never as "Url", and "customerID" not as "customerId".

```go
func (h *CustomerHandler) GetCustomerID(w http.ResponseWriter, r *http.Request) { ... }
```

#### 2.2.2. Variable Naming

- Single-letter names must not be used on non-local variables.
- The further from its declaration, the more descriptive the name should be. For a method receiver, one or two letters is sufficient. Common variables such as loop indices and readers can be a single letter (i, r). More unusual and global variables should have more descriptive names.
- If variable type is bool, its name must start with has, is, can, etc. (or Has, Is, Can, etc. for exported variables)

```go
// Good
type User struct {
    IsDeleted bool
}
 
// Bad
type User struct {
    Deleted bool
}
```

- Boolean variables must always use positive names.

```go
// Good
type User struct {
    IsDeleted bool
}
 
// Bad
type User struct {
    IsNotDeleted bool
}
```

#### 2.2.3. Package Naming

- Package name should be lowercase only, not plural, short, concise, evocative single-word names. It doesn't need to be unique across all source code, in case of a collision the importing package can choose a different name to use locally.

```go
package user // Good
package users // Bad
package Users // Bad
 
package mock // Good
package test_mocks // Bad
```

- Function declarations should not start with the package name, except when the function name is exact same as package name itself.

```go
log.Info() // Good
log.LogInfo() // Bad
```

- Imports should not be renamed, except for name collisions.

- In the event of name collision, you should rename the most local or project-specific import.

#### 2.2.4. Function Naming

- If the main purpose of functions or methods is returning a bool type value, the name of function or method must start with has, is, can, etc. (or Has, Is, Can, etc. for exported functions)
- An exported function should always accept interfaces and return structs (<https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8>)

#### 2.2.5. Getters and Setters

- Getter functions should not start with _Get_. For example, if a field is called owner (lower case, unexported), the getter method should be called Owner (upper case, exported), not GetOwner.
- Setter functions, if needed, can start with _Set_. For example SetOwner.

```go
type User struct {
    firstName   string
    lastName    string
    email       string
}
 
// Good
func (c *User) FirstName() string {
    return c.firstName
}
 
// Good
func (c *User) SetFirstName(name string) {
    firstName = name
}
 
// Bad
func (c *User) GetFirstName() string {
    return c.firstName
}
```

#### 2.2.6. Return Parameters Naming

- If the function returns two or three parameters of the same type, or if the meaning of a result isn't clear from context, you should use named return parameters to improve readability.

```go
// Good
func (f *Foo) Location() (lat, long float64, err error) {...}
 
// Bad
func (f *Foo) Location() (float64, float64, error) {...}
```

### 2.3. Declaring Empty Slices

- When declaring a slice, is recommended to declare a nil slice instead of an empty slice.

```go
// Prefer this
var t []string
 
// over this
t := []string{}
```

- The former does not allocate memory and points to nil, while the latter allocates memory and points to a slice with 0 elements.
- Note that there are circumstances where a *zero-length* slice is preferred (e.g. encoding JSON objects, a *nil* slice encodes to *null*, while *[]string{}* encodes to the JSON *array []*).

### 2.4. Defer

-   _defer_ functions should be used whenever dealing with resources that needs to be released, regardless of which path a function takes to return. This guarantees that you will never forget to release the resource. e.g. IO operations.
-   You must make sure _defer_ function does not include a resource that might be _nil_ and results in an _invalid memory address_ panic.

```go
// Bad
resp, err := c.send(ctx, "GET", endpoint, nil)
defer resp.Body.Close() // resp might be nil
if err != nil {
    return nil, err
}
 
// Bad
resp, err := c.send(ctx, "GET", endpoint, nil)
if err != nil {
    return nil, err
}
defer resp.Body.Close() // this can lead to memory leaks if c.send() returns both error and resp.Body
 
// Good
resp, err := c.send(ctx, "GET", endpoint, nil)
defer func() {
    if resp != nil && resp.Body != nil {
       resp.Body.Close()
    }
}()
if err != nil {
    return nil, err
}
```

### 2.5. Errors

- Error strings must not be capitalized (unless beginning with proper nouns or acronyms) or end with punctuation, since they are usually printed following other context. This does not apply to logging, which is implicitly line-oriented and not combined inside other messages.

```go
fmt.Errorf("not good") // Good
fmt.Errorf("Not good") // Bad
 
So that log.Printf("Reading %s: %v", filename, err) formats without a capital letter mid-message.
```

-   Errors should not be discarded using _ variables.

```go
// Bad
resp, _ := c.send(ctx, "GET", endpoint, nil)
 
// Good
resp, err := c.send(ctx, "GET", endpoint, nil)
if err != nil {
    // Handle error
}
```

### 2.6.Linting

Use the following linters at a minimum, because we feel that they
help to catch the most common issues and also establish a high bar for code
quality without being unnecessarily prescriptive:

- [errcheck] to ensure that errors are handled
- [goimports] to format code and manage imports
- [golint] to point out common style mistakes
- [govet] to analyze code for common mistakes
- [staticcheck] to do various static analysis checks

  [errcheck]: https://github.com/kisielk/errcheck
  [goimports]: https://godoc.org/golang.org/x/tools/cmd/goimports
  [golint]: https://github.com/golang/lint
  [govet]: https://golang.org/cmd/vet/
  [staticcheck]: https://staticcheck.io/

## 3. References

-   <https://golang.org/doc/effective_go.html>
-   <https://github.com/uber-go/guide/blob/master/style.md>
-   <https://github.com/golang/go/wiki/CodeReviewComments>
-   <https://blog.golang.org/package-names>
-   <https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8>
