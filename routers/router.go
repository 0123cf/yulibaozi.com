package routers

import (
	"github.com/devfeel/dotweb"
	"github.com/yulibaozi/yulibaozi.com/controllers"
)

// InitRoute 路由
func InitRoute(server *dotweb.HttpServer) {
	groupAPI := server.Group("/api")
	//文章部分
	artCtl := new(controllers.ArtiCleController)
	artAPI := groupAPI.Group("/art")
	artAPI.GET("/page", artCtl.Page) //分页获取文章
	artAPI.GET("/get", artCtl.Get)
	artAPI.GET("/hots", artCtl.Hot)
	artAPI.GET("/news", artCtl.NewN)
	artAPI.GET("/likes", artCtl.Like)
	artAPI.GET("/static", artCtl.Statistics)
	//标签和分类
	cateAPI := groupAPI.Group("/cate")
	cateCtl := new(controllers.CateController)
	cateAPI.GET("/hots", cateCtl.HotN)
	cateAPI.GET("/tags", cateCtl.Tags)
	cateAPI.GET("/cates", cateCtl.CatesAndArts)
	//评论部分
	comAPI := groupAPI.Group("/comment")
	commCtl := new(controllers.CommentController)
	comAPI.POST("/add", commCtl.Add)
	comAPI.GET("/tops", commCtl.TopN)
	comAPI.GET("/comments", commCtl.CommentsReplys)
	//获取首页信息
	groupAPI.GET("/home/get", new(controllers.HomeController).Single)
	//珍贵链接部分
	groupAPI.GET("/links", new(controllers.LinkController).List)
	//获取文字推荐部分
	groupAPI.GET("/rec/get", new(controllers.RecController).GetN)
	//获取轮播图部分
	groupAPI.GET("/slide/tops", new(controllers.SlideController).TopN)
}
