// Package apiclient provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package apiclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/oapi-codegen/runtime"
)

// Cluster Cluster represent a secure cloud stack Kubernetes cluster
type Cluster struct {
	// Context Context is defining the JSON-LD context for dereferencing the data as stated in the JSON-LD specification https://www.w3.org/TR/json-ld/#keywords
	Context *map[string]interface{} `json:"@context,omitempty"`

	// Id ID is identifying the node with an IRI as stated in the JSON-LD specification https://www.w3.org/TR/json-ld/#keywords
	Id *string `json:"@id,omitempty"`

	// Included Included will container linked resources included here for convenience
	Included *[]map[string]interface{} `json:"@included,omitempty"`

	// Type Type is optional explicit definition of the type of node as stated in the JSON-LD specification https://www.w3.org/TR/json-ld/#keywords
	Type *string `json:"@type,omitempty"`

	// ApiEndpoint ApiEndpoint is the uri for the Kubernetes api, this may be empty
	ApiEndpoint *string `json:"apiEndpoint,omitempty"`

	// Capacity Capacity is the capacity per node type, i.e., control-plane and worker nodes
	Capacity      *map[string]Capacity `json:"capacity,omitempty"`
	ClientVersion *Version             `json:"clientVersion,omitempty"`

	// ClusterType ClusterType specifies the type of cluster sharing
	ClusterType *ClusterType `json:"clusterType,omitempty"`

	// Components Components is a reference to the collection of components detected to be running in the cluster
	Components *string `json:"components,omitempty"`

	// Description Description is the humanreadable cluster description of the cluster
	Description *string `json:"description,omitempty"`

	// EnvironmentName EnvironmentName specifies the name of the environment to which the cluster is associated - this is generic determined by the cluster provider and tenants but suggested to be, e.g., "production"
	EnvironmentName *string `json:"environmentName,omitempty"`

	// Fqdn FQDN is deprecated, was a string identifying the cluster with a DNS valid string often also registered in DNS - this will not be provided going forward please use the NRN
	Fqdn *string `json:"fqdn,omitempty"`

	// Images Images is a reference to the collection of images discovered from the cluster
	Images *string `json:"images,omitempty"`

	// InfrastructureProvider InfrastructureProvider indicates the provider of the underlying infrastructure, e.g., azure, aws, netic etc. it may not be possible to determine the infrastructure provider
	InfrastructureProvider *string `json:"infrastructureProvider,omitempty"`

	// KubernetesProvider KubernetesProvider indicates the Kubernetes distribution, e.g., rke2, aks, eks, kubeadm, etc.
	KubernetesProvider *string  `json:"kubernetesProvider,omitempty"`
	KubernetesVersion  *Version `json:"kubernetesVersion,omitempty"`

	// Name Name uniquely identifying the cluster within the provider of the cluster
	Name *string `json:"name,omitempty"`

	// Namespaces Namespaces is a reference to the collection of namespace for the cluster
	Namespaces *string `json:"namespaces,omitempty"`

	// Nrn NRN is the Netic Resource Name uniquely identifying the cluster among all other Netic resource names
	Nrn *string `json:"nrn,omitempty"`

	// Provider Provider identification of cluster provider, i.e., the organizational unit responsible for providing the cluster to the tenants
	Provider *string `json:"provider,omitempty"`

	// ResilienceZone ResilienceZone identified the resilienze to which the cluster is associated
	ResilienceZone *string `json:"resilienceZone,omitempty"`

	// Subscription Subscription is the Netic subscription identifier used for accounting - this will only be included when the client is authorized for this information
	Subscription *string `json:"subscription,omitempty"`

	// Timestamp Timestamp indicating when data was collected
	Timestamp *time.Time `json:"timestamp,omitempty"`

	// Vulnerabilities Vulnerabilities is a reference to the collection of vulnerabilities detected from the cluster
	Vulnerabilities *string `json:"vulnerabilities,omitempty"`

	// Workloads Workloads is a reference to the collection of workloads for the cluster
	Workloads *string `json:"workloads,omitempty"`
}

// ClusterType ClusterType specifies the type of cluster sharing
type ClusterType = string

// Clusters Clusters represents a collection of clusters
type Clusters struct {
	// Context Context is defining the JSON-LD context for dereferencing the data as stated in the JSON-LD specification https://www.w3.org/TR/json-ld/#keywords
	Context *map[string]interface{} `json:"@context,omitempty"`

	// Id ID is identifying the node with an IRI as stated in the JSON-LD specification https://www.w3.org/TR/json-ld/#keywords
	Id *string `json:"@id,omitempty"`

	// Included Included will container linked resources included here for convenience
	Included *[]map[string]interface{} `json:"@included,omitempty"`

	// Type Type is optional explicit definition of the type of node as stated in the JSON-LD specification https://www.w3.org/TR/json-ld/#keywords
	Type *string `json:"@type,omitempty"`

	// Clusters Clusters is the identification of the clusters in the collection
	Clusters *[]string `json:"clusters,omitempty"`

	// Count Count is the number of clusters in the collection
	Count *int32 `json:"count,omitempty"`

	// Pagination Pagination contains data on other collection data
	Pagination *Pagination `json:"pagination,omitempty"`

	// Total TotalCount is the total number of clusters
	Total *int32 `json:"total,omitempty"`
}

// Pagination Pagination contains data on other collection data
type Pagination struct {
	// First First is link to the first page in the collection
	First *string `json:"first,omitempty"`

	// Last Last is link to the last page in the collection
	Last *string `json:"last,omitempty"`

	// Next Next is link to the next page, if any
	Next *string `json:"next,omitempty"`

	// Prev Prev is link to the previous page, if any
	Prev *string `json:"prev,omitempty"`
}

