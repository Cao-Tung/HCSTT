package controllers

import (
	"hcstt/models"
	"github.com/astaxie/beego"
	"fmt"
)

// Operations about data
type HCSTTController struct {
	beego.Controller
}

// @Title axtract Data
// @Description axtract Data
// @Param	data		path 	string	true		"The data you want "
// @Success 200 {int} models.DataResponse
// @Failure 403 data is empty
// @router /:data [get]
func (h *HCSTTController) GetData() {
	data := h.Ctx.Input.Param(":data")
	fmt.Println(data)
	dt := models.HandleData(data)
	h.Data["json"] = dt //map[string]string{"error": err}
	h.ServeJSON()
}
