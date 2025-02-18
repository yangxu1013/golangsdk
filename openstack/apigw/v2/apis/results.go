package apis

import (
	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/pagination"
)

type commonResult struct {
	golangsdk.Result
}

// CreateResult represents a result of the Create method.
type CreateResult struct {
	commonResult
}

// GetResult represents a result of the Get method.
type GetResult struct {
	commonResult
}

// UpdateResult represents a result of the Update method.
type UpdateResult struct {
	commonResult
}

// APIResp is a struct that represents the result of Create, Update, Get and List methods.
type APIResp struct {
	// API name which can contains of 3 to 64 characters, starting with a letter.
	// Only letters, digits, and underscores (_) are allowed.
	// Chinese characters must be in UTF-8 or Unicode format.
	Name string `json:"name"`
	// API type. The valid types are as following:
	//   1: public API
	//   2: private API
	Type int `json:"type"`
	// API version which can contains maximum of 16 characters.
	Version string `json:"version"`
	// Request protocol.
	//   HTTP
	//   HTTPS (default)
	//   BOTH: The API can be accessed through both HTTP and HTTPS.
	ReqProtocol string `json:"req_protocol"`
	// Request method. The valid values are GET, POST, PUT, DELETE, HEAD, PATCH, OPTIONS and ANY.
	ReqMethod string `json:"req_method"`
	// Request address, which can contain request parameters enclosed with brackets ({}). For example,
	//   /getUserInfo/{userId}.
	// The request address can contain special characters, such as asterisks (*), percent signs (%), hyphens (-), and underscores (_). It can contain a maximum of 512 characters and must comply with URI specifications.
	// The request address must comply with URI specifications.
	ReqURI string `json:"req_uri"`
	// Security authentication mode. The valid values are as following:
	//   NONE
	//   APP
	//   IAM
	//   AUTHORIZER
	AuthType string `json:"auth_type"`
	// Security authentication parameter.
	AuthOpt AuthOpt `json:"auth_opt"`
	// Indicates whether CORS is supported.
	// TRUE: supported
	// FALSE: not supported (default).
	Cors bool `json:"cors"`
	// Route matching mode.
	//   SWA: prefix match
	//   NORMAL: exact match (default).
	MatchMode string `json:"match_mode"`
	// Backend type. The valid types are as following:
	//   HTTP: web backend
	//   FUNCTION: FunctionGraph backend
	//   MOCK: Mock backend
	BackendType string `json:"backend_type"`
	// Description of the API, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description string `json:"remark"`
	// ID of the API group to which the API belongs.
	GroupId string `json:"group_id"`
	// API request body, which can be an example request body, media type, or parameters.
	// Ensure that the request body does not exceed 20,480 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	BodyDescription string `json:"body_remark"`
	// Example response for a successful request. Ensure that the response does not exceed 20,480 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	ResultNormalSample string `json:"result_normal_sample"`
	// Example response for a failed request. Ensure that the response does not exceed 20,480 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	ResultFailureSample string `json:"result_failure_sample"`
	// ID of the frontend custom authorizer.
	AuthorizerId string `json:"authorizer_id"`
	// Tags.
	Tags []string `json:"tags"`
	// Group response ID.
	ResponseId string `json:"response_id"`
	// API ID.
	ID string `json:"id"`
	// API status. 1: valid
	Status int `json:"status"`
	// Indicates whether to enable orchestration.
	ArrangeNecessary int `json:"arrange_necessary"`
	// Time when the API is registered.
	RegisterTime string `json:"register_time"`
	// Time when the API was last modified.
	UpdateTime string `json:"update_time"`
	// Name of the API group to which the API belongs.
	GroupName string `json:"group_name"`
	// Version of the API group to which the API belongs.
	// The default value is V1. Other versions are not supported.
	GroupVersion string `json:"group_version"`
	// ID of the environment in which the API has been published.
	// If there are multiple publication records, separate the environment IDs with vertical bars (|).
	RunEnvId string `json:"run_env_id"`
	// Name of the environment in which the API has been published.
	// If there are multiple publication records, separate the environment names with vertical bars (|).
	RunEnvName string `json:"run_env_name"`
	// Publication record ID.
	// You can separate multiple publication record IDs with vertical bars (|).
	PublishId string `json:"publish_id"`
	// FunctionGraph backend details.
	FuncInfo FuncGraph `json:"func_info"`
	// Mock backend details.
	MockInfo Mock `json:"mock_info"`
	// Web backend details.
	WebInfo Web `json:"backend_api"`
	// Request parameters.
	ReqParams []ReqParamResp `json:"req_params"`
	// Backend parameters.
	BackendParams []BackendParamResp `json:"backend_params"`
	// Mock policy backends.
	PolicyMocks []PolicyMockResp `json:"policy_mocks"`
	// FunctionGraph policy backends.
	PolicyFunctions []PolicyFuncGraphResp `json:"policy_functions"`
	// Web policy backends.
	PolicyWebs []PolicyWebResp `json:"policy_https"`
	// The following four parameters are only provided by the response of the API Versions.
	//   SlDomain
	//   SlDomains
	//   VersionId
	//   PublishTime.
	// Subdomain name that API Gateway automatically allocates to the API group.
	SlDomain string `json:"sl_domain"`
	// Subdomain names that API Gateway automatically allocates to the API group.
	SlDomains []string `json:"sl_domains"`
	// API version ID.
	VersionId string `json:"version_id"`
	//Time when the API version is published.
	PublishTime string `json:"publish_time"`
}

