package controller

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"github.com/gpmgo/gopm/modules/log"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"strings"
	"time"
	"x-archives/pipe/model"
	"x-archives/pipe/util"
)

var states = map[string]string{}

func redirectGithubLoginAction(c *gin.Context)  {
	result := util.NewResult()
	_, _, errs := gorequest.New().TLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		Get(util.HacPaiURL+"/oauth/pipe/client2").
		Set("user-agent", model.UserAgent).Timeout(10 * time.Second).EndStruct(result)
	if nil != errs{
		log.Print(2,"get oauth client id failed [code=%d ,msg=%s]",result.Code,result.Msg)
		c.Status(http.StatusNotFound)
		return
	}
	data := result.Data.(map[string]interface{})
	clientId := data["clientId"].(string)
	loginAuthURL := data["loginAuthURL"].(string)

	referer := c.Request.URL.Query().Get("referer")
	if "" == referer || !strings.Contains(referer, "://") {
		referer = model.Conf.Server + referer
	}
	if strings.HasSuffix(referer, "/") {
		referer = referer[:len(referer)-1]
	}
	state := util.RandString(16) + referer
	states[state] = state
	path := loginAuthURL + "?client_id=" + clientId + "&state=" + state + "&scope=public_repo,read:user,user:follow"

	//logger.Infof("redirect to github [" + path + "]")
	println("---redirect"+path)
	c.Redirect(http.StatusSeeOther, path)

}