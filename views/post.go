package views
// 发送人信息
type Sender struct {
	// 发送者 QQ 号
	UserId int64 `json:"user_id"`
	// 昵称
	Nickname string `json:"nickname"`
	// 性别，male 或 female 或 unknown
	Sex string `json:"sex"`
	// 年龄
	Age int32 `json:"age"`

	// 群消息信息
	Area  string `json:"area"`
	Level string `json:"level"`
	Role  string `json:"role"`
	Title string `json:"title"`
}

// 匿名信息
type Anonymous struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Flag string `json:"flag"`
}

// 通用qq消息结构：https://cqhttp.cc/docs/4.14/#/Post
type PostMsg struct {
	// 时间戳与自己的bot的qq号码
	Time   int64 `json:"time"`
	SelfId int64 `json:"self_id"`

	// 上报类型：
	// message 收到消息/群
	// notice  群、讨论组变动等通知类事件
	// request 加好友请求、加群请求／邀请
	PostType string `json:"post_type"`

	// 消息类型：
	// private 私戳
	// group   群组
	MessageType string `json:"message_type"`

	// 消息子类型：私聊与群组分别有定义
	SubType string `json:"sub_type"`

	// 消息 ID
	MessageId int32 `json:"message_id"`

	// 发送者 QQ 号
	UserId     int64  `json:"user_id"`
	Message    string `json:"message"`
	RawMessage string `json:"raw_message"`
	Sender     Sender `json:"sender"`

	// 群消息专有
	// 群组号码
	GroupId int64 `json:"group_id"`
	// 匿名消息
	Anonymous Anonymous `json:"anonymous"`
	// 字体
	Font int32 `json:"font"`
}