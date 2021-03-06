package configserver

import (
	"github.com/concourse/atc"
	"github.com/concourse/atc/config"
	"github.com/concourse/atc/db"
	"github.com/pivotal-golang/lager"
)

type Server struct {
	logger        lager.Logger
	teamDBFactory db.TeamDBFactory
	validate      ConfigValidator
}

type ConfigValidator func(atc.Config) ([]config.Warning, []string)

func NewServer(
	logger lager.Logger,
	teamDBFactory db.TeamDBFactory,
	validator ConfigValidator,
) *Server {
	return &Server{
		logger:        logger,
		teamDBFactory: teamDBFactory,
		validate:      validator,
	}
}
