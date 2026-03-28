package fiberext

import (
	"github.com/gofiber/fiber/v2"
)

// Continue responds with 100 Continue.
func Continue(ctx Context) error {
	return Respond(ctx, fiber.StatusContinue, nil)
}

// SwitchingProtocols responds with 101 Switching Protocols.
func SwitchingProtocols(ctx Context) error {
	return Respond(ctx, fiber.StatusSwitchingProtocols, nil)
}

// Processing responds with 102 Processing.
func Processing(ctx Context) error {
	return Respond(ctx, fiber.StatusProcessing, nil)
}

// EarlyHints responds with 103 Early Hints.
func EarlyHints(ctx Context) error {
	return Respond(ctx, fiber.StatusEarlyHints, nil)
}

// OK responds with 200 OK and serializes obj as JSON.
func OK(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusOK, obj...)
}

// Created responds with 201 Created and serializes obj as JSON.
func Created(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusCreated, obj...)
}

// Accepted responds with 202 Accepted and serializes obj as JSON.
func Accepted(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusAccepted, obj...)
}

// NonAuthoritativeInformation responds with 203 and serializes obj as JSON.
func NonAuthoritativeInformation(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusNonAuthoritativeInformation, obj...)
}

// NoContent responds with 204 No Content.
func NoContent(ctx Context) error {
	return Respond(ctx, fiber.StatusNoContent, nil)
}

// ResetContent responds with 205 Reset Content.
func ResetContent(ctx Context) error {
	return Respond(ctx, fiber.StatusResetContent, nil)
}

// PartialContent responds with 206 Partial Content and serializes obj as JSON.
func PartialContent(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusPartialContent, obj...)
}

// MultiStatus responds with 207 Multi-Status and serializes obj as JSON.
func MultiStatus(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusMultiStatus, obj...)
}

// AlreadyReported responds with 208 Already Reported and serializes obj as JSON.
func AlreadyReported(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusAlreadyReported, obj...)
}

// IMUsed responds with 226 IM Used and serializes obj as JSON.
func IMUsed(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusIMUsed, obj...)
}

// MultipleChoices responds with 300 Multiple Choices and serializes obj as JSON.
func MultipleChoices(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusMultipleChoices, obj...)
}

// MovedPermanently responds with 301 Moved Permanently.
func MovedPermanently(ctx Context) error {
	return Respond(ctx, fiber.StatusMovedPermanently, nil)
}

// Found responds with 302 Found.
func Found(ctx Context) error {
	return Respond(ctx, fiber.StatusFound, nil)
}

// SeeOther responds with 303 See Other.
func SeeOther(ctx Context) error {
	return Respond(ctx, fiber.StatusSeeOther, nil)
}

// NotModified responds with 304 Not Modified.
func NotModified(ctx Context) error {
	return Respond(ctx, fiber.StatusNotModified, nil)
}

// UseProxy responds with 305 Use Proxy.
func UseProxy(ctx Context) error {
	return Respond(ctx, fiber.StatusUseProxy, nil)
}

// SwitchProxy responds with 306 Switch Proxy.
func SwitchProxy(ctx Context) error {
	return Respond(ctx, fiber.StatusSwitchProxy, nil)
}

// TemporaryRedirect responds with 307 Temporary Redirect.
func TemporaryRedirect(ctx Context) error {
	return Respond(ctx, fiber.StatusTemporaryRedirect, nil)
}

// PermanentRedirect responds with 308 Permanent Redirect.
func PermanentRedirect(ctx Context) error {
	return Respond(ctx, fiber.StatusPermanentRedirect, nil)
}

// BadRequest responds with 400 Bad Request.
func BadRequest(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusBadRequest, obj...)
}

// Unauthorized responds with 401 Unauthorized.
func Unauthorized(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUnauthorized, obj...)
}

// PaymentRequired responds with 402 Payment Required.
func PaymentRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusPaymentRequired, obj...)
}

// Forbidden responds with 403 Forbidden.
func Forbidden(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusForbidden, obj...)
}

// NotFound responds with 404 Not Found.
func NotFound(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNotFound, obj...)
}

// MethodNotAllowed responds with 405 Method Not Allowed.
func MethodNotAllowed(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusMethodNotAllowed, obj...)
}

// NotAcceptable responds with 406 Not Acceptable.
func NotAcceptable(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNotAcceptable, obj...)
}

