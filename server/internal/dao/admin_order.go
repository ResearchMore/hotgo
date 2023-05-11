// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"hotgo/internal/dao/internal"
)

// internalAdminOrderDao is internal type for wrapping internal DAO implements.
type internalAdminOrderDao = *internal.AdminOrderDao

// adminOrderDao is the data access object for table hg_admin_order.
// You can define custom methods on it to extend its functionality as you wish.
type adminOrderDao struct {
	internalAdminOrderDao
}

var (
	// AdminOrder is globally public accessible object for table hg_admin_order operations.
	AdminOrder = adminOrderDao{
		internal.NewAdminOrderDao(),
	}
)

// Fill with you ideas below.