package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net"
	"os"

	_ "github.com/lib/pq"

	identityv1 "UltimateTeamX/proto/identity/v1"
	"UltimateTeamX/service/identity/internal/handler"

	"google.golang.org/grpc"
)

func main() {
	// --- Logger Initialization ---
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("Starting identity-svc...")

	// --- DB Connection ---
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	logger.Info("Connecting to DB", "host", os.Getenv("DB_HOST"), "port", os.Getenv("DB_PORT"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to DB Identity", "error", err)
		os.Exit(1)
	}
	defer func() {
		db.Close()
		logger.Info("DB connection closed")
	}()

	if err := db.Ping(); err != nil {
		logger.Error("DB ping failed", "error", err)
		os.Exit(1)
	}
	logger.Info("Connected to DB successfully")

	// --- Setup server gRPC ---
	port := ":50051"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.Error("Failed to listen", "port", port, "error", err)
		os.Exit(1)
	}
	logger.Info("gRPC listening", "port", port)

	grpcServer := grpc.NewServer()

	identityServer := &handler.IdentityServer{
		DB: db,
	}

	identityv1.RegisterIdentityServiceServer(grpcServer, identityServer)
	logger.Info("IdentityService registered with gRPC server")

	// --- Serve ---
	logger.Info("identity gRPC server running")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error("gRPC server failed", "error", err)
		os.Exit(1)
	}
}
