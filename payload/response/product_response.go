package response

import "git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"

type ProductResponse struct {
	ID       uint
	Name     string
	Price    uint
	Stock    uint
	Category models.Category
}
