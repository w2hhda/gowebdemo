
package controllers

import (
"github.com/astaxie/beego"
)


type MarketController struct {
	beego.Controller
}

func (c *MarketController) Get() {
	c.Data["json"] = "{\"\"}"
	c.ServeJSON()
}

