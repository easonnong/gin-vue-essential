package controller

import (
	"strconv"

	"github.com/easonnong/gin-vue-essential/model"
	"github.com/easonnong/gin-vue-essential/repository"
	"github.com/easonnong/gin-vue-essential/response"
	"github.com/easonnong/gin-vue-essential/vo"
	"github.com/gin-gonic/gin"
)

//定义增删改查的接口
type ICategoryController interface {
	RestController
}

type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() CategoryController {
	repository := repository.NewCategoryRepository()
	//添加自动迁移
	repository.DB.AutoMigrate(model.Category{})

	return CategoryController{
		Repository: repository,
	}
}

func (c CategoryController) Create(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory vo.CategoryRequest
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
		return
	}

	category, err := c.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
	}

	response.Success(ctx, "创建成功", gin.H{
		"category": category,
	})
}

func (c CategoryController) Update(ctx *gin.Context) {
	//绑定body中的参数
	var requestCategory model.Category
	/*此方法可扩展性差
	ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
		return
	} */
	if err := ctx.ShouldBind(&requestCategory); err != nil {
		response.Fail(ctx, "数据验证错误，分类名称必填", nil)
		return
	}

	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	updateCategory, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}

	//更新分类
	category, err := c.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		response.Fail(ctx, "修改失败", nil)
		return
	}

	response.Success(ctx, "修改成功", gin.H{"category": category})
}

func (c CategoryController) Show(ctx *gin.Context) {
	//获取path中的参数
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	category, err := c.Repository.SelectById(categoryId)
	if err != nil {
		response.Fail(ctx, "分类不存在", nil)
		return
	}

	response.Success(ctx, "查询成功", gin.H{"category": category})
}

func (c CategoryController) Delete(ctx *gin.Context) {
	//获取path中的参数
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))

	if err := c.Repository.DeleteById(categoryId); err != nil {
		response.Fail(ctx, "删除失败，请重试", nil)
		return
	}

	response.Success(ctx, "删除成功", nil)
}
