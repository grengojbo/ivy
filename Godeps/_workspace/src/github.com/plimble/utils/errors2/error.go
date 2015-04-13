package errors2

const (
	Internal = iota
	BadReq
	NotFound
	Unauthorized
	Forbidden
)

type Error interface {
	Code() int
	Error() string
	Type() int
}

type Errors2 struct {
	httpStatus int    `json:"-"`
	err        string `json:"error,omitempty"`
	message    string `json:"message"`
	ty         int    `json:"-"`
}

func (err *Errors2) Error() string {
	return err.message
}

func (err *Errors2) HttpStatus() int {
	return err.httpStatus
}

func (err *Errors2) Type() int {
	return err.ty
}

func NewAnyError() error {
	return NewError(500, "", "any error", Internal)
}

func NewError(status int, code, msg string, ty int) *Errors2 {
	return &Errors2{status, "", msg, ty}
}

func NewInternal(msg string) *Errors2 {
	return &Errors2{500, "", msg, Internal}
}

func NewBadReq(msg string) *Errors2 {
	return &Errors2{400, "", msg, BadReq}
}

func NewNotFound(msg string) *Errors2 {
	return &Errors2{404, "", msg, NotFound}
}

func NewUnauthorized(msg string) *Errors2 {
	return &Errors2{401, "", msg, Unauthorized}
}

func NewForbidden(msg string) *Errors2 {
	return &Errors2{403, "", msg, Forbidden}
}

//With err code

func NewErrorCode(status int, code, msg string, ty int) *Errors2 {
	return &Errors2{status, code, msg, ty}
}

func NewInternalCode(code, msg string) *Errors2 {
	return &Errors2{500, code, msg, Internal}
}

func NewBadReqCode(code, msg string) *Errors2 {
	return &Errors2{400, code, msg, BadReq}
}

func NewNotFoundCode(code, msg string) *Errors2 {
	return &Errors2{404, code, msg, NotFound}
}

func NewUnauthorizedCode(code, msg string) *Errors2 {
	return &Errors2{401, code, msg, Unauthorized}
}

func NewForbiddenCode(code, msg string) *Errors2 {
	return &Errors2{403, code, msg, Forbidden}
}
