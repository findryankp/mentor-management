package data

import (
	"immersiveApp/features/feedbacks"
	// "immersiveApp/features/mentees"
	"immersiveApp/features/statuses"
	"immersiveApp/features/users"
	"reflect"
)

func FeedbackEntityToFeedback(feedbackEntity feedbacks.FeedbackEntity) Feedback {
	return Feedback{
		Notes:    feedbackEntity.Notes,
		Proof:    feedbackEntity.Proof,
		UserId:   feedbackEntity.UserId,
		MenteeId: feedbackEntity.MenteeId,
		StatusId: feedbackEntity.StatusId,
		Approved: feedbackEntity.Approved,
	}
}

func FeedbackToFeedbackEntity(feedback Feedback) feedbacks.FeedbackEntity {
	result := feedbacks.FeedbackEntity{
		Id:        feedback.ID,
		Notes:     feedback.Notes,
		Proof:     feedback.Proof,
		UserId:    feedback.UserId,
		MenteeId:  feedback.MenteeId,
		StatusId:  feedback.StatusId,
		Approved:  feedback.Approved,
		CreatedAt: feedback.CreatedAt,
		UpdatedAt: feedback.UpdatedAt,
	}

	u := reflect.ValueOf(feedback.User)
	// m := reflect.ValueOf(feedback.Mentee)
	s := reflect.ValueOf(feedback.Status)

	if !u.IsZero() {
		result.User = users.UserEntity{
			Id:       feedback.User.ID,
			FullName: feedback.User.FullName,
		}
	}

	// if !m.IsZero() {
	// 	result.Mentee = mentees.MenteeEntity{
	// 		Id:       feedback.Mentee.ID,
	// 		FullName: feedback.Mentee.FullName,
	// 	}
	// }

	if !s.IsZero() {
		result.Status = statuses.StatusEntity{
			Id: feedback.Status.ID,
			Name: feedback.Status.Name,
		}
	}

	return result
}

func ListFeedbackToFeedbackEntity(feedback []Feedback) []feedbacks.FeedbackEntity {
	var feedbackEntity []feedbacks.FeedbackEntity
	for _, v := range feedback {
		feedbackEntity = append(feedbackEntity, FeedbackToFeedbackEntity(v))
	}
	return feedbackEntity
}
