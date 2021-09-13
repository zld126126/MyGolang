package controller

import (
	"bytes"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CaptchaHdl struct{}

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}

// @Summary 获取验证码
// @Tags 验证码
// @Description 获取验证码
// @Accept  json
// @Produce  json
// @Success 200 object CaptchaResponse
// @Router /base/captcha [get]
// curl -X GET "http://127.0.0.1:9090/base/captcha"
func (p *CaptchaHdl) GetCaptcha(c *gin.Context) {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = "/captcha/" + captchaId + ".png"
	c.JSON(http.StatusOK, captcha)
}

// @Summary 获取验证码图片
// @Tags 验证码
// @Description 获取验证码图片
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"图片地址"
// @Router /base/captcha/image [get]
// curl -X GET "http://127.0.0.1:9090/base/captcha/image/XXX.png"
func (p *CaptchaHdl) GetCaptchaImg(c *gin.Context) {
	captchaId := c.Param("captchaId")
	logrus.Println("GetCaptchaPng : " + captchaId)
	p.serveHTTP(c.Writer, c.Request)
}

// @Summary 校验验证码
// @Tags 验证码
// @Description 校验验证码
// @Accept  json
// @Produce  json
// @Success 200 {string} string	"ok"
// @Router /base/captcha/verify [post]
// curl -X POST "http://127.0.0.1:9090/base/captcha/verify/dVCqYbq7r2olKZfEtTvo/647489
func (p *CaptchaHdl) VerifyCaptcha(c *gin.Context) {
	captchaId := c.Param("captchaId")
	value := c.Param("value")
	if captchaId == "" || value == "" {
		c.String(http.StatusBadRequest, "参数错误")
	}
	if captcha.VerifyString(captchaId, value) {
		c.JSON(http.StatusOK, "验证成功")
	} else {
		c.JSON(http.StatusOK, "验证失败")
	}
}

func (p *CaptchaHdl) serveHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	log := logrus.WithField("file", file).WithField("ext", ext).WithField("id", id)
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	log.Println("reload : " + r.FormValue("reload"))
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if p.serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func (p *CaptchaHdl) serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