// ReqParamResp is an object struct that represents the elements of the front-end request parameter.
type ReqParamResp struct {
	// The parameter name, which contain of 1 to 32 characters and start with a letter.
	// Only letters, digits, hyphens (-), underscores (_), and periods (.) are allowed.
	Name string `json:"name"`
	// Parameter type. The valid types are as following:
	//   STRING
	//   NUMBER
	Type string `json:"type"`
	// Parameter location. The valid modes are as following:
	//   PATH
	//   QUERY
	//   HEADER
	Location string `json:"location"`
	// Default value.
	DefaultValue string `json:"default_value"`
	// Example value.
	SampleValue string `json:"sample_value"`
	// Indicates whether the parameter is required. The valid values are 1 (yes) and 2 (no).
	// The value of this parameter is 1 if Location is set to PATH, and 2 if Location is set to another value.
	Required int `json:"required"`
	// Indicates whether validity check is enabled. The valid modes are as following:
	// 1: enabled.
	// 2: disabled (default).
	ValidEnable int `json:"valid_enable"`
	// Description about the backend, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description string `json:"remark"`
	// Enumerated value.
	Enumerations string `json:"enumerations"`
	// Minimum value.
	// This parameter is valid when type is set to NUMBER.
	MinNum int `json:"min_num"`
	// Maximum value.
	// This parameter is valid when type is set to NUMBER.
	MaxNum int `json:"max_num"`
	// Minimum length.
	// This parameter is valid when type is set to STRING.
	MinSize int `json:"min_size"`
	// Maximum length.
	// This parameter is valid when type is set to STRING.
	MaxSize int `json:"max_size"`
	// Indicates whether to transparently transfer the parameter. The valid values are 1 (yes) and 2 (no).
	PassThrough int `json:"pass_through"`
	// Parameter ID.
	// Notes: This parameter is used for response.
	ID string `json:"id"`
}

// BackendParamResp is an object struct that represents the elements of the back-end parameter.
type BackendParamResp struct {
	// Parameter type. The valid types are as following:
	//   REQUEST: Backend parameter.
	//   CONSTANT: Constant parameter.
	//   SYSTEM: System parameter.
	Origin string `json:"origin"`
	// Parameter name, which can contains 1 to 32 characters, must start with a letter and can only contain letters, digits, hyphens (-), underscores (_),
	// and periods (.).
	Name string `json:"name"`
	// Description, which can contain a maximum of 255 characters.
	// Chinese characters must be in UTF-8 or Unicode format.
	Description string `json:"remark"`
	// Parameter location. The valid values are PATH, QUERY and HEADER.
	Location string `json:"location"`
	// Parameter value, which can contain a maximum of 255 characters.
	// If the origin type is REQUEST, the value of this parameter is the parameter name in req_params.
	// If the origin type is CONSTANT, the value is a constant.
	// If the origin type is SYSTEM, the value is a system parameter name.
	// System parameters include gateway parameters, front-end authentication parameters, and back-end authentication
	// parameters. You can set the frontend or backend authentication parameters after enabling custom frontend or
	// backend authentication.
	// The gateway parameters are as follows:
	//   $context.sourceIp: source IP address of the API caller.
	//   $context.stage: deployment environment in which the API is called.
	//   $context.apiId: API ID.
	//   $context.appId: ID of the app used by the API caller.
	//   $context.requestId: request ID generated when the API is called.
	//   $context.serverAddr: address of the gateway server.
	//   $context.serverName: name of the gateway server.
	//   $context.handleTime: time when the API request is processed.
	//   $context.providerAppId: ID of the app used by the API owner. This parameter is currently not supported.
	// Frontend authentication parameter: prefixed with "$context.authorizer.frontend.".
	// For example, to return "aaa" upon successful custom authentication, set this parameter to
	// "$context.authorizer.frontend.aaa".
	// Backend authentication parameter: prefixed with "$context.authorizer.backend.".
	// For example, to return "aaa" upon successful custom authentication, set this parameter to
	// "$context.authorizer.backend.aaa".
	Value string `json:"value"`
	// ID of the the specifies request parameter.
	ReqParamId string `json:"req_param_id"`
}

// PolicyMockResp is an object struct that represents the back-end policy of mock.
type PolicyMockResp struct {
	// Policy conditions.
	Conditions []APIConditionBase `json:"conditions" required:"true"`
	// Effective mode of the backend policy. The valid modes are as following:
	//   ALL: All conditions are met.
	//   ANY: Any condition is met.
	EffectMode string `json:"effect_mode" required:"true"`
	// Backend name, which consists of 3 to 64 characters and must start with a letter and can contain letters, digits,
	// and underscores (_).
	Name string `json:"name" required:"true"`
	// Authorizer ID.
	AuthorizerId string `json:"authorizer_id,omitempty"`
	// Backend parameters.
	BackendParams []BackendParamResp `json:"backend_params,omitempty"`
	// Response.
	ResultContent string `json:"result_content,omitempty"`
}

