package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"epc.com/api/constants"
	"epc.com/api/forms"
	"epc.com/api/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type MemberController struct{}

var memberModel = new(models.Member)

func (m MemberController) Login(c *gin.Context) {
	requestId := uuid.New().String()
	log.Printf("INFO : [RequestId : %v] Execute MemberController.login()", requestId)

	var form forms.Login
	if err := c.ShouldBind(&form); err != nil {
		var fieldname string
		var msg string
		for _, fieldError := range err.(validator.ValidationErrors) {
			fieldname = fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				msg = fmt.Sprintf("%v tidak boleh kosong.", strings.ToLower(fieldError.Field()))
			case "email":
				msg = fmt.Sprintf("%v tidak sesuai format (format:xxx@xx.xx)", strings.ToLower(fieldError.Field()))
			default:
				msg = fieldError.Error()
			}

			if msg != "" {
				break
			}
		}

		c.JSON(400, gin.H{
			"id":     requestId,
			"code":   http.StatusBadRequest,
			"status": http.StatusText(http.StatusBadRequest),
			"data": gin.H{
				"message": msg,
				"field":   strings.ToUpper(fieldname),
			},
		})
		return
	}

	member := memberModel.FindByEmail(form.Email)
	if member.Email != form.Email {
		c.JSON(400, gin.H{
			"id":     requestId,
			"code":   http.StatusBadRequest,
			"status": http.StatusText(http.StatusBadRequest),
			"data": gin.H{
				"message": "Alamat email belum terdaftar, segera daftarkan akunmu.",
				"field":   strings.ToUpper("Email"),
			},
		})
		return
	}

	if member.Password != form.Password {
		c.JSON(400, gin.H{
			"id":     requestId,
			"code":   http.StatusBadRequest,
			"status": http.StatusText(http.StatusBadRequest),
			"data": gin.H{
				"message": "Password salah, cek lagi ya.",
				"field":   strings.ToUpper("Password"),
			},
		})
		return
	}

	c.JSON(200, gin.H{
		"id":     requestId,
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data": gin.H{
			"token":       "random_token_here",
			"role":        constants.RoleStr(constants.ROLE_PM),
			"expiredDate": "2022-01-01T01:01:01.00z",
		},
	})
}

func (m MemberController) RegisterMember(c *gin.Context) {
	requestId := uuid.New().String()
	log.Printf("INFO : [RequestId : %v] Execute MemberController.registerMember()", requestId)
	c.JSON(200, gin.H{
		"id":     requestId,
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data": gin.H{
			"message": "Registrasi berhasil.",
			"data":    gin.H{},
		},
	})
}

func (m MemberController) RegisterPM(c *gin.Context) {
	requestId := uuid.New().String()
	log.Printf("INFO : [RequestId : %v] Execute MemberController.registerPM()", requestId)
	c.JSON(200, gin.H{
		"id":     requestId,
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data": gin.H{
			"message": "Registrasi berhasil.",
			"data":    gin.H{},
		},
	})
}

func (m MemberController) ForgotPassword(c *gin.Context) {
	requestId := uuid.New().String()
	log.Printf("INFO : [RequestId : %v] Execute MemberController.forgotPassword()", requestId)

	var form forms.ForgotPassword
	if err := c.ShouldBind(&form); err != nil {
		var fieldname string
		var msg string
		for _, fieldError := range err.(validator.ValidationErrors) {
			fieldname = fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				msg = fmt.Sprintf("%v tidak boleh kosong.", strings.ToLower(fieldError.Field()))
			case "email":
				msg = fmt.Sprintf("%v tidak sesuai format (format:xxx@xx.xx)", strings.ToLower(fieldError.Field()))
			default:
				msg = fieldError.Error()
			}

			if msg != "" {
				break
			}
		}

		c.JSON(400, gin.H{
			"id":     requestId,
			"code":   http.StatusBadRequest,
			"status": http.StatusText(http.StatusBadRequest),
			"data": gin.H{
				"message": msg,
				"field":   strings.ToUpper(fieldname),
			},
		})
		return
	}

	member := memberModel.FindByEmail(form.Email)
	if member.Email != form.Email {
		c.JSON(400, gin.H{
			"id":     requestId,
			"code":   http.StatusBadRequest,
			"status": http.StatusText(http.StatusBadRequest),
			"data": gin.H{
				"message": "Member tidak terdaftar.",
				"field":   strings.ToUpper("Email"),
			},
		})
	}

	// generate new password
	// send to member's email

	c.JSON(200, gin.H{
		"id":     requestId,
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data": gin.H{
			"message": fmt.Sprintf("Password baru Anda telah dikirim ke %v.", form.Email),
			"data":    gin.H{},
		},
	})
}

func (m MemberController) ChangePassword(c *gin.Context) {
	requestId := uuid.New().String()
	log.Printf("INFO : [RequestId : %v] Execute MemberController.changePassword()", requestId)

	var form forms.ChangePassword
	if err := c.ShouldBind(&form); err != nil {
		var fieldname string
		var msg string
		for _, fieldError := range err.(validator.ValidationErrors) {
			fieldname = fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				msg = fmt.Sprintf("%v tidak boleh kosong.", strings.ToLower(fieldError.Field()))
			case "email":
				msg = fmt.Sprintf("%v tidak sesuai format (format:xxx@xx.xx)", strings.ToLower(fieldError.Field()))
			default:
				msg = fieldError.Error()
			}

			if msg != "" {
				break
			}
		}

		c.JSON(400, gin.H{
			"id":     requestId,
			"code":   http.StatusBadRequest,
			"status": http.StatusText(http.StatusBadRequest),
			"data": gin.H{
				"message": msg,
				"field":   strings.ToUpper(fieldname),
			},
		})
		return
	}

	// change new password

	c.JSON(200, gin.H{
		"id":     requestId,
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data": gin.H{
			"message": "Anda berhasil mengganti password.",
			"data":    gin.H{},
		},
	})
}

func (m MemberController) Profile(c *gin.Context) {
	requestId := uuid.New().String()
	log.Printf("INFO : [RequestId : %v] Execute MemberController.profile()", requestId)
	c.JSON(200, gin.H{
		"id":     requestId,
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data": gin.H{
			"id":              1,
			"role":            "PM",
			"email":           "rizki.jayusman@metranet.co.id",
			"fullName":        "Rizki Jayusman",
			"leadName":        "Setiyo Jati",
			"overallProgress": 100,
		},
	})
}
