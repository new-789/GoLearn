package user

type Users struct {
	Username string
	Age      int
	Sex      string
	Edu      string
}

// 自定义构造函数,接收三个参数，返回值是指针类型的结构体变量名，以便在其它包中操作结构体中的字段,
// 注：自定义构造函数是为了方便在其他包中调用，所以函数名与返回值的首字母必须大写
func NewUser(username string, age int, sex string, edu string) *Users {
	// 初始化方式一
	//user := &Users{
	//	Username: username,
	//	Age:      age,
	//	Sex:      sex,
	//	Edu:      edu,
	//}

	// 初始化方式二: new 关键字
	user := new(Users)
	user.Username = username
	user.Age = age
	user.Sex = sex
	user.Edu = edu
	return user
	// 以上两种方式一样
}
