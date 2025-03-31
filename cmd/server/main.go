package main

import (
	"context"
	"crypto/rand"
	"github.com/aridae/goph-keeper/internal/logger"
	userusecases "github.com/aridae/goph-keeper/internal/server/controllers/user"
	"github.com/aridae/goph-keeper/internal/server/database"
	"github.com/aridae/goph-keeper/internal/server/pkg/jwt"
	"github.com/aridae/goph-keeper/internal/server/pkg/postgres"
	userrepo "github.com/aridae/goph-keeper/internal/server/repos/user"
	grpcserver "github.com/aridae/goph-keeper/internal/server/transport/grpc"
	secretsapi "github.com/aridae/goph-keeper/internal/server/transport/grpc/secrets-api"
	usersapi "github.com/aridae/goph-keeper/internal/server/transport/grpc/users-api"
	trmpgx "github.com/avito-tech/go-transaction-manager/drivers/pgxv5/v2"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go watchTerminationSignals(cancel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGHUP)

	// config-needed: db dsn, initial reconnect fallback, max retries count
	pgClient := mustInitPostgresClient(ctx, "", 5)

	err := database.PrepareSchema(ctx, pgClient)
	if err != nil {
		logger.Fatalf("failed to prepare database schema: %v", err)
	}

	userRepository := userrepo.NewRepository(pgClient, trmpgx.DefaultCtxGetter)

	// config-needed: jwt key
	jwtService := mustInitJWTService(ctx, "")

	userUseCasesController := userusecases.NewController(userRepository, jwtService)

	usersAPI := usersapi.New(userUseCasesController)
	secretsAPI := secretsapi.New(nil)

	// config-needed: grpc port
	grpcServer := grpcserver.NewServer(7012, usersAPI, secretsAPI)

	if err := grpcServer.Run(ctx); err != nil {
		logger.Fatalf("failed to start grpc server: %v", err)
	}
}

func mustInitPostgresClient(ctx context.Context, dsn string) *postgres.Client {
	client, err := postgres.NewClient(ctx, dsn,
		postgres.WithInitialReconnectBackoffOnFail(time.Second),
	)
	if err != nil {
		logger.Fatalf("failed to init postgres client: %v", err)
	}

	return client
}

func mustInitJWTService(_ context.Context, jwtKey string) *jwt.Service {
	if jwtKey == "" {
		randomFixedLenKey := make([]byte, 64)

		_, err := rand.Read(randomFixedLenKey)
		if err != nil {
			logger.Fatalf("failed to generate JWT key: %v", err)
		}

		jwtKey = string(randomFixedLenKey)
	}

	keyProvider := func(ctx context.Context) []byte {
		return []byte(jwtKey)
	}

	return jwt.NewService(keyProvider)
}

func watchTerminationSignals(cancel func(), signals ...os.Signal) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, signals...)

	<-signalCh

	logger.Infof("Got signal, shutting down...")

	// If you fail to cancel the context, the goroutine that WithCancel or WithTimeout created
	// will be retained in memory indefinitely (until the program shuts down), causing a memory leak.
	cancel()
}
