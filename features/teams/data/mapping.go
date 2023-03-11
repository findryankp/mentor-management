package data

import (
	"immersiveApp/features/teams"
)

func TeamEntityToTeam(teamEntity teams.TeamEntity) Team {
	return Team{
		Name: teamEntity.Name,
	}
}

func TeamToTeamEntity(team Team) teams.TeamEntity {
	return teams.TeamEntity{
		Id:        team.ID,
		Name:      team.Name,
		CreatedAt: team.CreatedAt,
		UpdatedAt: team.UpdatedAt,
	}
}

func ListTeamToTeamEntity(team []Team) []teams.TeamEntity {
	var teamEntity []teams.TeamEntity
	for _, v := range team {
		teamEntity = append(teamEntity, TeamToTeamEntity(v))
	}
	return teamEntity
}
