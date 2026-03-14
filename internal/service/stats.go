package service

import (
	"context"
	"real-holat/storage/repo"
)

func (s *reportService) MainPageStats(ctx context.Context) (*repo.MainPageStats, error) {

	stats, err := s.strg.Report().MainPageStats(ctx)
	if err != nil {
		return nil, err
	}

	return stats, nil
}
