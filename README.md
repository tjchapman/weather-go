## Weather CLI Tool built with Go

### Requirments:

- API Key from https://www.weatherapi.com/
- Add API Key to .env file

### Run the tool:

```
go run main.go
```

### To run the tool with your chosen location e.g. London:

```
go run main.go London
```

### Compile the package:

```
go build
```

### Move the pacakge so you just call 'weather' in your terminal (embed API Key into main.go if doing this)

```
sudo mv weather /usr/local/bin
```
