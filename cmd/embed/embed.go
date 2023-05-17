package embed

import "embed"

//go:embed build/*
var DashboardStaticFiles embed.FS
