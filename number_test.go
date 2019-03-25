package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordDepth(t *testing.T) {
	start := int64(277777788888899)
	n := FromInt64(start)
	assert.EqualValues(t, 11, n.MP())
}

func TestConstructor(t *testing.T) {
	testCases := []struct {
		start    int64
		expected Number
	}{{
		start: 277777788888899,
		expected: Number{
			Twos:   1,
			Sevens: 6,
			Eights: 6,
			Nines:  2,
		},
	}, {
		start: 1000,
		expected: Number{
			Zeros: true,
			Ones:  1,
		},
	}, {
		start: 0,
		expected: Number{
			Zeros: true,
		},
	}}

	for _, tc := range testCases {
		assert.Equal(t, tc.expected, FromInt64(tc.start), `mismatch on %d`, tc.start)
	}
}

func TestNext(t *testing.T) {
	assert.Equal(t, FromInt64(0), FromInt64(100).Next())
	assert.Equal(t, FromInt64(12), FromInt64(34).Next())
	assert.Equal(t, FromInt64(125), FromInt64(555).Next())
	assert.Equal(t, FromInt64(4996238671872), FromInt64(277777788888899).Next())
}

func TestMultiplicationPersistence(t *testing.T) {
	assert.EqualValues(t, 1, FromInt64(100).MP()) // 100 -> 0
	assert.EqualValues(t, 2, FromInt64(34).MP())  // 34 -> 12 -> 2
	assert.EqualValues(t, 3, FromInt64(555).MP()) // 555 -> 125 -> 10 -> 0
}

func TestPow(t *testing.T) {
	assert.EqualValues(t, 1, pow(2, 0))
	assert.EqualValues(t, 2, pow(2, 1))
	assert.EqualValues(t, 8, pow(2, 3))
	assert.EqualValues(t, 9, pow(3, 2))
}
