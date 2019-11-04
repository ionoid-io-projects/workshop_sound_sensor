# Introduction
Making it work

# Wiring

![Sound sensor](doc/img/scheme.png)

|sound sensor      |  Raspberry Pi  |
|------------------|----------------|
| Ground 		   | GND            |
| Vcc / + 		   | 5V             |
| DO 			   | GPIO 21        |

# How to build

Compile sound.go like this
```
go get github.com/stianeikeland/go-rpio
env GOOS=linux GOARCH=arm GOARM=6 go build sound.go
```
Copy the generated file to your Raspberry Pi device and execute it with this command

```
./sound
```

Happy blinking 