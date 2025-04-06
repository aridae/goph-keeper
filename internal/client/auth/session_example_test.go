package auth

import (
	"context"
	"fmt"
)

func ExampleSession_StoreToken_getToken() {
	ctx := context.Background()
	session := NewSession()

	err := session.StoreToken(ctx, "test-token")
	if err != nil {
		panic(err)
	}

	token := session.GetToken(ctx)
	fmt.Println(*token)
	// Output: test-token
}
