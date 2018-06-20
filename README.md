




#### Running individual service
- go to the specific services' directory
- run:
    docker build -t service_docker:1.0 .                    (change the name of the build name as convenience)
    docker run --rm -p 8080:8080 service_docker:1.0         (change the name of the build name as convenience)
    docker run --rm -t -i -p 8080:8080 service_docker:1.0


- additional commands:
    docker images : checking built images


## NOTES
- resources is used for json.Marshal to a model
- 