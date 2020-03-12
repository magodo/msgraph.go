// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// RoleAssignment The Role Assignment resource. Role assignments tie together a role definition with members and scopes. There can be one or more role assignments per role. This applies to custom and built-in roles.
type RoleAssignment struct {
	// Entity is the base model of RoleAssignment
	Entity
	// DisplayName The display or friendly name of the role Assignment.
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the Role Assignment.
	Description *string `json:"description,omitempty"`
	// ResourceScopes List of ids of role scope member security groups.  These are IDs from Azure Active Directory.
	ResourceScopes []string `json:"resourceScopes,omitempty"`
	// RoleDefinition undocumented
	RoleDefinition *RoleDefinition `json:"roleDefinition,omitempty"`
}

// RoleDefinition The Role Definition resource. The role definition is the foundation of role based access in Intune. The role combines an Intune resource such as a Mobile App and associated role permissions such as Create or Read for the resource. There are two types of roles, built-in and custom. Built-in roles cannot be modified. Both built-in roles and custom roles must have assignments to be enforced. Create custom roles if you want to define a role that allows any of the available resources and role permissions to be combined into a single role.
type RoleDefinition struct {
	// Entity is the base model of RoleDefinition
	Entity
	// DisplayName Display Name of the Role definition.
	DisplayName *string `json:"displayName,omitempty"`
	// Description Description of the Role definition.
	Description *string `json:"description,omitempty"`
	// RolePermissions List of Role Permissions this role is allowed to perform. These must match the actionName that is defined as part of the rolePermission.
	RolePermissions []RolePermission `json:"rolePermissions,omitempty"`
	// IsBuiltIn Type of Role. Set to True if it is built-in, or set to False if it is a custom role definition.
	IsBuiltIn *bool `json:"isBuiltIn,omitempty"`
	// RoleAssignments undocumented
	RoleAssignments []RoleAssignment `json:"roleAssignments,omitempty"`
}

// RolePermission undocumented
type RolePermission struct {
	// Object is the base model of RolePermission
	Object
	// ResourceActions Actions
	ResourceActions []ResourceAction `json:"resourceActions,omitempty"`
}
