# Getting Started
## Acquire Binaries
### Prebuilt Binaries For Linux
See [latest](https://github.com/kaiiorg/go-bif-examine/releases/latest).
## Do-It-Yourself
See [development setup](/docs/dev.md).

## Setting up Minio and Postgres
1. Use docker compose to run the [provided compose file](/_docker/docker-compose.yml).
   - Change the passwords for both in the file, if desired.
2. Log into minio, create a bucket, and access keys:
   1. Navigate to `http://localhost:9001`.
   2. Log in using the `MINIO_ROOT_USER` and `MINIO_ROOT_PASSWORD` defined in the compose file.
   3. Administrator -> Buckets -> Create Bucket +.
   4. Set the name to `go-bif-examine`.
   5. User -> Access Keys -> Create Access Key +.
   6. Copy/paste the Access Key and Secret Key to your config.
   7. Click create.
3. Log into postgres and create the database.
   1. Navigate to `http://localhost:8080`.
   2. Log into postgres
      - Make sure to select `postgres` type
      - Username: `postgres`
      - Password: whatever was configured in the compose file
   3. Click `Create database`
   4. Set the name to `go_bif_examine`
   5. Click Save

## Running go-bif-examine
1. Copy the [example](/configs/example.hcl) config.
2. Make sure hostnames and credentials for postgres and minio match what was configured in [Setting up Minio and Postgres](#setting-up-minio-and-postgres)
    - Hint: if you intend to run whisperer on a machine that isn't also running go-bif-examine, make sure `s3.host` is a value accessible from outside the machine running go-bif-examine!
3. Start go-bif-examine: `./go-bif-examine -config ./path/to/config.hcl`

## Running whisperer
1. Install [whisper](https://github.com/openai/whisper#setup)
2. At least once instance of whisperer: `./whisperer`
   - If you're not running whisperer on the same machine running go-bif-examine, add the following argument: `-gprc-server ${GRPC_HOST}:${GPRC_PORT}`
       - Where `GRPC_HOST` == IP of machine running go-bif-examine
       - Where `GPRC_PORT` == port configured in `grpc.port` in the go-bif-examine config

## Running the CLI
For all commands: if you're not running the CLI from the machine running add the following argument: `--gprc ${GRPC_HOST}:${GPRC_PORT}`
- Where `GRPC_HOST` == IP of machine running go-bif-examine
- Where `GPRC_PORT` == port configured in `grpc.port` in the go-bif-examine config


### Upload your key and bif files to be processed
These steps use the `auto` command and assume your key and bif files are laid out as defined by the key file. For example, Baldur's Gate II Enhanced Edition V2.6.6.0 has them laid out like:
```
path/to/game/install/dir/
├─ data/
│  ├─ 25AmbSnd.bif
│  ├─ 25Areas.bif
│  ├─ ...
├─ chitin.key
```

Variables:
- `PROJECT_NAME`: `Baldur's Gate II Enhanced Edition V2.6.6.0`
- `GAME_DIR`: `"/mnt/c/Program Files (x86)/Steam/steamapps/common/Baldur's Gate II Enhanced Edition"`
    - I'm running go-bif-examine from WSL, but have steam installed in Windows
- `KEY_NAME`: `chitin.key`
- `PROJECT_ID`: provided by go-bif-examine and printed to console by the CLI
- `BIF_NAME_IN_KEY`: `'data/AREA300C.bif'`

1. Upload key file and automatically try to upload bif files:
    1.  `./go-bif-examine-cli upload auto --project-name ${PROJECT_NAME} ${KEY_NAME}`
    2. This will print the project ID to console
2. Manually upload a bif file, if needed
    1. `./go-bif-examine-cli upload bif --project-id ${PROJECT_ID} --name-in-key ${BIF_NAME_IN_KEY} ${GAME_DIR}/${BIF_NAME_IN_KEY}`
3. Wait... potentially a long time.
    - I ran two instances of whisperer running in WSL on my AMD Ryzen 5800X equipped machine for just over 25 hours and churned through a bit over 4500 of 17877 audio resources

### Searching results
The CLI doesn't yet support searching for results, make manual queries via adminer for now. Sorry.

### Download resource
1. `./go-bif-examine-cli download resource ${ID_OF_RESOURCE}`
   - See the DB for the resource ID 
