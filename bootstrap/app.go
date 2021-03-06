/*
 * @Author: Jimpu
 * @Description: app
 */

package bootstrap

import (
	"bianwuji_demo/library/utils"
	"bianwuji_demo/models/page/consts"
	"bianwuji_demo/models/page/service/products"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// StartCron 初始化任务调度器
func StartCron() *CronTab {
	cronTab := InitScheduler()
	_ = cronTab.StartCron()
	return cronTab
}

// ResponseData 返回数据格式
type ResponseData struct {
	Status int         `json:"status"`  // 0 : 成功, 非零失败
	Msg    string      `json:"message"` // 描述
	Data   interface{} `json:"Data"`    //
}

// QueryData 查询数据
func QueryData(context *gin.Context) {
	// 解析参数
	p, err := CheckParam(context)
	if err != nil {
		context.String(http.StatusOK, "解析参数错误, error=%v", err)
		return
	}

	// 查询产品数据
	datas, err := p.QueryProducts()
	if err != nil {
		context.String(http.StatusOK, "%v", err)
		return
	}

	ret, _ := json.Marshal(ResponseData{
		Status: 0,
		Msg:    "success",
		Data:   datas,
	})
	context.String(http.StatusOK, "%v", string(ret))
}

// CheckParam 解析参数
func CheckParam(context *gin.Context) (ret products.Products, err error) {
	// 数据的维度, current: 当前, day: 天, month: 月
	dim := GetParam(context, "dim", "current")
	if dim != consts.DIM_MONTH && dim != consts.DIM_DAY && dim != consts.DIM_CURRENT {
		err = fmt.Errorf("param dim error")
		context.String(http.StatusOK, "%v", err)
		return
	}

	// 产品的类型,  t:温度计温度, p: 机械臂位置, pr: 酶标仪                                                                                                                                           // 查询的维度: month: 月, day: 天, current: 当前
	product := GetParam(context, "product", "t")
	if product != consts.P_TEMPERATURE && product != consts.P_POSITION && product != consts.P_PLATEREADER {
		err = fmt.Errorf("param product error")
		context.String(http.StatusOK, "%v", err)
		return
	}

	yes0 := utils.GetYesterday(0, 0, 0, 0)
	yes23 := utils.GetYesterday(23, 59, 59, 0)
	lastMonthStart, lastMonthEnd := utils.GetLastMonth()
	// 开始时间
	b := GetParam(context, "begin_time", "")
	if !utils.NumberValid(b) {
		err = fmt.Errorf("param begin_time is not number error")
		context.String(http.StatusOK, "%v", err)
		return
	}
	begin, _ := strconv.ParseInt(b, 10, 64)
	if product == consts.P_POSITION {
		if begin < yes0 || begin > yes23 {
			err = fmt.Errorf("param begin_time error, 必须是昨天[%v-%v]", yes0, yes23)
			context.String(http.StatusOK, "%v", err)
			return
		}
	} else if product == consts.P_PLATEREADER {
		if begin < lastMonthStart || begin > lastMonthEnd {
			err = fmt.Errorf("param begin_time error, 必须是上个月时间[%v-%v]", lastMonthStart, lastMonthEnd)
			context.String(http.StatusOK, "%v", err)
			return
		}
	}

	// 结束时间
	e := GetParam(context, "end_time", "")
	if !utils.NumberValid(e) {
		err = fmt.Errorf("param end_time is not number error")
		context.String(http.StatusOK, "%v", err)
		return
	}
	end, _ := strconv.ParseInt(e, 10, 64)
	if product == consts.P_POSITION {
		if end < yes0 || end > yes23 {
			err = fmt.Errorf("param end_time error, 必须是昨天时间[%v-%v]", yes0, yes23)
			context.String(http.StatusOK, "%v", err)
			return
		}
	} else if product == consts.P_PLATEREADER {
		if end < lastMonthStart || end > lastMonthEnd {
			err = fmt.Errorf("param end_time error, 必须是上个月时间[%v-%v]", lastMonthStart, lastMonthEnd)
			context.String(http.StatusOK, "%v", err)
			return
		}
	}
	if begin > end {
		err = fmt.Errorf("param begin > end error")
		context.String(http.StatusOK, "%v", err)
		return
	}

	ret = products.Products{
		Product:   product,
		Dim:       dim,
		BeginTime: begin,
		EndTime:   end,
	}
	return
}

// 获取接口普通参数。支持get/post
func GetParam(context *gin.Context, field string, defaultVal string) string {
	if context == nil || field == "" {
		return ""
	}
	switch context.Request.Method {
	case "", "GET":
		return context.DefaultQuery(field, defaultVal)
	case "POST":
		return context.DefaultPostForm(field, defaultVal)
	}
	return ""
}
