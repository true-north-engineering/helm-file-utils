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

- All exported names, non-trivial unexported declarations and complex code logic **SHOULD** have comments.
- Comments **SHOULD** be full sentences, even if it seems redundant.
- Comments **SHOULD** be brief and concise.

```go
// GetOrderByID fetch order for a given order ID.
func (service *orderService) GetOrderByID(ctx context.Context, orderID uint) (string, error) {
    ...
}
```

#### 2.1.2. Functions Comments

- Exported functions **MUST** have comments.
- Function comments **MUST** start with the name of the function and end in a period.

#### 2.1.3. Constant and Variable Comments

- Exported variables and constants **SHOULD** have comments.
- If there is a comment, it **MUST** begin with the name of what is being described and end in a period.
- If constants/variables are related and grouped, you **SHOULD** use a single comment for the group:

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

- Multiword names **MUST** use MixedCaps or mixedCaps rather than underscores.
- Declarations that are used only locally **MUST NOT** be exported.
- Words in names that are initialisms or acronyms (e.g. "URL" or "ID") **MUST** have a consistent case. For example, "URL" should appear as "URL" or "url", never as "Url", and "customerID" not as "customerId".

```go
func (h *CustomerHandler) GetCustomerID(w http.ResponseWriter, r *http.Request) { ... }
```

#### 2.2.2. Variable Naming

- Single-letter names **MUST NOT** be used on non-local variables.
- The further from its declaration, the more descriptive the name **SHOULD** be. For a method receiver, one or two letters is sufficient. Common variables such as loop indices and readers can be a single letter (i, r). More unusual and global variables should have more descriptive names.
- If variable type is bool, its name **MUST** start with has, is, can, etc. (or Has, Is, Can, etc. for exported variables)

```go
// Good
type Customer struct {
    IsDeleted bool
}
 
// Bad
type Customer struct {
    Deleted bool
}
```

- Boolean variables **MUST** always use positive names.

```go
// Good
type Customer struct {
    IsDeleted bool
}
 
// Bad
type Customer struct {
    IsNotDeleted bool
}
```

#### 2.2.3. Package Naming

- Package name **SHOULD** be lowercase only, not plural, short, concise, evocative single-word names. It doesn't need to be unique across all source code, in case of a collision the importing package can choose a different name to use locally.

```go
package customer // Good
package customers // Bad
package Customers // Bad
 
package mock // Good
package test_mocks // Bad
```

- Function declarations **SHOULD NOT** start with the package name, except when the function name is exact same as package name itself.

```go
log.Info() // Good
log.LogInfo() // Bad
```

- Imports **SHOULD NOT** be renamed, except for name collisions.

- In the event of name collision, you **SHOULD** rename the most local or project-specific import.

#### 2.2.4. Function Naming

- If the main purpose of functions or methods is returning a bool type value, the name of function or method **MUST** start with has, is, can, etc. (or Has, Is, Can, etc. for exported functions)
- An exported function **SHOULD** always accept interfaces and return structs (<https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8>)

#### 2.2.5. Getters and Setters

- Getter functions **SHOULD NOT** start with _Get_. For example, if a field is called owner (lower case, unexported), the getter method should be called Owner (upper case, exported), not GetOwner.
- Setter functions, if needed, can start with _Set_. For example SetOwner.

```go
type Customer struct {
    firstName   string
    lastName    string
    email       string
}
 
// Good
func (c *Customer) FirstName() string {
    return c.firstName
}
 
// Good
func (c *Customer) SetFirstName(name string) {
    firstName = name
}
 
// Bad
func (c *Customer) GetFirstName() string {
    return c.firstName
}
```

#### 2.2.6. Return Parameters Naming

- If the function returns two or three parameters of the same type, or if the meaning of a result isn't clear from context, you **SHOULD** use named return parameters to improve readability.

```go
// Good
func (f *Foo) Location() (lat, long float64, err error) {...}
 
// Bad
func (f *Foo) Location() (float64, float64, error) {...}
```

### 2.3. Declaring Empty Slices

- When declaring a slice, is **RECOMMENDED** to declare a nil slice instead of an empty slice.

```go
// Prefer this
var t []string
 
// over this
t := []string{}
```

- The former does not allocate memory and points to nil, while the latter allocates memory and points to a slice with 0 elements.
- Note that there are circumstances where a *zero-length* slice is preferred (e.g. encoding JSON objects, a *nil* slice encodes to *null*, while *[]string{}* encodes to the JSON *array []*).

### 2.4. Defer

-   _defer_ functions **SHOULD** be used whenever dealing with resources that needs to be released, regardless of which path a function takes to return. This guarantees that you will never forget to release the resource. e.g. IO operations.
-   You **MUST** make sure _defer_ function does not include a resource that might be _nil_ and results in an _invalid memory address_ panic.

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

- Error strings **MUST NOT** be capitalized (unless beginning with proper nouns or acronyms) or end with punctuation, since they are usually printed following other context. This does not apply to logging, which is implicitly line-oriented and not combined inside other messages.

```go
fmt.Errorf("something bad") // Good
fmt.Errorf("Something bad") // Bad
 
So that log.Printf("Reading %s: %v", filename, err) formats without a capital letter mid-message.
```

-   Errors **SHOULD NOT** be discarded using _ variables.

```go
// Bad
resp, _ := c.send(ctx, "GET", endpoint, nil)
 
// Good
resp, err := c.send(ctx, "GET", endpoint, nil)
if err != nil {
    // Handle error
}
```

-   Functions **SHOULD NOT** return values like _-1_, _0_ or _null_ as the only result to signal errors or missing results. A function should return an additional value to indicate whether its return values are valid. This return value may be an error, or a boolean when no explanation is needed. It should be the final return value.

```go
// Bad
resp, err := c.send(ctx, "GET", endpoint, nil)
if err != nil {
    return nil
}
return resp
 
// Good
resp, err := c.send(ctx, "GET", endpoint, nil)
if err != nil {
    return nil, err
}
return resp, nil
```

- The normal code path **SHOULD** be kept at a minimal indentation, by indenting the error handling and dealing with it first. This improves the readability of the code by permitting visually scanning the normal path quickly.

```go
// Bad
if err != nil {
    // error handling
} else {
    // normal code
}
 
// Bad
if err == nil {
    // normal code
} else {
    // error handling
}
 
// Good
if err != nil {
    // error handling
    return // or continue, etc.
}
// normal code
```
----------

## 3. References

-   <https://golang.org/doc/effective_go.html>
-   <https://github.com/golang/go/wiki/CodeReviewComments>
-   <https://blog.golang.org/package-names>
-   <https://medium.com/@cep21/what-accept-interfaces-return-structs-means-in-go-2fe879e25ee8>