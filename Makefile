build:
	go build -v -tags gtk_3_18 -gcflags "-N -l" -o bin/gtk3-app cmd/gtk3-main.go
