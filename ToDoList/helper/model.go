package helper

import (
	"time"

	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/domain"
	"github.com/ryhnfhrza/Golang-To-Do-List-API/model/web"
)

func ToAuthResponse(user domain.Users)web.AuthResponse{
	return web.AuthResponse{
		Username: user.Username,
	}

}

func ToAuthResponses(users [] domain.Users)[]web.AuthResponse{
	var authResponses []web.AuthResponse
	for _,u := range users{
		authResponses = append(authResponses, ToAuthResponse(u))
	}
	return authResponses
}

func ToLoginResponse(user domain.Users)web.AuthResponse{
	return web.AuthResponse{
		Username: user.Username,	
	}

}

func ToTasksResponse(tasks domain.Tasks)web.TaskResponse{

	createdAtFormatted := tasks.CreatedAt.Format(time.RFC3339)
	dueDateFormatted := ""
	if tasks.DueDate.Valid {
    dueDateFormatted = tasks.DueDate.Time.Format(time.RFC3339)
	}

	var completedStatus string
	if tasks.Completed == 1 {
		completedStatus = "Completed"
	} else {
		completedStatus = "Pending"
	}

	return web.TaskResponse{
		Title: tasks.Title,      
		Description: tasks.Description, 
		Completed:completedStatus,   
		DueDate: dueDateFormatted,     
		CreatedAt: createdAtFormatted,
	}

}

func ToTasksResponses(tasks []domain.Tasks) []web.TaskResponse {
	var taskResponses []web.TaskResponse
	for _, t := range tasks {
		taskResponses = append(taskResponses, ToTasksResponse(t))
	}
	return taskResponses
}

