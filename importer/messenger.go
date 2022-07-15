package importer

const (
	fbTypeGeneric     = "Generic"
	fbTypeShare       = "Share"
	fbTypeCall        = "Call"
	fbTypeSubscribe   = "Subscribe"
	fbTypeUnsubscribe = "Unsubscribe"

	fbThreadTypeRegular      = "Regular"
	fbThreadTypeRegularGroup = "RegularGroup"
)

type fbParticipant struct {
	Name string `json:"name"`
}

type fbPhoto struct {
	URI               string `json:"uri"`
	CreationTimestamp int64  `json:"creation_timestamp"`
}

type fbGIF struct {
	URI string `json:"uri"`
}

type fbShare struct {
	Link string `json:"link"`
}

type fbUser struct {
	Name string `json:"name"`
}

type fbReaction struct {
	Reaction string `json:"reaction"`
	Actor    string `json:"actor"`
}

type fbMessage struct {
	SenderName            string        `json:"sender_name"`
	TimestampMS           int64         `json:"timestamp_ms"`
	Content               string        `json:"content"`
	Photos                []*fbPhoto    `json:"photos"`
	GIFs                  []*fbGIF      `json:"gifs"`
	Share                 *fbShare      `json:"share"`
	Users                 []*fbUser     `json:"users"`
	Reactions             []*fbReaction `json:"reactions"`
	CallDuration          int           `json:"call_duration"`
	Type                  string        `json:"type"`
	IsUnsent              bool          `json:"is_unsent"`
	IsTakenDown           bool          `json:"is_taken_down"`
	BumpedMessageMetadata struct {
		BumpedMessage string `json:"bumped_message"`
		IsBumped      bool   `json:"is_bumped"`
	} `json:"bumped_message_metadata"`
}

type fbMagicWord struct {
	MagicWord           string `json:"magic_word"`
	CreationTimestampMS int64  `json:"creation_timestamp_ms"`
	AnimationEmoji      string `json:"animation_emoji"`
}

type fbContainer struct {
	Participants       []*fbParticipant `json:"participants"`
	Messages           []*fbMessage     `json:"messages"`
	Title              string           `json:"title"`
	IsStillParticipant bool             `json:"is_still_participant"`
	ThreadType         string           `json:"thread_type"`
	ThreadPath         string           `json:"thread_path"`
	MagicWords         []*fbMagicWord   `json:"magic_words"`
	Image              struct {
		URI               string `json:"uri"`
		CreationTimestamp int64  `json:"creation_timestamp"`
	} `json:"image"`
	JoinableMode struct {
		Mode int    `json:"mode"`
		Link string `json:"link"`
	} `json:"joinable_mode"`
}
