grpc {
    port = 50051
}

web {
    port = 8081
}

db {
    db_name  = "go_bif_examine"
    host     = "localhost"
    port     = 5432
    username = "postgres"
    password = "6a6009de9ad94e098b327db02706b3bc"
    ssl      = false
}

s3 {
    // Leave scheme, post, and port as their zero values to use AWS
    scheme = "http"
    host = "localhost"
    port = 9000
    region = "us-east-1" // MinIO does not use regions, but we still have to provide something
    force_path_style = true // Needed to force ${scheme}://${host}:${port}/${bucket}/${object_key} instead of ${scheme}://${bucket}.${host}:${port}/${object_key}
    bucket = "go-bif-examine" // Expects bucket to be setup ahead of time
    access_key = "KfOYD0GuCtCkNmRetguA"
    secret_key = "n62L0XIsUDOxufML4j5wnqdn8YxrbpnJWu78DiUt"
}
