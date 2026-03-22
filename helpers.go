package fiberext

import (
	"github.com/gofiber/fiber/v2"
)

// 1xx Informational

func Continue(ctx Context) error {
	return Respond(ctx, fiber.StatusContinue, nil)
}

func SwitchingProtocols(ctx Context) error {
	return Respond(ctx, fiber.StatusSwitchingProtocols, nil)
}

func Processing(ctx Context) error {
	return Respond(ctx, fiber.StatusProcessing, nil)
}

func EarlyHints(ctx Context) error {
	return Respond(ctx, fiber.StatusEarlyHints, nil)
}

// 2xx Success

func OK(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusOK, obj)
}

func Created(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusCreated, obj)
}

func Accepted(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusAccepted, obj)
}

func NonAuthoritativeInformation(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusNonAuthoritativeInformation, obj)
}

func NoContent(ctx Context) error {
	return Respond(ctx, fiber.StatusNoContent, nil)
}

func ResetContent(ctx Context) error {
	return Respond(ctx, fiber.StatusResetContent, nil)
}

func PartialContent(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusPartialContent, obj)
}

func MultiStatus(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusMultiStatus, obj)
}

func AlreadyReported(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusAlreadyReported, obj)
}

func IMUsed(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusIMUsed, obj)
}

// 3xx Redirection

func MultipleChoices(ctx Context, obj ...any) error {
	return Respond(ctx, fiber.StatusMultipleChoices, obj)
}

func MovedPermanently(ctx Context) error {
	return Respond(ctx, fiber.StatusMovedPermanently, nil)
}

func Found(ctx Context) error {
	return Respond(ctx, fiber.StatusFound, nil)
}

func SeeOther(ctx Context) error {
	return Respond(ctx, fiber.StatusSeeOther, nil)
}

func NotModified(ctx Context) error {
	return Respond(ctx, fiber.StatusNotModified, nil)
}

func UseProxy(ctx Context) error {
	return Respond(ctx, fiber.StatusUseProxy, nil)
}

func SwitchProxy(ctx Context) error {
	return Respond(ctx, fiber.StatusSwitchProxy, nil)
}

func TemporaryRedirect(ctx Context) error {
	return Respond(ctx, fiber.StatusTemporaryRedirect, nil)
}

func PermanentRedirect(ctx Context) error {
	return Respond(ctx, fiber.StatusPermanentRedirect, nil)
}

// 4xx Client Errors

func BadRequest(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusBadRequest, obj...)
}

func Unauthorized(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUnauthorized, obj...)
}

func PaymentRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusPaymentRequired, obj...)
}

func Forbidden(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusForbidden, obj...)
}

func NotFound(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNotFound, obj...)
}

func MethodNotAllowed(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusMethodNotAllowed, obj...)
}

func NotAcceptable(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNotAcceptable, obj...)
}

func ProxyAuthRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusProxyAuthRequired, obj...)
}

func RequestTimeout(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestTimeout, obj...)
}

func Conflict(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusConflict, obj...)
}

func Gone(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusGone, obj...)
}

func LengthRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusLengthRequired, obj...)
}

func PreconditionFailed(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusPreconditionFailed, obj...)
}

func RequestEntityTooLarge(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestEntityTooLarge, obj...)
}

func RequestURITooLong(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestURITooLong, obj...)
}

func UnsupportedMediaType(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUnsupportedMediaType, obj...)
}

func RequestedRangeNotSatisfiable(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestedRangeNotSatisfiable, obj...)
}

func ExpectationFailed(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusExpectationFailed, obj...)
}

func Teapot(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusTeapot, obj...)
}

func MisdirectedRequest(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusMisdirectedRequest, obj...)
}

func UnprocessableEntity(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUnprocessableEntity, obj...)
}

func Locked(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusLocked, obj...)
}

func FailedDependency(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusFailedDependency, obj...)
}

func TooEarly(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusTooEarly, obj...)
}

func UpgradeRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUpgradeRequired, obj...)
}

func PreconditionRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusPreconditionRequired, obj...)
}

func TooManyRequests(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusTooManyRequests, obj...)
}

func RequestHeaderFieldsTooLarge(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusRequestHeaderFieldsTooLarge, obj...)
}

func UnavailableForLegalReasons(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusUnavailableForLegalReasons, obj...)
}

// 5xx Server Errors

func RespondInternalError(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusInternalServerError, obj...)
}

func NotImplemented(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNotImplemented, obj...)
}

func BadGateway(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusBadGateway, obj...)
}

func ServiceUnavailable(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusServiceUnavailable, obj...)
}

func GatewayTimeout(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusGatewayTimeout, obj...)
}

func HTTPVersionNotSupported(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusHTTPVersionNotSupported, obj...)
}

func VariantAlsoNegotiates(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusVariantAlsoNegotiates, obj...)
}

func InsufficientStorage(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusInsufficientStorage, obj...)
}

func LoopDetected(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusLoopDetected, obj...)
}

func NotExtended(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNotExtended, obj...)
}

func NetworkAuthenticationRequired(ctx Context, obj ...any) error {
	return RespondError(ctx, fiber.StatusNetworkAuthenticationRequired, obj...)
}
