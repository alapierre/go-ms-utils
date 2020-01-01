package oauth

import (
	"fmt"
	"testing"
)

func TestSomeToken(t *testing.T) {
	res, err := CheckToken("http://localhost:9001/oauth", "b6d0b3d4-937f-4ded-9b6d-2829645e0b93")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}

	fmt.Printf("%#v", res)

	//if !res.Active {
	//	t.Errorf("token is not active, http status code = %d; want 200", res.StatusCode)
	//}
}

func TestGetToken(t *testing.T) {

	res, err := GetToken("http://localhost:9001/oauth", "admin", "admin", "client", "secret")

	if err != nil {
		fmt.Println("Error")
		panic(err)
	}

	fmt.Printf("%#v\n", res)
}
