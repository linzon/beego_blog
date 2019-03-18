package controllers


import (
	"github.com/Echosong/beego_blog/models"
	"github.com/Echosong/beego_blog/util"
)

type OfficeController struct {
	baseController
}

/**
首页
*/
func (c *BlogController) Index()  {
	c.list()
	c.TplName= c.controllerName+"/index.html"
}
