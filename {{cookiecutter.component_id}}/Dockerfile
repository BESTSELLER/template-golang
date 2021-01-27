FROM golang:1.15 as build

WORKDIR $GOPATH/src/github.com/{{cookiecutter.destination.organization}}/{{cookiecutter.component_id}}
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go install

FROM alpine

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/{{cookiecutter.destination.organization}}/{{cookiecutter.component_id}}"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

WORKDIR /
COPY --from=build /go/bin/{{cookiecutter.component_id}} /

CMD /{{cookiecutter.component_id}}
EXPOSE 3000
