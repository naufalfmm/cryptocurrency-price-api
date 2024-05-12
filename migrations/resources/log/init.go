package log

import (
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger"
	"github.com/naufalfmm/cryptocurrency-price-api/utils/logger/zeroLogger"
)

func NewLogger() (logger.Logger, error) {
	return zeroLogger.NewZeroLogger(
		zeroLogger.WithEnabled(true),
	)
}
