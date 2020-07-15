package helper

import (
	"log"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// UseAssertion : use assert
type UseAssertion struct {
	assert testing.T
}

// AssertEqual : must equal
func (t *UseAssertion) AssertEqual(expect interface{}, actual interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.Equal(expect, actual, err)
}

// AssertNotEqual : must not equal
func (t *UseAssertion) AssertNotEqual(expect interface{}, actual interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.NotEqual(expect, actual, err)
}

// AssertString : must string
func AssertString(expect interface{}, actual interface{}, err interface{}) error {
	if reflect.TypeOf(actual).String() != "string" {
		log.Panicln(err)
	}

	return nil
}

// AssertInt : must integer
func AssertInt(expect interface{}, actual interface{}, err interface{}) error {
	if reflect.TypeOf(actual).String() != "int" {
		log.Panicln(err)
	}

	return nil
}

// AssertFloat64 : must float64
func AssertFloat64(expect interface{}, actual interface{}, err interface{}) error {
	if reflect.TypeOf(actual).String() != "float64" {
		log.Panicln(err)
	}

	return nil
}

// AssertBool : must boolean
func AssertBool(expect interface{}, actual interface{}, err interface{}) error {
	if reflect.TypeOf(actual).String() != "bool" {
		log.Panicln(err)
	}

	return nil
}

// AssertEmpty : must empty
func (t *UseAssertion) AssertEmpty(expect interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.Empty(expect, err)
}

// AssertContains : must contains something
func (t *UseAssertion) AssertContains(expect interface{}, contain interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.Contains(expect, contain, err)
}

// AssertNotContains : must not contains something
func (t *UseAssertion) AssertNotContains(expect interface{}, contain interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.NotContains(expect, contain, err)
}

// AssertZero : must zero
func (t *UseAssertion) AssertZero(expect interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.Zero(expect, err)
}

// AssertNotZero : must not zero
func (t *UseAssertion) AssertNotZero(expect interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.NotZero(expect, err)
}

// AssertTrue : must true
func (t *UseAssertion) AssertTrue(expect bool, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.True(expect, err)
}

// AssertFalse : must false
func (t *UseAssertion) AssertFalse(expect bool, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.False(expect, err)
}

// AssertSame : must same
func (t *UseAssertion) AssertSame(expect interface{}, actual interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.Same(expect, actual, err)
}

// AssertNil : must nil
func (t *UseAssertion) AssertNil(expect interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.Nil(expect, err)
}

// AssertNotNil : must not nil
func (t *UseAssertion) AssertNotNil(expect interface{}, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.NotNil(expect, err)
}

// AssertError : must error
func (t *UseAssertion) AssertError(log error, err interface{}) bool {
	assert := assert.New(&t.assert)

	return assert.Error(log, err)
}

// AssertFail : must fail
func (t *UseAssertion) AssertFail(expect string) bool {
	assert := assert.New(&t.assert)

	return assert.Fail(expect)
}