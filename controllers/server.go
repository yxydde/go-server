package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yxydde/go-server/models"
)

// Information about Server
type ServerController struct {
	beego.Controller
}

// @Title Get
// @Description Get Server Information
func (s *ServerController) Get() {
	s.Data["json"] = models.GetServer()
	s.ServeJSON()
}
