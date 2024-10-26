# ip-geolocation
Use the [Abstract](https://app.abstractapi.com/) IP Geolocation services to obtain information about an IP address

## Usage
Run the application by providing an IP address on the command line or as piped input.

The application will call the [IP Geolocation API](https://app.abstractapi.com/api/ip-geolocation) and print the results which, for now, is a formatted version of the raw JSON

## API Key
The API requires a key and the application will attempt to load that key from a file called `.env` in the current directory.

The format of the `.env` is:
```
api_key: <api key goes here>
```

The key can be obtained by signing up with [Abstract](https://app.abstractapi.com/users/login) and obtaining a private key.

## Running the application 
Run from source code with the IP address as a command line argument or with piped input:
```shell
go run . [ipaddress]

ipcheck | go run .
```

> Note that piped input requires something to provide the IP address, such as the [go-ipcheck](https://github.com/Motivesoft/go-ipcheck) or [ipcheck](https://github.com/Motivesoft/ipcheck) GitHub projects, or using `curl https://ident.me`

## Building an executable
A portable executable can be built and, for example, put onto the user's path for easy access.

```shell
# Normal executable
go build .

# Normal executable but with some extraneous information removed to make it smaller
go build -ldflags "-s -w" .
```

> The built version operates identically to when [running the application](#running-the-application) from its source code.