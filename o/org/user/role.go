package user

type Roles string

const Admin = Roles("role_admin")
const Member = Roles("role_member")

func (r *Roles) GetRole(roleName string) Roles {
	return Roles(roleName)
}

func (r *Roles) CheckRoleAccess() {}
