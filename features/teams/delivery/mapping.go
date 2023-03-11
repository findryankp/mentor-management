package delivery

import (
	"immersiveApp/features/teams"
)

func TeamRequestToTeamEntity(teamRequest *TeamRequest) teams.TeamEntity {
	return teams.TeamEntity{
		Name: teamRequest.Name,
	}
}

func TeamEntityToTeamResponse(team teams.TeamEntity) TeamResponse {
	return TeamResponse{
		Id:   team.Id,
		Name: team.Name,
	}
}

func ListTeamEntityToTeamResponse(teamEntity []teams.TeamEntity) []TeamResponse {
	var dataResponses []TeamResponse
	for _, v := range teamEntity {
		dataResponses = append(dataResponses, TeamEntityToTeamResponse(v))
	}
	return dataResponses
}
