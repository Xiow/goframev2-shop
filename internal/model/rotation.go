package model

// RotationCreateUpdateBase 创建/修改内容基类
type RotationCreateUpdateBase struct {
	PicUrl string // 内容模型
	Sort   int    // 标题
	Link   string // 栏目ID

	//Rotation string   // 内容
	//Brief    string   // 摘要
	//Thumb    string   // 缩略图
	//Tags     []string // 标签名称列表，以JSON存储
	//Referer  string   // 内容来源，例如github/gitee
}

// RotationCreateInput 创建内容
type RotationCreateInput struct {
	RotationCreateUpdateBase
}

// RotationCreateOutput 创建内容返回结果
type RotationCreateOutput struct {
	RotationId int `json:"rotation_id"`
}

// ContentUpdateInput 修改内容
type RotationUpdateInput struct {
	RotationCreateUpdateBase
	Id uint
}
