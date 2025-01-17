package zoom

import "fmt"

// EndMeetingOptions are the options to delete a meeting
type EndMeetingOptions struct {
	MeetingID int `url:"-"`
}

// MeetingStatusUpdate data parameters for meeting status updates
type MeetingStatusUpdate struct {
	Action string `json:"action"`
}

// EndMeetingPath - v2 delete a meeting
const EndMeetingPath = "/meetings/%d/status"

// EndMeeting calls PUT /meetings/{meetingID}/status
func EndMeeting(opts EndMeetingOptions) error {
	return defaultClient.EndMeeting(opts)
}

// EndMeeting calls PUT /meetings/{meetingID}/status
func (c *Client) EndMeeting(opts EndMeetingOptions) error {
	return c.requestV2(requestV2Opts{
		Method: Put,
		Path:   fmt.Sprintf(EndMeetingPath, opts.MeetingID),
		DataParameters: &MeetingStatusUpdate{
			Action: "end",
		},
		HeadResponse: true,
	})
}
