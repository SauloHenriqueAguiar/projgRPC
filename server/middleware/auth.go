package middleware

import (
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthMiddleware é um middleware para autenticação de requisições gRPC.
func AuthMiddleware(authToken string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// Verifica se o token de autenticação está presente nas metadatas.
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			if tokens := md["authorization"]; len(tokens) > 0 {
				// Verifica se o token recebido é o mesmo que o esperado.
				if strings.TrimSpace(tokens[0]) == authToken {
					return handler(ctx, req) // Chama o handler se o token for válido.
				}
			}
		}

		// Se o token não for válido, retorna um erro de permissão negada.
		return nil, fmt.Errorf("permission denied")
	}
}

// AuthenticatedUnaryInterceptor cria um interceptor que verifica o token de autenticação.
func AuthenticatedUnaryInterceptor() grpc.UnaryServerInterceptor {
	const validToken = "my-secret-token" // Altere para seu token de autenticação real.
	return AuthMiddleware(validToken)
}
