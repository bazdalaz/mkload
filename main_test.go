package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakePrefix(t *testing.T) {
	assert := assert.New(t)
	res, _ := make_prefix("730")

	assert.Equal(res, " 73", "simple plant")

	res, _ = make_prefix("634")
	assert.Equal(res, " 44", "force write value")
}

func TestCreateString(t *testing.T) {
	assert := assert.New(t)

	res := create_string("73_R66", "73TBP66", "RUN" )
	assert.Equal(res,  "SEQCMD    73_R66 -LOAD  73TBP66 -SEMI\n")

	res = create_string("73ECOFRN", "73FRABSN", "RUN" )
	assert.Equal(res,  "SEQCMD  73ECOFRN -LOAD 73FRABSN -AUTO  -START\n")

}
