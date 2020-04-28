FROM scratch
ADD consulrouter /opt/consulrouter
ADD misc /opt/misc
ADD application.yml /opt/application.yml
EXPOSE 8080
WORKDIR /opt
ENTRYPOINT ["/opt/consulrouter"]