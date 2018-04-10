package sparql

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/knakk/digest"
	"github.com/knakk/kbp/rdf"
	"github.com/knakk/kbp/rdf/memory"
)

// Repo represent a RDF repository, assumed to be
// queryable via the SPARQL protocol over HTTP.
type Repo struct {
	endpoint   string
	client     *http.Client
	retryOn404 bool
}

// NewRepo creates a new representation of a RDF repository. It takes a
// variadic list of functional options which can alter the configuration
// of the repository.
func NewRepo(addr string, options ...func(*Repo) error) (*Repo, error) {
	r := Repo{
		endpoint: addr,
		client:   http.DefaultClient,
	}
	return &r, r.SetOption(options...)
}

// SetOption takes one or more option function and applies them in order to Repo.
func (r *Repo) SetOption(options ...func(*Repo) error) error {
	for _, opt := range options {
		if err := opt(r); err != nil {
			return err
		}
	}
	return nil
}

// DigestAuth configures Repo to use digest authentication on HTTP requests.
func DigestAuth(username, password string) func(*Repo) error {
	return func(r *Repo) error {
		r.client.Transport = digest.NewTransport(username, password)
		return nil
	}
}

// Timeout instructs the underlying HTTP transport to timeout after given duration.
func Timeout(t time.Duration) func(*Repo) error {
	return func(r *Repo) error {
		r.client.Timeout = t
		return nil
	}
}

// RetryOn404 will retry request when response status is 404. (This happens when
// virtuoso is performing checkpoints.)
func RetryOn404() func(*Repo) error {
	return func(r *Repo) error {
		r.retryOn404 = true
		return nil
	}
}

func (r *Repo) qselect(body string, n int) (*http.Response, error) {
	req, err := http.NewRequest(
		"POST",
		r.endpoint,
		bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header.Set("Accept", "application/sparql-results+json")
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	const maxTries = 5
	if resp.StatusCode == http.StatusNotFound && r.retryOn404 && n <= maxTries {
		resp.Body.Close()
		time.Sleep(time.Duration(100*n) * time.Millisecond)
		return r.qselect(body, n+1)
	}

	return resp, err
}

// Select performs a SPARQL HTTP request to the Repo, and returns the
// parsed application/sparql-results+json response.
func (r *Repo) Select(q string) (*Results, error) {
	form := url.Values{}
	form.Set("query", q)
	b := form.Encode()

	resp, err := r.qselect(b, 1)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(body)
		var msg string
		if err != nil {
			msg = "Failed to read response body"
		} else {
			if strings.TrimSpace(string(b)) != "" {
				msg = "Response body: \n" + string(b)
			}
		}
		return nil, fmt.Errorf("Select: SPARQL request failed: %s. "+msg, resp.Status)
	}
	results, err := ParseJSON(body)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// Update performs a SPARQL HTTP request to the Repo
func (r *Repo) Update(q string) error {
	form := url.Values{}
	form.Set("query", q)
	b := form.Encode()

	resp, err := r.construct(b, 1)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := gzip.NewReader(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(body)
		var msg string
		if err != nil {
			msg = "Failed to read response body"
		} else {
			if strings.TrimSpace(string(b)) != "" {
				msg = "Response body: \n" + string(b)
			}
		}
		return fmt.Errorf("Update: SPARQL request failed: %s. "+msg, resp.Status)
	}
	return nil
}

func (r *Repo) construct(body string, n int) (*http.Response, error) {
	req, err := http.NewRequest(
		"POST",
		r.endpoint,
		bytes.NewBufferString(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(body)))
	req.Header.Set("Accept-Encoding", "gzip")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}

	const maxTries = 5
	if resp.StatusCode == http.StatusNotFound && r.retryOn404 && n <= maxTries {
		resp.Body.Close()
		time.Sleep(time.Duration(100*n) * time.Millisecond)
		return r.construct(body, n+1)
	}

	return resp, err
}

// Construct performs a SPARQL HTTP request to the Repo, and returns the
// result graph.
func (r *Repo) Construct(q string) (*memory.Graph, error) {
	form := url.Values{}
	form.Set("query", q)
	form.Set("format", "text/plain")
	b := form.Encode()

	resp, err := r.construct(b, 1)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(body)
		var msg string
		if err != nil {
			msg = "Failed to read response body"
		} else {
			if strings.TrimSpace(string(b)) != "" {
				msg = "Response body: \n" + string(b)
			}
		}
		return nil, fmt.Errorf("Construct: SPARQL request failed: %s. "+msg, resp.Status)
	}
	dec := rdf.NewDecoder(body)
	g := memory.NewGraph()
	for tr, err := dec.Decode(); err != io.EOF; tr, err = dec.Decode() {
		if err != nil {
			return g, err
		}
		g.Insert(tr)
	}
	return g, nil
}
