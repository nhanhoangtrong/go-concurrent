package interfaces

import "deals-sync/pkg/config"

type Worker interface {
	Run(config *config.Config, count int)
}
