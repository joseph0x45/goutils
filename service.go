package goutils

import (
	"html/template"
	"log"
	"os"
	"os/user"
)

var serviceFile = `
[Unit]
Description={{.Description}}
After=network.target
Wants=network.target

[Service]
Type=simple

User={{.User}}

EnvironmentFile=-{{.Config}}

ExecStart=/usr/local/bin/{{.AppName}}
Restart=on-failure
RestartSec=5

NoNewPrivileges=true

StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
`

func GenerateServiceFile(description string) {
	appNameIsNotEmpty()
	t, err := template.New("service").Parse(serviceFile)
	if err != nil {
		log.Fatalln(err)
	}
	serviceFileName := appName + ".service"
	f, err := os.Create(serviceFileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	if err := t.Execute(f, map[string]string{
		"User":        user.Username,
		"Config":      AppConfigFile(user, "conf"),
		"Description": description,
		"AppName":     appName,
	}); err != nil {
		log.Fatalln(err)
	}
	if err := f.Sync(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Service file created")
}
