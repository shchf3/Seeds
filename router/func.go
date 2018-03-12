package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"Seeds/models"
	"Seeds/utils"
	"time"
)

type FunctionRouter struct {
	routerAble
}

func ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ret":  1,
		"data": "pong",
	})
}

func detectRules(context *gin.Context) {
	db := utils.GetMySQLInstance()
	var rules []models.DetectList
	db.Database.Find(&rules)

	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": rules,
	})
}

func relayRules(context *gin.Context) {
	db := utils.GetMySQLInstance()
	node, notFound := getNodeFromQuery(context)
	if notFound {
		return
	}

	var relies []models.Relay
	db.Database.Where("source_node_id = ?", node.Id).Find(&relies)

	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": relies,
	})
}

func getBlockIp(context *gin.Context) {
	db := utils.GetMySQLInstance()
	var ips []models.BlockIp

	db.Database.Where("datetime > ?", time.Now().Unix() - 60).Find(&ips)
	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": ips,
	})
}

func getUnblockIp(context *gin.Context) {
	db := utils.GetMySQLInstance()
	var ips []models.UnblockIp

	db.Database.Where("datetime > ?", time.Now().Unix() - 60).Find(&ips)
	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": ips,
	})
}

type BlockIpInstance struct {
	Ip string `json:"ip"`
}

type BlockIpInstanceJSON struct {
	Data []BlockIpInstance `json:"data"`
}

func addBlockIp(context *gin.Context) {
	db := utils.GetMySQLInstance()
	node, notFound := getNodeFromQuery(context)
	if notFound {
		return
	}

	var body BlockIpInstanceJSON
	context.BindJSON(&body)

	for _, item := range body.Data {
		var count int
		db.Database.Where(models.BlockIp{
			Ip: item.Ip,
		}).First(&count)
		if count >= 1 {
			continue
		}
		db.Database.Save(&models.BlockIp{
			Ip: item.Ip,
			NodeId: node.Id,
			Datetime: time.Now().Unix(),
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": "ok",
	})
}

type SpeedTestLog struct {
	TelecomePing string `json:"telecomping"`
	TelecomeUpload string `json:"telecomeupload"`
	TelecomeDownload string `json:"telecomedownload"`
	UnicomPing string `json:"unicomping"`
	UnicomUpload string `json:"unicomupload"`
	UnicomDownload string `json:"unicomdownload"`
	CmccPing string `json:"cmccping"`
	CmccUpload string `json:"cmccupload"`
	CmccDownload string `json:"cmccdownload"`
}

type SpeedTestLogJSON struct {
	Data []SpeedTestLog `json:"data"`
}

func addSpeedTest(context *gin.Context) {
	db := utils.GetMySQLInstance()
	node, notFound := getNodeFromQuery(context)
	if notFound {
		return
	}

	var body SpeedTestLogJSON
	context.BindJSON(&body)

	for _, item := range body.Data {
		db.Database.Save(&models.SpeedTest{
			TelecomePing: item.TelecomePing,
			TelecomeDownload: item.TelecomeDownload,
			TelecomeUpload: item.TelecomeUpload,
			UnicomPing: item.UnicomPing,
			UnicomDownload: item.UnicomDownload,
			UnicomUpload: item.UnicomUpload,
			CmccPing: item.CmccPing,
			CmccDownload: item.CmccDownload,
			CmccUpload: item.CmccUpload,
			NodeId: node.Id,
			DateTime: time.Now().Unix(),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": "ok",
	})
}

// TODO: implement this
func getAutoExec(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": []models.Auto{},
	})
}

// TODO: implement this
func addAutoExec(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"ret": 1,
		"data": "ok",
	})
}

func (FunctionRouter) create(engine *gin.RouterGroup) {
	funcGroup := engine.Group("/func")
	{
		funcGroup.GET("/ping", ping)
		funcGroup.GET("/detect_rules", detectRules)
		funcGroup.GET("/relay_rules", relayRules)
		funcGroup.GET("/block_ip", getBlockIp)
		funcGroup.GET("/unblock_ip", getUnblockIp)
		funcGroup.GET("/autoexec", getAutoExec)
		funcGroup.POST("/block_ip", addBlockIp)
		funcGroup.POST("/speedtest", addSpeedTest)
		funcGroup.POST("/autoexec", addAutoExec)
	}
}
