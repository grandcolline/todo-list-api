#!/bin/sh

# Check user environment variable
if [[ -z "${FIRESTORE_PROJECT_ID}" ]]; then
  echo "Missing FIRESTORE_PROJECT_ID environment variable" >&2
  exit 1
fi

# Config gcloud project
gcloud config set project ${FIRESTORE_PROJECT_ID}

# Start emulator
gcloud beta emulators firestore start --host-port=0.0.0.0:${FIRESTORE_PORT}

