# urlf

Simple and secure URL formatting.

## Install
`go get github.com/cmar0027/urlf@v1.0.0`

Godocs [here](https://pkg.go.dev/github.com/cmar0027/urlf@v1.0.0).

## Usage

### urlf.Sprintf()
Sprintf securely formats a url.

Works like fmt.Sprintf but only accepts `%p`, `%q` and `%%`:
 - `%p` calls **url.PathEscape** on the corresponding argument
 - `%q` calls **url.QueryEscape** on the corresponding argument
 - `%%` escapes an '%'

Example:
```go
urlf.Sprintf("/users/%p/orders/%p?after=%q&sort=%q", userId, orderId, lastId, sortField)
```
