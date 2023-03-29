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
