// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/rai-project/dlframework/httpapi/restapi/operations/authentication"
	"github.com/rai-project/dlframework/httpapi/restapi/operations/predict"
	"github.com/rai-project/dlframework/httpapi/restapi/operations/registry"
)

// NewDlframeworkAPI creates a new Dlframework instance
func NewDlframeworkAPI(spec *loads.Document) *DlframeworkAPI {
	return &DlframeworkAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		PredictCloseHandler: predict.CloseHandlerFunc(func(params predict.CloseParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictClose has not yet been implemented")
		}),
		PredictDatasetHandler: predict.DatasetHandlerFunc(func(params predict.DatasetParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictDataset has not yet been implemented")
		}),
		PredictDatasetStreamHandler: predict.DatasetStreamHandlerFunc(func(params predict.DatasetStreamParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictDatasetStream has not yet been implemented")
		}),
		RegistryFrameworkAgentsHandler: registry.FrameworkAgentsHandlerFunc(func(params registry.FrameworkAgentsParams) middleware.Responder {
			return middleware.NotImplemented("operation RegistryFrameworkAgents has not yet been implemented")
		}),
		RegistryFrameworkManifestsHandler: registry.FrameworkManifestsHandlerFunc(func(params registry.FrameworkManifestsParams) middleware.Responder {
			return middleware.NotImplemented("operation RegistryFrameworkManifests has not yet been implemented")
		}),
		PredictImagesHandler: predict.ImagesHandlerFunc(func(params predict.ImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictImages has not yet been implemented")
		}),
		PredictImagesStreamHandler: predict.ImagesStreamHandlerFunc(func(params predict.ImagesStreamParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictImagesStream has not yet been implemented")
		}),
		AuthenticationLoginHandler: authentication.LoginHandlerFunc(func(params authentication.LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthenticationLogin has not yet been implemented")
		}),
		RegistryModelAgentsHandler: registry.ModelAgentsHandlerFunc(func(params registry.ModelAgentsParams) middleware.Responder {
			return middleware.NotImplemented("operation RegistryModelAgents has not yet been implemented")
		}),
		RegistryModelManifestsHandler: registry.ModelManifestsHandlerFunc(func(params registry.ModelManifestsParams) middleware.Responder {
			return middleware.NotImplemented("operation RegistryModelManifests has not yet been implemented")
		}),
		PredictOpenHandler: predict.OpenHandlerFunc(func(params predict.OpenParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictOpen has not yet been implemented")
		}),
		PredictResetHandler: predict.ResetHandlerFunc(func(params predict.ResetParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictReset has not yet been implemented")
		}),
		AuthenticationSignupHandler: authentication.SignupHandlerFunc(func(params authentication.SignupParams) middleware.Responder {
			return middleware.NotImplemented("operation AuthenticationSignup has not yet been implemented")
		}),
		PredictUrlsHandler: predict.UrlsHandlerFunc(func(params predict.UrlsParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictUrls has not yet been implemented")
		}),
		PredictUrlsStreamHandler: predict.UrlsStreamHandlerFunc(func(params predict.UrlsStreamParams) middleware.Responder {
			return middleware.NotImplemented("operation PredictUrlsStream has not yet been implemented")
		}),
	}
}

/*DlframeworkAPI MLModelScope is a hardware/software agnostic platform to facilitate the evaluation, measurement, and introspection of ML models within AI pipelines. MLModelScope aids application developers in discovering and experimenting with models, data scientists developers in replicating and evaluating for publishing models, and system architects in understanding the performance of AI workloads. */
type DlframeworkAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// PredictCloseHandler sets the operation handler for the close operation
	PredictCloseHandler predict.CloseHandler
	// PredictDatasetHandler sets the operation handler for the dataset operation
	PredictDatasetHandler predict.DatasetHandler
	// PredictDatasetStreamHandler sets the operation handler for the dataset stream operation
	PredictDatasetStreamHandler predict.DatasetStreamHandler
	// RegistryFrameworkAgentsHandler sets the operation handler for the framework agents operation
	RegistryFrameworkAgentsHandler registry.FrameworkAgentsHandler
	// RegistryFrameworkManifestsHandler sets the operation handler for the framework manifests operation
	RegistryFrameworkManifestsHandler registry.FrameworkManifestsHandler
	// PredictImagesHandler sets the operation handler for the images operation
	PredictImagesHandler predict.ImagesHandler
	// PredictImagesStreamHandler sets the operation handler for the images stream operation
	PredictImagesStreamHandler predict.ImagesStreamHandler
	// AuthenticationLoginHandler sets the operation handler for the login operation
	AuthenticationLoginHandler authentication.LoginHandler
	// RegistryModelAgentsHandler sets the operation handler for the model agents operation
	RegistryModelAgentsHandler registry.ModelAgentsHandler
	// RegistryModelManifestsHandler sets the operation handler for the model manifests operation
	RegistryModelManifestsHandler registry.ModelManifestsHandler
	// PredictOpenHandler sets the operation handler for the open operation
	PredictOpenHandler predict.OpenHandler
	// PredictResetHandler sets the operation handler for the reset operation
	PredictResetHandler predict.ResetHandler
	// AuthenticationSignupHandler sets the operation handler for the signup operation
	AuthenticationSignupHandler authentication.SignupHandler
	// PredictUrlsHandler sets the operation handler for the urls operation
	PredictUrlsHandler predict.UrlsHandler
	// PredictUrlsStreamHandler sets the operation handler for the urls stream operation
	PredictUrlsStreamHandler predict.UrlsStreamHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *DlframeworkAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *DlframeworkAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *DlframeworkAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *DlframeworkAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *DlframeworkAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *DlframeworkAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *DlframeworkAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the DlframeworkAPI
func (o *DlframeworkAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.PredictCloseHandler == nil {
		unregistered = append(unregistered, "predict.CloseHandler")
	}

	if o.PredictDatasetHandler == nil {
		unregistered = append(unregistered, "predict.DatasetHandler")
	}

	if o.PredictDatasetStreamHandler == nil {
		unregistered = append(unregistered, "predict.DatasetStreamHandler")
	}

	if o.RegistryFrameworkAgentsHandler == nil {
		unregistered = append(unregistered, "registry.FrameworkAgentsHandler")
	}

	if o.RegistryFrameworkManifestsHandler == nil {
		unregistered = append(unregistered, "registry.FrameworkManifestsHandler")
	}

	if o.PredictImagesHandler == nil {
		unregistered = append(unregistered, "predict.ImagesHandler")
	}

	if o.PredictImagesStreamHandler == nil {
		unregistered = append(unregistered, "predict.ImagesStreamHandler")
	}

	if o.AuthenticationLoginHandler == nil {
		unregistered = append(unregistered, "authentication.LoginHandler")
	}

	if o.RegistryModelAgentsHandler == nil {
		unregistered = append(unregistered, "registry.ModelAgentsHandler")
	}

	if o.RegistryModelManifestsHandler == nil {
		unregistered = append(unregistered, "registry.ModelManifestsHandler")
	}

	if o.PredictOpenHandler == nil {
		unregistered = append(unregistered, "predict.OpenHandler")
	}

	if o.PredictResetHandler == nil {
		unregistered = append(unregistered, "predict.ResetHandler")
	}

	if o.AuthenticationSignupHandler == nil {
		unregistered = append(unregistered, "authentication.SignupHandler")
	}

	if o.PredictUrlsHandler == nil {
		unregistered = append(unregistered, "predict.UrlsHandler")
	}

	if o.PredictUrlsStreamHandler == nil {
		unregistered = append(unregistered, "predict.UrlsStreamHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *DlframeworkAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *DlframeworkAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *DlframeworkAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *DlframeworkAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *DlframeworkAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *DlframeworkAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the dlframework API
func (o *DlframeworkAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *DlframeworkAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/close"] = predict.NewClose(o.context, o.PredictCloseHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/dataset"] = predict.NewDataset(o.context, o.PredictDatasetHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/stream/dataset"] = predict.NewDatasetStream(o.context, o.PredictDatasetStreamHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/registry/frameworks/agent"] = registry.NewFrameworkAgents(o.context, o.RegistryFrameworkAgentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/registry/frameworks/manifest"] = registry.NewFrameworkManifests(o.context, o.RegistryFrameworkManifestsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/images"] = predict.NewImages(o.context, o.PredictImagesHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/stream/images"] = predict.NewImagesStream(o.context, o.PredictImagesStreamHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/login"] = authentication.NewLogin(o.context, o.AuthenticationLoginHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/registry/models/agent"] = registry.NewModelAgents(o.context, o.RegistryModelAgentsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/registry/models/manifest"] = registry.NewModelManifests(o.context, o.RegistryModelManifestsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/open"] = predict.NewOpen(o.context, o.PredictOpenHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/reset"] = predict.NewReset(o.context, o.PredictResetHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/auth/signup"] = authentication.NewSignup(o.context, o.AuthenticationSignupHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/urls"] = predict.NewUrls(o.context, o.PredictUrlsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/predict/stream/urls"] = predict.NewUrlsStream(o.context, o.PredictUrlsStreamHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *DlframeworkAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *DlframeworkAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *DlframeworkAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *DlframeworkAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
