package main

import (
	"database/sql"
	"fmt"
	"github.com/ulule/deepcopier"
)

type User struct {
	Name  string
	Email sql.NullString
}

func (u *User) MethodThatTakesContext(ctx map[string]interface{}) string {
	return "hello from this method " + ctx["foo"].(string)
}

type UserResource struct {
	DisplayName            string `deepcopier:"field:Name"`
	SkipMe                 string `deepcopier:"skip"`
	MethodThatTakesContext string `deepcopier:"context"`
	Email                  string `deepcopier:"force"`
}

func main() {
	user := &User{
		Name: "liuyang",
		Email: sql.NullString{
			Valid:  true,
			String: "userlinux@sina.com",
		},
	}

	resource := &UserResource{}
	//deepcopier.Copy(user).To(resource)

	deepcopier.Copy(user).WithContext(map[string]interface{}{"foo": "bar"}).To(resource)

	fmt.Println(resource.DisplayName)
	fmt.Println(resource.Email)
	fmt.Println(resource.MethodThatTakesContext)
	fmt.Println(resource.SkipMe)

}
