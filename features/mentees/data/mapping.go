package data

import (
	"immersiveApp/features/classes"
	// d "immersiveApp/features/feedbacks/data"
	// m "immersiveApp/features/feedbacks"
	"immersiveApp/features/mentees"
	"immersiveApp/features/statuses"
	"reflect"
)

func MenteeEntityToMentee(menteeEntity mentees.MenteeEntity) Mentee {
	return Mentee{
		ClassId:         menteeEntity.ClassId,
		FullName:        menteeEntity.FullName,
		NickName:        menteeEntity.NickName,
		Email:           menteeEntity.Email,
		Phone:           menteeEntity.Phone,
		CurrentAddress:  menteeEntity.CurrentAddress,
		HomeAddress:     menteeEntity.HomeAddress,
		Telegram:        menteeEntity.Telegram,
		StatusId:        menteeEntity.StatusId,
		Gender:          menteeEntity.Gender,
		EducationType:   menteeEntity.EducationType,
		Major:           menteeEntity.Major,
		Graduate:        menteeEntity.Graduate,
		Institution:     menteeEntity.Institution,
		EmergencyName:   menteeEntity.EmergencyName,
		EmergencyPhone:  menteeEntity.EmergencyPhone,
		EmergencyStatus: menteeEntity.EmergencyStatus,
		// Feedbacks:       menteeEntity.Feedback,
	}
}

func MenteeToMenteeEntity(mentee Mentee) mentees.MenteeEntity {
	result := mentees.MenteeEntity{
		Id:              mentee.ID,
		ClassId:         mentee.ClassId,
		// Class:           classes.ClassEntity{},
		FullName:        mentee.FullName,
		NickName:        mentee.NickName,
		Email:           mentee.Email,
		Phone:           mentee.Phone,
		CurrentAddress:  mentee.CurrentAddress,
		HomeAddress:     mentee.HomeAddress,
		Telegram:        mentee.Telegram,
		StatusId:        mentee.StatusId,
		// Status:          statuses.StatusEntity{},
		Gender:          mentee.Gender,
		EducationType:   mentee.EducationType,
		Major:           mentee.Major,
		Graduate:        mentee.Graduate,
		Institution:     mentee.Institution,
		EmergencyName:   mentee.EmergencyName,
		EmergencyPhone:  mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
		// Feedback:        m.FeedbackEntity{},
		CreatedAt:       mentee.CreatedAt,
		UpdatedAt:       mentee.UpdatedAt,
	}

	c := reflect.ValueOf(mentee.Class)
	s := reflect.ValueOf(mentee.Status)

	if !c.IsZero() {
		result.Class = classes.ClassEntity{
			Id:        mentee.Class.ID,
			ClassName: mentee.Class.ClassName,
		}
	}

	if !s.IsZero() {
		result.Status = statuses.StatusEntity{
			Id:   mentee.Status.ID,
			Name: mentee.Status.Name,
		}
	}

	return result
}

// func MenteeFeedbackEntity(mentee Mentee) mentees.MenteeEntity {
// 	var feedback d.Feedback
// 	result := mentees.MenteeEntity{
// 		Id:       mentee.ID,
// 		FullName: mentee.FullName,
// 		Feedback: m.FeedbackEntity {
// 			Id: feedback.ID,
// 			StatusId: feedback.StatusId,
// 			Notes: feedback.Notes,
// 			Proof: feedback.Proof,
// 	},
// }
// return result
// }

func ListMenteeToMenteeEntity(mentee []Mentee) []mentees.MenteeEntity {
	var menteeEntity []mentees.MenteeEntity
	for _, v := range mentee {
		menteeEntity = append(menteeEntity, MenteeToMenteeEntity(v))
	}
	return menteeEntity
}
