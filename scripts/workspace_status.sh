#!/bin/bash

echo STABLE_GIT_COMMIT $(git rev-parse HEAD)
echo DATE $(date -R -u)
echo DOCKER_TAG $(git rev-parse --abbrev-ref HEAD)-$(git rev-parse --short=6 HEAD)
