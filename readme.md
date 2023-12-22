# go-bif-examine
An overcomplicated system to examine [BIF](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/bif_v1.htm) and [KEY](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/key_v1.htm) files used in Bioware's Infinity Engine (used for Baldur's Gate 2), extract its audio files, pass them to [Whisper](https://github.com/openai/whisper), and allow searching and downloading of files based on the character's speech within.

This is mostly so I have an excuse to write something in Go.

## Important Resources
- [gibberlings3/iesdp](https://github.com/gibberlings3/iesdp/) for the detailed documentation on the BIF and KEY file formats
- [openai/whisper](https://github.com/openai/whisper) to do the heavy lifting to determine what was actually said in each audio file

## Components
1. [ ] go-bif-examine
    1. [ ] Front end API via gRPC
        - Might need to use HTTP because a web frontend using gRPC might be more effort than I want to put into this
        1. [X] Upload KEY file to new project
            1. [X] Parse KEY data
            2. [X] Determine what BIF files contain what audio files
            3. [X] Save this information to DB
        2. [ ] Upload many BIF files to project started by KEY file upload
            1. [X] Parse BIF data
                1. [X] Save BIF file to S3 compatible storage
                2. [X] Find the file's SHA256 hash to deduplicate data and to allow more than one version of the same file
                3. [X] Save BIF entry data to DB
            2. [ ] Don't save BIF file if the KEY file says it doesn't have any audio files in it
        3. [ ] Schedule jobs to be sent to whisperer
        4. [ ] Allow unscheduling of jobs to be sent to whisperer
        5. [ ] Search results of whisperer
        6. [ ] Download audio file content
    2. [ ] Check for and send jobs to instances of whisper to extract the speech from each audio file
        1. [X] Download only the audio data, not the entire bif file 
        2. [ ] Save results to db
2. [ ] whisperer
    1. [ ] Wrap [whisper](https://github.com/openai/whisper) with a simple API so `go-bif-examine` can utilize it without needing to resort to [exec](https://pkg.go.dev/os/exec)
        1. [ ] Alternatively, use exec so I don't need to write any python and can reuse Go code?
    2. [ ] Allow for multiple instances of whisper to be spread across more than one machine for faster processing
3. [ ] examine-fe
    1. [ ] Barebones frontend used to interact with go-bif-examine
    2. [ ] To be embedded within go-bif-examine; no need to deploy separately
4. [ ] go-bif-examin-cli
    - CLI tool to make gRPC calls. Mostly for development purposes so I don't need to figure out the web stuff immediately
    - [X] quick and dirty CLI to test that downloading resources work
    - Commands, using cobra:
      1. [ ] get
         - [ ] projects
      2. [ ] delete
         - [ ] Project
      3. [ ] upload
         - [ ] key
         - [ ] bif
         - [ ] auto: point to game dir and automatically upload key and bifs?
      4. [ ] download
         - [ ] resource: save to file
5. [X] minio; deployed via docker compose
6. [X] postgresql; deployed via docker compose
7. [ ] Tests for everything

## Whisperer <-> go-bif-examine communication notes
Thinking the easiest option for now would be to use exec in go to execute the existing CLI tool that [openai/whisper](https://github.com/openai/whisper) provides. 

For simplicities' sake, I think I'll just have the tool check if whisper is installed, the nope out if it isn't there and print install instructions. If I run it in a container most of the time, that shouldn't be an issue. 
And given the wonder of abstracting everything behind an interface, I can always change it later.

Remember that this is a personal project that I'm doing for fun; I want to focus on the fun stuff, which is writing Go.

### Job scheduling
Thinking the easiest option would be for whisperer to connect as a gRPC client to go-bif-examine, then ask it for jobs.
