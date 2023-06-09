// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package oapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"/5tb2jnVKQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
