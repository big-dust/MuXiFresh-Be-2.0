// Code generated by goctl. DO NOT EDIT.
package types

type GetReviewReq struct {
	Authorization string `header:"Authorization"`
	Year          int    `json:"year"`
	Group         string `json:"group"`
	Season        string `json:"season,options=[autumn,spring]"`
	Grade         string `json:"grade,optional"`
	School        string `json:"school,optional"`
	Status        string `json:"status,optional"`
}

type Row struct {
	ScheduleID      string `json:"schedule_id"`
	Name            string `json:"name"`
	Grade           string `json:"grader"`
	School          string `json:"school"`
	Group           string `json:"group"`
	ExamStuatus     string `json:"exam_status"`
	AdmissionStatus string `json:"admission_status"`
}

type GetReviewResp struct {
	Rows []Row `json:"rows"`
}

type SetAdmissionStatusReq struct {
	Authorization string `header:"Authorization"`
	ScheduleID    string `json:"schedule_id"`
	NewStatus     string `json:"new_status,options=[已报名,实习期,已转正]"`
}

type SetAdmissionStatusResp struct {
	Flag bool `json:"flag"`
}
