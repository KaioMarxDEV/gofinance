package database

import "gorm.io/gorm"

// instance of Database to be used on services/handlers/controllers
var DB *gorm.DB
