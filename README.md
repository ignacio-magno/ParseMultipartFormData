# Documentation

`go get github.com/ignacio-magno/parse_multipart_form_data`

Make multipart form from:

- body string with required headers boundary
- AWS events.ApiGatewayProxyRequest

### Example
```go
	request := NewBuilderMultipartFormDataFromString(map[string]string{
		"Content-Type": "multipart/form-data; boundary=--------------------------423813977796892037623442",
	}, string(b))

	f, err := request.Build()
	if err != nil {
		panic(err)
	}

	form, err := f.ReadForm(1 << 20)
	if err != nil {
		return
	}
```


code copied from http library