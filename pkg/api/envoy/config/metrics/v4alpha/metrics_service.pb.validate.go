// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/metrics/v4alpha/metrics_service.proto

package envoy_config_metrics_v4alpha

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

	v4alpha "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/config/core/v4alpha"
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

	_ = v4alpha.ApiVersion(0)
)

// Validate checks the field values on MetricsServiceConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MetricsServiceConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetGrpcService() == nil {
		return MetricsServiceConfigValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetGrpcService()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetricsServiceConfigValidationError{
				field:  "GrpcService",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := v4alpha.ApiVersion_name[int32(m.GetTransportApiVersion())]; !ok {
		return MetricsServiceConfigValidationError{
			field:  "TransportApiVersion",
			reason: "value must be one of the defined enum values",
		}
	}

	if v, ok := interface{}(m.GetReportCountersAsDeltas()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetricsServiceConfigValidationError{
				field:  "ReportCountersAsDeltas",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// MetricsServiceConfigValidationError is the validation error returned by
// MetricsServiceConfig.Validate if the designated constraints aren't met.
type MetricsServiceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricsServiceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricsServiceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricsServiceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricsServiceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricsServiceConfigValidationError) ErrorName() string {
	return "MetricsServiceConfigValidationError"
}

// Error satisfies the builtin error interface
func (e MetricsServiceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetricsServiceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricsServiceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricsServiceConfigValidationError{}
