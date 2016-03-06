package contegixclassic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/ajg/form"
	"github.com/hashicorp/go-cleanhttp"
)

// Client provides a client to the Contegix Cloud Classic API
// https://docs.contegix.com/display/CLOUDCLASSIC/
type Client struct {
	// Access Token
	Token string

	// URL to the base API to use
	URL string

	// HttpClient is the client to use. A client with
	// default values will be used if not provided.
	HTTP *http.Client
}

// NewClient returns a new Contegix Cloud client,
// requires an authorization token.
func NewClient(token string) (*Client, error) {
	return NewCustomClient(token, "https://classic.contegix.com")
}

// NewCustomClient returns a new Contigix client with a custom base URL.
// It requires an authorization token.
func NewCustomClient(token string, baseUrl string) (*Client, error) {
	client := Client{
		Token: token,
		URL:   baseUrl,
		HTTP:  cleanhttp.DefaultClient(),
	}
	return &client, nil
}

// Creates a new request with the params
func (c *Client) NewRequest(method string, endpoint string, params interface{}) (*http.Request, error) {
	fullUrl := c.URL + "/api/v1" + url.QueryEscape(endpoint)

	body, err := form.EncodeToString(authParams{AuthToken: c.Token})
	if err != nil {
		return nil, fmt.Errorf("Error parsing authentication parameters: %v", err)
	}

	if params != nil {
		sParams, err := form.EncodeToString(params)
		if err != nil {
			return nil, fmt.Errorf("Error encoding request body: %v", err)
		}
		body = body + "&" + sParams
	}

	if method == "GET" {
		fullUrl += "?" + body
		body = ""
	}

	u, err := url.Parse(fullUrl)

	if err != nil {
		return nil, fmt.Errorf("Error parsing base URL: %v", err)
	}

	// Build the request
	req, err := http.NewRequest(method, u.String(), strings.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %v", err)
	}

	req.Header.Add("Accept", "application/json")

	// If it's a not a get, add a content-type
	if method != "GET" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	}

	return req, nil

}

func (c *Client) DoRequest(method string, endpoint string, params interface{}, result interface{}) (*http.Response, error) {
	req, err := c.NewRequest(method, endpoint, params)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	switch i := result.(type) {
	case **string:
		err = readString(resp.Body, i)
	default:
		err = decodeJson(resp.Body, &result)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func readString(reader io.Reader, result **string) error {
	var bout bytes.Buffer
	if reader != nil {
		_, err := io.Copy(&bout, reader)
		if err != nil {
			return err
		}
	}
	value := bout.String()
	*result = &value
	return nil
}

func decodeJson(reader io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&obj)
	if err != nil {
		return err
	}
	return nil
}

// parseErr is used to take an error json resp
// and return a single string for use in error messages
func parseErr(resp *http.Response) error {
	return fmt.Errorf("API Error: %s", resp.Body)
}

// decodeBody is used to JSON decode a body
func decodeBody(resp *http.Response, out interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &out); err != nil {
		return err
	}

	return nil
}

func encodeBody(obj interface{}) (io.Reader, error) {
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	if err := enc.Encode(obj); err != nil {
		return nil, err
	}
	return buf, nil
}

// checkResp wraps http.Client.Do() and verifies that the
// request was successful. A non-200 request returns an error
// formatted to included any validation problems or otherwise
func checkResp(resp *http.Response, err error) (*http.Response, error) {
	// If the err is already there, there was an error higher
	// up the chain, so just return that
	if err != nil {
		return resp, err
	}

	switch i := resp.StatusCode; {
	case i == 200:
		return resp, nil
	case i == 201:
		return resp, nil
	case i == 202:
		return resp, nil
	case i == 204:
		return resp, nil
	case i == 422:
		return nil, fmt.Errorf("API Error: %s", resp.Status)
	case i == 400:
		return nil, parseErr(resp)
	default:
		return nil, fmt.Errorf("API Error: %s", resp.Status)
	}
}

// Provides an auth_token property that can be encoded.
type authParams struct {
	AuthToken string `form:"auth_token"`
}
