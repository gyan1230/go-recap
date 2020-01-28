# go-recap

GO RUN / BUILD / INSTALL

1. GO run compiles and run the application, without producing executables
2. Go build just compiles and produce executables : it compiles pkgs and dependecies.
3. Go build -o app will produce and executable called app in the current directory.
4. Go install behaves same like go build instead of creating executable in current directory , Crete exe in the $gopath/bin

GOOS=linux GOARCH=amd64 go build -o linuxapp
GOOS=darwin GOARCH=amd64 go build -o macapp