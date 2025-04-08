package grpcauth

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func ExampleExtractBearerTokenFromMetadata() {
	md := metadata.Pairs(authorizationHeader, "Bearer my-bearer-token")
	ctx := metadata.NewIncomingContext(context.Background(), md)

	token, err := ExtractBearerTokenFromMetadata(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Token:", token)
	// Output: Token: my-bearer-token
}

func Example_extractTokenFromMetaData_Basic() {
	md := metadata.Pairs(authorizationHeader, "Basic dXNlcjpwYXNzd29yZA==")
	ctx := metadata.NewIncomingContext(context.Background(), md)

	token, err := extractTokenFromMetaData(ctx, "Basic")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Token:", token)
	// Output: Token: dXNlcjpwYXNzd29yZA==
}

func ExamplePutBearerTokenToMetadata() {
	ctx := context.Background()
	token := "my-bearer-token"
	newCtx := PutBearerTokenToMetadata(ctx, token)
	md, ok := metadata.FromOutgoingContext(newCtx)
	if !ok {
		fmt.Println("Failed to retrieve metadata")
		return
	}
	fmt.Println(md.Get(authorizationHeader)[0])
	// Output: Bearer my-bearer-token
}

func Example_putTokenToMetadata() {
	ctx := context.Background()
	authScheme := "Basic"
	token := "my-basic-token"
	newCtx := putTokenToMetadata(ctx, authScheme, token)
	md, ok := metadata.FromOutgoingContext(newCtx)
	if !ok {
		fmt.Println("Failed to retrieve metadata")
		return
	}
	fmt.Println(md.Get(authorizationHeader)[0])
	// Output: Basic my-basic-token
}
