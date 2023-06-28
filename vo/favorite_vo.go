package vo

import (
	"fmt"
	"gin_mall/model"
)

type FavoriteVo struct {
	*ProductVo
}

func ToFavoriteVo(f *model.Favorite) *FavoriteVo {
	p := f.Product
	u := f.User

	fmt.Printf("-------------: u: %v\n", u)
	fmt.Printf("==============> p: %v\n", p)
	return &FavoriteVo{
		ProductVo: ToProductVo(p),
	}
}

func ToFavoriteListVo(f []*model.Favorite) []*FavoriteVo {
	vos := make([]*FavoriteVo, len(f))
	for i, f := range f {
		vos[i] = ToFavoriteVo(f)
	}
	return vos
}
