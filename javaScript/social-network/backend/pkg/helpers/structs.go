package helpers

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	DateOfBirth string `json:"dateofbirth"`
	Avatar      string `json:"avatar"`
	Nickname    string `json:"nickname"`
	About       string `json:"about"`
	Public      bool   `json:"public"`
	CurrentUser bool   `json:"currentuser"`
	Following   bool   `json:"following"`
	Pending     bool   `json:"pending"`
	Follower    bool   `json:"follower"`
}

type SessionResponse struct {
	SessionID string `json:"session_id"`
	UserID    int    `json:"user_id"`
}

type Post struct {
	PostID    int       `json:"postId"`
	UserID    int       `json:"userId"`
	Privacy   int       `json:"privacy"` //0 public, 1 follower, 2 specific, 3 group
	Followers []int     `json:"followers"`
	Text      string    `json:"text"`
	Image     string    `json:"image"`
	Comments  []Comment `json:"comments"`
	Created   string    `json:"created"`
}

type Comment struct {
	CommentID int    `json:"commentId"`
	PostID    int    `json:"postId"`
	UserID    int    `json:"userId"`
	Text      string `json:"text"`
	Image     string `json:"image"`
	Created   string `json:"created"`
}
type Group struct {
	GroupID     int    `json:"groupId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AdminID     int    `json:"adminId"`
	Joined      bool   `json:"joined"`    //2
	Invited     bool   `json:"invited"`   //1
	Requested   bool   `json:"requested"` //0
}

type GroupEvent struct {
	EventID     int    `json:"eventId"`
	GroupID     int    `json:"groupId"`
	CreatorID   int    `json:"creatorId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Created     string `json:"created"`
	Going       bool   `json:"going"`
	NotGoing    bool   `json:"notgoing"`
}

type Notification struct {
	// follow_id wants to follow you
	// you have been invited to group_id
	// user_id wants to join group id
	// event_id created in group_id
	NotificationType string     `json:"notificationType"`
	User             User       `json:"user"`
	UserID int `json:"userId"`
	GroupID              int        `json:"groupId"`
	GroupEvent       GroupEvent `json:"groupEvent"`
}

type Message struct {
	ToID    int    `json:"toId"`
	FromID  int    `json:"fromId"`
	Message string `json:"message"`
	Created string `json:"created"`
}
