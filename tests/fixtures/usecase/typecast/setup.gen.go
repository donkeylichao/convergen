// Code generated by github.com/reedom/convergen
// DO NOT EDIT.

package stringer

import (
	"github.com/reedom/convergen/tests/fixtures/data/domain"
	"github.com/reedom/convergen/tests/fixtures/data/model"
)

func DomainToModel(src *domain.Pet) (dst *model.Pet) {
	dst = &model.Pet{}
	dst.ID = uint64(src.ID)
	// no match: dst.Category
	dst.Name = src.Name
	// no match: dst.PhotoUrls
	dst.Status = string(src.Status)

	return
}

func ModelToDomain(src *model.Pet) (dst *domain.Pet) {
	dst = &domain.Pet{}
	dst.ID = uint(src.ID)
	// no match: dst.Category
	dst.Name = src.Name
	// no match: dst.PhotoUrls
	dst.Status = domain.PetStatus(src.Status)

	return
}
