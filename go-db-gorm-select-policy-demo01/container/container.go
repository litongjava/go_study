package container

import "gorm.io/gorm"

var Db *gorm.DB
var PoolToDSNMap = make(map[gorm.ConnPool]string)
