package delivery

import "immersiveApp/features/mentees"

func MenteeRequestToMenteeEntity(menteeRequest *MenteeRequest) mentees.MenteeEntity {
	return mentees.MenteeEntity{
		ClassId:         menteeRequest.ClassId,
		FullName:        menteeRequest.FullName,
		NickName:        menteeRequest.NickName,
		Email:           menteeRequest.Email,
		Phone:           menteeRequest.Phone,
		CurrentAddress:  menteeRequest.CurrentAddress,
		HomeAddress:     menteeRequest.HomeAddress,
		Telegram:        menteeRequest.Telegram,
		StatusId:        menteeRequest.StatusId,
		Gender:          menteeRequest.Gender,
		EducationType:   menteeRequest.EducationType,
		Major:           menteeRequest.Major,
		Graduate:        menteeRequest.Graduate,
		Institution:     menteeRequest.Institution,
		EmergencyName:   menteeRequest.EmergencyName,
		EmergencyPhone:  menteeRequest.EmergencyPhone,
		EmergencyStatus: menteeRequest.EmergencyStatus,
	}
}

func MenteeEntityToMenteeResponse(menteeEntity mentees.MenteeEntity) MenteeResponse {
	return MenteeResponse{
		Id:              menteeEntity.Id,
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
	}
}

// func ListMenteeFeedback

func ListMenteeEntityToMenteeResponse(dataCore []mentees.MenteeEntity) []MenteeResponse {
	var dataResponse []MenteeResponse
	for _, v := range dataCore {
		dataResponse = append(dataResponse, MenteeEntityToMenteeResponse(v))
	}
	return dataResponse
}
