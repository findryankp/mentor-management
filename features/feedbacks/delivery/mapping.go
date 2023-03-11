package delivery

import (
	"immersiveApp/features/feedbacks"
)

func FeedbackRequestToFeedbackEntity(feedbackRequest *FeedbackRequest) feedbacks.FeedbackEntity {
	return feedbacks.FeedbackEntity{
		Notes:    feedbackRequest.Notes,
		Proof:    feedbackRequest.Proof,
		UserId:   feedbackRequest.UserId,
		MenteeId: feedbackRequest.MenteeId,
		StatusId: feedbackRequest.StatusId,
		Approved: feedbackRequest.Approved,
	}
}

func FeedbackEntityToFeedbackResponse(feedbackEntity feedbacks.FeedbackEntity) FeedbackResponse {
	return FeedbackResponse{
		Id:       feedbackEntity.Id,
		Notes:    feedbackEntity.Notes,
		Proof:    feedbackEntity.Proof,
		UserId:   feedbackEntity.UserId,
		MenteeId: feedbackEntity.MenteeId,
		StatusId: feedbackEntity.StatusId,
		Approved: feedbackEntity.Approved,
	}
}

func ListFeedbackEntityToFeedbackResponse(feedbackEntity []feedbacks.FeedbackEntity) []FeedbackResponse {
	var dataResponse []FeedbackResponse
	for _, v := range feedbackEntity {
		dataResponse = append(dataResponse, FeedbackEntityToFeedbackResponse(v))
	}
	return dataResponse
}
