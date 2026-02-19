package semester

import (
	"dept-collector/internal/models"
	"dept-collector/internal/responseTypes"
)

func formatSemesterResponse(semesters []models.Semester) []responseTypes.SemesterResponse {
	var semesterResponse []responseTypes.SemesterResponse
	for _, semester := range semesters {
		semesterResponse = append(semesterResponse, responseTypes.SemesterResponse{
			ID:        semester.ID,
			Name:      semester.Name,
			StartDate: semester.StartDate,
			EndDate:   semester.EndDate,
			CreatedAt: semester.CreatedAt,
			UpdatedAt: semester.UpdatedAt,
		})
	}
	return semesterResponse
}
