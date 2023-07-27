package dashboard

import "embed"

//go:embed build/*
var DashboardStaticFiles embed.FS
