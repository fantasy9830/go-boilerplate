package database

import "go-boilerplate/internal/app/models"

// Seed Run the seeds.
func Seed() {
	// User
	adminUser := &models.User{
		Name:     "Admin",
		Username: "admin",
		Password: "admin",
		Email:    "admin@gmail.com",
	}
	demoUser := &models.User{
		Name:     "Demo",
		Username: "demo",
		Password: "demo",
		Email:    "demo@gmail.com",
	}

	// Permission
	readDashboard := &models.Permission{Name: "read_dashboard", DisplayName: "Read Dashboard", Description: "Read Dashboard"}
	createDashboard := &models.Permission{Name: "create_dashboard", DisplayName: "Create Dashboard", Description: "Create Dashboard"}
	editDashboard := &models.Permission{Name: "edit_dashboard", DisplayName: "Edit Dashboard", Description: "Edit Dashboard"}
	deleteDashboard := &models.Permission{Name: "delete_dashboard", DisplayName: "Delete Dashboard", Description: "Delete Dashboard"}

	readUsers := &models.Permission{Name: "read_users", DisplayName: "Read User", Description: "Read User"}
	createUsers := &models.Permission{Name: "create_users", DisplayName: "Create User", Description: "Create User"}
	editUsers := &models.Permission{Name: "edit_users", DisplayName: "Edit User", Description: "Edit User"}
	deleteUsers := &models.Permission{Name: "delete_users", DisplayName: "Delete User", Description: "Delete User"}

	readRoles := &models.Permission{Name: "read_roles", DisplayName: "Read Role", Description: "Read Role"}
	createRoles := &models.Permission{Name: "create_roles", DisplayName: "Create Role", Description: "Create Role"}
	editRoles := &models.Permission{Name: "edit_roles", DisplayName: "Edit Role", Description: "Edit Role"}
	deleteRoles := &models.Permission{Name: "delete_roles", DisplayName: "Delete Role", Description: "Delete Role"}

	readPermission := &models.Permission{Name: "read_permission", DisplayName: "Read Permission", Description: "Read Permission"}
	createPermission := &models.Permission{Name: "create_permission", DisplayName: "Create Permission", Description: "Create Permission"}
	editPermission := &models.Permission{Name: "edit_permission", DisplayName: "Edit Permission", Description: "Edit Permission"}
	deletePermission := &models.Permission{Name: "delete_permission", DisplayName: "Delete Permission", Description: "Delete Permission"}

	// Role
	admin := &models.Role{
		Name:        "admin",
		DisplayName: "Admin",
		Description: "Admin",
		Users: []*models.User{
			adminUser,
		},
		Permissions: []*models.Permission{
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
	general := &models.Role{
		Name:        "general",
		DisplayName: "General",
		Description: "General",
		Users: []*models.User{
			adminUser,
			demoUser,
		},
		Permissions: []*models.Permission{
			readDashboard,
			createDashboard,
			editDashboard,
			deleteDashboard,
		},
	}

	if db.HasTable(&models.User{}) {
		db.FirstOrCreate(adminUser, models.User{Name: adminUser.Name})
		db.FirstOrCreate(demoUser, models.User{Name: demoUser.Name})
	}

	if db.HasTable(&models.Permission{}) {
		db.FirstOrCreate(readUsers, models.Permission{Name: readUsers.Name})
		db.FirstOrCreate(createUsers, models.Permission{Name: createUsers.Name})
		db.FirstOrCreate(editUsers, models.Permission{Name: editUsers.Name})
		db.FirstOrCreate(deleteUsers, models.Permission{Name: deleteUsers.Name})
		db.FirstOrCreate(readRoles, models.Permission{Name: readRoles.Name})
		db.FirstOrCreate(createRoles, models.Permission{Name: createRoles.Name})
		db.FirstOrCreate(editRoles, models.Permission{Name: editRoles.Name})
		db.FirstOrCreate(deleteRoles, models.Permission{Name: deleteRoles.Name})
		db.FirstOrCreate(readPermission, models.Permission{Name: readPermission.Name})
		db.FirstOrCreate(createPermission, models.Permission{Name: createPermission.Name})
		db.FirstOrCreate(editPermission, models.Permission{Name: editPermission.Name})
		db.FirstOrCreate(deletePermission, models.Permission{Name: deletePermission.Name})

		db.FirstOrCreate(readDashboard, models.Permission{Name: readDashboard.Name})
		db.FirstOrCreate(createDashboard, models.Permission{Name: createDashboard.Name})
		db.FirstOrCreate(editDashboard, models.Permission{Name: editDashboard.Name})
		db.FirstOrCreate(deleteDashboard, models.Permission{Name: deleteDashboard.Name})
	}

	if db.HasTable(&models.Role{}) {
		db.FirstOrCreate(admin, models.Role{Name: admin.Name})
		db.FirstOrCreate(general, models.Role{Name: general.Name})
	}
}
