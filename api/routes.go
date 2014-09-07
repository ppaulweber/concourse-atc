package api

import "github.com/tedsuo/rata"

const (
	CreateBuild = "CreateBuild"
	BuildEvents = "BuildEvents"
	AbortBuild  = "AbortBuild"

	CreatePipe = "CreatePipe"
	WritePipe  = "WritePipe"
	ReadPipe   = "ReadPipe"
)

// pipeline = read-only
// builds & pipes api = read+write, irrespective of jobs
var Routes = rata.Routes{
	{Path: "/api/v1/builds", Method: "POST", Name: CreateBuild},
	{Path: "/api/v1/builds/:build_id/events", Method: "GET", Name: BuildEvents},
	// {Path: "/api/v1/builds/:build_id/abort", Method: "PUT", Name: AbortBuild},

	// {Path: "/api/v1/pipes", Method: "POST", Name: CreatePipe},
	// {Path: "/api/v1/pipes/:pipe", Method: "PUT", Name: WritePipe},
	// {Path: "/api/v1/pipes/:pipe", Method: "GET", Name: ReadPipe},
}