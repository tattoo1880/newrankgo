package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gotest/model"
	"gotest/utils"
	"sync"
)

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetXhs(c *gin.Context) {

	var xhs model.Xhs
	result, err := xhs.ReadAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func GetWeixin(c *gin.Context) {
	var weixin model.Weixin
	result, err := weixin.ReadAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func GetDy(c *gin.Context) {
	var dy model.Dy
	result, err := dy.ReadAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func GetSph(c *gin.Context) {
	var sph model.Sph
	result, err := sph.ReadAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func GetKs(c *gin.Context) {
	var ks model.Ks
	result, err := ks.ReadAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func GetBill(c *gin.Context) {
	var bill model.Bil
	result, err := bill.ReadAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error",
		})
		return

	}
	c.JSON(200, gin.H{
		"message": result,
	})
}

func Doit(c *gin.Context) {
	go utils.Doit()
	c.JSON(200, gin.H{
		"message": "success",
	})
}

func DeleteAll(c *gin.Context) {
	var xhs model.Xhs
	var weixin model.Weixin
	var dy model.Dy
	var sph model.Sph
	var ks model.Ks
	var bil model.Bil

	var wg sync.WaitGroup
	var mu sync.Mutex // 用于同步更新 result

	// 用于保存删除结果
	type Result struct {
		Xhs    string `json:"xhs"`
		Weixin string `json:"weixin"`
		Dy     string `json:"dy"`
		Sph    string `json:"sph"`
		Ks     string `json:"ks"`
		Bil    string `json:"bil"`
	}
	result := &Result{}
	wg.Add(6)

	// 启动并发 goroutines 来删除不同的数据
	go func() {
		defer wg.Done()
		if err := xhs.DeleteAll(); err != nil {
			mu.Lock() // 使用锁
			result.Xhs = "error"
			mu.Unlock()
		} else {
			mu.Lock() // 使用锁
			result.Xhs = "success"
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		if err := weixin.DeleteAll(); err != nil {
			mu.Lock()
			result.Weixin = "error"
			mu.Unlock()
		} else {
			mu.Lock()
			result.Weixin = "success"
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		if err := dy.DeleteAll(); err != nil {
			mu.Lock()
			result.Dy = "error"
			mu.Unlock()
		} else {
			mu.Lock()
			result.Dy = "success"
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		if err := sph.DeleteAll(); err != nil {
			mu.Lock()
			result.Sph = "error"
			mu.Unlock()
		} else {
			mu.Lock()
			result.Sph = "success"
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		if err := ks.DeleteAll(); err != nil {
			mu.Lock()
			result.Ks = "error"
			mu.Unlock()
		} else {
			mu.Lock()
			result.Ks = "success"
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		if err := bil.DeleteAll(); err != nil {
			mu.Lock()
			result.Bil = "error"
			mu.Unlock()
		} else {
			mu.Lock()
			result.Bil = "success"
			mu.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println(result)

	//todo 将结果返回给前端
	c.JSON(200, gin.H{
		"message": result,
	})
}
