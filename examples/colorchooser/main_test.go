package main_test

import (
	"os/exec"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wesovilabs/koazee"
)

func TestColorChooser(t *testing.T) {
	out, err := exec.Command("go", "run", ".").CombinedOutput()
	if assert.NoError(t, err) {
		actual := normaliseCmdOutput(out)

		exp := regexp.QuoteMeta("\x1b[") + "(.*)m(.*)" + regexp.QuoteMeta("\x1b[0m")
		assert.Regexp(t, exp, actual[0])
		assert.Regexp(t, exp, actual[1])
		assert.Regexp(t, exp, actual[2])
		assert.Regexp(t, exp, actual[3])
		assert.Regexp(t, exp, actual[4])
		assert.Regexp(t, exp, actual[5])

		colorCode1 := regexp.MustCompile(exp).FindStringSubmatch(actual[0])[1]
		colorCode2 := regexp.MustCompile(exp).FindStringSubmatch(actual[1])[1]
		colorCode3 := regexp.MustCompile(exp).FindStringSubmatch(actual[2])[1]
		colorCode4 := regexp.MustCompile(exp).FindStringSubmatch(actual[3])[1]
		colorCode5 := regexp.MustCompile(exp).FindStringSubmatch(actual[4])[1]
		colorCode6 := regexp.MustCompile(exp).FindStringSubmatch(actual[5])[1]

		assert.Equal(t, colorCode1, colorCode2)
		assert.NotEqual(t, colorCode1, colorCode3)

		customColors := koazee.StreamOf([]string{colorCode4, colorCode5, colorCode6}).RemoveDuplicates()
		c, err := customColors.Count()
		assert.Nil(t, err)
		assert.Equal(t, 2, c)

		r1, err := customColors.Contains("31")
		assert.Nil(t, err)
		assert.True(t, r1)

		r2, err := customColors.Contains("32")
		assert.Nil(t, err)
		assert.True(t, r2)
	}
}

func normaliseCmdOutput(in []byte) []string {
	out := string(in)
	return strings.Split(out, "\n")
}
