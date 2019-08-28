package main

import (
	"code/api"
	"code/db"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"gopkg.in/go-playground/validator.v9"
)

// func validation(d *api.Data) (err error) {

// 	if d.Name == "" {
// 		err = errors.New("name is required field")
// 	} else if d.Email == "" {
// 		err = errors.New("invalid email")
// 	} else if len(d.Password) < 8 {
// 		err = errors.New("password should be greater than 8")
// 	}

// 	return err
// }

func update(d *api.Data) (err error) {
	var r pg.Result
	conn := db.ConnectDB()
	r, err = conn.Model(d).
		Set("user_name = ?", d.Name).
		Set("password = ?", d.Password).
		Set("email = ?", d.Email).
		Where("user_id=?", d.ID).Update()
	rows := r.RowsAffected()
	//fmt.Println(rows)
	if rows == 0 {
		err := errors.New("value doesnt exist in relation")
		return err
	}
	if err != nil {
		return err
	}
	return
}
func put(c *gin.Context) {
	var d api.Data

	if body, err := c.GetRawData(); err == nil {
		if err = json.Unmarshal(body, &d); err == nil {
			//if err = validation(&d); err == nil {
			if err = validate.Struct(d); err == nil {
				if err = update(&d); err == nil {
					c.JSON(200, gin.H{
						"message": "updated successfully",
					})
				} else {
					c.JSON(200, gin.H{
						"message": err.Error(),
					})
				}
			} else {
				c.JSON(200, gin.H{
					"message": err.Error(),
				})
			}
		} else {
			c.JSON(200, gin.H{
				"message": err.Error(),
			})
		}

	}
}

var validate *validator.Validate

func main() {

	validate = validator.New()
	r := gin.Default()
	r.PUT("/puting", put)
	r.Run(":5000")

}
