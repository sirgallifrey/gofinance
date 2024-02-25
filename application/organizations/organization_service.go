package application_organizations

import "gofinance/domain"

type OrganizationService interface {
	New(org domain.NewOrganization) (error, string)
	Get(id string) (error, domain.Organization)
}
