package MultipartFormData

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

func TestFromString(t *testing.T) {
	b, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}

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

	for i := range form.File["file"] {
		open, err := form.File["file"][i].Open()
		if err != nil {
			panic(err)
		}

		body, err := io.ReadAll(open)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(fmt.Sprintf("test%v.pdf", i), body, 0644)
	}

}

func ServerTest() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		r.ParseMultipartForm(1 << 20)

		err = os.WriteFile("test.txt", body, 0644)
		if err != nil {
			return
		}
	}))

	if err != nil {
		return
	}
}
