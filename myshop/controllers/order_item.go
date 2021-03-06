package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"myshop/models"
	"myshop/util"
	"strconv"
)

// OrderItemController operations for OrderItem
type OrderItemController struct {
	beego.Controller
}

// URLMapping ...
func (c *OrderItemController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create OrderItem
// @Param	body		body 	models.OrderItem	true		"body for OrderItem content"
// @Success 201 {int} models.OrderItem
// @Failure 403 body is empty
// @router / [post]
func (c *OrderItemController) Post() {
	var v models.OrderItem
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddOrderItem(&v); err == nil {
			//c.Ctx.Output.SetStatus(201)
			c.Data["json"] = util.Success(v)
		} else {
			c.Data["json"] = util.Error(err)
		}
	} else {
		c.Data["json"] = util.Error(err)
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get OrderItem by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.OrderItem
// @Failure 403 :id is empty
// @router /:id [get]
func (c *OrderItemController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetOrderItemById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the OrderItem
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.OrderItem	true		"body for OrderItem content"
// @Success 200 {object} models.OrderItem
// @Failure 403 :id is not int
// @router /:id [put]
func (c *OrderItemController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.OrderItem{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateOrderItemById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the OrderItem
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *OrderItemController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteOrderItem(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
