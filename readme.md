# go-bif-examine
An overcomplicated system to examine [BIF](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/bif_v1.htm) and [KEY](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/key_v1.htm) files used in Bioware's Infinity Engine (used for Baldur's Gate 2, for example), extract its audio files, pass them to [Whisper](https://github.com/openai/whisper), and allow searching and downloading of files based on the character's speech within.

I wrote [something in C++](https://github.com/kaiiorg/bifex) a few years ago to do a subset of what this package will need to do, so I have a basic idea of what needs to be done.

This is mostly so I have an excuse to write something in Go.

## Important Resources
- [gibberlings3/iesdp](https://github.com/gibberlings3/iesdp/) for the detailed documentation on the BIF and KEY file formats
- [openai/whisper](https://github.com/openai/whisper) to do the heavy lifting to determine what was actually said in each audio file

## Components
1. go-bif-examine
    1. Allow uploading of KEY and BIF files via gRPC
        1. Save KEY and BIF data to minio
    1. Parse KEY data
        1. Determine what BIF files contain what audio files
    1. Parse BIF data
        1. Extract audio files, save them seperately for later processing by whisper and download by the user
    1. Send jobs to instances of whisper to extract the speech from each audio file
        1. Save results to db
1. examine-fe
    1. Barebones frontend used to interact with go-bif-examine
    1. To be embedded within go-bif-examine; no need to deploy seperately
    1. Don't judge me; I'm not a front end dev
1. whisperer
    1. Wrap [whisper](https://github.com/openai/whisper) with a simple API so go-bif-examine can utilize it without needing to resort to [exec](https://pkg.go.dev/os/exec)
    1. Potentially allow for multiple instances of whisper to be spread across more than one machine to allow for faster processing
1. minio
    1. I'm not using AWS if I don't need to
1. postgresql
    1. To do database things