package jwt

import (
	"context"
	"fmt"
)

func Example_service_GenerateToken() {
	secretKey := []byte("super-secret-key")

	service := NewService(func(ctx context.Context) []byte {
		return secretKey
	})

	clms := Claims{
		Subject: "test-user",
	}

	ctx := context.Background()

	tokenStr, err := service.GenerateToken(ctx, clms)
	if err != nil {
		fmt.Println("Failed to generate token:", err)
		return
	}

	fmt.Println("Generated token:", tokenStr)

	// Output:
	// Generated token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ0ZXN0LXVzZXIifQ.rn7-h434Nbx2EllSssPhumoAv2JxXSrZyVbBizAmolw
}

func Example_service_ParseToken() {
	secretKey := []byte("super-secret-key")

	service := NewService(func(ctx context.Context) []byte {
		return secretKey
	})

	clms := Claims{
		Subject: "test-user",
	}

	ctx := context.Background()

	tokenStr, err := service.GenerateToken(ctx, clms)
	if err != nil {
		fmt.Println("Failed to generate token:", err)
		return
	}

	parsedClms, err := service.ParseToken(ctx, tokenStr)
	if err != nil {
		fmt.Println("Failed to parse token:", err)
		return
	}

	fmt.Println("Parsed claims:", parsedClms)

	// Output:
	// Parsed claims: {test-user}
}
