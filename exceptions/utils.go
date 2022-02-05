package exceptions

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/pkg/errors"
)

// ResolveException 包装 recover 的返回值
func ResolveException(v interface{}) contracts.Exception {
	if v == nil {
		return nil
	}

	switch e := v.(type) {
	case contracts.Exception:
		return e
	case error:
		return WithError(e, contracts.Fields{})
	case string:
		return WithError(errors.New(e), contracts.Fields{})
	case contracts.Fields:
		if e["error"] != nil {
			return WithError(fmt.Errorf("%v", e["error"]), e)
		}
		if e["msg"] != nil {
			return WithError(fmt.Errorf("%v", e["msg"]), e)
		}
		if e["message"] != nil {
			return WithError(fmt.Errorf("%v", e["message"]), e)
		}
		if e["err"] != nil {
			return WithError(fmt.Errorf("%v", e["err"]), e)
		}
		return WithError(errors.New("Server exception"), e)
	default:
		return New("Server exception", contracts.Fields{"err": v})
	}
}
