package types

// degree module event types
const (
	EventTypeCreateDegree = "CreateDegree"
	// TODO: Create your event types
	// EventType<Action>    		= "action"

	// TODO: Create keys fo your events, the values will be derivided from the msg
	AttributeCreator  		= "creator"
	AttributeSubject  		= "subject"
	AttributeBatch  		= "batch"

	// TODO: Some events may not have values for that reason you want to emit that something happened.
	// AttributeValueDoubleSign = "double_sign"

	AttributeValueCategory = ModuleName
)
