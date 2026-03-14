package target

const AcceptRangesHeader = "accept-ranges"
const ContentLengthHeader = "content-length"

type Target struct {
	Url    string
	Client Client

	AcceptRanges  bool
	ContentLength int64
}

func New(u string) *Target {
	return &Target{Url: u}
}
