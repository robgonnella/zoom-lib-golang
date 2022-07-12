package zoom

import "fmt"

// UpdateMeetingOptions are the options to update a meeting with
type UpdateMeetingOptions struct {
	ID             string          `json:"-"`
	HostID         string          `json:"-"`
	Topic          string          `json:"topic,omitempty"`
	Type           MeetingType     `json:"type,omitempty"`
	StartTime      *Time           `json:"start_time,omitempty"`
	Duration       int             `json:"duration,omitempty"`
	Timezone       string          `json:"timezone,omitempty"`
	Password       string          `json:"password,omitempty"` // Max 10 characters. [a-z A-Z 0-9 @ - _ *]
	Agenda         string          `json:"agenda,omitempty"`
	TrackingFields []TrackingField `json:"tracking_fields,omitempty"`
	Settings       MeetingSettings `json:"settings,omitempty"`
}

// UpdateMeetingPath - v2 update a meeting
const UpdateMeetingPath = "/meetings/%s"

// UpdateMeeting calls PATCH /meetings/{meetingId}
func UpdateMeeting(opts UpdateMeetingOptions) (Meeting, error) {
	return defaultClient.UpdateMeeting(opts)
}

// UpdateMeeting calls PATCH /meetings/{meetingId}
// https://marketplace.zoom.us/docs/api-reference/zoom-api/meetings/meetingupdate
func (c *Client) UpdateMeeting(opts UpdateMeetingOptions) (Meeting, error) {
	var ret = Meeting{}
	return ret, c.requestV2(requestV2Opts{
		Method:         Patch,
		Path:           fmt.Sprintf(UpdateMeetingPath, opts.ID),
		DataParameters: &opts,
		Ret:            &ret,
	})
}
