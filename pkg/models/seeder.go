package models

// Seed Run the seeds.
func Seed() {
	// User
	adminUser := &User{
		Name:     "Admin",
		Username: "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
	}
	demoUser := &User{
		Name:     "Demo",
		Username: "demo",
		Password: "demo",
		Email:    "demo@gmail.com",
	}

	// Permission
	readDashboard := &Permission{Name: "read_dashboard", DisplayName: "Read Dashboard", Description: "Read Dashboard"}
	createDashboard := &Permission{Name: "create_dashboard", DisplayName: "Create Dashboard", Description: "Create Dashboard"}
	editDashboard := &Permission{Name: "edit_dashboard", DisplayName: "Edit Dashboard", Description: "Edit Dashboard"}
	deleteDashboard := &Permission{Name: "delete_dashboard", DisplayName: "Delete Dashboard", Description: "Delete Dashboard"}

	readUsers := &Permission{Name: "read_users", DisplayName: "Read User", Description: "Read User"}
	createUsers := &Permission{Name: "create_users", DisplayName: "Create User", Description: "Create User"}
	editUsers := &Permission{Name: "edit_users", DisplayName: "Edit User", Description: "Edit User"}
	deleteUsers := &Permission{Name: "delete_users", DisplayName: "Delete User", Description: "Delete User"}

	readRoles := &Permission{Name: "read_roles", DisplayName: "Read Role", Description: "Read Role"}
	createRoles := &Permission{Name: "create_roles", DisplayName: "Create Role", Description: "Create Role"}
	editRoles := &Permission{Name: "edit_roles", DisplayName: "Edit Role", Description: "Edit Role"}
	deleteRoles := &Permission{Name: "delete_roles", DisplayName: "Delete Role", Description: "Delete Role"}

	readPermission := &Permission{Name: "read_permission", DisplayName: "Read Permission", Description: "Read Permission"}
	createPermission := &Permission{Name: "create_permission", DisplayName: "Create Permission", Description: "Create Permission"}
	editPermission := &Permission{Name: "edit_permission", DisplayName: "Edit Permission", Description: "Edit Permission"}
	deletePermission := &Permission{Name: "delete_permission", DisplayName: "Delete Permission", Description: "Delete Permission"}

	// Role
	admin := &Role{
		Name:        "admin",
		DisplayName: "Admin",
		Description: "Admin",
		Users: []*User{
			adminUser,
		},
		Permissions: []*Permission{
			readUsers,
			createUsers,
			editUsers,
			deleteUsers,
			readRoles,
			createRoles,
			editRoles,
			deleteRoles,
			readPermission,
			createPermission,
			editPermission,
			deletePermission,
		},
	}
	general := &Role{
		Name:        "general",
		DisplayName: "General",
		Description: "General",
		Users: []*User{
			adminUser,
			demoUser,
		},
		Permissions: []*Permission{
			readDashboard,
			createDashboard,
			editDashboard,
			deleteDashboard,
		},
	}

	if db.HasTable(&User{}) {
		db.FirstOrCreate(adminUser, User{Name: adminUser.Name})
		db.FirstOrCreate(demoUser, User{Name: demoUser.Name})
	}

	if db.HasTable(&Permission{}) {
		db.FirstOrCreate(readUsers, Permission{Name: readUsers.Name})
		db.FirstOrCreate(createUsers, Permission{Name: createUsers.Name})
		db.FirstOrCreate(editUsers, Permission{Name: editUsers.Name})
		db.FirstOrCreate(deleteUsers, Permission{Name: deleteUsers.Name})
		db.FirstOrCreate(readRoles, Permission{Name: readRoles.Name})
		db.FirstOrCreate(createRoles, Permission{Name: createRoles.Name})
		db.FirstOrCreate(editRoles, Permission{Name: editRoles.Name})
		db.FirstOrCreate(deleteRoles, Permission{Name: deleteRoles.Name})
		db.FirstOrCreate(readPermission, Permission{Name: readPermission.Name})
		db.FirstOrCreate(createPermission, Permission{Name: createPermission.Name})
		db.FirstOrCreate(editPermission, Permission{Name: editPermission.Name})
		db.FirstOrCreate(deletePermission, Permission{Name: deletePermission.Name})

		db.FirstOrCreate(readDashboard, Permission{Name: readDashboard.Name})
		db.FirstOrCreate(createDashboard, Permission{Name: createDashboard.Name})
		db.FirstOrCreate(editDashboard, Permission{Name: editDashboard.Name})
		db.FirstOrCreate(deleteDashboard, Permission{Name: deleteDashboard.Name})
	}

	if db.HasTable(&Role{}) {
		db.FirstOrCreate(admin, Role{Name: admin.Name})
		db.FirstOrCreate(general, Role{Name: general.Name})
	}
}
