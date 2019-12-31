package auth

import (
	"fmt"
	"testing"
)

func TestSomeToken(t *testing.T) {
	res, err := CheckToken("http://localhost:9001/oauth", "2e4228e4-ffd7-41ff-abf5-aa5d105abd79")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}

	fmt.Println(res)

	if !res.Active {
		t.Errorf("token is not active, http status code = %d; want 200", res.StatusCode)
	}
}
