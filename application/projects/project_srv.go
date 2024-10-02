package features

import (
	"context"
	"requirementor/domain"
)

type ProjectService interface {
	CreateProject(ctx context.Context, name string) (domain.Project, error)
	GetProjects(ctx context.Context) ([]domain.Project, error)
}

// type ProjectServiceImpl struct {
// 	repo repo.Querier
// }

// func (srv *ProjectServiceImpl) CreateProject(ctx context.Context, name string) (result domain.Project, err error) {

// 	model, err := srv.repo.CreateProject(ctx, name)
// 	if err != nil {
// 		return result, err
// 	}

// 	result, err = model.ToDomain()
// 	if err != nil {
// 		return result, err
// 	}

// 	return result, nil
// }

// func (srv *ProjectServiceImpl) GetProjects(ctx context.Context) ([]domain.Project, error) {

// 	results, err := srv.repo.GetAllProjects(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var items []domain.Project
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
