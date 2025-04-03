package morsecode

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStringIntoMorseCode(t *testing.T) {
	morse, err := ReadStringIntoMorse("Hello World")
	require.NoError(t, err)
	require.NotEmpty(t, morse)
	assert.Equal(t, morse, ".... . .-.. .-.. --- / .-- --- .-. .-.. -..")

	morse, err = ReadStringIntoMorse("[]{}()123.")
	require.NoError(t, err)
	require.NotEmpty(t, morse)
	assert.Equal(t, morse, ".--..-- --..--. ..--..- --..-.. -.....- .-----. .---- ..--- ...-- .-.-.-")

	morse, err = ReadStringIntoMorse("&^")
	require.Error(t, err)
	require.Empty(t, morse)
}

func TestMorseCodeIntoString(t *testing.T) {
	text, err := ReadMorseIntoString(".... . .-.. .-.. --- / .-- --- .-. .-.. -..")
	require.NoError(t, err)
	require.NotEmpty(t, text)
	assert.Equal(t, text, "HELLO WORLD")

	text, err = ReadMorseIntoString(".--..-- --..--. ..--..- --..-.. -.....- .-----. .---- ..--- ...-- .-.-.-")
	require.NoError(t, err)
	require.NotEmpty(t, text)
	assert.Equal(t, text, "[]{}()123.")

	text, err = ReadMorseIntoString("Hello")
	require.Error(t, err)
	require.Empty(t, text)
}
