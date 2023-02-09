package parse_multipart_form_data

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"io"
	"mime"
	"mime/multipart"
)

type BuilderMultipartFormData struct {
	Headers map[string]string
	Body    io.Reader
}

func NewBuilderMultipartFormDataFromEventApiGatewayRequest(e events.APIGatewayProxyRequest) *BuilderMultipartFormData {
	if e.IsBase64Encoded {
		un64, err := base64.StdEncoding.DecodeString(e.Body)
		if err != nil {
			panic(err)
		}

		e.Body = string(un64)
	}

	return &BuilderMultipartFormData{
		Headers: e.Headers,
		Body:    bytes.NewReader([]byte(e.Body)),
	}
}

func NewBuilderMultipartFormDataFromString(headers map[string]string, body string) *BuilderMultipartFormData {
	return &BuilderMultipartFormData{
		Headers: headers,
		Body:    bytes.NewReader([]byte(body)),
	}
}

func (r *BuilderMultipartFormData) Build() (*multipart.Reader, error) {
	return r.mireader()
}

func (r *BuilderMultipartFormData) mireader() (*multipart.Reader, error) {
	v := r.Headers["Content-Type"]

	if v == "" {
		return nil, fmt.Errorf("missing Content-Type header")
	}

	if r.Body == nil {
		return nil, errors.New("missing form body")
	}
	d, params, err := mime.ParseMediaType(v)
	if err != nil || !(d == "multipart/form-data") {
		return nil, fmt.Errorf("bad Content-Type header: %v", err)
	}
	boundary, ok := params["boundary"]
	if !ok {
		return nil, fmt.Errorf("missing boundary param in Content-Type header")
	}
	return multipart.NewReader(r.Body, boundary), nil
}
