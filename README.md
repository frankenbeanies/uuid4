# UUID4

An RFC 4122 compliant UUID library

## License

[MIT](LICENSE)

## Usage

### New()

Generates a new UUID4

```go
uuid := uuid4.New()
```

### String()

Provides an RFC 4122 comliant string representation of the UUID4

```go
uuidStr := uuid4.New()
fmt.Println(uuidStr)
```