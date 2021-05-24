package validator

import (
	"github.com/faceair/jio"
)

func LoginSchema() jio.Schema {
	schema := jio.Object().Keys(jio.K{
		"loginName": jio.String().Min(4).Max(12).Regex(`^[a-zA-Z][a-zA-Z0-9_]*$`).Required(),
		"password":  jio.String().Min(6).Max(20).Required(),
	})
	return schema
}

func AddUserSchema() jio.Schema {
	schema := jio.Object().Keys(jio.K{
		"username":  jio.String().Min(1).Max(50).Required(),
		"loginName": jio.String().Min(4).Max(12).Regex(`^[a-zA-Z][a-zA-Z0-9_]*$`).Required(),
		"password":  jio.String().Min(6).Max(20).Required(),
	})
	return schema
}

func GetUserInfoSchema() jio.Schema {
	schema := jio.Object().Keys(jio.K{
		"id":        jio.Number().Integer().Min(1),
		"loginName": jio.String().Min(4).Max(12).Regex(`^[a-zA-Z][a-zA-Z0-9_]*$`),
	}).PrependTransform(OnlyOne("id", "loginName"))

	return schema
}

func GetUserListSchema() jio.Schema {
	schema := jio.Object().Keys(jio.K{
		"status":   jio.Number().Valid(-1, 0, 1).Default(-1),
		"pageNo":   jio.Number().Integer().Min(1).Default(1),
		"pageSize": jio.Number().Integer().Min(1).Max(500).Default(10),
	})
	return schema
}

func UpdateUserSchema() jio.Schema {
	schema := jio.Object().Keys(jio.K{
		"id":       jio.Number().Integer().Min(1).Required(),
		"username": jio.String().Min(1).Max(50),
		"password": jio.String().Min(6).Max(20),
	})
	return schema
}

func DeleteUsersSchema() jio.Schema {
	schema := jio.Object().Keys(jio.K{
		"idList": jio.Array().Transform(Unique()).Items(jio.Number().Integer().Min(1)).Required(),
	})
	return schema
}

func BatchUpdateUserStatus() jio.Schema {
	schema := jio.Object().Keys(jio.K{
		"idList": jio.Array().Transform(Unique()).Items(jio.Number().Integer().Min(1)).Required(),
		"status": jio.Number().Valid(0, 1).Default(1),
	})
	return schema
}
