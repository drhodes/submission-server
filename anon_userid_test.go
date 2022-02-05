package main

import (
	"testing"
)

func TestUserId(t *testing.T) {
	// In [2]: util.generate_jupyterhub_userid("abcdefghijklmnopqrstuvwxyz")
	// Out[2]: 'jupyter-abcdefghijklmnopqr-e13d8'

	userid := "abcdefghijklmnopqrstuvwxyz"
	result := generate_jupyterhub_userid(userid)
	expect := "jupyter-abcdefghijklmnopqr-e13d8"
	
	if result != expect {
		t.Errorf("got: %s, exp: %s", result, expect)
	}
}
