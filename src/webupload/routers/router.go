// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"webupload/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/webupload", &controllers.UploadController{}, "*:Webupload")
}
