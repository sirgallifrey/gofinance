package features

import (
	"context"
	"requirementor/domain"
)

type FeatureService interface {
	CreateFeature(ctx context.Context, feat domain.Feature) (domain.Feature, error)
	GetFeatures(ctx context.Context, projectId string) ([]domain.Feature, error)
}

// type FeatureServiceImpl struct {
// 	repo repo.Querier
// }

// func (srv *FeatureServiceImpl) CreateFeature(ctx context.Context, feat domain.Feature) (result domain.Feature, err error) {
// 	var params repo.CreateFeatureParams
// 	err = params.From(feat)
// 	if err != nil {
// 		return result, err
// 	}

// 	model, err := srv.repo.CreateFeature(ctx, params)
// 	if err != nil {
// 		return result, err
// 	}

// 	result, err = model.ToDomain()
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

// func (srv *FeatureServiceImpl) GetFeatures(ctx context.Context, projectId string) ([]domain.Feature, error) {
// 	uuid, err := repo.UUIDFromString(projectId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	results, err := srv.repo.GetFeatures(ctx, repo.GetFeaturesParams{
// 		ProjectID: uuid,
// 		Limit:     100,
// 		Offset:    0,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}

// 	var items []domain.Feature
// 	if results == nil {
// 		return items, nil
// 	}

// 	for _, item := range results {
// 		domainItem, err := item.ToDomain()
// 		if err != nil {
// 			return nil, err
// 		}
// 		items = append(items, domainItem)
// 	}

// 	return items, nil
// }
