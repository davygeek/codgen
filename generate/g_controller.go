package generate

const (
	CtrlTPL = `package controllers

import (
	"{{pkgPath}}/models"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
    "github.com/zssky/tc/http"
)

// {{ctrlName}}Controller operations for {{ctrlName}}
type {{ctrlName}}Controller struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description create {{ctrlName}}
// @Param	body		body 	models.{{ctrlName}}	true		"body for {{ctrlName}} content"
// @Success 201 {int} models.{{ctrlName}}
// @Failure 403 body is empty
func (c *{{ctrlName}}Controller) Post() {
	var v models.{{ctrlName}}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.Add{{ctrlName}}(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
            c.Data["json"] = http.HttpResponse{Code: 0, Message: "success", Data: v}
		} else {
            c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		}
	} else {
        c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
	}
	c.ServeJSON()
}

// Detail ...
// @Title Get the detail
// @Description get {{ctrlName}} by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.{{ctrlName}}
// @Failure 403 :id is empty
func (c *{{ctrlName}}Controller) Detail() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.Get{{ctrlName}}ById(id)
	if err != nil {
        c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
	} else {
        c.Data["json"] = http.HttpResponse{Code: 0, Message: "success", Data: v}
	}
	c.ServeJSON()
}

// List ...
// @Title Get all 
// @Description get {{ctrlName}}
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.{{ctrlName}}
// @Failure 403
func (c *{{ctrlName}}Controller) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sort"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

    if len(sortby) == 0 && len(order) != 0 {
        order = []string{}
    }

	l, count, err := models.GetAll{{ctrlName}}(query, fields, sortby, order, offset, limit)
	if err != nil {
        c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
	} else {
        if count == 0 {
            c.Data["json"] = http.RespData{Total: count, Rows: []string{}}
        } else {
            c.Data["json"] = http.RespData{Total: count, Rows: l}
        }
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the {{ctrlName}}
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.{{ctrlName}}	true		"body for {{ctrlName}} content"
// @Success 200 {object} models.{{ctrlName}}
// @Failure 403 :id is not int
func (c *{{ctrlName}}Controller) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.{{ctrlName}}{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.Update{{ctrlName}}ById(&v); err == nil {
            c.Data["json"] = http.HttpResponse{Code: 0, Message: "success"}
		} else {
            c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
		}
	} else {
        c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the {{ctrlName}}
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
func (c *{{ctrlName}}Controller) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.Delete{{ctrlName}}(id); err == nil {
        c.Data["json"] = http.HttpResponse{Code: 0, Message: "success"}
	} else {
        c.Data["json"] = http.HttpResponse{Code: 1, Message: err.Error()}
	}
	c.ServeJSON()
}
`
)
