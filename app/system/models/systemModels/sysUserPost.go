package systemModels

type SysUserPost struct {
	UserId int64 `db:"user_id"`
	PostId int64 `db:"post_id"`
}

func NewSysUserPost(userId int64, postId int64) *SysUserPost {
	return &SysUserPost{UserId: userId, PostId: postId}
}
