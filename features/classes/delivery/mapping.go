package delivery

import (
	"immersiveApp/features/classes"
	"immersiveApp/features/users/delivery"
	"reflect"
)

func ClassRequestToClassEntity(classRequest *ClassRequest) classes.ClassEntity {
	return classes.ClassEntity{
		ClassName:    classRequest.ClassName,
		PicUserId:    classRequest.PicUserId,
		DateStart:    classRequest.DateStart,
		DateGraduate: classRequest.DateGraduate,
	}
}

func ClassEntityToClassResponse(classEntity classes.ClassEntity) ClassResponse {
	result := ClassResponse{
		Id:           classEntity.Id,
		ClassName:    classEntity.ClassName,
		PicUserId:    classEntity.PicUserId,
		DateStart:    classEntity.DateStart,
		DateGraduate: classEntity.DateGraduate,
	}

	if !reflect.ValueOf(classEntity.User).IsZero() {
		result.User = delivery.UserResponse{
			FullName: classEntity.User.FullName,
			Email:    classEntity.User.Email,
		}
	}

	return result
}

func ListClassEntityToClassResponse(classEntity []classes.ClassEntity) []ClassResponse {
	var dataResponses []ClassResponse
	for _, v := range classEntity {
		dataResponses = append(dataResponses, ClassEntityToClassResponse(v))
	}
	return dataResponses
}
