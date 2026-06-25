# TODO App — Golang, Docker, CI/CD & AWS EC2

A modern ToDo web application built using Go's standard library (`net/http`) and server-side HTML templates. This project was created to learn backend development with Golang, Docker containerization, CI/CD automation using GitHub Actions, and deployment on AWS EC2.

![Go](https://img.shields.io/badge/Go-1.26-00ADD8)
![Docker](https://img.shields.io/badge/Docker-Ready-2496ED)
![AWS](https://img.shields.io/badge/AWS-EC2-FF9900)
![GitHub Actions](https://img.shields.io/badge/CI%2FCD-GitHub%20Actions-2088FF)

## Features

- Add new tasks
- Delete existing tasks
- Mark tasks as completed
- Dynamic UI using Go HTML templates
- Static asset serving (CSS)
- Health check endpoint
- Dockerized application
- Multi-stage Docker build using Alpine Linux
- GitHub Actions CI/CD pipeline
- Automated deployment pipeline to AWS EC2

> **Note:** Tasks are currently stored in memory and will be reset whenever the application restarts.

## Tech Stack

| Component | Technology |
|-----------|------------|
| Language | Go 1.26 |
| Web Server | net/http |
| Templates | html/template |
| Frontend | HTML, CSS |
| Containerization | Docker |
| CI/CD | GitHub Actions |
| Deployment | AWS EC2 |
| OS | Linux |

## Project Structure

```text
todoApp/
├── .github/
│   └── workflows/
│       ├── ci-cd.yml
├── static/
│   └── css/
│       └── style.css
├── templates/
│   └── index.html
├── main.go
├── Dockerfile
├── go.mod
└── README.md
