package features

import (
	"context"
	"requirementor/domain"
)

type FeatureRepo interface {
	CreateFeature(ctx context.Context, feat domain.Feature) (domain.Feature, error)
	GetFeatures(ctx context.Context, projectId string) ([]domain.Feature, error)
}

// type FeatureRepoImpl struct {
// 	db bun.IDB
// }

// func (srv *FeatureRepoImpl) CreateFeature(ctx context.Context, feat domain.Feature) (domain.Feature, error) {

// }

// func (srv *FeatureRepoImpl) GetFeature(ctx context.Context, projectId string) (domain.Feature, error) {

// }
