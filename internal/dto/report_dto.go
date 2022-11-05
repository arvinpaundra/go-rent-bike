package dto

type ReportDTO struct {
	UserId     string `json:"user_id" form:"user_id"`
	TitleIssue string `json:"title_issue" form:"title_issue"`
	BodyIssue  string `json:"body_issue" form:"body_issue"`
}
