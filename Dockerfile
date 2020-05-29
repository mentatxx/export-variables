FROM golang:latest as build
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o export-variables . 

FROM scratch
COPY --from=build /app/export-variables /
CMD ["/export-variables"]