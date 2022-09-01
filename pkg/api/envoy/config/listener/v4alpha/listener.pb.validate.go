// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/listener/v4alpha/listener.proto

package envoy_config_listener_v4alpha

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

	"github.com/golang/protobuf/ptypes"

	v4alpha "github.com/datawire/ambassador/v2/pkg/api/envoy/config/core/v4alpha"
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
	_ = ptypes.DynamicAny{}

	_ = v4alpha.TrafficDirection(0)
)

// Validate checks the field values on ListenerCollection with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListenerCollection) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetEntries() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerCollectionValidationError{
					field:  fmt.Sprintf("Entries[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListenerCollectionValidationError is the validation error returned by
// ListenerCollection.Validate if the designated constraints aren't met.
type ListenerCollectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerCollectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerCollectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerCollectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerCollectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerCollectionValidationError) ErrorName() string {
	return "ListenerCollectionValidationError"
}

// Error satisfies the builtin error interface
func (e ListenerCollectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListenerCollection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerCollectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerCollectionValidationError{}

// Validate checks the field values on Listener with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Listener) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if m.GetAddress() == nil {
		return ListenerValidationError{
			field:  "Address",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetFilterChains() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  fmt.Sprintf("FilterChains[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetUseOriginalDst()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "UseOriginalDst",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetDefaultFilterChain()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "DefaultFilterChain",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetPerConnectionBufferLimitBytes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "PerConnectionBufferLimitBytes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetHiddenEnvoyDeprecatedDeprecatedV1()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "HiddenEnvoyDeprecatedDeprecatedV1",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DrainType

	for idx, item := range m.GetListenerFilters() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  fmt.Sprintf("ListenerFilters[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetListenerFiltersTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "ListenerFiltersTimeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ContinueOnListenerFiltersTimeout

	if v, ok := interface{}(m.GetTransparent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "Transparent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetFreebind()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "Freebind",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetSocketOptions() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  fmt.Sprintf("SocketOptions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetTcpFastOpenQueueLength()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "TcpFastOpenQueueLength",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for TrafficDirection

	if v, ok := interface{}(m.GetUdpListenerConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "UdpListenerConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetApiListener()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "ApiListener",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetConnectionBalanceConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "ConnectionBalanceConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ReusePort

	for idx, item := range m.GetAccessLog() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListenerValidationError{
					field:  fmt.Sprintf("AccessLog[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if v, ok := interface{}(m.GetUdpWriterConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "UdpWriterConfig",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetTcpBacklogSize()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "TcpBacklogSize",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetBindToPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListenerValidationError{
				field:  "BindToPort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ListenerValidationError is the validation error returned by
// Listener.Validate if the designated constraints aren't met.
type ListenerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListenerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListenerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListenerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListenerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListenerValidationError) ErrorName() string { return "ListenerValidationError" }

// Error satisfies the builtin error interface
func (e ListenerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListenerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListenerValidationError{}

// Validate checks the field values on Listener_DeprecatedV1 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Listener_DeprecatedV1) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetBindToPort()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Listener_DeprecatedV1ValidationError{
				field:  "BindToPort",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// Listener_DeprecatedV1ValidationError is the validation error returned by
// Listener_DeprecatedV1.Validate if the designated constraints aren't met.
type Listener_DeprecatedV1ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Listener_DeprecatedV1ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Listener_DeprecatedV1ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Listener_DeprecatedV1ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Listener_DeprecatedV1ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Listener_DeprecatedV1ValidationError) ErrorName() string {
	return "Listener_DeprecatedV1ValidationError"
}

// Error satisfies the builtin error interface
func (e Listener_DeprecatedV1ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener_DeprecatedV1.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Listener_DeprecatedV1ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Listener_DeprecatedV1ValidationError{}

// Validate checks the field values on Listener_ConnectionBalanceConfig with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *Listener_ConnectionBalanceConfig) Validate() error {
	if m == nil {
		return nil
	}

	switch m.BalanceType.(type) {

	case *Listener_ConnectionBalanceConfig_ExactBalance_:

		if v, ok := interface{}(m.GetExactBalance()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return Listener_ConnectionBalanceConfigValidationError{
					field:  "ExactBalance",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return Listener_ConnectionBalanceConfigValidationError{
			field:  "BalanceType",
			reason: "value is required",
		}

	}

	return nil
}

// Listener_ConnectionBalanceConfigValidationError is the validation error
// returned by Listener_ConnectionBalanceConfig.Validate if the designated
// constraints aren't met.
type Listener_ConnectionBalanceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Listener_ConnectionBalanceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Listener_ConnectionBalanceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Listener_ConnectionBalanceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Listener_ConnectionBalanceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Listener_ConnectionBalanceConfigValidationError) ErrorName() string {
	return "Listener_ConnectionBalanceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e Listener_ConnectionBalanceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener_ConnectionBalanceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Listener_ConnectionBalanceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Listener_ConnectionBalanceConfigValidationError{}

// Validate checks the field values on
// Listener_ConnectionBalanceConfig_ExactBalance with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Listener_ConnectionBalanceConfig_ExactBalance) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// Listener_ConnectionBalanceConfig_ExactBalanceValidationError is the
// validation error returned by
// Listener_ConnectionBalanceConfig_ExactBalance.Validate if the designated
// constraints aren't met.
type Listener_ConnectionBalanceConfig_ExactBalanceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Listener_ConnectionBalanceConfig_ExactBalanceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Listener_ConnectionBalanceConfig_ExactBalanceValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e Listener_ConnectionBalanceConfig_ExactBalanceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Listener_ConnectionBalanceConfig_ExactBalanceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Listener_ConnectionBalanceConfig_ExactBalanceValidationError) ErrorName() string {
	return "Listener_ConnectionBalanceConfig_ExactBalanceValidationError"
}

// Error satisfies the builtin error interface
func (e Listener_ConnectionBalanceConfig_ExactBalanceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListener_ConnectionBalanceConfig_ExactBalance.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Listener_ConnectionBalanceConfig_ExactBalanceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Listener_ConnectionBalanceConfig_ExactBalanceValidationError{}