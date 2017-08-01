FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY ./app /go/src/gitlab.com/one-more/acropole/app
WORKDIR /go/src/gitlab.com/one-more/acropole/app

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = testing ]; then \
        go test; \
	elif [ ${APP_ENV} = development ]; then \
        go get github.com/pilu/fresh && \
        fresh; \
	else \
        app; \
	fi