FROM node:alpine as node-builder 

WORKDIR /cms-frontend

COPY package*.json ./ 
RUN npm install 

COPY . . 
RUN npm run build

FROM golang:1.19.1

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . . 

COPY --from=node-builder /cms-frontend /cms-frontend

RUN go build -v -o /usr/local/bin ./... 
RUN cms migrate


ENTRYPOINT ["cms", "serve", "--http=0.0.0.0:8090", "--dir=pb_data", "--publicDir=/cms-frontend/pb_public"]

LABEL org.opencontainers.image.source="https://github.com/climactivity/cy-cms-2"
