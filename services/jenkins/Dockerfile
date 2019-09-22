FROM jenkins/jenkins:lts-alpine
USER root
RUN apk update && \
	apk add --no-cache docker openrc sudo && \
	rc-update add docker boot && \
	addgroup jenkins docker && \
	addgroup jenkins root && \
	addgroup jenkins ping && \
	echo "jenkins ALL=NOPASSWD: ALL" >> /etc/sudoers

USER jenkins
RUN /usr/local/bin/install-plugins.sh workflow-aggregator \
	job-dsl git pipeline-utility-steps \
	blueocean build-blocker-plugin
COPY config/ /usr/share/jenkins/ref/init.groovy.d/
ENV JAVA_OPTS "-Djenkins.install.runSetupWizard=false"
