package grpcMiddleware

import (
	"context"
	"github.com/charmbracelet/log"
	"google.golang.org/grpc/stats"
)

type LoggingStatsHandler struct{}

func (l *LoggingStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	log.Infof("RPC request: %s", info.FullMethodName)
	return ctx
}

func (l *LoggingStatsHandler) HandleRPC(_ context.Context, stat stats.RPCStats) {
	switch s := stat.(type) {
	case *stats.End:
		log.Infof("RPC End: error=%v", s.Error)
	}
}
func (l *LoggingStatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	log.Infof("New connection from %v", info.RemoteAddr)
	return ctx
}

func (l *LoggingStatsHandler) HandleConn(_ context.Context, stat stats.ConnStats) {
	switch stat.(type) {
	case *stats.ConnEnd:
		log.Infof("Connection closed")
	}
}
