// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beeApiTest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)

	// beego.Router("/student/", &controllers.StudentController{}, "*:GetAll")
	// beego.Router("/student/:studentId", &controllers.StudentController{}, "*:Get")

	schoolNs := beego.NewNamespace("/ironSchool",
		beego.NSNamespace("/student",
			beego.NSRouter("/", &controllers.StudentController{}, "get:GetAll"),
			beego.NSRouter("/", &controllers.StudentController{}, "post:PostNewStudent"),
			beego.NSRouter("/:studentId", &controllers.StudentController{}, "get:Get"),
			beego.NSRouter("/:studentId", &controllers.StudentController{}, "delete:DelStudentByTID"),
			beego.NSRouter("/test", &controllers.StudentController{}, "put:PutRequestBodyTest"),
		),
		beego.NSNamespace("/class",
			beego.NSRouter("/", &controllers.ClassController{}, "get:GetAllClass"),
			beego.NSRouter("/", &controllers.ClassController{}, "post:PostNewClass"),
			beego.NSRouter("/:classId", &controllers.ClassController{}, "get:GetById"),
			beego.NSNamespace("/name",
				beego.NSRouter("/:className", &controllers.ClassController{}, "get:GetByName"),
			),
		),
	)

	beego.AddNamespace(ns)
	beego.AddNamespace(schoolNs)
}
