# Development

## References

+ [API](https://www.goatcounter.com/help/api)
+ [Swagger](https://www.goatcounter.com/api.json)


## [gnostic Go Generator plugin](https://github.com/google/gnostic-go-generator) (Archived)

Didn't work. Gave up!

```console
gnostic api.json --go-generator-out=goatcounter
Errors reading api.json
Plugin error: [client.go:87:41: expected '(', found '{' (and 10 more errors) server.go:95:30: expected '(', found '{' (and 10 more errors) provider.go:25:19: expected ';', found '{' (and 1 more errors) types.go:223:24: expected type, found '{' (and 7 more errors)]
```

```bash
# Even though github.com/google/gnostic-go-generator
# Module declares its path as: github.com/googleapis/gnostic-go-generator
go install github.com/googleapis/gnostic-go-generator@latest

which gnostic-go-generator
/home/dazwilkin/go/bin/gnostic-go-generator
```


