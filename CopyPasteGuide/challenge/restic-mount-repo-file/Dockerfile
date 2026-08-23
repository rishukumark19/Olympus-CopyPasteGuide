FROM public.ecr.aws/d3j8x8q7/olympus-base-go:latest
WORKDIR /app
COPY . .
ENV CHALLENGE_DOCKER=1
RUN go mod download
CMD ["/bin/bash"]
