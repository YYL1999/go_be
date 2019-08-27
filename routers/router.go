// @APIVersion 1.0.0
// @Title blog API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"go_demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/register", &controllers.UserController{}, "POST:Register")
	beego.Router("/login", &controllers.UserController{})
	beego.Router("/getall", &controllers.UserController{}, "GET:GetAllUser")
	beego.Router("/addarticle", &controllers.ArticleController{}, "POST:AddArticle")
	beego.Router("/getallarticle", &controllers.ArticleController{}, "GET:GetAllArticle")
	beego.Router("/getarticle/?:id", &controllers.ArticleController{}, "GET:GetOneArticle")
	beego.Router("/updatearticle", &controllers.ArticleController{}, "PUT:UpdateArticle")
	beego.Router("/delete/:id", &controllers.ArticleController{}, "DELETE:DeleteArticle")
}
