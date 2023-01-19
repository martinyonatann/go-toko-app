package utils

import (
	"errors"
	"strings"
	"sync"
)

type List []interface{}

// RecoverWith will catch error value passed when panic.
func RecoverWith(catch func(recv interface{})) {
	if r := recover(); r != nil && catch != nil {
		catch(r)
	}
}

// Require only once.
func Require(required ...string) (err error) {
	var requireHasCalled bool
	onceEnv := new(sync.Once)

	if len(required) < 1 {
		return nil
	} else if requireHasCalled {
		return errors.New("env: require has called")
	}

	onceEnv.Do(func() {
		r := make(map[string]struct{})
		for _, v := range required {
			v = strings.TrimSpace(v)
			if v != "" {
				r[v] = struct{}{}
			}
		}
		requireHasCalled = true
	})

	return nil
}
