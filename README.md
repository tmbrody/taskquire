# Taskquire

`Taskquire` is a web application that enables users to join teams and create tasks as they work though projects from different organizations. The app includes a `Generate Tasks` button that fancifies the basic task-generation process by using OpenAI's API to generate tasks that give users new ideas and perspectives on what needs to be tackled to accomplish any given project.

## Motivation

I've been using GitHub Copilot to act as a pair programmer to speed up my coding and help me tackle any challenges encountered. I wanted to take things a step further and create a copilot for task planning and generation to help users more quickly determine what they need to work on so they can spend more time working on their code.

## 🚀 Quick Start

1. Navigate to [taskquire.com](https://taskquire.com)
2. Create a user account
3. Create a team
4. Create an org
5. Create a project for the org
6. Add your team to the project
5. Use the `Create Task` or `Generate Tasks` buttons to start making tasks that will guide you through every phase of the project

## 📖 Usage

* Create an organization
* Create projects for that organization
* Assign teams to work on your projects
* Create a team
* Add users to be your team members
* Become a member of other teams
* Create tasks to set clear goals towards finishing a project
* Use the `Generate Tasks` button and get different perspectives on how to tackle a project

## 🤝 Contributing

### Clone the repo

```bash
git clone https://github.com/tmbrody/taskquire@latest
cd taskquire
```

### Build the project

```bash
cd backend
go build -o out && ./out
```

### Run the project

```bash
cd frontend
nvm run dev
```

### Run the tests

```bash
cd backend
go test ./...
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.