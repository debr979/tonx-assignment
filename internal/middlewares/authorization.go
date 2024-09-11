package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"tonx-assignment/internal/app/controllers"
	"tonx-assignment/internal/app/models"
	"tonx-assignment/internal/app/repositories"
	"tonx-assignment/pkg/utils"
)

func UserAuthorization(c *gin.Context) {
	runMode := os.Getenv(`RUN_MODE`)
	if runMode == `debug` {
		c.Set(`username`, `doug123`)
		c.Set(`isValid`, true)
		c.Set(`user_id`, 99999999999999)
		return
	}

	funcNameSplit := strings.Split(c.Request.URL.Path, "/")
	if funcNameSplit[len(funcNameSplit)-1] == `login` {
		return
	}

	accessToken := c.Request.Header.Get(`Authorization`)
	if accessToken == "" {
		controllers.Base.Request(c).Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	accessToken = strings.Split(accessToken, "Bearer ")[1]

	isValid, username, err := utils.JsonWebToken.VerifyJWToken(accessToken, `access_token`)
	if !isValid || username == "" || err != nil {
		controllers.Base.Request(c).Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}
	var user models.User
	err = repositories.UserRepository.IsUser(&user, username)
	if err != nil || user.Id <= 0 {
		controllers.Base.Request(c).Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	c.Set(`username`, user.Username)
	c.Set(`user_id`, user.Id)
	c.Set(`isValid`, isValid)

}

func MgrAuthorization(c *gin.Context) {
	runMode := os.Getenv(`RUN_MODE`)
	if runMode == `debug` {
		c.Set(`manager_name`, `doug123`)
		c.Set(`isValid`, true)
		c.Set(`manager_id`, 99999999999999)
		return
	}

	funcNameSplit := strings.Split(c.Request.URL.Path, "/")
	if funcNameSplit[len(funcNameSplit)-1] == `login` {
		return
	}

	accessToken := c.Request.Header.Get(`Authorization`)
	if accessToken == "" {
		controllers.Base.Request(c).Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	accessToken = strings.Split(accessToken, "Bearer ")[1]

	isValid, mgrName, err := utils.JsonWebToken.VerifyJWToken(accessToken, `access_token`)
	if !isValid || mgrName == "" || err != nil {
		controllers.Base.Request(c).Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}
	var mgr models.Manager
	err = repositories.MgrRepository.IsManager(&mgr, mgrName)
	if err != nil || mgr.Id <= 0 {
		controllers.Base.Request(c).Response(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
		return
	}

	c.Set(`manager_name`, mgrName)
	c.Set(`manager_id`, mgr.Id)
	c.Set(`isValid`, isValid)

}
