package param

import "errors"

//CreateProductCmd param
type UpdateProductCmd struct {
	Id          uint64    `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

var (
	//ErrId empty name
	ErrId = errors.New("id can not be null")
)

//Validate request param
func (p *UpdateProductCmd) Validate() error {
	if p.Name == "" {
		return ErrEmptyName
	}
	if p.Description == "" {
		return ErrDescription
	}
	if p.Id == 0 {
		return ErrId
	}
	return nil

}
