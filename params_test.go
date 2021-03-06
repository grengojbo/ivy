package ivy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseParams(t *testing.T) {
	assert := assert.New(t)

	p, err := ParseParams("")
	assert.NoError(err)
	assert.Equal(NewParams(), p)

	p, err = ParseParams("_")
	assert.NoError(err)
	assert.Equal(NewParams(), p)

	p, err = ParseParams("r_100x200,c_100x100,g_n,q_100")
	assert.NoError(err)
	assert.Equal(&Params{100, 200, 100, 100, "n", 100, true, true, true, false, ""}, p)

	p, err = ParseParams("r")
	assert.Equal("invalid parameter: r", err.Error())

	p, err = ParseParams("rr")
	assert.Equal("invalid parameter: rr", err.Error())

	p, err = ParseParams("w_100x100")
	assert.Equal("invalid parameter: w_100x100", err.Error())
}

func TestParamsGetParamDimentsion(t *testing.T) {
	assert := assert.New(t)

	width, height, err := getParamDimentsion("r", "100x200", 0)
	assert.Equal(100, width)
	assert.Equal(200, height)
	assert.NoError(err)

	width, height, err = getParamDimentsion("r", "-1x100", 0)
	assert.Equal(0, width)
	assert.Equal(0, height)
	assert.Equal("value -1x100 must be > 0: r", err.Error())

	width, height, err = getParamDimentsion("r", "100x-1", 0)
	assert.Equal(0, width)
	assert.Equal(0, height)
	assert.Equal("value 100x-1 must be > 0: r", err.Error())

	width, height, err = getParamDimentsion("r", "100x-1", 0)
	assert.Equal(0, width)
	assert.Equal(0, height)
	assert.Equal("value 100x-1 must be > 0: r", err.Error())

	width, height, err = getParamDimentsion("r", "ss", 0)
	assert.Equal(0, width)
	assert.Equal(0, height)
	assert.Equal("invalid value for r", err.Error())

	width, height, err = getParamDimentsion("r", "ax100", 0)
	assert.Equal(0, width)
	assert.Equal(0, height)
	assert.Equal("could not parse value for parameter: r", err.Error())

	width, height, err = getParamDimentsion("r", "100xa", 0)
	assert.Equal(0, width)
	assert.Equal(0, height)
	assert.Equal("could not parse value for parameter: r", err.Error())
}

func TestParamsString(t *testing.T) {
	assert := assert.New(t)

	params, err := ParseParams("r_100x200,c_100x100,g_n,q_100")
	assert.NoError(err)
	assert.Equal("100_200_100_100_100", params.String())

	params, err = ParseParams("r_100x200")
	assert.NoError(err)
	assert.Equal("100_200_0_0_-1", params.String())
}

func BenchmarkParseParams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseParams("r_100x200,c_100x100,g_n,q_100")
	}
}

func BenchmarkParamsToString(b *testing.B) {
	params, _ := ParseParams("r_100x200,c_100x100,g_n,q_100")
	for i := 0; i < b.N; i++ {
		params.String()
	}
}
