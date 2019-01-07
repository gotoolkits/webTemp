package server

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"html/template"
	"io"
	"net/http"
)

type SvrStatus struct {
	Stime      string `json:"start_time"`
	SniffStart bool   `json:"Sniff_Start"`
	// Queue          UQueue   `json:"Queue"`
	SyncErrUsers []string `json:"syncErrUsers"`
}

type SysInfo struct {
	AppID   string   `json:"AppID"`
	SysName string   `json:"AppName"`
	Version string   `json:"Version"`
	APIs    []string `json:"APIs"`
	Author  string   `json:"Author"`
}

type Switch struct {
	Token  string `json:"token" xml:"token" form:"token" query:"token"`
	HostIp string `json:"HostIp" xml:"HostIp" form:"HostIp" query:"HostIp"`
	Code   string `json:"Code" xml:"Code" form:"Code" query:"Code"`
}

var (
	svrstus SvrStatus
	uuid    = "c98bad34-e0f2-4eec-bf98-2eda26af935d"
	info    SysInfo
	HostIP  string

	upgrader = websocket.Upgrader{}
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func init() {

	info = SysInfo{
		AppID:   uuid,
		SysName: "opETL - DT Micro Factory",
		Version: "V1.0.1",
		APIs:    nil,
		Author:  "gotoolkits",
	}

}

func fnInfo(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, info, " ")
}

func fnHealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Success")
}

func fnStatus(c echo.Context) error {

	return c.JSONPretty(http.StatusOK, nil, " ")
}

func fnUpload(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, nil, " OK")
}

func FnGetLocalConns(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, nil, " ")

}

func FnGetRemoteConns(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, nil, " ")
}

func FnSniffOn(c echo.Context) error {
	return c.String(http.StatusOK, "done")
}

func FnSniffOff(c echo.Context) error {
	return c.String(http.StatusOK, "done")
}
