package grpc2

type Storage interface {
	// Get message by id. For admin only
	Get(uint) (*Message, error)

	//GetAll ...
	GetAll() []*Message

	//Update ...
	Update(Message, uint) (*Message, error)

	//Add ...
	Add(Message) *Message

	// Delete message by ID - for admin use
	Delete(uint) (*Message, error)

	// Len ...
	Len() uint

	//GetsUserMessage
	GetUserMessage(Message) ([]Message, error)

	//DeleteUserMessage deletes user message
	DeleteUserMessage(Message, []Message) ([]Message, error)
}
