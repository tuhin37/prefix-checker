# Prefix-Matcher
This application is a web server written in golang. Its loads static wordlist files (prefixes) during runtime and
given an input string of variable length, this application tries to match the longest prefix that was loaded from 
wordlist files and returns


# Instruction - to run
## 1. Run from source
### 1.1 Make sure you put your wordlist (prefix list) files in ./prefix-wordlists directory. There are some sample files already present there.
The program loads the contents of these files to memory during initialization. 
Note, the wordlist file shloud have .txt extainsion and they should not be empty files. In the files, every word/prefix is a new line

### 1.2 build and run
To run the program, you must go go installed. Then execute the following command
```sh
go mod vendor 
go build -mod=vendor
./truecaller-prefix
```
This should start a webserver on you local port 5000. 
### 1.3 Health check
```sh
curl --location 'http://localhost:5000/health'
```
Expected response
```json
{
    "app": "prefix-checker",
    "status": "healthy",
    "version": "1.0.0"
}
```
### 1.4 Check prefix
To check prefix make sure that the server is up and running. Now to check the largest matching prefix against a input string execute the following

Example of a successful match 
```sh
curl --location 'http://localhost:5000/check-prefix/humanknowledgeisthemostpreciouscomodity'
```
response
```json
{
    "prefix": "humanknowledge",
    "status": "successful"
}
```

Example of a no match
```sh
curl --location 'http://localhost:5000/check-prefix/boldlygowherenomanhasevergonebefore'
```
response
```json
{
    "prefix": "",
    "status": "error"
}
```

## 2. Run from docker image
The repo can be easyly converted into a docker image. The Dockerfile and makefile are already included to make this very easy.
Make sure you have docker installed. If your system can not execute make commands, then run them manually after copying from the makefile
### 2.1 Docker image build
To build docker image execute the following command
```sh
make docker-build
```
This should create a new docker image in your system called `fidays/check-prefix:latest`

### 2.2 Create container from local image
To run the docker image execute the following command
```sh
make docker-run
```
Now you can execute the previously mentioned curls as it is.

### 2.3 Create container of off public image from dockerhub 
I have also uploaded the docker image in docker hub. If you do not wish to build from repo, you can alwasy use the existing docker image. Its public.
To run the server from the dockerhub image run the following command
```sh
docker run --rm \
--name prefix-checker \
-v <path-to-wordlist-folder>:/app/prefix-wordlists/ \
-p 5000:5000 \
-it fidays/check-prefix:latest
```
Note: replace `<path-to-wordlist-folder>` with an absolute path on your system taht contains one ot may wordlist files
Now you can execute the same curls defined in `section 1.3 & section 1.4 `

## 3 Cleanup
To clean up lets delete the docker image that was either downloaded from dockerhub or build on gthe local system, run the following command
```sh
make docker-rmi
```
 













