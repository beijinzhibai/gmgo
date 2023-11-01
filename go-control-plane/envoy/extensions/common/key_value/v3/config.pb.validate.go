// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/common/key_value/v3/config.proto

package envoy_extensions_common_key_value_v3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on KeyValueStoreConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *KeyValueStoreConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetConfig() == nil {
		return KeyValueStoreConfigValidationError{
			field:  "Config",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KeyValueStoreConfigValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// KeyValueStoreConfigValidationError is the validation error returned by
// KeyValueStoreConfig.Validate if the designated constraints aren't met.
type KeyValueStoreConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyValueStoreConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyValueStoreConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyValueStoreConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyValueStoreConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyValueStoreConfigValidationError) ErrorName() string {
	return "KeyValueStoreConfigValidationError"
}

// Error satisfies the builtin error interface
func (e KeyValueStoreConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyValueStoreConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyValueStoreConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyValueStoreConfigValidationError{}
