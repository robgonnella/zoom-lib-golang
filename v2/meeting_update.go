package zoom

import "fmt"

// UpdateMeetingPath - v2 update a meeting
const UpdateMeetingPath = "/meetings/%d"

// UpdateMeeting calls PATCH /meetings/{meetingId}
func UpdateMeeting(m *Meeting) (Meeting, error) {
	return defaultClient.UpdateMeeting(m)
}

// UpdateMeeting calls PATCH /meetings/{meetingId}
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingupdate
func (c *Client) UpdateMeeting(m *Meeting) (Meeting, error) {
	var ret = Meeting{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Patch,
		Path:           fmt.Sprintf(UpdateMeetingPath, m.ID),
		DataParameters: m,
		Ret:            &ret,
	})
}
