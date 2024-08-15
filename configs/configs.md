# Config Guide

The configuration file is located in $HOME/.config/ohmygossh/gossh.toml

Below is an example of the configuration file, if you are not familiar with TOML. The `title`, `ssh`, and `color` sections are **required**, while the user (YOU!) can make as many sections as they want, provided it follows the format below. 

### title
- name: The text/big header/ usually ASCII art that you want in teh center of the terminal
- subtitle: The subtitle of the title, a small blurb or quote.
- tab: The tab name that you want to be displayed in the terminal

### ssh
- status: The status of the SSH connection, either true or false. If false the TUI displays locally on your window, if true it will display the remote server.
- host: The host of the SSH connection, usually an IP address or domain name.\
- port: The port of the SSH connection, usually 23 or 19 for this case.

### color
- active: The color of the active tab
- inactive: The color of the inactive tab

### projects
- file: The file name of the project, should be a directory to an exact markdown file, ie $HOME/.config/ohmygossh/projects/example.md
- name: The name of the project
- about: A small blurb about the project

Below is an example of the configuration file in TOML format.

```toml


[title]
name = "Golang"
subtitle = "Memory safe programming language"
tab = "oh my gossh"

[ssh]
status = "true"
host = "0.0.0.0"
port = "19"

[color]
active = "#00FF00"
inactive = "#000000"

[example]
file = "example"
name = "foo"
about = "Example project"

```

____________________