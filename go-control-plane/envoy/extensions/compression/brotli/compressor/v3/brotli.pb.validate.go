// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/compression/brotli/compressor/v3/brotli.proto

package envoy_extensions_compression_brotli_compressor_v3

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

// Validate checks the field values on Brotli with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Brotli) Validate() error {
	if m == nil {
		return nil
	}

	if wrapper := m.GetQuality(); wrapper != nil {

		if wrapper.GetValue() > 11 {
			return BrotliValidationError{
				field:  "Quality",
				reason: "value must be less than or equal to 11",
			}
		}

	}

	if _, ok := Brotli_EncoderMode_name[int32(m.GetEncoderMode())]; !ok {
		return BrotliValidationError{
			field:  "EncoderMode",
			reason: "value must be one of the defined enum values",
		}
	}

	if wrapper := m.GetWindowBits(); wrapper != nil {

		if val := wrapper.GetValue(); val < 10 || val > 24 {
			return BrotliValidationError{
				field:  "WindowBits",
				reason: "value must be inside range [10, 24]",
			}
		}

	}

	if wrapper := m.GetInputBlockBits(); wrapper != nil {

		if val := wrapper.GetValue(); val < 16 || val > 24 {
			return BrotliValidationError{
				field:  "InputBlockBits",
				reason: "value must be inside range [16, 24]",
			}
		}

	}

	if wrapper := m.GetChunkSize(); wrapper != nil {

		if val := wrapper.GetValue(); val < 4096 || val > 65536 {
			return BrotliValidationError{
				field:  "ChunkSize",
				reason: "value must be inside range [4096, 65536]",
			}
		}

	}

	// no validation rules for DisableLiteralContextModeling

	return nil
}

// BrotliValidationError is the validation error returned by Brotli.Validate if
// the designated constraints aren't met.
type BrotliValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BrotliValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BrotliValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BrotliValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BrotliValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BrotliValidationError) ErrorName() string { return "BrotliValidationError" }

// Error satisfies the builtin error interface
func (e BrotliValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBrotli.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BrotliValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BrotliValidationError{}
