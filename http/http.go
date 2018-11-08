package http

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/rai-project/dlframework/httpapi/restapi/operations"
	"github.com/rai-project/dlframework/httpapi/restapi/operations/authentication"
	"github.com/rai-project/dlframework/httpapi/restapi/operations/predict"
	"github.com/rai-project/dlframework/httpapi/restapi/operations/registry"
)

func ConfigureAPI(api *operations.DlframeworkAPI) http.Handler {
	api.ServeError = ServeError
	api.Logger = log.Debugf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AuthenticationLoginHandler = authentication.LoginHandlerFunc(LoginHandler)
	api.AuthenticationSignupHandler = authentication.SignupHandlerFunc(SignupHandler)

	api.RegistryFrameworkAgentsHandler = registry.FrameworkAgentsHandlerFunc(RegistryFrameworkAgentsHandler)
	api.RegistryFrameworkManifestsHandler = registry.FrameworkManifestsHandlerFunc(RegistryFrameworkManifestsHandler)
	api.RegistryModelAgentsHandler = registry.ModelAgentsHandlerFunc(RegistryModelAgentsHandler)
	api.RegistryModelManifestsHandler = registry.ModelManifestsHandlerFunc(RegistryModelManifestsHandler)

	predictHandler := &PredictHandler{}
	api.PredictOpenHandler = predict.OpenHandlerFunc(predictHandler.Open)
	api.PredictCloseHandler = predict.CloseHandlerFunc(predictHandler.Close)
	api.PredictResetHandler = predict.ResetHandlerFunc(predictHandler.Reset)
	api.PredictImagesHandler = predict.ImagesHandlerFunc(predictHandler.Images)
	api.PredictUrlsHandler = predict.UrlsHandlerFunc(predictHandler.URLs)
	api.PredictDatasetHandler = predict.DatasetHandlerFunc(predictHandler.Dataset)

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}
