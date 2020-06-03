package handler

import (
	"testing"
)

func TestHandler(t *testing.T) {
	var data = `1200F230040102A0000000000000040000001048468112122012340000100000001107221800000001161204171926FABCDE123ABD06414243000termid1210Community11112341234234`
	resp := Handle(data)
	if resp != allTimeAnswer {
		t.Error("All time answer is different")
	}
}
