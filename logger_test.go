package prognoslog

import (
	"bytes"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestLogger_JSON(t *testing.T) {
	obj := struct {
		Hi  string
		Bye string
	}{
		"hello",
		"goodbye",
	}

	var b bytes.Buffer

	log := &Logger{out: &b}
	log.JSON("hibye", obj)

	expected := "JSON hibye={\"Hi\":\"hello\",\"Bye\":\"goodbye\"}\n"
	assert.Equal(t, expected, b.String())
}

func TestLogger_JSONIfVerbose(t *testing.T) {
	obj := struct {
		Hi  string
		Bye string
	}{
		"hello",
		"goodbye",
	}

	var b bytes.Buffer

	cases := []struct {
		setDefault       bool
		setDefaultTo     bool
		expectNullOutput bool
	}{
		{false, false, true},
		{true, false, true},
		{true, true, false},
	}

	for _, c := range cases {
		log := &Logger{out: &b}
		if c.setDefault {
			log.SetVerbosity(c.setDefaultTo)
		}
		log.JSONIfVerbose("hibye", obj)

		expected := b.String()
		if c.expectNullOutput {
			expected = ""
		}
		assert.Equal(t, expected, b.String())
	}
}

func TestLogger_JSONString(t *testing.T) {
	json := `{"hello": "world"}`

	var b bytes.Buffer

	log := &Logger{out: &b}
	log.JSONString("hibye", json)

	expected := "JSON hibye={\"hello\":\"world\"}\n"
	assert.Equal(t, expected, b.String())
}

func TestLogger_JSONString_compacted(t *testing.T) {
	json := `{
		"hello": "world",
		"hey": "galaxy"
}`

	var b bytes.Buffer

	log := &Logger{out: &b}
	log.JSONString("hibye", json)

	expected := "JSON hibye={\"hello\":\"world\",\"hey\":\"galaxy\"}\n"
	assert.Equal(t, expected, b.String())
}

func TestLogger_JSONStringIfVerbose(t *testing.T) {
	json := `{"hello":"world"}`

	var b bytes.Buffer

	cases := []struct {
		setDefault       bool
		setDefaultTo     bool
		expectNullOutput bool
	}{
		{false, false, true},
		{true, false, true},
		{true, true, false},
	}

	for _, c := range cases {
		log := &Logger{out: &b}
		if c.setDefault {
			log.SetVerbosity(c.setDefaultTo)
		}
		log.JSONStringIfVerbose("hibye", json)

		expected := "JSON hibye={\"hello\":\"world\"}\n"
		if c.expectNullOutput {
			expected = ""
		}
		assert.Equal(t, expected, b.String())
	}
}

func TestLogger_KVP(t *testing.T) {
	var b bytes.Buffer

	log := &Logger{out: &b}
	log.KVP("hi", "bye")

	expected := "KVP hi=\"bye\"\n"
	assert.Equal(t, expected, b.String())
}

func TestLogger_KVPIfVerbose(t *testing.T) {
	var b bytes.Buffer

	cases := []struct {
		setDefault       bool
		setDefaultTo     bool
		expectNullOutput bool
	}{
		{false, false, true},
		{true, false, true},
		{true, true, false},
	}

	for _, c := range cases {
		log := &Logger{out: &b}
		if c.setDefault {
			log.SetVerbosity(c.setDefaultTo)
		}
		log.KVPIfVerbose("hi", "bye")

		expected := "KVP hi=\"bye\"\n"
		if c.expectNullOutput {
			expected = ""
		}
		assert.Equal(t, expected, b.String())
	}
}

func TestLogger_Txt(t *testing.T) {
	var b bytes.Buffer

	log := &Logger{out: &b}
	log.Txt("hi bye")

	expected := "TXT hi bye\n"
	assert.Equal(t, expected, b.String())
}

func TestLogger_TxtIfVerbose(t *testing.T) {
	var b bytes.Buffer

	cases := []struct {
		setDefault       bool
		setDefaultTo     bool
		expectNullOutput bool
	}{
		{false, false, true},
		{true, false, true},
		{true, true, false},
	}

	for _, c := range cases {
		log := &Logger{out: &b}
		if c.setDefault {
			log.SetVerbosity(c.setDefaultTo)
		}
		log.TxtIfVerbose("hi bye")

		expected := "TXT hi bye\n"
		if c.expectNullOutput {
			expected = ""
		}
		assert.Equal(t, expected, b.String())
	}
}

func TestLogger_Txt_withParams(t *testing.T) {
	var b bytes.Buffer

	log := &Logger{out: &b}
	log.Txt("hi bye, %s", "later")

	expected := "TXT hi bye, later\n"
	assert.Equal(t, expected, b.String())
}

func TestLogger_TxtIfVerbose_withParams(t *testing.T) {
	var b bytes.Buffer

	cases := []struct {
		setDefault       bool
		setDefaultTo     bool
		expectNullOutput bool
	}{
		{false, false, true},
		{true, false, true},
		{true, true, false},
	}

	for _, c := range cases {
		log := &Logger{out: &b}
		if c.setDefault {
			log.SetVerbosity(c.setDefaultTo)
		}
		log.TxtIfVerbose("hi bye, %s", "later")

		expected := "TXT hi bye, later\n"
		if c.expectNullOutput {
			expected = ""
		}
		assert.Equal(t, expected, b.String())
	}
}

func TestLogger_SetVerbosity(t *testing.T) {
	var b bytes.Buffer

	log := &Logger{out: &b}
	assert.False(t, log.isVerbose)

	log.SetVerbosity()
	assert.True(t, log.isVerbose)

	log.SetVerbosity(false)
	assert.False(t, log.isVerbose)

	log.SetVerbosity(true)
	assert.True(t, log.isVerbose)
}

func Test_enforce_noPanic(t *testing.T) {
	assert.NotPanics(t, func() { enforce(0, nil) })
}

func Test_enforce_panic(t *testing.T) {
	assert.Panics(t, func() { enforce(0, errors.New("something bad")) })
}
