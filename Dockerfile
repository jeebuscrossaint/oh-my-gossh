FROM golang:1.22.6

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN make release
#RUN make firsttime

RUN mkdir -p /root/.config/ohmygossh/
COPY configs/gossh.toml /root/.config/ohmygossh/gossh.toml
COPY configs/ascii.txt /root/.config/ohmygossh/ascii.txt
COPY ./assets/MDStyle.json /root/.config/ohmygossh/MDStyle.json

CMD CMD ["./oh-my-gossh"]
