package aftership

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var url = "https://api.aftership.com/v4"

// Aftership allows for connections to occur to the aftership tracking service
type Aftership interface {
	NewTracking(NewTrackingRequest) (TrackingResponse, error)
	GetTracking(slug string, trackingNumber string, includedFields ...string) (TrackingResponse, error)
}

// New creates a new Aftership client
func New(apiKey string) (Aftership, error) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: 20,
		},
		Timeout: 5 * time.Second,
	}

	// ping aftership to check api key validity
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("could not establish the http client")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("aftership-api-key", apiKey)

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != 200 {
		return nil, fmt.Errorf("could not connect to afterhip (check api key)")
	}

	return &aftership{apiKey: apiKey, client: client}, nil

}

// ================================================= request =================================================

type request struct {
	req    *http.Request
	client *http.Client
}

// do executes the request
func (r *request) do() ([]byte, error) {
	response, err := r.client.Do(r.req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// ================================================ aftership ================================================

// aftership is the client struct used to make requests
type aftership struct {
	apiKey string
	client *http.Client
}

// ============= helpers =============

// newRequest returns a pre-configured request
func (c *aftership) newRequest(method string, body io.Reader, pathParams ...string) (*request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", url, strings.Join(pathParams, "/")), body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("aftership-api-key", c.apiKey)

	return &request{req: req, client: c.client}, err
}

// ============= public =============

// NewTracking creates a new tracking record in aftership
func (c *aftership) NewTracking(r NewTrackingRequest) (resp TrackingResponse, err error) {
	js, err := r.toJSON()
	if err != nil {
		return
	}

	request, err := c.newRequest("POST", js, "trackings")
	if err != nil {
		return
	}

	body, err := request.do()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	return
}

// GetTracking fetches new tracking information
func (c *aftership) GetTracking(slug string, trackingNumber string,
	included ...string) (resp TrackingResponse, err error) {

	var r []byte
	if len(included) != 0 {
		r, err = json.Marshal(included)
		if err != nil {
			return
		}
	} else {
		r = []byte{}
	}
	request, err := c.newRequest("GET", bytes.NewReader(r), "trackings", slug, trackingNumber)
	if err != nil {
		return
	}

	body, err := request.do()
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	return
}
