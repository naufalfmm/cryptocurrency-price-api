package log

import (
	"github.com/naufalfmm/cryptocurrency-price-api/resources/config"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger/zeroLogger"
)

func NewLogger(c *config.EnvConfig) (logger.Logger, error) {
	return zeroLogger.NewZeroLogger(
		zeroLogger.WithEnabled(c.LogMode),
	)
}
