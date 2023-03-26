package exceptions

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/pkg/errors"
)

// ResolveException 包装 recover 的返回值
func ResolveException(v any) contracts.Exception {
	if v == nil {
		return nil
	}

	switch e := v.(type) {
	case contracts.Exception:
		return e
	case string:
		return WithError(errors.New(e))
	case contracts.Fields:
		if e["error"] != nil {
			return WithError(fmt.Errorf("%v", e["error"]))
		}
		if e["msg"] != nil {
			return WithError(fmt.Errorf("%v", e["msg"]))
		}
		if e["message"] != nil {
			return WithError(fmt.Errorf("%v", e["message"]))
		}
		if e["err"] != nil {
			return WithError(fmt.Errorf("%v", e["err"]))
		}
		return WithError(errors.New("Server exception"))
	default:
		return New("Server exception")
	}
}
