package e

var MsgFlags = map[int]string{
	Success:       "success",
	Error:         "fail",
	InvalidParams: "请求参数错误",
	AuthError:     "用户名或密码错误",
	
	TagExistError:    "标签已存在",
	TagGetError:      "标签获取失败",
	TagNotExistError: "标签不存在",
	TagGetAllError:   "获取所有标签失败",
	TagCreateError:   "添加标签失败",
	TagEditError:     "编辑标签失败",
	TagDeleteError:   "删除标签失败",
	
	CategoryExistError:    "分类已存在",
	CategoryGetError:      "分类获取失败",
	CategoryNotExistError: "分类不存在",
	CategoryGetAllError:   "获取所有分类失败",
	CategoryCreateError:   "添加分类失败",
	CategoryEditError:     "编辑分类失败",
	CategoryDeleteError:   "删除分类失败",
	
	ArticleExistError:    "文章已存在",
	ArticleGetError:      "文章获取失败",
	ArticleNotExistError: "文章不存在",
	ArticleGetAllError:   "获取所有文章失败",
	ArticleCountError:    "统计文章失败",
	ArticleCreateError:   "添加文章失败",
	ArticleEditError:     "编辑文章失败",
	ArticleDeleteError:   "删除文章失败",
	
	TokenCheckError:   "Token鉴权失败",
	TokenTimeoutError: "Token已超时",
	TokenGenError:     "Token生成失败",
	
	UserGetError:      "用户信息获取失败",
	UserEditError:     "用户信息更新失败",
	RestPasswordError: "修改密码失败",
	
	GetUploadImageError:  "获取上传文件信息失败",
	SaveUploadImageError: "保存上传文件失败",
	
	CommentGetAllError: "获取所有评论失败",
	SearchArticleError: "搜索文章失败",
	CommentGetError:    "评论获取失败",
	CommentDeleteError: "删除评论失败",
	
	SoupGetAllError:  "获取所有鸡汤失败",
	SoupCreateError:  "添加鸡汤失败",
	SoupGetError:     "鸡汤获取失败",
	SoupDeleteError:  "删除鸡汤失败",
	SoupEditError:    "编辑鸡汤失败",
	SoupGetRandError: "随机获取鸡汤失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	
	return MsgFlags[Error]
}