// Problem Problem is simple implementation of (RFC9457)[https://datatracker.ietf.org/doc/html/rfc9457]
type Problem struct {
	// Detail Detail is humanreadable explanation of the specific occurence of the problem RFC-9457#3.1.4
	Detail *string `json:"detail,omitempty"`

	// Instance Instance identifies the specific instance of the problem RFC-9457#3.1.5
	Instance *string `json:"instance,omitempty"`

	// Status Status is the http status code and must be consistent with the server status code RFC-9457#3.1.2
	Status *int32 `json:"status,omitempty"`

	// Title Title is short humanreadable summary RFC-9457#3.1.3
	Title *string `json:"title,omitempty"`

	// Type Type identify problem type RFC-9457#3.1.1
	Type *string `json:"type,omitempty"`
}

// Capacity defines model for capacity.
type Capacity struct {
	// Cores CPU is the total amount of allocatable millicores
	Cores *int64 `json:"cores,omitempty"`

	// Memory Memory is the total amount of allocatable bytes of memory
	Memory *int64 `json:"memory,omitempty"`

	// Nodes Nodes is the total number of nodes
	Nodes *int64 `json:"nodes,omitempty"`
}

// Version defines model for version.
type Version struct {
	// Commit Commit is the commit sha of the source code for the binary
	Commit *string `json:"commit,omitempty"`

	// Date Date is the date the binary was build
	Date *string `json:"date,omitempty"`

	// Version Version is the semver version
	Version *string `json:"version,omitempty"`
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// ListClusters request
	ListClusters(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetCluster request
	GetCluster(ctx context.Context, clusterId string, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) ListClusters(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListClustersRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetCluster(ctx context.Context, clusterId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetClusterRequest(c.Server, clusterId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewListClustersRequest generates requests for ListClusters
func NewListClustersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/clusters")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetClusterRequest generates requests for GetCluster
func NewGetClusterRequest(server string, clusterId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "clusterId", runtime.ParamLocationPath, clusterId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/clusters/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// ListClustersWithResponse request
	ListClustersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListClustersResponse, error)

	// GetClusterWithResponse request
	GetClusterWithResponse(ctx context.Context, clusterId string, reqEditors ...RequestEditorFn) (*GetClusterResponse, error)
}

type ListClustersResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	ApplicationldJSON400          *Problem
	ApplicationproblemJSON400     *Problem
	ApplicationldJSON401          *Problem
	ApplicationproblemJSON401     *Problem
	ApplicationldJSON500          *Problem
	ApplicationproblemJSON500     *Problem
	ApplicationldJSONDefault      *Clusters
	ApplicationproblemJSONDefault *Clusters
}

// Status returns HTTPResponse.Status
func (r ListClustersResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListClustersResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetClusterResponse struct {
	Body                          []byte
	HTTPResponse                  *http.Response
	ApplicationldJSON401          *Problem
	ApplicationproblemJSON401     *Problem
	ApplicationldJSON404          *Problem
	ApplicationproblemJSON404     *Problem
	ApplicationldJSON500          *Problem
	ApplicationproblemJSON500     *Problem
	ApplicationldJSONDefault      *Cluster
	ApplicationproblemJSONDefault *Cluster
}

// Status returns HTTPResponse.Status
func (r GetClusterResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetClusterResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// ListClustersWithResponse request returning *ListClustersResponse
func (c *ClientWithResponses) ListClustersWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListClustersResponse, error) {
	rsp, err := c.ListClusters(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListClustersResponse(rsp)
}

// GetClusterWithResponse request returning *GetClusterResponse
func (c *ClientWithResponses) GetClusterWithResponse(ctx context.Context, clusterId string, reqEditors ...RequestEditorFn) (*GetClusterResponse, error) {
	rsp, err := c.GetCluster(ctx, clusterId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetClusterResponse(rsp)
}

// ParseListClustersResponse parses an HTTP response from a ListClustersWithResponse call
func ParseListClustersResponse(rsp *http.Response) (*ListClustersResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListClustersResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.Header.Get("Content-Type") == "application/ld+json" && rsp.StatusCode == 400:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationldJSON400 = &dest

	case rsp.Header.Get("Content-Type") == "application/ld+json" && rsp.StatusCode == 401:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationldJSON401 = &dest

	case rsp.Header.Get("Content-Type") == "application/ld+json" && rsp.StatusCode == 500:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationldJSON500 = &dest

	case rsp.Header.Get("Content-Type") == "application/ld+json" && true:
		var dest Clusters
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationldJSONDefault = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && rsp.StatusCode == 400:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON400 = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && rsp.StatusCode == 401:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON401 = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && rsp.StatusCode == 500:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && true:
		var dest Clusters
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}

// ParseGetClusterResponse parses an HTTP response from a GetClusterWithResponse call
func ParseGetClusterResponse(rsp *http.Response) (*GetClusterResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetClusterResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case rsp.Header.Get("Content-Type") == "application/ld+json" && rsp.StatusCode == 401:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationldJSON401 = &dest

	case rsp.Header.Get("Content-Type") == "application/ld+json" && rsp.StatusCode == 404:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationldJSON404 = &dest

	case rsp.Header.Get("Content-Type") == "application/ld+json" && rsp.StatusCode == 500:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationldJSON500 = &dest

	case rsp.Header.Get("Content-Type") == "application/ld+json" && true:
		var dest Cluster
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationldJSONDefault = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && rsp.StatusCode == 401:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON401 = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && rsp.StatusCode == 404:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON404 = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && rsp.StatusCode == 500:
		var dest Problem
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSON500 = &dest

	case rsp.Header.Get("Content-Type") == "application/problem+json" && true:
		var dest Cluster
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.ApplicationproblemJSONDefault = &dest

	}

	return response, nil
}