// PolicyFuncGraphResp is an object struct that represents the back-end policy of function graph.
type PolicyFuncGraphResp struct {
	// Policy conditions.
	Conditions []APIConditionBase `json:"conditions" required:"true"`
	// Effective mode of the backend policy.
	//   ALL: All conditions are met.
	//   ANY: Any condition is met.
	EffectMode string `json:"effect_mode" required:"true"`
	// Function URN.
	FunctionUrn string `json:"function_urn" required:"true"`
	// Invocation mode. The valid modes are as following:
	//   async: asynchronous
	//   sync: synchronous
	InvocationType string `json:"invocation_type" required:"true"`
	// The backend name, which can consists of 3 to 64 characters and must start with a letter and can contain letters,
	// digits, and underscores (_).
	Name string `json:"name" required:"true"`
	// Authorizer ID.
	AuthorizerId string `json:"authorizer_id,omitempty"`
	// Backend parameters.
	BackendParams []BackendParamResp `json:"backend_params,omitempty"`
	// Timeout, in ms, which allowed for API Gateway to request the backend service.
	// The valid value is range from 1 to 600,000.
	Timeout int `json:"timeout,omitempty"`
	// Function version Ensure that the version does not exceed 64 characters.
	Version string `json:"version,omitempty"`
}

// PolicyWebResp is an object struct that represents the back-end policy of http and https.
type PolicyWebResp struct {
	// Request protocol. The value can be HTTP or HTTPS.
	ReqProtocol string `json:"req_protocol" required:"true"`
	// Request method. The valid methods are GET, POST, PUT, DELETE, HEAD, PATCH, OPTIONS and ANY.
	ReqMethod string `json:"req_method" required:"true"`
	// Request address, which can contain request parameters enclosed with brackets ({}). For example,
	//   /getUserInfo/{userId}.
	// The request address can contain special characters, such as asterisks (), percent signs (%), hyphens (-), and
	// underscores (_). It can contain a maximum of 512 characters and must comply with URI specifications.
	// The request address can contain environment variables, each starting with a letter and consisting of 3 to 32
	// characters. Only letters, digits, hyphens (-), and underscores (_) are allowed in environment variables.
	// The request address must comply with URI specifications.
	ReqURI string `json:"req_uri" required:"true"`
	// Endpoint of the policy backend.
	// An endpoint consists of a domain name or IP address and a port number, with not more than 255 characters.
	// It must be in the format "Domain name:Port number", for example, apig.example.com:7443.
	// If the port number is not specified, the default HTTPS port 443 or the default HTTP port 80 is used.
	// The endpoint can contain environment variables, each starting with letter and consisting of 3 to 32 characters.
	// Only letters, digits, hyphens (-), and underscores (_) are allowed.
	DomainURL string `json:"url_domain,omitempty"`
	// Timeout, in ms, which allowed for API Gateway to request the backend service.
	// The valid value is range from 1 to 600,000.
	Timeout int `json:"timeout,omitempty"`
	// Effective mode of the backend policy. The valid modes are as following:
	//   ALL: All conditions are met.
	//   ANY: Any condition is met.
	EffectMode string `json:"effect_mode" required:"true"`
	// Backend name, which contains of 3 to 64 and must start with a letter and can contain letters, digits, and
	// underscores (_).
	Name string `json:"name" required:"true"`
	// Backend parameters.
	BackendParams []BackendParamResp `json:"backend_params,omitempty"`
	// Policy conditions.
	Conditions []APIConditionBase `json:"conditions" required:"true"`
	// Authorizer ID.
	AuthorizerId string `json:"authorizer_id,omitempty"`
	// VPC channel details. This parameter is required if vpc_channel_status is set to 1.
	VpcChannelInfo VpcChannel `json:"vpc_channel_info,omitempty"`
	// Indicates whether to use a VPC channel. The valid value are as following:
	//   1: A VPC channel is used.
	//   2: No VPC channel is used.
	VpcChannelEnable int `json:"vpc_channel_status,omitempty"`
}

// Extract is a method to extract an response struct.
func (r commonResult) Extract() (*APIResp, error) {
	var s APIResp
	err := r.ExtractInto(&s)
	return &s, err
}

// APIPage represents the api pages of the List operation.
type APIPage struct {
	pagination.SinglePageBase
}

// ExtractApis is a method to extract an response struct list.
func ExtractApis(r pagination.Page) ([]APIResp, error) {
	var s []APIResp
	err := r.(APIPage).Result.ExtractIntoSlicePtr(&s, "apis")
	return s, err
}

// DeleteResult represents a result of the Delete method.
type DeleteResult struct {
	golangsdk.ErrResult
}
