package errors

import (
	"errors"
	"fmt"
)

// 对 error 类型进行拓展，携带 error type 信息。方便针对不同的错误类型进行不同的操作。

// 保留的错误类型定义
const (
	// NoErr 无错误
	NoErr = 0
	// UnknownErr 未知错误
	UnknownErr = -1
	// ParamErr 参数错误
	ParamErr = 1
	// DuplicatedErr 重复记录错误
	DuplicatedErr = 2
)

// Typer 携带错误类型的 error 接口
type Typer interface {
	error
	Type() int32
}

// Error sdk 内部错误实现
type Error struct {
	typ int32  // 错误类型
	msg string // 错误信息

	cause error // 引起错误的更深层错误
}

// Error 错误
func (err *Error) Error() string {
	if err == nil {
		return ""
	}

	if err.cause == nil {
		return err.msg
	}

	return fmt.Sprintf("%s: %s", err.msg, err.cause.Error())
}

// Type 返回错误类型
func (err *Error) Type() int32 {
	return err.typ
}

// Unwrap 实现标准库的 Unwarp 方法，用于获取引起错误的原因
func (err *Error) Unwrap() error {
	if err == nil {
		return nil
	}
	return err.cause
}

// New 创建 error
func New(typ int32, msg string) error {
	return &Error{
		typ: typ,
		msg: msg,
	}
}

// Errorf 创建 error
func Errorf(typ int32, format string, a ...interface{}) error {
	return &Error{
		typ: typ,
		msg: fmt.Sprintf(format, a...),
	}
}

// Wrap 包装一个 error 并指定具体的错误类型
func Wrap(typ int32, err error, msg string) error {
	return &Error{
		typ: typ,
		msg: msg,

		cause: err,
	}
}

// Wrapf 包装一个 error 并指定错误类型
func Wrapf(typ int32, err error, format string, a ...interface{}) error {
	return &Error{
		typ: typ,
		msg: fmt.Sprintf(format, a...),
	}
}

// WithMsg 向 error 中添加额外信息，不改变原来的 error 类型
func WithMsg(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}

// Type 获取 error 的错误类型
func Type(err error) int32 {
	if err == nil {
		return NoErr
	}

	if e := Typer(nil); errors.As(err, &e) {
		return e.Type()
	}

	return UnknownErr
}
