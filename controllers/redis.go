package controllers

import (
	"encoding/json"
	"fmt"
	"redisgohttp-beego/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about Users
type RedisController struct {
	beego.Controller
}

// @Title createUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (r *RedisController) Post() {
	var result models.Result
	var redis models.Redis
	fmt.Println(r.Ctx.Input.Data())
	json.Unmarshal(r.Ctx.Input.RequestBody, &redis)
	fmt.Println(r.GetString("Cmd"))
	fmt.Println(redis)
	o, err := models.SetRedis(redis)
	if err != nil {
		result.Sucess = strconv.Itoa(models.ResultFalse)
		result.Message = err.Error()
	} else {
		result.Sucess = strconv.Itoa(models.ResultTrue)
	}
	result.Info = o
	r.Data["json"] = result
	r.ServeJSON()
}

// @Title Get
// @Description Get value by key
// @Param	key		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Redis
// @Failure 403 :key is empty
// @router /get/:key [get]
func (r *RedisController) Get() {
	var result models.Result
	var redis models.Redis
	fmt.Println(r.Ctx.Input.Data())
	key := r.Ctx.Input.Param(":key")
	fmt.Println(r.GetString("Cmd"))
	fmt.Println(redis)
	redis.Cmd = "GET"
	redis.Key = key

	o, err := models.SetRedis(redis)
	if err != nil {
		result.Sucess = strconv.Itoa(int(models.ResultFalse))
		result.Message = err.Error()
	} else {
		result.Sucess = strconv.Itoa(int(models.ResultTrue))
	}
	result.Info = o
	r.Data["json"] = result
	r.ServeJSON()
}

// @Title Get
// @Description Get value by key
// @Param	key		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Redis
// @Failure 403 :key is empty
// @router /set/:key/:value [get]
func (r *RedisController) Set() {
	var result models.Result
	var redis models.Redis
	fmt.Println(r.Ctx.Input.Data())
	json.Unmarshal(r.Ctx.Input.RequestBody, &redis)
	fmt.Println(r.GetString("Cmd"))
	fmt.Println(redis)
	key := r.Ctx.Input.Param(":key")
	value := r.Ctx.Input.Param(":value")
	fmt.Println(r.GetString("Cmd"))
	fmt.Println(redis)
	redis.Cmd = "SET"
	redis.Key = key
	redis.Value = value
	o, err := models.SetRedis(redis)
	if err != nil {
		result.Sucess = strconv.Itoa(models.ResultFalse)
		result.Message = err.Error()
	} else {
		result.Sucess = strconv.Itoa(models.ResultTrue)
	}
	result.Info = o
	r.Data["json"] = result
	r.ServeJSON()
}
