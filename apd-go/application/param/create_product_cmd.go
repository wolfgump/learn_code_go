package param

import "errors"

//CreateProductCmd param
type CreateProductCmd struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

var (
	//ErrEmptyName empty name
	ErrEmptyName = errors.New("Name can not be null")
	//ErrDescription empty ErrDescription
	ErrDescription = errors.New("ErrDescription can not be null")
)

//Validate request param
func (p *CreateProductCmd) Validate() error {
	if p.Name == "" {
		return ErrEmptyName
	}
	if p.Description == "" {
		return ErrDescription
	}
	return nil

}
