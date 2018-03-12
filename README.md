# UUID4

An RFC 4122 compliant UUID library

## License

[MIT](LICENSE)

## Methods

### New()

Generates a new UUID4

```go
uuid := uuid4.New()
```

### String()

Provides an RFC 4122 comliant string representation of the UUID4

```go
uuidStr := uuid4.New().String()
fmt.Println(uuidStr)
```

### Bytes()

Provides the byte representation of UUID4

```go
uuidBytes := uuid4.New().Bytes()
fmt.Println(uuid)
```