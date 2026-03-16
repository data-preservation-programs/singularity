package contentprovider

import (
	"context"

	"github.com/data-preservation-programs/singularity/service"
	logging "github.com/ipfs/go-log/v2"
	"gorm.io/gorm"
)

var logger = logging.Logger("contentprovider")

type Service struct {
	servers []service.Server
}

type Config struct {
	HTTP HTTPConfig
}

type HTTPConfig struct {
	EnablePiece         bool
	EnablePieceMetadata bool
	EnableIPFS          bool
	Bind                string
}

func NewService(db *gorm.DB, config Config) (*Service, error) {
	s := &Service{}

	if config.HTTP.EnablePiece || config.HTTP.EnablePieceMetadata || config.HTTP.EnableIPFS {
		s.servers = append(s.servers, &HTTPServer{
			dbNoContext:         db,
			bind:                config.HTTP.Bind,
			enablePiece:         config.HTTP.EnablePiece,
			enablePieceMetadata: config.HTTP.EnablePieceMetadata,
			enableIPFS:          config.HTTP.EnableIPFS,
		})
	}

	return s, nil
}

func (s *Service) Start(ctx context.Context) error {
	return service.StartServers(ctx, logger, s.servers...)
}
