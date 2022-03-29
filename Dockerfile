FROM centos:7
COPY output/bin/alert /root/server
EXPOSE 8888
EXPOSE 8080
CMD /root/server > /root/log