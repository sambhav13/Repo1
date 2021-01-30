package mypkg1

import "testing"
import "github.com/stretchr/testify/assert"

func TestPkg1GetString(t *testing.T) {

	str := GetPkg1String()
	expectedStr := "inside the package 1 code"
	assert.Equal(t, expectedStr, str, "The two words should be the same.")

}
