# This is still a work in progress!

# go-bif-examine
An overcomplicated system to examine [BIF](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/bif_v1.htm) and [KEY](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/key_v1.htm) files used in Bioware's Infinity Engine 
(used for games like Baldur's Gate 2), extract its audio files, pass them to [Whisper](https://github.com/openai/whisper), 
and allow searching and downloading of files based on the character's speech within.

## Scope
This is mostly to have an excuse to write something in Go, so there's a handful of things that
I've left out of scope either because they're not really needed for a self-hosted system that lives
entirely behind a firewall, or I didn't think would be fun to work on. If this were a real 
production application generating revenue, I'd put a lot more time and effort into these things.

- Encryption between services
- Encrypted DB connection
- HTTPS for S3
- Using AWS S3 instead of minio
  - Using minio mostly so I don't have to pay any money for a hobby project. I also like self-hosting stuff when possible/reasonable
- Not using [exec](https://pkg.go.dev/os/exec) for the whisperer service
  - Since the interface uses gRPC and gRPC supports python, it'd be More Correct™ to rewrite it in python and call Whisper directly, but I don't want to write python.

## BYOA: Bring Your Own Assets
I do not (purposely) provide any game assets. I'd really rather avoid a letter from whoever owns 
the copywrite. You'll need to bring your own. I downloaded Baldur's Gate II: Enhanced Edition
from Steam and copied the files from there.

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
        3. [ ] Search results of whisperer
        4. [X] Download audio file content
    2. [ ] Save whisperer results to db
2. [ ] whisperer
   1. [ ] Using Exec so I don't need to write any python
   2. [X] Allow for multiple instances of whisper to be spread across more than one machine for faster processing
   3. [X] Download only the audio data, not the entire bif file
3. [ ] examine-fe
    - This has been shelved for the time being
    1. [ ] Barebones frontend used to interact with go-bif-examine
    2. [ ] To be embedded within go-bif-examine; no need to deploy separately
4. [ ] go-bif-examin-cli
    - CLI tool to make gRPC calls. Mostly for development purposes so I don't need to figure out the web stuff immediately
    - Commands, using cobra:
      1. [ ] get
         - [X] projects
      2. [ ] delete
         - [X] Project
      3. [ ] upload
         - [X] key
         - [X] bif
         - [ ] auto: point to game dir and automatically upload key and bifs?
      4. [ ] download
         - [X] resource: save to file
5. [X] minio; deployed via docker compose
6. [X] postgresql; deployed via docker compose
7. [ ] Tests for everything
