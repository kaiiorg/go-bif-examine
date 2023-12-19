# go-bif-examine
An overcomplicated system to examine [BIF](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/bif_v1.htm) and [KEY](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/key_v1.htm) files used in Bioware's Infinity Engine (used for Baldur's Gate 2), extract its audio files, pass them to [Whisper](https://github.com/openai/whisper), and allow searching and downloading of files based on the character's speech within.

This is mostly so I have an excuse to write something in Go.

## Important Resources
- [gibberlings3/iesdp](https://github.com/gibberlings3/iesdp/) for the detailed documentation on the BIF and KEY file formats
- [openai/whisper](https://github.com/openai/whisper) to do the heavy lifting to determine what was actually said in each audio file

## Components
1. [ ] go-bif-examine
    1. [ ] Allow uploading of KEY and BIF files via gRPC and saving to S3 compatible storage
        - Might need to allow uploading via HTTP; might be more effort than I want to put into it for a web frontend to support gRPC
    1. [X] Parse KEY data
        1. [X] Determine what BIF files contain what audio files
    1. [X] Parse BIF data
    1. [ ] Send jobs to instances of whisper to extract the speech from each audio file
        1. [ ] Save results to db
1. [ ] whisperer
    1. [ ] Wrap [whisper](https://github.com/openai/whisper) with a simple API so `go-bif-examine` can utilize it without needing to resort to [exec](https://pkg.go.dev/os/exec)
        1. [ ] Alternatively, use exec so I don't need to write any python?
    1. [ ] Allow for multiple instances of whisper to be spread across more than one machine for faster processing
1. [ ] examine-fe
    1. [ ] Barebones frontend used to interact with go-bif-examine
       1. [ ] To be embedded within go-bif-examine; no need to deploy separately
1. [X] minio; deployed via docker compose
1. [X] postgresql; deployed via docker compose
1. [ ] Tests for everything
