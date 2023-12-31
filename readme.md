# go-bif-examine
An overcomplicated system to examine [BIF](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/bif_v1.htm) and [KEY](https://gibberlings3.github.io/iesdp/file_formats/ie_formats/key_v1.htm) files used in Bioware's Infinity Engine 
(used for games like Baldur's Gate 2), extract its audio files, pass them to [Whisper](https://github.com/openai/whisper), 
and allow searching and downloading of files based on the character's speech within.

See the [article on my personal website](https://kaiiorg.wtf/go-bif-examine/) for more thoughts and details.

## Scope
This is mostly to have an excuse to write something in Go, so there's a handful of things that
I've left out of scope either because they're not really needed for a self-hosted system that lives
entirely behind a firewall, or I didn't think would be fun to work on. If this were a real 
production application generating revenue, I'd put a lot more time and effort into these things.

- Encryption between services.
- Encrypted DB connection.
- HTTPS for S3.
- Using AWS S3 instead of minio.
    - Using minio mostly so I don't have to pay any money for a hobby project. I also like self-hosting stuff when possible/reasonable.
- Not using [exec](https://pkg.go.dev/os/exec) for the whisperer service.
    - Since the interface uses gRPC and gRPC supports python, it'd be More Correct™ to rewrite it in python and call Whisper directly, but I don't want to write python..
- Initial test coverage, TDD, etc.
    - I'm learning how to use some technologies here. Trying to write tests for something when I have no idea how it works is kinda hard.
    - The KEY and BIF files will need some example files that aren't covered by someone else's copywrite.
    - I'll eventually come back and write a lot of tests, but I'm not ready to do that.

## BYOA: Bring Your Own Assets
I do not (purposely) provide any game assets. I'd really rather avoid a letter from whoever owns 
the copywrite. You'll need to bring your own. I downloaded Baldur's Gate II: Enhanced Edition
from Steam and copied the files from there.

## Important Resources
- [gibberlings3/iesdp](https://github.com/gibberlings3/iesdp/) for the detailed documentation on the BIF and KEY file formats
- [openai/whisper](https://github.com/openai/whisper) to do the heavy lifting to determine what was actually said in each audio file

## Dev Setup
See [Dev Setup](/docs/dev.md)

## Getting Started
See [Getting Started](/docs/getting_started.md)
