package validator

import (
	"bytes"
	"fmt"
	"go-skeleton/serializer"
	"go-skeleton/utils/log"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/faceair/jio"
	"github.com/gin-gonic/gin"
)

func Validate(schema jio.Schema) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := c.Request.Body
		data, err := ioutil.ReadAll(body)
		if err != nil {
			log.Errorf("failed to read request body when validating, err: %v", err)
			res := serializer.ReplyError(err)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}

		c.Request.Body.Close()

		log.Infof("request uri: %v, request message: %v", c.Request.RequestURI, string(data))

		_, err = jio.ValidateJSON(&data, schema)
		if err != nil {
			log.Errorf("params failed to pass validator, err: %v", err)
			res := serializer.ReplyError(err)
			c.JSON(http.StatusOK, res)
			c.Abort()
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		c.Next()
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// 自定义参数校验规则
////////////////////////////////////////////////////////////////////////////////////////////////////
func Unique() func(*jio.Context) {
	return func(ctx *jio.Context) {
		paramList := ctx.Value.([]interface{})
		var result []interface{}
		for i := range paramList {
			flag := true
			for j := range result {
				if paramList[i] == result[j] {
					ctx.Abort(fmt.Errorf("field `%s` duplicate elements in array", ctx.FieldPath()))
				}
			}
			if flag {
				result = append(result, paramList[i])
			}
		}
	}
}

func OnlyOne(keys ...string) func(*jio.Context) {
	return func(ctx *jio.Context) {
		ctxValue, ok := ctx.Value.(map[string]interface{})
		if !ok {
			ctx.Abort(fmt.Errorf("field `%s` value %v is not object", ctx.FieldPath(), ctx.Value))
			return
		}
		contains := make([]string, 0, 3)
		for _, key := range keys {
			_, ok := ctxValue[key]
			if ok {
				contains = append(contains, key)
			}
		}

		if len(contains) != 1 {
			ctx.Abort(fmt.Errorf("contain one and only one in [%v], contains: %v", strings.Join(keys, ","), strings.Join(contains, ",")))
			return
		}
	}
}

func Together(keys ...string) func(*jio.Context) {
	return func(ctx *jio.Context) {
		ctxValue, ok := ctx.Value.(map[string]interface{})
		if !ok {
			ctx.Abort(fmt.Errorf("field `%s` value %v is not object", ctx.FieldPath(), ctx.Value))
			return
		}

		contains := make([]string, 0, 3)
		for _, key := range keys {
			_, ok := ctxValue[key]
			if ok {
				contains = append(contains, key)
			}
		}

		if len(contains) != 0 {
			if len(contains) != len(keys) {
				ctx.Abort(fmt.Errorf("missing field in [%v]", strings.Join(keys, ",")))
			}
		}
	}
}

func Greater(refPath string) func(*jio.Context) {
	return func(ctx *jio.Context) {
		ctxValue := ctx.Value.(float64)
		refVal, ok := ctx.Ref(refPath)
		if !ok {
			ctx.Abort(fmt.Errorf("failed to get value from ref path: %v", refPath))
			return
		}
		if ctxValue <= refVal.(float64) {
			ctx.Abort(fmt.Errorf("field `%s` value %v is not greater than `%v`", ctx.FieldPath(), ctxValue, refPath))
		}
	}
}
