package models

// Migrate run the migrations.
func Migrate() {
	db.AutoMigrate(
		&User{},
		&Role{},
		&Permission{},
	)

	db.Table("user_roles").AddForeignKey("role_id", "roles(id)", "CASCADE", "RESTRICT")
	db.Table("user_permissions").AddForeignKey("permission_id", "permissions(id)", "CASCADE", "RESTRICT")
	db.Table("role_permissions").AddForeignKey("role_id", "roles(id)", "CASCADE", "RESTRICT")
	db.Table("role_permissions").AddForeignKey("permission_id", "permissions(id)", "CASCADE", "RESTRICT")
}

// Reverse reverse the migrations.
func Reverse() {
	db.Table("role_permissions").RemoveForeignKey("permission_id", "permissions(id)")
	db.Table("role_permissions").RemoveForeignKey("role_id", "roles(id)")
	db.Table("user_permissions").RemoveForeignKey("permission_id", "permissions(id)")
	db.Table("user_roles").RemoveForeignKey("role_id", "roles(id)")

	db.DropTableIfExists(
		"role_permissions",
		"user_permissions",
		"user_roles",
		&Permission{},
		&Role{},
		&User{},
	)
}
