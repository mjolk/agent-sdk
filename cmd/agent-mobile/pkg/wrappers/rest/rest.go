/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/hyperledger/aries-framework-go/pkg/common/log"
)

var logger = log.New("aries-agent-mobile/wrappers/rest")

type httpClient interface {
	Do(r *http.Request) (*http.Response, error)
}

func exec(u, token string, httpClient httpClient, endpoint endpoint, request []byte) ([]byte, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to parse url [%s]: %v", u, err)
	}

	parsedURL.Path = path.Join(parsedURL.Path, endpoint.Path)

	parsedURL.Path, err = embedParams(parsedURL.Path, request)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to embed params in request path: %v", err)
	}

	resp, err := makeHTTPRequest(httpClient, endpoint.Method,
		parsedURL.String(), token, request)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to make http request to [%s]: %v", parsedURL.String(), err)
	}

	return resp, nil
}

func makeHTTPRequest(httpClient httpClient, method, agentURL, token string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(context.Background(), method, agentURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create a new http request for [%s]: %w", agentURL, err)
	}

	req.Header.Set("Content-Type", "application/json")

	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	response, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error while making request to [%s]: %w", agentURL, err)
	}

	defer func() {
		e := response.Body.Close()
		if e != nil {
			logger.Warnf("failed to close response body: %w", e)
		}
	}()

	return ioutil.ReadAll(response.Body)
}

func embedParams(reqPath string, body []byte) (newURL string, err error) {
	params := []string{"piid", "id", "name"}
	newURL = reqPath

	for _, param := range params {
		newURL, err = embedParam(newURL, param, body)
		if err != nil {
			return
		}
	}

	return newURL, nil
}

func embedParam(reqPath, paramKey string, body []byte) (string, error) {
	toReplace := fmt.Sprintf("{%s}", paramKey)

	if body == nil || !strings.Contains(reqPath, toReplace) {
		return reqPath, nil
	}

	model := make(map[string]interface{})
	if err := json.Unmarshal(body, &model); err != nil {
		return reqPath, fmt.Errorf("failed to unmarshal request body: %w", err)
	}

	val, ok := model[paramKey]
	if !ok {
		return reqPath, fmt.Errorf("no %s found in request path", paramKey)
	}

	newURL := strings.ReplaceAll(reqPath, toReplace, fmt.Sprintf("%s", val))

	return newURL, nil
}
