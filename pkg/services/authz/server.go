package authz

import (
	"context"
	"fmt"

	authzv1 "github.com/grafana/authlib/authz/proto/v1"
	"github.com/grafana/authlib/claims"

	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/infra/tracing"
	"github.com/grafana/grafana/pkg/services/accesscontrol"
	"github.com/grafana/grafana/pkg/services/featuremgmt"
	"github.com/grafana/grafana/pkg/services/grpcserver"
)

var _ authzv1.AuthzServiceServer = (*legacyServer)(nil)

type legacyServer struct {
	authzv1.UnimplementedAuthzServiceServer

	acSvc  accesscontrol.Service
	logger log.Logger
	tracer tracing.Tracer
	cfg    *Cfg
}

func newLegacyServer(
	acSvc accesscontrol.Service, features featuremgmt.FeatureToggles,
	grpcServer grpcserver.Provider, tracer tracing.Tracer, cfg *Cfg,
) (*legacyServer, error) {
	if !features.IsEnabledGlobally(featuremgmt.FlagAuthZGRPCServer) {
		return nil, nil
	}

	s := &legacyServer{
		acSvc:  acSvc,
		logger: log.New("authz-grpc-server"),
		tracer: tracer,
		cfg:    cfg,
	}

	if cfg.listen {
		grpcServer.GetServer().RegisterService(&authzv1.AuthzService_ServiceDesc, s)
	}

	return s, nil
}

// AuthFuncOverride is a function that allows to override the default auth function.
// This override is only allowed in development mode as we skip all authentication checks.
func (s *legacyServer) AuthFuncOverride(ctx context.Context, _ string) (context.Context, error) {
	ctx, span := s.tracer.Start(ctx, "authz.AuthFuncOverride")
	defer span.End()

	if !s.cfg.allowInsecure {
		s.logger.Error("AuthFuncOverride is not allowed in production mode")
		return nil, tracing.Errorf(span, "AuthFuncOverride is not allowed in production mode")
	}
	return ctx, nil
}

// AuthorizeFuncOverride is a function that allows to override the default authorize function that checks the namespace of the caller.
// We skip all authorization checks in development mode. Once we have access tokens, we need to do namespace validation in the Read handler.
func (s *legacyServer) AuthorizeFuncOverride(ctx context.Context) error {
	_, span := s.tracer.Start(ctx, "authz.AuthorizeFuncOverride")
	defer span.End()

	if !s.cfg.allowInsecure {
		s.logger.Error("AuthorizeFuncOverride is not allowed in production mode")
		return tracing.Errorf(span, "AuthorizeFuncOverride is not allowed in production mode")
	}
	return nil
}

func (s *legacyServer) Read(ctx context.Context, req *authzv1.ReadRequest) (*authzv1.ReadResponse, error) {
	ctx, span := s.tracer.Start(ctx, "authz.grpc.Read")
	defer span.End()

	// FIXME: once we have access tokens, we need to do namespace validation here

	action := req.GetAction()
	subject := req.GetSubject()
	namespace := req.GetNamespace() // TODO can we consider the stackID as the orgID?

	info, err := claims.ParseNamespace(namespace)
	if err != nil || info.OrgID == 0 {
		return nil, fmt.Errorf("invalid namespace: %s", namespace)
	}

	ctxLogger := s.logger.FromContext(ctx)
	ctxLogger.Debug("Read", "action", action, "subject", subject, "namespace", namespace)

	permissions, err := s.acSvc.SearchUserPermissions(
		ctx,
		info.OrgID,
		accesscontrol.SearchOptions{Action: action, TypedID: subject},
	)
	if err != nil {
		ctxLogger.Error("failed to search user permissions", "error", err)
		return nil, tracing.Errorf(span, "failed to search user permissions: %w", err)
	}

	data := make([]*authzv1.ReadResponse_Data, 0, len(permissions))
	for _, perm := range permissions {
		data = append(data, &authzv1.ReadResponse_Data{Scope: perm.Scope})
	}
	return &authzv1.ReadResponse{
		Data:  data,
		Found: len(data) > 0,
	}, nil
}
