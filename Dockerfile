FROM scratch
EXPOSE 8080
ENTRYPOINT ["/gophercon-workshop"]
COPY ./bin/ /