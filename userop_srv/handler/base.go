package handler

import (
	"go-shop-srvs/userop_srv/proto"

	"gorm.io/gorm"
)

type UserOpServer struct {
	proto.UnimplementedAddressServer
	proto.UnimplementedMessageServer
	proto.UnimplementedUserFavServer
}

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
