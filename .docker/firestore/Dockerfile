FROM google/cloud-sdk:272.0.0-alpine

# Install Java 8 & Firestore emulator
RUN apk add --update --no-cache openjdk8-jre &&\
	gcloud components install cloud-firestore-emulator beta --quiet

COPY entrypoint.sh .

ENTRYPOINT ["./entrypoint.sh"]
