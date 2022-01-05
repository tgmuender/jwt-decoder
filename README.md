# Jwt Decoder

Utility to decode a Json Web Token. The encoded token is read from stdin, the decoded output will be written to stdout. The signature part is not verified.

## Building
```sh
$ go build -o jd src/JwtDecoder.go
```

## Sample Usage
```sh
$ echo "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" | ./jd
{
        "alg": "HS256",
        "typ": "JWT"
}
{
        "iat": 1516239022,
        "name": "John Doe",
        "sub": "1234567890"
}
```