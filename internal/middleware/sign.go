package middleware

//
// import (
// 	"elm/pkg/helper/md5"
// 	"elm/pkg/helper/resp"
// 	"elm/pkg/log"
// 	"github.com/gin-gonic/gin"
// 	"github.com/spf13/viper"
// 	"net/http"
// 	"sort"
// 	"strings"
// )
//
// func SignMiddleware(logger *log.Logger, conf *viper.Viper) gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		requiredHeaders := []string{"Timestamp", "Nonce", "Sign", "App-Version"}
//
// 		for _, header := range requiredHeaders {
// 			value, ok := ctx.Request.Header[header]
// 			if !ok || len(value) == 0 {
// 				resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
// 				ctx.Abort()
// 				return
// 			}
// 		}
//
// 		data := map[string]string{
// 			"AppKey":     conf.GetString("security.api_sign.app_key"),
// 			"Timestamp":  ctx.Request.Header.Get("Timestamp"),
// 			"Nonce":      ctx.Request.Header.Get("Nonce"),
// 			"AppVersion": ctx.Request.Header.Get("App-Version"),
// 		}
//
// 		var keys []string
// 		for k := range data {
// 			keys = append(keys, k)
// 		}
// 		sort.Slice(keys, func(i, j int) bool { return strings.ToLower(keys[i]) < strings.ToLower(keys[j]) })
//
// 		var str string
// 		for _, k := range keys {
// 			str += k + data[k]
// 		}
// 		str += conf.GetString("security.api_sign.app_security")
//
// 		if ctx.Request.Header.Get("Sign") != strings.ToUpper(md5.Md5(str)) {
// 			resp.HandleError(ctx, http.StatusBadRequest, 1, "sign error.", nil)
// 			ctx.Abort()
// 			return
// 		}
// 		ctx.Next()
// 	}
// }
