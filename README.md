# go-loader
Go-based Docker App Loader
Auto-runs uploaded builds with a Docker Container

# Structures
/ Home Page

/ping Check Docker Container and show status

/upload Folder of build files are loaded

/run Docker Build router

/file Can see uploaded build files

# Default
Application default run http://localhost:8080 and support Gin Framework

Output supported JSON format

# Functions
saveFileHandler - save build file

status - Check Docker Containers and report

home - Gin Framework default router

build - New Container build function

upload - staticFS (use default upload folder)
