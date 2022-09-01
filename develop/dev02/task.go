package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var str string
	if _, err := fmt.Scan(&str); err != nil {
		os.Exit(2)
	}

	str, err := Unpack(str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

func Unpack(str string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	ctx := new(Context)
	*ctx = Context{
		encoder: &EmptyEncoder{ctx: ctx},
	}

	reader := strings.NewReader(str)

	for ch, _, err := reader.ReadRune(); err == nil; ch, _, err = reader.ReadRune() {
		if err = ctx.EncodeRune(ch); err != nil {
			return "", err
		}
	}

	if err := ctx.Flush(); err != nil {
		return "", err
	}

	return ctx.sb.String(), nil
}

type RuneEncoder interface {
	EncodeRune(r rune) error
	Flush() error
}

// Context is a main part of the state pattern. The encoder changes in process.
type Context struct {
	sb      strings.Builder
	encoder RuneEncoder
	prev    struct {
		r     rune
		count int
	}
}

func (c *Context) EncodeRune(r rune) error {
	return c.encoder.EncodeRune(r)
}

func (c *Context) SetEncoder(enc RuneEncoder) {
	c.encoder = enc
}

func (c *Context) Flush() error {
	return c.encoder.Flush()
}

// EmptyEncoder works if prev is empty
type EmptyEncoder struct {
	ctx *Context
}

func (e *EmptyEncoder) EncodeRune(r rune) error {
	switch {
	case unicode.IsDigit(r):
		return errors.New("unexpected digit appeared")
	case r == '\\':
		e.ctx.SetEncoder(&EscapeEncoder{e.ctx})
	default:
		// not special symbols
		e.ctx.prev.r = r
		e.ctx.SetEncoder(&CommonEncoder{ctx: e.ctx})
	}

	return nil
}

func (e *EmptyEncoder) Flush() error {
	return nil
}

// EscapeEncoder works if prev symbol is backslash
type EscapeEncoder struct {
	ctx *Context
}

func (e *EscapeEncoder) EncodeRune(r rune) error {
	switch {
	case unicode.IsDigit(r) || r == '\\':
		e.ctx.prev.r = r
		e.ctx.SetEncoder(&CommonEncoder{e.ctx})
	default:
		return errors.New("unknown escape character" + string(r))
	}

	return nil
}

func (e *EscapeEncoder) Flush() error {
	return errors.New("no escape character at the end")
}

// CommonEncoder works if prev symbol is not special
type CommonEncoder struct {
	ctx *Context
}

func (e *CommonEncoder) flushPrev() {
	if e.ctx.prev.count > 0 {
		e.ctx.sb.WriteString(strings.Repeat(string(e.ctx.prev.r), e.ctx.prev.count))
	} else {
		e.ctx.sb.WriteRune(e.ctx.prev.r)
	}

	e.ctx.prev = struct {
		r     rune
		count int
	}{}
}

func (e *CommonEncoder) EncodeRune(r rune) error {
	switch {
	case unicode.IsDigit(r):
		e.ctx.prev.count = e.ctx.prev.count*10 + int(r-'0')
	case r == '\\':
		e.flushPrev()
		e.ctx.SetEncoder(&EscapeEncoder{e.ctx})
	default:
		// not special symbols
		e.flushPrev()
		e.ctx.prev.r = r
	}

	return nil
}

func (e *CommonEncoder) Flush() error {
	e.flushPrev()
	return nil
}