// ProxyAuthRequired responds with 407 Proxy Authentication Required.
func ProxyAuthRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusProxyAuthRequired, obj...)
}

// RequestTimeout responds with 408 Request Timeout.
func RequestTimeout(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestTimeout, obj...)
}

// Conflict responds with 409 Conflict.
func Conflict(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusConflict, obj...)
}

// Gone responds with 410 Gone.
func Gone(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusGone, obj...)
}

// LengthRequired responds with 411 Length Required.
func LengthRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusLengthRequired, obj...)
}

// PreconditionFailed responds with 412 Precondition Failed.
func PreconditionFailed(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusPreconditionFailed, obj...)
}

// RequestEntityTooLarge responds with 413 Request Entity Too Large.
func RequestEntityTooLarge(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestEntityTooLarge, obj...)
}

// RequestURITooLong responds with 414 Request-URI Too Long.
func RequestURITooLong(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestURITooLong, obj...)
}

// UnsupportedMediaType responds with 415 Unsupported Media Type.
func UnsupportedMediaType(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUnsupportedMediaType, obj...)
}

// RequestedRangeNotSatisfiable responds with 416 Range Not Satisfiable.
func RequestedRangeNotSatisfiable(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestedRangeNotSatisfiable, obj...)
}

// ExpectationFailed responds with 417 Expectation Failed.
func ExpectationFailed(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusExpectationFailed, obj...)
}

// Teapot responds with 418 I'm a Teapot.
func Teapot(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusTeapot, obj...)
}

// MisdirectedRequest responds with 421 Misdirected Request.
func MisdirectedRequest(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusMisdirectedRequest, obj...)
}

// UnprocessableEntity responds with 422 Unprocessable Entity.
func UnprocessableEntity(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUnprocessableEntity, obj...)
}

// Locked responds with 423 Locked.
func Locked(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusLocked, obj...)
}

// FailedDependency responds with 424 Failed Dependency.
func FailedDependency(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusFailedDependency, obj...)
}

// TooEarly responds with 425 Too Early.
func TooEarly(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusTooEarly, obj...)
}

// UpgradeRequired responds with 426 Upgrade Required.
func UpgradeRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUpgradeRequired, obj...)
}

// PreconditionRequired responds with 428 Precondition Required.
func PreconditionRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusPreconditionRequired, obj...)
}

// TooManyRequests responds with 429 Too Many Requests.
func TooManyRequests(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusTooManyRequests, obj...)
}

// RequestHeaderFieldsTooLarge responds with 431 Request Header Fields Too Large.
func RequestHeaderFieldsTooLarge(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestHeaderFieldsTooLarge, obj...)
}

// UnavailableForLegalReasons responds with 451 Unavailable For Legal Reasons.
func UnavailableForLegalReasons(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUnavailableForLegalReasons, obj...)
}

// RespondInternalError responds with 500 Internal Server Error.
func RespondInternalError(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusInternalServerError, obj...)
}

// NotImplemented responds with 501 Not Implemented.
func NotImplemented(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNotImplemented, obj...)
}

// BadGateway responds with 502 Bad Gateway.
func BadGateway(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusBadGateway, obj...)
}

// ServiceUnavailable responds with 503 Service Unavailable.
func ServiceUnavailable(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusServiceUnavailable, obj...)
}

// GatewayTimeout responds with 504 Gateway Timeout.
func GatewayTimeout(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusGatewayTimeout, obj...)
}

// HTTPVersionNotSupported responds with 505 HTTP Version Not Supported.
func HTTPVersionNotSupported(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusHTTPVersionNotSupported, obj...)
}

// VariantAlsoNegotiates responds with 506 Variant Also Negotiates.
func VariantAlsoNegotiates(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusVariantAlsoNegotiates, obj...)
}

// InsufficientStorage responds with 507 Insufficient Storage.
func InsufficientStorage(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusInsufficientStorage, obj...)
}

// LoopDetected responds with 508 Loop Detected.
func LoopDetected(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusLoopDetected, obj...)
}

// NotExtended responds with 510 Not Extended.
func NotExtended(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNotExtended, obj...)
}

// NetworkAuthenticationRequired responds with 511 Network Authentication Required.
func NetworkAuthenticationRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNetworkAuthenticationRequired, obj...)
}
