package dto

// type AddressDto struct {
// 	Aid uint `json:"aid" binding:"required"`
// }

type AddressUpdateDto struct {
	Aid     uint   `json:"aid" binding:"required"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type AddressCreateDto struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Address string `json:"address" binding:"required"`
}
