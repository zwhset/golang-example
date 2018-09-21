package main

// define error inteface
type error interface {
	Error() string
}

// package errors
// 通过errors.New(string)来获取一个错误类型
func New(text string) error {
	// 初始化一个错误类型并返回
	return &errorString{text}
}

type errorString struct {
	text string
}

// 只要type 有Error()方法就实现了 error 接口
// Error() 方法返回errorString类型的text字符串
// 即调用方不需要知道啥并且直接调 type.Error()方法即可获取错误的字符串
func (e *errorString) Error() string {
	return e.text
}
