FROM golang

COPY . /wrk
RUN cd /wrk && go build

ENTRYPOINT ["/wrk/basic"]