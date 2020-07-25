// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/monitorables/gitlab/api/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetCountIssues provides a mock function with given fields: params
func (_m *Repository) GetCountIssues(params *models.IssuesParams) (int, error) {
	ret := _m.Called(params)

	var r0 int
	if rf, ok := ret.Get(0).(func(*models.IssuesParams) int); ok {
		r0 = rf(params)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*models.IssuesParams) error); ok {
		r1 = rf(params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMergeRequest provides a mock function with given fields: projectID, mergeRequestID
func (_m *Repository) GetMergeRequest(projectID int, mergeRequestID int) (*models.MergeRequest, error) {
	ret := _m.Called(projectID, mergeRequestID)

	var r0 *models.MergeRequest
	if rf, ok := ret.Get(0).(func(int, int) *models.MergeRequest); ok {
		r0 = rf(projectID, mergeRequestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MergeRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(projectID, mergeRequestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMergeRequestPipelines provides a mock function with given fields: projectID, mergeRequestID
func (_m *Repository) GetMergeRequestPipelines(projectID int, mergeRequestID int) ([]int, error) {
	ret := _m.Called(projectID, mergeRequestID)

	var r0 []int
	if rf, ok := ret.Get(0).(func(int, int) []int); ok {
		r0 = rf(projectID, mergeRequestID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(projectID, mergeRequestID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMergeRequests provides a mock function with given fields: projectID
func (_m *Repository) GetMergeRequests(projectID int) ([]models.MergeRequest, error) {
	ret := _m.Called(projectID)

	var r0 []models.MergeRequest
	if rf, ok := ret.Get(0).(func(int) []models.MergeRequest); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.MergeRequest)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPipeline provides a mock function with given fields: projectID, pipelineID
func (_m *Repository) GetPipeline(projectID int, pipelineID int) (*models.Pipeline, error) {
	ret := _m.Called(projectID, pipelineID)

	var r0 *models.Pipeline
	if rf, ok := ret.Get(0).(func(int, int) *models.Pipeline); ok {
		r0 = rf(projectID, pipelineID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Pipeline)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(projectID, pipelineID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPipelines provides a mock function with given fields: projectID, ref
func (_m *Repository) GetPipelines(projectID int, ref string) ([]int, error) {
	ret := _m.Called(projectID, ref)

	var r0 []int
	if rf, ok := ret.Get(0).(func(int, string) []int); ok {
		r0 = rf(projectID, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, string) error); ok {
		r1 = rf(projectID, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProject provides a mock function with given fields: projectID
func (_m *Repository) GetProject(projectID int) (*models.Project, error) {
	ret := _m.Called(projectID)

	var r0 *models.Project
	if rf, ok := ret.Get(0).(func(int) *models.Project); ok {
		r0 = rf(projectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(projectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
