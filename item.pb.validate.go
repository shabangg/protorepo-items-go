// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: item.proto

package item

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
)

// define the regex for a UUID once up-front
var _item_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on AddonItem with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AddonItem) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Price

	// no validation rules for Available

	// no validation rules for Description

	return nil
}

// AddonItemValidationError is the validation error returned by
// AddonItem.Validate if the designated constraints aren't met.
type AddonItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddonItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddonItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddonItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddonItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddonItemValidationError) ErrorName() string { return "AddonItemValidationError" }

// Error satisfies the builtin error interface
func (e AddonItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddonItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddonItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddonItemValidationError{}

// Validate checks the field values on Addon with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Addon) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for RestId

	for idx, item := range m.GetAddonItem() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AddonValidationError{
					field:  fmt.Sprintf("AddonItem[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for XType

	return nil
}

// AddonValidationError is the validation error returned by Addon.Validate if
// the designated constraints aren't met.
type AddonValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddonValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddonValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddonValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddonValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddonValidationError) ErrorName() string { return "AddonValidationError" }

// Error satisfies the builtin error interface
func (e AddonValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddon.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddonValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddonValidationError{}

// Validate checks the field values on Item with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Item) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Price

	// no validation rules for Ratings

	// no validation rules for RatingsCount

	// no validation rules for OrderCount

	// no validation rules for CategoryId

	// no validation rules for SubCategoryId

	// no validation rules for Priority

	// no validation rules for Available

	// no validation rules for Veg

	for idx, item := range m.GetAddons() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ItemValidationError{
					field:  fmt.Sprintf("Addons[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for RestId

	if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ItemValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ItemValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ItemValidationError is the validation error returned by Item.Validate if the
// designated constraints aren't met.
type ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemValidationError) ErrorName() string { return "ItemValidationError" }

// Error satisfies the builtin error interface
func (e ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemValidationError{}

// Validate checks the field values on Id with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Id) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// IdValidationError is the validation error returned by Id.Validate if the
// designated constraints aren't met.
type IdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdValidationError) ErrorName() string { return "IdValidationError" }

// Error satisfies the builtin error interface
func (e IdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdValidationError{}

// Validate checks the field values on RestId with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RestId) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// RestIdValidationError is the validation error returned by RestId.Validate if
// the designated constraints aren't met.
type RestIdValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RestIdValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RestIdValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RestIdValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RestIdValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RestIdValidationError) ErrorName() string { return "RestIdValidationError" }

// Error satisfies the builtin error interface
func (e RestIdValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRestId.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RestIdValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RestIdValidationError{}