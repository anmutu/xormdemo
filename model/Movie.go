package model

type Movie struct {
	Id   int    `xorm:"not null pk autoincr comment('主键ID') INT(10)" form:"id"`
	Name string `xorm:"not null comment('电影名称') VARCHAR(50)" form:"name"`
}
