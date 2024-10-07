# Docker Containers CLI

## Overview

This project is a Go-based Command-Line Interface (CLI) tool that lists all the Docker containers running on your system. It interacts with Docker using the Docker Go SDK, providing a simple interface to manage and monitor containers.

## Features

- List all Docker containers running in a beautiful terminal table
- Display detailed information for each container, including:

    (a) Container ID

    (b) Image
    
    (c) Command
    
    (d) Creation time
    
    (e) Status
    
    (f) Ports

- Real-time updates of container status

## Prerequisites

- Go 1.16 or higher
- Docker installed and running on your system

## Installation

- Clone the repository: clone [This repo](https://github.com/GitataY/docker-containers-cli.git) 
- Change to the project directory:  `docker-containers-cli`
- Build the application: `build -o docker-containers-cli`


## Usage
Run the application: `./docker-containers-cli`

Once the application starts, you'll see a table with all your Docker containers. Use the following keys to interact with the interface:

- Arrow keys: Navigate through the table
Enter: Select a container (functionality to be implemented)
- Esc: Toggle focus on the table
- q or Ctrl+C: Quit the application

## Project Structure

- main.go: Entry point of the application, sets up the TUI and table model
- models.go: Contains the main logic for fetching Docker container data and updating the TUI

## Dependencies

This project uses the following external libraries:

- github.com/charmbracelet/bubbles: TUI components
- github.com/charmbracelet/bubbletea: TUI framework
- github.com/charmbracelet/lipgloss: Style definitions for TUI
- github.com/docker/docker: Docker API client

Dependencies are managed using Go modules. To install all dependencies, run: `mod tidy`

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

- Fork the repository
- Create your feature branch (`git checkout -b feature/AmazingFeature`)
- Commit your changes (`git commit -m 'Add some AmazingFeature'`)
- Push to the branch (`git push origin feature/AmazingFeature`)
- Open a Pull Request