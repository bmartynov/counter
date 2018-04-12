FROM nginx:1.13.12-alpine


ENV CONSUL_TEMPLATE_VERSION=0.19.4
ENV CONSUL_TEMPLATE_RELEASES=https://releases.hashicorp.com/consul-template

RUN wget ${CONSUL_TEMPLATE_RELEASES}/${CONSUL_TEMPLATE_VERSION}/consul-template_${CONSUL_TEMPLATE_VERSION}_linux_amd64.tgz && \
    tar -xzf consul-template_${CONSUL_TEMPLATE_VERSION}_linux_amd64.tgz consul-template -C /usr/bin/

COPY conf.d/app.conf /tmp/app.conf
#TODO: add nginx -t && entripoint
CMD ["consul-template", "-consul-addr", "consul:8500", "-template", "/tmp/app.conf:/etc/nginx/nginx.conf", "-exec", "nginx"]