package data

import (
	"immersiveApp/features/classes"
	"immersiveApp/features/users"
	"reflect"
)

func ClassEntityToClass(classEntity classes.ClassEntity) Class {
	return Class{
		ClassName:    classEntity.ClassName,
		PicUserId:    classEntity.PicUserId,
		DateStart:    classEntity.DateStart,
		DateGraduate: classEntity.DateGraduate,
	}
}

func ClassToClassEntity(class Class) classes.ClassEntity {
	result := classes.ClassEntity{
		Id:           class.ID,
		ClassName:    class.ClassName,
		PicUserId:    class.PicUserId,
		DateStart:    class.DateStart,
		DateGraduate: class.DateGraduate,
		CreatedAt:    class.CreatedAt,
		UpdatedAt:    class.UpdatedAt,
	}

	if !reflect.ValueOf(class.User).IsZero() {
		result.User = users.UserEntity{
			Id:          class.User.ID,
			TeamId:      class.User.TeamId,
			FullName:    class.User.FullName,
			Email:       class.User.Email,
			Role:        class.User.Role,
			Status:      class.User.Status,
			PhoneNumber: class.User.PhoneNumber,
			Address:     class.User.Address,
		}
	}
	return result
}

func ListClassToClassEntity(class []Class) []classes.ClassEntity {
	var classEntity []classes.ClassEntity
	for _, v := range class {
		classEntity = append(classEntity, ClassToClassEntity(v))
	}
	return classEntity
}
